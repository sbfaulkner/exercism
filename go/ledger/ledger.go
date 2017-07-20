package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

const testVersion = 4

// Entry is the data for a single ledger entry
type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

// outputData is used to communicate the resulting output back over a channel
type outputData struct {
	i int
	s string
	e error
}

// currencies contains the mapping of a currency to the appropriate currency symbol
var currencies = map[string]string{
	"EUR": "€",
	"USD": "$",
}

// localeData defines regional formats and headings
type localeData struct {
	headers            []interface{}
	dateFormat         string
	positiveFormat     string
	negativeFormat     string
	thousandsSeparator string
}

// formatCurrency formats a given number of cents using the provided currency symbol depending on the locale
func (l localeData) formatCurrency(cents int, currencySymbol string) string {
	format := l.positiveFormat

	if cents < 0 {
		cents *= -1
		format = l.negativeFormat
	}

	dollars := cents / 100

	var dollarText string

	if dollars > 999 {
		dollarParts := []string{}

		for dollars > 0 {
			dollarParts = append([]string{strconv.Itoa(dollars % 1000)}, dollarParts...)
			dollars /= 1000
		}

		dollarText = strings.Join(dollarParts, l.thousandsSeparator)
	} else {
		dollarText = strconv.Itoa(dollars)
	}

	cents %= 100

	return fmt.Sprintf(format, currencySymbol, dollarText, cents)
}

// formatDate formats a time.Time using a locale specific date format
func (l localeData) formatDate(t time.Time) string {
	return t.Format(l.dateFormat)
}

// locales is the localeData for the known regions
var locales = map[string]localeData{
	"en-US": {
		headers:            []interface{}{"Date", "Description", "Change"},
		dateFormat:         "01/02/2006",
		positiveFormat:     "%s%s.%02d ",
		negativeFormat:     "(%s%s.%02d)",
		thousandsSeparator: ",",
	},
	"nl-NL": {
		headers:            []interface{}{"Datum", "Omschrijving", "Verandering"},
		dateFormat:         "02-01-2006",
		positiveFormat:     "%s %s,%02d ",
		negativeFormat:     "%s %s,%02d-",
		thousandsSeparator: ".",
	},
}

// Errors for FormatLedger
var (
	ErrInvalidCurrency = errors.New("ledger: invalid currency")
	ErrInvalidDate     = errors.New("ledger: invalid date")
	ErrInvalidLocale   = errors.New("ledger: invalid locale")
)

// output formats for the ledger headers and line entries
const (
	headerFormat string = "%-10s | %-25s | %s\n"
	lineFormat   string = "%s | %-25s | %13s\n"
)

// byEntry provides an interface to sort the ledger entries
type byEntry []Entry

// Len returns the number of entries
func (e byEntry) Len() int {
	return len(e)
}

// Swap the two specified entries
func (e byEntry) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// Less returns true if an entry is less than the other based on Date, Description and Change in that order
func (e byEntry) Less(i, j int) bool {
	if e[i].Date != e[j].Date {
		return e[i].Date < e[j].Date
	}

	if e[i].Description != e[j].Description {
		return e[i].Description < e[j].Description
	}

	return e[i].Change < e[j].Change
}

// FormatLedger returns a string containing formatted ledger entries using a specified locale and currency
func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	currencySymbol, validCurrency := currencies[currency]
	if !validCurrency {
		return "", ErrInvalidCurrency
	}

	currentLocale, validLocale := locales[locale]
	if !validLocale {
		return "", ErrInvalidLocale
	}

	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)
	sort.Sort(byEntry(entriesCopy))

	// declare output string and add (localized) headers (ie. in either Netherlands Dutch or US English)
	s := fmt.Sprintf(headerFormat, currentLocale.headers...)

	// Parallelism, always a great idea
	co := make(chan outputData)
	for i, et := range entriesCopy {
		go formatEntry(i, et, co, currencySymbol, currentLocale)
	}

	// read from channel and insert lines in output collection at the correct index
	ss := make([]string, len(entriesCopy))
	for range entriesCopy {
		v := <-co
		if v.e != nil {
			return "", v.e
		}
		ss[v.i] = v.s
	}

	// append lines outpus string `s``
	for i := 0; i < len(entriesCopy); i++ {
		s += ss[i]
	}

	return s, nil
}

// formatEntry formats the provided entry given a currency symbol and locale data
func formatEntry(i int, entry Entry, co chan outputData, currencySymbol string, currentLocale localeData) {
	t, err := time.Parse("2006-01-02", entry.Date)
	if err != nil {
		co <- outputData{e: ErrInvalidDate}
	}

	d := currentLocale.formatDate(t)

	de := entry.Description
	if len(de) > 25 {
		de = de[:22] + "..."
	}

	a := currentLocale.formatCurrency(entry.Change, currencySymbol)

	co <- outputData{i: i, s: fmt.Sprintf(lineFormat, d, de, a)}
}

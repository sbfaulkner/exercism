package ledger

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const testVersion = 4

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type outputData struct {
	i int
	s string
	e error
}

var currencies = map[string]string{
	"EUR": "€",
	"USD": "$",
}

type localeData struct {
	locale             string
	headers            []interface{}
	dateFormat         string
	positiveFormat     string
	negativeFormat     string
	thousandsSeparator string
}

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

func (l localeData) formatDate(t time.Time) string {
	return t.Format(l.dateFormat)
}

var locales = map[string]localeData{
	"en-US": {
		locale:             "en-US",
		headers:            []interface{}{"Date", "Description", "Change"},
		dateFormat:         "01/02/2006",
		positiveFormat:     "%s%s.%02d ",
		negativeFormat:     "(%s%s.%02d)",
		thousandsSeparator: ",",
	},
	"nl-NL": {
		locale:             "nl-NL",
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

const (
	headerFormat string = "%-10s | %-25s | %s\n"
	lineFormat   string = "%s | %-25s | %13s\n"
)

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

	m1 := map[bool]int{true: 0, false: 1}
	m2 := map[bool]int{true: -1, false: 1}
	es := entriesCopy
	for len(es) > 1 {
		first, rest := es[0], es[1:]
		success := false
		for !success {
			success = true
			for i, e := range rest {
				if (m1[e.Date == first.Date]*m2[e.Date < first.Date]*4 +
					m1[e.Description == first.Description]*m2[e.Description < first.Description]*2 +
					m1[e.Change == first.Change]*m2[e.Change < first.Change]*1) < 0 {
					es[0], es[i+1] = es[i+1], es[0]
					success = false
				}
			}
		}
		es = es[1:]
	}

	// declare output string and add (localized) headers (ie. in either Netherlands Dutch or US English)
	s := fmt.Sprintf(headerFormat, currentLocale.headers...)

	// Parallelism, always a great idea
	co := make(chan outputData)
	for i, et := range entriesCopy {
		go processEntry(i, et, co, currencySymbol, currentLocale)
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

func processEntry(i int, entry Entry, co chan outputData, currencySymbol string, currentLocale localeData) {
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

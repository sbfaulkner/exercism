package ledger

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
	locale        string
	headers       []interface{}
	dateFormat    string
	dateSeparator string
}

var locales = map[string]localeData{
	"en-US": {
		locale:        "en-US",
		headers:       []interface{}{"Date", "Description", "Change"},
		dateFormat:    "03/07/2015",
		dateSeparator: "/",
	},
	"nl-NL": {
		locale:        "nl-NL",
		headers:       []interface{}{"Datum", "Omschrijving", "Verandering"},
		dateFormat:    "07-03-2015",
		dateSeparator: "-",
	},
}

// Errors for FormatLedger
var (
	ErrInvalidCurrency = errors.New("ledger: invalid currency")
	ErrInvalidDate     = errors.New("ledger: invalid date")
	ErrInvalidLocale   = errors.New("ledger: invalid locale")
)

const formatHeader string = "%-10s | %-25s | %s\n"

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
	s := fmt.Sprintf(formatHeader, currentLocale.headers...)

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
	if len(entry.Date) != 10 {
		co <- outputData{e: ErrInvalidDate}
	}
	d1, d2, d3, d4, d5 := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:10]
	if d2 != '-' {
		co <- outputData{e: ErrInvalidDate}
	}
	if d4 != '-' {
		co <- outputData{e: ErrInvalidDate}
	}
	de := entry.Description
	if len(de) > 25 {
		de = de[:22] + "..."
	} else {
		de = de + strings.Repeat(" ", 25-len(de))
	}
	var d string
	if currentLocale.locale == "nl-NL" {
		d = d5 + currentLocale.dateSeparator + d3 + currentLocale.dateSeparator + d1
	} else if currentLocale.locale == "en-US" {
		d = d3 + currentLocale.dateSeparator + d5 + currentLocale.dateSeparator + d1
	}
	negative := false
	cents := entry.Change
	if cents < 0 {
		cents = cents * -1
		negative = true
	}
	var a string
	if currentLocale.locale == "nl-NL" {
		a += currencySymbol
		a += " "
		centsStr := strconv.Itoa(cents)
		switch len(centsStr) {
		case 1:
			centsStr = "00" + centsStr
		case 2:
			centsStr = "0" + centsStr
		}
		rest := centsStr[:len(centsStr)-2]
		var parts []string
		for len(rest) > 3 {
			parts = append(parts, rest[len(rest)-3:])
			rest = rest[:len(rest)-3]
		}
		if len(rest) > 0 {
			parts = append(parts, rest)
		}
		for i := len(parts) - 1; i >= 0; i-- {
			a += parts[i] + "."
		}
		a = a[:len(a)-1]
		a += ","
		a += centsStr[len(centsStr)-2:]
		if negative {
			a += "-"
		} else {
			a += " "
		}
	} else if currentLocale.locale == "en-US" {
		if negative {
			a += "("
		}
		a += currencySymbol
		centsStr := strconv.Itoa(cents)
		switch len(centsStr) {
		case 1:
			centsStr = "00" + centsStr
		case 2:
			centsStr = "0" + centsStr
		}
		rest := centsStr[:len(centsStr)-2]
		var parts []string
		for len(rest) > 3 {
			parts = append(parts, rest[len(rest)-3:])
			rest = rest[:len(rest)-3]
		}
		if len(rest) > 0 {
			parts = append(parts, rest)
		}
		for i := len(parts) - 1; i >= 0; i-- {
			a += parts[i] + ","
		}
		a = a[:len(a)-1]
		a += "."
		a += centsStr[len(centsStr)-2:]
		if negative {
			a += ")"
		} else {
			a += " "
		}
	}
	var al int
	for range a {
		al++
	}
	co <- outputData{i: i, s: d + strings.Repeat(" ", 10-len(d)) + " | " + de + " | " +
		strings.Repeat(" ", 13-al) + a + "\n"}
}

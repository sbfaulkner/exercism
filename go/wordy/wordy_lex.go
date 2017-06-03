package wordy

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type tokenType int

const eof = -1

const (
	tokenError tokenType = iota

	tokenEOF

	tokenWord
	tokenNumber
	tokenQuestionMark
)

type token struct {
	typ tokenType
	val string
}

func (t token) String() string {
	switch t.typ {
	case tokenError:
		return t.val
	case tokenEOF:
		return "EOF"
	}

	return fmt.Sprintf("%q", t.val)
}

type stateFn func(*lexer) stateFn

type lexer struct {
	input  string
	start  int
	pos    int
	width  int
	state  stateFn
	tokens chan token
}

func lex(input string) *lexer {
	return &lexer{
		input:  input,
		state:  lexText,
		tokens: make(chan token, 2),
	}
}

func (l *lexer) getToken() token {
	for {
		select {
		case t := <-l.tokens:
			return t
		default:
			l.state = l.state(l)
		}
	}
}

func (l *lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.tokens)
}

func (l *lexer) emit(t tokenType) {
	l.tokens <- token{t, strings.ToLower(l.input[l.start:l.pos])}
	l.start = l.pos
}

func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

func (l *lexer) acceptRun(valid string) {
	for l.accept(valid) {
	}
}

func (l *lexer) acceptRangeTable(rangeTab *unicode.RangeTable) {
	for unicode.Is(rangeTab, l.next()) {
	}
	l.backup()
}

func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func lexText(l *lexer) stateFn {
	for {
		switch r := l.next(); {
		case r == eof:
			l.emit(tokenEOF)
			return nil
		case unicode.IsSpace(r):
			l.ignore()
		case unicode.IsLetter(r):
			l.backup()
			return lexWord
		case r == '+' || r == '-' || unicode.IsDigit(r):
			l.backup()
			return lexNumber
		case r == '?':
			l.backup()
			return lexQuestionMark
		}
	}
}

func lexWord(l *lexer) stateFn {
	l.acceptRangeTable(unicode.Letter)
	l.emit(tokenWord)
	return lexText
}

func lexNumber(l *lexer) stateFn {
	l.accept("+-")
	l.acceptRangeTable(unicode.Digit)
	l.emit(tokenNumber)
	return lexText
}

func lexQuestionMark(l *lexer) stateFn {
	l.pos++
	l.emit(tokenQuestionMark)
	return lexText
}

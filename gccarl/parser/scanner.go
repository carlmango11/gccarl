package parser

import (
	"fmt"
	"io"
	"strconv"
	"unicode"
)

type Scanner struct {
	text string
	i    int
}

func (s *Scanner) String() string {
	return s.text[:s.i] + ">" + s.text[s.i:]
}

func (s *Scanner) ReadLiteral(l string) (bool, error) {
	s.skipWhitespace()

	val, err := s.read(len(l))
	if err != nil {
		return false, err
	}

	return val == l, nil
}

var whitespaceChars = map[string]bool{
	" ":  true,
	"\n": true,
	"\r": true,
	"\t": true,
}

func (s *Scanner) skipWhitespace() {
	if s.i >= len(s.text) {
		return
	}

	for {
		ch := s.text[s.i : s.i+1]
		if !whitespaceChars[ch] {
			return
		}

		s.i++
	}
}

func (s *Scanner) Index() int {
	return s.i
}

func (s *Scanner) ParseNumber() (int64, error) {
	s.skipWhitespace()

	startCh, err := s.next()
	if err != nil {
		return 0, err
	}

	if !unicode.IsNumber(startCh) {
		return 0, fmt.Errorf("expected a number and got %q", startCh)
	}

	//var decimal bool
	var chs []rune

	for {
		ch, err := s.next()
		if err != nil {
			break
		}

		//if ch == '.' {
		//	if decimal {
		//		return 0, fmt.Errorf("more than 1 decimal point in number")
		//	}
		//
		//	decimal = true
		//}

		if !unicode.IsNumber(ch) {
			s.i--
			break
		}

		chs = append(chs, ch)
	}

	all := make([]rune, 1, len(chs)+1)
	all[0] = startCh
	all = append(all, chs...)

	return strconv.ParseInt(string(all), 10, 64)
}

func (s *Scanner) ParseChar() (byte, error) {
	ch, err := s.next()
	if err != nil {
		return 0, err
	}

	return string(ch)[0], nil
}

func (s *Scanner) ParseIdentifier() (Identifier, error) {
	s.skipWhitespace()

	startCh, err := s.next()
	if err != nil {
		return "", err
	}

	if !unicode.IsLetter(startCh) {
		s.i--
		return "", fmt.Errorf("identifier does not start with a letter (%q)", startCh)
	}

	var chs []rune
	for {
		ch, err := s.next()
		if err != nil {
			break
		}

		if !validIdenChar(ch) {
			s.i--
			break
		}

		chs = append(chs, ch)
	}

	return runesToIden(startCh, chs)
}

func validIdenChar(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_'
}

func runesToIden(start rune, rest []rune) (Identifier, error) {
	all := make([]rune, 1, len(rest)+1)
	all[0] = start
	all = append(all, rest...)

	return Identifier(all), nil
}

func (s *Scanner) next() (rune, error) {
	if s.i+1 > len(s.text) {
		return 0, io.EOF
	}

	r := []rune(s.text)[s.i]
	s.i++

	return r, nil
}

func (s *Scanner) read(n int) (string, error) {
	if s.i+n > len(s.text) {
		return "", io.EOF
	}

	i := s.i
	s.i += n

	return s.text[i:s.i], nil
}

func (s *Scanner) Finish() bool {
	s.skipWhitespace()

	return s.i == len(s.text)
}

func (s *Scanner) Reset(index int) {
	s.i = index
}

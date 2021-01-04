package boot

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"strings"
	"github.com/mitchellh/mapstructure"
)

// PlainParser is a parser for config files in an extremely simple format. Each
// line is tokenized as a single key/value pair. The first whitespace-delimited
// token in the line is interpreted as the flag name, and all remaining tokens
// are interpreted as the value. Any leading hyphens on the flag name are
// ignored.
func PlainParser(r io.Reader, set func(name, value string) error) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue // skip empties
		}

		if line[0] == '#' {
			continue // skip comments
		}

		var (
			name  string
			value string
			index = strings.IndexRune(line, ' ')
		)
		if index < 0 {
			name, value = line, "true" // boolean option
		} else {
			name, value = line[:index], strings.TrimSpace(line[index:])
		}

		if i := strings.Index(value, " #"); i >= 0 {
			value = strings.TrimSpace(value[:i])
		}

		if err := set(name, value); err != nil {
			return err
		}
	}
	return nil
}

func DotEnvParser(r io.Reader, set func(name, value string) error) error  {
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue // skip empties
		}

		if line[0] == '#' {
			continue // skip comments
		}

		var (
			name  string
			value string
			index = strings.IndexRune(line, '=')
		)
		if index < 0 {
			name, value = line, "true" // boolean option
		} else {
			name, value = line[:index], strings.TrimSpace(line[index+1:])
		}

		if i := strings.Index(value, " #"); i >= 0 {
			value = strings.TrimSpace(value[:i])
		}
		name = strings.ToLower(strings.ReplaceAll(name, "_", "-"))
		if err := set(name, value); err != nil {
			return err
		}
	}
	return nil
}

func configParser(r io.Reader, set func(name, value string) error) (err error) {
	var (
		m  = make(map[string]interface{})
		d = yaml.NewDecoder(r)
	)

	if err = mapstructure.Decode(r, d); err != nil && err != io.EOF {
		return ParseError{err}
	}

	fmt.Println(m)

	return nil
}

// ParseError wraps all errors originating from the YAML parser.
type ParseError struct {
	Inner error
}

// Error implenents the error interface.
func (e ParseError) Error() string {
	return fmt.Sprintf("error parsing YAML config: %v", e.Inner)
}

// Unwrap implements the errors.Wrapper interface, allowing errors.Is and
// errors.As to work with ParseErrors.
func (e ParseError) Unwrap() error {
	return e.Inner
}
package rcscan

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type RCfile struct {
	name   string
	values map[string]map[string]string
}

const (
	ERR_PREFIX      = "rcscan error, "
	DEFAULT_SECTION = "[__default__]"
)

func e(msg string) error {
	return errors.New(ERR_PREFIX + msg)
}

// Creates new rcfile object for filename specified
func New(filename string) (*RCfile, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	rc := &RCfile{name: filename, values: make(map[string]map[string]string)}
	section := DEFAULT_SECTION
	rc.values[DEFAULT_SECTION] = make(map[string]string)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || line[0] == ';' || line[0] == '#' || line[0] == '/' {
			continue
		}

		if line[0] == '[' {
			if line[len(line)-1] == ']' {
				section = line
				if _, ok := rc.values[line]; !ok {
					rc.values[line] = make(map[string]string)
				}
				continue
			} else {
				return nil, e("bad syntax, no ']' closure found at line: " + line)
			}
		}

		param := strings.SplitAfterN(line, "=", 2)
		if len(param) != 2 {
			return nil, e("bad syntax, no '=' found at line: " + line)
		}

		param[0] = strings.TrimSpace(strings.Trim(param[0], "="))
		param[1] = strings.TrimSpace(param[1])

		rc.values[section][param[0]] = param[1]
	}

	return rc, nil
}

// Returns parameter's value from the section specified. Section name can be specified either in '[]' or without.
func (r *RCfile) Get(section string, param string) (string, error) {

	if section == "" {
		section = DEFAULT_SECTION
	}

	if section[0] != '[' {
		section = "[" + section + "]"
	}

	if value, ok := r.values[section][param]; ok {
		return value, nil
	}

	return "", e("section or parameter not defined")
}

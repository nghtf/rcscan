package rcscan

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type RCScanner struct {
	File string
}

// Creates new scanner for rcfile specified. Returns error if target rcfile doesn't exist.
func New(file string) (*RCScanner, error) {
	var r RCScanner
	if _, err := os.Stat(file); err == nil {
		r.File = file
		return &r, nil
	} else {
		return nil, errors.New("file doesn't exist: " + file)
	}
}

// Returns parameter's value as string from the section specified
// Example 1: GetParam("[section]", "parameterA")
// Example 2: GetParam("section", "parameterB")
func (r *RCScanner) GetParam(section string, param string) (string, error) {

	logPrefix := "rcscan.GetParam(): "

	if (len(section) == 0) || (len(param) == 0) {
		return "", errors.New(logPrefix + "args must not be empty")
	}

	file, err := os.Open(r.File)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)

	if section[0] != '[' {
		section = "[" + section + "]"
	}

	inside := false
	value := ""

	for scanner.Scan() {
		str := scanner.Text()
		if str == section {
			inside = true
			continue
		}
		if inside {
			if strings.HasPrefix(str, param) {
				// Clean up value
				value = strings.TrimPrefix(str, param)
				value = strings.TrimSpace(value)
				value = strings.TrimPrefix(value, "=")
				value = strings.TrimSpace(value)
				return value, nil
			}
			if isSection(str) {
				break
			}
		}

	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", errors.New(logPrefix + param + " - param not found")
}

func isSection(str string) bool {
	if len(str) > 0 {
		if str[0] == '[' {
			return true
		}
	}
	return false
}

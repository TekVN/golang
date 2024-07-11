package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsNumeric(t *testing.T) {
	testcase := map[string]bool{
		"1":  true,
		"2":  true,
		"3":  true,
		"4":  true,
		"5":  true,
		"6":  true,
		"7":  true,
		"8":  true,
		"9":  true,
		"0":  true,
		"11": true,
		"a":  false,
		"b":  false,
		"c":  false,
		"d":  false,
		"e":  false,
		"f":  false,
		"111111111111111111111111111111111111111": true,
	}

	for key, val := range testcase {
		assert.Equal(t, val, IsNumeric(key))
	}
}

func TestIsInt(t *testing.T) {
	testcase := map[string]bool{
		"1":                     true,
		"2":                     true,
		"9":                     true,
		"0":                     true,
		"11":                    true,
		"a":                     false,
		"b":                     false,
		"c":                     false,
		"d":                     false,
		"1e":                    false,
		"e":                     false,
		"f":                     false,
		"1f":                    false,
		"111111111111111111111": false,
		"1.1":                   false,
		"1,2":                   false,
	}

	for key, val := range testcase {
		assert.Equal(t, val, IsInt(key), key)
	}
}

func TestIsFloat(t *testing.T) {
	testcase := map[string]bool{
		"1":                       false,
		"2":                       false,
		"9":                       false,
		"0":                       false,
		"11":                      false,
		"a":                       false,
		"b":                       false,
		"c":                       false,
		"d":                       false,
		"1e":                      false,
		"e":                       false,
		"f":                       false,
		"1f":                      false,
		"1111111.11.111111111111": false,
		"1.1":                     true,
		"1,2":                     false,
	}

	for key, val := range testcase {
		assert.Equal(t, val, IsFloat(key), key)
	}
}

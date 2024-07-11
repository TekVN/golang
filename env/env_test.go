package env

import (
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"strconv"
	"testing"
)

func TestGetEnv(t *testing.T) {
	testcase := map[string]string{
		"KEY1": "VAL1",
		"Key2": "VAL2",
		"keY3": "val3",
		"key4": "val4",
		"kEy5": "vAl5",
	}

	for key, val := range testcase {
		assert.NoError(t, os.Setenv(key, val))
		assert.Equal(t, val, GetEnv(key))
		assert.NotEmpty(t, val)
	}
	assert.Empty(t, GetEnv("KEY_NOT_EXISTS"))
}

func TestGetEnvWithDefault(t *testing.T) {
	assert.Equal(t, "default", GetEnvWithDefault("KEY_NOT_EXISTS", "default"))
}

func TestGetEnvInt(t *testing.T) {
	testcase := map[string]any{
		"KEY1": 1,
		"Key2": "a",
		"keY3": "3",
		"key4": "0",
		"kEy5": "",
		"kEY6": "a2",
	}
	for key, val := range testcase {
		if reflect.ValueOf(val).Kind() == reflect.String {
			assert.NoError(t, os.Setenv(key, val.(string)))
		} else {
			assert.NoError(t, os.Setenv(key, strconv.Itoa(val.(int))))
		}
		num, ok := GetEnvInt(key)

		if reflect.ValueOf(val).Kind() == reflect.String {
			if val.(string) == "" {
				assert.False(t, ok)
				assert.Equal(t, 0, num)
			} else {
				expect, err := strconv.Atoi(val.(string))
				if err != nil {
					assert.False(t, ok)
					assert.Equal(t, 0, num)
				} else {
					assert.True(t, ok)
					assert.Equal(t, expect, num)
				}
			}
		} else {
			assert.True(t, ok)
			assert.Equal(t, val.(int), num)
		}
	}
	num, ok := GetEnvInt("KEY_NOT_EXISTS")
	assert.False(t, ok)
	assert.Equal(t, 0, num)
}

func TestGetEnvIntWithDefault(t *testing.T) {
	assert.Equal(t, 1, GetEnvIntWithDefault("KEY_NOT_EXISTS", 1))
}

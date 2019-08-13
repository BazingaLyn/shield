package config

import "regexp"

type Config struct {
	conf         map[interface{}]interface{}
	placeHolders map[string]interface{}
	rex          *regexp.Regexp
}

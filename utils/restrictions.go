package utils

import (
	"strings"

	"github.com/chinmayagrawal775/forward_proxy/config"
)

func IsRestrictedHost(host string) bool {
	for _, rHost := range config.RestrictedHosts {
		if rHost == host {
			return true
		}
	}
	return false
}

func IsRestrictedWord(body string) bool {
	for _, word := range config.RestrictedWords {
		if strings.Contains(body, word) {
			return true
		}
	}
	return false
}

package validators

import (
	"strings"
	"unicode"
)

func isLocalSymbol(r rune) bool {
	return r == '.' || r == '-' || r == '_'
}

func isDomainSymbol(r rune) bool {
	return r == '.' || r == '-'
}

func IsValidEmail(email string) bool {
	atIndex := strings.IndexByte(email, '@')
	if atIndex < 1 || atIndex > len(email)-4 {
		return false // Invalid @ index position
	}

	local := email[:atIndex]
	domain := email[atIndex+1:]
	if len(local) > 64 || len(domain) > 255 {
		return false // Invalid local or domain length
	}

	if isDomainSymbol(rune(domain[0])) || isDomainSymbol(rune(domain[len(domain)-1])) {
		return false // Invalid domain symbol position
	} else if isLocalSymbol(rune(local[0])) || isLocalSymbol(rune(local[len(local)-1])) {
		return false // Invalid local symbol position
	}

	var localCache rune
	for _, r := range local {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			if !isLocalSymbol(r) || isLocalSymbol(localCache) {
				return false // Invalid local part structure
			}
		}
		localCache = r
	}

	var domainCache rune
	for i, r := range domain {
		if !(r|32 >= 'a' && r|32 <= 'z') && !(r <= '9' && r >= '0') {
			if !isDomainSymbol(r) || (r == '.' && (isDomainSymbol(domainCache) || isDomainSymbol(rune(domain[i+1])))) {
				return false // Invalid domain part structure
			}
		}
		domainCache = r
	}

	return strings.IndexByte(domain, '.') > -1
}

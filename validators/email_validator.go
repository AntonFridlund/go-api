package validators

import "strings"

const (
	localSymbols     = "-._"
	domainSymbols    = "-."
	extensionSymbols = ""
)

func isASCIILetter(r byte) bool {
	return r|32 >= 'a' && r|32 <= 'z'
}

func isASCIIDigit(r byte) bool {
	return r >= '0' && r <= '9'
}

func isValidLocalPart(local string) bool {
	for i := 0; i < len(local); i++ {
		if !isASCIILetter(local[i]) && !isASCIIDigit(local[i]) {
			if strings.IndexByte(localSymbols, local[i]) < 1 || i == len(local)-1 {
				return false // Invalid symbol or position
			} else if strings.IndexByte(localSymbols, local[i-1]) != -1 {
				return false // Invalid consecutive symbols
			}
		}
	}
	return true
}

func isValidDomainPart(domain string) bool {
	for i := 0; i < len(domain); i++ {
		if !isASCIILetter(domain[i]) && !isASCIIDigit(domain[i]) {
			if strings.IndexByte(domainSymbols, domain[i]) < 1 || i == len(domain)-1 {
				return false // Invalid symbol or position
			} else if strings.IndexByte(domainSymbols, domain[i-1]) != -1 {
				return !(domain[i] != '-' || domain[i-1] != '-' || domain[i-2] == '-')
			}
		}
	}
	return true
}

func isValidExtensionPart(extension string) bool {
	for i := 0; i < len(extension); i++ {
		if !isASCIILetter(extension[i]) {
			return false // Invalid extension part structure
		}
	}
	return true
}

// Email validation logic
func IsValidEmail(email string) bool {
	atIndex := strings.IndexByte(email, '@')
	if atIndex < 1 || atIndex > len(email)-4 {
		return false // Invalid @ position
	}

	lastDotIndex := strings.LastIndexByte(email, '.')
	if lastDotIndex < 3 || lastDotIndex > len(email)-2 {
		return false // Invalid last dot position
	}

	local := email[:atIndex]
	domain := email[atIndex+1 : lastDotIndex]
	extension := email[lastDotIndex+1:]

	if len(local) > 64 || len(domain) > 253 || len(extension) > 63 {
		return false // Invalid local, domain or extension length
	}

	return isValidLocalPart(local) && isValidDomainPart(domain) && isValidExtensionPart(extension)
}

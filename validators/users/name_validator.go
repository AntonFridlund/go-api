package users

import "unicode"

func IsValidFirstName(name string) bool {
	return validateName(name)
}

func IsValidLastName(name string) bool {
	return validateName(name)
}

func validateName(name string) bool {
	if len(name) < 1 || len(name) > 80 || name[0] == '-' || name[len(name)-1] == '-' {
		return false // Invalid name format
	}
	for i, r := range name {
		if !unicode.IsLetter(r) && r != '-' {
			return false // Only letters and hyphens allowed
		}
		if i > 0 && name[i] == '-' && name[i-1] == '-' {
			return false // Consecutive hyphens are not allowed
		}
	}
	return true
}

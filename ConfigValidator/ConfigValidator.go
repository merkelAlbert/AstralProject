package ConfigValidator

import ()

func ValidateAge(a uint) bool {
	if (a > 0) && (a < 120) {
		return true
	}
	return false
}

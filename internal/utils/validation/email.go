package validation

import "strings"

func IsValidEmail(email string) bool {
	strings.Split(email, "@")
	return len(strings.Split(email, "@")) == 2
}

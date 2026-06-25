package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateEmaill(t *testing.T) {
	type test struct {
		title    string
		email    string
		expected bool
	}

	tests := []test{
		{title: "Valild Email", email: "john.doe@gmail.com", expected: true},
		{title: "Invalild Email", email: "johndoe.com", expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			got := IsValidEmail(tt.email)
			require.Equal(t, tt.expected, got)
		})
	}
}

package encryption

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateAndValidate(t *testing.T) {
	plain := "admin123"
	hashed, err := GenerateFromPassword(plain)
	log.Println(hashed)
	require.Nil(t, err)
	require.NotEmpty(t, hashed)

	err = ValidatePassword(hashed, plain)
	require.Nil(t, err)
}

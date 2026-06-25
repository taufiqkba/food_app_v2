package generator

import "github.com/oklog/ulid/v2"

func GeneratePublicID() string {
	id := ulid.Make()
	return id.String()
}

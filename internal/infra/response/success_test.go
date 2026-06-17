package response

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSuccess(t *testing.T) {
	payload := map[string]interface{}{
		"foo": "baar",
	}
	resp := NewSuccessCreated("test",
		WithPayload(payload),
		WithStatusCode("201010100"),
	)
	resJson, _ := json.Marshal(resp)
	fmt.Printf("%+v\n", string(resJson))

	require.NotEmpty(t, resp)
	require.Equal(t, payload, resp.Payload)
	require.Equal(t, "201010100", resp.StatusCode)
}

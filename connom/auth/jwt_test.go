package auth

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestJwt(t *testing.T) {
	payload := LoginPayload{
		UID: 1024,
	}
	token, err := GenToken(payload)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("token", token)
	loginPayload, err := Parse(token)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("uid", loginPayload.UID)
	assert.Equal(t, int64(1024), loginPayload.UID)
}

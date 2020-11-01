package auth_test

import (
	"github.com/gira-games/client/pkg/client"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/asankov/gira/internal/auth"
	gassert "github.com/asankov/gira/internal/fixtures/assert"
)

var (
	expectedUser     = &client.User{Username: expectedUsername}
	expectedUsername = "username"

	authenticator = auth.NewAutheniticator("secret")
)

func TestToken(t *testing.T) {
	token, err := authenticator.NewTokenForUser(expectedUser)
	require.Nil(t, err)

	usr, err := authenticator.DecodeToken(token)

	require.Nil(t, err)
	assert.Equal(t, expectedUser.Username, usr.Username)
}

func TestDecodeTokenError(t *testing.T) {
	usr, err := authenticator.DecodeToken("o.o")

	assert.Nil(t, usr)
	gassert.Error(t, err, auth.ErrInvalidFormat)
}

func TestTokenExpired(t *testing.T) {
	token, err := authenticator.NewTokenForUserWithExpiration(expectedUser, 1*time.Millisecond)
	require.NoError(t, err)

	// wait for the token to expire
	time.Sleep(1 * time.Second)

	usr, err := authenticator.DecodeToken(token)
	assert.Nil(t, usr)
	gassert.Error(t, err, auth.ErrTokenExpired)
}

func TestInvalidSignature(t *testing.T) {
	newAuthenticator := auth.NewAutheniticator("some.other.secret")
	newToken, err := newAuthenticator.NewTokenForUser(&client.User{})
	assert.Nil(t, err)

	usr, err := authenticator.DecodeToken(newToken)
	assert.Nil(t, usr)
	gassert.Error(t, err, auth.ErrInvalidSignature)
}

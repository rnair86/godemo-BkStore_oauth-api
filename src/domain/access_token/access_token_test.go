package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstant(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "Expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.NotNil(t, at, "brand new access token should not be nil")

	assert.False(t, at.IsExpired(), "brand new access token should not be expired")

	assert.EqualValues(t, "", at.AccessToken, "brand new access token should not have defined access token id")

	assert.True(t, at.UserId == 0, "brand new access token should not have an associated user id")

	assert.True(t, at.ClientId == 0, "brand new access token should not have an associated client id")

}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring 3 hours from now should not be expired")

}

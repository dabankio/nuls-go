package nuls

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GetAccount(t *testing.T) {
	info, err := apiTest.GetAccount(testAccount)
	require.NoError(t, err, "api.GetAccount")
	t.Log(info)
}

func TestClient_CreateAccount(t *testing.T) {
	info, err := apiTest.CreateAccount(3, testPassword)
	require.NoError(t, err, "api.GetAccount")
	t.Log(info)
}

func TestClient_Balance(t *testing.T) {
	info, err := apiTest.Balance()
	require.NoError(t, err, "api.Balance")
	t.Logf("%+v", info)
}

func TestClient_ValidateAddress(t *testing.T) {
	info, err := apiTest.ValidateAddress(testAccount)
	require.NoError(t, err, "api.validateAddress")
	t.Logf("%+v", info)
}

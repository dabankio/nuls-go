package nuls

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_ValidateContractAddress(t *testing.T) {
	info, err := apiTest.ValidateContractAddress(testAccount)
	require.NoError(t, err, "api.ValidateContractAddress")
	t.Log(info)
}

package nuls

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GetTransaction(t *testing.T) {
	info, err := apiTest.GetTransaction("00205254e6716199effe6516f0369f22e4a4f97b108f103a1d764493f601b4e478f4")
	require.NoError(t, err, "api.GetTransaction")
	t.Logf("%+v", info)
}

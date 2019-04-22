package nuls

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_Transfer(t *testing.T) {
	info, err := apiTest.Transfer(testAccount, "TTamZRpu8ksoUFwHbsF3cuCAu8Njo8QG", "", AmountToInt64(4000.998))
	require.NoError(t, err, "api.Transfer")
	t.Logf("%+v", info)
}

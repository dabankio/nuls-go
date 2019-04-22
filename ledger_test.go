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

func TestClient_MultipleAddressTransfer(t *testing.T) {
	inputs := []Inputs{{
		Address:  "TTamZRpu8ksoUFwHbsF3cuCAu8Njo8QG",
		Password: "",
	}, {
		Address:  "TTaseiEdKMrZi6SBLx14267CQtpo65GY",
		Password: "",
	}}
	outPut := []OutPuts{{
		ToAddress: "TTawEGqAXHyQwJL6pcw4n3boHwhcUZRS",
		Amount:    AmountToInt64(10001),
	}}
	info, err := apiTest.MultipleAddressTransfer(inputs, outPut)
	require.NoError(t, err, "api.Transfer")
	t.Logf("%+v", info)
}

func TestClient_AddressBalance(t *testing.T) {
	info, err := apiTest.AddressBalance("TTamZRpu8ksoUFwHbsF3cuCAu8Njo8QG")
	require.NoError(t, err, "api.AddressBalance")
	t.Logf("%+v", info)
}

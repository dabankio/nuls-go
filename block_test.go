package nuls

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_Block(t *testing.T) {
	info, err := apiTest.Block(1337967, 1)
	require.NoError(t, err, "api.block")
	for i := range info.Data.List {
		for j := range info.Data.List[i].TxList {
			if info.Data.List[i].TxList[j].Type != 2 {
				continue
			}
			t.Logf("%+v", info.Data.List[i].TxList[j].Outputs)
		}
	}
	// t.Logf("%+v", info.Data.List)
}

func TestClient_BlockHeight(t *testing.T) {
	info, err := apiTest.BlockHeight()
	require.NoError(t, err, "api.blockHeight")
	t.Logf("%+v", info)
}

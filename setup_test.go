package nuls

// apiTest connects to nuls testnet node
var apiTest *Client

const (
	testAccount  = "TTawEGqAXHyQwJL6pcw4n3boHwhcUZRS"
	testPassword = "123456789qqq"
	testHost     = ""
)

func init() {
	apiTest = New(testHost)
	apiTest.Verbose = true
	apiTest.BeforeRequest = func(_ string, _ string, _ interface{}) error {
		return nil
	}
}

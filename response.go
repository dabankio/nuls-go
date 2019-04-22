package nuls

// Response shows up when there is something wrong.
type ErrorResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Code string `json:"code"`
		Msg  string `json:"msg"`
	} `json:"data"`
}

type AccountInfo struct {
	Success bool `json:"success"`
	Data    struct {
		PageNumber int `json:"pageNumber"`
		PageSize   int `json:"pageSize"`
		Total      int `json:"total"`
		Pages      int `json:"pages"`
		List       []struct {
			Address    string      `json:"address"`
			Alias      string      `json:"alias"`
			PubKey     string      `json:"pubKey"`
			Extend     interface{} `json:"extend"`
			CreateTime int64       `json:"createTime"`
			Encrypted  bool        `json:"encrypted"`
			Remark     string      `json:"remark"`
			Ok         bool        `json:"ok"`
		} `json:"list"`
	} `json:"data"`
}

type AccountBalance struct {
	Success bool `json:"success"`
	Data    struct {
		Balance struct {
			Value int64 `json:"value"`
		} `json:"balance"`
		Locked struct {
			Value int `json:"value"`
		} `json:"locked"`
		Usable struct {
			Value int64 `json:"value"`
		} `json:"usable"`
	} `json:"data"`
}

type AccountList struct {
	Success bool `json:"success"`
	Data    struct {
		List []string `json:"list"`
	} `json:"data"`
}

type ValidateAddress struct {
	Success bool `json:"success"`
	Data    struct {
		Value bool `json:"value"`
	} `json:"data"`
}

type transaction struct {
	Hash         string `json:"hash"`
	Type         int    `json:"type"`
	Time         int64  `json:"time"`
	BlockHeight  int    `json:"blockHeight"`
	Fee          int    `json:"fee"`
	Value        int64  `json:"value"`
	Remark       string `json:"remark"`
	ScriptSig    string `json:"scriptSig"`
	Status       int    `json:"status"`
	ConfirmCount int    `json:"confirmCount"`
	Size         int    `json:"size"`
	Inputs       []struct {
		FromHash  string `json:"fromHash"`
		FromIndex int    `json:"fromIndex"`
		Address   string `json:"address"`
		Value     int    `json:"value"`
	} `json:"inputs"`
	Outputs []struct {
		TxHash   *string `json:"txHash"`
		Index    *int    `json:"index"`
		Address  string  `json:"address"`
		Value    int64   `json:"value"`
		LockTime int     `json:"lockTime"`
		Status   *int    `json:"status"`
	} `json:"outputs"`
}

type Transaction struct {
	Success bool        `json:"success"`
	Data    transaction `json:"data"`
}

type BlockInfo struct {
	Success bool `json:"success"`
	Data    struct {
		List []struct {
			Hash                 string        `json:"hash"`
			PreHash              string        `json:"preHash"`
			MerkleHash           string        `json:"merkleHash"`
			StateRoot            string        `json:"stateRoot"`
			Time                 int64         `json:"time"`
			Height               int           `json:"height"`
			TxCount              int           `json:"txCount"`
			PackingAddress       string        `json:"packingAddress"`
			Extend               string        `json:"extend"`
			RoundIndex           int           `json:"roundIndex"`
			ConsensusMemberCount int           `json:"consensusMemberCount"`
			RoundStartTime       int64         `json:"roundStartTime"`
			PackingIndexOfRound  int           `json:"packingIndexOfRound"`
			Reward               int           `json:"reward"`
			Fee                  int           `json:"fee"`
			ConfirmCount         int           `json:"confirmCount"`
			Size                 int           `json:"size"`
			TxList               []transaction `json:"txList"`
			ScriptSig            string        `json:"scriptSig"`
		} `json:"list"`
	} `json:"data"`
}

type BlockHeight struct {
	Success bool `json:"success"`
	Data    struct {
		Value int `json:"value"`
	} `json:"data"`
}

type Transfer struct {
	Success bool `json:"success"`
	Data    struct {
		Value string `json:"value"`
	} `json:"data"`
}

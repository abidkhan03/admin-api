package response

type PhrasePattern struct {
	Tokens []TokenInfo `json:"tokens"`
}

type TokenInfo struct {
	SeqId   int      `json:"seq_id"`
	Word    string   `json:"word"`
	Pos     string   `json:"pos"`
	Classes []string `json:"classes"`
}

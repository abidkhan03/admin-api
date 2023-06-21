package freeling

type Response struct {
	Sentences []Sentence `json:"sentences"`
}

type Sentence struct {
	ID     int
	IdStr  string  `json:"id"`
	Tokens []Token `json:"tokens"`
}

type Token struct {
	ID       int
	IdStr    string `json:"id"`
	Begin    int
	BeginStr string `json:"begin"`
	End      int
	EndStr   string `json:"end"`
	Form     string `json:"form"`
	Lemma    string `json:"lemma"`
	Tag      string `json:"tag"`
	CTag     string `json:"ctag"`
	POS      string `json:"pos"`
	Type     string `json:"type"`
}

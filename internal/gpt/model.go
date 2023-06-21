package gpt

type Response struct {
	Sentences []Sentence `json:"sentences"`
}

type Sentence struct {
	SeqID   int      `json:"seq_id"`
	Begin   int      `json:"begin"`
	End     int      `json:"end"`
	Phrases []Phrase `json:"phrases"`
}

type Phrase struct {
	SeqID int    `json:"seq_id"`
	Begin int    `json:"begin"`
	End   int    `json:"end"`
	Words []Word `json:"words"`
}

type Word struct {
	SeqID int    `json:"seq_id"`
	Word  string `json:"word"`
	Lemma string `json:"lemma"`
	Tag   string `json:"tag"`
}

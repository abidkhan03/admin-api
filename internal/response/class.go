package response

type Class struct {
	Id          uint64   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Words       []string `json:"words"`
}

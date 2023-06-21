package response

type Response struct {
	Id      uint64 `json:"id,omitempty"`
	Status  uint64 `json:"status"`
	Message string `json:"message"`
}

package response

type StepResponse struct {
	ID       uint   `json:"id"`
	Content  string `json:"content"`
	Sequence uint   `json:"sequence"`
}

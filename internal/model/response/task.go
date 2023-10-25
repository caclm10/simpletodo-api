package response

type TaskResponse struct {
	ID       uint   `json:"id"`
	Content  string `json:"content"`
	Sequence uint   `json:"sequence"`
}

type TaskWithStepsResponse struct {
	ID       uint           `json:"id"`
	Content  string         `json:"content"`
	Sequence uint           `json:"sequence"`
	Steps    []StepResponse `json:"steps"`
}

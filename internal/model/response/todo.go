package response

type TodoResponse struct {
	ID       uint                    `json:"id"`
	Name     string                  `json:"name"`
	Sequence uint                    `json:"sequence"`
	Tasks    []TaskWithStepsResponse `json:"tasks"`
}

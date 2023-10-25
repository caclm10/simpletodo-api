package request

type StoreStepRequest struct {
	Content string `json:"name" validate:"required"`
}

type UpdateStepRequest struct {
	Content string `json:"name" validate:"required"`
}

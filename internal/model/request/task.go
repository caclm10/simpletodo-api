package request

type StoreTaskRequest struct {
	Content string `json:"name" validate:"required"`
}

type UpdateTaskRequest struct {
	Content string `json:"name" validate:"required"`
}

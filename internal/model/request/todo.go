package request

type StoreTodoRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateTodoRequest struct {
	Name string `json:"name" validate:"required"`
}

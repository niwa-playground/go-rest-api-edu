package request

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required, min=2, max=200"`
}

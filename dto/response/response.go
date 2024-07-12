package response

type Body struct {
	Message string            `json:"message"`
	Data    any               `json:"data,omitempty"`
	Errors  map[string]string `json:"pagination,omitempty"`
}

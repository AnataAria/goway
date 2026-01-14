package response

type Response[T any] struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

type PaginationMeta struct {
	CurrentPage int64 `json:"current_page"`
	TotalPage   int64 `json:"total_page"`
	TotalItems  int64 `json:"total_items"`
	Limit       int64 `json:"limit"`
}

type PaginatedResponse[T any] struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    []T            `json:"data"`
	Meta    PaginationMeta `json:"meta"`
}

type PaginationRequest struct {
	Page  int64 `json:"page" validate:"min=1"`
	Limit int64 `json:"limit" validate:"min=1,max=100"`
}

func NewPaginatedResponse[T any](data []T, totalItems, page, limit int64, message string) *PaginatedResponse[T] {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	totalPages := (totalItems + limit - 1) / limit

	if data == nil {
		data = make([]T, 0)
	}

	return &PaginatedResponse[T]{
		Code:    "SUCCESS",
		Message: message,
		Data:    data,
		Meta: PaginationMeta{
			CurrentPage: page,
			TotalPage:   totalPages,
			TotalItems:  totalItems,
			Limit:       limit,
		},
	}
}

func NewSuccessResponse[T any](data T, message string) *Response[T] {
	return &Response[T]{
		Code:    "SUCCESS",
		Message: message,
		Data:    data,
	}
}

func RespondOK[T any](data T) *Response[T] {
	return NewSuccessResponse(data, "Success")
}

func RespondCreated[T any](data T) *Response[T] {
	return NewSuccessResponse(data, "Resource created successfully")
}

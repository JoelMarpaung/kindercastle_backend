package payload

//go:generate easytags $GOFILE json

const (
	maxLimit = 60
)

type PagingAndFilterPayload struct {
	Search  string `query:"search" json:"search" example:"keyword"`
	Limit   int    `query:"limit" json:"limit" example:"1"`
	Page    int    `query:"page" json:"page" example:"1"`
	Usecase string `query:"usecase" json:"usecase" example:"history"`
}

func (params *PagingAndFilterPayload) Normalize() {
	if params.Limit > maxLimit || params.Limit == 0 {
		params.Limit = maxLimit
	}

	if params.Page <= 0 {
		params.Page = 1
	}
}

func (params *PagingAndFilterPayload) GetOffset() int {
	offset := (params.Page - 1) * params.Limit
	if offset < 0 {
		offset = 0
	}

	return offset
}

type ResponseData[T any] struct {
	Data T `json:"data"`
}

type ResponseList[T any] struct {
	Items []T `json:"items"`
}

type PaginatedResponse[T any] struct {
	Page      int    `json:"page" example:"1"`
	Limit     int    `json:"limit" example:"60"`
	Search    string `json:"search" example:"keyword"`
	TotalRow  int    `json:"total_row" example:"1"`
	TotalPage int    `json:"total_page" example:"1"`
	Items     []T    `json:"items"`
}

type ResponseError[T any] struct {
	Error T `json:"error"`
}

type FieldError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func (fe FieldError) Error() string {
	return fe.Reason
}

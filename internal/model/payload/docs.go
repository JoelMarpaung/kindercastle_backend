package payload

type (
	ResponseDataDetailBook = ResponseData[Book]
	ResponseListBook       = PaginatedResponse[Book]
)

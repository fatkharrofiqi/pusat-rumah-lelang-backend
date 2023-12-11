package requests

type AutocompletePaginationRequest struct {
	Title    string `json:"title"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
}

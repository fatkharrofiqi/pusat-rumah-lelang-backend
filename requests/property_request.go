package requests

type PropertyPaginationRequest struct {
	Title    string `json:"title"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

package request

type GetComboPropertyRequest struct {
	Title string `json:"title"`
	Page  int    `json:"page" form:"page"`
	Size  int    `json:"size" form:"size"`
}

type AutocompletePaginationRequest struct {
	Title    string `json:"title"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
}

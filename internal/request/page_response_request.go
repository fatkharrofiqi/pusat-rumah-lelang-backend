package request

type PageResponse[T any] struct {
	Data         []T          `json:"data,omitempty"`
	PageMetadata PageMetadata `json:"paging,omitempty"`
}

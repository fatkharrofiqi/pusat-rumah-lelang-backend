package request

type WebResponse[T any] struct {
	Data   T             `json:"data"`
	Paging *PageMetadata `json:"paging,omitempty"`
	Errors string        `json:"errors,omitempty"`
}

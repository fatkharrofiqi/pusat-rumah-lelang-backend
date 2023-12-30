package interfaces

type IGetAllGeneric[T any] interface {
	GetAll(page, pageSize int) ([]T, error)
}

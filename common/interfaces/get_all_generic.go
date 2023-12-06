package interfaces

type IGetAllGeneric[T any] interface {
	GetAll() ([]T, error)
}

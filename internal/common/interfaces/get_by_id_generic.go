package interfaces

type IGetByIdGeneric[T any] interface {
	GetById(id int64) (T, error)
}

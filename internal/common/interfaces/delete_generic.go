package interfaces

type IDeleteGeneric[T any] interface {
	Delete(id int64) error
}

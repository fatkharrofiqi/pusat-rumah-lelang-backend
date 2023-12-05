package interfaces

type IUpdateGeneric[T any] interface {
	Update(id int64, data *T) error
}

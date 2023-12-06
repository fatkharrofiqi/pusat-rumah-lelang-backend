package interfaces

type ICreateGeneric[T any] interface {
	Create(data *T) error
}

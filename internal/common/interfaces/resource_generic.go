package interfaces

type IGenericResource[T any] interface {
	ICreateGeneric[T]
	IUpdateGeneric[T]
	IGetByIdGeneric[T]
	IGetAllGeneric[T]
	IDeleteGeneric[T]
}

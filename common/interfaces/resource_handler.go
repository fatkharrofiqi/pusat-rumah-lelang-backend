package interfaces

type IResourceHandler interface {
	ICreateHandler
	IGetAllHandler
	IGetByIdHandler
	IUpdateHandler
	IDeleteHandler
}

package usecases

type BookRepository interface {
	Create()
	Edit()
	Delete()
	FindByID()
	FindAll()
}

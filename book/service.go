package book

type Service interface {
	FindAll()([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookrequest BookRequest)(Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func(s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}
func(s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}
func(s *service) Create(bookrequest BookRequest) (Book, error) {
	book := Book{
		Title: bookrequest.Title,
		Price: bookrequest.Price,
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}

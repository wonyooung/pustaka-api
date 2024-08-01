package book

type Service interface {
	FindAll()([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookrequest BookRequest)(Book, error)
	Delete(ID int)(Book, error)
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
		Description: bookrequest.Description,
		Rating: bookrequest.Rating,
		Discount: bookrequest.Discount,
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}
func(s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	newBook, err := s.repository.Delete(book)
	return newBook, err
}





package usecase

var URL = "http://localhost:8080/generate-salt"

type UseCase struct {
	User CreateAndRead
}

func New(User CreateAndRead) *UseCase {
	return &UseCase{
		User: NewUser(User),
	}
}

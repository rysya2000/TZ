package usecase

type UseCase struct {
	// service_1 Service_1
	Gen_Salt
}

func New() Service_1 {
	return &UseCase{}
		// service_1: New_Gen_Salt(service),}
}

// func (u *UseCase) GenerateSalt() string {
// 	return Generate(12)
// }


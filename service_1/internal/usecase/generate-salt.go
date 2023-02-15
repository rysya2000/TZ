package usecase

type Gen_Salt struct {
	service_1 Service_1
}

func New_Gen_Salt(service Service_1) *Gen_Salt {
	return &Gen_Salt{
		service_1: service,
	}
}

func (g *Gen_Salt) GenerateSalt() string {
	return Generate(12)
}

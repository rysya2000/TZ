package usecase

type (
	Service_1 interface {
		GenerateSalt() string
	}
	Store interface {
		StoreSomething() error
	}
)

package service

type Service struct {
	User  UserService
	Batch BatchService
}

func New(userService UserService, batchService BatchService) *Service {
	return &Service{
		User:  userService,
		Batch: batchService,
	}
}

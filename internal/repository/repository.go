package repository

type Repository struct {
	User  UserRepository
	Batch BatchRepository
}

func New(userRepository UserRepository, batchRepository BatchRepository) *Repository {
	return &Repository{
		User:  userRepository,
		Batch: batchRepository,
	}
}

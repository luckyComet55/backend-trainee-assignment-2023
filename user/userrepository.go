package user

type UserRepository struct {
	Db UserDatabase
}

func NewUserRepository(db UserDatabase) UserRepository {
	return UserRepository{
		Db: db,
	}
}

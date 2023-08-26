package user

type User struct {
	id int
}

func NewUser(id int) User {
	return User{
		id: id,
	}
}

func (u User) GetId() int {
	return u.id
}

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

type ErrUnsupportedMethod struct {
}

func (e ErrUnsupportedMethod) Error() string {
	return "method is unsupported on object, User has no name"
}

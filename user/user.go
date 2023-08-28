package user

var idCounter int = 0

type User struct {
	id int `ksql:"id"`
}

func NewUser() User {
	newId := idCounter
	idCounter++
	return User{
		id: newId,
	}
}

func (u User) GetId() int {
	return u.id
}

package user

import (
	"context"
	"fmt"

	db_ "github.com/luckyComet55/backend-trainee-assignment-2023/database"
	"github.com/vingarcia/ksql"
)

type UserActualDatabase struct {
	db    ksql.DB
	table ksql.Table
}

func NewUserActualDatabase(db ksql.DB) *UserActualDatabase {
	return &UserActualDatabase{
		db:    db,
		table: ksql.NewTable("users"),
	}
}

func (d *UserActualDatabase) GetObjectById(id int) (User, error) {
	var user User
	err := d.db.QueryOne(context.Background(), &user, "select * from users where id=?", id)
	if err != nil {
		fmt.Println(err)
	}
	return user, err
}

func (d *UserActualDatabase) CreateObject(user User) error {
	query := fmt.Sprintf("insert into users values('%d')", user.id)
	_, err := d.db.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
		err = db_.ErrUniqueConstraintFailed{Field: "id", Value: fmt.Sprintf("%d", user.id)}
	}
	return err
}

func (d *UserActualDatabase) DeleteObject(user User) error {
	err := d.db.Delete(context.Background(), d.table, &user)
	if err != nil {
		err = db_.ErrObjNotFound{}
	}
	return err
}

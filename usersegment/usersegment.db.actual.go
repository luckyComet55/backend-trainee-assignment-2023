package usersegment

import (
	"context"
	"fmt"

	db_ "github.com/luckyComet55/backend-trainee-assignment-2023/database"
	"github.com/vingarcia/ksql"
)

type UserSegmentActualDatabase struct {
	db    ksql.DB
	table ksql.Table
}

func NewUserSegmentActualDatabase(db ksql.DB) *UserSegmentActualDatabase {
	return &UserSegmentActualDatabase{
		db:    db,
		table: ksql.NewTable("user_segments"),
	}
}

func (d *UserSegmentActualDatabase) GetObjectById(id int) (UserSegment, error) {
	return UserSegment{}, db_.ErrUnsupportedMethod{}
}

func (d *UserSegmentActualDatabase) CreateObject(s UserSegment) error {
	_, err := d.db.Exec(context.Background(), "insert into user_segments(user_id, segment_name) values($1, $2)", s.UserId, s.SegmentName)
	if err != nil {
		fmt.Println(err)
		err = db_.ErrUniqueConstraintFailed{Field: "user_id&segment_name", Value: fmt.Sprintf("%d&%s", s.UserId, s.SegmentName)}
	}
	return err
}

func (d *UserSegmentActualDatabase) DeleteObject(s UserSegment) error {
	err := d.db.Delete(context.Background(), d.table, &s)
	if err != nil {
		err = db_.ErrObjNotFound{}
	}
	return err
}

func (d *UserSegmentActualDatabase) GetByUserId(id int) []UserSegment {
	var res []UserSegment
	err := d.db.Query(context.Background(), &res, "select * from user_segments where user_id=$1", id)
	if err != nil {
		fmt.Println(err)
		res = nil
	}
	return res
}

func (d *UserSegmentActualDatabase) GetBySegmentName(name string) []UserSegment {
	res := make([]UserSegment, 0)
	err := d.db.Query(context.Background(), &res, "select * from user_segments where segment_name=$1", name)
	if err != nil {
		fmt.Println(err)
		res = nil
	}
	return res
}

func (d *UserSegmentActualDatabase) DeleteByUserId(id int) error {
	_, err := d.db.Exec(context.Background(), "delete from user_segments where user_id=$1", id)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (d *UserSegmentActualDatabase) DeleteBySegmentName(name string) error {
	_, err := d.db.Exec(context.Background(), "delete from user_segments where segment_name=$1", name)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (d *UserSegmentActualDatabase) DeleteByUserIdWithSegmentName(user_id int, segement_name string) error {
	_, err := d.db.Exec(context.Background(), "delete from user_segments where user_id=$1 and segment_name=$2", user_id, segement_name)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

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
	var res UserSegment
	err := d.db.QueryOne(context.Background(), &res, "select * from user_segments where id=$1", id)
	if err != nil {
		err = db_.ErrObjNotFound{}
	}
	return res, err
}

func (d *UserSegmentActualDatabase) CreateObject(s UserSegment) error {
	_, err := d.db.Exec(context.Background(), "insert into user_segments values($1, $2, $3)", s.Id, s.UserId, s.SegmentId)
	if err != nil {
		fmt.Println(err)
		err = db_.ErrUniqueConstraintFailed{Field: "user_id&segment_id", Value: fmt.Sprintf("%d&%d", s.UserId, s.SegmentId)}
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

func (d *UserSegmentActualDatabase) GetBySegmentId(id int) []UserSegment {
	res := make([]UserSegment, 0)
	err := d.db.Query(context.Background(), &res, "select * from user_segments where segment_id=$1", id)
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

func (d *UserSegmentActualDatabase) DeleteBySegmentId(id int) error {
	_, err := d.db.Exec(context.Background(), "delete from user_segments where segment_id=$1", id)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (d *UserSegmentActualDatabase) DeleteByUserIdWithSegmentId(user_id, segement_id int) error {
	_, err := d.db.Exec(context.Background(), "delete from user_segments where user_id=$1 and segment_id=$2", user_id, segement_id)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

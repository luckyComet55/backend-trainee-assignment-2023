package segment

import (
	"context"
	"fmt"

	db_ "github.com/luckyComet55/backend-trainee-assignment-2023/database"
	"github.com/vingarcia/ksql"
)

type SegmentActualDatabase struct {
	db    ksql.DB
	table ksql.Table
}

func NewSegmentActualDatabase(db ksql.DB) *SegmentActualDatabase {
	return &SegmentActualDatabase{
		db:    db,
		table: ksql.NewTable("segments"),
	}
}

func (d *SegmentActualDatabase) GetObjectById(id int) (Segment, error) {
	var res Segment
	err := d.db.QueryOne(context.Background(), &res, "select * from segments where id=?", id)
	if err != nil {
		err = db_.ErrObjNotFound{}
	}
	return res, err
}

func (d *SegmentActualDatabase) CreateObject(s Segment) error {
	query := fmt.Sprintf("insert into segments values('%d', '%s')", s.Id, s.Name)
	fmt.Println(query)
	_, err := d.db.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
		err = db_.ErrUniqueConstraintFailed{Field: "name", Value: s.Name}
	}
	return err
}

func (d *SegmentActualDatabase) UpdateObject(s Segment) error {
	err := d.db.Patch(context.Background(), d.table, &s)
	if err != nil {
		err = db_.ErrObjNotFound{}
	}
	return err
}

func (d *SegmentActualDatabase) DeleteObject(s Segment) error {
	err := d.db.Delete(context.Background(), d.table, &s)
	if err != nil {
		err = db_.ErrObjNotFound{}
	}
	return err
}

func (d *SegmentActualDatabase) DeleteByName(name string) error {
	queryString := fmt.Sprintf("delete from segments where name='%s'", name)
	_, err := d.db.Exec(context.Background(), queryString)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (d *SegmentActualDatabase) GetByName(name string) (Segment, error) {
	var res Segment
	err := d.db.QueryOne(context.Background(), &res, "select * from segments where name=?", name)
	if err != nil {
		err = db_.ErrObjNotFound{}
	}
	return res, err
}

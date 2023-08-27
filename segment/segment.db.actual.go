package segment

import (
	"context"
	"fmt"

	db_ "github.com/luckyComet55/backend-trainee-assignment-2023/database"
	"github.com/vingarcia/ksql"
)

type SegmentActualDatabase struct {
	db                 ksql.DB
	table              ksql.Table
	queryStrFindById   string
	queryStrFindByName string
}

func NewSegmentActualDatabase(db ksql.DB) *SegmentActualDatabase {
	return &SegmentActualDatabase{
		db:                 db,
		table:              ksql.NewTable("segments"),
		queryStrFindById:   "select * from segments where id=?",
		queryStrFindByName: "select * from segments where name=?",
	}
}

func (d *SegmentActualDatabase) GetObjectById(id int) (Segment, error) {
	var res Segment
	err := d.db.QueryOne(context.Background(), &res, d.queryStrFindById, id)
	if err != nil {
		err = db_.ErrObjNotFound{}
	}
	return res, err
}

func (d *SegmentActualDatabase) CreateObject(s Segment) error {
	err := d.db.Insert(context.Background(), d.table, &s)
	if err != nil {
		fmt.Println(err)
		err = db_.ErrObjNotFound{}
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

func (d *SegmentActualDatabase) GetByName(name string) (Segment, error) {
	var res Segment
	err := d.db.QueryOne(context.Background(), &res, d.queryStrFindByName, name)
	if err != nil {
		err = db_.ErrObjNotFound{}
	}
	return res, err
}

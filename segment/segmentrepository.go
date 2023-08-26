package segment

type SegmentRepository struct {
	Db SegmentDatabase
}

func NewSegmentRepository(db SegmentDatabase) SegmentRepository {
	return SegmentRepository{
		Db: db,
	}
}

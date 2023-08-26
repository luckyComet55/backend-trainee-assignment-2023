package usersegment

type UserSegmentRepository struct {
	Db UserSegmentDatabase
}

func NewUserSegmentRepository(db UserSegmentDatabase) UserSegmentRepository {
	return UserSegmentRepository{
		Db: db,
	}
}

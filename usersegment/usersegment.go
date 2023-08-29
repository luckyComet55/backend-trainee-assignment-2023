package usersegment

type UserSegment struct {
	UserId      int    `ksql:"user_id"`
	SegmentName string `ksql:"segment_name"`
}

func NewUserSegment(userId int, segmentName string) UserSegment {
	return UserSegment{
		UserId:      userId,
		SegmentName: segmentName,
	}
}

func (u UserSegment) GetUserId() int {
	return u.UserId
}

func (u UserSegment) GetSegmentName() string {
	return u.SegmentName
}

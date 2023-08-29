package usersegment

var idCounter int = 0

type UserSegment struct {
	Id          int    `ksql:"id"`
	UserId      int    `ksql:"user_id"`
	SegmentName string `ksql:"segment_name"`
}

func NewUserSegment(userId int, segmentName string) UserSegment {
	newId := idCounter
	idCounter++
	return UserSegment{
		Id:          newId,
		UserId:      userId,
		SegmentName: segmentName,
	}
}

func (u UserSegment) GetId() int {
	return u.Id
}

func (u UserSegment) GetUserId() int {
	return u.UserId
}

func (u UserSegment) GetSegmentName() string {
	return u.SegmentName
}

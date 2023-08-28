package usersegment

var idCounter int = 0

type UserSegment struct {
	Id        int `ksql:"id"`
	UserId    int `ksql:"user_id"`
	SegmentId int `ksql:"segment_id"`
}

func NewUserSegment(userId, segmentId int) UserSegment {
	newId := idCounter
	idCounter++
	return UserSegment{
		Id:        newId,
		UserId:    userId,
		SegmentId: segmentId,
	}
}

func (u UserSegment) GetId() int {
	return u.Id
}

func (u UserSegment) GetUserId() int {
	return u.UserId
}

func (u UserSegment) GetSegmentId() int {
	return u.SegmentId
}

package usersegment

var idCounter int = 0

type UserSegment struct {
	id        int
	userId    int
	segmentId int
}

func NewUserSegment(userId, segmentId int) UserSegment {
	newId := idCounter
	idCounter++
	return UserSegment{
		id:        newId,
		userId:    userId,
		segmentId: segmentId,
	}
}

func (u UserSegment) GetId() int {
	return u.id
}

func (u UserSegment) GetUserId() int {
	return u.userId
}

func (u UserSegment) GetSegmentId() int {
	return u.segmentId
}

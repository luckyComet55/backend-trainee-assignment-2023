package usersegment

type UserSegment struct {
	id        int
	userId    int
	segmentId int
}

func NewUserSegment(id, userId, segmentId int) UserSegment {
	return UserSegment{
		id:        id,
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

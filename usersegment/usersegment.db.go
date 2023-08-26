package usersegment

import (
	db "github.com/luckyComet55/backend-trainee-assignment-2023/database"
)

type UserSegmentDatabase interface {
	db.Database[UserSegment]
	GetByUserId(int) []UserSegment
	GetBySegmentId(int) []UserSegment
	DeleteByUserId(int) error
	DeleteBySegmentId(int) error
	DeleteByUserIdWithSegmentId(int, int) error
}

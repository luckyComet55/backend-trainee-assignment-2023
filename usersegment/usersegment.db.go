package usersegment

import (
	db "github.com/luckyComet55/backend-trainee-assignment-2023/database"
)

type UserSegmentDatabase interface {
	db.Database[UserSegment]
	GetByUserId(int) []UserSegment
	GetBySegmentName(string) []UserSegment
	DeleteByUserId(int) error
	DeleteBySegmentName(string) error
	DeleteByUserIdWithSegmentName(int, string) error
}

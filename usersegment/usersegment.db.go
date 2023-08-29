package usersegment

import (
	db "github.com/luckyComet55/backend-trainee-assignment-2023/database"
	sg "github.com/luckyComet55/backend-trainee-assignment-2023/segment"
)

type UserSegmentDatabase interface {
	db.Database[UserSegment]
	GetByUserId(int) []UserSegment
	GetBySegmentName(string) []UserSegment
	GetUserActiveSegments(int) []sg.Segment
	DeleteByUserId(int) error
	DeleteBySegmentName(string) error
	DeleteByUserIdWithSegmentName(int, string) error
	SetUserSegmentInactive(int, string) error
}

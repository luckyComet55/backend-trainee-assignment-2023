package repository

import (
	sg "github.com/luckyComet55/backend-trainee-assignment-2023/segment"
	usr "github.com/luckyComet55/backend-trainee-assignment-2023/user"
)

type ServiceRepository interface {
	GetUsersBySegmentId(int) ([]usr.User, error)
	GetSegmentsByUserId(int) ([]sg.Segment, error)
	CheckNonExistantSegments([]string) ([]string, []int)
}

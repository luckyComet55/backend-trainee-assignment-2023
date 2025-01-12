package repository

import (
	sg "github.com/luckyComet55/backend-trainee-assignment-2023/segment"
)

type ServiceRepository interface {
	GetUserIdsBySegmentName(string) ([]int, error)
	GetSegmentsByUserId(int) ([]sg.Segment, error)
	GetUserActiveSegments(int) ([]string, error)
	CheckNonExistantSegments([]string) ([]string, []string)
	SetRandomSegmentAuditory(sg.Segment) error
}

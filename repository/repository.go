package repository

import (
	sg "github.com/luckyComet55/backend-trainee-assignment-2023/segment"
)

type ServiceRepository interface {
	GetUserIdsBySegmentId(int) ([]int, error)
	GetSegmentsByUserId(int) ([]sg.Segment, error)
	CheckNonExistantSegments([]string) ([]string, []int)
}

package segment

import (
	db "github.com/luckyComet55/backend-trainee-assignment-2023/database"
)

type SegmentDatabase interface {
	db.Database[Segment]
	GetByName(string) (Segment, error)
}

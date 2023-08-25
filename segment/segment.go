package segment

type Segment struct {
	Id   int
	Name string
}

func NewSegment(id int, name string) Segment {
	return Segment{
		Id:   id,
		Name: name,
	}
}

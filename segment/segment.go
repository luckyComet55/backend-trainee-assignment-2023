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

func (s Segment) GetId() int {
	return s.Id
}

func (s Segment) GetName() string {
	return s.Name
}

package segment

var idCounter int = 0

type Segment struct {
	id   int
	name string
}

func NewSegment(name string) Segment {
	newId := idCounter
	idCounter++
	return Segment{
		id:   newId,
		name: name,
	}
}

func (s Segment) GetId() int {
	return s.id
}

func (s Segment) GetName() string {
	return s.name
}

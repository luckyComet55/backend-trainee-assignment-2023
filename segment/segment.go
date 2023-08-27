package segment

var idCounter int = 0

type Segment struct {
	Id   int    `json:"-" ksql:"id"`
	Name string `json:"name" ksql:"name"`
}

func NewSegment(name string) Segment {
	newId := idCounter
	idCounter++
	return Segment{
		Id:   newId,
		Name: name,
	}
}

func (s Segment) GetId() int {
	return s.Id
}

func (s Segment) GetName() string {
	return s.Name
}

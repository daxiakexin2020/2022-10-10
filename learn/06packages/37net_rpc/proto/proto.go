package proto

type PServer struct {
	List map[string]*People
}

type People struct {
	Id   string
	Name string
	Age  int
}

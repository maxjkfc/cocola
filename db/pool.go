package db

type Pool interface {
	Add(Session) error
	List() map[int]string
	Status()
}

type pool struct {
	num int
	db  map[int]Session
}

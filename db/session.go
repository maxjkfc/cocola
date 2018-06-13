package db

type Session interface {
}

type session struct {
	config SessionConfig
}

type SessionConfig struct {
	DBtype   string
	Account  string
	Password string
	Database string
	Host     string
	Protocol string
}

func Dial(config SessionConfig) Session {
}

func (s *session) Connect() {
	switch config.DBtype {
	case Mongo:
	case Mariadb, Mysql:
	case redis:
	}
}

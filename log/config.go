package log

type Configer interface {
	Level(string) Configer
	Name(string) Configer
	Env(string) Configer
	AddCallerSkip(int) Configer
	Build() (Logger, error)
}

func NewZapConfig() Configer {
	return newZapConfig()
}

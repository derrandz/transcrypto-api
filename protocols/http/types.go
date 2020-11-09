package http

type (
	Service interface {
		Configure(arg ...interface{})
		Run() (bool, error)
		Stop() (bool, error)
	}
)

package domain

type (
	// TODO: not sure that logger interface should be here. Even not sure if it's worth having interface at all.
	LoggerInterface interface {
		Info(format string, v ...interface{})
		Error(format string, v ...interface{})
	}
)

package logger

type (
	Logger interface {
		Info(fields LogFields)
		Warning(fields LogFields)
		Error(fields LogFields)
	}
	LogFields struct {
		Section  string
		Function string
		Params   interface{}
		Message  string
	}
)

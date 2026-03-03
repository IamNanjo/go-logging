package logging

func NewLogger(config LoggerConfig) *Logger {
	l := Logger{}

	return &l
}

type LoggerConfig struct {
	// Start logs with RFC 3339 (ISO 8601) formatted current time
	PrependTime bool
}

type Logger struct {
	outPrefix  string
	outSuffix  string
	okPrefix   string
	okSuffix   string
	warnPrefix string
	warnSuffix string
	errPrefix  string
	errSuffix  string
}

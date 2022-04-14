package logger

func Infof(msg string, args ...interface{}) {
	defaultLogger.base.Infof(msg, args...)
}

func Infow(msg string, kvs ...interface{}) {
	defaultLogger.base.Infow(msg, kvs...)
}

func Debugf(msg string, args ...interface{}) {
	defaultLogger.base.Debugf(msg, args...)
}

func Debugw(msg string, kvs ...interface{}) {
	defaultLogger.base.Debugw(msg, kvs...)
}

func Warnf(msg string, args ...interface{}) {
	defaultLogger.base.Warnf(msg, args...)
}

func Warnw(msg string, kvs ...interface{}) {
	defaultLogger.base.Warnw(msg, kvs...)
}

func Errorf(msg string, args ...interface{}) {
	defaultLogger.base.Errorf(msg, args...)
}

func Errorw(msg string, kvs ...interface{}) {
	defaultLogger.base.Errorw(msg, kvs...)
}

func Fatalf(msg string, args ...interface{}) {
	defaultLogger.base.Fatalf(msg, args...)
}

func Fatalw(msg string, kvs ...interface{}) {
	defaultLogger.base.Fatalw(msg, kvs...)
}

func Panicf(msg string, args ...interface{}) {
	defaultLogger.base.Panicf(msg, args...)
}

func Panicw(msg string, kvs ...interface{}) {
	defaultLogger.base.Panicw(msg, kvs...)
}

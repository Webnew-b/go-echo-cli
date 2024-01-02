package database

import (
	"golang.org/x/net/context"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

type SqlLogger struct {
	Writer *log.Logger
	Level  logger.LogLevel
}

func NewLogger(writer *log.Logger, logMode logger.LogLevel) *SqlLogger {
	return &SqlLogger{
		Writer: writer,
		Level:  logMode,
	}
}

func (l *SqlLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.Level = level
	return &newLogger
}

func (l *SqlLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.Level >= logger.Info {
		l.Writer.Printf(msg, data...)
	}
}

func (l *SqlLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.Level >= logger.Warn {
		l.Writer.Printf(msg, data...)
	}
}

func (l *SqlLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.Level >= logger.Error {
		l.Writer.Printf(msg, data...)
	}
}

func (l *SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.Level <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		l.Writer.Printf("trace error: %s [%.2fms] [%d rows] %s", err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
	} else {
		l.Writer.Printf("trace info: [%.2fms] [%d rows] %s", float64(elapsed.Nanoseconds())/1e6, rows, sql)
	}
}

package utils

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Log *Logger

type Logger struct {
	logrus *logrus.Logger
}

func init() {
	Log = NewLogger()
}

func NewLogger() *Logger {
	log := logrus.New()

	log.SetReportCaller(true)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.Level = logrus.DebugLevel

	// 默认输出到log.txt
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Error("Failed to log to file, using default stderr")
	} else {
		log.SetOutput(file)
	}

	return &Logger{
		logrus: log,
	}
}

// 使日志同时输出到文件和控制台
func (l *Logger) OutputToConsole(alsoToConsole bool) {
	if alsoToConsole {
		l.logrus.SetOutput(io.MultiWriter(os.Stdout, l.logrus.Out))
	} else {
		// 只输出到原始设置，这里我们重新打开文件
		file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			l.logrus.Error("Failed to log to file, using default stderr")
			l.logrus.SetOutput(os.Stderr)
		} else {
			l.logrus.SetOutput(file)
		}
	}
}

func (l *Logger) SetOutput(output io.Writer) {
	l.logrus.SetOutput(output)
}

func (l *Logger) SetLevel(lvl logrus.Level) {
	l.logrus.SetLevel(lvl)
}

func (l *Logger) Panic(args ...interface{}) {
	l.logrus.Panic(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logrus.Fatal(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.logrus.Error(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.logrus.Warn(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.logrus.Info(args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.logrus.Debug(args...)
}

func (l *Logger) Printf(format string, args ...interface{}) {
    // 我们假设您想使用 InfoLevel 记录此消息
    l.logrus.Infof(format, args...)
}


func ContextLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request

		beginTime := time.Now()
		c.Next()
		endTime := time.Since(beginTime)

		Log.logrus.WithFields(logrus.Fields{
			"time_rfc3339": time.Now().Format(time.RFC3339),
			"remote_ip":    c.ClientIP(),
			"host":         req.Host,
			"uri":          req.RequestURI,
			"method":       req.Method,
			"referer":      req.Referer(),
			"user_agent":   req.UserAgent(),
			"status":       c.Writer.Status(),
			"latency":      fmt.Sprintf("%d µs", endTime.Microseconds()),
			// Assuming you will add these later, using placeholders for now
			"req_bytes":  "N/A",
			"resp_bytes": "N/A",
		}).Info("http request")
	}
}

package logger

import (
	"log"
	"os"
	"time"
)

var l *log.Logger

func Init() {
	l = log.New(os.Stdout, "", 0)
	Info("initialised loggigng")
}

func Info(msgs ...any) {
	print("INFO", msgs...)
}
func Warn(msgs ...any) {
	print("WARN", msgs...)
}
func Error(msgs ...any) {
	print("ERROR", msgs...)
}
func Fatal(msgs ...any) {
	print("FATAL", msgs...)
	panic("panicing due to the fatal error above")
}

func print(t string, msgs ...any) {
	if l == nil {
		panic("usage of uninitialised logger")
	}

	buff := time.Now().Format(time.RFC822Z)
	buff += "\t" + t + "\t"
	for _, m := range msgs {
		buff += m.(string)
	}
	buff += "\n"

	l.Print(buff)
}

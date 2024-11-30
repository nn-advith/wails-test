package clogger

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

func getTimeStamp() (string, string) {
	timenow := time.Now()
	dateslice := strings.Split(timenow.Format(time.DateOnly), "-")
	slices.Reverse(dateslice)
	date := strings.Join(dateslice, "-")
	return date, timenow.Format(time.TimeOnly)
}

type Logger struct {
	basedir string
}

func InitLogger() Logger {

	logdir, _ := filepath.Abs("../LOGS")
	newLogger := Logger{""}
	newLogger.basedir = logdir

	err := os.MkdirAll(logdir, 0777)
	if err != nil {
		fmt.Println("ERROR while configuring log directory::::", err)
		panic("ERROR while configuring log directory::::")
	}
	return newLogger
}

func (l *Logger) SetPath(path string) {

	l.basedir = path
}

func (l Logger) GetPath() string {
	return l.basedir
}

func (l Logger) fileLog(level string, msg string) {

	dstamp, tstamp := getTimeStamp()
	filename := "thegray_" + dstamp + ".log"
	file, err := os.OpenFile(filepath.Join(l.basedir, filename), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("ERROR while opening log file ", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("\n"+dstamp+" "+tstamp+" |: %s |: %s", level, msg))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

}

func (l Logger) Print(msg string) {
	l.fileLog("LOG", msg)
	fmt.Println("LOG |: ", msg)
}

func (l Logger) Debug(msg string) {
	l.fileLog("DEBUG", msg)
	fmt.Println("DEBUG |: ", msg)
}

func (l Logger) Error(msg string) {
	l.fileLog("ERROR", msg)
	fmt.Println("ERROR |: ", msg)
}

func (l Logger) Warning(msg string) {
	l.fileLog("WARNING", msg)
	fmt.Println("WARNING |: ", msg)
}

func (l Logger) Fatal(msg string) {
	l.fileLog("FATAL", msg)
	fmt.Println("FATAL |: ", msg)
	os.Exit(1)
}

func (l Logger) Info(msg string) {
	l.fileLog("INFO", msg)
	fmt.Println("INFO |: ", msg)
}

func (l Logger) Trace(msg string) {
	l.fileLog("TRACE", msg)
	fmt.Println("TRACE |: ", msg)
}

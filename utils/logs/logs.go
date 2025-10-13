package logs

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var (
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Info = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// InfoPrintln writes an info message to the file logger and inserts into DB (best-effort).
func InfoPrintln(v ...interface{}) {
	Info.Println(v...)
	file := callerFile()
	_ = InsertLog("INFO", fmt.Sprint(v...), file)
}

// WarningPrintln writes a warning message and inserts into DB (best-effort).
func WarningPrintln(v ...interface{}) {
	Warning.Println(v...)
	file := callerFile()
	_ = InsertLog("WARNING", fmt.Sprint(v...), file)
}

// ErrorPrintln writes an error message and inserts into DB (best-effort).
func ErrorPrintln(v ...interface{}) {
	Error.Println(v...)
	file := callerFile()
	_ = InsertLog("ERROR", fmt.Sprint(v...), file)
}

func callerFile() string {
	_, file, _, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	return file
}

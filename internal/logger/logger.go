package logger

import (
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

var (
	logfile   *os.File
	createdAt time.Time
)

func NewLogFile() {
	// Open the file for reading and writing, create it if it doesn't exist, clear it if it does
	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)
	os.MkdirAll(filepath.Join(exeDir, "../log"), os.ModePerm)
	createdAt = time.Now()
	logfilepath := filepath.Join(exeDir, "../log", createdAt.Format("20060102150405")+".log")
	logfile, _ = os.OpenFile(logfilepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
}

func Log(msg ...interface{}) {
	if logfile == nil {
		NewLogFile()
	}
	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)
	_, err := os.Stat(filepath.Join(exeDir, "../log", createdAt.Format("20060102150405")+".log"))
	if os.IsNotExist(err) || time.Since(createdAt).Hours() > 1 {
		NewLogFile()
	}
	// Get the file name and line number of the caller
	_, file, line, _ := runtime.Caller(1)
	file = filepath.Base(file)

	fullMessage := ""
	// Write the messages to the log file
	for _, m := range msg {
		switch v := m.(type) {
		case string:
			fullMessage += v
		case []string:
			fullMessage += "["
			for _, s := range v {
				fullMessage += s + ", "
			}
			fullMessage = fullMessage[:len(fullMessage)-2] + "]"
		case int:
			fullMessage += strconv.Itoa(v)
		case []int:
			fullMessage += "["
			for _, i := range v {
				fullMessage += strconv.Itoa(i) + ", "
			}
			fullMessage = fullMessage[:len(fullMessage)-2] + "]"
		case float32:
			fullMessage += strconv.FormatFloat(float64(v), 'f', -1, 64)
		case []float32:
			fullMessage += "["
			for _, f := range v {
				fullMessage += strconv.FormatFloat(float64(f), 'f', -1, 64) + ", "
			}
			fullMessage = fullMessage[:len(fullMessage)-2] + "]"
		case float64:
			fullMessage += strconv.FormatFloat(v, 'f', -1, 64)
		case []float64:
			fullMessage += "["
			for _, f := range v {
				fullMessage += strconv.FormatFloat(f, 'f', -1, 64) + ", "
			}
			fullMessage = fullMessage[:len(fullMessage)-2] + "]"
		case bool:
			fullMessage += strconv.FormatBool(v)
		case error:
			fullMessage += v.Error()
		default:
			fullMessage += ""
		}
		fullMessage += " "
	}
	logfile.WriteString("[" + time.Now().Format("2006-01-02 15:04:05") + "] " + fullMessage + "(" + file + " line: " + strconv.Itoa(line) + ")\n")
}

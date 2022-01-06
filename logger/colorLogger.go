package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

type LEVEL int
type COLOR int
type STYLE int

// level
const (
	DEBUG LEVEL = iota
	INFO
	WARN
	ERROR
	FATAL
)

// color
const (
	COLOR_RED     = COLOR(31)
	COLOR_BLUE    = COLOR(34)
	COLOR_YELLOW  = COLOR(33)
	COLOR_PURPLE  = COLOR(35)
	COLOR_DEFAULT = COLOR(39)
)

// style
const (
	STYLE_DEFAULT   = STYLE(0)
	STYLE_HIGHLIGHT = STYLE(1)
)

// flag
const (
	LOG_FLAGS = log.Ldate | log.Lmicroseconds | log.Lshortfile
	LOG_FLAG  = 0
)

type LogFile struct {
	sync.RWMutex
	dir      string
	fileName string
	filePath string
	time     time.Time
	file     *os.File
	logger   *log.Logger
}

var (
	logLevel   LEVEL = DEBUG
	logConsole bool  = true
	logPrefix  string
	logFile    *LogFile
)

func SetConsole(isConsole bool) {
	logConsole = isConsole
}

func SetPrefix(prefix string) {
	logPrefix = fmt.Sprintf("[%s] ", prefix)
}

func SetLevel(level LEVEL) {
	logLevel = level
}

func SetFile(fileDir string, fileName string) {

	dir := fileDir
	if fileDir[len(fileDir)-1] == '\\' || fileDir[len(fileDir)-1] == '/' {
		dir = fileDir[:len(fileDir)-1]
	}

	logFile = &LogFile{dir: dir, fileName: fileName, time: time.Now()}

	logFile.Lock()
	defer logFile.Unlock()

	os.MkdirAll(logFile.dir, os.ModePerm)

	logFile.filePath = fmt.Sprintf("%s/%s-%04d%02d%02d.log", logFile.dir, logFile.fileName,
		logFile.time.Year(), logFile.time.Month(), logFile.time.Day())

	var err error
	logFile.file, err = os.OpenFile(logFile.filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	logFile.logger = log.New(logFile.file, logPrefix, LOG_FLAGS)
}

func init() {
	log.SetFlags(LOG_FLAG)
}

func Debug(v ...interface{}) {
	if logLevel > DEBUG {
		return
	}

	context := strings.TrimRight(fmt.Sprintf("DEBUG: %s", fmt.Sprintln(v...)), "\n")

	if logFile != nil {
		logFile.RLock()
		defer logFile.RUnlock()
		logFile.logger.Output(2, context)
	}

	console(DEBUG, context)
}

func Info(v ...interface{}) {
	if logLevel > INFO {
		return
	}

	context := strings.TrimRight(fmt.Sprintf("INFO: %s", fmt.Sprintln(v...)), "\n")

	if logFile != nil {
		logFile.RLock()
		defer logFile.RUnlock()
		logFile.logger.Output(2, context)
	}

	console(INFO, context)
}

func Warn(v ...interface{}) {
	if logLevel > WARN {
		return
	}

	context := strings.TrimRight(fmt.Sprintf("WARN: %s", fmt.Sprintln(v...)), "\n")

	if logFile != nil {
		logFile.RLock()
		defer logFile.RUnlock()
		logFile.logger.Output(2, context)
	}

	console(WARN, context)
}

func Error(v ...interface{}) {
	if logLevel > ERROR {
		return
	}

	context := strings.TrimRight(fmt.Sprintf("ERROR: %s", fmt.Sprintln(v...)), "\n")

	if logFile != nil {
		logFile.RLock()
		defer logFile.RUnlock()
		logFile.logger.Output(2, context)
	}

	console(ERROR, context)
}

func Fatal(v ...interface{}) {
	if logLevel > FATAL {
		return
	}

	context := strings.TrimRight(fmt.Sprintf("FATAL: %s", fmt.Sprintln(v...)), "\n")

	if logFile != nil {
		logFile.RLock()
		defer logFile.RUnlock()
		logFile.logger.Output(2, context)
	}

	console(FATAL, context)
}

func Debugf(format string, v ...interface{}) {
	if logLevel > DEBUG {
		return
	}

	context := strings.TrimRight(fmt.Sprintf("DEBUG: %s", fmt.Sprintf(format, v...)), "\n")

	if logFile != nil {
		logFile.RLock()
		defer logFile.RUnlock()
		logFile.logger.Output(2, context)
	}

	console(DEBUG, context)
}

func Infof(format string, v ...interface{}) {
	if logLevel > INFO {
		return
	}

	context := strings.TrimRight(fmt.Sprintf("INFO: %s", fmt.Sprintf(format, v...)), "\n")

	if logFile != nil {
		logFile.RLock()
		defer logFile.RUnlock()
		logFile.logger.Output(2, context)
	}

	console(INFO, context)
}

func Warnf(format string, v ...interface{}) {
	if logLevel > WARN {
		return
	}

	context := strings.TrimRight(fmt.Sprintf("WARN: %s", fmt.Sprintf(format, v...)), "\n")

	if logFile != nil {
		logFile.RLock()
		defer logFile.RUnlock()
		logFile.logger.Output(2, context)
	}

	console(WARN, context)
}

func Errorf(format string, v ...interface{}) {
	if logLevel > ERROR {
		return
	}

	context := strings.TrimRight(fmt.Sprintf("ERROR: %s", fmt.Sprintf(format, v...)), "\n")

	if logFile != nil {
		logFile.RLock()
		defer logFile.RUnlock()
		logFile.logger.Output(2, context)
	}

	console(ERROR, context)
}

func Fatalf(format string, v ...interface{}) {
	if logLevel > FATAL {
		return
	}

	context := strings.TrimRight(fmt.Sprintf("FATAL: %s", fmt.Sprintf(format, v...)), "\n")

	if logFile != nil {
		logFile.RLock()
		defer logFile.RUnlock()
		logFile.logger.Output(2, context)
	}

	console(FATAL, context)
}

func console(level LEVEL, context string) {

	if !logConsole {
		return
	}

	now := time.Now()

	_, file, line, _ := runtime.Caller(2)
	name := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			name = file[i+1:]
			break
		}
	}
	file = name

	context = fmt.Sprintf("%s%04d/%02d/%02d %02d:%02d:%02d.%06d %s:%d %s", logPrefix,
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(),
		time.Duration(now.Nanosecond())/(time.Microsecond), file, line, context,
	)

	switch level {
	case DEBUG:
		log.Println(PrintColor(context, STYLE_DEFAULT, COLOR_DEFAULT, COLOR_DEFAULT))
	case INFO:
		log.Println(PrintColor(context, STYLE_DEFAULT, COLOR_BLUE, COLOR_DEFAULT))
	case WARN:
		log.Println(PrintColor(context, STYLE_DEFAULT, COLOR_YELLOW, COLOR_DEFAULT))
	case ERROR:
		log.Println(PrintColor(context, STYLE_HIGHLIGHT, COLOR_RED, COLOR_DEFAULT))
	case FATAL:
		log.Println(PrintColor(context, STYLE_HIGHLIGHT, COLOR_PURPLE, COLOR_DEFAULT))
	default:
		log.Println(PrintColor(context, STYLE_DEFAULT, COLOR_DEFAULT, COLOR_DEFAULT))
	}
}

func PrintColor(context string, s STYLE, f COLOR, c COLOR) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, int(s), int(f), int(c)+10, context, 0x1B)
}

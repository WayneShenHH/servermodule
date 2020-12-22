package zaphdr

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/constant"
	"github.com/WayneShenHH/servermodule/logger/zaphdr/stack"
	"github.com/WayneShenHH/servermodule/slackAlert"
	"github.com/WayneShenHH/servermodule/util/color"
)

const (
	callerSkipOffset = 3 // zaphdr.fatal & zaphdr.FatalOnError & logger.FatalOnError
)

// key list
const (
	timeKey        = "time"
	stackstraceKey = "stacktrace"
	codekey        = "serviceCode"
	msgKey         = "msg"
	levelKey       = "level"
)

// level list
const (
	debugLevel = 0
	infoLevel  = 1
	warnLevel  = 2
	errorLevel = 3
	fatalLevel = 4
)

// formatter list
const (
	jsonFormatter = 1
	stdFormatter  = 2
)

var levelNameMap = map[constant.Level]int{
	constant.Debug: debugLevel,
	constant.Info:  infoLevel,
	constant.Warn:  warnLevel,
	constant.Error: errorLevel,
	constant.Fatal: fatalLevel,
}

var levelIDMap = map[int]constant.Level{
	debugLevel: constant.Debug,
	infoLevel:  constant.Info,
	warnLevel:  constant.Warn,
	errorLevel: constant.Error,
	fatalLevel: constant.Fatal,
}

var levelColorMap = map[int]string{
	debugLevel: color.Blue.Add("DEBUG"),
	infoLevel:  color.Cyan.Add("INFO"),
	warnLevel:  color.Yellow.Add("WARN"),
	errorLevel: color.Red.Add("ERROR"),
	fatalLevel: color.Magenta.Add("FATAL"),
}

var formatterNameMap = map[constant.LogFormatter]int{
	constant.JsonFormatter: jsonFormatter,
	constant.StdFormatter:  stdFormatter,
}

var (
	workingDir string
)

// Logger of zap implement
type Logger struct {
	level         int
	serviceCode   constant.ServiceCode
	fatalCallback func(msg string)
	outputFile    *os.File
	formatter     func(level int, addStack bool, fields []interface{}) string
}

func init() {
	if dir, err := os.Getwd(); err == nil {
		workingDir = dir
	}
}

// New init logger
func New(cfg *config.LoggerConfig) *Logger {
	if cfg == nil {
		cfg = &config.LoggerConfig{
			StdLevel: "debug",
		}
	}

	lv, exist := levelNameMap[cfg.StdLevel]
	if !exist {
		lv = debugLevel
	}

	l := &Logger{
		level:       lv,
		serviceCode: cfg.ServiceCode,
	}

	switch formatterNameMap[cfg.Formatter] {
	case jsonFormatter:
		l.formatter = l.jsonOutput
	case stdFormatter:
		l.formatter = l.stdOutput
	default:
		l.formatter = l.stdOutput
	}

	return l
}

// Fatal console log err
func (l *Logger) fatal(fields ...interface{}) {
	output := l.formatter(fatalLevel, true, fields)
	if l.fatalCallback != nil {
		l.fatalCallback(output)
	}
	if slackAlert.IsSlackEnable() {
		slackAlert.SendError(output)
	}

	os.Exit(1)
}

// Error console log err
func (l *Logger) error(fields ...interface{}) string {
	e := l.formatter(errorLevel, true, fields)
	if slackAlert.IsSlackEnable() {
		slackAlert.SendWarns(e)
	}
	return e
}

// Warn console log warn
func (l *Logger) warn(addStack bool, fields ...interface{}) {
	e := l.formatter(warnLevel, addStack, fields)
	if slackAlert.IsSlackEnable() {
		slackAlert.SendWarns(e)
	}
}

// Info console log info
func (l *Logger) info(addStack bool, fields ...interface{}) {
	l.formatter(infoLevel, addStack, fields)
}

// Debug console log debug
func (l *Logger) debug(addStack bool, fields ...interface{}) {
	l.formatter(debugLevel, addStack, fields)
}

func now() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func stdnow() string {
	return color.Green.Add(time.Now().UTC().Format(time.RFC3339))
}

func fieldsToString(fields []interface{}) string {
	fs, split := "", " "
	for idx := range fields {
		switch val := fields[idx].(type) {
		case error:
			fs += fmt.Sprint(val.Error(), split)
		case string:
			fs += fmt.Sprint(val, split)
		default:
			js, _ := json.Marshal(val)
			fs += fmt.Sprint(string(js), split)
		}
	}

	return strings.Trim(fs, split)
}

func (l *Logger) jsonOutput(level int, addStack bool, fields []interface{}) string {
	if l.level > level {
		return ""
	}

	msg := fieldsToString(fields)
	outmap := map[string]interface{}{
		levelKey: levelIDMap[level],
		timeKey:  now(),
		msgKey:   msg,
	}

	if addStack {
		outmap[stackstraceKey] = stack.TakeStacktrace(callerSkipOffset + 1)
	}
	if l.serviceCode > 0 {
		outmap[codekey] = l.serviceCode
	}

	js, _ := json.Marshal(outmap)
	output := string(js)
	fmt.Println(output)
	if l.outputFile != nil {
		line := output + "\n"
		if _, err := l.outputFile.WriteString(line); err != nil {
			panic(err)
		}
	}
	return output
}

func (l *Logger) stdOutput(level int, addStack bool, fields []interface{}) string {
	if l.level > level {
		return ""
	}

	msg := fieldsToString(fields)

	var stackstr string
	if addStack {
		stackstr = "\n" + stack.TakeStacktrace(callerSkipOffset) + "\n"
	}

	var codeStr string
	if l.serviceCode > 0 {
		codeStr = fmt.Sprintf(" service-code: %v", l.serviceCode)
	}

	output := fmt.Sprintf("%v [%v]%v %v%v", stdnow(), levelColorMap[level], codeStr, msg, stackstr)
	fmt.Println(output)

	return output
}

// func getStacktrace() string {
// 	funcptr, file, line, _ := runtime.Caller(callerSkipOffset)
// 	file = "." + strings.TrimPrefix(file, workingDir)
// 	return fmt.Sprintf("%v:%d func:%v", file, line, runtime.FuncForPC(funcptr).Name())
// }

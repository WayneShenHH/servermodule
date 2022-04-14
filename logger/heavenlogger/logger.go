package heavenlogger

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/WayneShenHH/servermodule/logger/constants"
	"github.com/WayneShenHH/servermodule/util/color"
	"github.com/WayneShenHH/servermodule/util/stack"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Logger implement of paradise
type Logger struct {
	fatalCallback func(msg string)
	level         int
	serviceCode   constants.ServiceCode
	outputFile    *os.File
	formatter     func(level int, msg, traceId, transactionId, spanId string, isStack bool) string
}

// New instance
func New(level, formatter string, code int) *Logger {
	// check parameter valid
	lv, exist := levelNameMap[level]
	if !exist {
		lv = infoLevel
	}

	l := &Logger{
		level:       lv,
		serviceCode: constants.ServiceCode(code),
	}

	switch formatter {
	case JsonFormatter:
		l.formatter = l.jsonOutput
	case ConsoleFormatter:
		l.formatter = l.consoleOutput
	default:
		l.formatter = l.jsonOutput
	}

	return l
}

// formatter list
const (
	JsonFormatter    = "json"
	ConsoleFormatter = "console"
)

// level list
const (
	debug1Level = -4
	debug2Level = -3
	debug3Level = -2
	debug4Level = -1
	debugLevel  = 0
	infoLevel   = 1
	warnLevel   = 2
	errorLevel  = 3
	fatalLevel  = 4
)

// level name list
const (
	Debug1Level = "debug1"
	Debug2Level = "debug2"
	Debug3Level = "debug3"
	Debug4Level = "debug4"
	DebugLevel  = "debug"
	InfoLevel   = "info"
	WarnLevel   = "warn"
	ErrorLevel  = "error"
	FatalLevel  = "fatal"
)

type ApmTraceId string
type ApmTransactionId string
type ApmSpanId string

var levelNameMap = map[string]int{
	Debug1Level: debug1Level,
	Debug2Level: debug2Level,
	Debug3Level: debug3Level,
	Debug4Level: debug4Level,
	DebugLevel:  debugLevel,
	InfoLevel:   infoLevel,
	WarnLevel:   warnLevel,
	ErrorLevel:  errorLevel,
	FatalLevel:  fatalLevel,
}

var levelIDMap = map[int]string{
	debug1Level: Debug1Level,
	debug2Level: Debug2Level,
	debug3Level: Debug3Level,
	debug4Level: Debug4Level,
	debugLevel:  DebugLevel,
	infoLevel:   InfoLevel,
	warnLevel:   WarnLevel,
	errorLevel:  ErrorLevel,
	fatalLevel:  FatalLevel,
}

var colorMap = map[int]string{
	debug1Level: color.Blue1.Add("DEBUG"),
	debug2Level: color.Blue2.Add("DEBUG"),
	debug3Level: color.Blue3.Add("DEBUG"),
	debug4Level: color.Blue4.Add("DEBUG"),
	debugLevel:  color.Blue.Add("DEBUG"),
	infoLevel:   color.Cyan.Add("INFO"),
	warnLevel:   color.Yellow.Add("WARN"),
	errorLevel:  color.Red.Add("ERROR"),
	fatalLevel:  color.Magenta.Add("FATAL"),
}

// key list for json-output field
const (
	timeKey          = "time"
	stackstraceKey   = "stacktrace"
	codekey          = "serviceCode"
	msgKey           = "msg"
	levelKey         = "level"
	traceIdKey       = "trace.id"
	transactionIdKey = "transaction.id"
	spanIdKey        = "span.id"
)

var (
	workingDir string
)

const (
	callerSkipOffset = 4 // logger.InfoCallStack & heavenlogger.InfoCallStack & heavenlogger.log | heavenlogger.logf & consoleOutput | jsonOutput
)

func init() {
	if dir, err := os.Getwd(); err == nil {
		workingDir = dir
	}
}

func now() string {
	return color.Green.Add(time.Now().UTC().Format(time.RFC3339))
}

func (l *Logger) getServiceCode() string {
	if l.serviceCode > 0 {
		return fmt.Sprintf(" service-code: %v", l.serviceCode)
	}

	return ""
}

func fieldsToString(fields []interface{}) (string, string, string, string) {
	fs, split := "", " "
	traceId, transactionId, spanId := "", "", ""
	for idx := range fields {
		switch val := fields[idx].(type) {
		case constants.ServiceCode:
			continue
		case error:
			fs += fmt.Sprint(val.Error(), split)
		case string:
			fs += fmt.Sprint(val, split)
		case ApmTraceId:
			traceId = string(val)
		case ApmTransactionId:
			transactionId = string(val)
		case ApmSpanId:
			spanId = string(val)
		default:
			js, _ := json.Marshal(val)

			fs += fmt.Sprint(string(js), split)
		}
	}

	return strings.Trim(fs, split), traceId, transactionId, spanId
}

func formatArguments(args []interface{}) {
	for idx := range args {
		val := reflect.ValueOf(args[idx])

		if val.Kind() == reflect.Ptr {
			val = reflect.Indirect(val)
		}

		if !val.IsValid() {
			continue
		}

		switch args[idx].(type) {
		case error:
			args[idx] = args[idx].(error).Error()
			continue
		}

		switch val.Interface().(type) {
		case error:
			args[idx] = val.Interface().(error).Error()
			continue
		}

		switch val.Kind() {
		case reflect.Struct, reflect.Array, reflect.Map, reflect.Slice:
			jstr, _ := json.Marshal(args[idx])
			args[idx] = string(jstr)
		}
	}
}

func (l *Logger) log(level int, isStack bool, args ...interface{}) string {
	if l.level > level {
		return ""
	}
	msg, traceId, transactionId, spanId := fieldsToString(args)

	o := l.formatter(level, msg, traceId, transactionId, spanId, isStack)
	fmt.Print(o)

	if l.outputFile != nil {
		line := o + "\n"
		if _, err := l.outputFile.WriteString(line); err != nil {
			fmt.Println(err)
		}
	}

	return o
}

func (l *Logger) logf(level int, format string, isStack bool, args ...interface{}) string {
	if l.level > level {
		return ""
	}
	formatArguments(args)
	msg := fmt.Sprintf(format, args...)

	o := l.formatter(level, msg, "", "", "", isStack)
	fmt.Print(o)

	if l.outputFile != nil {
		if _, err := l.outputFile.WriteString(o); err != nil {
			fmt.Println(err)
		}
	}

	return o
}

func (l *Logger) consoleOutput(level int, msg, traceId, transactionId, spanId string, isStack bool) string {
	code := l.getServiceCode()

	printField := "%v [%v]%v %v"
	printValue := []interface{}{now(), colorMap[level], code, msg}

	if traceId != "" {
		printField += " trace.id:%v transaction.id:%v"
		printValue = append(printValue, traceId, transactionId)
		if spanId != "" {
			printField += " span.id:%v"
			printValue = append(printValue, spanId)
		}
	}
	printField += "\n"

	if isStack {
		stack := stack.TakeStacktrace(callerSkipOffset)
		printField += "%v\n"
		printValue = append(printValue, stack)
	}
	return fmt.Sprintf(printField, printValue...)
}

func (l *Logger) jsonOutput(level int, msg, traceId, transactionId, spanId string, isStack bool) string {
	outmap := map[string]interface{}{
		levelKey: levelIDMap[level],
		timeKey:  time.Now().UTC().Format(time.RFC3339),
		msgKey:   msg,
	}

	if traceId != "" {
		outmap[traceIdKey] = traceId
		outmap[transactionIdKey] = transactionId
		if spanId != "" {
			outmap[spanIdKey] = spanId
		}
	}

	if isStack {
		outmap[stackstraceKey] = stack.TakeStacktrace(callerSkipOffset)
	}
	if l.serviceCode > 0 {
		outmap[codekey] = l.serviceCode
	}

	js := jsonMarshal(outmap)
	return js
}

func jsonMarshal(obj interface{}) string {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(obj)
	if err != nil {
		fmt.Println(err)
	}
	return buffer.String()
}

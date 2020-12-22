package constant

// ServiceCode 錯誤代碼型別，呼叫 gologger 方法若傳入此型別，將自動填入服務代碼欄位
type ServiceCode int

// ServiceCode list
const (
	Scheduler ServiceCode = 105
)

// Level of logger
type Level string

// LogFormatter of logger
type LogFormatter string

// LogFormatter list
const (
	JsonFormatter LogFormatter = "json"
	StdFormatter  LogFormatter = "std"
)

// Level list
const (
	Debug Level = "debug"
	Info  Level = "info"
	Warn  Level = "warn"
	Error Level = "error"
	Fatal Level = "fatal"
)

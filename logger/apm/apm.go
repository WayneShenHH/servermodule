package apm

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"runtime"
	"strings"
	"sync"

	"github.com/WayneShenHH/servermodule/logger"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmhttp"
	"go.elastic.co/apm/transport"
)

const (
	defaultAgentServerURL = "http://localhost:8200"
)

var tracer *apm.Tracer
var once sync.Once
var initialized bool

type Apm struct {
	context     context.Context
	transaction *apm.Transaction
	span        *apm.Span
}

func Init(serviceName, serviceEnvironment, urls string) {
	once.Do(func() {
		s, err := transport.NewHTTPTransport()
		if err != nil {
			panic(err)
		}

		s.SetServerURL(initServerURLs(urls)...)

		t, err := apm.NewTracerOptions(apm.TracerOptions{
			ServiceName:        serviceName,
			ServiceEnvironment: serviceEnvironment,
			Transport:          s,
		})
		if err != nil {
			logger.Warnf("Failed to initial apm tracer: %v", err)
			return
		}

		tracer = t
		initialized = true
	})

}

func Tracer() *apm.Tracer {
	return tracer
}

func initServerURLs(value string) []*url.URL {
	var urls []*url.URL
	for _, field := range strings.Split(value, ",") {
		field = strings.TrimSpace(field)
		if field == "" {
			continue
		}
		u, err := url.Parse(field)
		if err != nil {
			continue
		}
		urls = append(urls, u)
	}

	defaultServerURL, _ := url.Parse(defaultAgentServerURL)

	if len(urls) == 0 {
		urls = []*url.URL{defaultServerURL}
	}
	return urls
}

// StartTransaction 開始一個 transaction
func StartTransaction(name, transactionType string) *Apm {
	if !initialized {
		return nil
	}

	transaction := tracer.StartTransaction(name, transactionType)
	ctx := apm.ContextWithTransaction(context.TODO(), transaction)
	return &Apm{
		context:     ctx,
		transaction: transaction,
		span:        nil,
	}
}

// StartTransactionOptions 接續上一個 transaction 追蹤
func StartTransactionOptions(name, transactionType, traceparent string) *Apm {
	if !initialized {
		return nil
	}

	traceContext, err := apmhttp.ParseTraceparentHeader(traceparent)
	if err != nil {
		logger.Warnf("Failed to parse traceparent header: %v", err)
		return StartTransaction(name, transactionType)
	}
	opts := apm.TransactionOptions{
		TraceContext: traceContext,
	}
	transaction := tracer.StartTransactionOptions(name, transactionType, opts)
	ctx := apm.ContextWithTransaction(context.TODO(), transaction)
	return &Apm{
		context:     ctx,
		transaction: transaction,
		span:        nil,
	}
}

// End 結束 transaction
func (a *Apm) End() {
	a.transaction.End()
}

// GetContext
func (a *Apm) GetContext() context.Context {
	return a.context
}

// GetTransaction 取得 transaction
func (a *Apm) GetTransaction() *apm.Transaction {
	return a.transaction
}

// GetTraceContext 取得此鏈路的 traceContext
func (a *Apm) GetTraceContext() string {
	traceContext := a.transaction.TraceContext()
	return apmhttp.FormatTraceparentHeader(traceContext)
}

// StartSpan 記錄一筆 span
func (a *Apm) StartSpan(name, spanType string, parent *apm.Span) *apm.Span {
	span := a.transaction.StartSpan(name, spanType, parent)
	a.span = span
	return span
}

// getTraceId 取得 tracerId
func (a *Apm) getTraceId() string {
	return a.transaction.TraceContext().Trace.String()
}

// getTransactionId 取得 apm 的 transactionId
func (a *Apm) getTransactionId() string {
	return a.transaction.TraceContext().Span.String()
}

// getSpanId 取得 apm 的 spanId
func (a *Apm) getSpanId() string {
	span := a.span
	if span != nil {
		return span.TraceContext().Span.String()
	}
	return ""
}

func (a *Apm) getLogId() (string, string, string) {
	return a.getTraceId(), a.getTransactionId(), a.getSpanId()
}

// Error
func (a *Apm) Error(args ...interface{}) {
	msgString := fmt.Sprint(args...)
	traceId, transactionId, spanId := a.getLogId()
	logger.ApmError(traceId, transactionId, spanId, msgString)

	err := tracer.NewError(errors.New(msgString))
	if a.span != nil {
		err.SetSpan(a.span)
	} else {
		err.SetTransaction(a.transaction)
	}
	err.Send()
}

// Warn
func (a *Apm) Warn(args ...interface{}) {
	msgString := fmt.Sprint(args...)
	traceId, transactionId, spanId := a.getLogId()
	logger.ApmWarn(traceId, transactionId, spanId, msgString)
}

// Debug
func (a *Apm) Debug(args ...interface{}) {
	msgString := fmt.Sprint(args...)
	traceId, transactionId, spanId := a.getLogId()
	logger.ApmDebug(traceId, transactionId, spanId, msgString)
}

// Info
func (a *Apm) Info(args ...interface{}) {
	msgString := fmt.Sprint(args...)
	traceId, transactionId, spanId := a.getLogId()
	logger.ApmInfo(traceId, transactionId, spanId, msgString)
}

// Errorf
func (a *Apm) Errorf(format string, args ...interface{}) {
	msgString := fmt.Sprintf(format, args...)
	traceId, transactionId, spanId := a.getLogId()
	logger.ApmError(traceId, transactionId, spanId, msgString)

	err := tracer.NewError(errors.New(msgString))
	if a.span != nil {
		err.SetSpan(a.span)
	} else {
		err.SetTransaction(a.transaction)
	}
	err.Send()
}

// Warnf
func (a *Apm) Warnf(format string, args ...interface{}) {
	msgString := fmt.Sprintf(format, args...)
	traceId, transactionId, spanId := a.getLogId()
	logger.ApmWarn(traceId, transactionId, spanId, msgString)
}

// Debugf
func (a *Apm) Debugf(format string, args ...interface{}) {
	msgString := fmt.Sprintf(format, args...)
	traceId, transactionId, spanId := a.getLogId()
	logger.ApmDebug(traceId, transactionId, spanId, msgString)
}

// Debug1f
func (a *Apm) Debug1f(format string, args ...interface{}) {
	msgString := fmt.Sprintf(format, args...)
	traceId, transactionId, spanId := a.getLogId()
	logger.ApmDebug1f(traceId, transactionId, spanId, msgString)
}

// Infof
func (a *Apm) Infof(format string, args ...interface{}) {
	msgString := fmt.Sprintf(format, args...)
	traceId, transactionId, spanId := a.getLogId()
	logger.ApmInfo(traceId, transactionId, spanId, msgString)
}

// Fatalf
func (a *Apm) Fatalf(format string, args ...interface{}) {
	msgString := fmt.Sprintf(format, args...)
	traceId, transactionId, spanId := a.getLogId()
	logger.ApmFatalf(traceId, transactionId, spanId, msgString)
}

// RunFuncName
func RunFuncName() string {
	pc, _, _, ok := runtime.Caller(1)

	if !ok {
		return "can_not_get_function_name"
	}

	f := runtime.FuncForPC(pc)
	funcName := strings.Split(f.Name(), ".")
	return funcName[len(funcName)-1]
}

func TransactionFromContext(ctx context.Context) *apm.Transaction {
	return apm.TransactionFromContext(ctx)
}

func GetTraceparentHeader(ctx apm.TraceContext) string {
	return apmhttp.FormatTraceparentHeader(ctx)
}

func ContextWithTransaction(parent context.Context, t *apm.Transaction) context.Context {
	return apm.ContextWithTransaction(parent, t)
}

func ContextWithSpan(parent context.Context, s *apm.Span) context.Context {
	return apm.ContextWithSpan(parent, s)
}

func StartSpan(ctx context.Context, name, spanType string) (*apm.Span, context.Context) {
	return apm.StartSpan(ctx, name, spanType)
}

package protocol

type Timestamp uint64
type Balance int64
type ErrorCode int

type SequenceId = uint64 // MQ序列號
type TraceId = string    // 每個請求從客戶端開始一直到結束的唯一識別碼，用於追蹤所有經過的方法

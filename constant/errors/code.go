package errors

// 定義服務的錯誤代碼
const (
	DatabaseCollection   = 1000 // 取得 Collection 失敗
	DatabaseReadDocument = 1001 // 讀取 Document 失敗
	InputValueInvalid    = 1002 // 傳入的數值不合法
	DataIncompleted      = 1003 // 資料不完整
	NotImplement         = 1004 // 尚未實做
	DatabaseUpdate       = 1005 // 資料更新失敗
	DatabaseMetaEmptyID  = 1006 // 資料庫回傳空的 Meta ID
)

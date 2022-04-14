package errorcode

import "github.com/WayneShenHH/servermodule/protocol"

// 定義服務的錯誤代碼
const (
	DatabaseCollection   protocol.ErrorCode = 1000 // 取得 Collection 失敗
	DatabaseReadDocument protocol.ErrorCode = 1001 // 讀取 Document 失敗
	InputValueInvalid    protocol.ErrorCode = 1002 // 傳入的數值不合法
	DataIncompleted      protocol.ErrorCode = 1003 // 資料不完整
	NotImplement         protocol.ErrorCode = 1004 // 尚未實做
	DatabaseUpdate       protocol.ErrorCode = 1005 // 資料更新失敗
	DatabaseMetaEmptyID  protocol.ErrorCode = 1006 // 資料庫回傳空的 Meta ID
)

const (
	Transport protocol.ErrorCode = 3000 // 資料傳輸失敗
)

const Success protocol.ErrorCode = 0

package util

import (
	"bytes"
	"encoding/binary"

	"github.com/WayneShenHH/servermodule/protocol"
	"github.com/WayneShenHH/servermodule/protocol/errorcode"
)

// DecodeResponse 解析成 code & payload
type DecodeResponse struct {
	EventCode protocol.EventCode
	Data      []byte
}

// EncodeData 編碼內容，包含 code
type EncodeData struct {
	EventCode protocol.EventCode
	Payload   interface{}
}

// Decode 解析 websocket/mq 收到的封包
//
// @param data 來源的原始資料
//
// @return *DecodeResponse 解析完成的 code & payload (raw data)
//
// @return errorcode
func Decode(data []byte) (*DecodeResponse, protocol.ErrorCode) {
	dataLen := len(data)
	if dataLen < 2 {
		return nil, errorcode.DataIncompleted
	}

	eventCode := binary.LittleEndian.Uint16(data[0:2])
	bytes := data[2:]

	res := &DecodeResponse{
		EventCode: protocol.EventCode(eventCode),
		Data:      bytes,
	}

	return res, errorcode.Success
}

// Encode 將準備送出到 websocket/mq server 的訊息包裝成溝通好的格式
//
// @param f 訊息內容
//
// @return []byte 編碼後的位元組陣列
//
// @return errorcode
func Encode(data *EncodeData) ([]byte, protocol.ErrorCode) {
	payload, err := Marshal(data.Payload)
	if err != nil {
		return nil, errorcode.DataIncompleted
	}

	var buffer bytes.Buffer

	// event code
	eventCodeByte := make([]byte, 2)
	binary.LittleEndian.PutUint16(eventCodeByte, uint16(data.EventCode))
	buffer.Write(eventCodeByte)

	// payload
	buffer.Write(payload)

	return buffer.Bytes(), errorcode.Success
}

package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

//  Deocder Abc
type HsDecoder struct {
	Endian binary.ByteOrder
}

// Decode decode function
func (dr HsDecoder) Decode(data []byte) (Msg, error) {
	header := new(HsHeader)
	r := bytes.NewReader(data)
	binary.Read(r, binary.BigEndian, header)
	var body Msg
	switch header.Body_length {
	case 0x38:
		body = new(HsOrder)
	case 0x48:
		body = new(HsTrade)
	case 0x1e8:
		body = new(HsStockSnap)
	default:
		body = nil
	}
	if body == nil {
		return body, errors.New("Error msg type") // 主动不解析
	}
	err := binary.Read(r, dr.Endian, body)
	if err != nil {
		fmt.Println("failed to decode :", data)
		return body, err // 解析失败
	}
	return body, nil // 成功
}

type HsHeader struct {
	Msg_type         uint32
	Body_length      uint32
	Nsq_seq_num      uint64
	Version          byte
	Reserved         byte
	Market_type      uint16
	Sub_msg_function uint16
}

type HsOrder struct {
	Channel_no    uint16
	Appl_seq_num  uint64
	Security_id   [8]byte
	Price         int64
	Transact_time int64
	Order_qty     int32
	Side          byte
	Order_type    byte
	Reserved      uint16
	Checksum      int32
}

func (m HsOrder) ToString(recvTime int64) string {
	return fmt.Sprintf(
		"%s, %d, %d, %d, %d, %d, %d, %c",
		m.Security_id[:6], m.Channel_no, m.Appl_seq_num,
		m.Transact_time, recvTime, m.Price, m.Order_qty, m.Order_type)
}

type HsTrade struct {
	Channel_no         uint16
	Appl_seq_num       uint64
	Bid_appl_seq_num   uint64
	Offer_appl_seq_num uint64
	Security_id        [8]byte
	Last_price         uint64
	Transact_time      uint64
	Last_qty           uint32
	Exec_type          byte
	Reserved           [3]byte
	Checksum           int32
}

func (m HsTrade) ToString(recvTime int64) string {
	var rate float32 = 1e-6
	return fmt.Sprintf(
		"%s, %d, %d, %d, %d, %f, %d, %c",
		m.Security_id[:6], m.Channel_no, m.Appl_seq_num, m.Transact_time, recvTime, float32(m.Last_price)*rate, m.Last_qty, m.Exec_type)
}

type PriQty struct {
	Price uint64
	Qty   uint64
}

type HsStockSnap struct {
	Channel_no         uint16
	Orig_time          uint64
	Security_id        [8]byte
	Trading_phase_code uint64
	Prev_close_px      uint64
	Num_trades         uint64
	Total_volume_trade uint64
	Total_value_trade  uint64
	Last_price         uint64
	Open_price         uint64
	High_price         uint64
	Low_price          uint64
	Up_price           uint64
	Down_price         uint64
	BidPriQty          [10]PriQty
	Wa_bid_price       uint64
	Total_bid_quantity uint64
	AskPriQty          [10]PriQty
	Wa_ask_price       uint64
	Total_ask_quantity uint64
	Pre_iopv           uint64
	Iopv               uint64
	Checksum           int32
}

func (m HsStockSnap) ToString(recvTime int64) string {
	var rate float32 = 1e-6
	excgid := 2
	return fmt.Sprintf(
		"%s, %d, %d, %d, %f, %d, %f, %d, %f, %d, %f, %d, %d, %d, %f, %f, %f, %f, %d",
		m.Security_id[:6], excgid, m.Orig_time, recvTime,
		float32(m.BidPriQty[0].Price)*rate, m.BidPriQty[0].Qty, float32(m.AskPriQty[0].Price)*rate, m.AskPriQty[0].Qty,
		float32(m.BidPriQty[9].Price)*rate, m.BidPriQty[9].Qty, float32(m.AskPriQty[9].Price)*rate, m.AskPriQty[9].Qty,
		m.Total_bid_quantity, m.Total_ask_quantity,
		float32(m.Open_price)*rate, float32(m.High_price)*rate, float32(m.Low_price)*rate, float32(m.Last_price)*rate, m.Total_volume_trade)
}

type HsIndexSnap struct {
	Channel_no         uint16
	Orig_time          uint64
	Security_id        [8]byte
	Trading_phase_code uint64
	Prev_close_px      uint64
	Num_trades         uint64
	Total_volume_trade uint64
	Total_value_trade  uint64
	Prev_close_index   uint64
	Last_index         uint64
	Open_index         uint64
	High_index         uint64
	Low_index          uint64
	Close_index        uint64
	Checksum           int32
}

type HsTradeVol struct {
	Channel_no         uint16
	Orig_time          uint64
	Security_id        [8]byte
	Trading_phase_code uint64
	Prev_close_px      uint64
	Num_trades         uint64
	Total_volume_trade uint64
	Total_value_trade  uint64
	Stock_num          int32
	Checksum           int32
}

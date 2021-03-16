// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs C_AkMsg.go

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type AkDecoder struct {
	Endian binary.ByteOrder
}

func (dr AkDecoder) Decode(data []byte) (Msg, error) {
	header := new(AkHeader)
	r := bytes.NewReader(data)
	binary.Read(r, dr.Endian, header)
	var body Msg
	switch header.MessageType {
	case 1:
		body = new(AkSnap)
	case 3:
		body = new(AkIndex)
	case 4:
		body = new(AkTradeSse)
	case 5:
		body = new(AkEntrust)
	default:
		body = nil
	}
	if body == nil {
		return body, fmt.Errorf("Error msg type : %x", header.MessageType) // 主动不解析
	}
	err := binary.Read(r, dr.Endian, body)
	if err != nil {
		fmt.Println("failed to decode :", data)
		return body, err // 解析失败
	}
	return body, nil // 成功
}

type AkHeader struct {
	MessageType uint8
}
type AkBestOrder struct {
	ExchangeID  uint8
	SecurityID  [8]byte
	Side        uint8
	Number      uint8
	Volume      [50]uint64
	ChannelNo   uint32
	Mdstreamid  [3]int8
	SendingTime [6]int8
}

func (m AkBestOrder) ToString(recvTime int64) string {
	return fmt.Sprintf("NaN, %d", recvTime)
}

type AkIndex struct {
	Pad_cgo_0        [4]byte
	ExchangeID       uint8
	SecurityID       [8]byte
	Flag             uint16
	TradingPhaseCode [8]int8
	TimeStamp        uint64
	TradeTime        uint32
	Resv             [4]uint8
	PreClosePrice    uint64
	OpenPrice        uint64
	LastPrice        uint64
	HighPrice        uint64
	LowPrice         uint64
	ClosePrice       uint64
	TradeNum         uint64
	TotalVolume      uint64
	TotalValue       uint64
	ChannelNo        uint32
	Mdstreamid       [3]int8
	Pad_cgo_1        [4]byte
	SendingTime      [6]int8
}

func (m AkIndex) ToString(recvTime int64) string {
	return fmt.Sprintf(
		"%d, %d, %d, %s, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d",
		m.ExchangeID, m.TimeStamp, m.ChannelNo, m.SecurityID[:6],
		m.LastPrice, m.PreClosePrice, m.OpenPrice, m.HighPrice, m.LowPrice, m.ClosePrice, m.TradeNum, m.TotalVolume, m.TotalValue, recvTime)
}

type AkEntrust struct {
	Pad_cgo_0    [4]byte
	ExchangeID   uint8
	SecurityID   [8]byte
	Side         byte
	OrderType    byte
	ApplSeqNum   uint64
	TransactTime uint64
	Price        uint64
	Qty          uint64
	ChannelNo    uint32
	Mdstreamid   [3]byte
}

func (m AkEntrust) ToString(recvTime int64) string {
	return fmt.Sprintf(
		"%s, %d, %d, %d, %d, %d, %d, %c",
		m.SecurityID[:6], m.ChannelNo, m.ApplSeqNum,
		m.TransactTime, recvTime, m.Price, m.Qty, m.OrderType)
}

type AkTrade struct {
	Pad_cgo_0       [4]byte
	ExchangeID      uint8
	SecurityID      [8]byte
	ExecType        uint8
	TradeBSFlag     int8
	ApplSeqNum      uint64
	TransactTime    uint64
	LastPrice       uint64
	LastQty         uint64
	TradeMoney      uint64
	BidapplSeqnum   uint64
	OfferapplSeqnum uint64
	ChannelNo       uint32
	Mdstreamid      [3]byte
	Pad_cgo_1       [4]byte
	SendingTime     [6]int8
}

func (m AkTrade) ToString(recvTime int64) string {
	return fmt.Sprintf(
		"%s, %d, %d, %d, %d, %d, %d, %d",
		m.SecurityID[:6], m.ChannelNo, m.ApplSeqNum, m.TransactTime, recvTime, m.LastPrice, m.LastQty, m.ExecType)
}

type BidAskPriceQty struct {
	Price uint64
	Qty   uint64
}

type AkSnap struct {
	Pad_cgo_0        [4]byte
	ExchangeID       uint8
	SecurityID       [8]byte
	Resv             uint8
	Pad_cgo_1        [2]byte
	TradingPhaseCode [8]int8
	InstrumentStatus [5]int8
	Resv2            [2]uint8
	TimeStamp        uint64
	PreClosePrice    uint64
	NumTrades        uint64
	TotalVolumeTrade uint64
	TotalValueTrade  uint64
	LastPrice        uint64
	OpenPrice        uint64
	ClosePrice       uint64
	HighPrice        uint64
	LowPrice         uint64
	UpperlmtPrice    uint64
	LowerlmtPrice    uint64
	BidAvgPrice      uint64
	BidTotalQty      uint64
	AskAvgPrice      uint64
	AskTotalQty      uint64
	BidInfo          [10]BidAskPriceQty
	AskInfo          [10]BidAskPriceQty
	ChannelNo        uint32
	Mdstreamid       [3]int8
	Pad_cgo_2        [4]byte
	SendingTime      [6]int8
}

func (m AkSnap) ToString(recvTime int64) string {
	return fmt.Sprintf(
		"%s, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d",
		m.SecurityID[:6], m.ExchangeID, m.TimeStamp, recvTime,
		m.BidInfo[0].Price, m.BidInfo[0].Qty, m.AskInfo[0].Price, m.AskInfo[0].Qty,
		m.BidInfo[9].Price, m.BidInfo[9].Qty, m.AskInfo[9].Price, m.AskInfo[9].Qty,
		m.BidTotalQty, m.AskTotalQty,
		m.OpenPrice, m.HighPrice, m.LowPrice, m.LastPrice, m.TotalVolumeTrade)
}

//AkTradeSse 逐笔成交
type AkTradeSse struct {
	Sequence        uint32
	TradeBSFlag     uint8
	MsgSeqID        uint32
	TradeIndex      uint32
	SecurityID      [6]byte
	ChannelNo       uint16
	TransactTime    uint32
	LastPrice       uint32
	LastQty         uint64
	TradeMoney      uint64
	BidapplSeqnum   uint32
	OfferapplSeqnum uint32
}

//ToString 消息格式化输出
func (m AkTradeSse) ToString(recvTime int64) string {
	return fmt.Sprintf(
		"%s, %d, %d, %d, %d, %d, %d, %d",
		m.SecurityID[:6], m.ChannelNo, m.TradeIndex, m.TransactTime, recvTime, m.LastPrice, m.LastQty, m.TradeBSFlag)
}

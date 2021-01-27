// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs C_AkMsg.go

package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type AkDecoder struct {
	Endian binary.ByteOrder
}

func (dr AkDecoder) Decode(data []byte) (Msg, error) {
	msg := new(AkMsg)
	r := bytes.NewReader(data)
	binary.Read(r, dr.Endian, &msg.Head)
	switch msg.Head.MessageType {
	case 1:
		msg.Body = new(AkSnap)
		msg.Type = "Depth"
	case 2:
		msg.Body = new(AkBestOrder)
		msg.Type = "other"
	case 3:
		msg.Body = new(AkIndex)
		msg.Type = "other"
	case 4:
		msg.Body = new(AkTrade)
		msg.Type = "Tick"
	case 5:
		msg.Body = new(AkEntrust)
		msg.Type = "Tick"
	default:
		msg.Body = nil
	}
	if msg.Body == nil {
		return msg, errors.New("Error msg type")
	}
	binary.Read(r, dr.Endian, msg.Body)
	return msg, nil
}

type AkMsg struct {
	Head AkMsgHeader
	Type string
	Body interface{ ToString(int64) string }
}

func (msg *AkMsg) SaveType() string {
	return msg.Type
}

func (msg *AkMsg) ToString(recvTime int64) string {
	return msg.Body.ToString(recvTime)
}

type AkMsgHeader struct {
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

func (m *AkBestOrder) ToString(recvTime int64) string {
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

func (m *AkIndex) ToString(recvTime int64) string {
	return fmt.Sprintf(
		"%d, %d, %d, %s, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d",
		m.ExchangeID, m.TimeStamp, m.ChannelNo, m.SecurityID,
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

func (m *AkEntrust) ToString(recvTime int64) string {
	return fmt.Sprintf(
		"%s, %d, %d, %d, %d, %d, %d, %c",
		m.SecurityID, m.ChannelNo, m.ApplSeqNum,
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

func (m *AkTrade) ToString(recvTime int64) string {
	var rate float32 = 1e-6
	return fmt.Sprintf(
		"%s, %d, %d, %d, %d, %f, %d, %d",
		m.SecurityID, m.ChannelNo, m.ApplSeqNum, m.TransactTime, recvTime, float32(m.LastPrice)*rate, m.LastQty, m.ExecType)
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

func (m *AkSnap) ToString(recvTime int64) string {
	var rate float32 = 1e-6
	return fmt.Sprintf(
		"%s, %d, %d, %d, %f, %d, %f, %d, %f, %d, %f, %d, %d, %d, %f, %f, %f, %f, %d",
		m.SecurityID, m.ExchangeID, m.TimeStamp, recvTime,
		float32(m.BidInfo[0].Price)*rate, m.BidInfo[0].Qty, float32(m.AskInfo[0].Price)*rate, m.AskInfo[0].Qty,
		float32(m.BidInfo[9].Price)*rate, m.BidInfo[9].Qty, float32(m.AskInfo[9].Price)*rate, m.AskInfo[9].Qty,
		m.BidTotalQty, m.AskTotalQty,
		float32(m.OpenPrice)*rate, float32(m.HighPrice)*rate, float32(m.LowPrice)*rate, float32(m.LastPrice)*rate, m.TotalVolumeTrade)
}
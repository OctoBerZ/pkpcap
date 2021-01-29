package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

// RsDecoder Deocder Abc
type RsDecoder struct {
	Endian binary.ByteOrder
}

// Decode decode function
func (dr RsDecoder) Decode(data []byte) (Msg, error) {
	header := new(RsHeader)
	r := bytes.NewReader(data)
	binary.Read(r, dr.Endian, header)
	var body Msg
	switch header.MsgType {
	case 2:
		body = new(RsEntrust)
		//msg.Type = "Tick"
	case 3:
		body = new(RsTrade)
		//msg.Type = "Tick"
	case 4:
		body = new(RsSnap)
		//msg.Type = "Depth"
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

//RsHeader Rishon UDP data header
type RsHeader struct {
	MagicNum [2]byte //魔数
	MsgID    uint32  //消息ID
	MsgType  uint16  //消息类型
	Proto    [2]byte //协议类型
	Source   [2]byte //行情源
	Len      uint32  //消息长处
}

//RsIndex Rishon index data struct
type RsIndex struct {
	MsgID      uint64  //消息ID
	ExchangeID [2]byte //交易所ID
	OrigTime   int64   //交易所时间
	ChannelNO  uint16  //频道代码
	SecurityID [8]byte //证券代码
	LastPx     int64   //最新价格
	PreClosePx int64   //昨收价
	OpenPx     int64   //今开盘
	HighPx     int64   //最高价
	LowPx      int64   //最低价
	ClosePx    int64   //今收盘价
	NumTrades  int64   //成交笔数
	TotalVol   int64   //总成交量
	TotalValue int64   //总成交额
	FpgaTime   int64   //硬件时间戳
}

func (m RsIndex) ToString(recvTime int64) string {
	return fmt.Sprintf(
		"%s, %d, %d, %s, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d",
		m.ExchangeID, m.OrigTime, m.ChannelNO, m.SecurityID[:6],
		m.LastPx, m.PreClosePx, m.OpenPx, m.HighPx, m.LowPx, m.ClosePx, m.NumTrades, m.TotalVol, m.TotalValue, recvTime)
}

//RsEntrust Rishon entrust data struct
type RsEntrust struct {
	MsgID        uint64  //消息ID
	ExchangeID   [2]byte //交易所ID
	ChannelNO    uint16  //频道代码
	AppSeqNum    int64   //消息记录号
	SecurityID   [8]byte //证券代码
	Price        int64   //委托价格
	OrderQty     int64   //委托数量
	Side         byte    //买卖方向
	TransactTime int64   //委托时间
	OrdType      byte    //订单类别
	FpgaTime     int64   //硬件时间戳
}

func (m RsEntrust) ToString(recvTime int64) string {
	return fmt.Sprintf(
		"%s, %d, %d, %d, %d, %d, %d, %c",
		m.SecurityID[:6], m.ChannelNO, m.AppSeqNum,
		m.TransactTime, recvTime, m.Price, m.OrderQty, m.OrdType)
}

//RsTrade Rishon trade data struct
type RsTrade struct {
	MsgID          uint64  //消息ID
	ExchangeID     [2]byte //交易所ID
	ChannelNO      uint16  //频道代码
	AppSeqNum      int64   //消息记录号
	BidAppSeqNum   int64   //买方委托索引
	OfferAppSeqNum int64   //卖方委托索引
	SecurityID     [8]byte //证券代码
	LastPx         int64   //成交价格
	LastQty        int64   //成交数量
	ExecType       byte    //成交类别
	TransactTime   int64   //成交时间
	FpgaTime       int64   //硬件时间戳
}

func (m RsTrade) ToString(recvTime int64) string {
	return fmt.Sprintf(
		"%s, %d, %d, %d, %d, %d, %d, %c",
		m.SecurityID[:6], m.ChannelNO, m.AppSeqNum, m.TransactTime, recvTime, m.LastPx, m.LastQty, m.ExecType)
}

//RsSnap Rishon snapshot data struct
type RsSnap struct {
	MsgID        uint64  //消息ID
	ExchangeID   [2]byte //交易所ID
	OrigTime     int64   //交易所时间
	ChannelNO    uint16  //频道代码
	SecurityID   [8]byte //证券代码
	TradingPhase [8]byte //交易阶段代码

	BidPrice  [10]int64 //10档申买价
	AskPrice  [10]int64 //10档申卖价
	BidVolume [10]int64 //10档申买量
	AskVolume [10]int64 //10档申卖量

	PreClosePx        int64 //昨收价
	OpenPx            int64 //开盘价
	HighPx            int64 //最高价
	LowPx             int64 //最低价
	LastPx            int64 //最新价
	NumTrade          int64 //成交笔数
	TotalVolume       int64 //总成交量
	TotalValue        int64 //总成交额
	StockPER1         int64 //ll
	LastSubPreClosePx int64 //ll
	LastSubPreLastPx  int64 //ll
	TotalBidVol       int64 //委托买入总量
	TotalAskVol       int64 //委托卖出总量
	MaBidPx           int64 //加权平均买入价
	MaAskPx           int64 //加权平均卖出价
	PreIOPV           int64 //ss
	IOPV              int64 //lll
	UpperLimPx        int64 //涨停价
	LowerLimPx        int64 //跌停价
	FpgaTime          int64 //硬件时间戳
}

func (m RsSnap) ToString(recvTime int64) string {
	var excgid int
	switch m.ExchangeID[1] {
	case 0x5a: // b'Z'
		excgid = 2
	case 0x48: // b'H'
		excgid = 1
	}
	return fmt.Sprintf(
		"%s, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d",
		m.SecurityID[:6], excgid, m.OrigTime, recvTime,
		m.BidPrice[0], m.BidVolume[0], m.AskPrice[0], m.AskVolume[0],
		m.BidPrice[9], m.BidVolume[9], m.AskPrice[9], m.AskVolume[9],
		m.TotalBidVol, m.TotalAskVol,
		m.OpenPx, m.HighPx, m.LowPx, m.LastPx, m.TotalVolume)
}

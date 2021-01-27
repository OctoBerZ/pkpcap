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
	msg := new(RsMsg)
	r := bytes.NewReader(data)
	binary.Read(r, dr.Endian, &msg.Head)
	switch msg.Head.MsgType {
	case 1:
		msg.Body = new(RsIndex)
		msg.Type = ""
	case 2:
		msg.Body = new(RsEntrust)
		msg.Type = "Tick"
	case 3:
		msg.Body = new(RsTrade)
		msg.Type = "Tick"
	case 4:
		msg.Body = new(RsSnap)
		msg.Type = "Depth"
	default:
		msg.Body = nil
	}
	if msg.Body == nil {
		return msg, errors.New("Error msg type")
	}
	binary.Read(r, dr.Endian, msg.Body)
	return msg, nil

}

//RsMsg Abc
type RsMsg struct {
	Head RsMsgHeader
	Type string
	Body interface{ ToString(int64) string }
}

func (msg *RsMsg) SaveType() string {
	return msg.Type
}

func (msg *RsMsg) ToString(recvTime int64) string {
	return msg.Body.ToString(recvTime)
}

//RsMsgHeader Rishon UDP data header
type RsMsgHeader struct {
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

func (m *RsIndex) ToString(recvTime int64) string {
	return fmt.Sprintf(
		"%s, %d, %d, %s, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d",
		m.ExchangeID, m.OrigTime, m.ChannelNO, m.SecurityID,
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

func (m *RsEntrust) ToString(recvTime int64) string {
	return fmt.Sprintf(
		"%s, %d, %d, %d, %d, %d, %d, %c",
		m.SecurityID, m.ChannelNO, m.AppSeqNum,
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

func (m *RsTrade) ToString(recvTime int64) string {
	var rate float32 = 1e-6
	return fmt.Sprintf(
		"%s, %d, %d, %d, %d, %f, %d, %c",
		m.SecurityID, m.ChannelNO, m.AppSeqNum, m.TransactTime, recvTime, float32(m.LastPx)*rate, m.LastQty, m.ExecType)
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

func (m *RsSnap) ToString(recvTime int64) string {
	var rate float32 = 1e-6
	return fmt.Sprintf(
		"%s, %s, %d, %d, %f, %d, %f, %d, %f, %d, %f, %d, %d, %d, %f, %f, %f, %f, %d",
		m.SecurityID, m.ExchangeID, m.OrigTime, recvTime,
		float32(m.BidPrice[0])*rate, m.BidVolume[0], float32(m.AskPrice[0])*rate, m.AskVolume[0],
		float32(m.BidPrice[9])*rate, m.BidVolume[9], float32(m.AskPrice[9])*rate, m.AskVolume[9],
		m.TotalBidVol, m.TotalAskVol,
		float32(m.OpenPx)*rate, float32(m.HighPx)*rate, float32(m.LowPx)*rate, float32(m.LastPx)*rate, m.TotalVolume)
}

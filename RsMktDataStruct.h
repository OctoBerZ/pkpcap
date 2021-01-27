#ifndef _RS_API_STRUCT_H_
#define _RS_API_STRUCT_H_
#include <stdint.h>
#include <cstring>

#pragma pack(push)
#pragma pack (1)

//指数行情快照
//type == 1
typedef struct _SECU_MARKET_DATA
{
	uint64				MsgID;
	char				ExchangeID[2];
	int64				OrigTime;
	uint16				ChannelNo;
	char				SecurityID[8];
	int64				LastPrice;
	int64				PreClosePrice;
	int64				OpenPrice;///今开盘
	int64				HighPrice;///最高价
	int64				LowPrice;///最低价
	int64				ClosePrice;///今收盘
	int64				NumTrades;///成交笔数
	int64				TotalVolumeTrade;///成交总量
	int64				TotalValueTrade;///成交总金额
	int64  				FpgaTime;
}SecuIndexDataField;

///逐笔委托数据信息
//type == 2
typedef struct	_SECU_ENTRUST_DATA
{
	uint64				MsgID;///消息ID
	char				ExchangeID[2];///交易所
	uint16				ChannelNo;///频道代码
	int64				ApplSeqNum;///消息记录号
	char				SecurityID[8];///证券代码
	int64				Price;///委托价格
	int64				OrderQty;///委托数量
	char				Side;///买卖方向
	int64				TransactTime;///委托时间
	char				OrdType;///订单类别
	int64  				FpgaTime;
}SecuEntrustDataField;

///逐笔成交数据信息
//type == 3
typedef struct	_SECU_TRADE_DATA
{
	uint64				MsgID;///消息ID
	char				ExchangeID[2];///交易所
	uint16				ChannelNo;///频道代码
	int64				ApplSeqNum;///消息记录号
	int64				BidApplSeqNum;///买方委托索引
	int64				OfferApplSeqNum;///卖方委托索引
	char				SecurityID[8];///证券代码
	int64				LastPx;///成交价格
	int64				LastQty;///成交数量
	char				ExecType;///成交类别
	int64				TransactTime;///成交时间
	int64  				FpgaTime;
}SecuTradeDataField;

////盘口
typedef struct _SNAP_SHOT_
{
	uint64		MsgID;
	char		ExchangeID[2];
	int64		OrigTime;
	uint16		ChannelNo;
	char		SecurityID[8];
	char		TradingPhaseCode[8];

	int64		BidPrice[10];
	int64		AskPrice[10];
	int64		BidVolume[10];
	int64		AskVolume[10];

	int64		PrevClosePx;
	int64		OpenPrice;
	int64		HighPrice;
	int64		LowPrice;
	int64		LastPrice;
	int64		NumTrades;
	int64		TotalVolumeTrade;
	int64		TotalValueTrade;
	int64		StockPER1;
	int64		LastSubPrevClosePrice;
	int64		LastSubPrevLastPrice;
	int64		TotalBidVolume;
	int64		TotalAskVolume;
	int64		MaBidPrice;
	int64		MaAskPrice;
	int64		PreIOPV;
	int64		IOPV;
	int64		UpperLimitPrice;
	int64		LowerLimitPrice;
	int64  		FpgaTime;
}SecuSnapshotDataField;
#pragma pack(pop)
#endif

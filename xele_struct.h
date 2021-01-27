/*****************************************************************************
艾科朗克行情结构体定义，字节序为大端
******************************************************************************/
#ifndef MD_STRUCT_H
#define MD_STRUCT_H
#include "stdint.h"

#define MSG_TYPE_SNAPSHOT               0x1   //快照 
#define MSG_TYPE_BEST_ORDERS            0x2   //订单明细，最多揭示50笔
#define MSG_TYPE_INDEX                  0x3   //指数行情
#define MSG_TYPE_TRADE                  0x4   //逐笔成交
#define MSG_TYPE_ORDER                  0x5   //逐笔委托

#define SNAPSHOT_LEVEL                  10
#define BEST_ORDERS_LEVEL               50

#pragma pack(1)

struct BidAskPriceQty {
  uint64_t price;    //申买、申卖价格，实际值深交除以1000000，上交除以1000
  uint64_t qty;      //申买、申卖数量，实际值深交除以100，上交除以1000
};

struct MsgHead {
  uint8_t  messageType;    //消息类型，快照为0x1
  uint32_t sequence;       //udp输出包序号，从1开始
  uint8_t  exchangeID;     //交易所id，上交所：1，深交所：2
  char     securityID[8];  //证券代码
};

/*
*行情快照
*/
struct MarketDataSnapshot {
  uint8_t  messageType;    //消息类型，快照为0x1
  uint32_t sequence;       //udp输出包序号，从1开始
  uint8_t  exchangeID;     //交易所id，上交所：1，深交所：2
  char     securityID[8];  //证券代码
  uint8_t  resv;           //保留字段
  /*
  字段有效标识:
  bit0:最近价有效，bit1:开盘价有效，bit2:最高价有效，bit3:最低价有效，bit4:涨停价有效，
  bit5:跌停价有效，bit6:买入委托数量加权平均价/总数量有效，bit7:卖出委托数量加权平均价/总数量有效，bit8:今日收盘价，深交默认填0，其余位保留
  */
  uint16_t flag;
  /*
  产品所处的交易阶段代码
  深交:
  第0 位：S=启动（开市前）,O=开盘集合竞价,T=连续竞价,B=休市,C=收盘集合竞价,E=已闭市,H=临时停牌,A=盘后交易,V=波动性中断
  第1 位：0=正常状态,1=全天停牌
  -------------------------------------------------
  上交L2
  该字段为8位字符串，左起每位表示特定的含义，无定义则填空格。
  第1位：‘S’表示启动（开市前）时段，‘C’表示开盘集合竞价时段，‘T’表示连续交易时段，‘E’表示闭市时段，‘P’表示产品停牌，
        ‘M’表示可恢复交易的熔断时段（盘中集合竞价），‘N’表示不可恢复交易的熔断时段（暂停交易至闭市），‘U’表示收盘集合竞价时段。
  第2位：‘0’表示此产品不可正常交易，‘1’表示此产品可正常交易，无意义填空格。
  第3位：‘0’表示未上市，‘1’表示已上市。
  第4位：‘0’表示此产品在当前时段不接受订单申报，‘1’ 表示此产品在当前时段可接受订单申报。无意义填空格
  */
  char tradingPhaseCode[8];
  /*
  当前品种交易状态（仅上交所L2有效）
  详见上交所接口说明文档
  */
  char instrumentStatus[5];
  uint8_t resv2[2];
  /*
  沪深时间戳
  深交L2
  origtime 数据生成时间（切片时间），精确到毫秒
  -------------------------------------------------------------
  上交L2
  DataTimeStamp 最新订单时间（秒）143025 表示 14:30:25
  */
  uint64_t timeStamp;
  uint64_t preClosePrice;                    //昨收价（来源消息头),实际值深交除以10000，上交除以1000
  uint64_t numTrades;                        //总成交笔数
  uint64_t totalVolumeTrade;                 //总成交量,实际值深交除以100，上交除以1000
  uint64_t totalValueTrade;                  //总成交金额，实际值深交除以10000，上交除以100000
  uint64_t lastPrice;                        //最近价，实际值深交除以1000000，上交除以1000
  uint64_t openPrice;                        //开盘价，实际值深交除以1000000，上交除以1000
  uint64_t closePrice;                       //今日收盘价（仅上交所L2有效），实际值上交除以1000
  uint64_t highPrice;                        //最高价，实际值深交除以1000000，上交除以1000
  uint64_t lowPrice;                         //最低价，实际值深交除以1000000，上交除以1000
  uint64_t upperlmtPrice;                    //涨停价，实际值深交除以1000000，上交除以1000
  uint64_t lowerlmtPrice;                    //跌停价，实际值深交除以1000000，上交除以1000
  uint64_t bidAvgPrice;                      //买入委托数量加权平均价，实际值深交除以1000000，上交除以1000
  uint64_t bidTotalQty;                      //买入委托总数量，实际值深交除以100，上交除以1000
  uint64_t askAvgPrice;                      //卖出委托数量加权平均价，实际值深交除以1000000，上交除以1000
  uint64_t askTotalQty;                      //卖出委托总数量，实际值深交除以100，上交除以1000
  struct BidAskPriceQty bidInfo[SNAPSHOT_LEVEL];    //申买信息
  struct BidAskPriceQty askInfo[SNAPSHOT_LEVEL];    //申卖信息
  uint32_t channelNo;                        //频道代码（只有深交所有）
  char mdstreamid[3];                        //行情类别（只有深交所有）
  uint32_t msgSeqID;                         //消息序号（只有上交所有）
  char SendingTime[6];                       //消息报文FIX head中时间戳（只有上交所有）
};

/*
订单明细，最多揭示50笔
*/
struct BestOrders {
  uint8_t  messageType;               //消息类型，订单明细为0x2
  uint32_t sequence;                  //udp输出包序号，从1开始
  uint8_t  exchangeID;                //交易所id，上交所：1，深交所：2
  char     securityID[8];             //证券代码
  /*
  沪深时间戳
  深交L2
  origtime 数据生成时间（切片时间），精确到毫秒
  -------------------------------------------------------------
  上交L2
  DataTimeStamp 最新订单时间（秒）143025 表示 14:30:25
  */
  uint64_t timeStamp;
  uint8_t side;                        //买卖标识:买：1，卖：2
  uint64_t price;                      //委托价格，实际值深交除以1000000，上交除以1000
  uint64_t orders;                     //申买/卖数量,实际值深交除以100，上交除以1000
  uint8_t number;                      //明细个数
  uint64_t volume[BEST_ORDERS_LEVEL];  //订单明细，最多揭示50笔,实际值深交除以100，上交除以1000
  uint32_t channelNo;                  //频道代码（只有深交所有）
  char mdstreamid[3];                  //行情类别（只有深交所有）
  uint32_t msgSeqID;                   //消息序号（只有上交所有）
  char SendingTime[6];                 //消息报文FIX head中时间戳（只有上交所有）
};

/*
指数行情
*/
struct Index {
  uint8_t  messageType;               //消息类型，指数为0x3
  uint32_t sequence;                  //udp输出包序号，从1开始
  uint8_t  exchangeID;                //交易所id，上交所：1，深交所：2
  char     securityID[8];             //证券代码
  /*
  字段有效标识:
  0bit:前收盘指数,1bit:今开盘指数,2bit:最新指数,3bit:最高指数,4bit:最低指数,
  5bit:今日收盘指数,6bit:成交笔数，上交所默认填0,7bit:成交总量,8bit:成交总金额,9bit:成交时间，深交所默认填0,
  10bit:产品所处的交易阶段代码，上交所默认填0,其余位保留
  */
  uint16_t  flag;
  /*
  产品所处的交易阶段代码(仅深交所有)
  第0 位：
  S=启动（开市前）,O=开盘集合竞价,T=连续竞价,B=休市,C=收盘集合竞价,
  E=已闭市,H=临时停牌,A=盘后交易,V=波动性中断
  第1 位：
  0=正常状态,1=全天停牌
  */
  char tradingPhaseCode[8]; 
  /*
  沪深时间戳
  深交L2
  origtime 数据生成时间（切片时间），精确到毫秒
  -------------------------------------------------------------
  上交L2
  DataTimeStamp 最新订单时间（秒）143025 表示 14:30:25
  */
  uint64_t timeStamp;
  uint32_t tradeTime;                 //成交时间（只有上交所有）
  uint8_t resv[4];                    //保留字段

  uint64_t preClosePrice;             //前盘指数（来源扩展字段）,实际值深交除以1000000，上交除以100000
  uint64_t openPrice;                 //开盘指数,实际值深交除以1000000，上交除以100000
  uint64_t lastPrice;                 //最新指数,实际值深交除以1000000，上交除以100000

  uint64_t highPrice;                 //最高指数,实际值深交除以1000000，上交除以100000
  uint64_t lowPrice;                  //最低指数,实际值深交除以1000000，上交除以100000
  uint64_t closePrice;                //今日收盘指数,实际值深交除以1000000，上交除以100000
  uint64_t tradeNum;                  //成交笔数（只有深交所有）
  uint64_t totalVolume;               //成交总量,实际值深交除以100，上交除以100000
  uint64_t totalValue;                //成交总金额,实际值深交除以10000，上交除以10

  uint32_t channelNo;                 //频道代码（只有深交所有）
  char mdstreamid[3];                 //行情类别（只有深交所有）
  uint32_t msgSeqID;                  //消息序号（只有上交所有）
  char SendingTime[6];                //消息报文FIX head中时间戳（只有上交所有）
};

/*
逐笔成交
*/
struct Trade {
  uint8_t  messageType;               //消息类型，逐笔成交为0x4
  uint32_t sequence;                  //udp输出包序号，从1开始
  uint8_t  exchangeID;                //交易所id，上交所：1，深交所：2
  char     securityID[8];             //证券代码
  uint8_t  execType;                  //成交类别 ： 0x1--撤销；0x2--成交（上交所默认成交）
  char tradeBSFlag;                   //内外盘标志（仅上交所有，深交填0） ：B-外盘，主动买，S-内盘，主动买，N-未知
  uint64_t applSeqNum;                //深交所：消息记录号，上交所：成交序号
  uint64_t transactTime;              //成交时间
  uint64_t lastPrice;                 //成交价格,实际值深交除以10000，上交除以1000
  uint64_t lastQty;                   //成交数量,实际值深交除以100，上交除以1000
  uint64_t tradeMoney;                //成交金额（仅上交所有，深交填0）,实际值除以100000

  uint64_t bidapplSeqnum;             //买方委托索引
  uint64_t offerapplSeqnum;           //卖方委托索引
  
  uint32_t channelNo;                 //频道代码（只有深交所有）
  char mdstreamid[3];                 //行情类别（只有深交所有）
  uint32_t msgSeqID;                  //消息序号（只有上交所有）
  char SendingTime[6];                //消息报文FIX head中时间戳（只有上交所有）
};


/*
逐笔委托,只有深交所有
*/
struct Order {
  uint8_t  messageType;               //消息类型，逐笔委托为0x5
  uint32_t sequence;                  //udp输出包序号，从1开始
  uint8_t  exchangeID;                //交易所id，上交所：1，深交所：2
  char     securityID[8];             //证券代码
  char side;                          //买卖方向：1=买，2=卖，G=借入，F=出借
  char  orderType;                    //订单类别：1=市价，2=限价，U=本方最优
  uint64_t applSeqNum;                //消息记录号
  uint64_t transactTime;              //委托时间
  uint64_t price;                     //价格,实际值除以10000
  uint64_t qty;                       //数量，实际值除以100

  uint32_t channelNo;                 //频道代码
  char mdstreamid[3];                 //行情类别
};

#pragma pack(8)

#endif

package dataStruct

type Student struct {
	Name  string
	Age   int
	Score float64
}
type TenBlockHeader struct {
	CreateTimestamp string `protobuf:"bytes,1,opt,name=CreateTimestamp,proto3" json:"CreateTimestamp,omitempty"` //创建时间戳
	KeyId           string `protobuf:"bytes,2,opt,name=keyId,proto3" json:"keyId,omitempty"`                     //当前区块的key值
	PreBlockHash    string `protobuf:"bytes,3,opt,name=PreBlockHash,proto3" json:"PreBlockHash,omitempty"`       //前一个hash区块hash值
	BlockHash       string `protobuf:"bytes,4,opt,name=BlockHash,proto3" json:"BlockHash,omitempty"`             //前一个hash区块hash值
	BlockHeight     int64  `protobuf:"varint,5,opt,name=BlockHeight,proto3" json:"BlockHeight,omitempty"`        //区块高度
	BlockType       string `protobuf:"bytes,6,opt,name=BlockType,proto3" json:"BlockType,omitempty"`             //区块类型
	LedgerType      string `json:"LedgerType"`
}

//创始块
type GenesisBlock struct {
	//2021-05-17_LedgerType_BlockType
	KeyId                string `json:"key"`                  //通过该字段，获取当前创世区块
	Height               int64  `json:"height"`               //创世区块就是区块高位为1
	GenesisBlockHash     string `json:"genesisBlockHash"`     //创世区块哈希值
	DataCounts           int32  `json:"dataCounts"`           //数据交易记录数
	DataSize             int64  `json:"dataSize"`             //总数据量大小
	ChildBlockCount      int32  `json:"childBlockCount"`      //子块数量
	CreateTimestamp      string `json:"createTimestamp"`      //创建时间戳
	UpdateTimestamp      string `json:"updateTimestamp"`      //更新时间戳
	CumulativeBlock      int64  `json:"cumulativeBlock"`      //累计区块总数
	Version              string `json:"version"`              //创世区块版本号
	BlockChainType       string `json:"blockChainType"`       //目前主要三条链标示不同的链
	CreateChainTimestamp string `json:"createChainTimestamp"` //链创建时间
	LedgerType           string `json:"ledgerType"`           //目前主要三条链标示不同的链
	CumulativeValue      int64  `json:"cumulativeValue"`      //累计价值量
	CumulativeNode       int64  `json:"cumulativeNode"`       //累计参与终端数
	CumulativeUser       int64  `json:"cumulativeUser"`       //累计用户数
	GroupMasterNodeCount int32  `json:"GroupMasterNodeCount"` //集群master节点数量
	GroupSlaveNodeCount  int32  `json:"GroupSlaveNodeCount"`  //集群slave节点数量
}

//存证结构体
type DataReceipt struct {
	EntityId            string   `json:"entityId"`
	Timestamp           string   `json:"timeStamp" validate:"required"`
	KeyId               string   `json:"keyId" validate:"required"`
	ReceiptValue        float64  `json:"receiptValue"`
	Version             string   `json:"version"`
	UserName            string   `json:"userName"`
	OperationType       string   `json:"operationType"`
	DataType            string   `json:"dataType" validate:"required"`
	ServiceType         string   `json:"serviceType"`
	FileName            string   `json:"fileName"`
	FileSize            float64  `json:"fileSize"`
	FileHash            string   `json:"fileHash"`
	Uri                 string   `json:"uri"`
	ParentKeyId         string   `json:"parentKeyId"`
	AttachmentFileUris  []string `json:"attachmentFileUris"`
	AttachmentTotalHash string   `json:"attachmentTotalHash"`
}

// 实时交易记录
type Transaction struct {
	Timestamp     string  `json:"timeStamp" validate:"required"`
	EntityId      string  `json:"entityId"`
	TransactionId string  `json:"transactionId" validate:"required"`
	Initiator     string  `json:"initiator"`
	Recipient     string  `json:"recipient"`
	TxAmount      float64 `json:"txAmount"`
	DataType      string  `json:"dataType" validate:"required"`
	ServiceType   string  `json:"serviceType"`
	Remark        string  `json:"remark"`
}

//普通块 增强块 天块
type MinuteBlock struct {
	CreateTimestamp  string `json:"createTimestamp"`  //创建时间戳
	Key              string `json:"key"`              //通过该字段，获取当前区块
	BlockHeight      int64  `json:"blockHeight"`      //区块高度
	DataType         string `json:"dataType"`         //数据类型
	DataValue        string `json:"dataValue"`        //数据价值
	UpdateTimestamp  string `json:"updateTimestamp"`  //更新时间戳
	DataHash         string `json:"dataHash"`         //数据哈希值
	BlockHash        string `json:"blockHash"`        //区块哈希值
	PreBlockHash     string `json:"preBlockHash"`     //前一个区块hash值
	Nonce            int32  `json:"nonce"`            //nonce 值
	Target           int32  `json:"target"`           //目标值
	CurrentDataCount int64  `json:"currentDataCount"` //当前数据记录量
	CurrentDataSize  int64  `json:"currentDataSize"`  //当前数据大小
	Version          string `json:"version"`          //版本号
	BlockType        string `json:"blockType"`        //区块类型
	LedgerType       string `json:"ledgerType"`       //账本类型

	DataReceipts []DataReceipt `json:"dataReceipts"` //存证元数据
	Transactions []Transaction `json:"transactions"` //存证元数据
}
type TenMinuteBlock struct {
	CreateTimestamp string        `json:"createTimestamp"` //创建时间戳
	Key             string        `json:"key"`             //通过该字段，获取当前区块
	PreBlockHash    string        `json:"preBlockHash"`    //前一个区块hash值
	BlockHash       string        `json:"blockHash"`       //区块哈希值
	BlockHeight     int64         `json:"blockHeight"`     //区块高度
	BlockType       string        `json:"blockType"`       //区块类型
	LedgerType      string        `json:"ledgerType"`      //账本类型
	Blocks          []MinuteBlock `json:"blocks"`
}
type DailyBlock struct {
	CreateTimestamp string           `json:"createTimestamp"` //创建时间戳
	Key             string           `json:"key"`             //通过该字段，获取当前区块
	PreBlockHash    string           `json:"preBlockHash"`    //前一个区块hash值
	BlockHash       string           `json:"blockHash"`       //区块哈希值
	BlockHeight     int64            `json:"blockHeight"`     //区块高度
	BlockType       string           `json:"blockType"`       //区块类型
	LedgerType      string           `json:"ledgerType"`      //账本类型
	Blocks          []TenMinuteBlock `json:"blocks"`
}

type MinuteData struct {
	//MD数据key 2021-04-20 : MD : 账本类型 : 分钟索引 : 节点ID
	Key                                string   `json:"key"`                  //
	ReceiptTimeStampLedgerTypeKeyId    []string `json:"receiptTimeStamp"`     //存证数据 2021-04-20 17:18:10.123 # 账本类型 # KeyId
	TransactionTimeStampLedgerTypeTxId []string `json:"transactionTimeStamp"` //交易 2021-04-20 17:18:10.123 # 账本类型 # TransactionId
}

type Services struct {
	LeaderConfig    EtcdUrl   `yaml:"leaderConfig"`
	FollowersConfig []EtcdUrl `yaml:"followersConfig"`
}
type EtcdUrl struct {
	HraftAddress string `yaml:"hraft_grpc_address"`
	BlockAddress string `yaml:"block_grpc_address"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
}

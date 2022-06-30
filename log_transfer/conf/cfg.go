package conf

// LogTransferCfg is a struct that contains two fields, Kafka and ES, which are both structs themselves.
// @property {Kafka}  - `Kafka`: The Kafka properties.
// @property {ES}  - `Kafka`: The Kafka properties.
type LogTransferCfg struct {
	KafkaCfg `ini:"kafka"`
	ESCfg    `ini:"es"`
}

// KafkaCfg is a struct with two fields, Address and Topic, both of type string.
// @property {string} Address - The address of the Kafka server.
// @property {string} Topic - The topic to which the message will be sent.
type KafkaCfg struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

// ESCfg is a struct with a single field, `Address`, which is a string.
// @property {string} Address - The address of the Elasticsearch server.
type ESCfg struct {
	Address  string `ini:"address"`
	ChanSize int    `ini:"chanSize"`
	Nums     int    `ini:"nums"`
}

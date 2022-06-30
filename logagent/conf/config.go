package conf

// `AppConf` is a struct that contains two fields, `KafkaConf` and `EtcdConf`, which are both structs.
// @property {KafkaConf}  - `KafkaConf`: The configuration of Kafka
// @property {EtcdConf}  - `KafkaConf`: The configuration of Kafka
type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
}

// KafkaConf is a struct that contains a string and an int.
// @property {string} Address - Kafka address
// @property {int} ChanMaxSize - The maximum number of messages that can be stored in the channel.
type KafkaConf struct {
	Address     string `ini:"address"`
	ChanMaxSize int    `ini:"chan_max_size"`
}

// EtcdConf is a struct with three fields, each of which is a string of type string.
// @property {string} Address - The address of the etcd server
// @property {string} Key - The key of the log file in etcd
// @property {int} Timeout - The timeout for the connection to etcd.
type EtcdConf struct {
	Address string `ini:"address"`
	Key     string `ini:"collect_log_key"`
	Timeout int    `ini:"timeout"`
}

//---------------- unused ↓ -----------------
// `TaillogConf` is a struct with a single field, `FileName`, which is a string.
// @property {string} FileName - The name of the log file to be tailed.
type TaillogConf struct {
	FileName string `ini:"filename"`
}

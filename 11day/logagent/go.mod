module learning.goland.com/studyGoland/11day/logagent

go 1.18

require (
	gopkg.in/ini.v1 v1.66.6
	learning.goland.com/studyGoland/11day/logagent/conf v0.0.0-00010101000000-000000000000
	learning.goland.com/studyGoland/11day/logagent/etcd v0.0.0-00010101000000-000000000000
	learning.goland.com/studyGoland/11day/logagent/kafka v0.0.0-00010101000000-000000000000
	learning.goland.com/studyGoland/11day/logagent/taillog v0.0.0-00010101000000-000000000000
)

replace learning.goland.com/studyGoland/11day/logagent/kafka => ./kafka

replace learning.goland.com/studyGoland/11day/logagent/taillog => ./taillog

replace learning.goland.com/studyGoland/11day/logagent/conf => ./conf

replace learning.goland.com/studyGoland/11day/logagent/etcd => ./etcd

require (
	github.com/Shopify/sarama v1.34.0 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/eapache/go-resiliency v1.2.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.0.0 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.2 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	go.etcd.io/etcd/api/v3 v3.5.4 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.4 // indirect
	go.etcd.io/etcd/client/v3 v3.5.4 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292 // indirect
	golang.org/x/net v0.0.0-20220607020251-c690dde0001d // indirect
	golang.org/x/sys v0.0.0-20220608164250-635b8c9b7f68 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220608133413-ed9918b62aac // indirect
	google.golang.org/grpc v1.47.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
)

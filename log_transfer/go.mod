module learning.goland.com/studyGoland/log_transfer

go 1.18

replace learning.goland.com/studyGoland/log_transfer/conf => ./conf

replace learning.goland.com/studyGoland/log_transfer/kafka => ./kafka

replace learning.goland.com/studyGoland/log_transfer/es => ./es

require (
	gopkg.in/ini.v1 v1.66.6
	learning.goland.com/studyGoland/log_transfer/conf v0.0.0-00010101000000-000000000000
	learning.goland.com/studyGoland/log_transfer/es v0.0.0-00010101000000-000000000000
	learning.goland.com/studyGoland/log_transfer/kafka v0.0.0-00010101000000-000000000000
)

require (
	github.com/Shopify/sarama v1.34.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/eapache/go-resiliency v1.2.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.0.0 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.2 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/compress v1.15.6 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/olivere/elastic/v7 v7.0.32 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/stretchr/testify v1.7.5 // indirect
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292 // indirect
	golang.org/x/net v0.0.0-20220520000938-2e3eb7b945c2 // indirect
	k8s.io/apimachinery v0.24.2 // indirect
	sigs.k8s.io/json v0.0.0-20211208200746-9f7c6b3444d2 // indirect
)

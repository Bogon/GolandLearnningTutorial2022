module learning.goland.com/studyGoland/log_transfer/kafka

go 1.18

require (
	github.com/Shopify/sarama v1.34.1
	k8s.io/apimachinery v0.24.2
	learning.goland.com/studyGoland/log_transfer/es v0.0.0-00010101000000-000000000000
)

require (
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/olivere/elastic/v7 v7.0.32 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	sigs.k8s.io/json v0.0.0-20211208200746-9f7c6b3444d2 // indirect
)

replace learning.goland.com/studyGoland/log_transfer/es => ./../es

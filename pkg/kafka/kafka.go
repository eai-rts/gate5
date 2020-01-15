package kafka

import (
	ckafka "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"net/http"
)

// RequestSaver ..
type RequestSaver interface {
	Save(req *http.Request) error
}

type producer struct {
	CKafka: *ckafka.Producer
}

func (p *Producer) Save(req *http.Request) error {
	return nil
}

func (p *Producer) NewRequestSaver() *RequestSaver{
	p, err := kafka.NewProducer(&ckafka.ConfigMap{"bootstrap.servers": broker})
	if (err != nil) {
		log.Fatal(err)
	}
	return &producer{CKafka: &p}
}


package kafka

import (
	"log"
	"net/http"

	ckafka "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// Publisher ..
type Publisher interface {
	Publish(req *http.Request) error
}

type producer struct {
	*ckafka.Producer
}

func newProducer() producer {
	p, err := ckafka.NewProducer(&ckafka.ConfigMap{})
	if err != nil {
		log.Fatal(err)
	}
	return producer{p}
}

func (p *producer) Publish(req *http.Request) error {
	return nil
}

// NewPublisher is a Publisher constructor
func NewPublisher() Publisher {
	p := newProducer()
	return &p
}

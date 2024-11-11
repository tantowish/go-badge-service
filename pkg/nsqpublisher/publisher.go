package nsqpublisher

import (
	"log"

	"github.com/nsqio/go-nsq"
)

// Publisher struct holds the NSQ Producer and Topic
type Publisher struct {
	Producer *nsq.Producer
	Topic    string
}

// NewPublisher initializes a new NSQ producer for a given topic
func NewPublisher(nsqAddr, topic string) (*Publisher, error) {
	// Create NSQ config
	config := nsq.NewConfig()

	// Initialize a new producer
	producer, err := nsq.NewProducer(nsqAddr, config)
	if err != nil {
		return nil, err
	}

	return &Publisher{
		Producer: producer,
		Topic:    topic,
	}, nil
}

// PublishMessage publishes a message to the NSQ topic
func (p *Publisher) PublishMessage(message string) error {
	err := p.Producer.Publish(p.Topic, []byte(message))
	if err != nil {
		log.Printf("Failed to publish message to topic %s: %v", p.Topic, err)
		return err
	}

	log.Printf("Published message to topic %s: %s", p.Topic, message)
	return nil
}

// StopPublisher stops the NSQ producer
func (p *Publisher) StopPublisher() {
	p.Producer.Stop()
}
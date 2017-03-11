package redis

import (
	"github.com/rai-project/pubsub"
	redis "gopkg.in/redis.v5"
)

type subscriber struct {
	channel string
	sub     *redis.PubSub
	conn    *connection
	msgs    chan pubsub.Message
	err     error
	done    chan bool
}

func NewSubscriber(conn *connection, channel string) (pubsub.Subscriber, error) {
	sub, err := conn.Subscribe(channel)
	if err != nil {
		return nil, err
	}
	return &subscriber{
		conn:    conn,
		sub:     sub,
		channel: channel,
	}, nil
}

func (s *subscriber) Start() <-chan pubsub.Message {
	s.msgs = make(chan pubsub.Message)
	go func() {
		for {
			select {
			case <-s.done:
				return
			default:
				msg, err := s.sub.ReceiveMessage()
				if err != nil {
					s.err = err
					continue
				}
				if msg.Payload == pubsub.EndPayload {
					close(s.msgs)
					s.msgs = nil
					return
				}
				s.msgs <- createMessage(msg)
			}
		}
	}()
	return s.msgs
}

func (s *subscriber) Err() error {
	return s.err
}

func (s *subscriber) Stop() error {
	go func() {
		s.done <- true
	}()
	if s.msgs != nil {
		close(s.msgs)
	}
	if s.sub != nil {
		err := s.sub.Unsubscribe(s.channel)
		if err != nil {
			log.WithError(err).
				WithField("channel", s.channel).
				Error("unable to unsubscribe from channel")
		}
		err = s.sub.Close()
		if err != nil {
			log.WithError(err).Error("unable to close subscription")
		}
	}
	return nil
}

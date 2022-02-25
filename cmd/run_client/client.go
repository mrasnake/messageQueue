package main

import (
	"bufio"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/urfave/cli/v2"
	"os"
)

type service struct {
	RequestFile string
	QueueConn string
	QueueName string
}

func DefineService(ctx *cli.Context) *service {
	out := &service{
		RequestFile: ctx.String("file"),
		QueueConn: ctx.String("connection"),
		QueueName: ctx.String("queue"),
	}
	return out
}


func (s *service) Run() error{

	conn, err := amqp.Dial(s.QueueConn)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		s.QueueName, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}
	
	lines, err := s.getLines()
	if err != nil {
		return fmt.Errorf("unable to get lines: %w", err)
	}
	
	for _, l := range lines{
		if err := sendMessage(l, q, ch); err != nil {
			return fmt.Errorf("unable to process line: %w", err)
		}
	}
	return nil
}


func (s *service) getLines() ([]string, error){
	file, err := os.Open(s.RequestFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	out := []string{}
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func sendMessage(msg string, queue amqp.Queue, channal *amqp.Channel) error {

	err := channal.Publish(
		"",     // exchange
		queue.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType:"text/plain",
			Body:        []byte(msg),
		})
	if err != nil {
		return fmt.Errorf("problem publishing to queue: %w", err)
	}

	return nil
}
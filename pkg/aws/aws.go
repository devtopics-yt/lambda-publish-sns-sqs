package aws

import (
	"context"
	"os"
	"pub-lambda/pkg/logger"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"go.uber.org/zap"
)

type Connection struct {
	sqs *sqs.Client
	sns *sns.Client

	queueURL string
	topicARN string
}

func New() (*Connection, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}

	return &Connection{
		sqs:      sqs.NewFromConfig(cfg),
		sns:      sns.NewFromConfig(cfg),
		queueURL: os.Getenv("QUEUE_URL"),
		topicARN: os.Getenv("TOPIC_ARN"),
	}, nil
}

func (c *Connection) SendSQSMessage(ctx context.Context, message string) error {
	log := logger.GetLoggerFromContext(ctx)

	output, err := c.sqs.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: &message,
		QueueUrl:    &c.queueURL,
	})

	log.Info("sent sqs message", zap.Any("output", output))

	return err
}

func (c *Connection) PublishSNSMessage(ctx context.Context, message string) error {
	log := logger.GetLoggerFromContext(ctx)

	output, err := c.sns.Publish(ctx, &sns.PublishInput{
		Message:  &message,
		TopicArn: &c.topicARN,
	})

	log.Info("published sns message", zap.Any("output", output))

	return err
}

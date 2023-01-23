package main

import (
	"context"
	"pub-lambda/pkg/aws"
	"pub-lambda/pkg/logger"

	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

// var db, caches, loggers, or other AWS services (clients to interface those)
var ctx context.Context

func init() {
	log, _ := zap.NewProduction()
	c, err := aws.New()
	if err != nil {
		log.Fatal("couldn't set up aws clients", zap.Error(err))
	}

	// inject those clients into a context, and start our lambda with that context
	ctx = context.Background()

	// inject logger and aws client into context
	ctx = logger.Inject(ctx, log)
	ctx = aws.Inject(ctx, c)
}

type Test struct {
	Hello string `json:"hello"`
}

func Handler(ctx context.Context, event interface{}) {
	log := logger.GetLoggerFromContext(ctx)
	log.Info("received event lambda", zap.Any("event", event))

	aws := aws.GetConnectionFromContext(ctx)

	err := aws.SendSQSMessage(ctx, "hello sqs")
	if err != nil {
		log.Error("couldn't send sqs", zap.Error(err))
	}
	err = aws.PublishSNSMessage(ctx, "hello sns")
	if err != nil {
		log.Error("couldn't send sns", zap.Error(err))
	}
}

func main() {
	lambda.StartWithContext(ctx, Handler)
}

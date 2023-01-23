package aws

import (
	"context"
)

type AWSKeyType string

var awsKey AWSKeyType = "AWS"

func Inject(ctx context.Context, c *Connection) context.Context {
	return context.WithValue(ctx, awsKey, c)
}

func GetConnectionFromContext(ctx context.Context) *Connection {
	c, _ := ctx.Value(awsKey).(*Connection)
	return c
}

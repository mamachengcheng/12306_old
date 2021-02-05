package client

import (
	"context"
	"google.golang.org/grpc"
	"time"
)

type Executor interface {
	exec(coon *grpc.ClientConn, ctx context.Context, request map[string]interface{}) (map[string]interface{}, error)
}

type GRPCClient struct {
	Address  string
	Executor Executor
}

func (c *GRPCClient) Send(request map[string]interface{}) (map[string]interface{}, error) {

	var err error
	if coon, err := grpc.Dial(c.Address, grpc.WithInsecure(), grpc.WithBlock()); err == nil {
		defer coon.Close()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		reply, err := c.Executor.exec(coon, ctx, request)
		return reply, err
	}
	return nil, err
}

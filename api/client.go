package api

import (
	"context"
	"errors"
	"time"

	"github.com/oriiolabs/pebble/pb"
	"google.golang.org/grpc"
)

type Client struct {
	rpc  pb.CacheClient
	conn *grpc.ClientConn
}

func NewClient(serverAddress string, opts []grpc.DialOption) (*Client, error) {
	if len(serverAddress) == 0 {
		return nil, errors.New("you must provide a server address to connect to")
	}

	c := &Client{}

	if len(opts) == 0 {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(opts, grpc.FailOnNonTempDialError(true))
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(serverAddress, opts...)
	if err != nil {
		return nil, err
	}

	c.conn = conn
	c.rpc = pb.NewCacheClient(c.conn)

	return c, nil
}

func (c *Client) Get(ctx context.Context, key string) ([]byte, error) {
	req := &pb.GetRequest{
		Key: key,
	}

	resp, err := c.rpc.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Value, nil
}

func (c *Client) Set(ctx context.Context, key string, value []byte) error {
	req := &pb.SetRequest{
		Key:   key,
		Value: value,
	}

	if _, err := c.rpc.Set(ctx, req); err != nil {
		return err
	}

	return nil
}

func (c *Client) SetTTL(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	req := &pb.SetTTLRequest{
		Key:   key,
		Value: value,
		Ttl:   int64(ttl.Seconds()),
	}

	if _, err := c.rpc.SetTTL(ctx, req); err != nil {
		return err
	}

	return nil
}

func (c *Client) Delete(ctx context.Context, key string) error {
	req := &pb.DeleteRequest{
		Key: key,
	}

	if _, err := c.rpc.Delete(ctx, req); err != nil {
		return err
	}

	return nil
}

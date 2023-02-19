package helloservice

import (
	"context"
	"io"
	"time"
)

type HelloService struct {
}

func (h HelloService) mustEmbedUnimplementedHelloServiceServer() {
	panic("implement me")
}

func (h HelloService) Hello(ctx context.Context, args *String) (*String, error) {
	time.Sleep(time.Second)
	reply := &String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func (h HelloService) Channel(stream HelloService_ChannelServer) error {
	for {
		recv, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &String{Value: "hello:" + recv.Value}
		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

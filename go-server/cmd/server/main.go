package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jinpeng/flutter-go-grpc/go-server/pkg/protocol/grpc"
	"github.com/jinpeng/flutter-go-grpc/go-server/pkg/service/v1"
)

func main() {
	if err := grpc.RunServer(context.Background(), v1.NewChatServiceServer(), "3000"); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
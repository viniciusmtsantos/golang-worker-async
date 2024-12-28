package grpc

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	"grpc-caller/grpc/server/pb"
	"grpc-caller/person"
)

func StartGrpcServer(ctx context.Context, port *string) {
	opts := []grpc.ServerOption{}

	opts = append(opts,
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             5 * time.Minute,
				PermitWithoutStream: true,
			},
		),
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				Time:    5 * time.Second,
				Timeout: 2 * time.Second,
			},
		),
	)

	grpcServer := grpc.NewServer(opts...)

	// Person
	pb.RegisterPersonRPCServer(grpcServer, &person.PersonGrpcHandlers{})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")
			grpcServer.GracefulStop()
			<-ctx.Done()
		}
	}()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		log.Fatalf("could not listen to %d: %v", port, err)
	}

	log.WithFields(log.Fields{"gRPC Port": *port}).Info("Server is running")

	reflection.Register(grpcServer)

	log.Fatal(grpcServer.Serve(listener))
}

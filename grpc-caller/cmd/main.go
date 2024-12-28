package main

import (
	"context"
	"flag"
	"grpc-caller/grpc"
	"log"
	_ "net/http/pprof"
)

type Config struct {
	GRPCPort      string
	LogLevel      int
	LogTimeFormat string
}

func main() {

	var cfg Config

	cfg.LogLevel = -1
	ctx := context.Background()

	flag.StringVar(&cfg.GRPCPort, "port", "50053", "gRPC port to bind")
	flag.IntVar(&cfg.LogLevel, "log-level", -1, "Global log level")
	flag.StringVar(&cfg.LogTimeFormat, "log-time-format", "02/01/2006 15:04:05Z07:00", "Print time format for logger e.g. 02/01/2006 15:04:05Z07:00")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		log.Fatalf("invalid tcp port for gRPC server: '%s'", cfg.GRPCPort)
		return
	}

	grpc.StartGrpcServer(ctx, &cfg.GRPCPort)

}

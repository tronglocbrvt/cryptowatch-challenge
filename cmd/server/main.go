package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"log"

	"github.com/cryptowatch_challenge/config"
	pb "github.com/cryptowatch_challenge/pb/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	// load configs
	cfg := config.Load()

	// grpc server
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
			otelgrpc.UnaryServerInterceptor(),
		)),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)

	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(s)
	// Register Prometheus metrics handler.
	l, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("failed when new zap logger")
	}
	svc := registerService(cfg, l)
	pb.RegisterCryptoWatchServer(s, svc)

	// handle signal
	_, ctxCancel := context.WithCancel(context.Background())
	go func() {
		osSignal := make(chan os.Signal, 1)
		signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)
		<-osSignal
		ctxCancel()
		// Wait for maximum 15s
		go func() {
			var durationSec time.Duration = 1
			timer := time.NewTimer(durationSec * time.Second)
			<-timer.C
			l.Fatal("Force shutdown due to timeout!")
		}()
	}()

	go func() {
		gw := NewServer(cfg)
		l.Info("HTTP server start listening", zap.Int("HTTP address", cfg.HTTPAddress))
		err := gw.RunGRPCGateway()
		if err != nil {
			l.Fatal("error listening to HTTP address", zap.Error(err))
			return
		}
	}()

	l.Info("GRPC server start listening", zap.Int("GRPC address", cfg.GRPCAddress))
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCAddress))
	if err != nil {
		l.Fatal("error listening to address", zap.Int("address", cfg.GRPCAddress), zap.Error(err))
		return
	}

	err = s.Serve(listener)
	if err != nil {
		l.Fatal("error when serve listener", zap.Error(err))
	}
}

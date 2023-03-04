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
	"github.com/cryptowatch_challenge/external"
	"github.com/cryptowatch_challenge/internal/stores"
	pb "github.com/cryptowatch_challenge/pb/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// load configs
	cfg := config.Load()

	// grpc server
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			otelgrpc.UnaryServerInterceptor(),
		)),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)

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
			timer := time.NewTimer(time.Second)
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

	db := mustConnectPostgreSQL(cfg)
	priceStore := stores.NewPriceStore(db)
	cryptoWatchClient := external.NewCryptoWatchClient(cfg)

	go cryptoWatchClient.ListenCryptoWatch(priceStore)

	err = s.Serve(listener)
	if err != nil {
		l.Fatal("error when serve listener", zap.Error(err))
	}
}

func mustConnectPostgreSQL(cfg *config.Config) *gorm.DB {
	logLevel := logger.Silent

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logLevel,    // Log level
			Colorful:      true,        // Disable color
		},
	)

	dsn := "host=localhost user=loc.truong dbname=crypto_watch port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	err = db.Raw("SELECT 1").Error
	if err != nil {
		log.Fatal("error querying SELECT 1", zap.Error(err))
	}

	return db
}

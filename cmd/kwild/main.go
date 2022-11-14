package main

import (
	"context"
	"fmt"
	"kwil/x/async"
	"kwil/x/execution"
	"kwil/x/schema"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"kwil/x/cfgx"
	"kwil/x/crypto"
	"kwil/x/deposits"
	"kwil/x/grpcx"
	"kwil/x/logx"
	"kwil/x/proto/apipb"
	"kwil/x/service/apisvc"

	kg "kwil/cmd/kwild-gateway/server"

	"github.com/oklog/run"
)

func execute(logger logx.Logger) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dc := cfgx.GetConfig().Select("deposit-settings")

	kr, err := crypto.NewKeyring(dc)
	if err != nil {
		return fmt.Errorf("failed to create keyring: %w", err)
	}

	acc, err := kr.GetDefaultAccount()
	if err != nil {
		return fmt.Errorf("failed to get default account: %w", err)
	}

	d, err := deposits.New(dc, logger, acc)
	if err != nil {
		return fmt.Errorf("failed to initialize new deposits: %w", err)
	}
	err = d.Listen(ctx)
	if err != nil {
		return fmt.Errorf("failed to listen to deposits: %w", err)
	}

	apiService := apisvc.NewService(d, schema.NewTestService(), execution.NewTestService())
	httpHandler := apisvc.NewHandler(logger)

	return serve(logger, httpHandler, apiService)
}

func serve(logger logx.Logger, httpHandler http.Handler, apiService apipb.KwilServiceServer) error {
	var g run.Group

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	g.Add(func() error {
		grpcServer := grpcx.NewServer(logger)
		apipb.RegisterKwilServiceServer(grpcServer, apiService)
		return grpcServer.Serve(listener)
	}, func(error) {
		listener.Close()
	})

	httpServer := http.Server{
		Addr:    ":8081",
		Handler: httpHandler,
	}
	g.Add(func() error {
		return httpServer.ListenAndServe()
	}, func(error) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = httpServer.Shutdown(ctx)
	})

	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})

	return g.Run()
}

func main() {
	logger := logx.New()

	stop := func(err error) {
		logger.Sugar().Error(err)
		os.Exit(1)
	}

	kwild := func() error {
		return execute(logger)
	}

	if !isGatewayEnabled() {
		if err := kwild(); err != nil {
			stop(err)
		}
	}

	async.Run(kg.Start).Catch(stop)

	<-async.Run(kwild).Catch(stop).DoneCh()
}

func isGatewayEnabled() bool {
	var args []string
	with_gateway_flag := false
	found := -2
	for i, arg := range os.Args {
		if i == found+1 {
			if arg == "true" {
				with_gateway_flag = true
			}
			continue
		}

		if arg != "--withgateway" {
			args = append(args, arg)
			continue
		}

		found = i
	}

	if with_gateway_flag {
		os.Args = args //make sure the flag and value are removed
	}

	return with_gateway_flag
}

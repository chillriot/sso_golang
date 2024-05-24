package app

import (
	"context"
	"flag"
	descAuthV1 "gitea.24example.ru/RosarStoreBackend/protobuf/pkg/sso_v1"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/closer"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/config"
	"gitea.24example.ru/RosarStoreBackend/sso_v1/internal/interceptor"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"sync"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	httpServer      *http.Server
}

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		if err := a.runGRPCServer(); err != nil {
			log.Error().Err(err).Msg("grpc server failed")
			os.Exit(1)
		}
	}()

	go func() {
		defer wg.Done()

		if err := a.runHTTPServer(); err != nil {
			log.Error().Err(err).Msg("http server failed")
			os.Exit(1)
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
		a.initGRPCServer,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	if err := config.Load(configPath); err != nil {
		return err
	}
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := descAuthV1.RegisterAuthServiceHandlerFromEndpoint(
		ctx,
		mux,
		a.serviceProvider.GRPCConfig().Address(),
		opts,
	)
	if err != nil {
		return err
	}

	err = descAuthV1.RegisterTokenServiceHandlerFromEndpoint(
		ctx,
		mux,
		a.serviceProvider.GRPCConfig().Address(),
		opts,
	)
	if err != nil {
		return err
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Content-Length", "X-Requested-With"},
		AllowCredentials: true,
	})

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.HTTPConfig().Address(),
		Handler: corsMiddleware.Handler(mux),
	}

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(interceptor.ValidateInterceptor),
	)

	reflection.Register(a.grpcServer)

	descAuthV1.RegisterAuthServiceServer(a.grpcServer, a.serviceProvider.AuthV1Impl(ctx))
	descAuthV1.RegisterTokenServiceServer(a.grpcServer, a.serviceProvider.AuthV1Impl(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	log.Printf("gRPC Server in running on: %s", a.serviceProvider.GRPCConfig().Address())

	lis, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	if err = a.grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (a *App) runHTTPServer() error {
	log.Printf("HTTP Server in running on: %s", a.serviceProvider.HTTPConfig().Address())

	if err := a.httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

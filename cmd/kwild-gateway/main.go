package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	v0 "github.com/kwilteam/kwil-db/internal/api/v0"
)

const (
	grpcEndpointEnv = "KWIL_GRPC_ENDPOINT"
)

func run() error {
	cmd := &cobra.Command{
		Use:   "api-gateway",
		Short: "api-gateway is a gRPC to HTTP gateway",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			mux := runtime.NewServeMux()
			opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
			err := v0.RegisterKwilServiceHandlerFromEndpoint(cmd.Context(), mux, viper.GetString("endpoint"), opts)
			if err != nil {
				return err
			}

			mux.HandlePath(http.MethodGet, "/api/v0/swagger.json", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
				v0.ServeSwaggerJSON(w, r)
			})

			mux.HandlePath(http.MethodGet, "/swagger/ui", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
				v0.ServeSwaggerUI(w, r)
			})

			return http.ListenAndServe(":8080", cors(mux))
		},
	}

	cmd.PersistentFlags().String("endpoint", "localhost:50051", "gRPC server endpoint")
	viper.BindPFlag("endpoint", cmd.PersistentFlags().Lookup("endpoint"))
	viper.BindEnv("endpoint", grpcEndpointEnv)

	return cmd.Execute()
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func allowedOrigin(origin string) bool {
	if viper.GetString("cors") == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}
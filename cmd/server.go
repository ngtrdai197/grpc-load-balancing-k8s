package cmd

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/ngtrdai197/grpc-lb/config"
	grpc_server "github.com/ngtrdai197/grpc-lb/pkg"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Serve gRPC server application",
	Long:  `Server gRPC`,
	Run: func(_ *cobra.Command, _ []string) {
		c, err := config.GetConfig(validator.New())
		if err != nil {
			panic(fmt.Errorf("config file invalidate with error: %w", err))
		}
		initGrpcServer(c)
	},
}

func initGrpcServer(c *config.Config) {
	s := grpc_server.NewServer(c)
	s.Start()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

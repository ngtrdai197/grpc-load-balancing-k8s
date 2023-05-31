package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/ngtrdai197/grpc-lb/config"
	"github.com/ngtrdai197/grpc-lb/pkg/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Serve gRPC client application",
	Long:  `Client gRPC`,
	Run: func(_ *cobra.Command, _ []string) {
		c, err := config.GetConfig(validator.New())
		if err != nil {
			panic(fmt.Errorf("config file invalidate with error: %w", err))
		}
		initGrpcClient(c)
	},
}

func initGrpcClient(c *config.Config) {
	conn, err := grpc.Dial(c.GrpcServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed connect to grpc client, error=%v", err)
	}

	client := pb.NewMessageServiceClient(conn)

	for i := 1; i <= 10; i++ {
		msg, err := client.SendMessage(context.Background(), &pb.Message{
			Msg: fmt.Sprintf("Msg sent from client %d", i),
		})
		if err != nil {
			log.Printf("sent msg to server occurs error=%v", err)
		}
		log.Printf("msg sent: %v", msg)
	}

}

func init() {
	rootCmd.AddCommand(clientCmd)
}

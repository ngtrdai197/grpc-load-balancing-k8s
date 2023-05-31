package cmd

import (
	"context"
	"fmt"
	"log"
	"strconv"

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
	Run: func(cmd *cobra.Command, _ []string) {
		c, err := config.GetConfig(validator.New())
		if err != nil {
			panic(fmt.Errorf("config file invalidate with error: %w", err))
		}
		numbOfMsg := cmd.Flags().Lookup("numOfMsg").Value.String()
		num, err := strconv.Atoi(numbOfMsg)
		if err != nil {
			log.Panicf("numOfMsg is not a number, error=%v", err)
		}
		initGrpcClient(c, num)
	},
}

func initGrpcClient(c *config.Config, num int) {
	conn, err := grpc.Dial(c.GrpcClientAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed connect to grpc client, error=%v", err)
	}

	client := pb.NewMessageServiceClient(conn)

	for i := 1; i <= num; i++ {
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

	clientCmd.Flags().String("numOfMsg", "", "Number of msg you want to send")

	err := clientCmd.MarkFlagRequired("numOfMsg")
	if err != nil {
		return
	}
}

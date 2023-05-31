package config

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	GrpcServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"  validate:"required"`
	GrpcClientAddress string `mapstructure:"GRPC_CLIENT_ADDRESS"  validate:"required"`
}

func GetConfig(validator *validator.Validate) (*Config, error) {
	c := &Config{
		GrpcServerAddress: viper.GetString("GRPC_SERVER_ADDRESS"),
		GrpcClientAddress: viper.GetString("GRPC_CLIENT_ADDRESS"),
	}
	if err := validator.StructCtx(context.Background(), c); err != nil {
		return nil, err
	}
	return c, nil
}

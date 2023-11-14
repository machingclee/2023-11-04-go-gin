package util

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type Env struct {
	DBDriver            string        `json:"DB_DRIVER"`
	DBSource            string        `json:"DB_SOURCE"`
	ServerAddress       string        `json:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `json:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `json:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(parentDir string) (*Env, error) {
	env := "local"
	env_override := os.Getenv("env")
	if env_override != "" {
		env = env_override
	}

	secretName := fmt.Sprintf("simple_bank_%s", env)
	region := "ap-northeast-1"

	fmt.Printf("Getting Environment variables [%s] from aws screts ... ", secretName)

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}
	svc := secretsmanager.NewFromConfig(config)
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}
	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		log.Fatal(err.Error())
	}
	var secretString string = *result.SecretString
	var envVariables = &Env{}
	json.Unmarshal([]byte(secretString), envVariables)

	fmt.Println("Env variables retrieved")
	return envVariables, nil
}

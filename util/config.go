package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type Env struct {
	DBDriver             string        `json:"DB_DRIVER"`
	DBSource             string        `json:"DB_SOURCE"`
	ServerAddress        string        `json:"SERVER_ADDRESS"`
	TokenSymmetricKey    string        `json:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `json:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `json:"REFRESH_TOKEN_DURATION"`
}

type intermediateEnv struct {
	AccessTokenDuration  string `json:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration string `json:"REFRESH_TOKEN_DURATION"`
}

var ENV *Env = nil

func LoadEnv() (*Env, error) {
	if ENV != nil {
		return ENV, nil
	}
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
	ENV = &Env{}
	envIntermediate := &intermediateEnv{}
	json.Unmarshal([]byte(secretString), ENV)
	json.Unmarshal([]byte(secretString), envIntermediate)

	ENV.AccessTokenDuration, err = time.ParseDuration(envIntermediate.AccessTokenDuration)
	if err != nil {
		log.Fatal(err)
	}
	ENV.RefreshTokenDuration, err = time.ParseDuration(envIntermediate.RefreshTokenDuration)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Env variables retrieved")
	return ENV, nil
}

func RunDbMigration() error {
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}
	env, err := LoadEnv()
	if err != nil {
		return err
	}
	fmt.Println("Running Migration Script")

	os.Setenv("DB_URL", env.DBSource)
	migrationScriptPath := filepath.Join(workingDir, "script_db_migrate_up.sh")

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		var winBash string = `C:\Program Files\Git\usr\bin\sh.exe`
		if where := os.Getenv("bin_where"); where != "" {
			winBash = where
		}
		cmd = exec.Command(winBash, migrationScriptPath)
	} else {
		cmd = exec.Command(migrationScriptPath)
	}

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err = cmd.Run()

	if err != nil {
		return err
	}

	fmt.Println("out:", outb.String())
	fmt.Println("err:", errb.String())

	return nil
}

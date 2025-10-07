package config

import (
	"errors"
	_const "mhride_backend/utils/const"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Env struct {
	DeploymentEnv        string
	MongodbUri           string
	Database             string
	RedisHost            string
	RedisPort            string
	RedisPassword        string
	RedisUserName        string
	GrpcAddr             string
	GrpcWebAddr          string
	AwsRegion            string
	AwsAccessKey         string
	AwsSecretAccessKey   string
	JwtSecretKey         string
	S3BucketName         string
	S3BucketUrl          string
	EmailEnabled         bool
	SenderEmail          string
	ZeptomailToken       string
	StripePublishableKey string
	StripeSecretKey      string
	StripeWebhookSecret  string
	ConnectyCubeAuthKey  string
	AwsSqsQueueUrl       string
	GoogleApiKey         string
	SmsEnabled           bool
	TwilioAccountSid     string
	TwilioAuthToken      string
	TwilioServiceId      string
}

func NewEnv() (*Env, error) {
	deploymentEnv := strings.TrimSpace(os.Getenv("DEPLOYMENT_ENV"))
	if len(deploymentEnv) == 0 {
		// default to dev if not set
		deploymentEnv = string(_const.Deployment_Dev)
	}

	// Decide which env file to load
	var envFile string
	switch deploymentEnv {
	case string(_const.Deployment_Production):
		envFile = ".env.prod"
	default:
		envFile = ".env.dev"
	}

	// Load the environment file
	err := godotenv.Load(envFile)
	if err != nil {
		return nil, errors.New("failed to load env file: " + envFile)
	}

	// Helper function to fetch env vars
	getEnv := func(key string) (string, error) {
		val := os.Getenv(key)
		if len(val) == 0 {
			return "", errors.New(key + " is empty")
		}
		return val, nil
	}

	emailEnabled, _ := getEnv("EMAIL_ENABLED")
	smsEnabled, _ := getEnv("SMS_ENABLED")

	return &Env{
		DeploymentEnv:        deploymentEnv,
		MongodbUri:           must(getEnv("MONGODB_URI")),
		Database:             must(getEnv("DATABASE")),
		RedisHost:            must(getEnv("REDIS_HOST")),
		RedisPort:            must(getEnv("REDIS_PORT")),
		RedisPassword:        must(getEnv("REDIS_PASSWORD")),
		RedisUserName:        must(getEnv("REDIS_USERNAME")),
		GrpcAddr:             must(getEnv("GRPC_ADDR")),
		GrpcWebAddr:          must(getEnv("GRPC_WEB_ADDR")),
		AwsRegion:            must(getEnv("AWS_REGION")),
		AwsAccessKey:         must(getEnv("AWS_ACCESS_KEY")),
		AwsSecretAccessKey:   must(getEnv("AWS_SECRET_ACCESS_KEY")),
		JwtSecretKey:         must(getEnv("JWT_SECRET_KEY")),
		S3BucketName:         must(getEnv("S3_BUCKET_NAME")),
		S3BucketUrl:          must(getEnv("S3_BUCKET_URL")),
		EmailEnabled:         emailEnabled == "true",
		SenderEmail:          must(getEnv("SENDER_EMAIL")),
		ZeptomailToken:       must(getEnv("ZEPTOMAIL_TOKEN")),
		GoogleApiKey:         must(getEnv("GOOGLE_API_KEY")),
		StripePublishableKey: must(getEnv("STRIPE_PUBLISHABLE_KEY")),
		StripeSecretKey:      must(getEnv("STRIPE_SECRET_KEY")),
		StripeWebhookSecret:  must(getEnv("STRIPE_WEBHOOK_SECRET")),
		ConnectyCubeAuthKey:  must(getEnv("CONNECTYCUBE_API_KEY")),
		AwsSqsQueueUrl:       must(getEnv("AWS_SQS_QUEUE_URL")),
		SmsEnabled:           smsEnabled == "true",
		TwilioAccountSid:     must(getEnv("TWILIO_ACCOUNT_SID")),
		TwilioAuthToken:      must(getEnv("TWILIO_AUTH_TOKEN")),
		TwilioServiceId:      must(getEnv("TWILIO_SERVICE_ID")),
	}, nil
}

// must panics if err is not nil (for simplicity)
func must(value string, err error) string {
	if err != nil {
		panic(err)
	}
	return value
}

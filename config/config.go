package config

import (
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadConfig() IConfig {
	envMap, err := godotenv.Read(".env")
	if err != nil {
		logrus.Println("error load environment:", err.Error())
	}

	return &config{
		minio: &minio{
			endPoint:     envMap["MINIO_ENDPOINT"],
			accessKey:    envMap["MINIO_ACCESS_KEY"],
			accessSecret: envMap["MINIO_ACCESS_SECRET"],
			secure: func() bool {
				secure, err := strconv.ParseBool(envMap["MINIO_SECURE"])
				if err != nil {
					logrus.Println("error parse secure:", err.Error())
				}
				return secure
			}(),
			bucket: envMap["MINIO_BUCKET"],
		},
	}
}

type IConfig interface {
	MiniO() IMiniO
}
type config struct {
	minio *minio
}

func (c *config) MiniO() IMiniO {
	return c.minio
}

type IMiniO interface {
	GetEndPoint() string
	GetAccessKey() string
	GetAccessSecret() string
	GetSecure() bool
	GetBucket() string
}

type minio struct {
	endPoint     string
	accessKey    string
	accessSecret string
	secure       bool
	bucket       string
}

func (m *minio) GetEndPoint() string {
	return m.endPoint
}

func (m *minio) GetAccessKey() string {
	return m.accessKey
}

func (m *minio) GetAccessSecret() string {
	return m.accessSecret
}

func (m *minio) GetSecure() bool {
	return m.secure
}

func (m *minio) GetBucket() string {
	return m.bucket
}

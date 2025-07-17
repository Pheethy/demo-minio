package main

import (
	"bytes"
	"context"
	"demo-minio/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.LoadConfig()
	minioClient, err := minio.New(cfg.MiniO().GetEndPoint(), &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MiniO().GetAccessKey(), cfg.MiniO().GetAccessSecret(), ""),
		Secure: cfg.MiniO().GetSecure(),
	})
	if err != nil {
		logrus.Println("connent minio failed:", err.Error())
	}

	ctx := context.Background()
	fileContent := []byte("Hello MiniO")
	fileName := "photo/test.txt"
	uploadInfo, err := minioClient.PutObject(ctx, cfg.MiniO().GetBucket(), fileName, bytes.NewReader(fileContent), int64(len(fileContent)), minio.PutObjectOptions{})
	if err != nil {
		logrus.Println("upload file failed:", err.Error())
	}

	logrus.Info("upload file success:", uploadInfo)
}

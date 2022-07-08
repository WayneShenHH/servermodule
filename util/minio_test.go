package util_test

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/WayneShenHH/servermodule/logger"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func Test_Minio(t *testing.T) {
	endpoint := "minio-dev.paradise-soft.com.tw/"
	accessKeyID := "CSzzvaDIfIgFaffo"
	secretAccessKey := "FLS47v8XeFuyu47rk0iguj8Ej3xuppmo"
	useSSL := true
	bucket := "dev-squirrel"

	ctx := context.Background()

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		logger.Error(err)
	}

	ans, err := minioClient.BucketExists(ctx, bucket)
	if err != nil {
		logger.Error(err)
	} else {
		logger.Info(ans)
	}

	objectCh := minioClient.ListObjects(ctx, bucket, minio.ListObjectsOptions{
		Prefix:    "kkgame",
		Recursive: true,
	})
	for object := range objectCh {
		if object.Err != nil {
			logger.Error(object.Err)
			return
		}
		logger.Info(object)
	}

	object, err := minioClient.GetObject(ctx, bucket, "kkgame/kkgame_01.png", minio.GetObjectOptions{})
	if err != nil {
		logger.Error(err)
		return
	}
	localFile, err := os.Create("/tmp/local-file.jpg")
	if err != nil {
		logger.Error(err)
		return
	}
	if _, err = io.Copy(localFile, object); err != nil {
		logger.Error(err)
		return
	}

	file, err := os.Open("/home/shen/Downloads/kkgame_01.png")
	if err != nil {
		logger.Error(err)
		return
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		logger.Error(err)
		return
	}

	uploadInfo, err := minioClient.PutObject(context.Background(), bucket, "kkgame/kkgame_01.png", file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Successfully uploaded bytes: ", uploadInfo)
	logger.Info(fmt.Sprintf("http://10.200.252.101:9000/%v/%v", uploadInfo.Bucket, uploadInfo.Key))
}

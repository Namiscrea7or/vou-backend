// pkg/storage/firebase.go
package storage

import (
	"context"
	"fmt"
	"io"
	"os"

	firebase "firebase.google.com/go"
)

func UploadFileToFirebaseStorage(bucketName string, filePath string, fileName string) (string, error) {
	ctx := context.Background()
	conf := &firebase.Config{StorageBucket: bucketName}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return "", err
	}

	client, err := app.Storage(ctx)
	if err != nil {
		return "", err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	wc := bucket.Object(fileName).NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	imageURL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media", bucketName, fileName)
	return imageURL, nil
}

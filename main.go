package main

import (
	"context"
	"fmt"
	"log"

	"github.com/BuildWithYou/fetroshop-api/app/helper/confighelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/BuildWithYou/fetroshop-api/app/helper/miniohelper"
	"github.com/BuildWithYou/fetroshop-api/app/injector"
	"github.com/minio/minio-go/v7"
)

func main() {

	// Initialize minio
	config := confighelper.GetConfig()
	cmsLogger := logger.NewCmsLogger(config)
	myMinio := miniohelper.GetMinio(config, cmsLogger)
	bucketIsExist, err := myMinio.Client.BucketExists(context.Background(), myMinio.BucketName)
	if err != nil {
		cmsLogger.Fatal(err.Error())
	}
	if !bucketIsExist {
		cmsLogger.Info("Bucket not exist, creating bucket...")
		err = myMinio.Client.MakeBucket(context.Background(), myMinio.BucketName, minio.MakeBucketOptions{})
		if err != nil {
			cmsLogger.Fatal(err)
		}
		cmsLogger.Info("Bucket created successfully")

		// Set public access policy for the bucket
		policy := `
      {
         "Version":"2012-10-17",
         "Statement":[
               {
                  "Effect":"Allow",
                  "Principal":{"AWS":["*"]},
                  "Action":["s3:GetObject"],
                  "Resource":
                     [
                        "arn:aws:s3:::` + myMinio.BucketName + `/public/*",
                        "arn:aws:s3:::` + myMinio.BucketName + `/testing-images/*"
                     ]
               }
            ]
      }`
		err = myMinio.Client.SetBucketPolicy(context.Background(), myMinio.BucketName, policy)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("Bucket %s created with public access.\n", myMinio.BucketName)
	}

	// Run cms server
	go func() {
		cmsApp := injector.InitializeCmsServer()
		cmsApp.Logger.Info("Starting cms server...")
		err := cmsApp.FiberApp.Listen(fmt.Sprintf("%s:%d", cmsApp.Host, cmsApp.Port))
		if err != nil {
			cmsApp.Logger.Panic(err.Error())
		}
	}()

	// Run web server
	webApp := injector.InitializeWebServer()
	webApp.Logger.Info("Starting web server...")
	err = webApp.FiberApp.Listen(fmt.Sprintf("%s:%d", webApp.Host, webApp.Port))
	if err != nil {
		webApp.Logger.Panic(err.Error())
	}
}

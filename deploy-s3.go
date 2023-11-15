package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "fmt"
)

func main() {
    // Create a new AWS session using your AWS credentials
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String("your-aws-region"),
    }))

    // Create an S3 service client
    s3Svc := s3.New(sess)

    // Specify the S3 bucket name
    bucketName := "my-unique-bucket-name"

    // Create the S3 bucket
    _, err := s3Svc.CreateBucket(&s3.CreateBucketInput{
        Bucket: aws.String(bucketName),
    })

    if err != nil {
        fmt.Println("Error creating S3 bucket:", err)
        return
    }

    fmt.Println("Created S3 bucket:", bucketName)
}

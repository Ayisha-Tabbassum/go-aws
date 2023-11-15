package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
    "fmt"
)

func main() {
    // Create a new AWS session using your AWS credentials
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String("your-aws-region"),
    }))

    // Create an EC2 service client
    ec2Svc := ec2.New(sess)

    // Specify the EC2 instance parameters
    runInput := &ec2.RunInstancesInput{
        ImageId:      aws.String("ami-12345678"), // Replace with your desired AMI ID
        InstanceType: aws.String("t2.micro"),    // Replace with your desired instance type
        MinCount:     aws.Int64(1),
        MaxCount:     aws.Int64(1),
    }

    // Launch the EC2 instance
    runResult, err := ec2Svc.RunInstances(runInput)
    if err != nil {
        fmt.Println("Error launching EC2 instance:", err)
        return
    }

    // Print the instance ID
    instanceID := runResult.Instances[0].InstanceId
    fmt.Println("Launched EC2 instance with ID:", *instanceID)
}

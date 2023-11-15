package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/eks"
    "fmt"
)

func main() {
    // Create a new AWS session using your AWS credentials
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String("your-aws-region"),
    }))

    // Create an EKS service client
    eksSvc := eks.New(sess)

    // Define the number of EKS clusters to create
    numClusters := 10

    for i := 1; i <= numClusters; i++ {
        // Specify cluster parameters
        clusterName := fmt.Sprintf("eks-cluster-%d", i) // Unique cluster name
        vpcID := "your-vpc-id"                          // Replace with your VPC ID
        subnetIDs := []*string{
            aws.String("subnet-1"),
            aws.String("subnet-2"),
            // Add more subnet IDs as needed
        }

        // Create the EKS cluster
        createClusterInput := &eks.CreateClusterInput{
            Name:    aws.String(clusterName),
            RoleArn: aws.String("arn:aws:iam::123456789012:role/eks-cluster-role"), // Replace with the IAM role ARN for EKS
            ResourcesVpcConfig: &eks.VpcConfigRequest{
                SubnetIds: []*string{vpcID},
            },
        }

        _, err := eksSvc.CreateCluster(createClusterInput)
        if err != nil {
            fmt.Printf("Error creating EKS cluster %s: %v\n", clusterName, err)
        } else {
            fmt.Printf("Created EKS cluster: %s\n", clusterName)
        }
    }
}

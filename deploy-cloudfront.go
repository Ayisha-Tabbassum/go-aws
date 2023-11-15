package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/cloudfront"
    "fmt"
)

func main() {
    // Create a new AWS session using your AWS credentials
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String("your-aws-region"),
    }))

    // Create a CloudFront service client
    cloudfrontSvc := cloudfront.New(sess)

    // Specify the CloudFront distribution configuration
    distributionConfig := &cloudfront.DistributionConfig{
        CallerReference: aws.String("unique-reference-id"), // Provide a unique reference ID
        DefaultCacheBehavior: &cloudfront.DefaultCacheBehavior{
            ForwardedValues: &cloudfront.ForwardedValues{
                QueryString: aws.Bool(true),
            },
            TargetOriginId: aws.String("your-origin-id"), // Replace with your Origin ID
            ViewerProtocolPolicy: aws.String("allow-all"),
        },
        Origins: &cloudfront.Origins{
            Items: []*cloudfront.Origin{
                {
                    Id:   aws.String("your-origin-id"), // Replace with your Origin ID
                    DomainName: aws.String("your-origin-domain"), // Replace with your Origin Domain Name
                    CustomOriginConfig: &cloudfront.CustomOriginConfig{
                        HTTPPort:              aws.Int64(80),
                        HTTPSPort:             aws.Int64(443),
                        OriginProtocolPolicy: aws.String("http-only"),
                    },
                },
            },
            Quantity: aws.Int64(1),
        },
    }

    // Create the CloudFront distribution
    createDistributionInput := &cloudfront.CreateDistributionInput{
        DistributionConfig: distributionConfig,
    }

    result, err := cloudfrontSvc.CreateDistribution(createDistributionInput)
    if err != nil {
        fmt.Println("Error creating CloudFront distribution:", err)
        return
    }

    fmt.Println("Created CloudFront distribution with ID:", *result.Distribution.Id)
}

package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/waf"
    "fmt"
)

func main() {
    // Create a new AWS session using your AWS credentials
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String("your-aws-region"),
    }))

    // Create a WAF service client
    wafSvc := waf.New(sess)

    // Specify the WebACL configuration
    webAclName := "MyWebAcl" // Replace with your desired WebACL name
    webAclRules := []*waf.Rule{
        {
            Action: &waf.Action{
                Type: aws.String("BLOCK"),
            },
            Priority: aws.Int64(1),
            RuleId:   aws.String("your-rule-id"), // Replace with your WAF rule ID
            Type:     aws.String("REGULAR"),
        },
        // Add more rules as needed
    }

    // Create the WebACL
    createWebAclInput := &waf.CreateWebACLInput{
        Name:          aws.String(webAclName),
        MetricName:    aws.String("MyWebAclMetricName"), // Replace with your desired metric name
        DefaultAction: &waf.Action{
            Type: aws.String("ALLOW"), // You can customize the default action
        },
        Rules: webAclRules,
    }

    result, err := wafSvc.CreateWebACL(createWebAclInput)
    if err != nil {
        fmt.Println("Error creating WebACL:", err)
        return
    }

    fmt.Println("Created WebACL with ID:", *result.WebACL.WebACLId)
}

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	client := s3control.NewFromConfig(cfg)

	out, err := client.PutStorageLensConfiguration(context.TODO(), &s3control.PutStorageLensConfigurationInput{
		AccountId: aws.String("REPLACE WITH YOUR ACCOUNT ID"),
		ConfigId:  aws.String("gotest"),
		StorageLensConfiguration: &types.StorageLensConfiguration{
			Id: aws.String("gotest"),
			AccountLevel: &types.AccountLevel{
				AdvancedCostOptimizationMetrics: &types.AdvancedCostOptimizationMetrics{IsEnabled: false}, // <-- THIS IS FALSE
				BucketLevel: &types.BucketLevel{
					AdvancedCostOptimizationMetrics: &types.AdvancedCostOptimizationMetrics{IsEnabled: false}, // <-- FALSE HERE TOO
				},
			},
		},
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%v\n", out)
}

package main

import (
	"github.com/pulumi/pulumi-ibm-api-connect/sdk/go/ibm-api-connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myRandomResource, err := ibm - api - connect.NewRandom(ctx, "myRandomResource", &ibm-api-connect.RandomArgs{
			Length: pulumi.Int(24),
		})
		if err != nil {
			return err
		}
		ctx.Export("output", map[string]interface{}{
			"value": myRandomResource.Result,
		})
		return nil
	})
}

package main

import (
	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a random pet name
		serverRandomPet, err := random.NewRandomPet(ctx, "random-name", nil)
		if err != nil {
			return err
		}

		// Create a random string
		randomString, err := random.NewRandomString(ctx, "random", &random.RandomStringArgs{
			Length:  pulumi.Int(6),
			Special: pulumi.Bool(false),
		})
		if err != nil {
			return err
		}

		// Create an Azure AD user
		user, err := azuread.NewUser(ctx, "golang", &azuread.UserArgs{
			DisplayName:       serverRandomPet.ID(),
			MailNickname:      randomString.Result,
			Password:          pulumi.String("SecretP@sswd99!"),
			UserPrincipalName: pulumi.Sprintf("%s@pulumici.onmicrosoft.com", randomString.Result),
		})
		if err != nil {
			return err
		}

		// Export the user ID
		ctx.Export("userid", user.ID())
		return nil
	})
}

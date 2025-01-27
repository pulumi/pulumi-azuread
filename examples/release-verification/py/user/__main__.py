import pulumi
import pulumi_azuread as azuread
import pulumi_random as random

# Create a random pet name
server_random_pet = random.RandomPet("random-name")

# Create a random string
random_string = random.RandomString("random", length=6, special=False)

# Create an Azure AD user
user = azuread.User("python",
                    display_name=server_random_pet.id,
                    mail_nickname=random_string.result,
                    password="SecretP@sswd99!",
                    user_principal_name=pulumi.Output.concat(random_string.result, "@pulumici.onmicrosoft.com"))

# Export the user ID
pulumi.export("userid", user.id)

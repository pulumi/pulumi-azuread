import * as azuread from "@pulumi/azuread";
import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";

const serverRandomPet = new random.RandomPet("random-name");

const randomString = new random.RandomString("random", {
    length: 6,
    special: false,
});

const user = new azuread.User("me", {
    displayName: serverRandomPet.id,
    mailNickname: randomString.result,
    password: "SecretP@sswd99!",
    userPrincipalName: pulumi.interpolate`${randomString.result}@pulumici.onmicrosoft.com`,
});

export const userid: pulumi.Output<string> = user.id;

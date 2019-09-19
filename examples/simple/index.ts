// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

import * as azuread from "@pulumi/azuread";
import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";

const config = new pulumi.Config();
const password = config.require("password");

const serverRandomPet = new random.RandomPet("random-name");

const randomString = new random.RandomString("random", {
    length: 6,
    special: false,
});

const user = new azuread.User("me", {
    displayName: serverRandomPet.id,
    mailNickname: randomString.result,
    password: password,
    userPrincipalName: pulumi.interpolate`${randomString.result}@pulumi.com`,
});

export const userid: pulumi.Output<string> = user.id;

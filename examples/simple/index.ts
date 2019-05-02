// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

import * as azuread from "@pulumi/azuread";
import * as pulumi from "@pulumi/pulumi";

const user = new azuread.User("me", {
    displayName: "John Doe",
    mailNickname: "johnd",
    password: "SecretP@sswd99!",
    userPrincipalName: "johndoe@pulumi.com",
});

export const userid: pulumi.Output<string> = user.id;
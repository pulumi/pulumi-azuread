[![Build Status](https://travis-ci.com/pulumi/pulumi-azuread.svg?token=eHg7Zp5zdDDJfTjY8ejq&branch=master)](https://travis-ci.com/pulumi/pulumi-azuread)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fazuread.svg)](https://npmjs.com/package/@pulumi/azuread)
[![Python version](https://badge.fury.io/py/pulumi-azuread.svg)](https://pypi.org/project/pulumi-azuread)
[![GoDoc](https://godoc.org/github.com/pulumi/pulumi-azuread?status.svg)](https://godoc.org/github.com/pulumi/pulumi-azuread)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-azuread/blob/master/LICENSE)

# Microsoft Azure Active Directory Resource Provider

The Microsoft Azure AD resource provider for Pulumi lets you use Azure Active Directory resources in your cloud programs.  To use
this package, please [install the Pulumi CLI first](https://pulumi.io/). For a streamlined Pulumi walkthrough, including language runtime installation and Azure configuration, click "Get Started" below.

<div>
    <a href="https://www.pulumi.com/docs/get-started/azure" title="Get Started">
       <img src="https://www.pulumi.com/images/get-started.svg" width="120">
    </a>
</div>

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/azuread

or `yarn`:

    $ yarn add @pulumi/azuread

### Python 3

To use from Python, install using `pip`:

    $ pip install pulumi_azuread

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-azuread/sdk/go/...

## Configuration

The following configuration points are available:

- `azuread:clientId` - The Client ID which should be used. This can also be sourced from the `ARM_CLIENT_ID` Environment 
   Variable.
- `azuread:subscriptionId` - The Subscription ID which should be used. This can also be sourced from the `ARM_SUBSCRIPTION_ID` 
   Environment Variable.
- `azuread:tenantId` - The Tenant ID which should be used. This can also be sourced from the `ARM_TENANT_ID` Environment 
   Variable.
- `azuread:clientSecret` - The Client Secret which should be used. This can also be sourced from the `ARM_CLIENT_SECRET` 
   Environment Variable.
- `azuread:certificatePassword` - The password associated with the Client Certificate. This can also be sourced from 
   the `ARM_CLIENT_CERTIFICATE_PASSWORD` Environment Variable.
- `azuread:clientCertificatePath` - The path to the Client Certificate associated with the Service Principal which should 
   be used. This can also be sourced from the `ARM_CLIENT_CERTIFICATE_PATH` Environment Variable.
- `azuread:environment` -  The Cloud Environment which be used. Possible values are public, usgovernment, german and china. 
   Defaults to `public`. This can also be sourced from the `ARM_ENVIRONMENT` environment variable.
- `azuread:msiEndpoint` - The path to a custom endpoint for Managed Service Identity - in most circumstances this should
   be detected automatically. This can also be sourced from the `ARM_MSI_ENDPOINT` Environment Variable.
- `azuread:useMsi` - Should Managed Service Identity be used for Authentication? This can also be sourced from the 
   `ARM_USE_MSI` Environment Variable. Defaults to `false`.
   
## Reference

For detailed reference documentation, please visit [the API docs][1].

## Build from source

### Add dependencies

In the root of the repository, run:

- Download the `install-common-toolchain.sh` script from [here](https://github.com/pulumi/scripts/blob/master/ci/install-common-toolchain.sh) (or clone the repo) and run it in your terminal.
- `make ensure`

### Build the provider:

- `make`


[1]: https://pulumi.io/reference/pkg/nodejs/@pulumi/azuread/index.html

[![Actions Status](https://github.com/pulumi/pulumi-azuread/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-azuread/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fazuread.svg)](https://npmjs.com/package/@pulumi/azuread)
[![NuGet version](https://badge.fury.io/nu/pulumi.azuread.svg)](https://badge.fury.io/nu/pulumi.azured)
[![Python version](https://badge.fury.io/py/pulumi-azuread.svg)](https://pypi.org/project/pulumi-azuread)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-azuread/sdk/v6/go)](https://pkg.go.dev/github.com/pulumi/pulumi-azuread/sdk/v6/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-azuread/blob/master/LICENSE)

# Microsoft Azure Active Directory Resource Provider

The Microsoft Azure AD resource provider for Pulumi lets you use Azure Active Directory resources in your cloud programs.  To use
this package, please [install the Pulumi CLI first](https://pulumi.io/). For a streamlined Pulumi walkthrough, including language runtime installation and Azure configuration, click "Get Started" below.

<div>
    <a href="https://www.pulumi.com/docs/get-started/azure" title="Get Started">
       <img src="https://www.pulumi.com/images/get-started.svg?" width="120">
    </a>
</div>

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    npm install @pulumi/azuread

or `yarn`:

    yarn add @pulumi/azuread

### Python 3

To use from Python, install using `pip`:

    pip install pulumi-azuread

### Go

To use from Go, use `go get` to grab the latest version of the library

    go get github.com/pulumi/pulumi-azuread/sdk/v6

### .NET

To use from .NET, install using `dotnet add package`:

    dotnet add package Pulumi.Azuread

## Configuration

The following configuration points are available:

- `azuread:clientId` - The Client ID which should be used. This can also be sourced from the `ARM_CLIENT_ID` Environment
   Variable.
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

For further information, please visit [the AzureAD provider docs](https://www.pulumi.com/registry/packages/azuread/) or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/registry/packages/azuread/api-docs/).

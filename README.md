[![Build Status](https://travis-ci.com/pulumi/pulumi-azure.svg?token=eHg7Zp5zdDDJfTjY8ejq&branch=master)](https://travis-ci.com/pulumi/pulumi-azure)

# Microsoft Azure Resource Provider

The Microsoft Azure resource provider for Pulumi lets you use Azure resources in your cloud programs.  To use
this package, please [install the Pulumi CLI first](https://pulumi.io/).

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

## Build from source

### Add dependencies

In the root of the repository, run:

- Download the `install-common-toolchain.sh` script from [here](https://github.com/pulumi/scripts/blob/master/ci/install-common-toolchain.sh) (or clone the repo) and run it in your terminal.
- `GO111MODULE=on go get github.com/pulumi/pulumi-terraform@master`
- `GO111MODULE=on go get github.com/terraform-providers/terraform-provider-azuread`
- `GO111MODULE=on go mod vendor`
- `make ensure`

### Build the provider:

- Edit `resources.go` to map each resource, and specify provider information
- Enumerate any examples in `examples/examples_test.go`
- `make`

## Reference

For detailed reference documentation, please visit [the API docs][1].


[1]: https://pulumi.io/reference/pkg/nodejs/@pulumi/azuread/index.html

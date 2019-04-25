Terraform Bridge Provider Boilerplate
=====================================

This repository contains boilerplate code for building a new Pulumi
provider which wraps an existing Terraform provider, if the existing
provider uses *Go Modules*.

Modify this README to describe:

-  The type of resources the provider manages
-  Add a build status image from Travis at the top of the README
-  Update package names in the information below
-  Add any important documentation of concepts (e.g. the “serverless”
   components in the AWS provider).

Creating a Pulumi Terraform Bridge Provider
-------------------------------------------

*Note: Go 1.12 is needed to build Pulumi providers using Go Modules.
Currently, we recommend pinning the version in ``.travis.yml`` to
``1.12.1`` to work around an issue with running later versions on Travis
CI.*

First, clone this repo with the name of the desired provider in place of
``azuread``:

::

   git clone https://github.com/pulumi/pulumi-tf-provider-boilerplate pulumi-azuread

Next, replace references to ``azuread`` with the name of your provider:
- Search/replace the string ``azuread`` with the name of your provider
throughout this repo - Rename the
``cmd/pulumi-{resource,tfgen}-azuread`` directories to match the
provider name - Replace the module name in ``go.mod`` to reflect the
repository name. - If the pulumi provider name differs from the
Terraform provider name, set ``TF_NAME`` in ``Makefile`` to the
Terraform name, leaving ``PACK`` set to the Pulumi name.

   Note: If the name of the desired Pulumi provider differs from the
   name of the Terraform provider, you will need to carefully
   distinguish between the references - see
   https://github.com/pulumi/pulumi-azure for an example.

Add dependencies
~~~~~~~~~~~~~~~~

In the root of the repository, run:

-  ``GO111MODULE=on go get github.com/pulumi/pulumi-terraform@master``
-  Install ``bazaar`` using ``brew install bazaar`` or your favorite
   package manager.
-  ``GO111MODULE=on go get github.com/terraform-providers/terraform-provider-azuread``
   (where ``azuread`` is the name of the provider)

   -  If this fails with an error like this:

   ::

      # contrib.go.opencensus.io/exporter/ocagent ../../../../pkg/mod/contrib.go.opencensus.io/exporter/ocagent@v0.4.2/transform_stats_to_metrics.go:210:61: cannot use data.ExemplarsPerBucket (type []*metricdata.Exemplar) as type []*exemplar.Exemplar in argument to bucketsToProtoBucket

   -  ..then copy the ``replace`` section from
      `this <https://github.com/pulumi/pulumi-azure/blob/master/go.mod>`__
      into the ``go.mod`` file in this repo. It should look something
      like this:

   ::

      replace (
        contrib.go.opencensus.io/exporter/ocagent => contrib.go.opencensus.io/exporter/ocagent v0.4.12
        github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
        github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2-0.20190403091019-9b3cdde74fbe
        github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
        github.com/terraform-providers/terraform-provider-azurerm => github.com/pulumi/terraform-provider-azurerm v0.0.0-20190417123607-dd01e8265e07
      )

-  ``GO111MODULE=on go mod vendor``
-  ``make ensure``

Build the provider:
~~~~~~~~~~~~~~~~~~~

-  Edit ``resources.go`` to map each resource, and specify provider
   information
-  Enumerate any examples in ``examples/examples_test.go``
-  ``make``

Installing
----------

This package is available in many languages in the standard packaging
formats.

Node.js (Java/TypeScript)
~~~~~~~~~~~~~~~~~~~~~~~~~

To use from JavaScript or TypeScript in Node.js, install using either
``npm``:

::

   $ npm install @pulumi/azuread

or ``yarn``:

::

   $ yarn add @pulumi/azuread

Python
~~~~~~

To use from Python, install using ``pip``:

::

   $ pip install pulumi_azuread

Go
~~

To use from Go, use ``go get`` to grab the latest version of the library

::

   $ go get github.com/pulumi/pulumi-azuread/sdk/go/...

Reference
---------

For detailed reference documentation, please visit `the API
docs <https://pulumi.io/reference/pkg/nodejs/@pulumi/x/index.html>`__.

|Build Status|

Microsoft Azure Active Directory Resource Provider
==================================================

The Microsoft Azure AD resource provider for Pulumi lets you use Azure
Active Directory resources in your cloud programs. To use this package,
please `install the Pulumi CLI first <https://pulumi.io/>`__.

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

Python 3
~~~~~~~~

To use from Python, install using ``pip``:

::

   $ pip install pulumi_azuread

Go
~~

To use from Go, use ``go get`` to grab the latest version of the library

::

   $ go get github.com/pulumi/pulumi-azuread/sdk/go/...

Build from source
-----------------

Add dependencies
~~~~~~~~~~~~~~~~

In the root of the repository, run:

-  Download the ``install-common-toolchain.sh`` script from
   `here <https://github.com/pulumi/scripts/blob/master/ci/install-common-toolchain.sh>`__
   (or clone the repo) and run it in your terminal.
-  ``make ensure``

Build the provider:
~~~~~~~~~~~~~~~~~~~

-  ``make``

Reference
---------

For detailed reference documentation, please visit `the API
docs <https://pulumi.io/reference/pkg/nodejs/@pulumi/azuread/index.html>`__.

.. |Build Status| image:: https://travis-ci.com/pulumi/pulumi-azuread.svg?token=eHg7Zp5zdDDJfTjY8ejq&branch=master
   :target: https://travis-ci.com/pulumi/pulumi-azuread

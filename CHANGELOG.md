CHANGELOG
=========

## HEAD (Unreleased)
* Upgrade to pulumi-terraform-bridge v2.19.0  
  **Please Note:** This includes a bug fix that stops mutating resources options in the nodejs provider
* Avoid storing config from the environment into the state

---

## 3.3.0 (2021-02-01)
* Upgrading pulumi-terraform-bridge to v2.18.1
* Upgrade to v1.3.0 of the AzureAD Terraform Provider

## 3.2.0 (2021-01-19)
* Upgrade to v1.2.2 of the AzureAD Terraform Provider

## 3.1.1 (2021-01-13)
* Upgrade to pulumi-terraform-bridge v2.17.0
* Upgrade to Pulumi v2.17.0

## 3.1.0 (2020-12-10)
* Upgrade to v1.1.1 of the AzureAD Terraform Provider

## 3.0.1 (2020-12-08)
* Upgrade to pulumi-terraform-bridge v2.15.3

## 3.0.0 (2020-12-08)
* Upgrade to pulumi-terraform-bridge v2.15.2
* Upgrade to v1.0.0 of the AzureAD Terraform Provider
  * *PLEASE NOTE:*  
    There are a number of breaking changes in this upgrade:
    * `azuread.applicationPassword` has deprecated the `applicationId` parameter.
    * `azuread.getGroup` the `name` property is now case-insensitive.
    * `azuread.Provider` `subscriptionId` is no longer a valid provider argument.

## 2.6.1 (2020-11-09)
* Upgrade to pulumi-terraform-bridge v2.13.0

## 2.6.0 (2020-10-23)
* Upgrade to Pulumi v2.12.0 and pulumi-terraform-bridge v2.11.0
* Improving the accuracy of previews leading to a more accurate understanding of what will actually change rather than assuming all output properties will change.  
  ** PLEASE NOTE:**  
  This new preview functionality can be disabled by setting `PULUMI_DISABLE_PROVIDER_PREVIEW` to `1` or `false`.

## 2.5.1 (2020-08-26)
* Upgrade to pulumi-terraform-bridge v2.7.3

## 2.5.0 (2020-08-24)
* Upgrade to Pulumi v2.9.0, which adds type annotations and input/output classes to Python
* Upgrade to pulumi-terraform-bridge v2.7.2

## 2.4.0 (2020-07-10)
* Upgrade to v0.11.0 of the AzureAD Terraform Provider

## 2.3.0 (2020-06-08)
* Upgrade to v0.10.0 of the AzureAD Terraform Provider

## 2.2.1 (2020-05-28)
* Upgrade to Pulumi v2.3.0
* Upgrade to pulumi-terraform-bridge v2.4.0

## 2.2.0 (2020-05-15)
* Upgrade to v0.9.0 of the AzureAD Terraform Provider

## 2.1.1 (2020-05-11)
* Upgrade to pulumi-terraform-bridge v2.3.1

## 2.1.0 (2020-04-29)
* Regenerate datasource examples to be async
* Upgrade to pulumi-terraform-bridge v2.1.0

## 2.0.0 (2020-04-16)
* Upgrade to Pulumi v2.0.0
* Upgrade to pulumi-terraform-bridge v2.0.0

## 1.8.0 (2020-03-30)
* Upgrade to Pulumi v1.13.1
* Upgrade to pulumi-terraform-bridge v1.8.4

## 1.7.0 (2020-03-23)
* Upgrade to v0.8.0 of the AzureAD Terraform Provider

## 1.6.0 (2020-03-16)
* Upgrade to Pulumi v1.12.1
* Upgrade to pulumi-terraform-bridge v1.8.2

## 1.5.0 (2020-01-29)
* Upgrade to pulumi-terraform-bridge v1.6.4

## 1.4.0 (2019-12-12)
* Upgrade to pulumi-terraform-bridge v1.5.2

## 1.3.0 (2019-12-02)
* Upgrade to pulumi-terraform-bridge v1.4.3

## 1.2.0 (2019-11-16)
* Upgrade to v0.7.0 of the AzureAD Terraform Provider

## 1.1.0 (2019-11-07)
* Upgrade to support go1.13
* Regenerate SDK against tf2pulumi 0.6.0
* Add a **preview** .NET SDK

## 1.0.0 (2019-09-03)
* Use 1.0 version of Pulumi dependency

## 1.0.0-rc.1 (2019-08-28)
* Upgrade pulumi-terraform to 3f206601e7

## 1.0.0-beta.2 (2019-08-26)
* Upgrade to pulumi-terraform@14e049b09d_
* Upgrade to v0.6.0 of the Azuread Terraform Provider
* Update pulumi-terraform reference to 58c7473d0
* Use 1.0-beta versions of python dependencies

## 1.0.0-beta.1 (2019-08-13)
* No significant changes

## 0.18.4 (2019-08-09)
* Update to pulumi-terraform@9db2fc93cd

## 0.18.3 (2019-07-24)
* Upgrade to v0.5.1 of the Azuread Terraform Provider
* Upgrade to 98fabcf5067b of pulumi-terraform

## 0.18.2 (2019-06-21)
* Update to pulumi-terraform@3635bed3a5 which stops maps containing `.` being treated as nested maps.

## 0.18.1 (2019-06-06)
* Update to v0.4.0 of the AzureAD Terraform provider
* Add TypeScript type guards for each resource class

## 0.18.0 (2019-05-02)
* Initial release of the Azure AD provider based on v0.3.1 of the AzureAD Terraform provider


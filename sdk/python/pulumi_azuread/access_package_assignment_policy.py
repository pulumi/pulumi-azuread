# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['AccessPackageAssignmentPolicyArgs', 'AccessPackageAssignmentPolicy']

@pulumi.input_type
class AccessPackageAssignmentPolicyArgs:
    def __init__(__self__, *,
                 access_package_id: pulumi.Input[str],
                 description: pulumi.Input[str],
                 display_name: pulumi.Input[str],
                 approval_settings: Optional[pulumi.Input['AccessPackageAssignmentPolicyApprovalSettingsArgs']] = None,
                 assignment_review_settings: Optional[pulumi.Input['AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs']] = None,
                 duration_in_days: Optional[pulumi.Input[int]] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 extension_enabled: Optional[pulumi.Input[bool]] = None,
                 questions: Optional[pulumi.Input[Sequence[pulumi.Input['AccessPackageAssignmentPolicyQuestionArgs']]]] = None,
                 requestor_settings: Optional[pulumi.Input['AccessPackageAssignmentPolicyRequestorSettingsArgs']] = None):
        """
        The set of arguments for constructing a AccessPackageAssignmentPolicy resource.
        :param pulumi.Input[str] access_package_id: The ID of the access package that will contain the policy.
        :param pulumi.Input[str] description: The description of the policy.
        :param pulumi.Input[str] display_name: The display name of the policy.
        :param pulumi.Input['AccessPackageAssignmentPolicyApprovalSettingsArgs'] approval_settings: An `approval_settings` block to specify whether approvals are required and how they are obtained, as documented below.
        :param pulumi.Input['AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs'] assignment_review_settings: An `assignment_review_settings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
        :param pulumi.Input[int] duration_in_days: How many days this assignment is valid for.
        :param pulumi.Input[str] expiration_date: The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
        :param pulumi.Input[bool] extension_enabled: Whether users will be able to request extension of their access to this package before their access expires.
        :param pulumi.Input[Sequence[pulumi.Input['AccessPackageAssignmentPolicyQuestionArgs']]] questions: One or more `question` blocks for the requestor, as documented below.
        :param pulumi.Input['AccessPackageAssignmentPolicyRequestorSettingsArgs'] requestor_settings: A `requestor_settings` block to configure the users who can request access, as documented below.
        """
        pulumi.set(__self__, "access_package_id", access_package_id)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "display_name", display_name)
        if approval_settings is not None:
            pulumi.set(__self__, "approval_settings", approval_settings)
        if assignment_review_settings is not None:
            pulumi.set(__self__, "assignment_review_settings", assignment_review_settings)
        if duration_in_days is not None:
            pulumi.set(__self__, "duration_in_days", duration_in_days)
        if expiration_date is not None:
            pulumi.set(__self__, "expiration_date", expiration_date)
        if extension_enabled is not None:
            pulumi.set(__self__, "extension_enabled", extension_enabled)
        if questions is not None:
            pulumi.set(__self__, "questions", questions)
        if requestor_settings is not None:
            pulumi.set(__self__, "requestor_settings", requestor_settings)

    @property
    @pulumi.getter(name="accessPackageId")
    def access_package_id(self) -> pulumi.Input[str]:
        """
        The ID of the access package that will contain the policy.
        """
        return pulumi.get(self, "access_package_id")

    @access_package_id.setter
    def access_package_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_package_id", value)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        The description of the policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The display name of the policy.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="approvalSettings")
    def approval_settings(self) -> Optional[pulumi.Input['AccessPackageAssignmentPolicyApprovalSettingsArgs']]:
        """
        An `approval_settings` block to specify whether approvals are required and how they are obtained, as documented below.
        """
        return pulumi.get(self, "approval_settings")

    @approval_settings.setter
    def approval_settings(self, value: Optional[pulumi.Input['AccessPackageAssignmentPolicyApprovalSettingsArgs']]):
        pulumi.set(self, "approval_settings", value)

    @property
    @pulumi.getter(name="assignmentReviewSettings")
    def assignment_review_settings(self) -> Optional[pulumi.Input['AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs']]:
        """
        An `assignment_review_settings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
        """
        return pulumi.get(self, "assignment_review_settings")

    @assignment_review_settings.setter
    def assignment_review_settings(self, value: Optional[pulumi.Input['AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs']]):
        pulumi.set(self, "assignment_review_settings", value)

    @property
    @pulumi.getter(name="durationInDays")
    def duration_in_days(self) -> Optional[pulumi.Input[int]]:
        """
        How many days this assignment is valid for.
        """
        return pulumi.get(self, "duration_in_days")

    @duration_in_days.setter
    def duration_in_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "duration_in_days", value)

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> Optional[pulumi.Input[str]]:
        """
        The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
        """
        return pulumi.get(self, "expiration_date")

    @expiration_date.setter
    def expiration_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_date", value)

    @property
    @pulumi.getter(name="extensionEnabled")
    def extension_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether users will be able to request extension of their access to this package before their access expires.
        """
        return pulumi.get(self, "extension_enabled")

    @extension_enabled.setter
    def extension_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "extension_enabled", value)

    @property
    @pulumi.getter
    def questions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AccessPackageAssignmentPolicyQuestionArgs']]]]:
        """
        One or more `question` blocks for the requestor, as documented below.
        """
        return pulumi.get(self, "questions")

    @questions.setter
    def questions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AccessPackageAssignmentPolicyQuestionArgs']]]]):
        pulumi.set(self, "questions", value)

    @property
    @pulumi.getter(name="requestorSettings")
    def requestor_settings(self) -> Optional[pulumi.Input['AccessPackageAssignmentPolicyRequestorSettingsArgs']]:
        """
        A `requestor_settings` block to configure the users who can request access, as documented below.
        """
        return pulumi.get(self, "requestor_settings")

    @requestor_settings.setter
    def requestor_settings(self, value: Optional[pulumi.Input['AccessPackageAssignmentPolicyRequestorSettingsArgs']]):
        pulumi.set(self, "requestor_settings", value)


@pulumi.input_type
class _AccessPackageAssignmentPolicyState:
    def __init__(__self__, *,
                 access_package_id: Optional[pulumi.Input[str]] = None,
                 approval_settings: Optional[pulumi.Input['AccessPackageAssignmentPolicyApprovalSettingsArgs']] = None,
                 assignment_review_settings: Optional[pulumi.Input['AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 duration_in_days: Optional[pulumi.Input[int]] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 extension_enabled: Optional[pulumi.Input[bool]] = None,
                 questions: Optional[pulumi.Input[Sequence[pulumi.Input['AccessPackageAssignmentPolicyQuestionArgs']]]] = None,
                 requestor_settings: Optional[pulumi.Input['AccessPackageAssignmentPolicyRequestorSettingsArgs']] = None):
        """
        Input properties used for looking up and filtering AccessPackageAssignmentPolicy resources.
        :param pulumi.Input[str] access_package_id: The ID of the access package that will contain the policy.
        :param pulumi.Input['AccessPackageAssignmentPolicyApprovalSettingsArgs'] approval_settings: An `approval_settings` block to specify whether approvals are required and how they are obtained, as documented below.
        :param pulumi.Input['AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs'] assignment_review_settings: An `assignment_review_settings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
        :param pulumi.Input[str] description: The description of the policy.
        :param pulumi.Input[str] display_name: The display name of the policy.
        :param pulumi.Input[int] duration_in_days: How many days this assignment is valid for.
        :param pulumi.Input[str] expiration_date: The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
        :param pulumi.Input[bool] extension_enabled: Whether users will be able to request extension of their access to this package before their access expires.
        :param pulumi.Input[Sequence[pulumi.Input['AccessPackageAssignmentPolicyQuestionArgs']]] questions: One or more `question` blocks for the requestor, as documented below.
        :param pulumi.Input['AccessPackageAssignmentPolicyRequestorSettingsArgs'] requestor_settings: A `requestor_settings` block to configure the users who can request access, as documented below.
        """
        if access_package_id is not None:
            pulumi.set(__self__, "access_package_id", access_package_id)
        if approval_settings is not None:
            pulumi.set(__self__, "approval_settings", approval_settings)
        if assignment_review_settings is not None:
            pulumi.set(__self__, "assignment_review_settings", assignment_review_settings)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if duration_in_days is not None:
            pulumi.set(__self__, "duration_in_days", duration_in_days)
        if expiration_date is not None:
            pulumi.set(__self__, "expiration_date", expiration_date)
        if extension_enabled is not None:
            pulumi.set(__self__, "extension_enabled", extension_enabled)
        if questions is not None:
            pulumi.set(__self__, "questions", questions)
        if requestor_settings is not None:
            pulumi.set(__self__, "requestor_settings", requestor_settings)

    @property
    @pulumi.getter(name="accessPackageId")
    def access_package_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the access package that will contain the policy.
        """
        return pulumi.get(self, "access_package_id")

    @access_package_id.setter
    def access_package_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_package_id", value)

    @property
    @pulumi.getter(name="approvalSettings")
    def approval_settings(self) -> Optional[pulumi.Input['AccessPackageAssignmentPolicyApprovalSettingsArgs']]:
        """
        An `approval_settings` block to specify whether approvals are required and how they are obtained, as documented below.
        """
        return pulumi.get(self, "approval_settings")

    @approval_settings.setter
    def approval_settings(self, value: Optional[pulumi.Input['AccessPackageAssignmentPolicyApprovalSettingsArgs']]):
        pulumi.set(self, "approval_settings", value)

    @property
    @pulumi.getter(name="assignmentReviewSettings")
    def assignment_review_settings(self) -> Optional[pulumi.Input['AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs']]:
        """
        An `assignment_review_settings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
        """
        return pulumi.get(self, "assignment_review_settings")

    @assignment_review_settings.setter
    def assignment_review_settings(self, value: Optional[pulumi.Input['AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs']]):
        pulumi.set(self, "assignment_review_settings", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the policy.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="durationInDays")
    def duration_in_days(self) -> Optional[pulumi.Input[int]]:
        """
        How many days this assignment is valid for.
        """
        return pulumi.get(self, "duration_in_days")

    @duration_in_days.setter
    def duration_in_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "duration_in_days", value)

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> Optional[pulumi.Input[str]]:
        """
        The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
        """
        return pulumi.get(self, "expiration_date")

    @expiration_date.setter
    def expiration_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_date", value)

    @property
    @pulumi.getter(name="extensionEnabled")
    def extension_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether users will be able to request extension of their access to this package before their access expires.
        """
        return pulumi.get(self, "extension_enabled")

    @extension_enabled.setter
    def extension_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "extension_enabled", value)

    @property
    @pulumi.getter
    def questions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AccessPackageAssignmentPolicyQuestionArgs']]]]:
        """
        One or more `question` blocks for the requestor, as documented below.
        """
        return pulumi.get(self, "questions")

    @questions.setter
    def questions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AccessPackageAssignmentPolicyQuestionArgs']]]]):
        pulumi.set(self, "questions", value)

    @property
    @pulumi.getter(name="requestorSettings")
    def requestor_settings(self) -> Optional[pulumi.Input['AccessPackageAssignmentPolicyRequestorSettingsArgs']]:
        """
        A `requestor_settings` block to configure the users who can request access, as documented below.
        """
        return pulumi.get(self, "requestor_settings")

    @requestor_settings.setter
    def requestor_settings(self, value: Optional[pulumi.Input['AccessPackageAssignmentPolicyRequestorSettingsArgs']]):
        pulumi.set(self, "requestor_settings", value)


class AccessPackageAssignmentPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_package_id: Optional[pulumi.Input[str]] = None,
                 approval_settings: Optional[pulumi.Input[pulumi.InputType['AccessPackageAssignmentPolicyApprovalSettingsArgs']]] = None,
                 assignment_review_settings: Optional[pulumi.Input[pulumi.InputType['AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 duration_in_days: Optional[pulumi.Input[int]] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 extension_enabled: Optional[pulumi.Input[bool]] = None,
                 questions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccessPackageAssignmentPolicyQuestionArgs']]]]] = None,
                 requestor_settings: Optional[pulumi.Input[pulumi.InputType['AccessPackageAssignmentPolicyRequestorSettingsArgs']]] = None,
                 __props__=None):
        """
        Manages an assignment policy for an access package within Identity Governance in Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.

        When authenticated with a user principal, this resource requires `Global Administrator` directory role, or one of the `Catalog Owner` and `Access Package Manager` role in Identity Governance.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_group = azuread.Group("exampleGroup",
            display_name="group-name",
            security_enabled=True)
        example_access_package_catalog = azuread.AccessPackageCatalog("exampleAccessPackageCatalog",
            display_name="example-catalog",
            description="Example catalog")
        example_access_package = azuread.AccessPackage("exampleAccessPackage",
            catalog_id=example_access_package_catalog.id,
            display_name="access-package",
            description="Access Package")
        example_access_package_assignment_policy = azuread.AccessPackageAssignmentPolicy("exampleAccessPackageAssignmentPolicy",
            access_package_id=example_access_package.id,
            display_name="assignment-policy",
            description="My assignment policy",
            duration_in_days=90,
            requestor_settings=azuread.AccessPackageAssignmentPolicyRequestorSettingsArgs(
                scope_type="AllExistingDirectoryMemberUsers",
            ),
            approval_settings=azuread.AccessPackageAssignmentPolicyApprovalSettingsArgs(
                approval_required=True,
                approval_stages=[azuread.AccessPackageAssignmentPolicyApprovalSettingsApprovalStageArgs(
                    approval_timeout_in_days=14,
                    primary_approvers=[azuread.AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverArgs(
                        object_id=example_group.object_id,
                        subject_type="groupMembers",
                    )],
                )],
            ),
            assignment_review_settings=azuread.AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs(
                enabled=True,
                review_frequency="weekly",
                duration_in_days=3,
                review_type="Self",
                access_review_timeout_behavior="keepAccess",
            ),
            questions=[azuread.AccessPackageAssignmentPolicyQuestionArgs(
                text=azuread.AccessPackageAssignmentPolicyQuestionTextArgs(
                    default_text="hello, how are you?",
                ),
            )])
        ```

        ## Import

        An access package assignment policy can be imported using the ID, e.g.

        ```sh
         $ pulumi import azuread:index/accessPackageAssignmentPolicy:AccessPackageAssignmentPolicy example 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_package_id: The ID of the access package that will contain the policy.
        :param pulumi.Input[pulumi.InputType['AccessPackageAssignmentPolicyApprovalSettingsArgs']] approval_settings: An `approval_settings` block to specify whether approvals are required and how they are obtained, as documented below.
        :param pulumi.Input[pulumi.InputType['AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs']] assignment_review_settings: An `assignment_review_settings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
        :param pulumi.Input[str] description: The description of the policy.
        :param pulumi.Input[str] display_name: The display name of the policy.
        :param pulumi.Input[int] duration_in_days: How many days this assignment is valid for.
        :param pulumi.Input[str] expiration_date: The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
        :param pulumi.Input[bool] extension_enabled: Whether users will be able to request extension of their access to this package before their access expires.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccessPackageAssignmentPolicyQuestionArgs']]]] questions: One or more `question` blocks for the requestor, as documented below.
        :param pulumi.Input[pulumi.InputType['AccessPackageAssignmentPolicyRequestorSettingsArgs']] requestor_settings: A `requestor_settings` block to configure the users who can request access, as documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccessPackageAssignmentPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an assignment policy for an access package within Identity Governance in Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.

        When authenticated with a user principal, this resource requires `Global Administrator` directory role, or one of the `Catalog Owner` and `Access Package Manager` role in Identity Governance.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_group = azuread.Group("exampleGroup",
            display_name="group-name",
            security_enabled=True)
        example_access_package_catalog = azuread.AccessPackageCatalog("exampleAccessPackageCatalog",
            display_name="example-catalog",
            description="Example catalog")
        example_access_package = azuread.AccessPackage("exampleAccessPackage",
            catalog_id=example_access_package_catalog.id,
            display_name="access-package",
            description="Access Package")
        example_access_package_assignment_policy = azuread.AccessPackageAssignmentPolicy("exampleAccessPackageAssignmentPolicy",
            access_package_id=example_access_package.id,
            display_name="assignment-policy",
            description="My assignment policy",
            duration_in_days=90,
            requestor_settings=azuread.AccessPackageAssignmentPolicyRequestorSettingsArgs(
                scope_type="AllExistingDirectoryMemberUsers",
            ),
            approval_settings=azuread.AccessPackageAssignmentPolicyApprovalSettingsArgs(
                approval_required=True,
                approval_stages=[azuread.AccessPackageAssignmentPolicyApprovalSettingsApprovalStageArgs(
                    approval_timeout_in_days=14,
                    primary_approvers=[azuread.AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverArgs(
                        object_id=example_group.object_id,
                        subject_type="groupMembers",
                    )],
                )],
            ),
            assignment_review_settings=azuread.AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs(
                enabled=True,
                review_frequency="weekly",
                duration_in_days=3,
                review_type="Self",
                access_review_timeout_behavior="keepAccess",
            ),
            questions=[azuread.AccessPackageAssignmentPolicyQuestionArgs(
                text=azuread.AccessPackageAssignmentPolicyQuestionTextArgs(
                    default_text="hello, how are you?",
                ),
            )])
        ```

        ## Import

        An access package assignment policy can be imported using the ID, e.g.

        ```sh
         $ pulumi import azuread:index/accessPackageAssignmentPolicy:AccessPackageAssignmentPolicy example 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param AccessPackageAssignmentPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccessPackageAssignmentPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_package_id: Optional[pulumi.Input[str]] = None,
                 approval_settings: Optional[pulumi.Input[pulumi.InputType['AccessPackageAssignmentPolicyApprovalSettingsArgs']]] = None,
                 assignment_review_settings: Optional[pulumi.Input[pulumi.InputType['AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 duration_in_days: Optional[pulumi.Input[int]] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 extension_enabled: Optional[pulumi.Input[bool]] = None,
                 questions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccessPackageAssignmentPolicyQuestionArgs']]]]] = None,
                 requestor_settings: Optional[pulumi.Input[pulumi.InputType['AccessPackageAssignmentPolicyRequestorSettingsArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccessPackageAssignmentPolicyArgs.__new__(AccessPackageAssignmentPolicyArgs)

            if access_package_id is None and not opts.urn:
                raise TypeError("Missing required property 'access_package_id'")
            __props__.__dict__["access_package_id"] = access_package_id
            __props__.__dict__["approval_settings"] = approval_settings
            __props__.__dict__["assignment_review_settings"] = assignment_review_settings
            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["duration_in_days"] = duration_in_days
            __props__.__dict__["expiration_date"] = expiration_date
            __props__.__dict__["extension_enabled"] = extension_enabled
            __props__.__dict__["questions"] = questions
            __props__.__dict__["requestor_settings"] = requestor_settings
        super(AccessPackageAssignmentPolicy, __self__).__init__(
            'azuread:index/accessPackageAssignmentPolicy:AccessPackageAssignmentPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_package_id: Optional[pulumi.Input[str]] = None,
            approval_settings: Optional[pulumi.Input[pulumi.InputType['AccessPackageAssignmentPolicyApprovalSettingsArgs']]] = None,
            assignment_review_settings: Optional[pulumi.Input[pulumi.InputType['AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            duration_in_days: Optional[pulumi.Input[int]] = None,
            expiration_date: Optional[pulumi.Input[str]] = None,
            extension_enabled: Optional[pulumi.Input[bool]] = None,
            questions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccessPackageAssignmentPolicyQuestionArgs']]]]] = None,
            requestor_settings: Optional[pulumi.Input[pulumi.InputType['AccessPackageAssignmentPolicyRequestorSettingsArgs']]] = None) -> 'AccessPackageAssignmentPolicy':
        """
        Get an existing AccessPackageAssignmentPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_package_id: The ID of the access package that will contain the policy.
        :param pulumi.Input[pulumi.InputType['AccessPackageAssignmentPolicyApprovalSettingsArgs']] approval_settings: An `approval_settings` block to specify whether approvals are required and how they are obtained, as documented below.
        :param pulumi.Input[pulumi.InputType['AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs']] assignment_review_settings: An `assignment_review_settings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
        :param pulumi.Input[str] description: The description of the policy.
        :param pulumi.Input[str] display_name: The display name of the policy.
        :param pulumi.Input[int] duration_in_days: How many days this assignment is valid for.
        :param pulumi.Input[str] expiration_date: The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
        :param pulumi.Input[bool] extension_enabled: Whether users will be able to request extension of their access to this package before their access expires.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccessPackageAssignmentPolicyQuestionArgs']]]] questions: One or more `question` blocks for the requestor, as documented below.
        :param pulumi.Input[pulumi.InputType['AccessPackageAssignmentPolicyRequestorSettingsArgs']] requestor_settings: A `requestor_settings` block to configure the users who can request access, as documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccessPackageAssignmentPolicyState.__new__(_AccessPackageAssignmentPolicyState)

        __props__.__dict__["access_package_id"] = access_package_id
        __props__.__dict__["approval_settings"] = approval_settings
        __props__.__dict__["assignment_review_settings"] = assignment_review_settings
        __props__.__dict__["description"] = description
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["duration_in_days"] = duration_in_days
        __props__.__dict__["expiration_date"] = expiration_date
        __props__.__dict__["extension_enabled"] = extension_enabled
        __props__.__dict__["questions"] = questions
        __props__.__dict__["requestor_settings"] = requestor_settings
        return AccessPackageAssignmentPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessPackageId")
    def access_package_id(self) -> pulumi.Output[str]:
        """
        The ID of the access package that will contain the policy.
        """
        return pulumi.get(self, "access_package_id")

    @property
    @pulumi.getter(name="approvalSettings")
    def approval_settings(self) -> pulumi.Output[Optional['outputs.AccessPackageAssignmentPolicyApprovalSettings']]:
        """
        An `approval_settings` block to specify whether approvals are required and how they are obtained, as documented below.
        """
        return pulumi.get(self, "approval_settings")

    @property
    @pulumi.getter(name="assignmentReviewSettings")
    def assignment_review_settings(self) -> pulumi.Output[Optional['outputs.AccessPackageAssignmentPolicyAssignmentReviewSettings']]:
        """
        An `assignment_review_settings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
        """
        return pulumi.get(self, "assignment_review_settings")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name of the policy.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="durationInDays")
    def duration_in_days(self) -> pulumi.Output[Optional[int]]:
        """
        How many days this assignment is valid for.
        """
        return pulumi.get(self, "duration_in_days")

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> pulumi.Output[Optional[str]]:
        """
        The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
        """
        return pulumi.get(self, "expiration_date")

    @property
    @pulumi.getter(name="extensionEnabled")
    def extension_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether users will be able to request extension of their access to this package before their access expires.
        """
        return pulumi.get(self, "extension_enabled")

    @property
    @pulumi.getter
    def questions(self) -> pulumi.Output[Optional[Sequence['outputs.AccessPackageAssignmentPolicyQuestion']]]:
        """
        One or more `question` blocks for the requestor, as documented below.
        """
        return pulumi.get(self, "questions")

    @property
    @pulumi.getter(name="requestorSettings")
    def requestor_settings(self) -> pulumi.Output[Optional['outputs.AccessPackageAssignmentPolicyRequestorSettings']]:
        """
        A `requestor_settings` block to configure the users who can request access, as documented below.
        """
        return pulumi.get(self, "requestor_settings")


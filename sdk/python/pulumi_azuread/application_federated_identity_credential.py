# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ApplicationFederatedIdentityCredentialArgs', 'ApplicationFederatedIdentityCredential']

@pulumi.input_type
class ApplicationFederatedIdentityCredentialArgs:
    def __init__(__self__, *,
                 application_object_id: pulumi.Input[str],
                 audiences: pulumi.Input[Sequence[pulumi.Input[str]]],
                 display_name: pulumi.Input[str],
                 issuer: pulumi.Input[str],
                 subject: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ApplicationFederatedIdentityCredential resource.
        :param pulumi.Input[str] application_object_id: The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] audiences: List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
        :param pulumi.Input[str] display_name: A unique display name for the federated identity credential. Changing this forces a new resource to be created.
        :param pulumi.Input[str] issuer: The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
        :param pulumi.Input[str] subject: The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
        :param pulumi.Input[str] description: A description for the federated identity credential.
        """
        pulumi.set(__self__, "application_object_id", application_object_id)
        pulumi.set(__self__, "audiences", audiences)
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "issuer", issuer)
        pulumi.set(__self__, "subject", subject)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="applicationObjectId")
    def application_object_id(self) -> pulumi.Input[str]:
        """
        The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "application_object_id")

    @application_object_id.setter
    def application_object_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "application_object_id", value)

    @property
    @pulumi.getter
    def audiences(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
        """
        return pulumi.get(self, "audiences")

    @audiences.setter
    def audiences(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "audiences", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        A unique display name for the federated identity credential. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def issuer(self) -> pulumi.Input[str]:
        """
        The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
        """
        return pulumi.get(self, "issuer")

    @issuer.setter
    def issuer(self, value: pulumi.Input[str]):
        pulumi.set(self, "issuer", value)

    @property
    @pulumi.getter
    def subject(self) -> pulumi.Input[str]:
        """
        The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
        """
        return pulumi.get(self, "subject")

    @subject.setter
    def subject(self, value: pulumi.Input[str]):
        pulumi.set(self, "subject", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the federated identity credential.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _ApplicationFederatedIdentityCredentialState:
    def __init__(__self__, *,
                 application_object_id: Optional[pulumi.Input[str]] = None,
                 audiences: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 credential_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 subject: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApplicationFederatedIdentityCredential resources.
        :param pulumi.Input[str] application_object_id: The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] audiences: List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
        :param pulumi.Input[str] credential_id: A UUID used to uniquely identify this federated identity credential.
        :param pulumi.Input[str] description: A description for the federated identity credential.
        :param pulumi.Input[str] display_name: A unique display name for the federated identity credential. Changing this forces a new resource to be created.
        :param pulumi.Input[str] issuer: The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
        :param pulumi.Input[str] subject: The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
        """
        if application_object_id is not None:
            pulumi.set(__self__, "application_object_id", application_object_id)
        if audiences is not None:
            pulumi.set(__self__, "audiences", audiences)
        if credential_id is not None:
            pulumi.set(__self__, "credential_id", credential_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if issuer is not None:
            pulumi.set(__self__, "issuer", issuer)
        if subject is not None:
            pulumi.set(__self__, "subject", subject)

    @property
    @pulumi.getter(name="applicationObjectId")
    def application_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "application_object_id")

    @application_object_id.setter
    def application_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_object_id", value)

    @property
    @pulumi.getter
    def audiences(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
        """
        return pulumi.get(self, "audiences")

    @audiences.setter
    def audiences(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "audiences", value)

    @property
    @pulumi.getter(name="credentialId")
    def credential_id(self) -> Optional[pulumi.Input[str]]:
        """
        A UUID used to uniquely identify this federated identity credential.
        """
        return pulumi.get(self, "credential_id")

    @credential_id.setter
    def credential_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "credential_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the federated identity credential.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique display name for the federated identity credential. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def issuer(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
        """
        return pulumi.get(self, "issuer")

    @issuer.setter
    def issuer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuer", value)

    @property
    @pulumi.getter
    def subject(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
        """
        return pulumi.get(self, "subject")

    @subject.setter
    def subject(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subject", value)


class ApplicationFederatedIdentityCredential(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_object_id: Optional[pulumi.Input[str]] = None,
                 audiences: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 subject: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_application = azuread.Application("exampleApplication", display_name="example")
        example_application_federated_identity_credential = azuread.ApplicationFederatedIdentityCredential("exampleApplicationFederatedIdentityCredential",
            application_object_id=example_application.object_id,
            display_name="my-repo-deploy",
            description="Deployments for my-repo",
            audiences=["api://AzureADTokenExchange"],
            issuer="https://token.actions.githubusercontent.com",
            subject="repo:my-organization/my-repo:environment:prod")
        ```

        ## Import

        Federated Identity Credentials can be imported using the object ID of the associated application and the ID of the federated identity credential, e.g.

        ```sh
         $ pulumi import azuread:index/applicationFederatedIdentityCredential:ApplicationFederatedIdentityCredential test 00000000-0000-0000-0000-000000000000/federatedIdentityCredential/11111111-1111-1111-1111-111111111111
        ```

         -> This ID format is unique to Terraform and is composed of the application's object ID, the string "federatedIdentityCredential" and the credential ID in the format `{ObjectId}/federatedIdentityCredential/{CredentialId}`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_object_id: The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] audiences: List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
        :param pulumi.Input[str] description: A description for the federated identity credential.
        :param pulumi.Input[str] display_name: A unique display name for the federated identity credential. Changing this forces a new resource to be created.
        :param pulumi.Input[str] issuer: The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
        :param pulumi.Input[str] subject: The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationFederatedIdentityCredentialArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_application = azuread.Application("exampleApplication", display_name="example")
        example_application_federated_identity_credential = azuread.ApplicationFederatedIdentityCredential("exampleApplicationFederatedIdentityCredential",
            application_object_id=example_application.object_id,
            display_name="my-repo-deploy",
            description="Deployments for my-repo",
            audiences=["api://AzureADTokenExchange"],
            issuer="https://token.actions.githubusercontent.com",
            subject="repo:my-organization/my-repo:environment:prod")
        ```

        ## Import

        Federated Identity Credentials can be imported using the object ID of the associated application and the ID of the federated identity credential, e.g.

        ```sh
         $ pulumi import azuread:index/applicationFederatedIdentityCredential:ApplicationFederatedIdentityCredential test 00000000-0000-0000-0000-000000000000/federatedIdentityCredential/11111111-1111-1111-1111-111111111111
        ```

         -> This ID format is unique to Terraform and is composed of the application's object ID, the string "federatedIdentityCredential" and the credential ID in the format `{ObjectId}/federatedIdentityCredential/{CredentialId}`.

        :param str resource_name: The name of the resource.
        :param ApplicationFederatedIdentityCredentialArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationFederatedIdentityCredentialArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_object_id: Optional[pulumi.Input[str]] = None,
                 audiences: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 subject: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationFederatedIdentityCredentialArgs.__new__(ApplicationFederatedIdentityCredentialArgs)

            if application_object_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_object_id'")
            __props__.__dict__["application_object_id"] = application_object_id
            if audiences is None and not opts.urn:
                raise TypeError("Missing required property 'audiences'")
            __props__.__dict__["audiences"] = audiences
            __props__.__dict__["description"] = description
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            if issuer is None and not opts.urn:
                raise TypeError("Missing required property 'issuer'")
            __props__.__dict__["issuer"] = issuer
            if subject is None and not opts.urn:
                raise TypeError("Missing required property 'subject'")
            __props__.__dict__["subject"] = subject
            __props__.__dict__["credential_id"] = None
        super(ApplicationFederatedIdentityCredential, __self__).__init__(
            'azuread:index/applicationFederatedIdentityCredential:ApplicationFederatedIdentityCredential',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_object_id: Optional[pulumi.Input[str]] = None,
            audiences: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            credential_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            issuer: Optional[pulumi.Input[str]] = None,
            subject: Optional[pulumi.Input[str]] = None) -> 'ApplicationFederatedIdentityCredential':
        """
        Get an existing ApplicationFederatedIdentityCredential resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_object_id: The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] audiences: List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
        :param pulumi.Input[str] credential_id: A UUID used to uniquely identify this federated identity credential.
        :param pulumi.Input[str] description: A description for the federated identity credential.
        :param pulumi.Input[str] display_name: A unique display name for the federated identity credential. Changing this forces a new resource to be created.
        :param pulumi.Input[str] issuer: The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
        :param pulumi.Input[str] subject: The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationFederatedIdentityCredentialState.__new__(_ApplicationFederatedIdentityCredentialState)

        __props__.__dict__["application_object_id"] = application_object_id
        __props__.__dict__["audiences"] = audiences
        __props__.__dict__["credential_id"] = credential_id
        __props__.__dict__["description"] = description
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["issuer"] = issuer
        __props__.__dict__["subject"] = subject
        return ApplicationFederatedIdentityCredential(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationObjectId")
    def application_object_id(self) -> pulumi.Output[str]:
        """
        The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "application_object_id")

    @property
    @pulumi.getter
    def audiences(self) -> pulumi.Output[Sequence[str]]:
        """
        List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
        """
        return pulumi.get(self, "audiences")

    @property
    @pulumi.getter(name="credentialId")
    def credential_id(self) -> pulumi.Output[str]:
        """
        A UUID used to uniquely identify this federated identity credential.
        """
        return pulumi.get(self, "credential_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description for the federated identity credential.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        A unique display name for the federated identity credential. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def issuer(self) -> pulumi.Output[str]:
        """
        The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
        """
        return pulumi.get(self, "issuer")

    @property
    @pulumi.getter
    def subject(self) -> pulumi.Output[str]:
        """
        The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
        """
        return pulumi.get(self, "subject")


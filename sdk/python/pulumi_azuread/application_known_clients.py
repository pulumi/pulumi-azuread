# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ApplicationKnownClientsArgs', 'ApplicationKnownClients']

@pulumi.input_type
class ApplicationKnownClientsArgs:
    def __init__(__self__, *,
                 application_id: pulumi.Input[str],
                 known_client_ids: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        The set of arguments for constructing a ApplicationKnownClients resource.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] known_client_ids: A set of client IDs for the known applications.
        """
        pulumi.set(__self__, "application_id", application_id)
        pulumi.set(__self__, "known_client_ids", known_client_ids)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Input[str]:
        """
        The resource ID of the application registration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="knownClientIds")
    def known_client_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A set of client IDs for the known applications.
        """
        return pulumi.get(self, "known_client_ids")

    @known_client_ids.setter
    def known_client_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "known_client_ids", value)


@pulumi.input_type
class _ApplicationKnownClientsState:
    def __init__(__self__, *,
                 application_id: Optional[pulumi.Input[str]] = None,
                 known_client_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ApplicationKnownClients resources.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] known_client_ids: A set of client IDs for the known applications.
        """
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if known_client_ids is not None:
            pulumi.set(__self__, "known_client_ids", known_client_ids)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource ID of the application registration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="knownClientIds")
    def known_client_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of client IDs for the known applications.
        """
        return pulumi.get(self, "known_client_ids")

    @known_client_ids.setter
    def known_client_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "known_client_ids", value)


class ApplicationKnownClients(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 known_client_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_application_registration = azuread.ApplicationRegistration("exampleApplicationRegistration", display_name="example")
        client = azuread.ApplicationRegistration("client", display_name="example client")
        example_application_known_clients = azuread.ApplicationKnownClients("exampleApplicationKnownClients",
            application_id=example_application_registration.id,
            known_client_ids=[client.client_id])
        ```

        ## Import

        Application Known Clients can be imported using the object ID of the application in the following format.

        ```sh
        $ pulumi import azuread:index/applicationKnownClients:ApplicationKnownClients example /applications/00000000-0000-0000-0000-000000000000/knownClients
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] known_client_ids: A set of client IDs for the known applications.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationKnownClientsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_application_registration = azuread.ApplicationRegistration("exampleApplicationRegistration", display_name="example")
        client = azuread.ApplicationRegistration("client", display_name="example client")
        example_application_known_clients = azuread.ApplicationKnownClients("exampleApplicationKnownClients",
            application_id=example_application_registration.id,
            known_client_ids=[client.client_id])
        ```

        ## Import

        Application Known Clients can be imported using the object ID of the application in the following format.

        ```sh
        $ pulumi import azuread:index/applicationKnownClients:ApplicationKnownClients example /applications/00000000-0000-0000-0000-000000000000/knownClients
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationKnownClientsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationKnownClientsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 known_client_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationKnownClientsArgs.__new__(ApplicationKnownClientsArgs)

            if application_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_id'")
            __props__.__dict__["application_id"] = application_id
            if known_client_ids is None and not opts.urn:
                raise TypeError("Missing required property 'known_client_ids'")
            __props__.__dict__["known_client_ids"] = known_client_ids
        super(ApplicationKnownClients, __self__).__init__(
            'azuread:index/applicationKnownClients:ApplicationKnownClients',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            known_client_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'ApplicationKnownClients':
        """
        Get an existing ApplicationKnownClients resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] known_client_ids: A set of client IDs for the known applications.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationKnownClientsState.__new__(_ApplicationKnownClientsState)

        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["known_client_ids"] = known_client_ids
        return ApplicationKnownClients(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        The resource ID of the application registration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="knownClientIds")
    def known_client_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        A set of client IDs for the known applications.
        """
        return pulumi.get(self, "known_client_ids")


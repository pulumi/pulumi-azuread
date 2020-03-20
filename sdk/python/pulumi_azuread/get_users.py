# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetUsersResult:
    """
    A collection of values returned by getUsers.
    """
    def __init__(__self__, id=None, mail_nicknames=None, object_ids=None, user_principal_names=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if mail_nicknames and not isinstance(mail_nicknames, list):
            raise TypeError("Expected argument 'mail_nicknames' to be a list")
        __self__.mail_nicknames = mail_nicknames
        """
        The email aliases of the Azure AD Users.
        """
        if object_ids and not isinstance(object_ids, list):
            raise TypeError("Expected argument 'object_ids' to be a list")
        __self__.object_ids = object_ids
        """
        The Object IDs of the Azure AD Users.
        """
        if user_principal_names and not isinstance(user_principal_names, list):
            raise TypeError("Expected argument 'user_principal_names' to be a list")
        __self__.user_principal_names = user_principal_names
        """
        The User Principal Names of the Azure AD Users.
        """
class AwaitableGetUsersResult(GetUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUsersResult(
            id=self.id,
            mail_nicknames=self.mail_nicknames,
            object_ids=self.object_ids,
            user_principal_names=self.user_principal_names)

def get_users(mail_nicknames=None,object_ids=None,user_principal_names=None,opts=None):
    """
    Gets Object IDs or UPNs for multiple Azure Active Directory users.

    > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-azuread/blob/master/website/docs/d/users.html.markdown.


    :param list mail_nicknames: The email aliases of the Azure AD Users.
    :param list object_ids: The Object IDs of the Azure AD Users.
    :param list user_principal_names: The User Principal Names of the Azure AD Users.
    """
    __args__ = dict()


    __args__['mailNicknames'] = mail_nicknames
    __args__['objectIds'] = object_ids
    __args__['userPrincipalNames'] = user_principal_names
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azuread:index/getUsers:getUsers', __args__, opts=opts).value

    return AwaitableGetUsersResult(
        id=__ret__.get('id'),
        mail_nicknames=__ret__.get('mailNicknames'),
        object_ids=__ret__.get('objectIds'),
        user_principal_names=__ret__.get('userPrincipalNames'))

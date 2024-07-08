# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs

__all__ = [
    'GetUsersResult',
    'AwaitableGetUsersResult',
    'get_users',
    'get_users_output',
]

@pulumi.output_type
class GetUsersResult:
    """
    A collection of values returned by getUsers.
    """
    def __init__(__self__, employee_ids=None, id=None, ignore_missing=None, mail_nicknames=None, mails=None, object_ids=None, return_all=None, user_principal_names=None, users=None):
        if employee_ids and not isinstance(employee_ids, list):
            raise TypeError("Expected argument 'employee_ids' to be a list")
        pulumi.set(__self__, "employee_ids", employee_ids)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ignore_missing and not isinstance(ignore_missing, bool):
            raise TypeError("Expected argument 'ignore_missing' to be a bool")
        pulumi.set(__self__, "ignore_missing", ignore_missing)
        if mail_nicknames and not isinstance(mail_nicknames, list):
            raise TypeError("Expected argument 'mail_nicknames' to be a list")
        pulumi.set(__self__, "mail_nicknames", mail_nicknames)
        if mails and not isinstance(mails, list):
            raise TypeError("Expected argument 'mails' to be a list")
        pulumi.set(__self__, "mails", mails)
        if object_ids and not isinstance(object_ids, list):
            raise TypeError("Expected argument 'object_ids' to be a list")
        pulumi.set(__self__, "object_ids", object_ids)
        if return_all and not isinstance(return_all, bool):
            raise TypeError("Expected argument 'return_all' to be a bool")
        pulumi.set(__self__, "return_all", return_all)
        if user_principal_names and not isinstance(user_principal_names, list):
            raise TypeError("Expected argument 'user_principal_names' to be a list")
        pulumi.set(__self__, "user_principal_names", user_principal_names)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter(name="employeeIds")
    def employee_ids(self) -> Sequence[str]:
        """
        The employee identifiers assigned to the users by the organisation.
        """
        return pulumi.get(self, "employee_ids")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ignoreMissing")
    def ignore_missing(self) -> Optional[bool]:
        return pulumi.get(self, "ignore_missing")

    @property
    @pulumi.getter(name="mailNicknames")
    def mail_nicknames(self) -> Sequence[str]:
        """
        The email aliases of the users.
        """
        return pulumi.get(self, "mail_nicknames")

    @property
    @pulumi.getter
    def mails(self) -> Sequence[str]:
        """
        The SMTP email addresses of the users.
        """
        return pulumi.get(self, "mails")

    @property
    @pulumi.getter(name="objectIds")
    def object_ids(self) -> Sequence[str]:
        """
        The object IDs of the users.
        """
        return pulumi.get(self, "object_ids")

    @property
    @pulumi.getter(name="returnAll")
    def return_all(self) -> Optional[bool]:
        return pulumi.get(self, "return_all")

    @property
    @pulumi.getter(name="userPrincipalNames")
    def user_principal_names(self) -> Sequence[str]:
        """
        The user principal names (UPNs) of the users.
        """
        return pulumi.get(self, "user_principal_names")

    @property
    @pulumi.getter
    def users(self) -> Sequence['outputs.GetUsersUserResult']:
        """
        A list of users. Each `user` object provides the attributes documented below.
        """
        return pulumi.get(self, "users")


class AwaitableGetUsersResult(GetUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUsersResult(
            employee_ids=self.employee_ids,
            id=self.id,
            ignore_missing=self.ignore_missing,
            mail_nicknames=self.mail_nicknames,
            mails=self.mails,
            object_ids=self.object_ids,
            return_all=self.return_all,
            user_principal_names=self.user_principal_names,
            users=self.users)


def get_users(employee_ids: Optional[Sequence[str]] = None,
              ignore_missing: Optional[bool] = None,
              mail_nicknames: Optional[Sequence[str]] = None,
              mails: Optional[Sequence[str]] = None,
              object_ids: Optional[Sequence[str]] = None,
              return_all: Optional[bool] = None,
              user_principal_names: Optional[Sequence[str]] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUsersResult:
    """
    Gets basic information for multiple Azure Active Directory users.

    ## API Permissions

    The following API permissions are required in order to use this data source.

    When authenticated with a service principal, this data source requires one of the following application roles: `User.ReadBasic.All`, `User.Read.All` or `Directory.Read.All`

    When authenticated with a user principal, this data source does not require any additional roles.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuread as azuread

    users = azuread.get_users(user_principal_names=[
        "kat@example.com",
        "byte@example.com",
    ])
    ```


    :param Sequence[str] employee_ids: The employee identifiers assigned to the users by the organisation.
    :param bool ignore_missing: Ignore missing users and return users that were found. The data source will still fail if no users are found. Cannot be specified with `return_all`. Defaults to `false`.
    :param Sequence[str] mail_nicknames: The email aliases of the users.
    :param Sequence[str] mails: The SMTP email addresses of the users.
    :param Sequence[str] object_ids: The object IDs of the users.
    :param bool return_all: When `true`, the data source will return all users. Cannot be used with `ignore_missing`. Defaults to `false`.
    :param Sequence[str] user_principal_names: The user principal names (UPNs) of the users.
           
           > Either `return_all`, or one of `user_principal_names`, `object_ids`, `mail_nicknames`, `mails`, or `employee_ids` must be specified. These _may_ be specified as an empty list, in which case no results will be returned.
    """
    __args__ = dict()
    __args__['employeeIds'] = employee_ids
    __args__['ignoreMissing'] = ignore_missing
    __args__['mailNicknames'] = mail_nicknames
    __args__['mails'] = mails
    __args__['objectIds'] = object_ids
    __args__['returnAll'] = return_all
    __args__['userPrincipalNames'] = user_principal_names
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuread:index/getUsers:getUsers', __args__, opts=opts, typ=GetUsersResult).value

    return AwaitableGetUsersResult(
        employee_ids=pulumi.get(__ret__, 'employee_ids'),
        id=pulumi.get(__ret__, 'id'),
        ignore_missing=pulumi.get(__ret__, 'ignore_missing'),
        mail_nicknames=pulumi.get(__ret__, 'mail_nicknames'),
        mails=pulumi.get(__ret__, 'mails'),
        object_ids=pulumi.get(__ret__, 'object_ids'),
        return_all=pulumi.get(__ret__, 'return_all'),
        user_principal_names=pulumi.get(__ret__, 'user_principal_names'),
        users=pulumi.get(__ret__, 'users'))


@_utilities.lift_output_func(get_users)
def get_users_output(employee_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     ignore_missing: Optional[pulumi.Input[Optional[bool]]] = None,
                     mail_nicknames: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     mails: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     object_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     return_all: Optional[pulumi.Input[Optional[bool]]] = None,
                     user_principal_names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUsersResult]:
    """
    Gets basic information for multiple Azure Active Directory users.

    ## API Permissions

    The following API permissions are required in order to use this data source.

    When authenticated with a service principal, this data source requires one of the following application roles: `User.ReadBasic.All`, `User.Read.All` or `Directory.Read.All`

    When authenticated with a user principal, this data source does not require any additional roles.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuread as azuread

    users = azuread.get_users(user_principal_names=[
        "kat@example.com",
        "byte@example.com",
    ])
    ```


    :param Sequence[str] employee_ids: The employee identifiers assigned to the users by the organisation.
    :param bool ignore_missing: Ignore missing users and return users that were found. The data source will still fail if no users are found. Cannot be specified with `return_all`. Defaults to `false`.
    :param Sequence[str] mail_nicknames: The email aliases of the users.
    :param Sequence[str] mails: The SMTP email addresses of the users.
    :param Sequence[str] object_ids: The object IDs of the users.
    :param bool return_all: When `true`, the data source will return all users. Cannot be used with `ignore_missing`. Defaults to `false`.
    :param Sequence[str] user_principal_names: The user principal names (UPNs) of the users.
           
           > Either `return_all`, or one of `user_principal_names`, `object_ids`, `mail_nicknames`, `mails`, or `employee_ids` must be specified. These _may_ be specified as an empty list, in which case no results will be returned.
    """
    ...

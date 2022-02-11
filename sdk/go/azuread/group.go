// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a group within Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires one of the following application roles: `Group.ReadWrite.All` or `Directory.ReadWrite.All`
//
// If using the `assignableToRole` property, this resource additionally requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`
//
// If specifying owners for a group, which are user principals, this resource additionally requires one of the following application roles: `User.Read.All`, `User.ReadWrite.All`, `Directory.Read.All` or `Directory.ReadWrite.All`
//
// When authenticated with a user principal, this resource requires one of the following directory roles: `Groups Administrator`, `User Administrator` or `Global Administrator`
//
// The `externalSendersAllowed`, `autoSubscribeNewMembers`, `hideFromAddressLists` and `hideFromOutlookClients` properties can only be configured when authenticating as a user and cannot be configured when authenticating as a service principal. Additionally, the user being used for authentication must be a Member of the tenant where the group is being managed and _not_ a Guest. This is a known API issue; please see the [Microsoft Graph Known Issues](https://docs.microsoft.com/en-us/graph/known-issues#groups) official documentation.
//
// ## Import
//
// Groups can be imported using their object ID, e.g.
//
// ```sh
//  $ pulumi import azuread:index/group:Group my_group 00000000-0000-0000-0000-000000000000
// ```
type Group struct {
	pulumi.CustomResourceState

	// Indicates whether this group can be assigned to an Azure Active Directory role. Can only be `true` for security-enabled groups. Changing this forces a new resource to be created.
	AssignableToRole pulumi.BoolPtrOutput `pulumi:"assignableToRole"`
	// Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Can only be set for Unified groups.
	AutoSubscribeNewMembers pulumi.BoolOutput `pulumi:"autoSubscribeNewMembers"`
	// A set of behaviors for a Microsoft 365 group. Possible values are `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details. Changing this forces a new resource to be created.
	Behaviors pulumi.StringArrayOutput `pulumi:"behaviors"`
	// The description for the group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name for the group.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// A `dynamicMembership` block as documented below. Required when `types` contains `DynamicMembership`. Cannot be used with the `members` property.
	DynamicMembership GroupDynamicMembershipPtrOutput `pulumi:"dynamicMembership"`
	// Indicates whether people external to the organization can send messages to the group. Can only be set for Unified groups.
	ExternalSendersAllowed pulumi.BoolOutput `pulumi:"externalSendersAllowed"`
	// Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Can only be set for Unified groups.
	HideFromAddressLists pulumi.BoolOutput `pulumi:"hideFromAddressLists"`
	// Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Can only be set for Unified groups.
	HideFromOutlookClients pulumi.BoolOutput `pulumi:"hideFromOutlookClients"`
	// The SMTP address for the group.
	Mail pulumi.StringOutput `pulumi:"mail"`
	// Whether the group is a mail enabled, with a shared group mailbox. At least one of `mailEnabled` or `securityEnabled` must be specified. Only Microsoft 365 groups can be mail enabled (see the `types` property).
	MailEnabled pulumi.BoolPtrOutput `pulumi:"mailEnabled"`
	// The mail alias for the group, unique in the organisation. Required for mail-enabled groups. Changing this forces a new resource to be created.
	MailNickname pulumi.StringOutput `pulumi:"mailNickname"`
	// A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals. Cannot be used with the `dynamicMembership` block.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The object ID of the group.
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesDomainName pulumi.StringOutput `pulumi:"onpremisesDomainName"`
	// The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesNetbiosName pulumi.StringOutput `pulumi:"onpremisesNetbiosName"`
	// The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesSamAccountName pulumi.StringOutput `pulumi:"onpremisesSamAccountName"`
	// The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesSecurityIdentifier pulumi.StringOutput `pulumi:"onpremisesSecurityIdentifier"`
	// Whether this group is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
	OnpremisesSyncEnabled pulumi.BoolOutput `pulumi:"onpremisesSyncEnabled"`
	// A set of owners who own this group. Supported object types are Users or Service Principals
	Owners pulumi.StringArrayOutput `pulumi:"owners"`
	// The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
	PreferredLanguage pulumi.StringOutput `pulumi:"preferredLanguage"`
	// If `true`, will return an error if an existing group is found with the same name. Defaults to `false`.
	PreventDuplicateNames pulumi.BoolPtrOutput `pulumi:"preventDuplicateNames"`
	// A set of provisioning options for a Microsoft 365 group. The only supported value is `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details. Changing this forces a new resource to be created.
	ProvisioningOptions pulumi.StringArrayOutput `pulumi:"provisioningOptions"`
	// List of email addresses for the group that direct to the same group mailbox.
	ProxyAddresses pulumi.StringArrayOutput `pulumi:"proxyAddresses"`
	// Whether the group is a security group for controlling access to in-app resources. At least one of `securityEnabled` or `mailEnabled` must be specified. A Microsoft 365 group can be security enabled _and_ mail enabled (see the `types` property).
	SecurityEnabled pulumi.BoolPtrOutput `pulumi:"securityEnabled"`
	// The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. By default, no theme is set.
	Theme pulumi.StringPtrOutput `pulumi:"theme"`
	// A set of group types to configure for the group. Supported values are `DynamicMembership`, which denotes a group with dynamic membership, and `Unified`, which specifies a Microsoft 365 group. Required when `mailEnabled` is true. Changing this forces a new resource to be created.
	Types pulumi.StringArrayOutput `pulumi:"types"`
	// The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility and this value must be set when the group is created. By default, security groups will receive `Private` visibility and Microsoft 365 groups will receive `Public` visibility.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource Group
	err := ctx.RegisterResource("azuread:index/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("azuread:index/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// Indicates whether this group can be assigned to an Azure Active Directory role. Can only be `true` for security-enabled groups. Changing this forces a new resource to be created.
	AssignableToRole *bool `pulumi:"assignableToRole"`
	// Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Can only be set for Unified groups.
	AutoSubscribeNewMembers *bool `pulumi:"autoSubscribeNewMembers"`
	// A set of behaviors for a Microsoft 365 group. Possible values are `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details. Changing this forces a new resource to be created.
	Behaviors []string `pulumi:"behaviors"`
	// The description for the group.
	Description *string `pulumi:"description"`
	// The display name for the group.
	DisplayName *string `pulumi:"displayName"`
	// A `dynamicMembership` block as documented below. Required when `types` contains `DynamicMembership`. Cannot be used with the `members` property.
	DynamicMembership *GroupDynamicMembership `pulumi:"dynamicMembership"`
	// Indicates whether people external to the organization can send messages to the group. Can only be set for Unified groups.
	ExternalSendersAllowed *bool `pulumi:"externalSendersAllowed"`
	// Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Can only be set for Unified groups.
	HideFromAddressLists *bool `pulumi:"hideFromAddressLists"`
	// Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Can only be set for Unified groups.
	HideFromOutlookClients *bool `pulumi:"hideFromOutlookClients"`
	// The SMTP address for the group.
	Mail *string `pulumi:"mail"`
	// Whether the group is a mail enabled, with a shared group mailbox. At least one of `mailEnabled` or `securityEnabled` must be specified. Only Microsoft 365 groups can be mail enabled (see the `types` property).
	MailEnabled *bool `pulumi:"mailEnabled"`
	// The mail alias for the group, unique in the organisation. Required for mail-enabled groups. Changing this forces a new resource to be created.
	MailNickname *string `pulumi:"mailNickname"`
	// A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals. Cannot be used with the `dynamicMembership` block.
	Members []string `pulumi:"members"`
	// The object ID of the group.
	ObjectId *string `pulumi:"objectId"`
	// The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesDomainName *string `pulumi:"onpremisesDomainName"`
	// The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesNetbiosName *string `pulumi:"onpremisesNetbiosName"`
	// The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesSamAccountName *string `pulumi:"onpremisesSamAccountName"`
	// The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesSecurityIdentifier *string `pulumi:"onpremisesSecurityIdentifier"`
	// Whether this group is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
	OnpremisesSyncEnabled *bool `pulumi:"onpremisesSyncEnabled"`
	// A set of owners who own this group. Supported object types are Users or Service Principals
	Owners []string `pulumi:"owners"`
	// The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
	PreferredLanguage *string `pulumi:"preferredLanguage"`
	// If `true`, will return an error if an existing group is found with the same name. Defaults to `false`.
	PreventDuplicateNames *bool `pulumi:"preventDuplicateNames"`
	// A set of provisioning options for a Microsoft 365 group. The only supported value is `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details. Changing this forces a new resource to be created.
	ProvisioningOptions []string `pulumi:"provisioningOptions"`
	// List of email addresses for the group that direct to the same group mailbox.
	ProxyAddresses []string `pulumi:"proxyAddresses"`
	// Whether the group is a security group for controlling access to in-app resources. At least one of `securityEnabled` or `mailEnabled` must be specified. A Microsoft 365 group can be security enabled _and_ mail enabled (see the `types` property).
	SecurityEnabled *bool `pulumi:"securityEnabled"`
	// The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. By default, no theme is set.
	Theme *string `pulumi:"theme"`
	// A set of group types to configure for the group. Supported values are `DynamicMembership`, which denotes a group with dynamic membership, and `Unified`, which specifies a Microsoft 365 group. Required when `mailEnabled` is true. Changing this forces a new resource to be created.
	Types []string `pulumi:"types"`
	// The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility and this value must be set when the group is created. By default, security groups will receive `Private` visibility and Microsoft 365 groups will receive `Public` visibility.
	Visibility *string `pulumi:"visibility"`
}

type GroupState struct {
	// Indicates whether this group can be assigned to an Azure Active Directory role. Can only be `true` for security-enabled groups. Changing this forces a new resource to be created.
	AssignableToRole pulumi.BoolPtrInput
	// Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Can only be set for Unified groups.
	AutoSubscribeNewMembers pulumi.BoolPtrInput
	// A set of behaviors for a Microsoft 365 group. Possible values are `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details. Changing this forces a new resource to be created.
	Behaviors pulumi.StringArrayInput
	// The description for the group.
	Description pulumi.StringPtrInput
	// The display name for the group.
	DisplayName pulumi.StringPtrInput
	// A `dynamicMembership` block as documented below. Required when `types` contains `DynamicMembership`. Cannot be used with the `members` property.
	DynamicMembership GroupDynamicMembershipPtrInput
	// Indicates whether people external to the organization can send messages to the group. Can only be set for Unified groups.
	ExternalSendersAllowed pulumi.BoolPtrInput
	// Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Can only be set for Unified groups.
	HideFromAddressLists pulumi.BoolPtrInput
	// Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Can only be set for Unified groups.
	HideFromOutlookClients pulumi.BoolPtrInput
	// The SMTP address for the group.
	Mail pulumi.StringPtrInput
	// Whether the group is a mail enabled, with a shared group mailbox. At least one of `mailEnabled` or `securityEnabled` must be specified. Only Microsoft 365 groups can be mail enabled (see the `types` property).
	MailEnabled pulumi.BoolPtrInput
	// The mail alias for the group, unique in the organisation. Required for mail-enabled groups. Changing this forces a new resource to be created.
	MailNickname pulumi.StringPtrInput
	// A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals. Cannot be used with the `dynamicMembership` block.
	Members pulumi.StringArrayInput
	// The object ID of the group.
	ObjectId pulumi.StringPtrInput
	// The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesDomainName pulumi.StringPtrInput
	// The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesNetbiosName pulumi.StringPtrInput
	// The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesSamAccountName pulumi.StringPtrInput
	// The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesSecurityIdentifier pulumi.StringPtrInput
	// Whether this group is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
	OnpremisesSyncEnabled pulumi.BoolPtrInput
	// A set of owners who own this group. Supported object types are Users or Service Principals
	Owners pulumi.StringArrayInput
	// The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
	PreferredLanguage pulumi.StringPtrInput
	// If `true`, will return an error if an existing group is found with the same name. Defaults to `false`.
	PreventDuplicateNames pulumi.BoolPtrInput
	// A set of provisioning options for a Microsoft 365 group. The only supported value is `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details. Changing this forces a new resource to be created.
	ProvisioningOptions pulumi.StringArrayInput
	// List of email addresses for the group that direct to the same group mailbox.
	ProxyAddresses pulumi.StringArrayInput
	// Whether the group is a security group for controlling access to in-app resources. At least one of `securityEnabled` or `mailEnabled` must be specified. A Microsoft 365 group can be security enabled _and_ mail enabled (see the `types` property).
	SecurityEnabled pulumi.BoolPtrInput
	// The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. By default, no theme is set.
	Theme pulumi.StringPtrInput
	// A set of group types to configure for the group. Supported values are `DynamicMembership`, which denotes a group with dynamic membership, and `Unified`, which specifies a Microsoft 365 group. Required when `mailEnabled` is true. Changing this forces a new resource to be created.
	Types pulumi.StringArrayInput
	// The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility and this value must be set when the group is created. By default, security groups will receive `Private` visibility and Microsoft 365 groups will receive `Public` visibility.
	Visibility pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// Indicates whether this group can be assigned to an Azure Active Directory role. Can only be `true` for security-enabled groups. Changing this forces a new resource to be created.
	AssignableToRole *bool `pulumi:"assignableToRole"`
	// Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Can only be set for Unified groups.
	AutoSubscribeNewMembers *bool `pulumi:"autoSubscribeNewMembers"`
	// A set of behaviors for a Microsoft 365 group. Possible values are `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details. Changing this forces a new resource to be created.
	Behaviors []string `pulumi:"behaviors"`
	// The description for the group.
	Description *string `pulumi:"description"`
	// The display name for the group.
	DisplayName string `pulumi:"displayName"`
	// A `dynamicMembership` block as documented below. Required when `types` contains `DynamicMembership`. Cannot be used with the `members` property.
	DynamicMembership *GroupDynamicMembership `pulumi:"dynamicMembership"`
	// Indicates whether people external to the organization can send messages to the group. Can only be set for Unified groups.
	ExternalSendersAllowed *bool `pulumi:"externalSendersAllowed"`
	// Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Can only be set for Unified groups.
	HideFromAddressLists *bool `pulumi:"hideFromAddressLists"`
	// Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Can only be set for Unified groups.
	HideFromOutlookClients *bool `pulumi:"hideFromOutlookClients"`
	// Whether the group is a mail enabled, with a shared group mailbox. At least one of `mailEnabled` or `securityEnabled` must be specified. Only Microsoft 365 groups can be mail enabled (see the `types` property).
	MailEnabled *bool `pulumi:"mailEnabled"`
	// The mail alias for the group, unique in the organisation. Required for mail-enabled groups. Changing this forces a new resource to be created.
	MailNickname *string `pulumi:"mailNickname"`
	// A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals. Cannot be used with the `dynamicMembership` block.
	Members []string `pulumi:"members"`
	// A set of owners who own this group. Supported object types are Users or Service Principals
	Owners []string `pulumi:"owners"`
	// If `true`, will return an error if an existing group is found with the same name. Defaults to `false`.
	PreventDuplicateNames *bool `pulumi:"preventDuplicateNames"`
	// A set of provisioning options for a Microsoft 365 group. The only supported value is `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details. Changing this forces a new resource to be created.
	ProvisioningOptions []string `pulumi:"provisioningOptions"`
	// Whether the group is a security group for controlling access to in-app resources. At least one of `securityEnabled` or `mailEnabled` must be specified. A Microsoft 365 group can be security enabled _and_ mail enabled (see the `types` property).
	SecurityEnabled *bool `pulumi:"securityEnabled"`
	// The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. By default, no theme is set.
	Theme *string `pulumi:"theme"`
	// A set of group types to configure for the group. Supported values are `DynamicMembership`, which denotes a group with dynamic membership, and `Unified`, which specifies a Microsoft 365 group. Required when `mailEnabled` is true. Changing this forces a new resource to be created.
	Types []string `pulumi:"types"`
	// The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility and this value must be set when the group is created. By default, security groups will receive `Private` visibility and Microsoft 365 groups will receive `Public` visibility.
	Visibility *string `pulumi:"visibility"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// Indicates whether this group can be assigned to an Azure Active Directory role. Can only be `true` for security-enabled groups. Changing this forces a new resource to be created.
	AssignableToRole pulumi.BoolPtrInput
	// Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Can only be set for Unified groups.
	AutoSubscribeNewMembers pulumi.BoolPtrInput
	// A set of behaviors for a Microsoft 365 group. Possible values are `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details. Changing this forces a new resource to be created.
	Behaviors pulumi.StringArrayInput
	// The description for the group.
	Description pulumi.StringPtrInput
	// The display name for the group.
	DisplayName pulumi.StringInput
	// A `dynamicMembership` block as documented below. Required when `types` contains `DynamicMembership`. Cannot be used with the `members` property.
	DynamicMembership GroupDynamicMembershipPtrInput
	// Indicates whether people external to the organization can send messages to the group. Can only be set for Unified groups.
	ExternalSendersAllowed pulumi.BoolPtrInput
	// Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Can only be set for Unified groups.
	HideFromAddressLists pulumi.BoolPtrInput
	// Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Can only be set for Unified groups.
	HideFromOutlookClients pulumi.BoolPtrInput
	// Whether the group is a mail enabled, with a shared group mailbox. At least one of `mailEnabled` or `securityEnabled` must be specified. Only Microsoft 365 groups can be mail enabled (see the `types` property).
	MailEnabled pulumi.BoolPtrInput
	// The mail alias for the group, unique in the organisation. Required for mail-enabled groups. Changing this forces a new resource to be created.
	MailNickname pulumi.StringPtrInput
	// A set of members who should be present in this group. Supported object types are Users, Groups or Service Principals. Cannot be used with the `dynamicMembership` block.
	Members pulumi.StringArrayInput
	// A set of owners who own this group. Supported object types are Users or Service Principals
	Owners pulumi.StringArrayInput
	// If `true`, will return an error if an existing group is found with the same name. Defaults to `false`.
	PreventDuplicateNames pulumi.BoolPtrInput
	// A set of provisioning options for a Microsoft 365 group. The only supported value is `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details. Changing this forces a new resource to be created.
	ProvisioningOptions pulumi.StringArrayInput
	// Whether the group is a security group for controlling access to in-app resources. At least one of `securityEnabled` or `mailEnabled` must be specified. A Microsoft 365 group can be security enabled _and_ mail enabled (see the `types` property).
	SecurityEnabled pulumi.BoolPtrInput
	// The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. By default, no theme is set.
	Theme pulumi.StringPtrInput
	// A set of group types to configure for the group. Supported values are `DynamicMembership`, which denotes a group with dynamic membership, and `Unified`, which specifies a Microsoft 365 group. Required when `mailEnabled` is true. Changing this forces a new resource to be created.
	Types pulumi.StringArrayInput
	// The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility and this value must be set when the group is created. By default, security groups will receive `Private` visibility and Microsoft 365 groups will receive `Public` visibility.
	Visibility pulumi.StringPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

// GroupArrayInput is an input type that accepts GroupArray and GroupArrayOutput values.
// You can construct a concrete instance of `GroupArrayInput` via:
//
//          GroupArray{ GroupArgs{...} }
type GroupArrayInput interface {
	pulumi.Input

	ToGroupArrayOutput() GroupArrayOutput
	ToGroupArrayOutputWithContext(context.Context) GroupArrayOutput
}

type GroupArray []GroupInput

func (GroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (i GroupArray) ToGroupArrayOutput() GroupArrayOutput {
	return i.ToGroupArrayOutputWithContext(context.Background())
}

func (i GroupArray) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupArrayOutput)
}

// GroupMapInput is an input type that accepts GroupMap and GroupMapOutput values.
// You can construct a concrete instance of `GroupMapInput` via:
//
//          GroupMap{ "key": GroupArgs{...} }
type GroupMapInput interface {
	pulumi.Input

	ToGroupMapOutput() GroupMapOutput
	ToGroupMapOutputWithContext(context.Context) GroupMapOutput
}

type GroupMap map[string]GroupInput

func (GroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (i GroupMap) ToGroupMapOutput() GroupMapOutput {
	return i.ToGroupMapOutputWithContext(context.Background())
}

func (i GroupMap) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMapOutput)
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

type GroupArrayOutput struct{ *pulumi.OutputState }

func (GroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (o GroupArrayOutput) ToGroupArrayOutput() GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) Index(i pulumi.IntInput) GroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Group {
		return vs[0].([]*Group)[vs[1].(int)]
	}).(GroupOutput)
}

type GroupMapOutput struct{ *pulumi.OutputState }

func (GroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (o GroupMapOutput) ToGroupMapOutput() GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return o
}

func (o GroupMapOutput) MapIndex(k pulumi.StringInput) GroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Group {
		return vs[0].(map[string]*Group)[vs[1].(string)]
	}).(GroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupArrayInput)(nil)).Elem(), GroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMapInput)(nil)).Elem(), GroupMap{})
	pulumi.RegisterOutputType(GroupOutput{})
	pulumi.RegisterOutputType(GroupArrayOutput{})
	pulumi.RegisterOutputType(GroupMapOutput{})
}

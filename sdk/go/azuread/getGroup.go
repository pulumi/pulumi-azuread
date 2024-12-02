// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about an Azure Active Directory group.
//
// ## API Permissions
//
// The following API permissions are required in order to use this data source.
//
// When authenticated with a service principal, this data source requires one of the following application roles: `Group.Read.All` or `Directory.Read.All`
//
// When authenticated with a user principal, this data source does not require any additional roles.
//
// ## Example Usage
//
// ### By Group Display Name)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := azuread.LookupGroup(ctx, &azuread.LookupGroupArgs{
//				DisplayName:     pulumi.StringRef("MyGroupName"),
//				SecurityEnabled: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGroupResult
	err := ctx.Invoke("azuread:index/getGroup:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroup.
type LookupGroupArgs struct {
	// The display name for the group.
	DisplayName *string `pulumi:"displayName"`
	// Whether to include transitive members (a flat list of all nested members). Defaults to `false`.
	IncludeTransitiveMembers *bool `pulumi:"includeTransitiveMembers"`
	// Whether the group is mail-enabled.
	MailEnabled *bool `pulumi:"mailEnabled"`
	// The mail alias for the group, unique in the organisation.
	MailNickname *string `pulumi:"mailNickname"`
	// Specifies the object ID of the group.
	ObjectId *string `pulumi:"objectId"`
	// Whether the group is a security group.
	//
	// > One of `displayName`, `objectId` or `mailNickname` must be specified.
	SecurityEnabled *bool `pulumi:"securityEnabled"`
}

// A collection of values returned by getGroup.
type LookupGroupResult struct {
	// Indicates whether this group can be assigned to an Azure Active Directory role.
	AssignableToRole bool `pulumi:"assignableToRole"`
	// Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Only set for Unified groups.
	AutoSubscribeNewMembers bool `pulumi:"autoSubscribeNewMembers"`
	// A list of behaviors for a Microsoft 365 group, such as `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details.
	Behaviors []string `pulumi:"behaviors"`
	// The optional description of the group.
	Description string `pulumi:"description"`
	// The display name for the group.
	DisplayName string `pulumi:"displayName"`
	// A `dynamicMembership` block as documented below.
	DynamicMemberships []GetGroupDynamicMembership `pulumi:"dynamicMemberships"`
	// Indicates whether people external to the organization can send messages to the group. Only set for Unified groups.
	ExternalSendersAllowed bool `pulumi:"externalSendersAllowed"`
	// Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Only set for Unified groups.
	HideFromAddressLists bool `pulumi:"hideFromAddressLists"`
	// Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Only set for Unified groups.
	HideFromOutlookClients bool `pulumi:"hideFromOutlookClients"`
	// The provider-assigned unique ID for this managed resource.
	Id                       string `pulumi:"id"`
	IncludeTransitiveMembers *bool  `pulumi:"includeTransitiveMembers"`
	// The SMTP address for the group.
	Mail string `pulumi:"mail"`
	// Whether the group is mail-enabled.
	MailEnabled bool `pulumi:"mailEnabled"`
	// The mail alias for the group, unique in the organisation.
	MailNickname string `pulumi:"mailNickname"`
	// List of object IDs of the group members. When `includeTransitiveMembers` is `true`, contains a list of object IDs of all transitive group members.
	Members []string `pulumi:"members"`
	// The object ID of the group.
	ObjectId string `pulumi:"objectId"`
	// The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesDomainName string `pulumi:"onpremisesDomainName"`
	// The on-premises group type that the AAD group will be written as, when writeback is enabled. Possible values are `UniversalDistributionGroup`, `UniversalMailEnabledSecurityGroup`, or `UniversalSecurityGroup`.
	OnpremisesGroupType string `pulumi:"onpremisesGroupType"`
	// The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesNetbiosName string `pulumi:"onpremisesNetbiosName"`
	// The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesSamAccountName string `pulumi:"onpremisesSamAccountName"`
	// The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesSecurityIdentifier string `pulumi:"onpremisesSecurityIdentifier"`
	// Whether this group is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
	OnpremisesSyncEnabled bool `pulumi:"onpremisesSyncEnabled"`
	// List of object IDs of the group owners.
	Owners []string `pulumi:"owners"`
	// The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
	PreferredLanguage string `pulumi:"preferredLanguage"`
	// A list of provisioning options for a Microsoft 365 group, such as `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details.
	ProvisioningOptions []string `pulumi:"provisioningOptions"`
	// List of email addresses for the group that direct to the same group mailbox.
	ProxyAddresses []string `pulumi:"proxyAddresses"`
	// Whether the group is a security group.
	SecurityEnabled bool `pulumi:"securityEnabled"`
	// The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. When no theme is set, the value is `null`.
	Theme string `pulumi:"theme"`
	// A list of group types configured for the group. Supported values are `DynamicMembership`, which denotes a group with dynamic membership, and `Unified`, which specifies a Microsoft 365 group.
	Types []string `pulumi:"types"`
	// The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility.
	Visibility string `pulumi:"visibility"`
	// Whether the group will be written back to the configured on-premises Active Directory when Azure AD Connect is used.
	WritebackEnabled bool `pulumi:"writebackEnabled"`
}

func LookupGroupOutput(ctx *pulumi.Context, args LookupGroupOutputArgs, opts ...pulumi.InvokeOption) LookupGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupGroupResultOutput, error) {
			args := v.(LookupGroupArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupGroupResult
			secret, deps, err := ctx.InvokePackageRawWithDeps("azuread:index/getGroup:getGroup", args, &rv, "", opts...)
			if err != nil {
				return LookupGroupResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupGroupResultOutput)
			output = pulumi.OutputWithDependencies(ctx.Context(), output, deps...).(LookupGroupResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupGroupResultOutput), nil
			}
			return output, nil
		}).(LookupGroupResultOutput)
}

// A collection of arguments for invoking getGroup.
type LookupGroupOutputArgs struct {
	// The display name for the group.
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	// Whether to include transitive members (a flat list of all nested members). Defaults to `false`.
	IncludeTransitiveMembers pulumi.BoolPtrInput `pulumi:"includeTransitiveMembers"`
	// Whether the group is mail-enabled.
	MailEnabled pulumi.BoolPtrInput `pulumi:"mailEnabled"`
	// The mail alias for the group, unique in the organisation.
	MailNickname pulumi.StringPtrInput `pulumi:"mailNickname"`
	// Specifies the object ID of the group.
	ObjectId pulumi.StringPtrInput `pulumi:"objectId"`
	// Whether the group is a security group.
	//
	// > One of `displayName`, `objectId` or `mailNickname` must be specified.
	SecurityEnabled pulumi.BoolPtrInput `pulumi:"securityEnabled"`
}

func (LookupGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupArgs)(nil)).Elem()
}

// A collection of values returned by getGroup.
type LookupGroupResultOutput struct{ *pulumi.OutputState }

func (LookupGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupResult)(nil)).Elem()
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutput() LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutputWithContext(ctx context.Context) LookupGroupResultOutput {
	return o
}

// Indicates whether this group can be assigned to an Azure Active Directory role.
func (o LookupGroupResultOutput) AssignableToRole() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.AssignableToRole }).(pulumi.BoolOutput)
}

// Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Only set for Unified groups.
func (o LookupGroupResultOutput) AutoSubscribeNewMembers() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.AutoSubscribeNewMembers }).(pulumi.BoolOutput)
}

// A list of behaviors for a Microsoft 365 group, such as `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details.
func (o LookupGroupResultOutput) Behaviors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.Behaviors }).(pulumi.StringArrayOutput)
}

// The optional description of the group.
func (o LookupGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// The display name for the group.
func (o LookupGroupResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// A `dynamicMembership` block as documented below.
func (o LookupGroupResultOutput) DynamicMemberships() GetGroupDynamicMembershipArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []GetGroupDynamicMembership { return v.DynamicMemberships }).(GetGroupDynamicMembershipArrayOutput)
}

// Indicates whether people external to the organization can send messages to the group. Only set for Unified groups.
func (o LookupGroupResultOutput) ExternalSendersAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.ExternalSendersAllowed }).(pulumi.BoolOutput)
}

// Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Only set for Unified groups.
func (o LookupGroupResultOutput) HideFromAddressLists() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.HideFromAddressLists }).(pulumi.BoolOutput)
}

// Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Only set for Unified groups.
func (o LookupGroupResultOutput) HideFromOutlookClients() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.HideFromOutlookClients }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) IncludeTransitiveMembers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *bool { return v.IncludeTransitiveMembers }).(pulumi.BoolPtrOutput)
}

// The SMTP address for the group.
func (o LookupGroupResultOutput) Mail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Mail }).(pulumi.StringOutput)
}

// Whether the group is mail-enabled.
func (o LookupGroupResultOutput) MailEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.MailEnabled }).(pulumi.BoolOutput)
}

// The mail alias for the group, unique in the organisation.
func (o LookupGroupResultOutput) MailNickname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.MailNickname }).(pulumi.StringOutput)
}

// List of object IDs of the group members. When `includeTransitiveMembers` is `true`, contains a list of object IDs of all transitive group members.
func (o LookupGroupResultOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.Members }).(pulumi.StringArrayOutput)
}

// The object ID of the group.
func (o LookupGroupResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

// The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
func (o LookupGroupResultOutput) OnpremisesDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.OnpremisesDomainName }).(pulumi.StringOutput)
}

// The on-premises group type that the AAD group will be written as, when writeback is enabled. Possible values are `UniversalDistributionGroup`, `UniversalMailEnabledSecurityGroup`, or `UniversalSecurityGroup`.
func (o LookupGroupResultOutput) OnpremisesGroupType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.OnpremisesGroupType }).(pulumi.StringOutput)
}

// The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.
func (o LookupGroupResultOutput) OnpremisesNetbiosName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.OnpremisesNetbiosName }).(pulumi.StringOutput)
}

// The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.
func (o LookupGroupResultOutput) OnpremisesSamAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.OnpremisesSamAccountName }).(pulumi.StringOutput)
}

// The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
func (o LookupGroupResultOutput) OnpremisesSecurityIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.OnpremisesSecurityIdentifier }).(pulumi.StringOutput)
}

// Whether this group is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
func (o LookupGroupResultOutput) OnpremisesSyncEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.OnpremisesSyncEnabled }).(pulumi.BoolOutput)
}

// List of object IDs of the group owners.
func (o LookupGroupResultOutput) Owners() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.Owners }).(pulumi.StringArrayOutput)
}

// The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
func (o LookupGroupResultOutput) PreferredLanguage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.PreferredLanguage }).(pulumi.StringOutput)
}

// A list of provisioning options for a Microsoft 365 group, such as `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details.
func (o LookupGroupResultOutput) ProvisioningOptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.ProvisioningOptions }).(pulumi.StringArrayOutput)
}

// List of email addresses for the group that direct to the same group mailbox.
func (o LookupGroupResultOutput) ProxyAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.ProxyAddresses }).(pulumi.StringArrayOutput)
}

// Whether the group is a security group.
func (o LookupGroupResultOutput) SecurityEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.SecurityEnabled }).(pulumi.BoolOutput)
}

// The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. When no theme is set, the value is `null`.
func (o LookupGroupResultOutput) Theme() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Theme }).(pulumi.StringOutput)
}

// A list of group types configured for the group. Supported values are `DynamicMembership`, which denotes a group with dynamic membership, and `Unified`, which specifies a Microsoft 365 group.
func (o LookupGroupResultOutput) Types() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.Types }).(pulumi.StringArrayOutput)
}

// The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility.
func (o LookupGroupResultOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Visibility }).(pulumi.StringOutput)
}

// Whether the group will be written back to the configured on-premises Active Directory when Azure AD Connect is used.
func (o LookupGroupResultOutput) WritebackEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.WritebackEnabled }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupResultOutput{})
}

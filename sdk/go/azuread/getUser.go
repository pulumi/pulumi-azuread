// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about an Azure Active Directory user.
//
// ## API Permissions
//
// The following API permissions are required in order to use this data source.
//
// When authenticated with a service principal, this data source requires one of the following application roles: `User.Read.All` or `Directory.Read.All`
//
// When authenticated with a user principal, this data source does not require any additional roles.
//
// ## Example Usage
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
//			_, err := azuread.LookupUser(ctx, &azuread.LookupUserArgs{
//				UserPrincipalName: pulumi.StringRef("user@example.com"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserResult
	err := ctx.Invoke("azuread:index/getUser:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUser.
type LookupUserArgs struct {
	// The employee identifier assigned to the user by the organisation.
	EmployeeId *string `pulumi:"employeeId"`
	// The SMTP address for the user.
	Mail *string `pulumi:"mail"`
	// The email alias of the user.
	MailNickname *string `pulumi:"mailNickname"`
	// The object ID of the user.
	ObjectId *string `pulumi:"objectId"`
	// The user principal name (UPN) of the user.
	//
	// > One of `userPrincipalName`, `objectId`, `mail`, `mailNickname` or `employeeId` must be specified.
	UserPrincipalName *string `pulumi:"userPrincipalName"`
}

// A collection of values returned by getUser.
type LookupUserResult struct {
	// Whether or not the account is enabled.
	AccountEnabled bool `pulumi:"accountEnabled"`
	// The age group of the user. Supported values are `Adult`, `NotAdult` and `Minor`.
	AgeGroup string `pulumi:"ageGroup"`
	// A list of telephone numbers for the user.
	BusinessPhones []string `pulumi:"businessPhones"`
	// The city in which the user is located.
	City string `pulumi:"city"`
	// The company name which the user is associated. This property can be useful for describing the company that an external user comes from.
	CompanyName string `pulumi:"companyName"`
	// Whether consent has been obtained for minors. Supported values are `Granted`, `Denied` and `NotRequired`.
	ConsentProvidedForMinor string `pulumi:"consentProvidedForMinor"`
	// The cost center associated with the user.
	CostCenter string `pulumi:"costCenter"`
	// The country/region in which the user is located, e.g. `US` or `UK`.
	Country string `pulumi:"country"`
	// Indicates whether the user account was created as a regular school or work account (`null`), an external account (`Invitation`), a local account for an Azure Active Directory B2C tenant (`LocalAccount`) or self-service sign-up using email verification (`EmailVerified`).
	CreationType string `pulumi:"creationType"`
	// The name for the department in which the user works.
	Department string `pulumi:"department"`
	// The display name of the user.
	DisplayName string `pulumi:"displayName"`
	// The name of the division in which the user works.
	Division string `pulumi:"division"`
	// The hire date of the user, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
	EmployeeHireDate string `pulumi:"employeeHireDate"`
	// The employee identifier assigned to the user by the organisation.
	EmployeeId string `pulumi:"employeeId"`
	// Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.
	EmployeeType string `pulumi:"employeeType"`
	// For an external user invited to the tenant, this property represents the invited user's invitation status. Possible values are `PendingAcceptance` or `Accepted`.
	ExternalUserState string `pulumi:"externalUserState"`
	// The fax number of the user.
	FaxNumber string `pulumi:"faxNumber"`
	// The given name (first name) of the user.
	GivenName string `pulumi:"givenName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.
	ImAddresses []string `pulumi:"imAddresses"`
	// The user’s job title.
	JobTitle string `pulumi:"jobTitle"`
	// The SMTP address for the user.
	Mail string `pulumi:"mail"`
	// The email alias of the user.
	MailNickname string `pulumi:"mailNickname"`
	// The object ID of the user's manager.
	ManagerId string `pulumi:"managerId"`
	// The primary cellular telephone number for the user.
	MobilePhone string `pulumi:"mobilePhone"`
	// The object ID of the user.
	ObjectId string `pulumi:"objectId"`
	// The office location in the user's place of business.
	OfficeLocation string `pulumi:"officeLocation"`
	// The on-premises distinguished name (DN) of the user, synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesDistinguishedName string `pulumi:"onpremisesDistinguishedName"`
	// The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesDomainName string `pulumi:"onpremisesDomainName"`
	// The value used to associate an on-premise Active Directory user account with their Azure AD user object.
	OnpremisesImmutableId string `pulumi:"onpremisesImmutableId"`
	// The on-premise SAM account name of the user.
	OnpremisesSamAccountName string `pulumi:"onpremisesSamAccountName"`
	// The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesSecurityIdentifier string `pulumi:"onpremisesSecurityIdentifier"`
	// Whether this user is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
	OnpremisesSyncEnabled bool `pulumi:"onpremisesSyncEnabled"`
	// The on-premise user principal name of the user.
	OnpremisesUserPrincipalName string `pulumi:"onpremisesUserPrincipalName"`
	// A list of additional email addresses for the user.
	OtherMails []string `pulumi:"otherMails"`
	// The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.
	PostalCode string `pulumi:"postalCode"`
	// The user's preferred language, in ISO 639-1 notation.
	PreferredLanguage string `pulumi:"preferredLanguage"`
	// List of email addresses for the user that direct to the same mailbox.
	ProxyAddresses []string `pulumi:"proxyAddresses"`
	// Whether or not the Outlook global address list should include this user.
	ShowInAddressList bool `pulumi:"showInAddressList"`
	// The state or province in the user's address.
	State string `pulumi:"state"`
	// The street address of the user's place of business.
	StreetAddress string `pulumi:"streetAddress"`
	// The user's surname (family name or last name).
	Surname string `pulumi:"surname"`
	// The usage location of the user.
	UsageLocation string `pulumi:"usageLocation"`
	// The user principal name (UPN) of the user.
	UserPrincipalName string `pulumi:"userPrincipalName"`
	// The user type in the directory. Possible values are `Guest` or `Member`.
	UserType string `pulumi:"userType"`
}

func LookupUserOutput(ctx *pulumi.Context, args LookupUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupUserResultOutput, error) {
			args := v.(LookupUserArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azuread:index/getUser:getUser", args, LookupUserResultOutput{}, options).(LookupUserResultOutput), nil
		}).(LookupUserResultOutput)
}

// A collection of arguments for invoking getUser.
type LookupUserOutputArgs struct {
	// The employee identifier assigned to the user by the organisation.
	EmployeeId pulumi.StringPtrInput `pulumi:"employeeId"`
	// The SMTP address for the user.
	Mail pulumi.StringPtrInput `pulumi:"mail"`
	// The email alias of the user.
	MailNickname pulumi.StringPtrInput `pulumi:"mailNickname"`
	// The object ID of the user.
	ObjectId pulumi.StringPtrInput `pulumi:"objectId"`
	// The user principal name (UPN) of the user.
	//
	// > One of `userPrincipalName`, `objectId`, `mail`, `mailNickname` or `employeeId` must be specified.
	UserPrincipalName pulumi.StringPtrInput `pulumi:"userPrincipalName"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}

// A collection of values returned by getUser.
type LookupUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserResult)(nil)).Elem()
}

func (o LookupUserResultOutput) ToLookupUserResultOutput() LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToLookupUserResultOutputWithContext(ctx context.Context) LookupUserResultOutput {
	return o
}

// Whether or not the account is enabled.
func (o LookupUserResultOutput) AccountEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupUserResult) bool { return v.AccountEnabled }).(pulumi.BoolOutput)
}

// The age group of the user. Supported values are `Adult`, `NotAdult` and `Minor`.
func (o LookupUserResultOutput) AgeGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.AgeGroup }).(pulumi.StringOutput)
}

// A list of telephone numbers for the user.
func (o LookupUserResultOutput) BusinessPhones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []string { return v.BusinessPhones }).(pulumi.StringArrayOutput)
}

// The city in which the user is located.
func (o LookupUserResultOutput) City() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.City }).(pulumi.StringOutput)
}

// The company name which the user is associated. This property can be useful for describing the company that an external user comes from.
func (o LookupUserResultOutput) CompanyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.CompanyName }).(pulumi.StringOutput)
}

// Whether consent has been obtained for minors. Supported values are `Granted`, `Denied` and `NotRequired`.
func (o LookupUserResultOutput) ConsentProvidedForMinor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.ConsentProvidedForMinor }).(pulumi.StringOutput)
}

// The cost center associated with the user.
func (o LookupUserResultOutput) CostCenter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.CostCenter }).(pulumi.StringOutput)
}

// The country/region in which the user is located, e.g. `US` or `UK`.
func (o LookupUserResultOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Country }).(pulumi.StringOutput)
}

// Indicates whether the user account was created as a regular school or work account (`null`), an external account (`Invitation`), a local account for an Azure Active Directory B2C tenant (`LocalAccount`) or self-service sign-up using email verification (`EmailVerified`).
func (o LookupUserResultOutput) CreationType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.CreationType }).(pulumi.StringOutput)
}

// The name for the department in which the user works.
func (o LookupUserResultOutput) Department() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Department }).(pulumi.StringOutput)
}

// The display name of the user.
func (o LookupUserResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The name of the division in which the user works.
func (o LookupUserResultOutput) Division() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Division }).(pulumi.StringOutput)
}

// The hire date of the user, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
func (o LookupUserResultOutput) EmployeeHireDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.EmployeeHireDate }).(pulumi.StringOutput)
}

// The employee identifier assigned to the user by the organisation.
func (o LookupUserResultOutput) EmployeeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.EmployeeId }).(pulumi.StringOutput)
}

// Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.
func (o LookupUserResultOutput) EmployeeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.EmployeeType }).(pulumi.StringOutput)
}

// For an external user invited to the tenant, this property represents the invited user's invitation status. Possible values are `PendingAcceptance` or `Accepted`.
func (o LookupUserResultOutput) ExternalUserState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.ExternalUserState }).(pulumi.StringOutput)
}

// The fax number of the user.
func (o LookupUserResultOutput) FaxNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.FaxNumber }).(pulumi.StringOutput)
}

// The given name (first name) of the user.
func (o LookupUserResultOutput) GivenName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.GivenName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.
func (o LookupUserResultOutput) ImAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []string { return v.ImAddresses }).(pulumi.StringArrayOutput)
}

// The user’s job title.
func (o LookupUserResultOutput) JobTitle() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.JobTitle }).(pulumi.StringOutput)
}

// The SMTP address for the user.
func (o LookupUserResultOutput) Mail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Mail }).(pulumi.StringOutput)
}

// The email alias of the user.
func (o LookupUserResultOutput) MailNickname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.MailNickname }).(pulumi.StringOutput)
}

// The object ID of the user's manager.
func (o LookupUserResultOutput) ManagerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.ManagerId }).(pulumi.StringOutput)
}

// The primary cellular telephone number for the user.
func (o LookupUserResultOutput) MobilePhone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.MobilePhone }).(pulumi.StringOutput)
}

// The object ID of the user.
func (o LookupUserResultOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.ObjectId }).(pulumi.StringOutput)
}

// The office location in the user's place of business.
func (o LookupUserResultOutput) OfficeLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.OfficeLocation }).(pulumi.StringOutput)
}

// The on-premises distinguished name (DN) of the user, synchronised from the on-premises directory when Azure AD Connect is used.
func (o LookupUserResultOutput) OnpremisesDistinguishedName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.OnpremisesDistinguishedName }).(pulumi.StringOutput)
}

// The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
func (o LookupUserResultOutput) OnpremisesDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.OnpremisesDomainName }).(pulumi.StringOutput)
}

// The value used to associate an on-premise Active Directory user account with their Azure AD user object.
func (o LookupUserResultOutput) OnpremisesImmutableId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.OnpremisesImmutableId }).(pulumi.StringOutput)
}

// The on-premise SAM account name of the user.
func (o LookupUserResultOutput) OnpremisesSamAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.OnpremisesSamAccountName }).(pulumi.StringOutput)
}

// The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
func (o LookupUserResultOutput) OnpremisesSecurityIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.OnpremisesSecurityIdentifier }).(pulumi.StringOutput)
}

// Whether this user is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
func (o LookupUserResultOutput) OnpremisesSyncEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupUserResult) bool { return v.OnpremisesSyncEnabled }).(pulumi.BoolOutput)
}

// The on-premise user principal name of the user.
func (o LookupUserResultOutput) OnpremisesUserPrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.OnpremisesUserPrincipalName }).(pulumi.StringOutput)
}

// A list of additional email addresses for the user.
func (o LookupUserResultOutput) OtherMails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []string { return v.OtherMails }).(pulumi.StringArrayOutput)
}

// The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.
func (o LookupUserResultOutput) PostalCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.PostalCode }).(pulumi.StringOutput)
}

// The user's preferred language, in ISO 639-1 notation.
func (o LookupUserResultOutput) PreferredLanguage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.PreferredLanguage }).(pulumi.StringOutput)
}

// List of email addresses for the user that direct to the same mailbox.
func (o LookupUserResultOutput) ProxyAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []string { return v.ProxyAddresses }).(pulumi.StringArrayOutput)
}

// Whether or not the Outlook global address list should include this user.
func (o LookupUserResultOutput) ShowInAddressList() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupUserResult) bool { return v.ShowInAddressList }).(pulumi.BoolOutput)
}

// The state or province in the user's address.
func (o LookupUserResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.State }).(pulumi.StringOutput)
}

// The street address of the user's place of business.
func (o LookupUserResultOutput) StreetAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.StreetAddress }).(pulumi.StringOutput)
}

// The user's surname (family name or last name).
func (o LookupUserResultOutput) Surname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Surname }).(pulumi.StringOutput)
}

// The usage location of the user.
func (o LookupUserResultOutput) UsageLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.UsageLocation }).(pulumi.StringOutput)
}

// The user principal name (UPN) of the user.
func (o LookupUserResultOutput) UserPrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.UserPrincipalName }).(pulumi.StringOutput)
}

// The user type in the directory. Possible values are `Guest` or `Member`.
func (o LookupUserResultOutput) UserType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.UserType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}

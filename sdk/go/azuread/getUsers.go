// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets basic information for multiple Azure Active Directory users.
//
// ## API Permissions
//
// The following API permissions are required in order to use this data source.
//
// When authenticated with a service principal, this data source requires one of the following application roles: `User.ReadBasic.All`, `User.Read.All` or `Directory.Read.All`
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
//			_, err := azuread.GetUsers(ctx, &azuread.GetUsersArgs{
//				UserPrincipalNames: []string{
//					"kat@example.com",
//					"byte@example.com",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetUsers(ctx *pulumi.Context, args *GetUsersArgs, opts ...pulumi.InvokeOption) (*GetUsersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetUsersResult
	err := ctx.Invoke("azuread:index/getUsers:getUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUsers.
type GetUsersArgs struct {
	// The employee identifiers assigned to the users by the organisation.
	EmployeeIds []string `pulumi:"employeeIds"`
	// Ignore missing users and return users that were found. The data source will still fail if no users are found. Cannot be specified with `returnAll`. Defaults to `false`.
	IgnoreMissing *bool `pulumi:"ignoreMissing"`
	// The email aliases of the users.
	MailNicknames []string `pulumi:"mailNicknames"`
	// The SMTP email addresses of the users.
	Mails []string `pulumi:"mails"`
	// The object IDs of the users.
	ObjectIds []string `pulumi:"objectIds"`
	// When `true`, the data source will return all users. Cannot be used with `ignoreMissing`. Defaults to `false`.
	ReturnAll *bool `pulumi:"returnAll"`
	// The user principal names (UPNs) of the users.
	//
	// > Either `returnAll`, or one of `userPrincipalNames`, `objectIds`, `mailNicknames`, `mails`, or `employeeIds` must be specified. These _may_ be specified as an empty list, in which case no results will be returned.
	UserPrincipalNames []string `pulumi:"userPrincipalNames"`
}

// A collection of values returned by getUsers.
type GetUsersResult struct {
	// The employee identifiers assigned to the users by the organisation.
	EmployeeIds []string `pulumi:"employeeIds"`
	// The provider-assigned unique ID for this managed resource.
	Id            string `pulumi:"id"`
	IgnoreMissing *bool  `pulumi:"ignoreMissing"`
	// The email aliases of the users.
	MailNicknames []string `pulumi:"mailNicknames"`
	// The SMTP email addresses of the users.
	Mails []string `pulumi:"mails"`
	// The object IDs of the users.
	ObjectIds []string `pulumi:"objectIds"`
	ReturnAll *bool    `pulumi:"returnAll"`
	// The user principal names (UPNs) of the users.
	UserPrincipalNames []string `pulumi:"userPrincipalNames"`
	// A list of users. Each `user` object provides the attributes documented below.
	Users []GetUsersUser `pulumi:"users"`
}

func GetUsersOutput(ctx *pulumi.Context, args GetUsersOutputArgs, opts ...pulumi.InvokeOption) GetUsersResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetUsersResultOutput, error) {
			args := v.(GetUsersArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azuread:index/getUsers:getUsers", args, GetUsersResultOutput{}, options).(GetUsersResultOutput), nil
		}).(GetUsersResultOutput)
}

// A collection of arguments for invoking getUsers.
type GetUsersOutputArgs struct {
	// The employee identifiers assigned to the users by the organisation.
	EmployeeIds pulumi.StringArrayInput `pulumi:"employeeIds"`
	// Ignore missing users and return users that were found. The data source will still fail if no users are found. Cannot be specified with `returnAll`. Defaults to `false`.
	IgnoreMissing pulumi.BoolPtrInput `pulumi:"ignoreMissing"`
	// The email aliases of the users.
	MailNicknames pulumi.StringArrayInput `pulumi:"mailNicknames"`
	// The SMTP email addresses of the users.
	Mails pulumi.StringArrayInput `pulumi:"mails"`
	// The object IDs of the users.
	ObjectIds pulumi.StringArrayInput `pulumi:"objectIds"`
	// When `true`, the data source will return all users. Cannot be used with `ignoreMissing`. Defaults to `false`.
	ReturnAll pulumi.BoolPtrInput `pulumi:"returnAll"`
	// The user principal names (UPNs) of the users.
	//
	// > Either `returnAll`, or one of `userPrincipalNames`, `objectIds`, `mailNicknames`, `mails`, or `employeeIds` must be specified. These _may_ be specified as an empty list, in which case no results will be returned.
	UserPrincipalNames pulumi.StringArrayInput `pulumi:"userPrincipalNames"`
}

func (GetUsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersArgs)(nil)).Elem()
}

// A collection of values returned by getUsers.
type GetUsersResultOutput struct{ *pulumi.OutputState }

func (GetUsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersResult)(nil)).Elem()
}

func (o GetUsersResultOutput) ToGetUsersResultOutput() GetUsersResultOutput {
	return o
}

func (o GetUsersResultOutput) ToGetUsersResultOutputWithContext(ctx context.Context) GetUsersResultOutput {
	return o
}

// The employee identifiers assigned to the users by the organisation.
func (o GetUsersResultOutput) EmployeeIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []string { return v.EmployeeIds }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetUsersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetUsersResultOutput) IgnoreMissing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *bool { return v.IgnoreMissing }).(pulumi.BoolPtrOutput)
}

// The email aliases of the users.
func (o GetUsersResultOutput) MailNicknames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []string { return v.MailNicknames }).(pulumi.StringArrayOutput)
}

// The SMTP email addresses of the users.
func (o GetUsersResultOutput) Mails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []string { return v.Mails }).(pulumi.StringArrayOutput)
}

// The object IDs of the users.
func (o GetUsersResultOutput) ObjectIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []string { return v.ObjectIds }).(pulumi.StringArrayOutput)
}

func (o GetUsersResultOutput) ReturnAll() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetUsersResult) *bool { return v.ReturnAll }).(pulumi.BoolPtrOutput)
}

// The user principal names (UPNs) of the users.
func (o GetUsersResultOutput) UserPrincipalNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []string { return v.UserPrincipalNames }).(pulumi.StringArrayOutput)
}

// A list of users. Each `user` object provides the attributes documented below.
func (o GetUsersResultOutput) Users() GetUsersUserArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []GetUsersUser { return v.Users }).(GetUsersUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUsersResultOutput{})
}

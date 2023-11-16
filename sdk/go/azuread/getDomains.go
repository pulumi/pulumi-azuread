// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about existing Domains within Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this data source.
//
// When authenticated with a service principal, this data source requires one of the following application roles: `Domain.Read.All` or `Directory.Read.All`
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
//	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			aadDomains, err := azuread.GetDomains(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			var splat0 []*string
//			for _, val0 := range aadDomains.Domains {
//				splat0 = append(splat0, val0.DomainName)
//			}
//			ctx.Export("domainNames", splat0)
//			return nil
//		})
//	}
//
// ```
func GetDomains(ctx *pulumi.Context, args *GetDomainsArgs, opts ...pulumi.InvokeOption) (*GetDomainsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDomainsResult
	err := ctx.Invoke("azuread:index/getDomains:getDomains", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomains.
type GetDomainsArgs struct {
	// Set to `true` to only return domains whose DNS is managed by Microsoft 365. Defaults to `false`.
	AdminManaged *bool `pulumi:"adminManaged"`
	// Set to `true` if unverified Azure AD domains should be included. Defaults to `false`.
	IncludeUnverified *bool `pulumi:"includeUnverified"`
	// Set to `true` to only return the default domain.
	OnlyDefault *bool `pulumi:"onlyDefault"`
	// Set to `true` to only return the initial domain, which is your primary Azure Active Directory tenant domain. Defaults to `false`.
	OnlyInitial *bool `pulumi:"onlyInitial"`
	// Set to `true` to only return verified root domains. Excludes subdomains and unverified domains.
	OnlyRoot *bool `pulumi:"onlyRoot"`
	// A list of supported services that must be supported by a domain. Possible values include `Email`, `Sharepoint`, `EmailInternalRelayOnly`, `OfficeCommunicationsOnline`, `SharePointDefaultDomain`, `FullRedelegation`, `SharePointPublic`, `OrgIdAuthentication`, `Yammer` and `Intune`.
	//
	// > **Note on filters** If `includeUnverified` is set to `true`, you cannot specify `onlyDefault` or `onlyInitial`. Additionally, you cannot combine `onlyDefault` with `onlyInitial`.
	SupportsServices []string `pulumi:"supportsServices"`
}

// A collection of values returned by getDomains.
type GetDomainsResult struct {
	// Whether the DNS for the domain is managed by Microsoft 365.
	AdminManaged *bool `pulumi:"adminManaged"`
	// A list of tenant domains. Each `domain` object provides the attributes documented below.
	Domains []GetDomainsDomain `pulumi:"domains"`
	// The provider-assigned unique ID for this managed resource.
	Id                string   `pulumi:"id"`
	IncludeUnverified *bool    `pulumi:"includeUnverified"`
	OnlyDefault       *bool    `pulumi:"onlyDefault"`
	OnlyInitial       *bool    `pulumi:"onlyInitial"`
	OnlyRoot          *bool    `pulumi:"onlyRoot"`
	SupportsServices  []string `pulumi:"supportsServices"`
}

func GetDomainsOutput(ctx *pulumi.Context, args GetDomainsOutputArgs, opts ...pulumi.InvokeOption) GetDomainsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDomainsResult, error) {
			args := v.(GetDomainsArgs)
			r, err := GetDomains(ctx, &args, opts...)
			var s GetDomainsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDomainsResultOutput)
}

// A collection of arguments for invoking getDomains.
type GetDomainsOutputArgs struct {
	// Set to `true` to only return domains whose DNS is managed by Microsoft 365. Defaults to `false`.
	AdminManaged pulumi.BoolPtrInput `pulumi:"adminManaged"`
	// Set to `true` if unverified Azure AD domains should be included. Defaults to `false`.
	IncludeUnverified pulumi.BoolPtrInput `pulumi:"includeUnverified"`
	// Set to `true` to only return the default domain.
	OnlyDefault pulumi.BoolPtrInput `pulumi:"onlyDefault"`
	// Set to `true` to only return the initial domain, which is your primary Azure Active Directory tenant domain. Defaults to `false`.
	OnlyInitial pulumi.BoolPtrInput `pulumi:"onlyInitial"`
	// Set to `true` to only return verified root domains. Excludes subdomains and unverified domains.
	OnlyRoot pulumi.BoolPtrInput `pulumi:"onlyRoot"`
	// A list of supported services that must be supported by a domain. Possible values include `Email`, `Sharepoint`, `EmailInternalRelayOnly`, `OfficeCommunicationsOnline`, `SharePointDefaultDomain`, `FullRedelegation`, `SharePointPublic`, `OrgIdAuthentication`, `Yammer` and `Intune`.
	//
	// > **Note on filters** If `includeUnverified` is set to `true`, you cannot specify `onlyDefault` or `onlyInitial`. Additionally, you cannot combine `onlyDefault` with `onlyInitial`.
	SupportsServices pulumi.StringArrayInput `pulumi:"supportsServices"`
}

func (GetDomainsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsArgs)(nil)).Elem()
}

// A collection of values returned by getDomains.
type GetDomainsResultOutput struct{ *pulumi.OutputState }

func (GetDomainsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsResult)(nil)).Elem()
}

func (o GetDomainsResultOutput) ToGetDomainsResultOutput() GetDomainsResultOutput {
	return o
}

func (o GetDomainsResultOutput) ToGetDomainsResultOutputWithContext(ctx context.Context) GetDomainsResultOutput {
	return o
}

// Whether the DNS for the domain is managed by Microsoft 365.
func (o GetDomainsResultOutput) AdminManaged() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *bool { return v.AdminManaged }).(pulumi.BoolPtrOutput)
}

// A list of tenant domains. Each `domain` object provides the attributes documented below.
func (o GetDomainsResultOutput) Domains() GetDomainsDomainArrayOutput {
	return o.ApplyT(func(v GetDomainsResult) []GetDomainsDomain { return v.Domains }).(GetDomainsDomainArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDomainsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDomainsResultOutput) IncludeUnverified() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *bool { return v.IncludeUnverified }).(pulumi.BoolPtrOutput)
}

func (o GetDomainsResultOutput) OnlyDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *bool { return v.OnlyDefault }).(pulumi.BoolPtrOutput)
}

func (o GetDomainsResultOutput) OnlyInitial() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *bool { return v.OnlyInitial }).(pulumi.BoolPtrOutput)
}

func (o GetDomainsResultOutput) OnlyRoot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *bool { return v.OnlyRoot }).(pulumi.BoolPtrOutput)
}

func (o GetDomainsResultOutput) SupportsServices() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDomainsResult) []string { return v.SupportsServices }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDomainsResultOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a Named Location within Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this data source.
//
// When authenticated with a service principal, this resource requires the following application roles: `Policy.Read.All`
//
// When authenticated with a user principal, this resource requires one of the following directory roles: `Conditional Access Administrator` or `Global Reader`
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
//			_, err := azuread.LookupNamedLocation(ctx, &azuread.LookupNamedLocationArgs{
//				DisplayName: "My Named Location",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Attributes Reference
//
// The following attributes are exported:
//
// * `country` - A `country` block as documented below, which describes a country-based named location.
// * `id` - The ID of the named location.
// * `ip` - An `ip` block as documented below, which describes an IP-based named location.
// *
// ***
//
// `country` block exports the following:
//
// * `countriesAndRegions` - List of countries and/or regions in two-letter format specified by ISO 3166-2.
// * `includeUnknownCountriesAndRegions` - Whether IP addresses that don't map to a country or region are included in the named location.
//
// ***
//
// `ip` block exports the following:
//
// * `ipRanges` - List of IP address ranges in IPv4 CIDR format (e.g. `1.2.3.4/32`) or any allowable IPv6 format from IETF RFC596.
// * `trusted` - Whether the named location is trusted.
func LookupNamedLocation(ctx *pulumi.Context, args *LookupNamedLocationArgs, opts ...pulumi.InvokeOption) (*LookupNamedLocationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNamedLocationResult
	err := ctx.Invoke("azuread:index/getNamedLocation:getNamedLocation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNamedLocation.
type LookupNamedLocationArgs struct {
	// Specifies the display named of the named location to look up.
	DisplayName string `pulumi:"displayName"`
}

// A collection of values returned by getNamedLocation.
type LookupNamedLocationResult struct {
	Countries   []GetNamedLocationCountry `pulumi:"countries"`
	DisplayName string                    `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id  string               `pulumi:"id"`
	Ips []GetNamedLocationIp `pulumi:"ips"`
}

func LookupNamedLocationOutput(ctx *pulumi.Context, args LookupNamedLocationOutputArgs, opts ...pulumi.InvokeOption) LookupNamedLocationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNamedLocationResultOutput, error) {
			args := v.(LookupNamedLocationArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupNamedLocationResult
			secret, deps, err := ctx.InvokePackageRawWithDeps("azuread:index/getNamedLocation:getNamedLocation", args, &rv, "", opts...)
			if err != nil {
				return LookupNamedLocationResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupNamedLocationResultOutput)
			output = pulumi.OutputWithDependencies(ctx.Context(), output, deps...).(LookupNamedLocationResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupNamedLocationResultOutput), nil
			}
			return output, nil
		}).(LookupNamedLocationResultOutput)
}

// A collection of arguments for invoking getNamedLocation.
type LookupNamedLocationOutputArgs struct {
	// Specifies the display named of the named location to look up.
	DisplayName pulumi.StringInput `pulumi:"displayName"`
}

func (LookupNamedLocationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamedLocationArgs)(nil)).Elem()
}

// A collection of values returned by getNamedLocation.
type LookupNamedLocationResultOutput struct{ *pulumi.OutputState }

func (LookupNamedLocationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamedLocationResult)(nil)).Elem()
}

func (o LookupNamedLocationResultOutput) ToLookupNamedLocationResultOutput() LookupNamedLocationResultOutput {
	return o
}

func (o LookupNamedLocationResultOutput) ToLookupNamedLocationResultOutputWithContext(ctx context.Context) LookupNamedLocationResultOutput {
	return o
}

func (o LookupNamedLocationResultOutput) Countries() GetNamedLocationCountryArrayOutput {
	return o.ApplyT(func(v LookupNamedLocationResult) []GetNamedLocationCountry { return v.Countries }).(GetNamedLocationCountryArrayOutput)
}

func (o LookupNamedLocationResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamedLocationResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNamedLocationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamedLocationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNamedLocationResultOutput) Ips() GetNamedLocationIpArrayOutput {
	return o.ApplyT(func(v LookupNamedLocationResult) []GetNamedLocationIp { return v.Ips }).(GetNamedLocationIpArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamedLocationResultOutput{})
}

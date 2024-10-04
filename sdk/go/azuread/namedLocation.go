// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
//			_, err := azuread.NewNamedLocation(ctx, "example-ip", &azuread.NamedLocationArgs{
//				DisplayName: pulumi.String("IP Named Location"),
//				Ip: &azuread.NamedLocationIpArgs{
//					IpRanges: pulumi.StringArray{
//						pulumi.String("1.1.1.1/32"),
//						pulumi.String("2.2.2.2/32"),
//					},
//					Trusted: pulumi.Bool(true),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewNamedLocation(ctx, "example-country", &azuread.NamedLocationArgs{
//				DisplayName: pulumi.String("Country Named Location"),
//				Country: &azuread.NamedLocationCountryArgs{
//					CountriesAndRegions: pulumi.StringArray{
//						pulumi.String("GB"),
//						pulumi.String("US"),
//					},
//					IncludeUnknownCountriesAndRegions: pulumi.Bool(false),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Named Locations can be imported using the `id`, e.g.
//
// ```sh
// $ pulumi import azuread:index/namedLocation:NamedLocation my_location 00000000-0000-0000-0000-000000000000
// ```
type NamedLocation struct {
	pulumi.CustomResourceState

	// A `country` block as documented below, which configures a country-based named location.
	Country NamedLocationCountryPtrOutput `pulumi:"country"`
	// The friendly name for this named location.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// An `ip` block as documented below, which configures an IP-based named location.
	//
	// > Exactly one of `ip` or `country` must be specified. Changing between these forces a new resource to be created.
	Ip NamedLocationIpPtrOutput `pulumi:"ip"`
}

// NewNamedLocation registers a new resource with the given unique name, arguments, and options.
func NewNamedLocation(ctx *pulumi.Context,
	name string, args *NamedLocationArgs, opts ...pulumi.ResourceOption) (*NamedLocation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NamedLocation
	err := ctx.RegisterResource("azuread:index/namedLocation:NamedLocation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamedLocation gets an existing NamedLocation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamedLocation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamedLocationState, opts ...pulumi.ResourceOption) (*NamedLocation, error) {
	var resource NamedLocation
	err := ctx.ReadResource("azuread:index/namedLocation:NamedLocation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamedLocation resources.
type namedLocationState struct {
	// A `country` block as documented below, which configures a country-based named location.
	Country *NamedLocationCountry `pulumi:"country"`
	// The friendly name for this named location.
	DisplayName *string `pulumi:"displayName"`
	// An `ip` block as documented below, which configures an IP-based named location.
	//
	// > Exactly one of `ip` or `country` must be specified. Changing between these forces a new resource to be created.
	Ip *NamedLocationIp `pulumi:"ip"`
}

type NamedLocationState struct {
	// A `country` block as documented below, which configures a country-based named location.
	Country NamedLocationCountryPtrInput
	// The friendly name for this named location.
	DisplayName pulumi.StringPtrInput
	// An `ip` block as documented below, which configures an IP-based named location.
	//
	// > Exactly one of `ip` or `country` must be specified. Changing between these forces a new resource to be created.
	Ip NamedLocationIpPtrInput
}

func (NamedLocationState) ElementType() reflect.Type {
	return reflect.TypeOf((*namedLocationState)(nil)).Elem()
}

type namedLocationArgs struct {
	// A `country` block as documented below, which configures a country-based named location.
	Country *NamedLocationCountry `pulumi:"country"`
	// The friendly name for this named location.
	DisplayName string `pulumi:"displayName"`
	// An `ip` block as documented below, which configures an IP-based named location.
	//
	// > Exactly one of `ip` or `country` must be specified. Changing between these forces a new resource to be created.
	Ip *NamedLocationIp `pulumi:"ip"`
}

// The set of arguments for constructing a NamedLocation resource.
type NamedLocationArgs struct {
	// A `country` block as documented below, which configures a country-based named location.
	Country NamedLocationCountryPtrInput
	// The friendly name for this named location.
	DisplayName pulumi.StringInput
	// An `ip` block as documented below, which configures an IP-based named location.
	//
	// > Exactly one of `ip` or `country` must be specified. Changing between these forces a new resource to be created.
	Ip NamedLocationIpPtrInput
}

func (NamedLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namedLocationArgs)(nil)).Elem()
}

type NamedLocationInput interface {
	pulumi.Input

	ToNamedLocationOutput() NamedLocationOutput
	ToNamedLocationOutputWithContext(ctx context.Context) NamedLocationOutput
}

func (*NamedLocation) ElementType() reflect.Type {
	return reflect.TypeOf((**NamedLocation)(nil)).Elem()
}

func (i *NamedLocation) ToNamedLocationOutput() NamedLocationOutput {
	return i.ToNamedLocationOutputWithContext(context.Background())
}

func (i *NamedLocation) ToNamedLocationOutputWithContext(ctx context.Context) NamedLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedLocationOutput)
}

// NamedLocationArrayInput is an input type that accepts NamedLocationArray and NamedLocationArrayOutput values.
// You can construct a concrete instance of `NamedLocationArrayInput` via:
//
//	NamedLocationArray{ NamedLocationArgs{...} }
type NamedLocationArrayInput interface {
	pulumi.Input

	ToNamedLocationArrayOutput() NamedLocationArrayOutput
	ToNamedLocationArrayOutputWithContext(context.Context) NamedLocationArrayOutput
}

type NamedLocationArray []NamedLocationInput

func (NamedLocationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NamedLocation)(nil)).Elem()
}

func (i NamedLocationArray) ToNamedLocationArrayOutput() NamedLocationArrayOutput {
	return i.ToNamedLocationArrayOutputWithContext(context.Background())
}

func (i NamedLocationArray) ToNamedLocationArrayOutputWithContext(ctx context.Context) NamedLocationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedLocationArrayOutput)
}

// NamedLocationMapInput is an input type that accepts NamedLocationMap and NamedLocationMapOutput values.
// You can construct a concrete instance of `NamedLocationMapInput` via:
//
//	NamedLocationMap{ "key": NamedLocationArgs{...} }
type NamedLocationMapInput interface {
	pulumi.Input

	ToNamedLocationMapOutput() NamedLocationMapOutput
	ToNamedLocationMapOutputWithContext(context.Context) NamedLocationMapOutput
}

type NamedLocationMap map[string]NamedLocationInput

func (NamedLocationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NamedLocation)(nil)).Elem()
}

func (i NamedLocationMap) ToNamedLocationMapOutput() NamedLocationMapOutput {
	return i.ToNamedLocationMapOutputWithContext(context.Background())
}

func (i NamedLocationMap) ToNamedLocationMapOutputWithContext(ctx context.Context) NamedLocationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedLocationMapOutput)
}

type NamedLocationOutput struct{ *pulumi.OutputState }

func (NamedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamedLocation)(nil)).Elem()
}

func (o NamedLocationOutput) ToNamedLocationOutput() NamedLocationOutput {
	return o
}

func (o NamedLocationOutput) ToNamedLocationOutputWithContext(ctx context.Context) NamedLocationOutput {
	return o
}

// A `country` block as documented below, which configures a country-based named location.
func (o NamedLocationOutput) Country() NamedLocationCountryPtrOutput {
	return o.ApplyT(func(v *NamedLocation) NamedLocationCountryPtrOutput { return v.Country }).(NamedLocationCountryPtrOutput)
}

// The friendly name for this named location.
func (o NamedLocationOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *NamedLocation) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// An `ip` block as documented below, which configures an IP-based named location.
//
// > Exactly one of `ip` or `country` must be specified. Changing between these forces a new resource to be created.
func (o NamedLocationOutput) Ip() NamedLocationIpPtrOutput {
	return o.ApplyT(func(v *NamedLocation) NamedLocationIpPtrOutput { return v.Ip }).(NamedLocationIpPtrOutput)
}

type NamedLocationArrayOutput struct{ *pulumi.OutputState }

func (NamedLocationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NamedLocation)(nil)).Elem()
}

func (o NamedLocationArrayOutput) ToNamedLocationArrayOutput() NamedLocationArrayOutput {
	return o
}

func (o NamedLocationArrayOutput) ToNamedLocationArrayOutputWithContext(ctx context.Context) NamedLocationArrayOutput {
	return o
}

func (o NamedLocationArrayOutput) Index(i pulumi.IntInput) NamedLocationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NamedLocation {
		return vs[0].([]*NamedLocation)[vs[1].(int)]
	}).(NamedLocationOutput)
}

type NamedLocationMapOutput struct{ *pulumi.OutputState }

func (NamedLocationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NamedLocation)(nil)).Elem()
}

func (o NamedLocationMapOutput) ToNamedLocationMapOutput() NamedLocationMapOutput {
	return o
}

func (o NamedLocationMapOutput) ToNamedLocationMapOutputWithContext(ctx context.Context) NamedLocationMapOutput {
	return o
}

func (o NamedLocationMapOutput) MapIndex(k pulumi.StringInput) NamedLocationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NamedLocation {
		return vs[0].(map[string]*NamedLocation)[vs[1].(string)]
	}).(NamedLocationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NamedLocationInput)(nil)).Elem(), &NamedLocation{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamedLocationArrayInput)(nil)).Elem(), NamedLocationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamedLocationMapInput)(nil)).Elem(), NamedLocationMap{})
	pulumi.RegisterOutputType(NamedLocationOutput{})
	pulumi.RegisterOutputType(NamedLocationArrayOutput{})
	pulumi.RegisterOutputType(NamedLocationMapOutput{})
}

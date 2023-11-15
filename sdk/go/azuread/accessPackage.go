// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Access Package within Identity Governance in Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.
//
// When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner`, `Access package manager` or `Global Administrator`
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
//			exampleAccessPackageCatalog, err := azuread.NewAccessPackageCatalog(ctx, "exampleAccessPackageCatalog", &azuread.AccessPackageCatalogArgs{
//				DisplayName: pulumi.String("example-catalog"),
//				Description: pulumi.String("Example catalog"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewAccessPackage(ctx, "exampleAccessPackage", &azuread.AccessPackageArgs{
//				CatalogId:   exampleAccessPackageCatalog.ID(),
//				DisplayName: pulumi.String("access-package"),
//				Description: pulumi.String("Access Package"),
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
// Access Packages can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import azuread:index/accessPackage:AccessPackage example_package 00000000-0000-0000-0000-000000000000
//
// ```
type AccessPackage struct {
	pulumi.CustomResourceState

	// The ID of the Catalog this access package will be created in.
	CatalogId pulumi.StringOutput `pulumi:"catalogId"`
	// The description of the access package.
	Description pulumi.StringOutput `pulumi:"description"`
	// The display name of the access package.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Whether the access package is hidden from the requestor.
	Hidden pulumi.BoolPtrOutput `pulumi:"hidden"`
}

// NewAccessPackage registers a new resource with the given unique name, arguments, and options.
func NewAccessPackage(ctx *pulumi.Context,
	name string, args *AccessPackageArgs, opts ...pulumi.ResourceOption) (*AccessPackage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CatalogId == nil {
		return nil, errors.New("invalid value for required argument 'CatalogId'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccessPackage
	err := ctx.RegisterResource("azuread:index/accessPackage:AccessPackage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessPackage gets an existing AccessPackage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessPackageState, opts ...pulumi.ResourceOption) (*AccessPackage, error) {
	var resource AccessPackage
	err := ctx.ReadResource("azuread:index/accessPackage:AccessPackage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessPackage resources.
type accessPackageState struct {
	// The ID of the Catalog this access package will be created in.
	CatalogId *string `pulumi:"catalogId"`
	// The description of the access package.
	Description *string `pulumi:"description"`
	// The display name of the access package.
	DisplayName *string `pulumi:"displayName"`
	// Whether the access package is hidden from the requestor.
	Hidden *bool `pulumi:"hidden"`
}

type AccessPackageState struct {
	// The ID of the Catalog this access package will be created in.
	CatalogId pulumi.StringPtrInput
	// The description of the access package.
	Description pulumi.StringPtrInput
	// The display name of the access package.
	DisplayName pulumi.StringPtrInput
	// Whether the access package is hidden from the requestor.
	Hidden pulumi.BoolPtrInput
}

func (AccessPackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPackageState)(nil)).Elem()
}

type accessPackageArgs struct {
	// The ID of the Catalog this access package will be created in.
	CatalogId string `pulumi:"catalogId"`
	// The description of the access package.
	Description string `pulumi:"description"`
	// The display name of the access package.
	DisplayName string `pulumi:"displayName"`
	// Whether the access package is hidden from the requestor.
	Hidden *bool `pulumi:"hidden"`
}

// The set of arguments for constructing a AccessPackage resource.
type AccessPackageArgs struct {
	// The ID of the Catalog this access package will be created in.
	CatalogId pulumi.StringInput
	// The description of the access package.
	Description pulumi.StringInput
	// The display name of the access package.
	DisplayName pulumi.StringInput
	// Whether the access package is hidden from the requestor.
	Hidden pulumi.BoolPtrInput
}

func (AccessPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPackageArgs)(nil)).Elem()
}

type AccessPackageInput interface {
	pulumi.Input

	ToAccessPackageOutput() AccessPackageOutput
	ToAccessPackageOutputWithContext(ctx context.Context) AccessPackageOutput
}

func (*AccessPackage) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPackage)(nil)).Elem()
}

func (i *AccessPackage) ToAccessPackageOutput() AccessPackageOutput {
	return i.ToAccessPackageOutputWithContext(context.Background())
}

func (i *AccessPackage) ToAccessPackageOutputWithContext(ctx context.Context) AccessPackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPackageOutput)
}

// AccessPackageArrayInput is an input type that accepts AccessPackageArray and AccessPackageArrayOutput values.
// You can construct a concrete instance of `AccessPackageArrayInput` via:
//
//	AccessPackageArray{ AccessPackageArgs{...} }
type AccessPackageArrayInput interface {
	pulumi.Input

	ToAccessPackageArrayOutput() AccessPackageArrayOutput
	ToAccessPackageArrayOutputWithContext(context.Context) AccessPackageArrayOutput
}

type AccessPackageArray []AccessPackageInput

func (AccessPackageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessPackage)(nil)).Elem()
}

func (i AccessPackageArray) ToAccessPackageArrayOutput() AccessPackageArrayOutput {
	return i.ToAccessPackageArrayOutputWithContext(context.Background())
}

func (i AccessPackageArray) ToAccessPackageArrayOutputWithContext(ctx context.Context) AccessPackageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPackageArrayOutput)
}

// AccessPackageMapInput is an input type that accepts AccessPackageMap and AccessPackageMapOutput values.
// You can construct a concrete instance of `AccessPackageMapInput` via:
//
//	AccessPackageMap{ "key": AccessPackageArgs{...} }
type AccessPackageMapInput interface {
	pulumi.Input

	ToAccessPackageMapOutput() AccessPackageMapOutput
	ToAccessPackageMapOutputWithContext(context.Context) AccessPackageMapOutput
}

type AccessPackageMap map[string]AccessPackageInput

func (AccessPackageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessPackage)(nil)).Elem()
}

func (i AccessPackageMap) ToAccessPackageMapOutput() AccessPackageMapOutput {
	return i.ToAccessPackageMapOutputWithContext(context.Background())
}

func (i AccessPackageMap) ToAccessPackageMapOutputWithContext(ctx context.Context) AccessPackageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPackageMapOutput)
}

type AccessPackageOutput struct{ *pulumi.OutputState }

func (AccessPackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPackage)(nil)).Elem()
}

func (o AccessPackageOutput) ToAccessPackageOutput() AccessPackageOutput {
	return o
}

func (o AccessPackageOutput) ToAccessPackageOutputWithContext(ctx context.Context) AccessPackageOutput {
	return o
}

// The ID of the Catalog this access package will be created in.
func (o AccessPackageOutput) CatalogId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPackage) pulumi.StringOutput { return v.CatalogId }).(pulumi.StringOutput)
}

// The description of the access package.
func (o AccessPackageOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPackage) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The display name of the access package.
func (o AccessPackageOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPackage) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Whether the access package is hidden from the requestor.
func (o AccessPackageOutput) Hidden() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessPackage) pulumi.BoolPtrOutput { return v.Hidden }).(pulumi.BoolPtrOutput)
}

type AccessPackageArrayOutput struct{ *pulumi.OutputState }

func (AccessPackageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessPackage)(nil)).Elem()
}

func (o AccessPackageArrayOutput) ToAccessPackageArrayOutput() AccessPackageArrayOutput {
	return o
}

func (o AccessPackageArrayOutput) ToAccessPackageArrayOutputWithContext(ctx context.Context) AccessPackageArrayOutput {
	return o
}

func (o AccessPackageArrayOutput) Index(i pulumi.IntInput) AccessPackageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccessPackage {
		return vs[0].([]*AccessPackage)[vs[1].(int)]
	}).(AccessPackageOutput)
}

type AccessPackageMapOutput struct{ *pulumi.OutputState }

func (AccessPackageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessPackage)(nil)).Elem()
}

func (o AccessPackageMapOutput) ToAccessPackageMapOutput() AccessPackageMapOutput {
	return o
}

func (o AccessPackageMapOutput) ToAccessPackageMapOutputWithContext(ctx context.Context) AccessPackageMapOutput {
	return o
}

func (o AccessPackageMapOutput) MapIndex(k pulumi.StringInput) AccessPackageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccessPackage {
		return vs[0].(map[string]*AccessPackage)[vs[1].(string)]
	}).(AccessPackageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPackageInput)(nil)).Elem(), &AccessPackage{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPackageArrayInput)(nil)).Elem(), AccessPackageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPackageMapInput)(nil)).Elem(), AccessPackageMap{})
	pulumi.RegisterOutputType(AccessPackageOutput{})
	pulumi.RegisterOutputType(AccessPackageArrayOutput{})
	pulumi.RegisterOutputType(AccessPackageMapOutput{})
}

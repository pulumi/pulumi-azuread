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

// Manages the resources added to access package catalogs within Identity Governance in Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.
//
// When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner` or `Global Administrator`
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
//			_, err := azuread.NewGroup(ctx, "example", &azuread.GroupArgs{
//				DisplayName:     pulumi.String("example-group"),
//				SecurityEnabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewAccessPackageCatalog(ctx, "example", &azuread.AccessPackageCatalogArgs{
//				DisplayName: pulumi.String("example-catalog"),
//				Description: pulumi.String("Example catalog"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewAccessPackageResourceCatalogAssociation(ctx, "example", &azuread.AccessPackageResourceCatalogAssociationArgs{
//				CatalogId:            pulumi.Any(exampleCatalog.Id),
//				ResourceOriginId:     pulumi.Any(exampleGroup.ObjectId),
//				ResourceOriginSystem: pulumi.String("AadGroup"),
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
// The resource and catalog association can be imported using the catalog ID and the resource origin ID, e.g.
//
// ```sh
// $ pulumi import azuread:index/accessPackageResourceCatalogAssociation:AccessPackageResourceCatalogAssociation example 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111
// ```
//
//	-> This ID format is unique to Terraform and is composed of the Catalog ID and the Resource Origin ID in the format `{CatalogID}/{ResourceOriginID}`.
type AccessPackageResourceCatalogAssociation struct {
	pulumi.CustomResourceState

	// The unique ID of the access package catalog. Changing this forces a new resource to be created.
	CatalogId pulumi.StringOutput `pulumi:"catalogId"`
	// The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
	ResourceOriginId pulumi.StringOutput `pulumi:"resourceOriginId"`
	// The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
	ResourceOriginSystem pulumi.StringOutput `pulumi:"resourceOriginSystem"`
}

// NewAccessPackageResourceCatalogAssociation registers a new resource with the given unique name, arguments, and options.
func NewAccessPackageResourceCatalogAssociation(ctx *pulumi.Context,
	name string, args *AccessPackageResourceCatalogAssociationArgs, opts ...pulumi.ResourceOption) (*AccessPackageResourceCatalogAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CatalogId == nil {
		return nil, errors.New("invalid value for required argument 'CatalogId'")
	}
	if args.ResourceOriginId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceOriginId'")
	}
	if args.ResourceOriginSystem == nil {
		return nil, errors.New("invalid value for required argument 'ResourceOriginSystem'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccessPackageResourceCatalogAssociation
	err := ctx.RegisterResource("azuread:index/accessPackageResourceCatalogAssociation:AccessPackageResourceCatalogAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessPackageResourceCatalogAssociation gets an existing AccessPackageResourceCatalogAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessPackageResourceCatalogAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessPackageResourceCatalogAssociationState, opts ...pulumi.ResourceOption) (*AccessPackageResourceCatalogAssociation, error) {
	var resource AccessPackageResourceCatalogAssociation
	err := ctx.ReadResource("azuread:index/accessPackageResourceCatalogAssociation:AccessPackageResourceCatalogAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessPackageResourceCatalogAssociation resources.
type accessPackageResourceCatalogAssociationState struct {
	// The unique ID of the access package catalog. Changing this forces a new resource to be created.
	CatalogId *string `pulumi:"catalogId"`
	// The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
	ResourceOriginId *string `pulumi:"resourceOriginId"`
	// The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
	ResourceOriginSystem *string `pulumi:"resourceOriginSystem"`
}

type AccessPackageResourceCatalogAssociationState struct {
	// The unique ID of the access package catalog. Changing this forces a new resource to be created.
	CatalogId pulumi.StringPtrInput
	// The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
	ResourceOriginId pulumi.StringPtrInput
	// The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
	ResourceOriginSystem pulumi.StringPtrInput
}

func (AccessPackageResourceCatalogAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPackageResourceCatalogAssociationState)(nil)).Elem()
}

type accessPackageResourceCatalogAssociationArgs struct {
	// The unique ID of the access package catalog. Changing this forces a new resource to be created.
	CatalogId string `pulumi:"catalogId"`
	// The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
	ResourceOriginId string `pulumi:"resourceOriginId"`
	// The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
	ResourceOriginSystem string `pulumi:"resourceOriginSystem"`
}

// The set of arguments for constructing a AccessPackageResourceCatalogAssociation resource.
type AccessPackageResourceCatalogAssociationArgs struct {
	// The unique ID of the access package catalog. Changing this forces a new resource to be created.
	CatalogId pulumi.StringInput
	// The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
	ResourceOriginId pulumi.StringInput
	// The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
	ResourceOriginSystem pulumi.StringInput
}

func (AccessPackageResourceCatalogAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPackageResourceCatalogAssociationArgs)(nil)).Elem()
}

type AccessPackageResourceCatalogAssociationInput interface {
	pulumi.Input

	ToAccessPackageResourceCatalogAssociationOutput() AccessPackageResourceCatalogAssociationOutput
	ToAccessPackageResourceCatalogAssociationOutputWithContext(ctx context.Context) AccessPackageResourceCatalogAssociationOutput
}

func (*AccessPackageResourceCatalogAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPackageResourceCatalogAssociation)(nil)).Elem()
}

func (i *AccessPackageResourceCatalogAssociation) ToAccessPackageResourceCatalogAssociationOutput() AccessPackageResourceCatalogAssociationOutput {
	return i.ToAccessPackageResourceCatalogAssociationOutputWithContext(context.Background())
}

func (i *AccessPackageResourceCatalogAssociation) ToAccessPackageResourceCatalogAssociationOutputWithContext(ctx context.Context) AccessPackageResourceCatalogAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPackageResourceCatalogAssociationOutput)
}

// AccessPackageResourceCatalogAssociationArrayInput is an input type that accepts AccessPackageResourceCatalogAssociationArray and AccessPackageResourceCatalogAssociationArrayOutput values.
// You can construct a concrete instance of `AccessPackageResourceCatalogAssociationArrayInput` via:
//
//	AccessPackageResourceCatalogAssociationArray{ AccessPackageResourceCatalogAssociationArgs{...} }
type AccessPackageResourceCatalogAssociationArrayInput interface {
	pulumi.Input

	ToAccessPackageResourceCatalogAssociationArrayOutput() AccessPackageResourceCatalogAssociationArrayOutput
	ToAccessPackageResourceCatalogAssociationArrayOutputWithContext(context.Context) AccessPackageResourceCatalogAssociationArrayOutput
}

type AccessPackageResourceCatalogAssociationArray []AccessPackageResourceCatalogAssociationInput

func (AccessPackageResourceCatalogAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessPackageResourceCatalogAssociation)(nil)).Elem()
}

func (i AccessPackageResourceCatalogAssociationArray) ToAccessPackageResourceCatalogAssociationArrayOutput() AccessPackageResourceCatalogAssociationArrayOutput {
	return i.ToAccessPackageResourceCatalogAssociationArrayOutputWithContext(context.Background())
}

func (i AccessPackageResourceCatalogAssociationArray) ToAccessPackageResourceCatalogAssociationArrayOutputWithContext(ctx context.Context) AccessPackageResourceCatalogAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPackageResourceCatalogAssociationArrayOutput)
}

// AccessPackageResourceCatalogAssociationMapInput is an input type that accepts AccessPackageResourceCatalogAssociationMap and AccessPackageResourceCatalogAssociationMapOutput values.
// You can construct a concrete instance of `AccessPackageResourceCatalogAssociationMapInput` via:
//
//	AccessPackageResourceCatalogAssociationMap{ "key": AccessPackageResourceCatalogAssociationArgs{...} }
type AccessPackageResourceCatalogAssociationMapInput interface {
	pulumi.Input

	ToAccessPackageResourceCatalogAssociationMapOutput() AccessPackageResourceCatalogAssociationMapOutput
	ToAccessPackageResourceCatalogAssociationMapOutputWithContext(context.Context) AccessPackageResourceCatalogAssociationMapOutput
}

type AccessPackageResourceCatalogAssociationMap map[string]AccessPackageResourceCatalogAssociationInput

func (AccessPackageResourceCatalogAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessPackageResourceCatalogAssociation)(nil)).Elem()
}

func (i AccessPackageResourceCatalogAssociationMap) ToAccessPackageResourceCatalogAssociationMapOutput() AccessPackageResourceCatalogAssociationMapOutput {
	return i.ToAccessPackageResourceCatalogAssociationMapOutputWithContext(context.Background())
}

func (i AccessPackageResourceCatalogAssociationMap) ToAccessPackageResourceCatalogAssociationMapOutputWithContext(ctx context.Context) AccessPackageResourceCatalogAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPackageResourceCatalogAssociationMapOutput)
}

type AccessPackageResourceCatalogAssociationOutput struct{ *pulumi.OutputState }

func (AccessPackageResourceCatalogAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPackageResourceCatalogAssociation)(nil)).Elem()
}

func (o AccessPackageResourceCatalogAssociationOutput) ToAccessPackageResourceCatalogAssociationOutput() AccessPackageResourceCatalogAssociationOutput {
	return o
}

func (o AccessPackageResourceCatalogAssociationOutput) ToAccessPackageResourceCatalogAssociationOutputWithContext(ctx context.Context) AccessPackageResourceCatalogAssociationOutput {
	return o
}

// The unique ID of the access package catalog. Changing this forces a new resource to be created.
func (o AccessPackageResourceCatalogAssociationOutput) CatalogId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPackageResourceCatalogAssociation) pulumi.StringOutput { return v.CatalogId }).(pulumi.StringOutput)
}

// The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
func (o AccessPackageResourceCatalogAssociationOutput) ResourceOriginId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPackageResourceCatalogAssociation) pulumi.StringOutput { return v.ResourceOriginId }).(pulumi.StringOutput)
}

// The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
func (o AccessPackageResourceCatalogAssociationOutput) ResourceOriginSystem() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPackageResourceCatalogAssociation) pulumi.StringOutput { return v.ResourceOriginSystem }).(pulumi.StringOutput)
}

type AccessPackageResourceCatalogAssociationArrayOutput struct{ *pulumi.OutputState }

func (AccessPackageResourceCatalogAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessPackageResourceCatalogAssociation)(nil)).Elem()
}

func (o AccessPackageResourceCatalogAssociationArrayOutput) ToAccessPackageResourceCatalogAssociationArrayOutput() AccessPackageResourceCatalogAssociationArrayOutput {
	return o
}

func (o AccessPackageResourceCatalogAssociationArrayOutput) ToAccessPackageResourceCatalogAssociationArrayOutputWithContext(ctx context.Context) AccessPackageResourceCatalogAssociationArrayOutput {
	return o
}

func (o AccessPackageResourceCatalogAssociationArrayOutput) Index(i pulumi.IntInput) AccessPackageResourceCatalogAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccessPackageResourceCatalogAssociation {
		return vs[0].([]*AccessPackageResourceCatalogAssociation)[vs[1].(int)]
	}).(AccessPackageResourceCatalogAssociationOutput)
}

type AccessPackageResourceCatalogAssociationMapOutput struct{ *pulumi.OutputState }

func (AccessPackageResourceCatalogAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessPackageResourceCatalogAssociation)(nil)).Elem()
}

func (o AccessPackageResourceCatalogAssociationMapOutput) ToAccessPackageResourceCatalogAssociationMapOutput() AccessPackageResourceCatalogAssociationMapOutput {
	return o
}

func (o AccessPackageResourceCatalogAssociationMapOutput) ToAccessPackageResourceCatalogAssociationMapOutputWithContext(ctx context.Context) AccessPackageResourceCatalogAssociationMapOutput {
	return o
}

func (o AccessPackageResourceCatalogAssociationMapOutput) MapIndex(k pulumi.StringInput) AccessPackageResourceCatalogAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccessPackageResourceCatalogAssociation {
		return vs[0].(map[string]*AccessPackageResourceCatalogAssociation)[vs[1].(string)]
	}).(AccessPackageResourceCatalogAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPackageResourceCatalogAssociationInput)(nil)).Elem(), &AccessPackageResourceCatalogAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPackageResourceCatalogAssociationArrayInput)(nil)).Elem(), AccessPackageResourceCatalogAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPackageResourceCatalogAssociationMapInput)(nil)).Elem(), AccessPackageResourceCatalogAssociationMap{})
	pulumi.RegisterOutputType(AccessPackageResourceCatalogAssociationOutput{})
	pulumi.RegisterOutputType(AccessPackageResourceCatalogAssociationArrayOutput{})
	pulumi.RegisterOutputType(AccessPackageResourceCatalogAssociationMapOutput{})
}

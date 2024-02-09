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
//			authorized, err := azuread.NewApplicationRegistration(ctx, "authorized", &azuread.ApplicationRegistrationArgs{
//				DisplayName: pulumi.String("example-authorized-app"),
//			})
//			if err != nil {
//				return err
//			}
//			authorizer, err := azuread.NewApplication(ctx, "authorizer", &azuread.ApplicationArgs{
//				DisplayName: pulumi.String("example-authorizing-app"),
//				Api: &azuread.ApplicationApiArgs{
//					Oauth2PermissionScopes: azuread.ApplicationApiOauth2PermissionScopeArray{
//						&azuread.ApplicationApiOauth2PermissionScopeArgs{
//							AdminConsentDescription: pulumi.String("Administer the application"),
//							AdminConsentDisplayName: pulumi.String("Administer"),
//							Enabled:                 pulumi.Bool(true),
//							Id:                      pulumi.String("00000000-0000-0000-0000-000000000000"),
//							Type:                    pulumi.String("Admin"),
//							Value:                   pulumi.String("administer"),
//						},
//						&azuread.ApplicationApiOauth2PermissionScopeArgs{
//							AdminConsentDescription: pulumi.String("Access the application"),
//							AdminConsentDisplayName: pulumi.String("Access"),
//							Enabled:                 pulumi.Bool(true),
//							Id:                      pulumi.String("11111111-1111-1111-1111-111111111111"),
//							Type:                    pulumi.String("User"),
//							UserConsentDescription:  pulumi.String("Access the application"),
//							UserConsentDisplayName:  pulumi.String("Access"),
//							Value:                   pulumi.String("user_impersonation"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewApplicationPreAuthorized(ctx, "example", &azuread.ApplicationPreAuthorizedArgs{
//				ApplicationId:      authorizer.ID(),
//				AuthorizedClientId: authorized.ClientId,
//				PermissionIds: pulumi.StringArray{
//					pulumi.String("00000000-0000-0000-0000-000000000000"),
//					pulumi.String("11111111-1111-1111-1111-111111111111"),
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
// Pre-authorized applications can be imported using the object ID of the authorizing application and the application ID of the application being authorized, e.g.
//
// ```sh
// $ pulumi import azuread:index/applicationPreAuthorized:ApplicationPreAuthorized example 00000000-0000-0000-0000-000000000000/preAuthorizedApplication/11111111-1111-1111-1111-111111111111
// ```
//
//	-> This ID format is unique to Terraform and is composed of the authorizing application's object ID, the string "preAuthorizedApplication" and the authorized application's application ID (client ID) in the format `{ObjectId}/preAuthorizedApplication/{ApplicationId}`.
type ApplicationPreAuthorized struct {
	pulumi.CustomResourceState

	// The resource ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// The object ID of the application to which this pre-authorized application should be added
	//
	// Deprecated: The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
	ApplicationObjectId pulumi.StringOutput `pulumi:"applicationObjectId"`
	// The application ID of the pre-authorized application
	//
	// Deprecated: The `authorized_app_id` property has been replaced with the `authorized_client_id` property and will be removed in version 3.0 of the AzureAD provider
	AuthorizedAppId pulumi.StringOutput `pulumi:"authorizedAppId"`
	// The client ID of the application being authorized. Changing this field forces a new resource to be created.
	AuthorizedClientId pulumi.StringOutput `pulumi:"authorizedClientId"`
	// A set of permission scope IDs required by the authorized application.
	PermissionIds pulumi.StringArrayOutput `pulumi:"permissionIds"`
}

// NewApplicationPreAuthorized registers a new resource with the given unique name, arguments, and options.
func NewApplicationPreAuthorized(ctx *pulumi.Context,
	name string, args *ApplicationPreAuthorizedArgs, opts ...pulumi.ResourceOption) (*ApplicationPreAuthorized, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PermissionIds == nil {
		return nil, errors.New("invalid value for required argument 'PermissionIds'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationPreAuthorized
	err := ctx.RegisterResource("azuread:index/applicationPreAuthorized:ApplicationPreAuthorized", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationPreAuthorized gets an existing ApplicationPreAuthorized resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationPreAuthorized(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationPreAuthorizedState, opts ...pulumi.ResourceOption) (*ApplicationPreAuthorized, error) {
	var resource ApplicationPreAuthorized
	err := ctx.ReadResource("azuread:index/applicationPreAuthorized:ApplicationPreAuthorized", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationPreAuthorized resources.
type applicationPreAuthorizedState struct {
	// The resource ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
	ApplicationId *string `pulumi:"applicationId"`
	// The object ID of the application to which this pre-authorized application should be added
	//
	// Deprecated: The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
	ApplicationObjectId *string `pulumi:"applicationObjectId"`
	// The application ID of the pre-authorized application
	//
	// Deprecated: The `authorized_app_id` property has been replaced with the `authorized_client_id` property and will be removed in version 3.0 of the AzureAD provider
	AuthorizedAppId *string `pulumi:"authorizedAppId"`
	// The client ID of the application being authorized. Changing this field forces a new resource to be created.
	AuthorizedClientId *string `pulumi:"authorizedClientId"`
	// A set of permission scope IDs required by the authorized application.
	PermissionIds []string `pulumi:"permissionIds"`
}

type ApplicationPreAuthorizedState struct {
	// The resource ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
	ApplicationId pulumi.StringPtrInput
	// The object ID of the application to which this pre-authorized application should be added
	//
	// Deprecated: The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
	ApplicationObjectId pulumi.StringPtrInput
	// The application ID of the pre-authorized application
	//
	// Deprecated: The `authorized_app_id` property has been replaced with the `authorized_client_id` property and will be removed in version 3.0 of the AzureAD provider
	AuthorizedAppId pulumi.StringPtrInput
	// The client ID of the application being authorized. Changing this field forces a new resource to be created.
	AuthorizedClientId pulumi.StringPtrInput
	// A set of permission scope IDs required by the authorized application.
	PermissionIds pulumi.StringArrayInput
}

func (ApplicationPreAuthorizedState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationPreAuthorizedState)(nil)).Elem()
}

type applicationPreAuthorizedArgs struct {
	// The resource ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
	ApplicationId *string `pulumi:"applicationId"`
	// The object ID of the application to which this pre-authorized application should be added
	//
	// Deprecated: The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
	ApplicationObjectId *string `pulumi:"applicationObjectId"`
	// The application ID of the pre-authorized application
	//
	// Deprecated: The `authorized_app_id` property has been replaced with the `authorized_client_id` property and will be removed in version 3.0 of the AzureAD provider
	AuthorizedAppId *string `pulumi:"authorizedAppId"`
	// The client ID of the application being authorized. Changing this field forces a new resource to be created.
	AuthorizedClientId *string `pulumi:"authorizedClientId"`
	// A set of permission scope IDs required by the authorized application.
	PermissionIds []string `pulumi:"permissionIds"`
}

// The set of arguments for constructing a ApplicationPreAuthorized resource.
type ApplicationPreAuthorizedArgs struct {
	// The resource ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
	ApplicationId pulumi.StringPtrInput
	// The object ID of the application to which this pre-authorized application should be added
	//
	// Deprecated: The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
	ApplicationObjectId pulumi.StringPtrInput
	// The application ID of the pre-authorized application
	//
	// Deprecated: The `authorized_app_id` property has been replaced with the `authorized_client_id` property and will be removed in version 3.0 of the AzureAD provider
	AuthorizedAppId pulumi.StringPtrInput
	// The client ID of the application being authorized. Changing this field forces a new resource to be created.
	AuthorizedClientId pulumi.StringPtrInput
	// A set of permission scope IDs required by the authorized application.
	PermissionIds pulumi.StringArrayInput
}

func (ApplicationPreAuthorizedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationPreAuthorizedArgs)(nil)).Elem()
}

type ApplicationPreAuthorizedInput interface {
	pulumi.Input

	ToApplicationPreAuthorizedOutput() ApplicationPreAuthorizedOutput
	ToApplicationPreAuthorizedOutputWithContext(ctx context.Context) ApplicationPreAuthorizedOutput
}

func (*ApplicationPreAuthorized) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPreAuthorized)(nil)).Elem()
}

func (i *ApplicationPreAuthorized) ToApplicationPreAuthorizedOutput() ApplicationPreAuthorizedOutput {
	return i.ToApplicationPreAuthorizedOutputWithContext(context.Background())
}

func (i *ApplicationPreAuthorized) ToApplicationPreAuthorizedOutputWithContext(ctx context.Context) ApplicationPreAuthorizedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPreAuthorizedOutput)
}

// ApplicationPreAuthorizedArrayInput is an input type that accepts ApplicationPreAuthorizedArray and ApplicationPreAuthorizedArrayOutput values.
// You can construct a concrete instance of `ApplicationPreAuthorizedArrayInput` via:
//
//	ApplicationPreAuthorizedArray{ ApplicationPreAuthorizedArgs{...} }
type ApplicationPreAuthorizedArrayInput interface {
	pulumi.Input

	ToApplicationPreAuthorizedArrayOutput() ApplicationPreAuthorizedArrayOutput
	ToApplicationPreAuthorizedArrayOutputWithContext(context.Context) ApplicationPreAuthorizedArrayOutput
}

type ApplicationPreAuthorizedArray []ApplicationPreAuthorizedInput

func (ApplicationPreAuthorizedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationPreAuthorized)(nil)).Elem()
}

func (i ApplicationPreAuthorizedArray) ToApplicationPreAuthorizedArrayOutput() ApplicationPreAuthorizedArrayOutput {
	return i.ToApplicationPreAuthorizedArrayOutputWithContext(context.Background())
}

func (i ApplicationPreAuthorizedArray) ToApplicationPreAuthorizedArrayOutputWithContext(ctx context.Context) ApplicationPreAuthorizedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPreAuthorizedArrayOutput)
}

// ApplicationPreAuthorizedMapInput is an input type that accepts ApplicationPreAuthorizedMap and ApplicationPreAuthorizedMapOutput values.
// You can construct a concrete instance of `ApplicationPreAuthorizedMapInput` via:
//
//	ApplicationPreAuthorizedMap{ "key": ApplicationPreAuthorizedArgs{...} }
type ApplicationPreAuthorizedMapInput interface {
	pulumi.Input

	ToApplicationPreAuthorizedMapOutput() ApplicationPreAuthorizedMapOutput
	ToApplicationPreAuthorizedMapOutputWithContext(context.Context) ApplicationPreAuthorizedMapOutput
}

type ApplicationPreAuthorizedMap map[string]ApplicationPreAuthorizedInput

func (ApplicationPreAuthorizedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationPreAuthorized)(nil)).Elem()
}

func (i ApplicationPreAuthorizedMap) ToApplicationPreAuthorizedMapOutput() ApplicationPreAuthorizedMapOutput {
	return i.ToApplicationPreAuthorizedMapOutputWithContext(context.Background())
}

func (i ApplicationPreAuthorizedMap) ToApplicationPreAuthorizedMapOutputWithContext(ctx context.Context) ApplicationPreAuthorizedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPreAuthorizedMapOutput)
}

type ApplicationPreAuthorizedOutput struct{ *pulumi.OutputState }

func (ApplicationPreAuthorizedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationPreAuthorized)(nil)).Elem()
}

func (o ApplicationPreAuthorizedOutput) ToApplicationPreAuthorizedOutput() ApplicationPreAuthorizedOutput {
	return o
}

func (o ApplicationPreAuthorizedOutput) ToApplicationPreAuthorizedOutputWithContext(ctx context.Context) ApplicationPreAuthorizedOutput {
	return o
}

// The resource ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
func (o ApplicationPreAuthorizedOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPreAuthorized) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// The object ID of the application to which this pre-authorized application should be added
//
// Deprecated: The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
func (o ApplicationPreAuthorizedOutput) ApplicationObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPreAuthorized) pulumi.StringOutput { return v.ApplicationObjectId }).(pulumi.StringOutput)
}

// The application ID of the pre-authorized application
//
// Deprecated: The `authorized_app_id` property has been replaced with the `authorized_client_id` property and will be removed in version 3.0 of the AzureAD provider
func (o ApplicationPreAuthorizedOutput) AuthorizedAppId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPreAuthorized) pulumi.StringOutput { return v.AuthorizedAppId }).(pulumi.StringOutput)
}

// The client ID of the application being authorized. Changing this field forces a new resource to be created.
func (o ApplicationPreAuthorizedOutput) AuthorizedClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationPreAuthorized) pulumi.StringOutput { return v.AuthorizedClientId }).(pulumi.StringOutput)
}

// A set of permission scope IDs required by the authorized application.
func (o ApplicationPreAuthorizedOutput) PermissionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationPreAuthorized) pulumi.StringArrayOutput { return v.PermissionIds }).(pulumi.StringArrayOutput)
}

type ApplicationPreAuthorizedArrayOutput struct{ *pulumi.OutputState }

func (ApplicationPreAuthorizedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationPreAuthorized)(nil)).Elem()
}

func (o ApplicationPreAuthorizedArrayOutput) ToApplicationPreAuthorizedArrayOutput() ApplicationPreAuthorizedArrayOutput {
	return o
}

func (o ApplicationPreAuthorizedArrayOutput) ToApplicationPreAuthorizedArrayOutputWithContext(ctx context.Context) ApplicationPreAuthorizedArrayOutput {
	return o
}

func (o ApplicationPreAuthorizedArrayOutput) Index(i pulumi.IntInput) ApplicationPreAuthorizedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationPreAuthorized {
		return vs[0].([]*ApplicationPreAuthorized)[vs[1].(int)]
	}).(ApplicationPreAuthorizedOutput)
}

type ApplicationPreAuthorizedMapOutput struct{ *pulumi.OutputState }

func (ApplicationPreAuthorizedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationPreAuthorized)(nil)).Elem()
}

func (o ApplicationPreAuthorizedMapOutput) ToApplicationPreAuthorizedMapOutput() ApplicationPreAuthorizedMapOutput {
	return o
}

func (o ApplicationPreAuthorizedMapOutput) ToApplicationPreAuthorizedMapOutputWithContext(ctx context.Context) ApplicationPreAuthorizedMapOutput {
	return o
}

func (o ApplicationPreAuthorizedMapOutput) MapIndex(k pulumi.StringInput) ApplicationPreAuthorizedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationPreAuthorized {
		return vs[0].(map[string]*ApplicationPreAuthorized)[vs[1].(string)]
	}).(ApplicationPreAuthorizedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationPreAuthorizedInput)(nil)).Elem(), &ApplicationPreAuthorized{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationPreAuthorizedArrayInput)(nil)).Elem(), ApplicationPreAuthorizedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationPreAuthorizedMapInput)(nil)).Elem(), ApplicationPreAuthorizedMap{})
	pulumi.RegisterOutputType(ApplicationPreAuthorizedOutput{})
	pulumi.RegisterOutputType(ApplicationPreAuthorizedArrayOutput{})
	pulumi.RegisterOutputType(ApplicationPreAuthorizedMapOutput{})
}

// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuread/sdk/v6/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an app role assignment for a group, user or service principal. Can be used to grant admin consent for application permissions.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires one of the following application roles: `AppRoleAssignment.ReadWrite.All` and `Application.Read.All`, or `AppRoleAssignment.ReadWrite.All` and `Directory.Read.All`, or `Application.ReadWrite.All`, or `Directory.ReadWrite.All`
//
// When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`
//
// ## Example Usage
//
// *App role assignment for accessing Microsoft Graph*
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
//			wellKnown, err := azuread.GetApplicationPublishedAppIds(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			msgraph, err := azuread.NewServicePrincipal(ctx, "msgraph", &azuread.ServicePrincipalArgs{
//				ClientId:    pulumi.String(wellKnown.Result.MicrosoftGraph),
//				UseExisting: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			example, err := azuread.NewApplication(ctx, "example", &azuread.ApplicationArgs{
//				DisplayName: pulumi.String("example"),
//				RequiredResourceAccesses: azuread.ApplicationRequiredResourceAccessArray{
//					&azuread.ApplicationRequiredResourceAccessArgs{
//						ResourceAppId: pulumi.String(wellKnown.Result.MicrosoftGraph),
//						ResourceAccesses: azuread.ApplicationRequiredResourceAccessResourceAccessArray{
//							&azuread.ApplicationRequiredResourceAccessResourceAccessArgs{
//								Id: msgraph.AppRoleIds.ApplyT(func(appRoleIds map[string]string) (string, error) {
//									return appRoleIds.User.Read.All, nil
//								}).(pulumi.StringOutput),
//								Type: pulumi.String("Role"),
//							},
//							&azuread.ApplicationRequiredResourceAccessResourceAccessArgs{
//								Id: msgraph.Oauth2PermissionScopeIds.ApplyT(func(oauth2PermissionScopeIds map[string]string) (string, error) {
//									return oauth2PermissionScopeIds.User.ReadWrite, nil
//								}).(pulumi.StringOutput),
//								Type: pulumi.String("Scope"),
//							},
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleServicePrincipal, err := azuread.NewServicePrincipal(ctx, "example", &azuread.ServicePrincipalArgs{
//				ClientId: example.ClientId,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewAppRoleAssignment(ctx, "example", &azuread.AppRoleAssignmentArgs{
//				AppRoleId: msgraph.AppRoleIds.ApplyT(func(appRoleIds map[string]string) (string, error) {
//					return appRoleIds.User.Read.All, nil
//				}).(pulumi.StringOutput),
//				PrincipalObjectId: exampleServicePrincipal.ObjectId,
//				ResourceObjectId:  msgraph.ObjectId,
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
// *App role assignment for internal application*
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
//			internal, err := azuread.NewApplication(ctx, "internal", &azuread.ApplicationArgs{
//				DisplayName: pulumi.String("internal"),
//				AppRoles: azuread.ApplicationAppRoleTypeArray{
//					&azuread.ApplicationAppRoleTypeArgs{
//						AllowedMemberTypes: pulumi.StringArray{
//							pulumi.String("Application"),
//						},
//						Description: pulumi.String("Apps can query the database"),
//						DisplayName: pulumi.String("Query"),
//						Enabled:     pulumi.Bool(true),
//						Id:          pulumi.String("00000000-0000-0000-0000-111111111111"),
//						Value:       pulumi.String("Query.All"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			internalServicePrincipal, err := azuread.NewServicePrincipal(ctx, "internal", &azuread.ServicePrincipalArgs{
//				ClientId: internal.ClientId,
//			})
//			if err != nil {
//				return err
//			}
//			example, err := azuread.NewApplication(ctx, "example", &azuread.ApplicationArgs{
//				DisplayName: pulumi.String("example"),
//				RequiredResourceAccesses: azuread.ApplicationRequiredResourceAccessArray{
//					&azuread.ApplicationRequiredResourceAccessArgs{
//						ResourceAppId: internal.ClientId,
//						ResourceAccesses: azuread.ApplicationRequiredResourceAccessResourceAccessArray{
//							&azuread.ApplicationRequiredResourceAccessResourceAccessArgs{
//								Id: internalServicePrincipal.AppRoleIds.ApplyT(func(appRoleIds map[string]string) (string, error) {
//									return appRoleIds.Query.All, nil
//								}).(pulumi.StringOutput),
//								Type: pulumi.String("Role"),
//							},
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleServicePrincipal, err := azuread.NewServicePrincipal(ctx, "example", &azuread.ServicePrincipalArgs{
//				ClientId: example.ClientId,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewAppRoleAssignment(ctx, "example", &azuread.AppRoleAssignmentArgs{
//				AppRoleId: internalServicePrincipal.AppRoleIds.ApplyT(func(appRoleIds map[string]string) (string, error) {
//					return appRoleIds.Query.All, nil
//				}).(pulumi.StringOutput),
//				PrincipalObjectId: exampleServicePrincipal.ObjectId,
//				ResourceObjectId:  internalServicePrincipal.ObjectId,
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
// *Assign a user and group to an internal application*
//
// ## Import
//
// App role assignments can be imported using the object ID of the service principal representing the resource and the ID of the app role assignment (note: _not_ the ID of the app role), e.g.
//
// ```sh
// $ pulumi import azuread:index/appRoleAssignment:AppRoleAssignment example /servicePrincipals/00000000-0000-0000-0000-000000000000/appRoleAssignedTo/aaBBcDDeFG6h5JKLMN2PQrrssTTUUvWWxxxxxyyyzzz
// ```
//
// -> This ID format is unique to Terraform and is composed of the Resource Service Principal Object ID and the ID of the App Role Assignment in the format `/servicePrincipals/{ResourcePrincipalID}/appRoleAssignedTo/{AppRoleAssignmentID}`.
type AppRoleAssignment struct {
	pulumi.CustomResourceState

	// The ID of the app role to be assigned, or the default role ID `00000000-0000-0000-0000-000000000000`. Changing this forces a new resource to be created.
	AppRoleId pulumi.StringOutput `pulumi:"appRoleId"`
	// The display name of the principal to which the app role is assigned.
	PrincipalDisplayName pulumi.StringOutput `pulumi:"principalDisplayName"`
	// The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	PrincipalObjectId pulumi.StringOutput `pulumi:"principalObjectId"`
	// The object type of the principal to which the app role is assigned.
	PrincipalType pulumi.StringOutput `pulumi:"principalType"`
	// The display name of the application representing the resource.
	ResourceDisplayName pulumi.StringOutput `pulumi:"resourceDisplayName"`
	// The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
	ResourceObjectId pulumi.StringOutput `pulumi:"resourceObjectId"`
}

// NewAppRoleAssignment registers a new resource with the given unique name, arguments, and options.
func NewAppRoleAssignment(ctx *pulumi.Context,
	name string, args *AppRoleAssignmentArgs, opts ...pulumi.ResourceOption) (*AppRoleAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppRoleId == nil {
		return nil, errors.New("invalid value for required argument 'AppRoleId'")
	}
	if args.PrincipalObjectId == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalObjectId'")
	}
	if args.ResourceObjectId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceObjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AppRoleAssignment
	err := ctx.RegisterResource("azuread:index/appRoleAssignment:AppRoleAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppRoleAssignment gets an existing AppRoleAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppRoleAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppRoleAssignmentState, opts ...pulumi.ResourceOption) (*AppRoleAssignment, error) {
	var resource AppRoleAssignment
	err := ctx.ReadResource("azuread:index/appRoleAssignment:AppRoleAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppRoleAssignment resources.
type appRoleAssignmentState struct {
	// The ID of the app role to be assigned, or the default role ID `00000000-0000-0000-0000-000000000000`. Changing this forces a new resource to be created.
	AppRoleId *string `pulumi:"appRoleId"`
	// The display name of the principal to which the app role is assigned.
	PrincipalDisplayName *string `pulumi:"principalDisplayName"`
	// The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	PrincipalObjectId *string `pulumi:"principalObjectId"`
	// The object type of the principal to which the app role is assigned.
	PrincipalType *string `pulumi:"principalType"`
	// The display name of the application representing the resource.
	ResourceDisplayName *string `pulumi:"resourceDisplayName"`
	// The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
	ResourceObjectId *string `pulumi:"resourceObjectId"`
}

type AppRoleAssignmentState struct {
	// The ID of the app role to be assigned, or the default role ID `00000000-0000-0000-0000-000000000000`. Changing this forces a new resource to be created.
	AppRoleId pulumi.StringPtrInput
	// The display name of the principal to which the app role is assigned.
	PrincipalDisplayName pulumi.StringPtrInput
	// The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	PrincipalObjectId pulumi.StringPtrInput
	// The object type of the principal to which the app role is assigned.
	PrincipalType pulumi.StringPtrInput
	// The display name of the application representing the resource.
	ResourceDisplayName pulumi.StringPtrInput
	// The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
	ResourceObjectId pulumi.StringPtrInput
}

func (AppRoleAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*appRoleAssignmentState)(nil)).Elem()
}

type appRoleAssignmentArgs struct {
	// The ID of the app role to be assigned, or the default role ID `00000000-0000-0000-0000-000000000000`. Changing this forces a new resource to be created.
	AppRoleId string `pulumi:"appRoleId"`
	// The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	PrincipalObjectId string `pulumi:"principalObjectId"`
	// The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
	ResourceObjectId string `pulumi:"resourceObjectId"`
}

// The set of arguments for constructing a AppRoleAssignment resource.
type AppRoleAssignmentArgs struct {
	// The ID of the app role to be assigned, or the default role ID `00000000-0000-0000-0000-000000000000`. Changing this forces a new resource to be created.
	AppRoleId pulumi.StringInput
	// The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	PrincipalObjectId pulumi.StringInput
	// The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
	ResourceObjectId pulumi.StringInput
}

func (AppRoleAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appRoleAssignmentArgs)(nil)).Elem()
}

type AppRoleAssignmentInput interface {
	pulumi.Input

	ToAppRoleAssignmentOutput() AppRoleAssignmentOutput
	ToAppRoleAssignmentOutputWithContext(ctx context.Context) AppRoleAssignmentOutput
}

func (*AppRoleAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**AppRoleAssignment)(nil)).Elem()
}

func (i *AppRoleAssignment) ToAppRoleAssignmentOutput() AppRoleAssignmentOutput {
	return i.ToAppRoleAssignmentOutputWithContext(context.Background())
}

func (i *AppRoleAssignment) ToAppRoleAssignmentOutputWithContext(ctx context.Context) AppRoleAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppRoleAssignmentOutput)
}

// AppRoleAssignmentArrayInput is an input type that accepts AppRoleAssignmentArray and AppRoleAssignmentArrayOutput values.
// You can construct a concrete instance of `AppRoleAssignmentArrayInput` via:
//
//	AppRoleAssignmentArray{ AppRoleAssignmentArgs{...} }
type AppRoleAssignmentArrayInput interface {
	pulumi.Input

	ToAppRoleAssignmentArrayOutput() AppRoleAssignmentArrayOutput
	ToAppRoleAssignmentArrayOutputWithContext(context.Context) AppRoleAssignmentArrayOutput
}

type AppRoleAssignmentArray []AppRoleAssignmentInput

func (AppRoleAssignmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppRoleAssignment)(nil)).Elem()
}

func (i AppRoleAssignmentArray) ToAppRoleAssignmentArrayOutput() AppRoleAssignmentArrayOutput {
	return i.ToAppRoleAssignmentArrayOutputWithContext(context.Background())
}

func (i AppRoleAssignmentArray) ToAppRoleAssignmentArrayOutputWithContext(ctx context.Context) AppRoleAssignmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppRoleAssignmentArrayOutput)
}

// AppRoleAssignmentMapInput is an input type that accepts AppRoleAssignmentMap and AppRoleAssignmentMapOutput values.
// You can construct a concrete instance of `AppRoleAssignmentMapInput` via:
//
//	AppRoleAssignmentMap{ "key": AppRoleAssignmentArgs{...} }
type AppRoleAssignmentMapInput interface {
	pulumi.Input

	ToAppRoleAssignmentMapOutput() AppRoleAssignmentMapOutput
	ToAppRoleAssignmentMapOutputWithContext(context.Context) AppRoleAssignmentMapOutput
}

type AppRoleAssignmentMap map[string]AppRoleAssignmentInput

func (AppRoleAssignmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppRoleAssignment)(nil)).Elem()
}

func (i AppRoleAssignmentMap) ToAppRoleAssignmentMapOutput() AppRoleAssignmentMapOutput {
	return i.ToAppRoleAssignmentMapOutputWithContext(context.Background())
}

func (i AppRoleAssignmentMap) ToAppRoleAssignmentMapOutputWithContext(ctx context.Context) AppRoleAssignmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppRoleAssignmentMapOutput)
}

type AppRoleAssignmentOutput struct{ *pulumi.OutputState }

func (AppRoleAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppRoleAssignment)(nil)).Elem()
}

func (o AppRoleAssignmentOutput) ToAppRoleAssignmentOutput() AppRoleAssignmentOutput {
	return o
}

func (o AppRoleAssignmentOutput) ToAppRoleAssignmentOutputWithContext(ctx context.Context) AppRoleAssignmentOutput {
	return o
}

// The ID of the app role to be assigned, or the default role ID `00000000-0000-0000-0000-000000000000`. Changing this forces a new resource to be created.
func (o AppRoleAssignmentOutput) AppRoleId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppRoleAssignment) pulumi.StringOutput { return v.AppRoleId }).(pulumi.StringOutput)
}

// The display name of the principal to which the app role is assigned.
func (o AppRoleAssignmentOutput) PrincipalDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AppRoleAssignment) pulumi.StringOutput { return v.PrincipalDisplayName }).(pulumi.StringOutput)
}

// The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
func (o AppRoleAssignmentOutput) PrincipalObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppRoleAssignment) pulumi.StringOutput { return v.PrincipalObjectId }).(pulumi.StringOutput)
}

// The object type of the principal to which the app role is assigned.
func (o AppRoleAssignmentOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v *AppRoleAssignment) pulumi.StringOutput { return v.PrincipalType }).(pulumi.StringOutput)
}

// The display name of the application representing the resource.
func (o AppRoleAssignmentOutput) ResourceDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AppRoleAssignment) pulumi.StringOutput { return v.ResourceDisplayName }).(pulumi.StringOutput)
}

// The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
func (o AppRoleAssignmentOutput) ResourceObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppRoleAssignment) pulumi.StringOutput { return v.ResourceObjectId }).(pulumi.StringOutput)
}

type AppRoleAssignmentArrayOutput struct{ *pulumi.OutputState }

func (AppRoleAssignmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppRoleAssignment)(nil)).Elem()
}

func (o AppRoleAssignmentArrayOutput) ToAppRoleAssignmentArrayOutput() AppRoleAssignmentArrayOutput {
	return o
}

func (o AppRoleAssignmentArrayOutput) ToAppRoleAssignmentArrayOutputWithContext(ctx context.Context) AppRoleAssignmentArrayOutput {
	return o
}

func (o AppRoleAssignmentArrayOutput) Index(i pulumi.IntInput) AppRoleAssignmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppRoleAssignment {
		return vs[0].([]*AppRoleAssignment)[vs[1].(int)]
	}).(AppRoleAssignmentOutput)
}

type AppRoleAssignmentMapOutput struct{ *pulumi.OutputState }

func (AppRoleAssignmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppRoleAssignment)(nil)).Elem()
}

func (o AppRoleAssignmentMapOutput) ToAppRoleAssignmentMapOutput() AppRoleAssignmentMapOutput {
	return o
}

func (o AppRoleAssignmentMapOutput) ToAppRoleAssignmentMapOutputWithContext(ctx context.Context) AppRoleAssignmentMapOutput {
	return o
}

func (o AppRoleAssignmentMapOutput) MapIndex(k pulumi.StringInput) AppRoleAssignmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppRoleAssignment {
		return vs[0].(map[string]*AppRoleAssignment)[vs[1].(string)]
	}).(AppRoleAssignmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppRoleAssignmentInput)(nil)).Elem(), &AppRoleAssignment{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppRoleAssignmentArrayInput)(nil)).Elem(), AppRoleAssignmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppRoleAssignmentMapInput)(nil)).Elem(), AppRoleAssignmentMap{})
	pulumi.RegisterOutputType(AppRoleAssignmentOutput{})
	pulumi.RegisterOutputType(AppRoleAssignmentArrayOutput{})
	pulumi.RegisterOutputType(AppRoleAssignmentMapOutput{})
}

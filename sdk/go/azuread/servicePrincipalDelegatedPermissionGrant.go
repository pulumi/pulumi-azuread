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

// Manages a delegated permission grant for a service principal, on behalf of a single user, or all users.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires the following application role: `Directory.ReadWrite.All`
//
// When authenticated with a user principal, this resource requires one the following directory role: `Global Administrator`
//
// ## Example Usage
//
// *Delegated permission grant for all users*
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
//			wellKnown, err := azuread.GetApplicationPublishedAppIds(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			msgraph, err := azuread.NewServicePrincipal(ctx, "msgraph", &azuread.ServicePrincipalArgs{
//				ApplicationId: *pulumi.String(wellKnown.Result.MicrosoftGraph),
//				UseExisting:   pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			exampleApplication, err := azuread.NewApplication(ctx, "exampleApplication", &azuread.ApplicationArgs{
//				DisplayName: pulumi.String("example"),
//				RequiredResourceAccesses: azuread.ApplicationRequiredResourceAccessArray{
//					&azuread.ApplicationRequiredResourceAccessArgs{
//						ResourceAppId: *pulumi.String(wellKnown.Result.MicrosoftGraph),
//						ResourceAccesses: azuread.ApplicationRequiredResourceAccessResourceAccessArray{
//							&azuread.ApplicationRequiredResourceAccessResourceAccessArgs{
//								Id: msgraph.Oauth2PermissionScopeIds.ApplyT(func(oauth2PermissionScopeIds map[string]string) (string, error) {
//									return oauth2PermissionScopeIds.Openid, nil
//								}).(pulumi.StringOutput),
//								Type: pulumi.String("Scope"),
//							},
//							&azuread.ApplicationRequiredResourceAccessResourceAccessArgs{
//								Id: msgraph.Oauth2PermissionScopeIds.ApplyT(func(oauth2PermissionScopeIds map[string]string) (string, error) {
//									return oauth2PermissionScopeIds.User.Read, nil
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
//			exampleServicePrincipal, err := azuread.NewServicePrincipal(ctx, "exampleServicePrincipal", &azuread.ServicePrincipalArgs{
//				ApplicationId: exampleApplication.ApplicationId,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewServicePrincipalDelegatedPermissionGrant(ctx, "exampleServicePrincipalDelegatedPermissionGrant", &azuread.ServicePrincipalDelegatedPermissionGrantArgs{
//				ServicePrincipalObjectId:         exampleServicePrincipal.ObjectId,
//				ResourceServicePrincipalObjectId: msgraph.ObjectId,
//				ClaimValues: pulumi.StringArray{
//					pulumi.String("openid"),
//					pulumi.String("User.Read.All"),
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
// *Delegated permission grant for a single user*
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
//			wellKnown, err := azuread.GetApplicationPublishedAppIds(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			msgraph, err := azuread.NewServicePrincipal(ctx, "msgraph", &azuread.ServicePrincipalArgs{
//				ApplicationId: *pulumi.String(wellKnown.Result.MicrosoftGraph),
//				UseExisting:   pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			exampleApplication, err := azuread.NewApplication(ctx, "exampleApplication", &azuread.ApplicationArgs{
//				DisplayName: pulumi.String("example"),
//				RequiredResourceAccesses: azuread.ApplicationRequiredResourceAccessArray{
//					&azuread.ApplicationRequiredResourceAccessArgs{
//						ResourceAppId: *pulumi.String(wellKnown.Result.MicrosoftGraph),
//						ResourceAccesses: azuread.ApplicationRequiredResourceAccessResourceAccessArray{
//							&azuread.ApplicationRequiredResourceAccessResourceAccessArgs{
//								Id: msgraph.Oauth2PermissionScopeIds.ApplyT(func(oauth2PermissionScopeIds map[string]string) (string, error) {
//									return oauth2PermissionScopeIds.Openid, nil
//								}).(pulumi.StringOutput),
//								Type: pulumi.String("Scope"),
//							},
//							&azuread.ApplicationRequiredResourceAccessResourceAccessArgs{
//								Id: msgraph.Oauth2PermissionScopeIds.ApplyT(func(oauth2PermissionScopeIds map[string]string) (string, error) {
//									return oauth2PermissionScopeIds.User.Read, nil
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
//			exampleServicePrincipal, err := azuread.NewServicePrincipal(ctx, "exampleServicePrincipal", &azuread.ServicePrincipalArgs{
//				ApplicationId: exampleApplication.ApplicationId,
//			})
//			if err != nil {
//				return err
//			}
//			exampleUser, err := azuread.NewUser(ctx, "exampleUser", &azuread.UserArgs{
//				DisplayName:       pulumi.String("J. Doe"),
//				UserPrincipalName: pulumi.String("jdoe@hashicorp.com"),
//				MailNickname:      pulumi.String("jdoe"),
//				Password:          pulumi.String("SecretP@sswd99!"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewServicePrincipalDelegatedPermissionGrant(ctx, "exampleServicePrincipalDelegatedPermissionGrant", &azuread.ServicePrincipalDelegatedPermissionGrantArgs{
//				ServicePrincipalObjectId:         exampleServicePrincipal.ObjectId,
//				ResourceServicePrincipalObjectId: msgraph.ObjectId,
//				ClaimValues: pulumi.StringArray{
//					pulumi.String("openid"),
//					pulumi.String("User.Read.All"),
//				},
//				UserObjectId: exampleUser.ObjectId,
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
// Delegated permission grants can be imported using their ID, e.g.
//
// ```sh
//
//	$ pulumi import azuread:index/servicePrincipalDelegatedPermissionGrant:ServicePrincipalDelegatedPermissionGrant example aaBBcDDeFG6h5JKLMN2PQrrssTTUUvWWxxxxxyyyzzz
//
// ```
type ServicePrincipalDelegatedPermissionGrant struct {
	pulumi.CustomResourceState

	// A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
	ClaimValues pulumi.StringArrayOutput `pulumi:"claimValues"`
	// The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
	ResourceServicePrincipalObjectId pulumi.StringOutput `pulumi:"resourceServicePrincipalObjectId"`
	// The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
	ServicePrincipalObjectId pulumi.StringOutput `pulumi:"servicePrincipalObjectId"`
	// The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
	//
	// > **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `userObjectId` property.
	UserObjectId pulumi.StringPtrOutput `pulumi:"userObjectId"`
}

// NewServicePrincipalDelegatedPermissionGrant registers a new resource with the given unique name, arguments, and options.
func NewServicePrincipalDelegatedPermissionGrant(ctx *pulumi.Context,
	name string, args *ServicePrincipalDelegatedPermissionGrantArgs, opts ...pulumi.ResourceOption) (*ServicePrincipalDelegatedPermissionGrant, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClaimValues == nil {
		return nil, errors.New("invalid value for required argument 'ClaimValues'")
	}
	if args.ResourceServicePrincipalObjectId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceServicePrincipalObjectId'")
	}
	if args.ServicePrincipalObjectId == nil {
		return nil, errors.New("invalid value for required argument 'ServicePrincipalObjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServicePrincipalDelegatedPermissionGrant
	err := ctx.RegisterResource("azuread:index/servicePrincipalDelegatedPermissionGrant:ServicePrincipalDelegatedPermissionGrant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicePrincipalDelegatedPermissionGrant gets an existing ServicePrincipalDelegatedPermissionGrant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicePrincipalDelegatedPermissionGrant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicePrincipalDelegatedPermissionGrantState, opts ...pulumi.ResourceOption) (*ServicePrincipalDelegatedPermissionGrant, error) {
	var resource ServicePrincipalDelegatedPermissionGrant
	err := ctx.ReadResource("azuread:index/servicePrincipalDelegatedPermissionGrant:ServicePrincipalDelegatedPermissionGrant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicePrincipalDelegatedPermissionGrant resources.
type servicePrincipalDelegatedPermissionGrantState struct {
	// A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
	ClaimValues []string `pulumi:"claimValues"`
	// The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
	ResourceServicePrincipalObjectId *string `pulumi:"resourceServicePrincipalObjectId"`
	// The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
	ServicePrincipalObjectId *string `pulumi:"servicePrincipalObjectId"`
	// The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
	//
	// > **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `userObjectId` property.
	UserObjectId *string `pulumi:"userObjectId"`
}

type ServicePrincipalDelegatedPermissionGrantState struct {
	// A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
	ClaimValues pulumi.StringArrayInput
	// The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
	ResourceServicePrincipalObjectId pulumi.StringPtrInput
	// The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
	ServicePrincipalObjectId pulumi.StringPtrInput
	// The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
	//
	// > **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `userObjectId` property.
	UserObjectId pulumi.StringPtrInput
}

func (ServicePrincipalDelegatedPermissionGrantState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePrincipalDelegatedPermissionGrantState)(nil)).Elem()
}

type servicePrincipalDelegatedPermissionGrantArgs struct {
	// A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
	ClaimValues []string `pulumi:"claimValues"`
	// The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
	ResourceServicePrincipalObjectId string `pulumi:"resourceServicePrincipalObjectId"`
	// The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
	ServicePrincipalObjectId string `pulumi:"servicePrincipalObjectId"`
	// The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
	//
	// > **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `userObjectId` property.
	UserObjectId *string `pulumi:"userObjectId"`
}

// The set of arguments for constructing a ServicePrincipalDelegatedPermissionGrant resource.
type ServicePrincipalDelegatedPermissionGrantArgs struct {
	// A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
	ClaimValues pulumi.StringArrayInput
	// The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
	ResourceServicePrincipalObjectId pulumi.StringInput
	// The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
	ServicePrincipalObjectId pulumi.StringInput
	// The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
	//
	// > **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `userObjectId` property.
	UserObjectId pulumi.StringPtrInput
}

func (ServicePrincipalDelegatedPermissionGrantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePrincipalDelegatedPermissionGrantArgs)(nil)).Elem()
}

type ServicePrincipalDelegatedPermissionGrantInput interface {
	pulumi.Input

	ToServicePrincipalDelegatedPermissionGrantOutput() ServicePrincipalDelegatedPermissionGrantOutput
	ToServicePrincipalDelegatedPermissionGrantOutputWithContext(ctx context.Context) ServicePrincipalDelegatedPermissionGrantOutput
}

func (*ServicePrincipalDelegatedPermissionGrant) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalDelegatedPermissionGrant)(nil)).Elem()
}

func (i *ServicePrincipalDelegatedPermissionGrant) ToServicePrincipalDelegatedPermissionGrantOutput() ServicePrincipalDelegatedPermissionGrantOutput {
	return i.ToServicePrincipalDelegatedPermissionGrantOutputWithContext(context.Background())
}

func (i *ServicePrincipalDelegatedPermissionGrant) ToServicePrincipalDelegatedPermissionGrantOutputWithContext(ctx context.Context) ServicePrincipalDelegatedPermissionGrantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalDelegatedPermissionGrantOutput)
}

// ServicePrincipalDelegatedPermissionGrantArrayInput is an input type that accepts ServicePrincipalDelegatedPermissionGrantArray and ServicePrincipalDelegatedPermissionGrantArrayOutput values.
// You can construct a concrete instance of `ServicePrincipalDelegatedPermissionGrantArrayInput` via:
//
//	ServicePrincipalDelegatedPermissionGrantArray{ ServicePrincipalDelegatedPermissionGrantArgs{...} }
type ServicePrincipalDelegatedPermissionGrantArrayInput interface {
	pulumi.Input

	ToServicePrincipalDelegatedPermissionGrantArrayOutput() ServicePrincipalDelegatedPermissionGrantArrayOutput
	ToServicePrincipalDelegatedPermissionGrantArrayOutputWithContext(context.Context) ServicePrincipalDelegatedPermissionGrantArrayOutput
}

type ServicePrincipalDelegatedPermissionGrantArray []ServicePrincipalDelegatedPermissionGrantInput

func (ServicePrincipalDelegatedPermissionGrantArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePrincipalDelegatedPermissionGrant)(nil)).Elem()
}

func (i ServicePrincipalDelegatedPermissionGrantArray) ToServicePrincipalDelegatedPermissionGrantArrayOutput() ServicePrincipalDelegatedPermissionGrantArrayOutput {
	return i.ToServicePrincipalDelegatedPermissionGrantArrayOutputWithContext(context.Background())
}

func (i ServicePrincipalDelegatedPermissionGrantArray) ToServicePrincipalDelegatedPermissionGrantArrayOutputWithContext(ctx context.Context) ServicePrincipalDelegatedPermissionGrantArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalDelegatedPermissionGrantArrayOutput)
}

// ServicePrincipalDelegatedPermissionGrantMapInput is an input type that accepts ServicePrincipalDelegatedPermissionGrantMap and ServicePrincipalDelegatedPermissionGrantMapOutput values.
// You can construct a concrete instance of `ServicePrincipalDelegatedPermissionGrantMapInput` via:
//
//	ServicePrincipalDelegatedPermissionGrantMap{ "key": ServicePrincipalDelegatedPermissionGrantArgs{...} }
type ServicePrincipalDelegatedPermissionGrantMapInput interface {
	pulumi.Input

	ToServicePrincipalDelegatedPermissionGrantMapOutput() ServicePrincipalDelegatedPermissionGrantMapOutput
	ToServicePrincipalDelegatedPermissionGrantMapOutputWithContext(context.Context) ServicePrincipalDelegatedPermissionGrantMapOutput
}

type ServicePrincipalDelegatedPermissionGrantMap map[string]ServicePrincipalDelegatedPermissionGrantInput

func (ServicePrincipalDelegatedPermissionGrantMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePrincipalDelegatedPermissionGrant)(nil)).Elem()
}

func (i ServicePrincipalDelegatedPermissionGrantMap) ToServicePrincipalDelegatedPermissionGrantMapOutput() ServicePrincipalDelegatedPermissionGrantMapOutput {
	return i.ToServicePrincipalDelegatedPermissionGrantMapOutputWithContext(context.Background())
}

func (i ServicePrincipalDelegatedPermissionGrantMap) ToServicePrincipalDelegatedPermissionGrantMapOutputWithContext(ctx context.Context) ServicePrincipalDelegatedPermissionGrantMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalDelegatedPermissionGrantMapOutput)
}

type ServicePrincipalDelegatedPermissionGrantOutput struct{ *pulumi.OutputState }

func (ServicePrincipalDelegatedPermissionGrantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalDelegatedPermissionGrant)(nil)).Elem()
}

func (o ServicePrincipalDelegatedPermissionGrantOutput) ToServicePrincipalDelegatedPermissionGrantOutput() ServicePrincipalDelegatedPermissionGrantOutput {
	return o
}

func (o ServicePrincipalDelegatedPermissionGrantOutput) ToServicePrincipalDelegatedPermissionGrantOutputWithContext(ctx context.Context) ServicePrincipalDelegatedPermissionGrantOutput {
	return o
}

// A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
func (o ServicePrincipalDelegatedPermissionGrantOutput) ClaimValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServicePrincipalDelegatedPermissionGrant) pulumi.StringArrayOutput { return v.ClaimValues }).(pulumi.StringArrayOutput)
}

// The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
func (o ServicePrincipalDelegatedPermissionGrantOutput) ResourceServicePrincipalObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalDelegatedPermissionGrant) pulumi.StringOutput {
		return v.ResourceServicePrincipalObjectId
	}).(pulumi.StringOutput)
}

// The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
func (o ServicePrincipalDelegatedPermissionGrantOutput) ServicePrincipalObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalDelegatedPermissionGrant) pulumi.StringOutput {
		return v.ServicePrincipalObjectId
	}).(pulumi.StringOutput)
}

// The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
//
// > **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `userObjectId` property.
func (o ServicePrincipalDelegatedPermissionGrantOutput) UserObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalDelegatedPermissionGrant) pulumi.StringPtrOutput { return v.UserObjectId }).(pulumi.StringPtrOutput)
}

type ServicePrincipalDelegatedPermissionGrantArrayOutput struct{ *pulumi.OutputState }

func (ServicePrincipalDelegatedPermissionGrantArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePrincipalDelegatedPermissionGrant)(nil)).Elem()
}

func (o ServicePrincipalDelegatedPermissionGrantArrayOutput) ToServicePrincipalDelegatedPermissionGrantArrayOutput() ServicePrincipalDelegatedPermissionGrantArrayOutput {
	return o
}

func (o ServicePrincipalDelegatedPermissionGrantArrayOutput) ToServicePrincipalDelegatedPermissionGrantArrayOutputWithContext(ctx context.Context) ServicePrincipalDelegatedPermissionGrantArrayOutput {
	return o
}

func (o ServicePrincipalDelegatedPermissionGrantArrayOutput) Index(i pulumi.IntInput) ServicePrincipalDelegatedPermissionGrantOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServicePrincipalDelegatedPermissionGrant {
		return vs[0].([]*ServicePrincipalDelegatedPermissionGrant)[vs[1].(int)]
	}).(ServicePrincipalDelegatedPermissionGrantOutput)
}

type ServicePrincipalDelegatedPermissionGrantMapOutput struct{ *pulumi.OutputState }

func (ServicePrincipalDelegatedPermissionGrantMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePrincipalDelegatedPermissionGrant)(nil)).Elem()
}

func (o ServicePrincipalDelegatedPermissionGrantMapOutput) ToServicePrincipalDelegatedPermissionGrantMapOutput() ServicePrincipalDelegatedPermissionGrantMapOutput {
	return o
}

func (o ServicePrincipalDelegatedPermissionGrantMapOutput) ToServicePrincipalDelegatedPermissionGrantMapOutputWithContext(ctx context.Context) ServicePrincipalDelegatedPermissionGrantMapOutput {
	return o
}

func (o ServicePrincipalDelegatedPermissionGrantMapOutput) MapIndex(k pulumi.StringInput) ServicePrincipalDelegatedPermissionGrantOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServicePrincipalDelegatedPermissionGrant {
		return vs[0].(map[string]*ServicePrincipalDelegatedPermissionGrant)[vs[1].(string)]
	}).(ServicePrincipalDelegatedPermissionGrantOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalDelegatedPermissionGrantInput)(nil)).Elem(), &ServicePrincipalDelegatedPermissionGrant{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalDelegatedPermissionGrantArrayInput)(nil)).Elem(), ServicePrincipalDelegatedPermissionGrantArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalDelegatedPermissionGrantMapInput)(nil)).Elem(), ServicePrincipalDelegatedPermissionGrantMap{})
	pulumi.RegisterOutputType(ServicePrincipalDelegatedPermissionGrantOutput{})
	pulumi.RegisterOutputType(ServicePrincipalDelegatedPermissionGrantArrayOutput{})
	pulumi.RegisterOutputType(ServicePrincipalDelegatedPermissionGrantMapOutput{})
}

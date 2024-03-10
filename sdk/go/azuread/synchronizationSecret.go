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

// Manages synchronization secrets associated with a service principal (enterprise application) within Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.All` or `Directory.ReadWrite.All`
//
// ## Example Usage
//
// *Basic example*
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
//			example, err := azuread.GetApplicationTemplate(ctx, &azuread.GetApplicationTemplateArgs{
//				DisplayName: pulumi.StringRef("Azure Databricks SCIM Provisioning Connector"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleApplication, err := azuread.NewApplication(ctx, "example", &azuread.ApplicationArgs{
//				DisplayName: pulumi.String("example"),
//				TemplateId:  pulumi.String(example.TemplateId),
//				FeatureTags: azuread.ApplicationFeatureTagArray{
//					&azuread.ApplicationFeatureTagArgs{
//						Enterprise: pulumi.Bool(true),
//						Gallery:    pulumi.Bool(true),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleServicePrincipal, err := azuread.NewServicePrincipal(ctx, "example", &azuread.ServicePrincipalArgs{
//				ApplicationId: exampleApplication.ApplicationId,
//				UseExisting:   pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewSynchronizationSecret(ctx, "example", &azuread.SynchronizationSecretArgs{
//				ServicePrincipalId: exampleServicePrincipal.ID(),
//				Credentials: azuread.SynchronizationSecretCredentialArray{
//					&azuread.SynchronizationSecretCredentialArgs{
//						Key:   pulumi.String("BaseAddress"),
//						Value: pulumi.String("abc"),
//					},
//					&azuread.SynchronizationSecretCredentialArgs{
//						Key:   pulumi.String("SecretToken"),
//						Value: pulumi.String("some-token"),
//					},
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
// This resource does not support importing.
type SynchronizationSecret struct {
	pulumi.CustomResourceState

	// One or more `credential` blocks as documented below.
	Credentials SynchronizationSecretCredentialArrayOutput `pulumi:"credentials"`
	// The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
	ServicePrincipalId pulumi.StringOutput `pulumi:"servicePrincipalId"`
}

// NewSynchronizationSecret registers a new resource with the given unique name, arguments, and options.
func NewSynchronizationSecret(ctx *pulumi.Context,
	name string, args *SynchronizationSecretArgs, opts ...pulumi.ResourceOption) (*SynchronizationSecret, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServicePrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'ServicePrincipalId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SynchronizationSecret
	err := ctx.RegisterResource("azuread:index/synchronizationSecret:SynchronizationSecret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSynchronizationSecret gets an existing SynchronizationSecret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSynchronizationSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SynchronizationSecretState, opts ...pulumi.ResourceOption) (*SynchronizationSecret, error) {
	var resource SynchronizationSecret
	err := ctx.ReadResource("azuread:index/synchronizationSecret:SynchronizationSecret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SynchronizationSecret resources.
type synchronizationSecretState struct {
	// One or more `credential` blocks as documented below.
	Credentials []SynchronizationSecretCredential `pulumi:"credentials"`
	// The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
	ServicePrincipalId *string `pulumi:"servicePrincipalId"`
}

type SynchronizationSecretState struct {
	// One or more `credential` blocks as documented below.
	Credentials SynchronizationSecretCredentialArrayInput
	// The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
	ServicePrincipalId pulumi.StringPtrInput
}

func (SynchronizationSecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*synchronizationSecretState)(nil)).Elem()
}

type synchronizationSecretArgs struct {
	// One or more `credential` blocks as documented below.
	Credentials []SynchronizationSecretCredential `pulumi:"credentials"`
	// The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
	ServicePrincipalId string `pulumi:"servicePrincipalId"`
}

// The set of arguments for constructing a SynchronizationSecret resource.
type SynchronizationSecretArgs struct {
	// One or more `credential` blocks as documented below.
	Credentials SynchronizationSecretCredentialArrayInput
	// The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
	ServicePrincipalId pulumi.StringInput
}

func (SynchronizationSecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*synchronizationSecretArgs)(nil)).Elem()
}

type SynchronizationSecretInput interface {
	pulumi.Input

	ToSynchronizationSecretOutput() SynchronizationSecretOutput
	ToSynchronizationSecretOutputWithContext(ctx context.Context) SynchronizationSecretOutput
}

func (*SynchronizationSecret) ElementType() reflect.Type {
	return reflect.TypeOf((**SynchronizationSecret)(nil)).Elem()
}

func (i *SynchronizationSecret) ToSynchronizationSecretOutput() SynchronizationSecretOutput {
	return i.ToSynchronizationSecretOutputWithContext(context.Background())
}

func (i *SynchronizationSecret) ToSynchronizationSecretOutputWithContext(ctx context.Context) SynchronizationSecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynchronizationSecretOutput)
}

// SynchronizationSecretArrayInput is an input type that accepts SynchronizationSecretArray and SynchronizationSecretArrayOutput values.
// You can construct a concrete instance of `SynchronizationSecretArrayInput` via:
//
//	SynchronizationSecretArray{ SynchronizationSecretArgs{...} }
type SynchronizationSecretArrayInput interface {
	pulumi.Input

	ToSynchronizationSecretArrayOutput() SynchronizationSecretArrayOutput
	ToSynchronizationSecretArrayOutputWithContext(context.Context) SynchronizationSecretArrayOutput
}

type SynchronizationSecretArray []SynchronizationSecretInput

func (SynchronizationSecretArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SynchronizationSecret)(nil)).Elem()
}

func (i SynchronizationSecretArray) ToSynchronizationSecretArrayOutput() SynchronizationSecretArrayOutput {
	return i.ToSynchronizationSecretArrayOutputWithContext(context.Background())
}

func (i SynchronizationSecretArray) ToSynchronizationSecretArrayOutputWithContext(ctx context.Context) SynchronizationSecretArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynchronizationSecretArrayOutput)
}

// SynchronizationSecretMapInput is an input type that accepts SynchronizationSecretMap and SynchronizationSecretMapOutput values.
// You can construct a concrete instance of `SynchronizationSecretMapInput` via:
//
//	SynchronizationSecretMap{ "key": SynchronizationSecretArgs{...} }
type SynchronizationSecretMapInput interface {
	pulumi.Input

	ToSynchronizationSecretMapOutput() SynchronizationSecretMapOutput
	ToSynchronizationSecretMapOutputWithContext(context.Context) SynchronizationSecretMapOutput
}

type SynchronizationSecretMap map[string]SynchronizationSecretInput

func (SynchronizationSecretMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SynchronizationSecret)(nil)).Elem()
}

func (i SynchronizationSecretMap) ToSynchronizationSecretMapOutput() SynchronizationSecretMapOutput {
	return i.ToSynchronizationSecretMapOutputWithContext(context.Background())
}

func (i SynchronizationSecretMap) ToSynchronizationSecretMapOutputWithContext(ctx context.Context) SynchronizationSecretMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynchronizationSecretMapOutput)
}

type SynchronizationSecretOutput struct{ *pulumi.OutputState }

func (SynchronizationSecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SynchronizationSecret)(nil)).Elem()
}

func (o SynchronizationSecretOutput) ToSynchronizationSecretOutput() SynchronizationSecretOutput {
	return o
}

func (o SynchronizationSecretOutput) ToSynchronizationSecretOutputWithContext(ctx context.Context) SynchronizationSecretOutput {
	return o
}

// One or more `credential` blocks as documented below.
func (o SynchronizationSecretOutput) Credentials() SynchronizationSecretCredentialArrayOutput {
	return o.ApplyT(func(v *SynchronizationSecret) SynchronizationSecretCredentialArrayOutput { return v.Credentials }).(SynchronizationSecretCredentialArrayOutput)
}

// The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
func (o SynchronizationSecretOutput) ServicePrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *SynchronizationSecret) pulumi.StringOutput { return v.ServicePrincipalId }).(pulumi.StringOutput)
}

type SynchronizationSecretArrayOutput struct{ *pulumi.OutputState }

func (SynchronizationSecretArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SynchronizationSecret)(nil)).Elem()
}

func (o SynchronizationSecretArrayOutput) ToSynchronizationSecretArrayOutput() SynchronizationSecretArrayOutput {
	return o
}

func (o SynchronizationSecretArrayOutput) ToSynchronizationSecretArrayOutputWithContext(ctx context.Context) SynchronizationSecretArrayOutput {
	return o
}

func (o SynchronizationSecretArrayOutput) Index(i pulumi.IntInput) SynchronizationSecretOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SynchronizationSecret {
		return vs[0].([]*SynchronizationSecret)[vs[1].(int)]
	}).(SynchronizationSecretOutput)
}

type SynchronizationSecretMapOutput struct{ *pulumi.OutputState }

func (SynchronizationSecretMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SynchronizationSecret)(nil)).Elem()
}

func (o SynchronizationSecretMapOutput) ToSynchronizationSecretMapOutput() SynchronizationSecretMapOutput {
	return o
}

func (o SynchronizationSecretMapOutput) ToSynchronizationSecretMapOutputWithContext(ctx context.Context) SynchronizationSecretMapOutput {
	return o
}

func (o SynchronizationSecretMapOutput) MapIndex(k pulumi.StringInput) SynchronizationSecretOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SynchronizationSecret {
		return vs[0].(map[string]*SynchronizationSecret)[vs[1].(string)]
	}).(SynchronizationSecretOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SynchronizationSecretInput)(nil)).Elem(), &SynchronizationSecret{})
	pulumi.RegisterInputType(reflect.TypeOf((*SynchronizationSecretArrayInput)(nil)).Elem(), SynchronizationSecretArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SynchronizationSecretMapInput)(nil)).Elem(), SynchronizationSecretMap{})
	pulumi.RegisterOutputType(SynchronizationSecretOutput{})
	pulumi.RegisterOutputType(SynchronizationSecretArrayOutput{})
	pulumi.RegisterOutputType(SynchronizationSecretMapOutput{})
}

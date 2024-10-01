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

// Manages synchronization job on demand provisioning associated with a service principal (enterprise application) within Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires one of the following application roles: `Synchronization.ReadWrite.All`
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
//			current, err := azuread.GetClientConfig(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			exampleGroup, err := azuread.NewGroup(ctx, "example", &azuread.GroupArgs{
//				DisplayName: pulumi.String("example"),
//				Owners: pulumi.StringArray{
//					pulumi.String(current.ObjectId),
//				},
//				SecurityEnabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
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
//				ClientId:    exampleApplication.ClientId,
//				UseExisting: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewSynchronizationSecret(ctx, "example", &azuread.SynchronizationSecretArgs{
//				ServicePrincipalId: exampleServicePrincipal.ID(),
//				Credentials: azuread.SynchronizationSecretCredentialArray{
//					&azuread.SynchronizationSecretCredentialArgs{
//						Key:   pulumi.String("BaseAddress"),
//						Value: pulumi.String("https://adb-example.azuredatabricks.net/api/2.0/preview/scim"),
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
//			exampleSynchronizationJob, err := azuread.NewSynchronizationJob(ctx, "example", &azuread.SynchronizationJobArgs{
//				ServicePrincipalId: exampleServicePrincipal.ID(),
//				TemplateId:         pulumi.String("dataBricks"),
//				Enabled:            pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewSynchronizationJobProvisionOnDemand(ctx, "example", &azuread.SynchronizationJobProvisionOnDemandArgs{
//				ServicePrincipalId:   exampleServicePrincipal.ID(),
//				SynchronizationJobId: exampleSynchronizationJob.ID(),
//				Parameters: azuread.SynchronizationJobProvisionOnDemandParameterArray{
//					&azuread.SynchronizationJobProvisionOnDemandParameterArgs{
//						RuleId: pulumi.String(""),
//						Subjects: azuread.SynchronizationJobProvisionOnDemandParameterSubjectArray{
//							&azuread.SynchronizationJobProvisionOnDemandParameterSubjectArgs{
//								ObjectId:       exampleGroup.ObjectId,
//								ObjectTypeName: pulumi.String("Group"),
//							},
//						},
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
type SynchronizationJobProvisionOnDemand struct {
	pulumi.CustomResourceState

	// One or more `parameter` blocks as documented below.
	Parameters SynchronizationJobProvisionOnDemandParameterArrayOutput `pulumi:"parameters"`
	// The object ID of the service principal for the synchronization job.
	ServicePrincipalId pulumi.StringOutput `pulumi:"servicePrincipalId"`
	// Identifier of the synchronization template this job is based on.
	SynchronizationJobId pulumi.StringOutput    `pulumi:"synchronizationJobId"`
	Triggers             pulumi.StringMapOutput `pulumi:"triggers"`
}

// NewSynchronizationJobProvisionOnDemand registers a new resource with the given unique name, arguments, and options.
func NewSynchronizationJobProvisionOnDemand(ctx *pulumi.Context,
	name string, args *SynchronizationJobProvisionOnDemandArgs, opts ...pulumi.ResourceOption) (*SynchronizationJobProvisionOnDemand, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Parameters == nil {
		return nil, errors.New("invalid value for required argument 'Parameters'")
	}
	if args.ServicePrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'ServicePrincipalId'")
	}
	if args.SynchronizationJobId == nil {
		return nil, errors.New("invalid value for required argument 'SynchronizationJobId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SynchronizationJobProvisionOnDemand
	err := ctx.RegisterResource("azuread:index/synchronizationJobProvisionOnDemand:SynchronizationJobProvisionOnDemand", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSynchronizationJobProvisionOnDemand gets an existing SynchronizationJobProvisionOnDemand resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSynchronizationJobProvisionOnDemand(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SynchronizationJobProvisionOnDemandState, opts ...pulumi.ResourceOption) (*SynchronizationJobProvisionOnDemand, error) {
	var resource SynchronizationJobProvisionOnDemand
	err := ctx.ReadResource("azuread:index/synchronizationJobProvisionOnDemand:SynchronizationJobProvisionOnDemand", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SynchronizationJobProvisionOnDemand resources.
type synchronizationJobProvisionOnDemandState struct {
	// One or more `parameter` blocks as documented below.
	Parameters []SynchronizationJobProvisionOnDemandParameter `pulumi:"parameters"`
	// The object ID of the service principal for the synchronization job.
	ServicePrincipalId *string `pulumi:"servicePrincipalId"`
	// Identifier of the synchronization template this job is based on.
	SynchronizationJobId *string           `pulumi:"synchronizationJobId"`
	Triggers             map[string]string `pulumi:"triggers"`
}

type SynchronizationJobProvisionOnDemandState struct {
	// One or more `parameter` blocks as documented below.
	Parameters SynchronizationJobProvisionOnDemandParameterArrayInput
	// The object ID of the service principal for the synchronization job.
	ServicePrincipalId pulumi.StringPtrInput
	// Identifier of the synchronization template this job is based on.
	SynchronizationJobId pulumi.StringPtrInput
	Triggers             pulumi.StringMapInput
}

func (SynchronizationJobProvisionOnDemandState) ElementType() reflect.Type {
	return reflect.TypeOf((*synchronizationJobProvisionOnDemandState)(nil)).Elem()
}

type synchronizationJobProvisionOnDemandArgs struct {
	// One or more `parameter` blocks as documented below.
	Parameters []SynchronizationJobProvisionOnDemandParameter `pulumi:"parameters"`
	// The object ID of the service principal for the synchronization job.
	ServicePrincipalId string `pulumi:"servicePrincipalId"`
	// Identifier of the synchronization template this job is based on.
	SynchronizationJobId string            `pulumi:"synchronizationJobId"`
	Triggers             map[string]string `pulumi:"triggers"`
}

// The set of arguments for constructing a SynchronizationJobProvisionOnDemand resource.
type SynchronizationJobProvisionOnDemandArgs struct {
	// One or more `parameter` blocks as documented below.
	Parameters SynchronizationJobProvisionOnDemandParameterArrayInput
	// The object ID of the service principal for the synchronization job.
	ServicePrincipalId pulumi.StringInput
	// Identifier of the synchronization template this job is based on.
	SynchronizationJobId pulumi.StringInput
	Triggers             pulumi.StringMapInput
}

func (SynchronizationJobProvisionOnDemandArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*synchronizationJobProvisionOnDemandArgs)(nil)).Elem()
}

type SynchronizationJobProvisionOnDemandInput interface {
	pulumi.Input

	ToSynchronizationJobProvisionOnDemandOutput() SynchronizationJobProvisionOnDemandOutput
	ToSynchronizationJobProvisionOnDemandOutputWithContext(ctx context.Context) SynchronizationJobProvisionOnDemandOutput
}

func (*SynchronizationJobProvisionOnDemand) ElementType() reflect.Type {
	return reflect.TypeOf((**SynchronizationJobProvisionOnDemand)(nil)).Elem()
}

func (i *SynchronizationJobProvisionOnDemand) ToSynchronizationJobProvisionOnDemandOutput() SynchronizationJobProvisionOnDemandOutput {
	return i.ToSynchronizationJobProvisionOnDemandOutputWithContext(context.Background())
}

func (i *SynchronizationJobProvisionOnDemand) ToSynchronizationJobProvisionOnDemandOutputWithContext(ctx context.Context) SynchronizationJobProvisionOnDemandOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynchronizationJobProvisionOnDemandOutput)
}

// SynchronizationJobProvisionOnDemandArrayInput is an input type that accepts SynchronizationJobProvisionOnDemandArray and SynchronizationJobProvisionOnDemandArrayOutput values.
// You can construct a concrete instance of `SynchronizationJobProvisionOnDemandArrayInput` via:
//
//	SynchronizationJobProvisionOnDemandArray{ SynchronizationJobProvisionOnDemandArgs{...} }
type SynchronizationJobProvisionOnDemandArrayInput interface {
	pulumi.Input

	ToSynchronizationJobProvisionOnDemandArrayOutput() SynchronizationJobProvisionOnDemandArrayOutput
	ToSynchronizationJobProvisionOnDemandArrayOutputWithContext(context.Context) SynchronizationJobProvisionOnDemandArrayOutput
}

type SynchronizationJobProvisionOnDemandArray []SynchronizationJobProvisionOnDemandInput

func (SynchronizationJobProvisionOnDemandArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SynchronizationJobProvisionOnDemand)(nil)).Elem()
}

func (i SynchronizationJobProvisionOnDemandArray) ToSynchronizationJobProvisionOnDemandArrayOutput() SynchronizationJobProvisionOnDemandArrayOutput {
	return i.ToSynchronizationJobProvisionOnDemandArrayOutputWithContext(context.Background())
}

func (i SynchronizationJobProvisionOnDemandArray) ToSynchronizationJobProvisionOnDemandArrayOutputWithContext(ctx context.Context) SynchronizationJobProvisionOnDemandArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynchronizationJobProvisionOnDemandArrayOutput)
}

// SynchronizationJobProvisionOnDemandMapInput is an input type that accepts SynchronizationJobProvisionOnDemandMap and SynchronizationJobProvisionOnDemandMapOutput values.
// You can construct a concrete instance of `SynchronizationJobProvisionOnDemandMapInput` via:
//
//	SynchronizationJobProvisionOnDemandMap{ "key": SynchronizationJobProvisionOnDemandArgs{...} }
type SynchronizationJobProvisionOnDemandMapInput interface {
	pulumi.Input

	ToSynchronizationJobProvisionOnDemandMapOutput() SynchronizationJobProvisionOnDemandMapOutput
	ToSynchronizationJobProvisionOnDemandMapOutputWithContext(context.Context) SynchronizationJobProvisionOnDemandMapOutput
}

type SynchronizationJobProvisionOnDemandMap map[string]SynchronizationJobProvisionOnDemandInput

func (SynchronizationJobProvisionOnDemandMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SynchronizationJobProvisionOnDemand)(nil)).Elem()
}

func (i SynchronizationJobProvisionOnDemandMap) ToSynchronizationJobProvisionOnDemandMapOutput() SynchronizationJobProvisionOnDemandMapOutput {
	return i.ToSynchronizationJobProvisionOnDemandMapOutputWithContext(context.Background())
}

func (i SynchronizationJobProvisionOnDemandMap) ToSynchronizationJobProvisionOnDemandMapOutputWithContext(ctx context.Context) SynchronizationJobProvisionOnDemandMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynchronizationJobProvisionOnDemandMapOutput)
}

type SynchronizationJobProvisionOnDemandOutput struct{ *pulumi.OutputState }

func (SynchronizationJobProvisionOnDemandOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SynchronizationJobProvisionOnDemand)(nil)).Elem()
}

func (o SynchronizationJobProvisionOnDemandOutput) ToSynchronizationJobProvisionOnDemandOutput() SynchronizationJobProvisionOnDemandOutput {
	return o
}

func (o SynchronizationJobProvisionOnDemandOutput) ToSynchronizationJobProvisionOnDemandOutputWithContext(ctx context.Context) SynchronizationJobProvisionOnDemandOutput {
	return o
}

// One or more `parameter` blocks as documented below.
func (o SynchronizationJobProvisionOnDemandOutput) Parameters() SynchronizationJobProvisionOnDemandParameterArrayOutput {
	return o.ApplyT(func(v *SynchronizationJobProvisionOnDemand) SynchronizationJobProvisionOnDemandParameterArrayOutput {
		return v.Parameters
	}).(SynchronizationJobProvisionOnDemandParameterArrayOutput)
}

// The object ID of the service principal for the synchronization job.
func (o SynchronizationJobProvisionOnDemandOutput) ServicePrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *SynchronizationJobProvisionOnDemand) pulumi.StringOutput { return v.ServicePrincipalId }).(pulumi.StringOutput)
}

// Identifier of the synchronization template this job is based on.
func (o SynchronizationJobProvisionOnDemandOutput) SynchronizationJobId() pulumi.StringOutput {
	return o.ApplyT(func(v *SynchronizationJobProvisionOnDemand) pulumi.StringOutput { return v.SynchronizationJobId }).(pulumi.StringOutput)
}

func (o SynchronizationJobProvisionOnDemandOutput) Triggers() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SynchronizationJobProvisionOnDemand) pulumi.StringMapOutput { return v.Triggers }).(pulumi.StringMapOutput)
}

type SynchronizationJobProvisionOnDemandArrayOutput struct{ *pulumi.OutputState }

func (SynchronizationJobProvisionOnDemandArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SynchronizationJobProvisionOnDemand)(nil)).Elem()
}

func (o SynchronizationJobProvisionOnDemandArrayOutput) ToSynchronizationJobProvisionOnDemandArrayOutput() SynchronizationJobProvisionOnDemandArrayOutput {
	return o
}

func (o SynchronizationJobProvisionOnDemandArrayOutput) ToSynchronizationJobProvisionOnDemandArrayOutputWithContext(ctx context.Context) SynchronizationJobProvisionOnDemandArrayOutput {
	return o
}

func (o SynchronizationJobProvisionOnDemandArrayOutput) Index(i pulumi.IntInput) SynchronizationJobProvisionOnDemandOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SynchronizationJobProvisionOnDemand {
		return vs[0].([]*SynchronizationJobProvisionOnDemand)[vs[1].(int)]
	}).(SynchronizationJobProvisionOnDemandOutput)
}

type SynchronizationJobProvisionOnDemandMapOutput struct{ *pulumi.OutputState }

func (SynchronizationJobProvisionOnDemandMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SynchronizationJobProvisionOnDemand)(nil)).Elem()
}

func (o SynchronizationJobProvisionOnDemandMapOutput) ToSynchronizationJobProvisionOnDemandMapOutput() SynchronizationJobProvisionOnDemandMapOutput {
	return o
}

func (o SynchronizationJobProvisionOnDemandMapOutput) ToSynchronizationJobProvisionOnDemandMapOutputWithContext(ctx context.Context) SynchronizationJobProvisionOnDemandMapOutput {
	return o
}

func (o SynchronizationJobProvisionOnDemandMapOutput) MapIndex(k pulumi.StringInput) SynchronizationJobProvisionOnDemandOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SynchronizationJobProvisionOnDemand {
		return vs[0].(map[string]*SynchronizationJobProvisionOnDemand)[vs[1].(string)]
	}).(SynchronizationJobProvisionOnDemandOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SynchronizationJobProvisionOnDemandInput)(nil)).Elem(), &SynchronizationJobProvisionOnDemand{})
	pulumi.RegisterInputType(reflect.TypeOf((*SynchronizationJobProvisionOnDemandArrayInput)(nil)).Elem(), SynchronizationJobProvisionOnDemandArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SynchronizationJobProvisionOnDemandMapInput)(nil)).Elem(), SynchronizationJobProvisionOnDemandMap{})
	pulumi.RegisterOutputType(SynchronizationJobProvisionOnDemandOutput{})
	pulumi.RegisterOutputType(SynchronizationJobProvisionOnDemandArrayOutput{})
	pulumi.RegisterOutputType(SynchronizationJobProvisionOnDemandMapOutput{})
}

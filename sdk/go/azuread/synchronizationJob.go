// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages a synchronization job associated with a service principal (enterprise application) within Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.All` or `Directory.ReadWrite.All`
//
// ## Import
//
// Synchronization jobs can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import azuread:index/synchronizationJob:SynchronizationJob example 00000000-0000-0000-0000-000000000000/job/dataBricks.f5532fc709734b1a90e8a1fa9fd03a82.8442fd39-2183-419c-8732-74b6ce866bd5
//
// ```
//
//	-> This ID format is unique to Terraform and is composed of the Service Principal Object ID and the ID of the Synchronization Job Id in the format `{servicePrincipalId}/job/{jobId}`.
type SynchronizationJob struct {
	pulumi.CustomResourceState

	// Whether or not the provisioning job is enabled. Default state is `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// A `schedule` list as documented below.
	Schedules SynchronizationJobScheduleArrayOutput `pulumi:"schedules"`
	// The object ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId pulumi.StringOutput `pulumi:"servicePrincipalId"`
	// Identifier of the synchronization template this job is based on.
	TemplateId pulumi.StringOutput `pulumi:"templateId"`
}

// NewSynchronizationJob registers a new resource with the given unique name, arguments, and options.
func NewSynchronizationJob(ctx *pulumi.Context,
	name string, args *SynchronizationJobArgs, opts ...pulumi.ResourceOption) (*SynchronizationJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServicePrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'ServicePrincipalId'")
	}
	if args.TemplateId == nil {
		return nil, errors.New("invalid value for required argument 'TemplateId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SynchronizationJob
	err := ctx.RegisterResource("azuread:index/synchronizationJob:SynchronizationJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSynchronizationJob gets an existing SynchronizationJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSynchronizationJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SynchronizationJobState, opts ...pulumi.ResourceOption) (*SynchronizationJob, error) {
	var resource SynchronizationJob
	err := ctx.ReadResource("azuread:index/synchronizationJob:SynchronizationJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SynchronizationJob resources.
type synchronizationJobState struct {
	// Whether or not the provisioning job is enabled. Default state is `true`.
	Enabled *bool `pulumi:"enabled"`
	// A `schedule` list as documented below.
	Schedules []SynchronizationJobSchedule `pulumi:"schedules"`
	// The object ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId *string `pulumi:"servicePrincipalId"`
	// Identifier of the synchronization template this job is based on.
	TemplateId *string `pulumi:"templateId"`
}

type SynchronizationJobState struct {
	// Whether or not the provisioning job is enabled. Default state is `true`.
	Enabled pulumi.BoolPtrInput
	// A `schedule` list as documented below.
	Schedules SynchronizationJobScheduleArrayInput
	// The object ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId pulumi.StringPtrInput
	// Identifier of the synchronization template this job is based on.
	TemplateId pulumi.StringPtrInput
}

func (SynchronizationJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*synchronizationJobState)(nil)).Elem()
}

type synchronizationJobArgs struct {
	// Whether or not the provisioning job is enabled. Default state is `true`.
	Enabled *bool `pulumi:"enabled"`
	// The object ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId string `pulumi:"servicePrincipalId"`
	// Identifier of the synchronization template this job is based on.
	TemplateId string `pulumi:"templateId"`
}

// The set of arguments for constructing a SynchronizationJob resource.
type SynchronizationJobArgs struct {
	// Whether or not the provisioning job is enabled. Default state is `true`.
	Enabled pulumi.BoolPtrInput
	// The object ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId pulumi.StringInput
	// Identifier of the synchronization template this job is based on.
	TemplateId pulumi.StringInput
}

func (SynchronizationJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*synchronizationJobArgs)(nil)).Elem()
}

type SynchronizationJobInput interface {
	pulumi.Input

	ToSynchronizationJobOutput() SynchronizationJobOutput
	ToSynchronizationJobOutputWithContext(ctx context.Context) SynchronizationJobOutput
}

func (*SynchronizationJob) ElementType() reflect.Type {
	return reflect.TypeOf((**SynchronizationJob)(nil)).Elem()
}

func (i *SynchronizationJob) ToSynchronizationJobOutput() SynchronizationJobOutput {
	return i.ToSynchronizationJobOutputWithContext(context.Background())
}

func (i *SynchronizationJob) ToSynchronizationJobOutputWithContext(ctx context.Context) SynchronizationJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynchronizationJobOutput)
}

func (i *SynchronizationJob) ToOutput(ctx context.Context) pulumix.Output[*SynchronizationJob] {
	return pulumix.Output[*SynchronizationJob]{
		OutputState: i.ToSynchronizationJobOutputWithContext(ctx).OutputState,
	}
}

// SynchronizationJobArrayInput is an input type that accepts SynchronizationJobArray and SynchronizationJobArrayOutput values.
// You can construct a concrete instance of `SynchronizationJobArrayInput` via:
//
//	SynchronizationJobArray{ SynchronizationJobArgs{...} }
type SynchronizationJobArrayInput interface {
	pulumi.Input

	ToSynchronizationJobArrayOutput() SynchronizationJobArrayOutput
	ToSynchronizationJobArrayOutputWithContext(context.Context) SynchronizationJobArrayOutput
}

type SynchronizationJobArray []SynchronizationJobInput

func (SynchronizationJobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SynchronizationJob)(nil)).Elem()
}

func (i SynchronizationJobArray) ToSynchronizationJobArrayOutput() SynchronizationJobArrayOutput {
	return i.ToSynchronizationJobArrayOutputWithContext(context.Background())
}

func (i SynchronizationJobArray) ToSynchronizationJobArrayOutputWithContext(ctx context.Context) SynchronizationJobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynchronizationJobArrayOutput)
}

func (i SynchronizationJobArray) ToOutput(ctx context.Context) pulumix.Output[[]*SynchronizationJob] {
	return pulumix.Output[[]*SynchronizationJob]{
		OutputState: i.ToSynchronizationJobArrayOutputWithContext(ctx).OutputState,
	}
}

// SynchronizationJobMapInput is an input type that accepts SynchronizationJobMap and SynchronizationJobMapOutput values.
// You can construct a concrete instance of `SynchronizationJobMapInput` via:
//
//	SynchronizationJobMap{ "key": SynchronizationJobArgs{...} }
type SynchronizationJobMapInput interface {
	pulumi.Input

	ToSynchronizationJobMapOutput() SynchronizationJobMapOutput
	ToSynchronizationJobMapOutputWithContext(context.Context) SynchronizationJobMapOutput
}

type SynchronizationJobMap map[string]SynchronizationJobInput

func (SynchronizationJobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SynchronizationJob)(nil)).Elem()
}

func (i SynchronizationJobMap) ToSynchronizationJobMapOutput() SynchronizationJobMapOutput {
	return i.ToSynchronizationJobMapOutputWithContext(context.Background())
}

func (i SynchronizationJobMap) ToSynchronizationJobMapOutputWithContext(ctx context.Context) SynchronizationJobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynchronizationJobMapOutput)
}

func (i SynchronizationJobMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SynchronizationJob] {
	return pulumix.Output[map[string]*SynchronizationJob]{
		OutputState: i.ToSynchronizationJobMapOutputWithContext(ctx).OutputState,
	}
}

type SynchronizationJobOutput struct{ *pulumi.OutputState }

func (SynchronizationJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SynchronizationJob)(nil)).Elem()
}

func (o SynchronizationJobOutput) ToSynchronizationJobOutput() SynchronizationJobOutput {
	return o
}

func (o SynchronizationJobOutput) ToSynchronizationJobOutputWithContext(ctx context.Context) SynchronizationJobOutput {
	return o
}

func (o SynchronizationJobOutput) ToOutput(ctx context.Context) pulumix.Output[*SynchronizationJob] {
	return pulumix.Output[*SynchronizationJob]{
		OutputState: o.OutputState,
	}
}

// Whether or not the provisioning job is enabled. Default state is `true`.
func (o SynchronizationJobOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SynchronizationJob) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// A `schedule` list as documented below.
func (o SynchronizationJobOutput) Schedules() SynchronizationJobScheduleArrayOutput {
	return o.ApplyT(func(v *SynchronizationJob) SynchronizationJobScheduleArrayOutput { return v.Schedules }).(SynchronizationJobScheduleArrayOutput)
}

// The object ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
func (o SynchronizationJobOutput) ServicePrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *SynchronizationJob) pulumi.StringOutput { return v.ServicePrincipalId }).(pulumi.StringOutput)
}

// Identifier of the synchronization template this job is based on.
func (o SynchronizationJobOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *SynchronizationJob) pulumi.StringOutput { return v.TemplateId }).(pulumi.StringOutput)
}

type SynchronizationJobArrayOutput struct{ *pulumi.OutputState }

func (SynchronizationJobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SynchronizationJob)(nil)).Elem()
}

func (o SynchronizationJobArrayOutput) ToSynchronizationJobArrayOutput() SynchronizationJobArrayOutput {
	return o
}

func (o SynchronizationJobArrayOutput) ToSynchronizationJobArrayOutputWithContext(ctx context.Context) SynchronizationJobArrayOutput {
	return o
}

func (o SynchronizationJobArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SynchronizationJob] {
	return pulumix.Output[[]*SynchronizationJob]{
		OutputState: o.OutputState,
	}
}

func (o SynchronizationJobArrayOutput) Index(i pulumi.IntInput) SynchronizationJobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SynchronizationJob {
		return vs[0].([]*SynchronizationJob)[vs[1].(int)]
	}).(SynchronizationJobOutput)
}

type SynchronizationJobMapOutput struct{ *pulumi.OutputState }

func (SynchronizationJobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SynchronizationJob)(nil)).Elem()
}

func (o SynchronizationJobMapOutput) ToSynchronizationJobMapOutput() SynchronizationJobMapOutput {
	return o
}

func (o SynchronizationJobMapOutput) ToSynchronizationJobMapOutputWithContext(ctx context.Context) SynchronizationJobMapOutput {
	return o
}

func (o SynchronizationJobMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SynchronizationJob] {
	return pulumix.Output[map[string]*SynchronizationJob]{
		OutputState: o.OutputState,
	}
}

func (o SynchronizationJobMapOutput) MapIndex(k pulumi.StringInput) SynchronizationJobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SynchronizationJob {
		return vs[0].(map[string]*SynchronizationJob)[vs[1].(string)]
	}).(SynchronizationJobOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SynchronizationJobInput)(nil)).Elem(), &SynchronizationJob{})
	pulumi.RegisterInputType(reflect.TypeOf((*SynchronizationJobArrayInput)(nil)).Elem(), SynchronizationJobArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SynchronizationJobMapInput)(nil)).Elem(), SynchronizationJobMap{})
	pulumi.RegisterOutputType(SynchronizationJobOutput{})
	pulumi.RegisterOutputType(SynchronizationJobArrayOutput{})
	pulumi.RegisterOutputType(SynchronizationJobMapOutput{})
}

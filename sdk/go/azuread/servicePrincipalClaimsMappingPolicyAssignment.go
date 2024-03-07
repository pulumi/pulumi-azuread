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

// Manages a Claims Mapping Policy Assignment within Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ApplicationConfiguration` and `Policy.Read.All`
//
// When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := azuread.NewServicePrincipalClaimsMappingPolicyAssignment(ctx, "app", &azuread.ServicePrincipalClaimsMappingPolicyAssignmentArgs{
//				ClaimsMappingPolicyId: pulumi.Any(myPolicy.Id),
//				ServicePrincipalId:    pulumi.Any(myPrincipal.Id),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Claims Mapping Policy can be imported using the `id`, in the form `service-principal-uuid/claimsMappingPolicy/claims-mapping-policy-uuid`, e.g:
//
// ```sh
// $ pulumi import azuread:index/servicePrincipalClaimsMappingPolicyAssignment:ServicePrincipalClaimsMappingPolicyAssignment app 00000000-0000-0000-0000-000000000000/claimsMappingPolicy/11111111-0000-0000-0000-000000000000
// ```
type ServicePrincipalClaimsMappingPolicyAssignment struct {
	pulumi.CustomResourceState

	// The ID of the claims mapping policy to assign.
	ClaimsMappingPolicyId pulumi.StringOutput `pulumi:"claimsMappingPolicyId"`
	// The object ID of the service principal for the policy assignment.
	ServicePrincipalId pulumi.StringOutput `pulumi:"servicePrincipalId"`
}

// NewServicePrincipalClaimsMappingPolicyAssignment registers a new resource with the given unique name, arguments, and options.
func NewServicePrincipalClaimsMappingPolicyAssignment(ctx *pulumi.Context,
	name string, args *ServicePrincipalClaimsMappingPolicyAssignmentArgs, opts ...pulumi.ResourceOption) (*ServicePrincipalClaimsMappingPolicyAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClaimsMappingPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'ClaimsMappingPolicyId'")
	}
	if args.ServicePrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'ServicePrincipalId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServicePrincipalClaimsMappingPolicyAssignment
	err := ctx.RegisterResource("azuread:index/servicePrincipalClaimsMappingPolicyAssignment:ServicePrincipalClaimsMappingPolicyAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicePrincipalClaimsMappingPolicyAssignment gets an existing ServicePrincipalClaimsMappingPolicyAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicePrincipalClaimsMappingPolicyAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicePrincipalClaimsMappingPolicyAssignmentState, opts ...pulumi.ResourceOption) (*ServicePrincipalClaimsMappingPolicyAssignment, error) {
	var resource ServicePrincipalClaimsMappingPolicyAssignment
	err := ctx.ReadResource("azuread:index/servicePrincipalClaimsMappingPolicyAssignment:ServicePrincipalClaimsMappingPolicyAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicePrincipalClaimsMappingPolicyAssignment resources.
type servicePrincipalClaimsMappingPolicyAssignmentState struct {
	// The ID of the claims mapping policy to assign.
	ClaimsMappingPolicyId *string `pulumi:"claimsMappingPolicyId"`
	// The object ID of the service principal for the policy assignment.
	ServicePrincipalId *string `pulumi:"servicePrincipalId"`
}

type ServicePrincipalClaimsMappingPolicyAssignmentState struct {
	// The ID of the claims mapping policy to assign.
	ClaimsMappingPolicyId pulumi.StringPtrInput
	// The object ID of the service principal for the policy assignment.
	ServicePrincipalId pulumi.StringPtrInput
}

func (ServicePrincipalClaimsMappingPolicyAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePrincipalClaimsMappingPolicyAssignmentState)(nil)).Elem()
}

type servicePrincipalClaimsMappingPolicyAssignmentArgs struct {
	// The ID of the claims mapping policy to assign.
	ClaimsMappingPolicyId string `pulumi:"claimsMappingPolicyId"`
	// The object ID of the service principal for the policy assignment.
	ServicePrincipalId string `pulumi:"servicePrincipalId"`
}

// The set of arguments for constructing a ServicePrincipalClaimsMappingPolicyAssignment resource.
type ServicePrincipalClaimsMappingPolicyAssignmentArgs struct {
	// The ID of the claims mapping policy to assign.
	ClaimsMappingPolicyId pulumi.StringInput
	// The object ID of the service principal for the policy assignment.
	ServicePrincipalId pulumi.StringInput
}

func (ServicePrincipalClaimsMappingPolicyAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePrincipalClaimsMappingPolicyAssignmentArgs)(nil)).Elem()
}

type ServicePrincipalClaimsMappingPolicyAssignmentInput interface {
	pulumi.Input

	ToServicePrincipalClaimsMappingPolicyAssignmentOutput() ServicePrincipalClaimsMappingPolicyAssignmentOutput
	ToServicePrincipalClaimsMappingPolicyAssignmentOutputWithContext(ctx context.Context) ServicePrincipalClaimsMappingPolicyAssignmentOutput
}

func (*ServicePrincipalClaimsMappingPolicyAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalClaimsMappingPolicyAssignment)(nil)).Elem()
}

func (i *ServicePrincipalClaimsMappingPolicyAssignment) ToServicePrincipalClaimsMappingPolicyAssignmentOutput() ServicePrincipalClaimsMappingPolicyAssignmentOutput {
	return i.ToServicePrincipalClaimsMappingPolicyAssignmentOutputWithContext(context.Background())
}

func (i *ServicePrincipalClaimsMappingPolicyAssignment) ToServicePrincipalClaimsMappingPolicyAssignmentOutputWithContext(ctx context.Context) ServicePrincipalClaimsMappingPolicyAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalClaimsMappingPolicyAssignmentOutput)
}

// ServicePrincipalClaimsMappingPolicyAssignmentArrayInput is an input type that accepts ServicePrincipalClaimsMappingPolicyAssignmentArray and ServicePrincipalClaimsMappingPolicyAssignmentArrayOutput values.
// You can construct a concrete instance of `ServicePrincipalClaimsMappingPolicyAssignmentArrayInput` via:
//
//	ServicePrincipalClaimsMappingPolicyAssignmentArray{ ServicePrincipalClaimsMappingPolicyAssignmentArgs{...} }
type ServicePrincipalClaimsMappingPolicyAssignmentArrayInput interface {
	pulumi.Input

	ToServicePrincipalClaimsMappingPolicyAssignmentArrayOutput() ServicePrincipalClaimsMappingPolicyAssignmentArrayOutput
	ToServicePrincipalClaimsMappingPolicyAssignmentArrayOutputWithContext(context.Context) ServicePrincipalClaimsMappingPolicyAssignmentArrayOutput
}

type ServicePrincipalClaimsMappingPolicyAssignmentArray []ServicePrincipalClaimsMappingPolicyAssignmentInput

func (ServicePrincipalClaimsMappingPolicyAssignmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePrincipalClaimsMappingPolicyAssignment)(nil)).Elem()
}

func (i ServicePrincipalClaimsMappingPolicyAssignmentArray) ToServicePrincipalClaimsMappingPolicyAssignmentArrayOutput() ServicePrincipalClaimsMappingPolicyAssignmentArrayOutput {
	return i.ToServicePrincipalClaimsMappingPolicyAssignmentArrayOutputWithContext(context.Background())
}

func (i ServicePrincipalClaimsMappingPolicyAssignmentArray) ToServicePrincipalClaimsMappingPolicyAssignmentArrayOutputWithContext(ctx context.Context) ServicePrincipalClaimsMappingPolicyAssignmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalClaimsMappingPolicyAssignmentArrayOutput)
}

// ServicePrincipalClaimsMappingPolicyAssignmentMapInput is an input type that accepts ServicePrincipalClaimsMappingPolicyAssignmentMap and ServicePrincipalClaimsMappingPolicyAssignmentMapOutput values.
// You can construct a concrete instance of `ServicePrincipalClaimsMappingPolicyAssignmentMapInput` via:
//
//	ServicePrincipalClaimsMappingPolicyAssignmentMap{ "key": ServicePrincipalClaimsMappingPolicyAssignmentArgs{...} }
type ServicePrincipalClaimsMappingPolicyAssignmentMapInput interface {
	pulumi.Input

	ToServicePrincipalClaimsMappingPolicyAssignmentMapOutput() ServicePrincipalClaimsMappingPolicyAssignmentMapOutput
	ToServicePrincipalClaimsMappingPolicyAssignmentMapOutputWithContext(context.Context) ServicePrincipalClaimsMappingPolicyAssignmentMapOutput
}

type ServicePrincipalClaimsMappingPolicyAssignmentMap map[string]ServicePrincipalClaimsMappingPolicyAssignmentInput

func (ServicePrincipalClaimsMappingPolicyAssignmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePrincipalClaimsMappingPolicyAssignment)(nil)).Elem()
}

func (i ServicePrincipalClaimsMappingPolicyAssignmentMap) ToServicePrincipalClaimsMappingPolicyAssignmentMapOutput() ServicePrincipalClaimsMappingPolicyAssignmentMapOutput {
	return i.ToServicePrincipalClaimsMappingPolicyAssignmentMapOutputWithContext(context.Background())
}

func (i ServicePrincipalClaimsMappingPolicyAssignmentMap) ToServicePrincipalClaimsMappingPolicyAssignmentMapOutputWithContext(ctx context.Context) ServicePrincipalClaimsMappingPolicyAssignmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalClaimsMappingPolicyAssignmentMapOutput)
}

type ServicePrincipalClaimsMappingPolicyAssignmentOutput struct{ *pulumi.OutputState }

func (ServicePrincipalClaimsMappingPolicyAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalClaimsMappingPolicyAssignment)(nil)).Elem()
}

func (o ServicePrincipalClaimsMappingPolicyAssignmentOutput) ToServicePrincipalClaimsMappingPolicyAssignmentOutput() ServicePrincipalClaimsMappingPolicyAssignmentOutput {
	return o
}

func (o ServicePrincipalClaimsMappingPolicyAssignmentOutput) ToServicePrincipalClaimsMappingPolicyAssignmentOutputWithContext(ctx context.Context) ServicePrincipalClaimsMappingPolicyAssignmentOutput {
	return o
}

// The ID of the claims mapping policy to assign.
func (o ServicePrincipalClaimsMappingPolicyAssignmentOutput) ClaimsMappingPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalClaimsMappingPolicyAssignment) pulumi.StringOutput {
		return v.ClaimsMappingPolicyId
	}).(pulumi.StringOutput)
}

// The object ID of the service principal for the policy assignment.
func (o ServicePrincipalClaimsMappingPolicyAssignmentOutput) ServicePrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServicePrincipalClaimsMappingPolicyAssignment) pulumi.StringOutput {
		return v.ServicePrincipalId
	}).(pulumi.StringOutput)
}

type ServicePrincipalClaimsMappingPolicyAssignmentArrayOutput struct{ *pulumi.OutputState }

func (ServicePrincipalClaimsMappingPolicyAssignmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePrincipalClaimsMappingPolicyAssignment)(nil)).Elem()
}

func (o ServicePrincipalClaimsMappingPolicyAssignmentArrayOutput) ToServicePrincipalClaimsMappingPolicyAssignmentArrayOutput() ServicePrincipalClaimsMappingPolicyAssignmentArrayOutput {
	return o
}

func (o ServicePrincipalClaimsMappingPolicyAssignmentArrayOutput) ToServicePrincipalClaimsMappingPolicyAssignmentArrayOutputWithContext(ctx context.Context) ServicePrincipalClaimsMappingPolicyAssignmentArrayOutput {
	return o
}

func (o ServicePrincipalClaimsMappingPolicyAssignmentArrayOutput) Index(i pulumi.IntInput) ServicePrincipalClaimsMappingPolicyAssignmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServicePrincipalClaimsMappingPolicyAssignment {
		return vs[0].([]*ServicePrincipalClaimsMappingPolicyAssignment)[vs[1].(int)]
	}).(ServicePrincipalClaimsMappingPolicyAssignmentOutput)
}

type ServicePrincipalClaimsMappingPolicyAssignmentMapOutput struct{ *pulumi.OutputState }

func (ServicePrincipalClaimsMappingPolicyAssignmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePrincipalClaimsMappingPolicyAssignment)(nil)).Elem()
}

func (o ServicePrincipalClaimsMappingPolicyAssignmentMapOutput) ToServicePrincipalClaimsMappingPolicyAssignmentMapOutput() ServicePrincipalClaimsMappingPolicyAssignmentMapOutput {
	return o
}

func (o ServicePrincipalClaimsMappingPolicyAssignmentMapOutput) ToServicePrincipalClaimsMappingPolicyAssignmentMapOutputWithContext(ctx context.Context) ServicePrincipalClaimsMappingPolicyAssignmentMapOutput {
	return o
}

func (o ServicePrincipalClaimsMappingPolicyAssignmentMapOutput) MapIndex(k pulumi.StringInput) ServicePrincipalClaimsMappingPolicyAssignmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServicePrincipalClaimsMappingPolicyAssignment {
		return vs[0].(map[string]*ServicePrincipalClaimsMappingPolicyAssignment)[vs[1].(string)]
	}).(ServicePrincipalClaimsMappingPolicyAssignmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalClaimsMappingPolicyAssignmentInput)(nil)).Elem(), &ServicePrincipalClaimsMappingPolicyAssignment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalClaimsMappingPolicyAssignmentArrayInput)(nil)).Elem(), ServicePrincipalClaimsMappingPolicyAssignmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalClaimsMappingPolicyAssignmentMapInput)(nil)).Elem(), ServicePrincipalClaimsMappingPolicyAssignmentMap{})
	pulumi.RegisterOutputType(ServicePrincipalClaimsMappingPolicyAssignmentOutput{})
	pulumi.RegisterOutputType(ServicePrincipalClaimsMappingPolicyAssignmentArrayOutput{})
	pulumi.RegisterOutputType(ServicePrincipalClaimsMappingPolicyAssignmentMapOutput{})
}

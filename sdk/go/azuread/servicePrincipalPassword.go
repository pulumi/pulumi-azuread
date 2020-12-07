// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Password associated with a Service Principal within Azure Active Directory.
//
// > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azuread/sdk/v3/go/azuread"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleApplication, err := azuread.NewApplication(ctx, "exampleApplication", nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleServicePrincipal, err := azuread.NewServicePrincipal(ctx, "exampleServicePrincipal", &azuread.ServicePrincipalArgs{
// 			ApplicationId: exampleApplication.ApplicationId,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuread.NewServicePrincipalPassword(ctx, "exampleServicePrincipalPassword", &azuread.ServicePrincipalPasswordArgs{
// 			ServicePrincipalId: exampleServicePrincipal.ID(),
// 			Description:        pulumi.String("My managed password"),
// 			Value:              pulumi.String(fmt.Sprintf("%v%v%v", "VT=uSgbTanZhyz@", "%", "nL9Hpd+Tfay_MRV#")),
// 			EndDate:            pulumi.String("2099-01-01T01:02:03Z"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Passwords can be imported using the `object id` of a Service Principal and the `key id` of the password, e.g.
//
// ```sh
//  $ pulumi import azuread:index/servicePrincipalPassword:ServicePrincipalPassword test 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111
// ```
type ServicePrincipalPassword struct {
	pulumi.CustomResourceState

	// A description for the Password.
	Description pulumi.StringOutput `pulumi:"description"`
	// The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate pulumi.StringOutput `pulumi:"endDate"`
	// A relative duration for which the Password is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
	EndDateRelative pulumi.StringPtrOutput `pulumi:"endDateRelative"`
	// A GUID used to uniquely identify this Key. If not specified a GUID will be created. Changing this field forces a new resource to be created.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// The ID of the Service Principal for which this password should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId pulumi.StringOutput `pulumi:"servicePrincipalId"`
	// The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	StartDate pulumi.StringOutput `pulumi:"startDate"`
	// The Password for this Service Principal.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewServicePrincipalPassword registers a new resource with the given unique name, arguments, and options.
func NewServicePrincipalPassword(ctx *pulumi.Context,
	name string, args *ServicePrincipalPasswordArgs, opts ...pulumi.ResourceOption) (*ServicePrincipalPassword, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServicePrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'ServicePrincipalId'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	var resource ServicePrincipalPassword
	err := ctx.RegisterResource("azuread:index/servicePrincipalPassword:ServicePrincipalPassword", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicePrincipalPassword gets an existing ServicePrincipalPassword resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicePrincipalPassword(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicePrincipalPasswordState, opts ...pulumi.ResourceOption) (*ServicePrincipalPassword, error) {
	var resource ServicePrincipalPassword
	err := ctx.ReadResource("azuread:index/servicePrincipalPassword:ServicePrincipalPassword", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicePrincipalPassword resources.
type servicePrincipalPasswordState struct {
	// A description for the Password.
	Description *string `pulumi:"description"`
	// The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate *string `pulumi:"endDate"`
	// A relative duration for which the Password is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
	EndDateRelative *string `pulumi:"endDateRelative"`
	// A GUID used to uniquely identify this Key. If not specified a GUID will be created. Changing this field forces a new resource to be created.
	KeyId *string `pulumi:"keyId"`
	// The ID of the Service Principal for which this password should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId *string `pulumi:"servicePrincipalId"`
	// The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	StartDate *string `pulumi:"startDate"`
	// The Password for this Service Principal.
	Value *string `pulumi:"value"`
}

type ServicePrincipalPasswordState struct {
	// A description for the Password.
	Description pulumi.StringPtrInput
	// The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate pulumi.StringPtrInput
	// A relative duration for which the Password is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
	EndDateRelative pulumi.StringPtrInput
	// A GUID used to uniquely identify this Key. If not specified a GUID will be created. Changing this field forces a new resource to be created.
	KeyId pulumi.StringPtrInput
	// The ID of the Service Principal for which this password should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId pulumi.StringPtrInput
	// The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	StartDate pulumi.StringPtrInput
	// The Password for this Service Principal.
	Value pulumi.StringPtrInput
}

func (ServicePrincipalPasswordState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePrincipalPasswordState)(nil)).Elem()
}

type servicePrincipalPasswordArgs struct {
	// A description for the Password.
	Description *string `pulumi:"description"`
	// The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate *string `pulumi:"endDate"`
	// A relative duration for which the Password is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
	EndDateRelative *string `pulumi:"endDateRelative"`
	// A GUID used to uniquely identify this Key. If not specified a GUID will be created. Changing this field forces a new resource to be created.
	KeyId *string `pulumi:"keyId"`
	// The ID of the Service Principal for which this password should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId string `pulumi:"servicePrincipalId"`
	// The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	StartDate *string `pulumi:"startDate"`
	// The Password for this Service Principal.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a ServicePrincipalPassword resource.
type ServicePrincipalPasswordArgs struct {
	// A description for the Password.
	Description pulumi.StringPtrInput
	// The End Date which the Password is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
	EndDate pulumi.StringPtrInput
	// A relative duration for which the Password is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
	EndDateRelative pulumi.StringPtrInput
	// A GUID used to uniquely identify this Key. If not specified a GUID will be created. Changing this field forces a new resource to be created.
	KeyId pulumi.StringPtrInput
	// The ID of the Service Principal for which this password should be created. Changing this field forces a new resource to be created.
	ServicePrincipalId pulumi.StringInput
	// The Start Date which the Password is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
	StartDate pulumi.StringPtrInput
	// The Password for this Service Principal.
	Value pulumi.StringInput
}

func (ServicePrincipalPasswordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePrincipalPasswordArgs)(nil)).Elem()
}

type ServicePrincipalPasswordInput interface {
	pulumi.Input

	ToServicePrincipalPasswordOutput() ServicePrincipalPasswordOutput
	ToServicePrincipalPasswordOutputWithContext(ctx context.Context) ServicePrincipalPasswordOutput
}

func (ServicePrincipalPassword) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalPassword)(nil)).Elem()
}

func (i ServicePrincipalPassword) ToServicePrincipalPasswordOutput() ServicePrincipalPasswordOutput {
	return i.ToServicePrincipalPasswordOutputWithContext(context.Background())
}

func (i ServicePrincipalPassword) ToServicePrincipalPasswordOutputWithContext(ctx context.Context) ServicePrincipalPasswordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalPasswordOutput)
}

type ServicePrincipalPasswordOutput struct {
	*pulumi.OutputState
}

func (ServicePrincipalPasswordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalPasswordOutput)(nil)).Elem()
}

func (o ServicePrincipalPasswordOutput) ToServicePrincipalPasswordOutput() ServicePrincipalPasswordOutput {
	return o
}

func (o ServicePrincipalPasswordOutput) ToServicePrincipalPasswordOutputWithContext(ctx context.Context) ServicePrincipalPasswordOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ServicePrincipalPasswordOutput{})
}

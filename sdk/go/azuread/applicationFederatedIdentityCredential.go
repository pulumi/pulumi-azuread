// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"errors"
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
//			exampleApplication, err := azuread.NewApplication(ctx, "exampleApplication", &azuread.ApplicationArgs{
//				DisplayName: pulumi.String("example"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = azuread.NewApplicationFederatedIdentityCredential(ctx, "exampleApplicationFederatedIdentityCredential", &azuread.ApplicationFederatedIdentityCredentialArgs{
//				ApplicationObjectId: exampleApplication.ObjectId,
//				DisplayName:         pulumi.String("my-repo-deploy"),
//				Description:         pulumi.String("Deployments for my-repo"),
//				Audiences: pulumi.StringArray{
//					pulumi.String("api://AzureADTokenExchange"),
//				},
//				Issuer:  pulumi.String("https://token.actions.githubusercontent.com"),
//				Subject: pulumi.String("repo:my-organization/my-repo:environment:prod"),
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
// Federated Identity Credentials can be imported using the object ID of the associated application and the ID of the federated identity credential, e.g.
//
// ```sh
//
//	$ pulumi import azuread:index/applicationFederatedIdentityCredential:ApplicationFederatedIdentityCredential test 00000000-0000-0000-0000-000000000000/federatedIdentityCredential/11111111-1111-1111-1111-111111111111
//
// ```
//
//	-> This ID format is unique to Terraform and is composed of the application's object ID, the string "federatedIdentityCredential" and the credential ID in the format `{ObjectId}/federatedIdentityCredential/{CredentialId}`.
type ApplicationFederatedIdentityCredential struct {
	pulumi.CustomResourceState

	// The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
	ApplicationObjectId pulumi.StringOutput `pulumi:"applicationObjectId"`
	// List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
	Audiences pulumi.StringArrayOutput `pulumi:"audiences"`
	// A UUID used to uniquely identify this federated identity credential.
	CredentialId pulumi.StringOutput `pulumi:"credentialId"`
	// A description for the federated identity credential.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A unique display name for the federated identity credential. Changing this forces a new resource to be created.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
	Issuer pulumi.StringOutput `pulumi:"issuer"`
	// The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
	Subject pulumi.StringOutput `pulumi:"subject"`
}

// NewApplicationFederatedIdentityCredential registers a new resource with the given unique name, arguments, and options.
func NewApplicationFederatedIdentityCredential(ctx *pulumi.Context,
	name string, args *ApplicationFederatedIdentityCredentialArgs, opts ...pulumi.ResourceOption) (*ApplicationFederatedIdentityCredential, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationObjectId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationObjectId'")
	}
	if args.Audiences == nil {
		return nil, errors.New("invalid value for required argument 'Audiences'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Issuer == nil {
		return nil, errors.New("invalid value for required argument 'Issuer'")
	}
	if args.Subject == nil {
		return nil, errors.New("invalid value for required argument 'Subject'")
	}
	var resource ApplicationFederatedIdentityCredential
	err := ctx.RegisterResource("azuread:index/applicationFederatedIdentityCredential:ApplicationFederatedIdentityCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationFederatedIdentityCredential gets an existing ApplicationFederatedIdentityCredential resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationFederatedIdentityCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationFederatedIdentityCredentialState, opts ...pulumi.ResourceOption) (*ApplicationFederatedIdentityCredential, error) {
	var resource ApplicationFederatedIdentityCredential
	err := ctx.ReadResource("azuread:index/applicationFederatedIdentityCredential:ApplicationFederatedIdentityCredential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationFederatedIdentityCredential resources.
type applicationFederatedIdentityCredentialState struct {
	// The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
	ApplicationObjectId *string `pulumi:"applicationObjectId"`
	// List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
	Audiences []string `pulumi:"audiences"`
	// A UUID used to uniquely identify this federated identity credential.
	CredentialId *string `pulumi:"credentialId"`
	// A description for the federated identity credential.
	Description *string `pulumi:"description"`
	// A unique display name for the federated identity credential. Changing this forces a new resource to be created.
	DisplayName *string `pulumi:"displayName"`
	// The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
	Issuer *string `pulumi:"issuer"`
	// The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
	Subject *string `pulumi:"subject"`
}

type ApplicationFederatedIdentityCredentialState struct {
	// The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
	ApplicationObjectId pulumi.StringPtrInput
	// List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
	Audiences pulumi.StringArrayInput
	// A UUID used to uniquely identify this federated identity credential.
	CredentialId pulumi.StringPtrInput
	// A description for the federated identity credential.
	Description pulumi.StringPtrInput
	// A unique display name for the federated identity credential. Changing this forces a new resource to be created.
	DisplayName pulumi.StringPtrInput
	// The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
	Issuer pulumi.StringPtrInput
	// The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
	Subject pulumi.StringPtrInput
}

func (ApplicationFederatedIdentityCredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationFederatedIdentityCredentialState)(nil)).Elem()
}

type applicationFederatedIdentityCredentialArgs struct {
	// The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
	ApplicationObjectId string `pulumi:"applicationObjectId"`
	// List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
	Audiences []string `pulumi:"audiences"`
	// A description for the federated identity credential.
	Description *string `pulumi:"description"`
	// A unique display name for the federated identity credential. Changing this forces a new resource to be created.
	DisplayName string `pulumi:"displayName"`
	// The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
	Issuer string `pulumi:"issuer"`
	// The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
	Subject string `pulumi:"subject"`
}

// The set of arguments for constructing a ApplicationFederatedIdentityCredential resource.
type ApplicationFederatedIdentityCredentialArgs struct {
	// The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
	ApplicationObjectId pulumi.StringInput
	// List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
	Audiences pulumi.StringArrayInput
	// A description for the federated identity credential.
	Description pulumi.StringPtrInput
	// A unique display name for the federated identity credential. Changing this forces a new resource to be created.
	DisplayName pulumi.StringInput
	// The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
	Issuer pulumi.StringInput
	// The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
	Subject pulumi.StringInput
}

func (ApplicationFederatedIdentityCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationFederatedIdentityCredentialArgs)(nil)).Elem()
}

type ApplicationFederatedIdentityCredentialInput interface {
	pulumi.Input

	ToApplicationFederatedIdentityCredentialOutput() ApplicationFederatedIdentityCredentialOutput
	ToApplicationFederatedIdentityCredentialOutputWithContext(ctx context.Context) ApplicationFederatedIdentityCredentialOutput
}

func (*ApplicationFederatedIdentityCredential) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationFederatedIdentityCredential)(nil)).Elem()
}

func (i *ApplicationFederatedIdentityCredential) ToApplicationFederatedIdentityCredentialOutput() ApplicationFederatedIdentityCredentialOutput {
	return i.ToApplicationFederatedIdentityCredentialOutputWithContext(context.Background())
}

func (i *ApplicationFederatedIdentityCredential) ToApplicationFederatedIdentityCredentialOutputWithContext(ctx context.Context) ApplicationFederatedIdentityCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationFederatedIdentityCredentialOutput)
}

// ApplicationFederatedIdentityCredentialArrayInput is an input type that accepts ApplicationFederatedIdentityCredentialArray and ApplicationFederatedIdentityCredentialArrayOutput values.
// You can construct a concrete instance of `ApplicationFederatedIdentityCredentialArrayInput` via:
//
//	ApplicationFederatedIdentityCredentialArray{ ApplicationFederatedIdentityCredentialArgs{...} }
type ApplicationFederatedIdentityCredentialArrayInput interface {
	pulumi.Input

	ToApplicationFederatedIdentityCredentialArrayOutput() ApplicationFederatedIdentityCredentialArrayOutput
	ToApplicationFederatedIdentityCredentialArrayOutputWithContext(context.Context) ApplicationFederatedIdentityCredentialArrayOutput
}

type ApplicationFederatedIdentityCredentialArray []ApplicationFederatedIdentityCredentialInput

func (ApplicationFederatedIdentityCredentialArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationFederatedIdentityCredential)(nil)).Elem()
}

func (i ApplicationFederatedIdentityCredentialArray) ToApplicationFederatedIdentityCredentialArrayOutput() ApplicationFederatedIdentityCredentialArrayOutput {
	return i.ToApplicationFederatedIdentityCredentialArrayOutputWithContext(context.Background())
}

func (i ApplicationFederatedIdentityCredentialArray) ToApplicationFederatedIdentityCredentialArrayOutputWithContext(ctx context.Context) ApplicationFederatedIdentityCredentialArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationFederatedIdentityCredentialArrayOutput)
}

// ApplicationFederatedIdentityCredentialMapInput is an input type that accepts ApplicationFederatedIdentityCredentialMap and ApplicationFederatedIdentityCredentialMapOutput values.
// You can construct a concrete instance of `ApplicationFederatedIdentityCredentialMapInput` via:
//
//	ApplicationFederatedIdentityCredentialMap{ "key": ApplicationFederatedIdentityCredentialArgs{...} }
type ApplicationFederatedIdentityCredentialMapInput interface {
	pulumi.Input

	ToApplicationFederatedIdentityCredentialMapOutput() ApplicationFederatedIdentityCredentialMapOutput
	ToApplicationFederatedIdentityCredentialMapOutputWithContext(context.Context) ApplicationFederatedIdentityCredentialMapOutput
}

type ApplicationFederatedIdentityCredentialMap map[string]ApplicationFederatedIdentityCredentialInput

func (ApplicationFederatedIdentityCredentialMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationFederatedIdentityCredential)(nil)).Elem()
}

func (i ApplicationFederatedIdentityCredentialMap) ToApplicationFederatedIdentityCredentialMapOutput() ApplicationFederatedIdentityCredentialMapOutput {
	return i.ToApplicationFederatedIdentityCredentialMapOutputWithContext(context.Background())
}

func (i ApplicationFederatedIdentityCredentialMap) ToApplicationFederatedIdentityCredentialMapOutputWithContext(ctx context.Context) ApplicationFederatedIdentityCredentialMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationFederatedIdentityCredentialMapOutput)
}

type ApplicationFederatedIdentityCredentialOutput struct{ *pulumi.OutputState }

func (ApplicationFederatedIdentityCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationFederatedIdentityCredential)(nil)).Elem()
}

func (o ApplicationFederatedIdentityCredentialOutput) ToApplicationFederatedIdentityCredentialOutput() ApplicationFederatedIdentityCredentialOutput {
	return o
}

func (o ApplicationFederatedIdentityCredentialOutput) ToApplicationFederatedIdentityCredentialOutputWithContext(ctx context.Context) ApplicationFederatedIdentityCredentialOutput {
	return o
}

// The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
func (o ApplicationFederatedIdentityCredentialOutput) ApplicationObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationFederatedIdentityCredential) pulumi.StringOutput { return v.ApplicationObjectId }).(pulumi.StringOutput)
}

// List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
func (o ApplicationFederatedIdentityCredentialOutput) Audiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationFederatedIdentityCredential) pulumi.StringArrayOutput { return v.Audiences }).(pulumi.StringArrayOutput)
}

// A UUID used to uniquely identify this federated identity credential.
func (o ApplicationFederatedIdentityCredentialOutput) CredentialId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationFederatedIdentityCredential) pulumi.StringOutput { return v.CredentialId }).(pulumi.StringOutput)
}

// A description for the federated identity credential.
func (o ApplicationFederatedIdentityCredentialOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationFederatedIdentityCredential) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique display name for the federated identity credential. Changing this forces a new resource to be created.
func (o ApplicationFederatedIdentityCredentialOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationFederatedIdentityCredential) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
func (o ApplicationFederatedIdentityCredentialOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationFederatedIdentityCredential) pulumi.StringOutput { return v.Issuer }).(pulumi.StringOutput)
}

// The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
func (o ApplicationFederatedIdentityCredentialOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationFederatedIdentityCredential) pulumi.StringOutput { return v.Subject }).(pulumi.StringOutput)
}

type ApplicationFederatedIdentityCredentialArrayOutput struct{ *pulumi.OutputState }

func (ApplicationFederatedIdentityCredentialArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationFederatedIdentityCredential)(nil)).Elem()
}

func (o ApplicationFederatedIdentityCredentialArrayOutput) ToApplicationFederatedIdentityCredentialArrayOutput() ApplicationFederatedIdentityCredentialArrayOutput {
	return o
}

func (o ApplicationFederatedIdentityCredentialArrayOutput) ToApplicationFederatedIdentityCredentialArrayOutputWithContext(ctx context.Context) ApplicationFederatedIdentityCredentialArrayOutput {
	return o
}

func (o ApplicationFederatedIdentityCredentialArrayOutput) Index(i pulumi.IntInput) ApplicationFederatedIdentityCredentialOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationFederatedIdentityCredential {
		return vs[0].([]*ApplicationFederatedIdentityCredential)[vs[1].(int)]
	}).(ApplicationFederatedIdentityCredentialOutput)
}

type ApplicationFederatedIdentityCredentialMapOutput struct{ *pulumi.OutputState }

func (ApplicationFederatedIdentityCredentialMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationFederatedIdentityCredential)(nil)).Elem()
}

func (o ApplicationFederatedIdentityCredentialMapOutput) ToApplicationFederatedIdentityCredentialMapOutput() ApplicationFederatedIdentityCredentialMapOutput {
	return o
}

func (o ApplicationFederatedIdentityCredentialMapOutput) ToApplicationFederatedIdentityCredentialMapOutputWithContext(ctx context.Context) ApplicationFederatedIdentityCredentialMapOutput {
	return o
}

func (o ApplicationFederatedIdentityCredentialMapOutput) MapIndex(k pulumi.StringInput) ApplicationFederatedIdentityCredentialOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationFederatedIdentityCredential {
		return vs[0].(map[string]*ApplicationFederatedIdentityCredential)[vs[1].(string)]
	}).(ApplicationFederatedIdentityCredentialOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationFederatedIdentityCredentialInput)(nil)).Elem(), &ApplicationFederatedIdentityCredential{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationFederatedIdentityCredentialArrayInput)(nil)).Elem(), ApplicationFederatedIdentityCredentialArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationFederatedIdentityCredentialMapInput)(nil)).Elem(), ApplicationFederatedIdentityCredentialMap{})
	pulumi.RegisterOutputType(ApplicationFederatedIdentityCredentialOutput{})
	pulumi.RegisterOutputType(ApplicationFederatedIdentityCredentialArrayOutput{})
	pulumi.RegisterOutputType(ApplicationFederatedIdentityCredentialMapOutput{})
}

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

// Manages an invitation of a guest user within Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires one of the following application roles: `User.Invite.All`, `User.ReadWrite.All` or `Directory.ReadWrite.All`
//
// When authenticated with a user principal, this resource requires one of the following directory roles: `Guest Inviter`, `User Administrator` or `Global Administrator`
//
// ## Example Usage
//
// *Basic example*
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
//			_, err := azuread.NewInvitation(ctx, "example", &azuread.InvitationArgs{
//				UserEmailAddress: pulumi.String("jdoe@example.com"),
//				RedirectUrl:      pulumi.String("https://portal.azure.com"),
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
// *Invitation with standard message*
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
//			_, err := azuread.NewInvitation(ctx, "example", &azuread.InvitationArgs{
//				UserEmailAddress: pulumi.String("jdoe@example.com"),
//				RedirectUrl:      pulumi.String("https://portal.azure.com"),
//				Message: &azuread.InvitationMessageArgs{
//					Language: pulumi.String("en-US"),
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
// <!--End PulumiCodeChooser -->
//
// *Invitation with custom message body and an additional recipient*
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
//			_, err := azuread.NewInvitation(ctx, "example", &azuread.InvitationArgs{
//				UserDisplayName:  pulumi.String("Bob Bobson"),
//				UserEmailAddress: pulumi.String("bbobson@example.com"),
//				RedirectUrl:      pulumi.String("https://portal.azure.com"),
//				Message: &azuread.InvitationMessageArgs{
//					AdditionalRecipients: pulumi.String("aaliceberg@example.com"),
//					Body:                 pulumi.String("Hello there! You are invited to join my Azure tenant!"),
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// This resource does not support importing.
type Invitation struct {
	pulumi.CustomResourceState

	// A `message` block as documented below, which configures the message being sent to the invited user. If this block is omitted, no message will be sent.
	Message InvitationMessagePtrOutput `pulumi:"message"`
	// The URL the user can use to redeem their invitation.
	RedeemUrl pulumi.StringOutput `pulumi:"redeemUrl"`
	// The URL that the user should be redirected to once the invitation is redeemed.
	RedirectUrl pulumi.StringOutput `pulumi:"redirectUrl"`
	// The display name of the user being invited.
	UserDisplayName pulumi.StringPtrOutput `pulumi:"userDisplayName"`
	// The email address of the user being invited.
	UserEmailAddress pulumi.StringOutput `pulumi:"userEmailAddress"`
	// Object ID of the invited user.
	UserId pulumi.StringOutput `pulumi:"userId"`
	// The user type of the user being invited. Must be one of `Guest` or `Member`. Only Global Administrators can invite users as members. Defaults to `Guest`.
	UserType pulumi.StringPtrOutput `pulumi:"userType"`
}

// NewInvitation registers a new resource with the given unique name, arguments, and options.
func NewInvitation(ctx *pulumi.Context,
	name string, args *InvitationArgs, opts ...pulumi.ResourceOption) (*Invitation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RedirectUrl == nil {
		return nil, errors.New("invalid value for required argument 'RedirectUrl'")
	}
	if args.UserEmailAddress == nil {
		return nil, errors.New("invalid value for required argument 'UserEmailAddress'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Invitation
	err := ctx.RegisterResource("azuread:index/invitation:Invitation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInvitation gets an existing Invitation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInvitation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InvitationState, opts ...pulumi.ResourceOption) (*Invitation, error) {
	var resource Invitation
	err := ctx.ReadResource("azuread:index/invitation:Invitation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Invitation resources.
type invitationState struct {
	// A `message` block as documented below, which configures the message being sent to the invited user. If this block is omitted, no message will be sent.
	Message *InvitationMessage `pulumi:"message"`
	// The URL the user can use to redeem their invitation.
	RedeemUrl *string `pulumi:"redeemUrl"`
	// The URL that the user should be redirected to once the invitation is redeemed.
	RedirectUrl *string `pulumi:"redirectUrl"`
	// The display name of the user being invited.
	UserDisplayName *string `pulumi:"userDisplayName"`
	// The email address of the user being invited.
	UserEmailAddress *string `pulumi:"userEmailAddress"`
	// Object ID of the invited user.
	UserId *string `pulumi:"userId"`
	// The user type of the user being invited. Must be one of `Guest` or `Member`. Only Global Administrators can invite users as members. Defaults to `Guest`.
	UserType *string `pulumi:"userType"`
}

type InvitationState struct {
	// A `message` block as documented below, which configures the message being sent to the invited user. If this block is omitted, no message will be sent.
	Message InvitationMessagePtrInput
	// The URL the user can use to redeem their invitation.
	RedeemUrl pulumi.StringPtrInput
	// The URL that the user should be redirected to once the invitation is redeemed.
	RedirectUrl pulumi.StringPtrInput
	// The display name of the user being invited.
	UserDisplayName pulumi.StringPtrInput
	// The email address of the user being invited.
	UserEmailAddress pulumi.StringPtrInput
	// Object ID of the invited user.
	UserId pulumi.StringPtrInput
	// The user type of the user being invited. Must be one of `Guest` or `Member`. Only Global Administrators can invite users as members. Defaults to `Guest`.
	UserType pulumi.StringPtrInput
}

func (InvitationState) ElementType() reflect.Type {
	return reflect.TypeOf((*invitationState)(nil)).Elem()
}

type invitationArgs struct {
	// A `message` block as documented below, which configures the message being sent to the invited user. If this block is omitted, no message will be sent.
	Message *InvitationMessage `pulumi:"message"`
	// The URL that the user should be redirected to once the invitation is redeemed.
	RedirectUrl string `pulumi:"redirectUrl"`
	// The display name of the user being invited.
	UserDisplayName *string `pulumi:"userDisplayName"`
	// The email address of the user being invited.
	UserEmailAddress string `pulumi:"userEmailAddress"`
	// The user type of the user being invited. Must be one of `Guest` or `Member`. Only Global Administrators can invite users as members. Defaults to `Guest`.
	UserType *string `pulumi:"userType"`
}

// The set of arguments for constructing a Invitation resource.
type InvitationArgs struct {
	// A `message` block as documented below, which configures the message being sent to the invited user. If this block is omitted, no message will be sent.
	Message InvitationMessagePtrInput
	// The URL that the user should be redirected to once the invitation is redeemed.
	RedirectUrl pulumi.StringInput
	// The display name of the user being invited.
	UserDisplayName pulumi.StringPtrInput
	// The email address of the user being invited.
	UserEmailAddress pulumi.StringInput
	// The user type of the user being invited. Must be one of `Guest` or `Member`. Only Global Administrators can invite users as members. Defaults to `Guest`.
	UserType pulumi.StringPtrInput
}

func (InvitationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*invitationArgs)(nil)).Elem()
}

type InvitationInput interface {
	pulumi.Input

	ToInvitationOutput() InvitationOutput
	ToInvitationOutputWithContext(ctx context.Context) InvitationOutput
}

func (*Invitation) ElementType() reflect.Type {
	return reflect.TypeOf((**Invitation)(nil)).Elem()
}

func (i *Invitation) ToInvitationOutput() InvitationOutput {
	return i.ToInvitationOutputWithContext(context.Background())
}

func (i *Invitation) ToInvitationOutputWithContext(ctx context.Context) InvitationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InvitationOutput)
}

// InvitationArrayInput is an input type that accepts InvitationArray and InvitationArrayOutput values.
// You can construct a concrete instance of `InvitationArrayInput` via:
//
//	InvitationArray{ InvitationArgs{...} }
type InvitationArrayInput interface {
	pulumi.Input

	ToInvitationArrayOutput() InvitationArrayOutput
	ToInvitationArrayOutputWithContext(context.Context) InvitationArrayOutput
}

type InvitationArray []InvitationInput

func (InvitationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Invitation)(nil)).Elem()
}

func (i InvitationArray) ToInvitationArrayOutput() InvitationArrayOutput {
	return i.ToInvitationArrayOutputWithContext(context.Background())
}

func (i InvitationArray) ToInvitationArrayOutputWithContext(ctx context.Context) InvitationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InvitationArrayOutput)
}

// InvitationMapInput is an input type that accepts InvitationMap and InvitationMapOutput values.
// You can construct a concrete instance of `InvitationMapInput` via:
//
//	InvitationMap{ "key": InvitationArgs{...} }
type InvitationMapInput interface {
	pulumi.Input

	ToInvitationMapOutput() InvitationMapOutput
	ToInvitationMapOutputWithContext(context.Context) InvitationMapOutput
}

type InvitationMap map[string]InvitationInput

func (InvitationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Invitation)(nil)).Elem()
}

func (i InvitationMap) ToInvitationMapOutput() InvitationMapOutput {
	return i.ToInvitationMapOutputWithContext(context.Background())
}

func (i InvitationMap) ToInvitationMapOutputWithContext(ctx context.Context) InvitationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InvitationMapOutput)
}

type InvitationOutput struct{ *pulumi.OutputState }

func (InvitationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Invitation)(nil)).Elem()
}

func (o InvitationOutput) ToInvitationOutput() InvitationOutput {
	return o
}

func (o InvitationOutput) ToInvitationOutputWithContext(ctx context.Context) InvitationOutput {
	return o
}

// A `message` block as documented below, which configures the message being sent to the invited user. If this block is omitted, no message will be sent.
func (o InvitationOutput) Message() InvitationMessagePtrOutput {
	return o.ApplyT(func(v *Invitation) InvitationMessagePtrOutput { return v.Message }).(InvitationMessagePtrOutput)
}

// The URL the user can use to redeem their invitation.
func (o InvitationOutput) RedeemUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Invitation) pulumi.StringOutput { return v.RedeemUrl }).(pulumi.StringOutput)
}

// The URL that the user should be redirected to once the invitation is redeemed.
func (o InvitationOutput) RedirectUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Invitation) pulumi.StringOutput { return v.RedirectUrl }).(pulumi.StringOutput)
}

// The display name of the user being invited.
func (o InvitationOutput) UserDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Invitation) pulumi.StringPtrOutput { return v.UserDisplayName }).(pulumi.StringPtrOutput)
}

// The email address of the user being invited.
func (o InvitationOutput) UserEmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Invitation) pulumi.StringOutput { return v.UserEmailAddress }).(pulumi.StringOutput)
}

// Object ID of the invited user.
func (o InvitationOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *Invitation) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

// The user type of the user being invited. Must be one of `Guest` or `Member`. Only Global Administrators can invite users as members. Defaults to `Guest`.
func (o InvitationOutput) UserType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Invitation) pulumi.StringPtrOutput { return v.UserType }).(pulumi.StringPtrOutput)
}

type InvitationArrayOutput struct{ *pulumi.OutputState }

func (InvitationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Invitation)(nil)).Elem()
}

func (o InvitationArrayOutput) ToInvitationArrayOutput() InvitationArrayOutput {
	return o
}

func (o InvitationArrayOutput) ToInvitationArrayOutputWithContext(ctx context.Context) InvitationArrayOutput {
	return o
}

func (o InvitationArrayOutput) Index(i pulumi.IntInput) InvitationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Invitation {
		return vs[0].([]*Invitation)[vs[1].(int)]
	}).(InvitationOutput)
}

type InvitationMapOutput struct{ *pulumi.OutputState }

func (InvitationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Invitation)(nil)).Elem()
}

func (o InvitationMapOutput) ToInvitationMapOutput() InvitationMapOutput {
	return o
}

func (o InvitationMapOutput) ToInvitationMapOutputWithContext(ctx context.Context) InvitationMapOutput {
	return o
}

func (o InvitationMapOutput) MapIndex(k pulumi.StringInput) InvitationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Invitation {
		return vs[0].(map[string]*Invitation)[vs[1].(string)]
	}).(InvitationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InvitationInput)(nil)).Elem(), &Invitation{})
	pulumi.RegisterInputType(reflect.TypeOf((*InvitationArrayInput)(nil)).Elem(), InvitationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InvitationMapInput)(nil)).Elem(), InvitationMap{})
	pulumi.RegisterOutputType(InvitationOutput{})
	pulumi.RegisterOutputType(InvitationArrayOutput{})
	pulumi.RegisterOutputType(InvitationMapOutput{})
}

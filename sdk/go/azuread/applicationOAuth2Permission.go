// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an OAuth2 Permission (also known as a Scope) associated with an Application within Azure Active Directory.
//
// > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuread/sdk/v3/go/azuread/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleApplication, err := azuread.NewApplication(ctx, "exampleApplication", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuread.NewApplicationOAuth2Permission(ctx, "exampleApplicationOAuth2Permission", &azuread.ApplicationOAuth2PermissionArgs{
// 			ApplicationObjectId:     exampleApplication.ID(),
// 			AdminConsentDescription: pulumi.String("Administer the application"),
// 			AdminConsentDisplayName: pulumi.String("Administer"),
// 			IsEnabled:               pulumi.Bool(true),
// 			Type:                    pulumi.String("User"),
// 			UserConsentDescription:  pulumi.String("Administer the application"),
// 			UserConsentDisplayName:  pulumi.String("Administer"),
// 			Value:                   pulumi.String("administer"),
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
// OAuth2 Permissions can be imported using the `object id` of an Application and the `id` of the Permission, e.g.
//
// ```sh
//  $ pulumi import azuread:index/applicationOAuth2Permission:ApplicationOAuth2Permission test 00000000-0000-0000-0000-000000000000/scope/11111111-1111-1111-1111-111111111111
// ```
type ApplicationOAuth2Permission struct {
	pulumi.CustomResourceState

	// Permission help text that appears in the admin consent and app assignment experiences.
	AdminConsentDescription pulumi.StringOutput `pulumi:"adminConsentDescription"`
	// Display name for the permission that appears in the admin consent and app assignment experiences.
	AdminConsentDisplayName pulumi.StringOutput `pulumi:"adminConsentDisplayName"`
	// The Object ID of the Application for which this Permission should be created. Changing this field forces a new resource to be created.
	ApplicationObjectId pulumi.StringOutput `pulumi:"applicationObjectId"`
	// Determines if the Permission is enabled. Defaults to `true`.
	IsEnabled pulumi.BoolPtrOutput `pulumi:"isEnabled"`
	// Specifies a custom UUID for the Permission. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
	PermissionId pulumi.StringOutput `pulumi:"permissionId"`
	// Specifies whether this scope permission can be consented to by an end user, or whether it is a tenant-wide permission that must be consented to by an Administrator. Possible values are "User" or "Admin".
	Type pulumi.StringOutput `pulumi:"type"`
	// Permission help text that appears in the end user consent experience.
	UserConsentDescription pulumi.StringOutput `pulumi:"userConsentDescription"`
	// Display name for the permission that appears in the end user consent experience.
	UserConsentDisplayName pulumi.StringOutput `pulumi:"userConsentDisplayName"`
	// The value of the scope claim that the resource application should expect in the OAuth 2.0 access token.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewApplicationOAuth2Permission registers a new resource with the given unique name, arguments, and options.
func NewApplicationOAuth2Permission(ctx *pulumi.Context,
	name string, args *ApplicationOAuth2PermissionArgs, opts ...pulumi.ResourceOption) (*ApplicationOAuth2Permission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminConsentDescription == nil {
		return nil, errors.New("invalid value for required argument 'AdminConsentDescription'")
	}
	if args.AdminConsentDisplayName == nil {
		return nil, errors.New("invalid value for required argument 'AdminConsentDisplayName'")
	}
	if args.ApplicationObjectId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationObjectId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.UserConsentDescription == nil {
		return nil, errors.New("invalid value for required argument 'UserConsentDescription'")
	}
	if args.UserConsentDisplayName == nil {
		return nil, errors.New("invalid value for required argument 'UserConsentDisplayName'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	var resource ApplicationOAuth2Permission
	err := ctx.RegisterResource("azuread:index/applicationOAuth2Permission:ApplicationOAuth2Permission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationOAuth2Permission gets an existing ApplicationOAuth2Permission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationOAuth2Permission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationOAuth2PermissionState, opts ...pulumi.ResourceOption) (*ApplicationOAuth2Permission, error) {
	var resource ApplicationOAuth2Permission
	err := ctx.ReadResource("azuread:index/applicationOAuth2Permission:ApplicationOAuth2Permission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationOAuth2Permission resources.
type applicationOAuth2PermissionState struct {
	// Permission help text that appears in the admin consent and app assignment experiences.
	AdminConsentDescription *string `pulumi:"adminConsentDescription"`
	// Display name for the permission that appears in the admin consent and app assignment experiences.
	AdminConsentDisplayName *string `pulumi:"adminConsentDisplayName"`
	// The Object ID of the Application for which this Permission should be created. Changing this field forces a new resource to be created.
	ApplicationObjectId *string `pulumi:"applicationObjectId"`
	// Determines if the Permission is enabled. Defaults to `true`.
	IsEnabled *bool `pulumi:"isEnabled"`
	// Specifies a custom UUID for the Permission. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
	PermissionId *string `pulumi:"permissionId"`
	// Specifies whether this scope permission can be consented to by an end user, or whether it is a tenant-wide permission that must be consented to by an Administrator. Possible values are "User" or "Admin".
	Type *string `pulumi:"type"`
	// Permission help text that appears in the end user consent experience.
	UserConsentDescription *string `pulumi:"userConsentDescription"`
	// Display name for the permission that appears in the end user consent experience.
	UserConsentDisplayName *string `pulumi:"userConsentDisplayName"`
	// The value of the scope claim that the resource application should expect in the OAuth 2.0 access token.
	Value *string `pulumi:"value"`
}

type ApplicationOAuth2PermissionState struct {
	// Permission help text that appears in the admin consent and app assignment experiences.
	AdminConsentDescription pulumi.StringPtrInput
	// Display name for the permission that appears in the admin consent and app assignment experiences.
	AdminConsentDisplayName pulumi.StringPtrInput
	// The Object ID of the Application for which this Permission should be created. Changing this field forces a new resource to be created.
	ApplicationObjectId pulumi.StringPtrInput
	// Determines if the Permission is enabled. Defaults to `true`.
	IsEnabled pulumi.BoolPtrInput
	// Specifies a custom UUID for the Permission. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
	PermissionId pulumi.StringPtrInput
	// Specifies whether this scope permission can be consented to by an end user, or whether it is a tenant-wide permission that must be consented to by an Administrator. Possible values are "User" or "Admin".
	Type pulumi.StringPtrInput
	// Permission help text that appears in the end user consent experience.
	UserConsentDescription pulumi.StringPtrInput
	// Display name for the permission that appears in the end user consent experience.
	UserConsentDisplayName pulumi.StringPtrInput
	// The value of the scope claim that the resource application should expect in the OAuth 2.0 access token.
	Value pulumi.StringPtrInput
}

func (ApplicationOAuth2PermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationOAuth2PermissionState)(nil)).Elem()
}

type applicationOAuth2PermissionArgs struct {
	// Permission help text that appears in the admin consent and app assignment experiences.
	AdminConsentDescription string `pulumi:"adminConsentDescription"`
	// Display name for the permission that appears in the admin consent and app assignment experiences.
	AdminConsentDisplayName string `pulumi:"adminConsentDisplayName"`
	// The Object ID of the Application for which this Permission should be created. Changing this field forces a new resource to be created.
	ApplicationObjectId string `pulumi:"applicationObjectId"`
	// Determines if the Permission is enabled. Defaults to `true`.
	IsEnabled *bool `pulumi:"isEnabled"`
	// Specifies a custom UUID for the Permission. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
	PermissionId *string `pulumi:"permissionId"`
	// Specifies whether this scope permission can be consented to by an end user, or whether it is a tenant-wide permission that must be consented to by an Administrator. Possible values are "User" or "Admin".
	Type string `pulumi:"type"`
	// Permission help text that appears in the end user consent experience.
	UserConsentDescription string `pulumi:"userConsentDescription"`
	// Display name for the permission that appears in the end user consent experience.
	UserConsentDisplayName string `pulumi:"userConsentDisplayName"`
	// The value of the scope claim that the resource application should expect in the OAuth 2.0 access token.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a ApplicationOAuth2Permission resource.
type ApplicationOAuth2PermissionArgs struct {
	// Permission help text that appears in the admin consent and app assignment experiences.
	AdminConsentDescription pulumi.StringInput
	// Display name for the permission that appears in the admin consent and app assignment experiences.
	AdminConsentDisplayName pulumi.StringInput
	// The Object ID of the Application for which this Permission should be created. Changing this field forces a new resource to be created.
	ApplicationObjectId pulumi.StringInput
	// Determines if the Permission is enabled. Defaults to `true`.
	IsEnabled pulumi.BoolPtrInput
	// Specifies a custom UUID for the Permission. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
	PermissionId pulumi.StringPtrInput
	// Specifies whether this scope permission can be consented to by an end user, or whether it is a tenant-wide permission that must be consented to by an Administrator. Possible values are "User" or "Admin".
	Type pulumi.StringInput
	// Permission help text that appears in the end user consent experience.
	UserConsentDescription pulumi.StringInput
	// Display name for the permission that appears in the end user consent experience.
	UserConsentDisplayName pulumi.StringInput
	// The value of the scope claim that the resource application should expect in the OAuth 2.0 access token.
	Value pulumi.StringInput
}

func (ApplicationOAuth2PermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationOAuth2PermissionArgs)(nil)).Elem()
}

type ApplicationOAuth2PermissionInput interface {
	pulumi.Input

	ToApplicationOAuth2PermissionOutput() ApplicationOAuth2PermissionOutput
	ToApplicationOAuth2PermissionOutputWithContext(ctx context.Context) ApplicationOAuth2PermissionOutput
}

func (*ApplicationOAuth2Permission) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationOAuth2Permission)(nil))
}

func (i *ApplicationOAuth2Permission) ToApplicationOAuth2PermissionOutput() ApplicationOAuth2PermissionOutput {
	return i.ToApplicationOAuth2PermissionOutputWithContext(context.Background())
}

func (i *ApplicationOAuth2Permission) ToApplicationOAuth2PermissionOutputWithContext(ctx context.Context) ApplicationOAuth2PermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOAuth2PermissionOutput)
}

type ApplicationOAuth2PermissionOutput struct {
	*pulumi.OutputState
}

func (ApplicationOAuth2PermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationOAuth2Permission)(nil))
}

func (o ApplicationOAuth2PermissionOutput) ToApplicationOAuth2PermissionOutput() ApplicationOAuth2PermissionOutput {
	return o
}

func (o ApplicationOAuth2PermissionOutput) ToApplicationOAuth2PermissionOutputWithContext(ctx context.Context) ApplicationOAuth2PermissionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApplicationOAuth2PermissionOutput{})
}

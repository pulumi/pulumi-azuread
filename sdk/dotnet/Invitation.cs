// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    /// <summary>
    /// Manages an invitation of a guest user within Azure Active Directory.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires one of the following application roles: `User.Invite.All`, `User.ReadWrite.All` or `Directory.ReadWrite.All`
    /// 
    /// When authenticated with a user principal, this resource requires one of the following directory roles: `Guest Inviter`, `User Administrator` or `Global Administrator`
    /// 
    /// ## Example Usage
    /// 
    /// *Basic example*
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureAD.Invitation("example", new()
    ///     {
    ///         RedirectUrl = "https://portal.azure.com",
    ///         UserEmailAddress = "jdoe@hashicorp.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// *Invitation with standard message*
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureAD.Invitation("example", new()
    ///     {
    ///         Message = new AzureAD.Inputs.InvitationMessageArgs
    ///         {
    ///             Language = "en-US",
    ///         },
    ///         RedirectUrl = "https://portal.azure.com",
    ///         UserEmailAddress = "jdoe@hashicorp.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// *Invitation with custom message body and an additional recipient*
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureAD.Invitation("example", new()
    ///     {
    ///         Message = new AzureAD.Inputs.InvitationMessageArgs
    ///         {
    ///             AdditionalRecipients = "aaliceberg@hashicorp.com",
    ///             Body = "Hello there! You are invited to join my Azure tenant!",
    ///         },
    ///         RedirectUrl = "https://portal.azure.com",
    ///         UserDisplayName = "Bob Bobson",
    ///         UserEmailAddress = "bbobson@hashicorp.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource does not support importing.
    /// </summary>
    [AzureADResourceType("azuread:index/invitation:Invitation")]
    public partial class Invitation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A `message` block as documented below, which configures the message being sent to the invited user. If this block is omitted, no message will be sent.
        /// </summary>
        [Output("message")]
        public Output<Outputs.InvitationMessage?> Message { get; private set; } = null!;

        /// <summary>
        /// The URL the user can use to redeem their invitation.
        /// </summary>
        [Output("redeemUrl")]
        public Output<string> RedeemUrl { get; private set; } = null!;

        /// <summary>
        /// The URL that the user should be redirected to once the invitation is redeemed.
        /// </summary>
        [Output("redirectUrl")]
        public Output<string> RedirectUrl { get; private set; } = null!;

        /// <summary>
        /// The display name of the user being invited.
        /// </summary>
        [Output("userDisplayName")]
        public Output<string?> UserDisplayName { get; private set; } = null!;

        /// <summary>
        /// The email address of the user being invited.
        /// </summary>
        [Output("userEmailAddress")]
        public Output<string> UserEmailAddress { get; private set; } = null!;

        /// <summary>
        /// Object ID of the invited user.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;

        /// <summary>
        /// The user type of the user being invited. Must be one of `Guest` or `Member`. Only Global Administrators can invite users as members. Defaults to `Guest`.
        /// </summary>
        [Output("userType")]
        public Output<string?> UserType { get; private set; } = null!;


        /// <summary>
        /// Create a Invitation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Invitation(string name, InvitationArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/invitation:Invitation", name, args ?? new InvitationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Invitation(string name, Input<string> id, InvitationState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/invitation:Invitation", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Invitation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Invitation Get(string name, Input<string> id, InvitationState? state = null, CustomResourceOptions? options = null)
        {
            return new Invitation(name, id, state, options);
        }
    }

    public sealed class InvitationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `message` block as documented below, which configures the message being sent to the invited user. If this block is omitted, no message will be sent.
        /// </summary>
        [Input("message")]
        public Input<Inputs.InvitationMessageArgs>? Message { get; set; }

        /// <summary>
        /// The URL that the user should be redirected to once the invitation is redeemed.
        /// </summary>
        [Input("redirectUrl", required: true)]
        public Input<string> RedirectUrl { get; set; } = null!;

        /// <summary>
        /// The display name of the user being invited.
        /// </summary>
        [Input("userDisplayName")]
        public Input<string>? UserDisplayName { get; set; }

        /// <summary>
        /// The email address of the user being invited.
        /// </summary>
        [Input("userEmailAddress", required: true)]
        public Input<string> UserEmailAddress { get; set; } = null!;

        /// <summary>
        /// The user type of the user being invited. Must be one of `Guest` or `Member`. Only Global Administrators can invite users as members. Defaults to `Guest`.
        /// </summary>
        [Input("userType")]
        public Input<string>? UserType { get; set; }

        public InvitationArgs()
        {
        }
        public static new InvitationArgs Empty => new InvitationArgs();
    }

    public sealed class InvitationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `message` block as documented below, which configures the message being sent to the invited user. If this block is omitted, no message will be sent.
        /// </summary>
        [Input("message")]
        public Input<Inputs.InvitationMessageGetArgs>? Message { get; set; }

        /// <summary>
        /// The URL the user can use to redeem their invitation.
        /// </summary>
        [Input("redeemUrl")]
        public Input<string>? RedeemUrl { get; set; }

        /// <summary>
        /// The URL that the user should be redirected to once the invitation is redeemed.
        /// </summary>
        [Input("redirectUrl")]
        public Input<string>? RedirectUrl { get; set; }

        /// <summary>
        /// The display name of the user being invited.
        /// </summary>
        [Input("userDisplayName")]
        public Input<string>? UserDisplayName { get; set; }

        /// <summary>
        /// The email address of the user being invited.
        /// </summary>
        [Input("userEmailAddress")]
        public Input<string>? UserEmailAddress { get; set; }

        /// <summary>
        /// Object ID of the invited user.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        /// <summary>
        /// The user type of the user being invited. Must be one of `Guest` or `Member`. Only Global Administrators can invite users as members. Defaults to `Guest`.
        /// </summary>
        [Input("userType")]
        public Input<string>? UserType { get; set; }

        public InvitationState()
        {
        }
        public static new InvitationState Empty => new InvitationState();
    }
}

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
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleApplication = new AzureAD.Application("exampleApplication", new AzureAD.ApplicationArgs
    ///         {
    ///             DisplayName = "example",
    ///         });
    ///         var exampleApplicationFederatedIdentityCredential = new AzureAD.ApplicationFederatedIdentityCredential("exampleApplicationFederatedIdentityCredential", new AzureAD.ApplicationFederatedIdentityCredentialArgs
    ///         {
    ///             ApplicationObjectId = exampleApplication.ObjectId,
    ///             DisplayName = "my-repo-deploy",
    ///             Description = "Deployments for my-repo",
    ///             Audiences = 
    ///             {
    ///                 "api://AzureADTokenExchange",
    ///             },
    ///             Issuer = "https://token.actions.githubusercontent.com",
    ///             Subject = "repo:my-organization/my-repo:environment:prod",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Federated Identity Credentials can be imported using the object ID of the associated application and the ID of the federated identity credential, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azuread:index/applicationFederatedIdentityCredential:ApplicationFederatedIdentityCredential test 00000000-0000-0000-0000-000000000000/federatedIdentityCredential/11111111-1111-1111-1111-111111111111
    /// ```
    /// 
    ///  -&gt; This ID format is unique to Terraform and is composed of the application's object ID, the string "federatedIdentityCredential" and the credential ID in the format `{ObjectId}/federatedIdentityCredential/{CredentialId}`.
    /// </summary>
    [AzureADResourceType("azuread:index/applicationFederatedIdentityCredential:ApplicationFederatedIdentityCredential")]
    public partial class ApplicationFederatedIdentityCredential : Pulumi.CustomResource
    {
        /// <summary>
        /// The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
        /// </summary>
        [Output("applicationObjectId")]
        public Output<string> ApplicationObjectId { get; private set; } = null!;

        /// <summary>
        /// List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
        /// </summary>
        [Output("audiences")]
        public Output<ImmutableArray<string>> Audiences { get; private set; } = null!;

        /// <summary>
        /// A UUID used to uniquely identify this federated identity credential.
        /// </summary>
        [Output("credentialId")]
        public Output<string> CredentialId { get; private set; } = null!;

        /// <summary>
        /// A description for the federated identity credential.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A unique display name for the federated identity credential. Changing this forces a new resource to be created.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
        /// </summary>
        [Output("issuer")]
        public Output<string> Issuer { get; private set; } = null!;

        /// <summary>
        /// The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
        /// </summary>
        [Output("subject")]
        public Output<string> Subject { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationFederatedIdentityCredential resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationFederatedIdentityCredential(string name, ApplicationFederatedIdentityCredentialArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/applicationFederatedIdentityCredential:ApplicationFederatedIdentityCredential", name, args ?? new ApplicationFederatedIdentityCredentialArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationFederatedIdentityCredential(string name, Input<string> id, ApplicationFederatedIdentityCredentialState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/applicationFederatedIdentityCredential:ApplicationFederatedIdentityCredential", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationFederatedIdentityCredential resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationFederatedIdentityCredential Get(string name, Input<string> id, ApplicationFederatedIdentityCredentialState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationFederatedIdentityCredential(name, id, state, options);
        }
    }

    public sealed class ApplicationFederatedIdentityCredentialArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("applicationObjectId", required: true)]
        public Input<string> ApplicationObjectId { get; set; } = null!;

        [Input("audiences", required: true)]
        private InputList<string>? _audiences;

        /// <summary>
        /// List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
        /// </summary>
        public InputList<string> Audiences
        {
            get => _audiences ?? (_audiences = new InputList<string>());
            set => _audiences = value;
        }

        /// <summary>
        /// A description for the federated identity credential.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A unique display name for the federated identity credential. Changing this forces a new resource to be created.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
        /// </summary>
        [Input("issuer", required: true)]
        public Input<string> Issuer { get; set; } = null!;

        /// <summary>
        /// The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
        /// </summary>
        [Input("subject", required: true)]
        public Input<string> Subject { get; set; } = null!;

        public ApplicationFederatedIdentityCredentialArgs()
        {
        }
    }

    public sealed class ApplicationFederatedIdentityCredentialState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("applicationObjectId")]
        public Input<string>? ApplicationObjectId { get; set; }

        [Input("audiences")]
        private InputList<string>? _audiences;

        /// <summary>
        /// List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
        /// </summary>
        public InputList<string> Audiences
        {
            get => _audiences ?? (_audiences = new InputList<string>());
            set => _audiences = value;
        }

        /// <summary>
        /// A UUID used to uniquely identify this federated identity credential.
        /// </summary>
        [Input("credentialId")]
        public Input<string>? CredentialId { get; set; }

        /// <summary>
        /// A description for the federated identity credential.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A unique display name for the federated identity credential. Changing this forces a new resource to be created.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
        /// </summary>
        [Input("subject")]
        public Input<string>? Subject { get; set; }

        public ApplicationFederatedIdentityCredentialState()
        {
        }
    }
}

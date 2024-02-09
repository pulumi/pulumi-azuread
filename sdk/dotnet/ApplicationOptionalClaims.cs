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
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleApplicationRegistration = new AzureAD.ApplicationRegistration("exampleApplicationRegistration", new()
    ///     {
    ///         DisplayName = "example",
    ///     });
    /// 
    ///     var exampleApplicationOptionalClaims = new AzureAD.ApplicationOptionalClaims("exampleApplicationOptionalClaims", new()
    ///     {
    ///         ApplicationId = exampleApplicationRegistration.Id,
    ///         AccessTokens = new[]
    ///         {
    ///             new AzureAD.Inputs.ApplicationOptionalClaimsAccessTokenArgs
    ///             {
    ///                 Name = "myclaim",
    ///             },
    ///             new AzureAD.Inputs.ApplicationOptionalClaimsAccessTokenArgs
    ///             {
    ///                 Name = "otherclaim",
    ///             },
    ///         },
    ///         IdTokens = new[]
    ///         {
    ///             new AzureAD.Inputs.ApplicationOptionalClaimsIdTokenArgs
    ///             {
    ///                 Name = "userclaim",
    ///                 Source = "user",
    ///                 Essential = true,
    ///                 AdditionalProperties = new[]
    ///                 {
    ///                     "emit_as_roles",
    ///                 },
    ///             },
    ///         },
    ///         Saml2Tokens = new[]
    ///         {
    ///             new AzureAD.Inputs.ApplicationOptionalClaimsSaml2TokenArgs
    ///             {
    ///                 Name = "samlexample",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Application Optional Claims can be imported using the object ID of the application, in the following format.
    /// 
    /// ```sh
    /// $ pulumi import azuread:index/applicationOptionalClaims:ApplicationOptionalClaims example /applications/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureADResourceType("azuread:index/applicationOptionalClaims:ApplicationOptionalClaims")]
    public partial class ApplicationOptionalClaims : global::Pulumi.CustomResource
    {
        /// <summary>
        /// One or more `access_token` blocks as documented below.
        /// </summary>
        [Output("accessTokens")]
        public Output<ImmutableArray<Outputs.ApplicationOptionalClaimsAccessToken>> AccessTokens { get; private set; } = null!;

        /// <summary>
        /// The resource ID of the application registration. Changing this forces a new resource to be created.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// One or more `id_token` blocks as documented below.
        /// </summary>
        [Output("idTokens")]
        public Output<ImmutableArray<Outputs.ApplicationOptionalClaimsIdToken>> IdTokens { get; private set; } = null!;

        /// <summary>
        /// One or more `saml2_token` blocks as documented below.
        /// 
        /// &gt; At least one of `access_token`, `id_token` or `saml2_token` must be specified
        /// </summary>
        [Output("saml2Tokens")]
        public Output<ImmutableArray<Outputs.ApplicationOptionalClaimsSaml2Token>> Saml2Tokens { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationOptionalClaims resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationOptionalClaims(string name, ApplicationOptionalClaimsArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/applicationOptionalClaims:ApplicationOptionalClaims", name, args ?? new ApplicationOptionalClaimsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationOptionalClaims(string name, Input<string> id, ApplicationOptionalClaimsState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/applicationOptionalClaims:ApplicationOptionalClaims", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationOptionalClaims resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationOptionalClaims Get(string name, Input<string> id, ApplicationOptionalClaimsState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationOptionalClaims(name, id, state, options);
        }
    }

    public sealed class ApplicationOptionalClaimsArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessTokens")]
        private InputList<Inputs.ApplicationOptionalClaimsAccessTokenArgs>? _accessTokens;

        /// <summary>
        /// One or more `access_token` blocks as documented below.
        /// </summary>
        public InputList<Inputs.ApplicationOptionalClaimsAccessTokenArgs> AccessTokens
        {
            get => _accessTokens ?? (_accessTokens = new InputList<Inputs.ApplicationOptionalClaimsAccessTokenArgs>());
            set => _accessTokens = value;
        }

        /// <summary>
        /// The resource ID of the application registration. Changing this forces a new resource to be created.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        [Input("idTokens")]
        private InputList<Inputs.ApplicationOptionalClaimsIdTokenArgs>? _idTokens;

        /// <summary>
        /// One or more `id_token` blocks as documented below.
        /// </summary>
        public InputList<Inputs.ApplicationOptionalClaimsIdTokenArgs> IdTokens
        {
            get => _idTokens ?? (_idTokens = new InputList<Inputs.ApplicationOptionalClaimsIdTokenArgs>());
            set => _idTokens = value;
        }

        [Input("saml2Tokens")]
        private InputList<Inputs.ApplicationOptionalClaimsSaml2TokenArgs>? _saml2Tokens;

        /// <summary>
        /// One or more `saml2_token` blocks as documented below.
        /// 
        /// &gt; At least one of `access_token`, `id_token` or `saml2_token` must be specified
        /// </summary>
        public InputList<Inputs.ApplicationOptionalClaimsSaml2TokenArgs> Saml2Tokens
        {
            get => _saml2Tokens ?? (_saml2Tokens = new InputList<Inputs.ApplicationOptionalClaimsSaml2TokenArgs>());
            set => _saml2Tokens = value;
        }

        public ApplicationOptionalClaimsArgs()
        {
        }
        public static new ApplicationOptionalClaimsArgs Empty => new ApplicationOptionalClaimsArgs();
    }

    public sealed class ApplicationOptionalClaimsState : global::Pulumi.ResourceArgs
    {
        [Input("accessTokens")]
        private InputList<Inputs.ApplicationOptionalClaimsAccessTokenGetArgs>? _accessTokens;

        /// <summary>
        /// One or more `access_token` blocks as documented below.
        /// </summary>
        public InputList<Inputs.ApplicationOptionalClaimsAccessTokenGetArgs> AccessTokens
        {
            get => _accessTokens ?? (_accessTokens = new InputList<Inputs.ApplicationOptionalClaimsAccessTokenGetArgs>());
            set => _accessTokens = value;
        }

        /// <summary>
        /// The resource ID of the application registration. Changing this forces a new resource to be created.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        [Input("idTokens")]
        private InputList<Inputs.ApplicationOptionalClaimsIdTokenGetArgs>? _idTokens;

        /// <summary>
        /// One or more `id_token` blocks as documented below.
        /// </summary>
        public InputList<Inputs.ApplicationOptionalClaimsIdTokenGetArgs> IdTokens
        {
            get => _idTokens ?? (_idTokens = new InputList<Inputs.ApplicationOptionalClaimsIdTokenGetArgs>());
            set => _idTokens = value;
        }

        [Input("saml2Tokens")]
        private InputList<Inputs.ApplicationOptionalClaimsSaml2TokenGetArgs>? _saml2Tokens;

        /// <summary>
        /// One or more `saml2_token` blocks as documented below.
        /// 
        /// &gt; At least one of `access_token`, `id_token` or `saml2_token` must be specified
        /// </summary>
        public InputList<Inputs.ApplicationOptionalClaimsSaml2TokenGetArgs> Saml2Tokens
        {
            get => _saml2Tokens ?? (_saml2Tokens = new InputList<Inputs.ApplicationOptionalClaimsSaml2TokenGetArgs>());
            set => _saml2Tokens = value;
        }

        public ApplicationOptionalClaimsState()
        {
        }
        public static new ApplicationOptionalClaimsState Empty => new ApplicationOptionalClaimsState();
    }
}

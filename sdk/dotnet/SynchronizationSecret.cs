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
    /// Manages synchronization secrets associated with a service principal (enterprise application) within Azure Active Directory.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.All` or `Directory.ReadWrite.All`
    /// 
    /// ## Example Usage
    /// 
    /// *Basic example*
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = AzureAD.GetApplicationTemplate.Invoke(new()
    ///     {
    ///         DisplayName = "Azure Databricks SCIM Provisioning Connector",
    ///     });
    /// 
    ///     var exampleApplicationFromTemplate = new AzureAD.ApplicationFromTemplate("example", new()
    ///     {
    ///         DisplayName = "example",
    ///         TemplateId = example.Apply(getApplicationTemplateResult =&gt; getApplicationTemplateResult.TemplateId),
    ///     });
    /// 
    ///     var exampleGetServicePrincipal = AzureAD.GetServicePrincipal.Invoke(new()
    ///     {
    ///         ObjectId = exampleApplicationFromTemplate.ServicePrincipalObjectId,
    ///     });
    /// 
    ///     var exampleSynchronizationSecret = new AzureAD.SynchronizationSecret("example", new()
    ///     {
    ///         ServicePrincipalId = exampleGetServicePrincipal.Apply(getServicePrincipalResult =&gt; getServicePrincipalResult.Id),
    ///         Credentials = new[]
    ///         {
    ///             new AzureAD.Inputs.SynchronizationSecretCredentialArgs
    ///             {
    ///                 Key = "BaseAddress",
    ///                 Value = "abc",
    ///             },
    ///             new AzureAD.Inputs.SynchronizationSecretCredentialArgs
    ///             {
    ///                 Key = "SecretToken",
    ///                 Value = "some-token",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource does not support importing.
    /// </summary>
    [AzureADResourceType("azuread:index/synchronizationSecret:SynchronizationSecret")]
    public partial class SynchronizationSecret : global::Pulumi.CustomResource
    {
        /// <summary>
        /// One or more `credential` blocks as documented below.
        /// </summary>
        [Output("credentials")]
        public Output<ImmutableArray<Outputs.SynchronizationSecretCredential>> Credentials { get; private set; } = null!;

        /// <summary>
        /// The ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
        /// </summary>
        [Output("servicePrincipalId")]
        public Output<string> ServicePrincipalId { get; private set; } = null!;


        /// <summary>
        /// Create a SynchronizationSecret resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SynchronizationSecret(string name, SynchronizationSecretArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/synchronizationSecret:SynchronizationSecret", name, args ?? new SynchronizationSecretArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SynchronizationSecret(string name, Input<string> id, SynchronizationSecretState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/synchronizationSecret:SynchronizationSecret", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SynchronizationSecret resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SynchronizationSecret Get(string name, Input<string> id, SynchronizationSecretState? state = null, CustomResourceOptions? options = null)
        {
            return new SynchronizationSecret(name, id, state, options);
        }
    }

    public sealed class SynchronizationSecretArgs : global::Pulumi.ResourceArgs
    {
        [Input("credentials")]
        private InputList<Inputs.SynchronizationSecretCredentialArgs>? _credentials;

        /// <summary>
        /// One or more `credential` blocks as documented below.
        /// </summary>
        public InputList<Inputs.SynchronizationSecretCredentialArgs> Credentials
        {
            get => _credentials ?? (_credentials = new InputList<Inputs.SynchronizationSecretCredentialArgs>());
            set => _credentials = value;
        }

        /// <summary>
        /// The ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("servicePrincipalId", required: true)]
        public Input<string> ServicePrincipalId { get; set; } = null!;

        public SynchronizationSecretArgs()
        {
        }
        public static new SynchronizationSecretArgs Empty => new SynchronizationSecretArgs();
    }

    public sealed class SynchronizationSecretState : global::Pulumi.ResourceArgs
    {
        [Input("credentials")]
        private InputList<Inputs.SynchronizationSecretCredentialGetArgs>? _credentials;

        /// <summary>
        /// One or more `credential` blocks as documented below.
        /// </summary>
        public InputList<Inputs.SynchronizationSecretCredentialGetArgs> Credentials
        {
            get => _credentials ?? (_credentials = new InputList<Inputs.SynchronizationSecretCredentialGetArgs>());
            set => _credentials = value;
        }

        /// <summary>
        /// The ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("servicePrincipalId")]
        public Input<string>? ServicePrincipalId { get; set; }

        public SynchronizationSecretState()
        {
        }
        public static new SynchronizationSecretState Empty => new SynchronizationSecretState();
    }
}

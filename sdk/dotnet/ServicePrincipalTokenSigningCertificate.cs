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
    /// *Using default settings*
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureAD.Application("example", new()
    ///     {
    ///         DisplayName = "example",
    ///     });
    /// 
    ///     var exampleServicePrincipal = new AzureAD.ServicePrincipal("example", new()
    ///     {
    ///         ClientId = example.ClientId,
    ///     });
    /// 
    ///     var exampleServicePrincipalTokenSigningCertificate = new AzureAD.ServicePrincipalTokenSigningCertificate("example", new()
    ///     {
    ///         ServicePrincipalId = exampleServicePrincipal.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// *Using custom settings*
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureAD.Application("example", new()
    ///     {
    ///         DisplayName = "example",
    ///     });
    /// 
    ///     var exampleServicePrincipal = new AzureAD.ServicePrincipal("example", new()
    ///     {
    ///         ClientId = example.ClientId,
    ///     });
    /// 
    ///     var exampleServicePrincipalTokenSigningCertificate = new AzureAD.ServicePrincipalTokenSigningCertificate("example", new()
    ///     {
    ///         ServicePrincipalId = exampleServicePrincipal.Id,
    ///         DisplayName = "CN=example.com",
    ///         EndDate = "2023-05-01T01:02:03Z",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Token signing certificates can be imported using the object ID of the associated service principal and the key ID of the verify certificate credential, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuread:index/servicePrincipalTokenSigningCertificate:ServicePrincipalTokenSigningCertificate example 00000000-0000-0000-0000-000000000000/tokenSigningCertificate/11111111-1111-1111-1111-111111111111
    /// ```
    /// 
    /// -&gt; This ID format is unique to Terraform and is composed of the service principal's object ID, the string "tokenSigningCertificate" and the verify certificate's key ID in the format `{ServicePrincipalObjectId}/tokenSigningCertificate/{CertificateKeyId}`.
    /// </summary>
    [AzureADResourceType("azuread:index/servicePrincipalTokenSigningCertificate:ServicePrincipalTokenSigningCertificate")]
    public partial class ServicePrincipalTokenSigningCertificate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies a friendly name for the certificate. Must start with `CN=`. Changing this field forces a new resource to be created.
        /// 
        /// &gt; If not specified, it will default to `CN=Microsoft Azure Federated SSO Certificate`.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        /// </summary>
        [Output("endDate")]
        public Output<string> EndDate { get; private set; } = null!;

        /// <summary>
        /// A UUID used to uniquely identify the verify certificate.
        /// </summary>
        [Output("keyId")]
        public Output<string> KeyId { get; private set; } = null!;

        /// <summary>
        /// The ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
        /// </summary>
        [Output("servicePrincipalId")]
        public Output<string> ServicePrincipalId { get; private set; } = null!;

        /// <summary>
        /// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
        /// </summary>
        [Output("startDate")]
        public Output<string> StartDate { get; private set; } = null!;

        /// <summary>
        /// A SHA-1 generated thumbprint of the token signing certificate, which can be used to set the preferred signing certificate for a service principal.
        /// </summary>
        [Output("thumbprint")]
        public Output<string> Thumbprint { get; private set; } = null!;

        /// <summary>
        /// The certificate data, which is PEM encoded but does not include the header `-----BEGIN CERTIFICATE-----\n` or the footer `\n-----END CERTIFICATE-----`.
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;


        /// <summary>
        /// Create a ServicePrincipalTokenSigningCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServicePrincipalTokenSigningCertificate(string name, ServicePrincipalTokenSigningCertificateArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/servicePrincipalTokenSigningCertificate:ServicePrincipalTokenSigningCertificate", name, args ?? new ServicePrincipalTokenSigningCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServicePrincipalTokenSigningCertificate(string name, Input<string> id, ServicePrincipalTokenSigningCertificateState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/servicePrincipalTokenSigningCertificate:ServicePrincipalTokenSigningCertificate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "value",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServicePrincipalTokenSigningCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServicePrincipalTokenSigningCertificate Get(string name, Input<string> id, ServicePrincipalTokenSigningCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new ServicePrincipalTokenSigningCertificate(name, id, state, options);
        }
    }

    public sealed class ServicePrincipalTokenSigningCertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies a friendly name for the certificate. Must start with `CN=`. Changing this field forces a new resource to be created.
        /// 
        /// &gt; If not specified, it will default to `CN=Microsoft Azure Federated SSO Certificate`.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        /// </summary>
        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        /// <summary>
        /// The ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("servicePrincipalId", required: true)]
        public Input<string> ServicePrincipalId { get; set; } = null!;

        public ServicePrincipalTokenSigningCertificateArgs()
        {
        }
        public static new ServicePrincipalTokenSigningCertificateArgs Empty => new ServicePrincipalTokenSigningCertificateArgs();
    }

    public sealed class ServicePrincipalTokenSigningCertificateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies a friendly name for the certificate. Must start with `CN=`. Changing this field forces a new resource to be created.
        /// 
        /// &gt; If not specified, it will default to `CN=Microsoft Azure Federated SSO Certificate`.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        /// </summary>
        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        /// <summary>
        /// A UUID used to uniquely identify the verify certificate.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// The ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("servicePrincipalId")]
        public Input<string>? ServicePrincipalId { get; set; }

        /// <summary>
        /// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        /// <summary>
        /// A SHA-1 generated thumbprint of the token signing certificate, which can be used to set the preferred signing certificate for a service principal.
        /// </summary>
        [Input("thumbprint")]
        public Input<string>? Thumbprint { get; set; }

        [Input("value")]
        private Input<string>? _value;

        /// <summary>
        /// The certificate data, which is PEM encoded but does not include the header `-----BEGIN CERTIFICATE-----\n` or the footer `\n-----END CERTIFICATE-----`.
        /// </summary>
        public Input<string>? Value
        {
            get => _value;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _value = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ServicePrincipalTokenSigningCertificateState()
        {
        }
        public static new ServicePrincipalTokenSigningCertificateState Empty => new ServicePrincipalTokenSigningCertificateState();
    }
}

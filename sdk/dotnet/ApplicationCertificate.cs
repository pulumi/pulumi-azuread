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
    /// ## Import
    /// 
    /// Certificates can be imported using the object ID of the associated application and the key ID of the certificate credential, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azuread:index/applicationCertificate:ApplicationCertificate test 00000000-0000-0000-0000-000000000000/certificate/11111111-1111-1111-1111-111111111111
    /// ```
    /// 
    ///  -&gt; This ID format is unique to Terraform and is composed of the application's object ID, the string "certificate" and the certificate's key ID in the format `{ObjectId}/certificate/{CertificateKeyId}`.
    /// </summary>
    [AzureADResourceType("azuread:index/applicationCertificate:ApplicationCertificate")]
    public partial class ApplicationCertificate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The object ID of the application for which this certificate should be created. Changing this field forces a new resource to be created.
        /// </summary>
        [Output("applicationObjectId")]
        public Output<string> ApplicationObjectId { get; private set; } = null!;

        /// <summary>
        /// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
        /// 
        /// &gt; **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurerm_key_vault_certificate resource.
        /// </summary>
        [Output("encoding")]
        public Output<string?> Encoding { get; private set; } = null!;

        /// <summary>
        /// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.
        /// </summary>
        [Output("endDate")]
        public Output<string> EndDate { get; private set; } = null!;

        /// <summary>
        /// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        /// 
        /// &gt; One of `end_date` or `end_date_relative` must be set. The maximum allowed duration is determined by Azure AD.
        /// </summary>
        [Output("endDateRelative")]
        public Output<string?> EndDateRelative { get; private set; } = null!;

        /// <summary>
        /// A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
        /// </summary>
        [Output("keyId")]
        public Output<string> KeyId { get; private set; } = null!;

        /// <summary>
        /// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
        /// </summary>
        [Output("startDate")]
        public Output<string> StartDate { get; private set; } = null!;

        /// <summary>
        /// The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        /// <summary>
        /// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationCertificate(string name, ApplicationCertificateArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/applicationCertificate:ApplicationCertificate", name, args ?? new ApplicationCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationCertificate(string name, Input<string> id, ApplicationCertificateState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/applicationCertificate:ApplicationCertificate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationCertificate Get(string name, Input<string> id, ApplicationCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationCertificate(name, id, state, options);
        }
    }

    public sealed class ApplicationCertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The object ID of the application for which this certificate should be created. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("applicationObjectId", required: true)]
        public Input<string> ApplicationObjectId { get; set; } = null!;

        /// <summary>
        /// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
        /// 
        /// &gt; **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurerm_key_vault_certificate resource.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        /// <summary>
        /// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        /// <summary>
        /// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        /// 
        /// &gt; One of `end_date` or `end_date_relative` must be set. The maximum allowed duration is determined by Azure AD.
        /// </summary>
        [Input("endDateRelative")]
        public Input<string>? EndDateRelative { get; set; }

        /// <summary>
        /// A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        /// <summary>
        /// The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("value", required: true)]
        private Input<string>? _value;

        /// <summary>
        /// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
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

        public ApplicationCertificateArgs()
        {
        }
        public static new ApplicationCertificateArgs Empty => new ApplicationCertificateArgs();
    }

    public sealed class ApplicationCertificateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The object ID of the application for which this certificate should be created. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("applicationObjectId")]
        public Input<string>? ApplicationObjectId { get; set; }

        /// <summary>
        /// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
        /// 
        /// &gt; **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurerm_key_vault_certificate resource.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        /// <summary>
        /// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If omitted, the API will decide a suitable expiry date, which is typically around 2 years from the start date. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        /// <summary>
        /// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        /// 
        /// &gt; One of `end_date` or `end_date_relative` must be set. The maximum allowed duration is determined by Azure AD.
        /// </summary>
        [Input("endDateRelative")]
        public Input<string>? EndDateRelative { get; set; }

        /// <summary>
        /// A UUID used to uniquely identify this certificate. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        /// <summary>
        /// The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("value")]
        private Input<string>? _value;

        /// <summary>
        /// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
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

        public ApplicationCertificateState()
        {
        }
        public static new ApplicationCertificateState Empty => new ApplicationCertificateState();
    }
}

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
    /// Manages a Certificate associated with an Application within Azure Active Directory.
    /// 
    /// &gt; **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.
    /// 
    /// ## Import
    /// 
    /// Certificates can be imported using the `object id` of an Application and the `key id` of the certificate, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azuread:index/applicationCertificate:ApplicationCertificate test 00000000-0000-0000-0000-000000000000/certificate/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [AzureADResourceType("azuread:index/applicationCertificate:ApplicationCertificate")]
    public partial class ApplicationCertificate : Pulumi.CustomResource
    {
        /// <summary>
        /// The Object ID of the Application for which this Certificate should be created. Changing this field forces a new resource to be created.
        /// </summary>
        [Output("applicationObjectId")]
        public Output<string> ApplicationObjectId { get; private set; } = null!;

        /// <summary>
        /// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
        /// </summary>
        [Output("encoding")]
        public Output<string?> Encoding { get; private set; } = null!;

        /// <summary>
        /// The End Date which the Certificate is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        /// </summary>
        [Output("endDate")]
        public Output<string> EndDate { get; private set; } = null!;

        /// <summary>
        /// A relative duration for which the Certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        /// </summary>
        [Output("endDateRelative")]
        public Output<string?> EndDateRelative { get; private set; } = null!;

        /// <summary>
        /// A GUID used to uniquely identify this Certificate. If not specified a GUID will be created. Changing this field forces a new resource to be created.
        /// </summary>
        [Output("keyId")]
        public Output<string> KeyId { get; private set; } = null!;

        /// <summary>
        /// The Start Date which the Certificate is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
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

    public sealed class ApplicationCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Object ID of the Application for which this Certificate should be created. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("applicationObjectId", required: true)]
        public Input<string> ApplicationObjectId { get; set; } = null!;

        /// <summary>
        /// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        /// <summary>
        /// The End Date which the Certificate is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        /// </summary>
        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        /// <summary>
        /// A relative duration for which the Certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("endDateRelative")]
        public Input<string>? EndDateRelative { get; set; }

        /// <summary>
        /// A GUID used to uniquely identify this Certificate. If not specified a GUID will be created. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// The Start Date which the Certificate is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        /// <summary>
        /// The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ApplicationCertificateArgs()
        {
        }
    }

    public sealed class ApplicationCertificateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Object ID of the Application for which this Certificate should be created. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("applicationObjectId")]
        public Input<string>? ApplicationObjectId { get; set; }

        /// <summary>
        /// Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        /// <summary>
        /// The End Date which the Certificate is valid until, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        /// </summary>
        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        /// <summary>
        /// A relative duration for which the Certificate is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("endDateRelative")]
        public Input<string>? EndDateRelative { get; set; }

        /// <summary>
        /// A GUID used to uniquely identify this Certificate. If not specified a GUID will be created. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// The Start Date which the Certificate is valid from, formatted as a RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        /// <summary>
        /// The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public ApplicationCertificateState()
        {
        }
    }
}

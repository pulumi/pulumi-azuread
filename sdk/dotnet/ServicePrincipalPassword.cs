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
    /// Manages a password credential associated with a service principal within Azure Active Directory. See also the azuread.ApplicationPassword resource.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.All` or `Directory.ReadWrite.All`
    /// 
    /// When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`
    /// 
    /// ## Import
    /// 
    /// This resource does not support importing.
    /// </summary>
    [AzureADResourceType("azuread:index/servicePrincipalPassword:ServicePrincipalPassword")]
    public partial class ServicePrincipalPassword : Pulumi.CustomResource
    {
        /// <summary>
        /// A display name for the password.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        /// </summary>
        [Output("endDate")]
        public Output<string> EndDate { get; private set; } = null!;

        /// <summary>
        /// A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        /// </summary>
        [Output("endDateRelative")]
        public Output<string?> EndDateRelative { get; private set; } = null!;

        /// <summary>
        /// A UUID used to uniquely identify this password credential.
        /// </summary>
        [Output("keyId")]
        public Output<string> KeyId { get; private set; } = null!;

        /// <summary>
        /// A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
        /// </summary>
        [Output("rotateWhenChanged")]
        public Output<ImmutableDictionary<string, string>?> RotateWhenChanged { get; private set; } = null!;

        /// <summary>
        /// The object ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
        /// </summary>
        [Output("servicePrincipalId")]
        public Output<string> ServicePrincipalId { get; private set; } = null!;

        /// <summary>
        /// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        /// </summary>
        [Output("startDate")]
        public Output<string> StartDate { get; private set; } = null!;

        /// <summary>
        /// The password for this service principal, which is generated by Azure Active Directory.
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;


        /// <summary>
        /// Create a ServicePrincipalPassword resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServicePrincipalPassword(string name, ServicePrincipalPasswordArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/servicePrincipalPassword:ServicePrincipalPassword", name, args ?? new ServicePrincipalPasswordArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServicePrincipalPassword(string name, Input<string> id, ServicePrincipalPasswordState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/servicePrincipalPassword:ServicePrincipalPassword", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServicePrincipalPassword resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServicePrincipalPassword Get(string name, Input<string> id, ServicePrincipalPasswordState? state = null, CustomResourceOptions? options = null)
        {
            return new ServicePrincipalPassword(name, id, state, options);
        }
    }

    public sealed class ServicePrincipalPasswordArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A display name for the password.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        /// </summary>
        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        /// <summary>
        /// A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("endDateRelative")]
        public Input<string>? EndDateRelative { get; set; }

        [Input("rotateWhenChanged")]
        private InputMap<string>? _rotateWhenChanged;

        /// <summary>
        /// A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
        /// </summary>
        public InputMap<string> RotateWhenChanged
        {
            get => _rotateWhenChanged ?? (_rotateWhenChanged = new InputMap<string>());
            set => _rotateWhenChanged = value;
        }

        /// <summary>
        /// The object ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("servicePrincipalId", required: true)]
        public Input<string> ServicePrincipalId { get; set; } = null!;

        /// <summary>
        /// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        public ServicePrincipalPasswordArgs()
        {
        }
    }

    public sealed class ServicePrincipalPasswordState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A display name for the password.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        /// </summary>
        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        /// <summary>
        /// A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("endDateRelative")]
        public Input<string>? EndDateRelative { get; set; }

        /// <summary>
        /// A UUID used to uniquely identify this password credential.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        [Input("rotateWhenChanged")]
        private InputMap<string>? _rotateWhenChanged;

        /// <summary>
        /// A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
        /// </summary>
        public InputMap<string> RotateWhenChanged
        {
            get => _rotateWhenChanged ?? (_rotateWhenChanged = new InputMap<string>());
            set => _rotateWhenChanged = value;
        }

        /// <summary>
        /// The object ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
        /// </summary>
        [Input("servicePrincipalId")]
        public Input<string>? ServicePrincipalId { get; set; }

        /// <summary>
        /// The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        /// <summary>
        /// The password for this service principal, which is generated by Azure Active Directory.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public ServicePrincipalPasswordState()
        {
        }
    }
}

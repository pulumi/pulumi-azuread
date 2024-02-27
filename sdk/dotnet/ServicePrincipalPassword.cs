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
    /// *Basic example*
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Azuread = Pulumi.Azuread;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Azuread.Index.Application.Application("example", new()
    ///     {
    ///         DisplayName = "example",
    ///     });
    /// 
    ///     var exampleServicePrincipal = new Azuread.Index.ServicePrincipal.ServicePrincipal("example", new()
    ///     {
    ///         ApplicationId = example.ApplicationId,
    ///     });
    /// 
    ///     var exampleServicePrincipalPassword = new Azuread.Index.ServicePrincipalPassword.ServicePrincipalPassword("example", new()
    ///     {
    ///         ServicePrincipalId = exampleServicePrincipal.ObjectId,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// *Time-based rotation*
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Azuread = Pulumi.Azuread;
    /// using Time = Pulumi.Time;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Azuread.Index.Application.Application("example", new()
    ///     {
    ///         DisplayName = "example",
    ///     });
    /// 
    ///     var exampleServicePrincipal = new Azuread.Index.ServicePrincipal.ServicePrincipal("example", new()
    ///     {
    ///         ApplicationId = example.ApplicationId,
    ///     });
    /// 
    ///     var exampleRotating = new Time.Index.Rotating.Rotating("example", new()
    ///     {
    ///         RotationDays = 7,
    ///     });
    /// 
    ///     var exampleServicePrincipalPassword = new Azuread.Index.ServicePrincipalPassword.ServicePrincipalPassword("example", new()
    ///     {
    ///         ServicePrincipalId = exampleServicePrincipal.ObjectId,
    ///         RotateWhenChanged = 
    ///         {
    ///             { "rotation", exampleRotating.Id },
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
    [AzureADResourceType("azuread:index/servicePrincipalPassword:ServicePrincipalPassword")]
    public partial class ServicePrincipalPassword : global::Pulumi.CustomResource
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

    public sealed class ServicePrincipalPasswordArgs : global::Pulumi.ResourceArgs
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
        public static new ServicePrincipalPasswordArgs Empty => new ServicePrincipalPasswordArgs();
    }

    public sealed class ServicePrincipalPasswordState : global::Pulumi.ResourceArgs
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

        [Input("value")]
        private Input<string>? _value;

        /// <summary>
        /// The password for this service principal, which is generated by Azure Active Directory.
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

        public ServicePrincipalPasswordState()
        {
        }
        public static new ServicePrincipalPasswordState Empty => new ServicePrincipalPasswordState();
    }
}

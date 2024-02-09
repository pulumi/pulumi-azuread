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
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureAD.ApplicationRegistration("example", new()
    ///     {
    ///         DisplayName = "example",
    ///     });
    /// 
    ///     var exampleAdministrator = new Random.RandomUuid("exampleAdministrator");
    /// 
    ///     var exampleAdminister = new AzureAD.ApplicationAppRole("exampleAdminister", new()
    ///     {
    ///         ApplicationId = example.Id,
    ///         RoleId = exampleAdministrator.Id,
    ///         AllowedMemberTypes = new[]
    ///         {
    ///             "User",
    ///         },
    ///         Description = "My role description",
    ///         DisplayName = "Administer",
    ///         Value = "admin",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &gt; **Tip** For managing more app roles, create additional instances of this resource
    /// 
    /// *Usage with azuread.Application resource*
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
    ///     var exampleAdminister = new AzureAD.ApplicationAppRole("exampleAdminister", new()
    ///     {
    ///         ApplicationId = example.Id,
    ///     });
    /// 
    ///     // ...
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Application App Roles can be imported using the object ID of the application and the ID of the app role, in the following format.
    /// 
    /// ```sh
    /// $ pulumi import azuread:index/applicationAppRole:ApplicationAppRole example /applications/00000000-0000-0000-0000-000000000000/appRoles/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [AzureADResourceType("azuread:index/applicationAppRole:ApplicationAppRole")]
    public partial class ApplicationAppRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A set of values to specify whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications by setting to `Application`, or to both.
        /// </summary>
        [Output("allowedMemberTypes")]
        public Output<ImmutableArray<string>> AllowedMemberTypes { get; private set; } = null!;

        /// <summary>
        /// The resource ID of the application registration. Changing this forces a new resource to be created.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// Description of the app role that appears when the role is being assigned, and if the role functions as an application permissions, during the consent experiences.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Display name for the app role that appears during app role assignment and in consent experiences.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the app role
        /// </summary>
        [Output("roleId")]
        public Output<string> RoleId { get; private set; } = null!;

        /// <summary>
        /// The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
        /// 
        /// &gt; **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
        /// </summary>
        [Output("value")]
        public Output<string?> Value { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationAppRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationAppRole(string name, ApplicationAppRoleArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/applicationAppRole:ApplicationAppRole", name, args ?? new ApplicationAppRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationAppRole(string name, Input<string> id, ApplicationAppRoleState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/applicationAppRole:ApplicationAppRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationAppRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationAppRole Get(string name, Input<string> id, ApplicationAppRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationAppRole(name, id, state, options);
        }
    }

    public sealed class ApplicationAppRoleArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedMemberTypes", required: true)]
        private InputList<string>? _allowedMemberTypes;

        /// <summary>
        /// A set of values to specify whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications by setting to `Application`, or to both.
        /// </summary>
        public InputList<string> AllowedMemberTypes
        {
            get => _allowedMemberTypes ?? (_allowedMemberTypes = new InputList<string>());
            set => _allowedMemberTypes = value;
        }

        /// <summary>
        /// The resource ID of the application registration. Changing this forces a new resource to be created.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// Description of the app role that appears when the role is being assigned, and if the role functions as an application permissions, during the consent experiences.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// Display name for the app role that appears during app role assignment and in consent experiences.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The unique identifier of the app role
        /// </summary>
        [Input("roleId", required: true)]
        public Input<string> RoleId { get; set; } = null!;

        /// <summary>
        /// The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
        /// 
        /// &gt; **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public ApplicationAppRoleArgs()
        {
        }
        public static new ApplicationAppRoleArgs Empty => new ApplicationAppRoleArgs();
    }

    public sealed class ApplicationAppRoleState : global::Pulumi.ResourceArgs
    {
        [Input("allowedMemberTypes")]
        private InputList<string>? _allowedMemberTypes;

        /// <summary>
        /// A set of values to specify whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications by setting to `Application`, or to both.
        /// </summary>
        public InputList<string> AllowedMemberTypes
        {
            get => _allowedMemberTypes ?? (_allowedMemberTypes = new InputList<string>());
            set => _allowedMemberTypes = value;
        }

        /// <summary>
        /// The resource ID of the application registration. Changing this forces a new resource to be created.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// Description of the app role that appears when the role is being assigned, and if the role functions as an application permissions, during the consent experiences.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Display name for the app role that appears during app role assignment and in consent experiences.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The unique identifier of the app role
        /// </summary>
        [Input("roleId")]
        public Input<string>? RoleId { get; set; }

        /// <summary>
        /// The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
        /// 
        /// &gt; **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public ApplicationAppRoleState()
        {
        }
        public static new ApplicationAppRoleState Empty => new ApplicationAppRoleState();
    }
}

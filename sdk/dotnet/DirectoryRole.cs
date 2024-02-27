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
    /// Manages a Directory Role within Azure Active Directory. Directory Roles are also known as Administrator Roles.
    /// 
    /// Directory Roles are built-in to Azure Active Directory and are immutable. However, by default they are not activated in a tenant (except for the Global Administrator role). This resource ensures a directory role is activated from its associated role template, and exports the object ID of the role, so that role assignments can be made for it.
    /// 
    /// Once activated, directory roles cannot be deactivated and so this resource does not perform any actions on destroy.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`
    /// 
    /// When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`
    /// 
    /// ## Example Usage
    /// 
    /// *Activate a directory role by its template ID*
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Azuread = Pulumi.Azuread;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Azuread.Index.DirectoryRole.DirectoryRole("example", new()
    ///     {
    ///         TemplateId = "00000000-0000-0000-0000-000000000000",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// *Activate a directory role by display name*
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Azuread = Pulumi.Azuread;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Azuread.Index.DirectoryRole.DirectoryRole("example", new()
    ///     {
    ///         DisplayName = "Printer administrator",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource does not support importing.
    /// </summary>
    [AzureADResourceType("azuread:index/directoryRole:DirectoryRole")]
    public partial class DirectoryRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the directory role.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the directory role to activate. Changing this forces a new resource to be created.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The object ID of the directory role.
        /// </summary>
        [Output("objectId")]
        public Output<string> ObjectId { get; private set; } = null!;

        /// <summary>
        /// The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.
        /// 
        /// &gt; Either `display_name` or `template_id` must be specified.
        /// </summary>
        [Output("templateId")]
        public Output<string> TemplateId { get; private set; } = null!;


        /// <summary>
        /// Create a DirectoryRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DirectoryRole(string name, DirectoryRoleArgs? args = null, CustomResourceOptions? options = null)
            : base("azuread:index/directoryRole:DirectoryRole", name, args ?? new DirectoryRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DirectoryRole(string name, Input<string> id, DirectoryRoleState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/directoryRole:DirectoryRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DirectoryRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DirectoryRole Get(string name, Input<string> id, DirectoryRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new DirectoryRole(name, id, state, options);
        }
    }

    public sealed class DirectoryRoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The display name of the directory role to activate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.
        /// 
        /// &gt; Either `display_name` or `template_id` must be specified.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        public DirectoryRoleArgs()
        {
        }
        public static new DirectoryRoleArgs Empty => new DirectoryRoleArgs();
    }

    public sealed class DirectoryRoleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the directory role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the directory role to activate. Changing this forces a new resource to be created.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The object ID of the directory role.
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        /// <summary>
        /// The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.
        /// 
        /// &gt; Either `display_name` or `template_id` must be specified.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        public DirectoryRoleState()
        {
        }
        public static new DirectoryRoleState Empty => new DirectoryRoleState();
    }
}

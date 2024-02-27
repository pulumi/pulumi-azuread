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
    /// Manages a single directory role eligibility schedule request within Azure Active Directory.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// The calling principal requires one of the following application roles: `RoleEligibilitySchedule.ReadWrite.Directory` or `RoleManagement.ReadWrite.Directory`.
    /// 
    /// The calling principal requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`.
    /// 
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
    ///     var example = AzureAD.GetUser.Invoke(new()
    ///     {
    ///         UserPrincipalName = "jdoe@example.com",
    ///     });
    /// 
    ///     var exampleDirectoryRole = new AzureAD.DirectoryRole("example", new()
    ///     {
    ///         DisplayName = "Application Administrator",
    ///     });
    /// 
    ///     var exampleDirectoryRoleEligibilityScheduleRequest = new AzureAD.DirectoryRoleEligibilityScheduleRequest("example", new()
    ///     {
    ///         RoleDefinitionId = exampleDirectoryRole.TemplateId,
    ///         PrincipalId = exampleAzureadUser.ObjectId,
    ///         DirectoryScopeId = "/",
    ///         Justification = "Example",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &gt; Note the use of the `template_id` attribute when referencing built-in roles.
    /// 
    /// ## Import
    /// 
    /// Directory role eligibility schedule requests can be imported using the ID of the assignment, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuread:index/directoryRoleEligibilityScheduleRequest:DirectoryRoleEligibilityScheduleRequest example 822ec710-4c9f-4f71-a27a-451759cc7522
    /// ```
    /// </summary>
    [AzureADResourceType("azuread:index/directoryRoleEligibilityScheduleRequest:DirectoryRoleEligibilityScheduleRequest")]
    public partial class DirectoryRoleEligibilityScheduleRequest : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Identifier of the directory object representing the scope of the role eligibility. Changing this forces a new resource to be created.
        /// </summary>
        [Output("directoryScopeId")]
        public Output<string> DirectoryScopeId { get; private set; } = null!;

        /// <summary>
        /// Justification for why the principal is granted the role eligibility. Changing this forces a new resource to be created.
        /// </summary>
        [Output("justification")]
        public Output<string> Justification { get; private set; } = null!;

        /// <summary>
        /// The object ID of the principal to granted the role eligibility. Changing this forces a new resource to be created.
        /// </summary>
        [Output("principalId")]
        public Output<string> PrincipalId { get; private set; } = null!;

        /// <summary>
        /// The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
        /// </summary>
        [Output("roleDefinitionId")]
        public Output<string> RoleDefinitionId { get; private set; } = null!;


        /// <summary>
        /// Create a DirectoryRoleEligibilityScheduleRequest resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DirectoryRoleEligibilityScheduleRequest(string name, DirectoryRoleEligibilityScheduleRequestArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/directoryRoleEligibilityScheduleRequest:DirectoryRoleEligibilityScheduleRequest", name, args ?? new DirectoryRoleEligibilityScheduleRequestArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DirectoryRoleEligibilityScheduleRequest(string name, Input<string> id, DirectoryRoleEligibilityScheduleRequestState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/directoryRoleEligibilityScheduleRequest:DirectoryRoleEligibilityScheduleRequest", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DirectoryRoleEligibilityScheduleRequest resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DirectoryRoleEligibilityScheduleRequest Get(string name, Input<string> id, DirectoryRoleEligibilityScheduleRequestState? state = null, CustomResourceOptions? options = null)
        {
            return new DirectoryRoleEligibilityScheduleRequest(name, id, state, options);
        }
    }

    public sealed class DirectoryRoleEligibilityScheduleRequestArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier of the directory object representing the scope of the role eligibility. Changing this forces a new resource to be created.
        /// </summary>
        [Input("directoryScopeId", required: true)]
        public Input<string> DirectoryScopeId { get; set; } = null!;

        /// <summary>
        /// Justification for why the principal is granted the role eligibility. Changing this forces a new resource to be created.
        /// </summary>
        [Input("justification", required: true)]
        public Input<string> Justification { get; set; } = null!;

        /// <summary>
        /// The object ID of the principal to granted the role eligibility. Changing this forces a new resource to be created.
        /// </summary>
        [Input("principalId", required: true)]
        public Input<string> PrincipalId { get; set; } = null!;

        /// <summary>
        /// The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
        /// </summary>
        [Input("roleDefinitionId", required: true)]
        public Input<string> RoleDefinitionId { get; set; } = null!;

        public DirectoryRoleEligibilityScheduleRequestArgs()
        {
        }
        public static new DirectoryRoleEligibilityScheduleRequestArgs Empty => new DirectoryRoleEligibilityScheduleRequestArgs();
    }

    public sealed class DirectoryRoleEligibilityScheduleRequestState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier of the directory object representing the scope of the role eligibility. Changing this forces a new resource to be created.
        /// </summary>
        [Input("directoryScopeId")]
        public Input<string>? DirectoryScopeId { get; set; }

        /// <summary>
        /// Justification for why the principal is granted the role eligibility. Changing this forces a new resource to be created.
        /// </summary>
        [Input("justification")]
        public Input<string>? Justification { get; set; }

        /// <summary>
        /// The object ID of the principal to granted the role eligibility. Changing this forces a new resource to be created.
        /// </summary>
        [Input("principalId")]
        public Input<string>? PrincipalId { get; set; }

        /// <summary>
        /// The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
        /// </summary>
        [Input("roleDefinitionId")]
        public Input<string>? RoleDefinitionId { get; set; }

        public DirectoryRoleEligibilityScheduleRequestState()
        {
        }
        public static new DirectoryRoleEligibilityScheduleRequestState Empty => new DirectoryRoleEligibilityScheduleRequestState();
    }
}

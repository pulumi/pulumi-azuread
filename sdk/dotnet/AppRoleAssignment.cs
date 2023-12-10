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
    /// Manages an app role assignment for a group, user or service principal. Can be used to grant admin consent for application permissions.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires one of the following application roles: `AppRoleAssignment.ReadWrite.All` and `Application.Read.All`, or `AppRoleAssignment.ReadWrite.All` and `Directory.Read.All`, or `Application.ReadWrite.All`, or `Directory.ReadWrite.All`
    /// 
    /// When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`
    /// 
    /// ## Example Usage
    /// 
    /// *App role assignment for accessing Microsoft Graph*
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureAD = Pulumi.%[1]s;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var wellKnown = AzureAD.GetApplicationPublishedAppIds.Invoke();
    /// 
    ///     var msgraph = new AzureAD.ServicePrincipal("msgraph", new()
    ///     {
    ///         ApplicationId = wellKnown.Apply(getApplicationPublishedAppIdsResult =&gt; getApplicationPublishedAppIdsResult.Result?.MicrosoftGraph),
    ///         UseExisting = true,
    ///     });
    /// 
    ///     var exampleApplication = new AzureAD.Application("exampleApplication", new()
    ///     {
    ///         DisplayName = "example",
    ///         RequiredResourceAccesses = new[]
    ///         {
    ///             new AzureAD.Inputs.ApplicationRequiredResourceAccessArgs
    ///             {
    ///                 ResourceAppId = wellKnown.Apply(getApplicationPublishedAppIdsResult =&gt; getApplicationPublishedAppIdsResult.Result?.MicrosoftGraph),
    ///                 ResourceAccesses = new[]
    ///                 {
    ///                     new AzureAD.Inputs.ApplicationRequiredResourceAccessResourceAccessArgs
    ///                     {
    ///                         Id = msgraph.AppRoleIds.Apply(appRoleIds =&gt; appRoleIds.User_Read_All),
    ///                         Type = "Role",
    ///                     },
    ///                     new AzureAD.Inputs.ApplicationRequiredResourceAccessResourceAccessArgs
    ///                     {
    ///                         Id = msgraph.Oauth2PermissionScopeIds.Apply(oauth2PermissionScopeIds =&gt; oauth2PermissionScopeIds.User_ReadWrite),
    ///                         Type = "Scope",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleServicePrincipal = new AzureAD.ServicePrincipal("exampleServicePrincipal", new()
    ///     {
    ///         ApplicationId = exampleApplication.ApplicationId,
    ///     });
    /// 
    ///     var exampleAppRoleAssignment = new AzureAD.AppRoleAssignment("exampleAppRoleAssignment", new()
    ///     {
    ///         AppRoleId = msgraph.AppRoleIds.Apply(appRoleIds =&gt; appRoleIds.User_Read_All),
    ///         PrincipalObjectId = exampleServicePrincipal.ObjectId,
    ///         ResourceObjectId = msgraph.ObjectId,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// *App role assignment for internal application*
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var internalApplication = new AzureAD.Application("internalApplication", new()
    ///     {
    ///         DisplayName = "internal",
    ///         AppRoles = new[]
    ///         {
    ///             new AzureAD.Inputs.ApplicationAppRoleArgs
    ///             {
    ///                 AllowedMemberTypes = new[]
    ///                 {
    ///                     "Application",
    ///                 },
    ///                 Description = "Apps can query the database",
    ///                 DisplayName = "Query",
    ///                 Enabled = true,
    ///                 Id = "00000000-0000-0000-0000-111111111111",
    ///                 Value = "Query.All",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var internalServicePrincipal = new AzureAD.ServicePrincipal("internalServicePrincipal", new()
    ///     {
    ///         ApplicationId = internalApplication.ApplicationId,
    ///     });
    /// 
    ///     var exampleApplication = new AzureAD.Application("exampleApplication", new()
    ///     {
    ///         DisplayName = "example",
    ///         RequiredResourceAccesses = new[]
    ///         {
    ///             new AzureAD.Inputs.ApplicationRequiredResourceAccessArgs
    ///             {
    ///                 ResourceAppId = internalApplication.ApplicationId,
    ///                 ResourceAccesses = new[]
    ///                 {
    ///                     new AzureAD.Inputs.ApplicationRequiredResourceAccessResourceAccessArgs
    ///                     {
    ///                         Id = internalServicePrincipal.AppRoleIds.Apply(appRoleIds =&gt; appRoleIds.Query_All),
    ///                         Type = "Role",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleServicePrincipal = new AzureAD.ServicePrincipal("exampleServicePrincipal", new()
    ///     {
    ///         ApplicationId = exampleApplication.ApplicationId,
    ///     });
    /// 
    ///     var exampleAppRoleAssignment = new AzureAD.AppRoleAssignment("exampleAppRoleAssignment", new()
    ///     {
    ///         AppRoleId = internalServicePrincipal.AppRoleIds.Apply(appRoleIds =&gt; appRoleIds.Query_All),
    ///         PrincipalObjectId = exampleServicePrincipal.ObjectId,
    ///         ResourceObjectId = internalServicePrincipal.ObjectId,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// *Assign a user and group to an internal application*
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureAD = Pulumi.%[1]s;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleDomains = AzureAD.GetDomains.Invoke(new()
    ///     {
    ///         OnlyInitial = true,
    ///     });
    /// 
    ///     var internalApplication = new AzureAD.Application("internalApplication", new()
    ///     {
    ///         DisplayName = "internal",
    ///         AppRoles = new[]
    ///         {
    ///             new AzureAD.Inputs.ApplicationAppRoleArgs
    ///             {
    ///                 AllowedMemberTypes = new[]
    ///                 {
    ///                     "Application",
    ///                     "User",
    ///                 },
    ///                 Description = "Admins can perform all task actions",
    ///                 DisplayName = "Admin",
    ///                 Enabled = true,
    ///                 Id = "00000000-0000-0000-0000-222222222222",
    ///                 Value = "Admin.All",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var internalServicePrincipal = new AzureAD.ServicePrincipal("internalServicePrincipal", new()
    ///     {
    ///         ApplicationId = internalApplication.ApplicationId,
    ///     });
    /// 
    ///     var exampleGroup = new AzureAD.Group("exampleGroup", new()
    ///     {
    ///         DisplayName = "example",
    ///         SecurityEnabled = true,
    ///     });
    /// 
    ///     var exampleAppRoleAssignment = new AzureAD.AppRoleAssignment("exampleAppRoleAssignment", new()
    ///     {
    ///         AppRoleId = internalServicePrincipal.AppRoleIds.Apply(appRoleIds =&gt; appRoleIds.Admin_All),
    ///         PrincipalObjectId = exampleGroup.ObjectId,
    ///         ResourceObjectId = internalServicePrincipal.ObjectId,
    ///     });
    /// 
    ///     var exampleUser = new AzureAD.User("exampleUser", new()
    ///     {
    ///         DisplayName = "D. Duck",
    ///         Password = "SecretP@sswd99!",
    ///         UserPrincipalName = $"d.duck@{exampleDomains.Apply(getDomainsResult =&gt; getDomainsResult.Domains[0]?.DomainName)}",
    ///     });
    /// 
    ///     var exampleIndex_appRoleAssignmentAppRoleAssignment = new AzureAD.AppRoleAssignment("exampleIndex/appRoleAssignmentAppRoleAssignment", new()
    ///     {
    ///         AppRoleId = internalServicePrincipal.AppRoleIds.Apply(appRoleIds =&gt; appRoleIds.Admin_All),
    ///         PrincipalObjectId = exampleUser.ObjectId,
    ///         ResourceObjectId = internalServicePrincipal.ObjectId,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// *Assign a group to the default app role for an internal application*
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var internalApplication = new AzureAD.Application("internalApplication", new()
    ///     {
    ///         DisplayName = "internal",
    ///     });
    /// 
    ///     var internalServicePrincipal = new AzureAD.ServicePrincipal("internalServicePrincipal", new()
    ///     {
    ///         ApplicationId = internalApplication.ApplicationId,
    ///     });
    /// 
    ///     var exampleGroup = new AzureAD.Group("exampleGroup", new()
    ///     {
    ///         DisplayName = "example",
    ///         SecurityEnabled = true,
    ///     });
    /// 
    ///     var exampleAppRoleAssignment = new AzureAD.AppRoleAssignment("exampleAppRoleAssignment", new()
    ///     {
    ///         AppRoleId = "00000000-0000-0000-0000-000000000000",
    ///         PrincipalObjectId = exampleGroup.ObjectId,
    ///         ResourceObjectId = internalServicePrincipal.ObjectId,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// App role assignments can be imported using the object ID of the service principal representing the resource and the ID of the app role assignment (note_not_ the ID of the app role), e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azuread:index/appRoleAssignment:AppRoleAssignment example 00000000-0000-0000-0000-000000000000/appRoleAssignment/aaBBcDDeFG6h5JKLMN2PQrrssTTUUvWWxxxxxyyyzzz
    /// ```
    /// 
    ///  -&gt; This ID format is unique to Terraform and is composed of the Resource Service Principal Object ID and the ID of the App Role Assignment in the format `{ResourcePrincipalID}/appRoleAssignment/{AppRoleAssignmentID}`.
    /// </summary>
    [AzureADResourceType("azuread:index/appRoleAssignment:AppRoleAssignment")]
    public partial class AppRoleAssignment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the app role to be assigned, or the default role ID `00000000-0000-0000-0000-000000000000`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("appRoleId")]
        public Output<string> AppRoleId { get; private set; } = null!;

        /// <summary>
        /// The display name of the principal to which the app role is assigned.
        /// </summary>
        [Output("principalDisplayName")]
        public Output<string> PrincipalDisplayName { get; private set; } = null!;

        /// <summary>
        /// The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        /// </summary>
        [Output("principalObjectId")]
        public Output<string> PrincipalObjectId { get; private set; } = null!;

        /// <summary>
        /// The object type of the principal to which the app role is assigned.
        /// </summary>
        [Output("principalType")]
        public Output<string> PrincipalType { get; private set; } = null!;

        /// <summary>
        /// The display name of the application representing the resource.
        /// </summary>
        [Output("resourceDisplayName")]
        public Output<string> ResourceDisplayName { get; private set; } = null!;

        /// <summary>
        /// The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceObjectId")]
        public Output<string> ResourceObjectId { get; private set; } = null!;


        /// <summary>
        /// Create a AppRoleAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppRoleAssignment(string name, AppRoleAssignmentArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/appRoleAssignment:AppRoleAssignment", name, args ?? new AppRoleAssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppRoleAssignment(string name, Input<string> id, AppRoleAssignmentState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/appRoleAssignment:AppRoleAssignment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppRoleAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppRoleAssignment Get(string name, Input<string> id, AppRoleAssignmentState? state = null, CustomResourceOptions? options = null)
        {
            return new AppRoleAssignment(name, id, state, options);
        }
    }

    public sealed class AppRoleAssignmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the app role to be assigned, or the default role ID `00000000-0000-0000-0000-000000000000`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("appRoleId", required: true)]
        public Input<string> AppRoleId { get; set; } = null!;

        /// <summary>
        /// The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        /// </summary>
        [Input("principalObjectId", required: true)]
        public Input<string> PrincipalObjectId { get; set; } = null!;

        /// <summary>
        /// The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceObjectId", required: true)]
        public Input<string> ResourceObjectId { get; set; } = null!;

        public AppRoleAssignmentArgs()
        {
        }
        public static new AppRoleAssignmentArgs Empty => new AppRoleAssignmentArgs();
    }

    public sealed class AppRoleAssignmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the app role to be assigned, or the default role ID `00000000-0000-0000-0000-000000000000`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("appRoleId")]
        public Input<string>? AppRoleId { get; set; }

        /// <summary>
        /// The display name of the principal to which the app role is assigned.
        /// </summary>
        [Input("principalDisplayName")]
        public Input<string>? PrincipalDisplayName { get; set; }

        /// <summary>
        /// The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        /// </summary>
        [Input("principalObjectId")]
        public Input<string>? PrincipalObjectId { get; set; }

        /// <summary>
        /// The object type of the principal to which the app role is assigned.
        /// </summary>
        [Input("principalType")]
        public Input<string>? PrincipalType { get; set; }

        /// <summary>
        /// The display name of the application representing the resource.
        /// </summary>
        [Input("resourceDisplayName")]
        public Input<string>? ResourceDisplayName { get; set; }

        /// <summary>
        /// The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceObjectId")]
        public Input<string>? ResourceObjectId { get; set; }

        public AppRoleAssignmentState()
        {
        }
        public static new AppRoleAssignmentState Empty => new AppRoleAssignmentState();
    }
}

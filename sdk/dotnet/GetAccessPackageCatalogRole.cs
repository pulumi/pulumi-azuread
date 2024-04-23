// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    public static class GetAccessPackageCatalogRole
    {
        /// <summary>
        /// Gets information about an access package catalog role.
        /// 
        /// ## API Permissions
        /// 
        /// The following API permissions are required in order to use this data source.
        /// 
        /// When authenticated with a service principal, this data source requires one of the following application roles: `EntitlementManagement.Read.All` or `Directory.Read.All`
        /// 
        /// When authenticated with a user principal, this data source does not require any additional roles.
        /// 
        /// ## Example Usage
        /// 
        /// ### By Group Display Name)
        /// 
        /// *Look up by display name*
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureAD.GetAccessPackageCatalogRole.Invoke(new()
        ///     {
        ///         DisplayName = "Catalog owner",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// *Look up by object ID*
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureAD.GetAccessPackageCatalogRole.Invoke(new()
        ///     {
        ///         ObjectId = "00000000-0000-0000-0000-000000000000",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAccessPackageCatalogRoleResult> InvokeAsync(GetAccessPackageCatalogRoleArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccessPackageCatalogRoleResult>("azuread:index/getAccessPackageCatalogRole:getAccessPackageCatalogRole", args ?? new GetAccessPackageCatalogRoleArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an access package catalog role.
        /// 
        /// ## API Permissions
        /// 
        /// The following API permissions are required in order to use this data source.
        /// 
        /// When authenticated with a service principal, this data source requires one of the following application roles: `EntitlementManagement.Read.All` or `Directory.Read.All`
        /// 
        /// When authenticated with a user principal, this data source does not require any additional roles.
        /// 
        /// ## Example Usage
        /// 
        /// ### By Group Display Name)
        /// 
        /// *Look up by display name*
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureAD.GetAccessPackageCatalogRole.Invoke(new()
        ///     {
        ///         DisplayName = "Catalog owner",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// *Look up by object ID*
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureAD.GetAccessPackageCatalogRole.Invoke(new()
        ///     {
        ///         ObjectId = "00000000-0000-0000-0000-000000000000",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAccessPackageCatalogRoleResult> Invoke(GetAccessPackageCatalogRoleInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccessPackageCatalogRoleResult>("azuread:index/getAccessPackageCatalogRole:getAccessPackageCatalogRole", args ?? new GetAccessPackageCatalogRoleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccessPackageCatalogRoleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the display name of the role.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        /// <summary>
        /// Specifies the object ID of the role.
        /// 
        /// &gt; One of `display_name` or `object_id` must be specified.
        /// </summary>
        [Input("objectId")]
        public string? ObjectId { get; set; }

        public GetAccessPackageCatalogRoleArgs()
        {
        }
        public static new GetAccessPackageCatalogRoleArgs Empty => new GetAccessPackageCatalogRoleArgs();
    }

    public sealed class GetAccessPackageCatalogRoleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the display name of the role.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Specifies the object ID of the role.
        /// 
        /// &gt; One of `display_name` or `object_id` must be specified.
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        public GetAccessPackageCatalogRoleInvokeArgs()
        {
        }
        public static new GetAccessPackageCatalogRoleInvokeArgs Empty => new GetAccessPackageCatalogRoleInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccessPackageCatalogRoleResult
    {
        /// <summary>
        /// The description of the role.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The display name of the role.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The object ID of the role.
        /// </summary>
        public readonly string ObjectId;
        /// <summary>
        /// The object ID of the role.
        /// </summary>
        public readonly string TemplateId;

        [OutputConstructor]
        private GetAccessPackageCatalogRoleResult(
            string description,

            string displayName,

            string id,

            string objectId,

            string templateId)
        {
            Description = description;
            DisplayName = displayName;
            Id = id;
            ObjectId = objectId;
            TemplateId = templateId;
        }
    }
}

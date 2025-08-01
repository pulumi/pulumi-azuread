// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    public static class GetAccessPackageCatalog
    {
        /// <summary>
        /// i
        /// Use this resource to retrieve information for an existing access package catalog within Identity Governance in Azure Active Directory.
        /// 
        /// ## API Permissions
        /// 
        /// The following API permissions are required in order to use this data source.
        /// 
        /// When authenticated with a service principal, this data source requires one of the following application roles: `EntitlementManagement.Read.All`, or `EntitlementManagement.ReadWrite.All`.
        /// 
        /// When authenticated with a user principal, this data source requires one of the following directory roles: `Catalog owner`, `Catalog reader`, `Global Reader`, or `Global Administrator`.
        /// 
        /// ## Example Usage
        /// 
        /// *Look up by ID*
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureAD.GetAccessPackageCatalog.Invoke(new()
        ///     {
        ///         ObjectId = "00000000-0000-0000-0000-000000000000",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// *Look up by DisplayName*
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureAD.GetAccessPackageCatalog.Invoke(new()
        ///     {
        ///         DisplayName = "My access package Catalog",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAccessPackageCatalogResult> InvokeAsync(GetAccessPackageCatalogArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccessPackageCatalogResult>("azuread:index/getAccessPackageCatalog:getAccessPackageCatalog", args ?? new GetAccessPackageCatalogArgs(), options.WithDefaults());

        /// <summary>
        /// i
        /// Use this resource to retrieve information for an existing access package catalog within Identity Governance in Azure Active Directory.
        /// 
        /// ## API Permissions
        /// 
        /// The following API permissions are required in order to use this data source.
        /// 
        /// When authenticated with a service principal, this data source requires one of the following application roles: `EntitlementManagement.Read.All`, or `EntitlementManagement.ReadWrite.All`.
        /// 
        /// When authenticated with a user principal, this data source requires one of the following directory roles: `Catalog owner`, `Catalog reader`, `Global Reader`, or `Global Administrator`.
        /// 
        /// ## Example Usage
        /// 
        /// *Look up by ID*
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureAD.GetAccessPackageCatalog.Invoke(new()
        ///     {
        ///         ObjectId = "00000000-0000-0000-0000-000000000000",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// *Look up by DisplayName*
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureAD.GetAccessPackageCatalog.Invoke(new()
        ///     {
        ///         DisplayName = "My access package Catalog",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAccessPackageCatalogResult> Invoke(GetAccessPackageCatalogInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccessPackageCatalogResult>("azuread:index/getAccessPackageCatalog:getAccessPackageCatalog", args ?? new GetAccessPackageCatalogInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// i
        /// Use this resource to retrieve information for an existing access package catalog within Identity Governance in Azure Active Directory.
        /// 
        /// ## API Permissions
        /// 
        /// The following API permissions are required in order to use this data source.
        /// 
        /// When authenticated with a service principal, this data source requires one of the following application roles: `EntitlementManagement.Read.All`, or `EntitlementManagement.ReadWrite.All`.
        /// 
        /// When authenticated with a user principal, this data source requires one of the following directory roles: `Catalog owner`, `Catalog reader`, `Global Reader`, or `Global Administrator`.
        /// 
        /// ## Example Usage
        /// 
        /// *Look up by ID*
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureAD.GetAccessPackageCatalog.Invoke(new()
        ///     {
        ///         ObjectId = "00000000-0000-0000-0000-000000000000",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// *Look up by DisplayName*
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureAD.GetAccessPackageCatalog.Invoke(new()
        ///     {
        ///         DisplayName = "My access package Catalog",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAccessPackageCatalogResult> Invoke(GetAccessPackageCatalogInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccessPackageCatalogResult>("azuread:index/getAccessPackageCatalog:getAccessPackageCatalog", args ?? new GetAccessPackageCatalogInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccessPackageCatalogArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The display name of the access package catalog.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        /// <summary>
        /// The ID of this access package catalog.
        /// 
        /// &gt; One of `display_name` or `object_id` must be specified.
        /// </summary>
        [Input("objectId")]
        public string? ObjectId { get; set; }

        public GetAccessPackageCatalogArgs()
        {
        }
        public static new GetAccessPackageCatalogArgs Empty => new GetAccessPackageCatalogArgs();
    }

    public sealed class GetAccessPackageCatalogInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The display name of the access package catalog.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The ID of this access package catalog.
        /// 
        /// &gt; One of `display_name` or `object_id` must be specified.
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        public GetAccessPackageCatalogInvokeArgs()
        {
        }
        public static new GetAccessPackageCatalogInvokeArgs Empty => new GetAccessPackageCatalogInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccessPackageCatalogResult
    {
        /// <summary>
        /// The description of the access package catalog.
        /// </summary>
        public readonly string Description;
        public readonly string DisplayName;
        /// <summary>
        /// Whether the access packages in this catalog can be requested by users outside the tenant.
        /// </summary>
        public readonly bool ExternallyVisible;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ObjectId;
        /// <summary>
        /// Whether the access packages in this catalog are available for management.
        /// </summary>
        public readonly bool Published;

        [OutputConstructor]
        private GetAccessPackageCatalogResult(
            string description,

            string displayName,

            bool externallyVisible,

            string id,

            string objectId,

            bool published)
        {
            Description = description;
            DisplayName = displayName;
            ExternallyVisible = externallyVisible;
            Id = id;
            ObjectId = objectId;
            Published = published;
        }
    }
}

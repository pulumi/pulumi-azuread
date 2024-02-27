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
    /// Manages the resources added to access package catalogs within Identity Governance in Azure Active Directory.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.
    /// 
    /// When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner` or `Global Administrator`
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Azuread = Pulumi.Azuread;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Azuread.Index.Group.Group("example", new()
    ///     {
    ///         DisplayName = "example-group",
    ///         SecurityEnabled = true,
    ///     });
    /// 
    ///     var exampleAccessPackageCatalog = new Azuread.Index.AccessPackageCatalog.AccessPackageCatalog("example", new()
    ///     {
    ///         DisplayName = "example-catalog",
    ///         Description = "Example catalog",
    ///     });
    /// 
    ///     var exampleAccessPackageResourceCatalogAssociation = new Azuread.Index.AccessPackageResourceCatalogAssociation.AccessPackageResourceCatalogAssociation("example", new()
    ///     {
    ///         CatalogId = exampleCatalog.Id,
    ///         ResourceOriginId = exampleGroup.ObjectId,
    ///         ResourceOriginSystem = "AadGroup",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// The resource and catalog association can be imported using the catalog ID and the resource origin ID, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuread:index/accessPackageResourceCatalogAssociation:AccessPackageResourceCatalogAssociation example 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111
    /// ```
    /// 
    ///  -&gt; This ID format is unique to Terraform and is composed of the Catalog ID and the Resource Origin ID in the format `{CatalogID}/{ResourceOriginID}`.
    /// </summary>
    [AzureADResourceType("azuread:index/accessPackageResourceCatalogAssociation:AccessPackageResourceCatalogAssociation")]
    public partial class AccessPackageResourceCatalogAssociation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The unique ID of the access package catalog. Changing this forces a new resource to be created.
        /// </summary>
        [Output("catalogId")]
        public Output<string> CatalogId { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceOriginId")]
        public Output<string> ResourceOriginId { get; private set; } = null!;

        /// <summary>
        /// The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceOriginSystem")]
        public Output<string> ResourceOriginSystem { get; private set; } = null!;


        /// <summary>
        /// Create a AccessPackageResourceCatalogAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessPackageResourceCatalogAssociation(string name, AccessPackageResourceCatalogAssociationArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/accessPackageResourceCatalogAssociation:AccessPackageResourceCatalogAssociation", name, args ?? new AccessPackageResourceCatalogAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessPackageResourceCatalogAssociation(string name, Input<string> id, AccessPackageResourceCatalogAssociationState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/accessPackageResourceCatalogAssociation:AccessPackageResourceCatalogAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccessPackageResourceCatalogAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessPackageResourceCatalogAssociation Get(string name, Input<string> id, AccessPackageResourceCatalogAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new AccessPackageResourceCatalogAssociation(name, id, state, options);
        }
    }

    public sealed class AccessPackageResourceCatalogAssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique ID of the access package catalog. Changing this forces a new resource to be created.
        /// </summary>
        [Input("catalogId", required: true)]
        public Input<string> CatalogId { get; set; } = null!;

        /// <summary>
        /// The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceOriginId", required: true)]
        public Input<string> ResourceOriginId { get; set; } = null!;

        /// <summary>
        /// The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceOriginSystem", required: true)]
        public Input<string> ResourceOriginSystem { get; set; } = null!;

        public AccessPackageResourceCatalogAssociationArgs()
        {
        }
        public static new AccessPackageResourceCatalogAssociationArgs Empty => new AccessPackageResourceCatalogAssociationArgs();
    }

    public sealed class AccessPackageResourceCatalogAssociationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique ID of the access package catalog. Changing this forces a new resource to be created.
        /// </summary>
        [Input("catalogId")]
        public Input<string>? CatalogId { get; set; }

        /// <summary>
        /// The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceOriginId")]
        public Input<string>? ResourceOriginId { get; set; }

        /// <summary>
        /// The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceOriginSystem")]
        public Input<string>? ResourceOriginSystem { get; set; }

        public AccessPackageResourceCatalogAssociationState()
        {
        }
        public static new AccessPackageResourceCatalogAssociationState Empty => new AccessPackageResourceCatalogAssociationState();
    }
}

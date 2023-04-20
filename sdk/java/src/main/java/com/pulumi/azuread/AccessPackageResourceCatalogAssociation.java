// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.AccessPackageResourceCatalogAssociationArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.AccessPackageResourceCatalogAssociationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages the resources added to access package catalogs within Identity Governance in Azure Active Directory.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.
 * 
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner` or `Global Administrator`
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.Group;
 * import com.pulumi.azuread.GroupArgs;
 * import com.pulumi.azuread.AccessPackageCatalog;
 * import com.pulumi.azuread.AccessPackageCatalogArgs;
 * import com.pulumi.azuread.AccessPackageResourceCatalogAssociation;
 * import com.pulumi.azuread.AccessPackageResourceCatalogAssociationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var exampleGroup = new Group(&#34;exampleGroup&#34;, GroupArgs.builder()        
 *             .displayName(&#34;example-group&#34;)
 *             .securityEnabled(true)
 *             .build());
 * 
 *         var exampleAccessPackageCatalog = new AccessPackageCatalog(&#34;exampleAccessPackageCatalog&#34;, AccessPackageCatalogArgs.builder()        
 *             .displayName(&#34;example-catalog&#34;)
 *             .description(&#34;Example catalog&#34;)
 *             .build());
 * 
 *         var exampleAccessPackageResourceCatalogAssociation = new AccessPackageResourceCatalogAssociation(&#34;exampleAccessPackageResourceCatalogAssociation&#34;, AccessPackageResourceCatalogAssociationArgs.builder()        
 *             .catalogId(azuread_access_package_catalog.example_catalog().id())
 *             .resourceOriginId(azuread_group.example_group().object_id())
 *             .resourceOriginSystem(&#34;AadGroup&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * The resource and catalog association can be imported using the catalog ID and the resource origin ID, e.g.
 * 
 * ```sh
 *  $ pulumi import azuread:index/accessPackageResourceCatalogAssociation:AccessPackageResourceCatalogAssociation example 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111
 * ```
 * 
 *  -&gt; This ID format is unique to Terraform and is composed of the Catalog ID and the Resource Origin ID in the format `{CatalogID}/{ResourceOriginID}`.
 * 
 */
@ResourceType(type="azuread:index/accessPackageResourceCatalogAssociation:AccessPackageResourceCatalogAssociation")
public class AccessPackageResourceCatalogAssociation extends com.pulumi.resources.CustomResource {
    /**
     * The unique ID of the access package catalog. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="catalogId", type=String.class, parameters={})
    private Output<String> catalogId;

    /**
     * @return The unique ID of the access package catalog. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> catalogId() {
        return this.catalogId;
    }
    /**
     * The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="resourceOriginId", type=String.class, parameters={})
    private Output<String> resourceOriginId;

    /**
     * @return The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> resourceOriginId() {
        return this.resourceOriginId;
    }
    /**
     * The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="resourceOriginSystem", type=String.class, parameters={})
    private Output<String> resourceOriginSystem;

    /**
     * @return The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> resourceOriginSystem() {
        return this.resourceOriginSystem;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AccessPackageResourceCatalogAssociation(String name) {
        this(name, AccessPackageResourceCatalogAssociationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AccessPackageResourceCatalogAssociation(String name, AccessPackageResourceCatalogAssociationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AccessPackageResourceCatalogAssociation(String name, AccessPackageResourceCatalogAssociationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/accessPackageResourceCatalogAssociation:AccessPackageResourceCatalogAssociation", name, args == null ? AccessPackageResourceCatalogAssociationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AccessPackageResourceCatalogAssociation(String name, Output<String> id, @Nullable AccessPackageResourceCatalogAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/accessPackageResourceCatalogAssociation:AccessPackageResourceCatalogAssociation", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static AccessPackageResourceCatalogAssociation get(String name, Output<String> id, @Nullable AccessPackageResourceCatalogAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AccessPackageResourceCatalogAssociation(name, id, state, options);
    }
}

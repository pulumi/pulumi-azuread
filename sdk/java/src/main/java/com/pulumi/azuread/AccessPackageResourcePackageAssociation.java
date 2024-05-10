// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.AccessPackageResourcePackageAssociationArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.AccessPackageResourcePackageAssociationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages the resources added to access packages within Identity Governance in Azure Active Directory.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.
 * 
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner`, `Access package manager` or `Global Administrator`.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 * import com.pulumi.azuread.AccessPackage;
 * import com.pulumi.azuread.AccessPackageArgs;
 * import com.pulumi.azuread.AccessPackageResourcePackageAssociation;
 * import com.pulumi.azuread.AccessPackageResourcePackageAssociationArgs;
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
 *         var example = new Group("example", GroupArgs.builder()        
 *             .displayName("example-group")
 *             .securityEnabled(true)
 *             .build());
 * 
 *         var exampleAccessPackageCatalog = new AccessPackageCatalog("exampleAccessPackageCatalog", AccessPackageCatalogArgs.builder()        
 *             .displayName("example-catalog")
 *             .description("Example catalog")
 *             .build());
 * 
 *         var exampleAccessPackageResourceCatalogAssociation = new AccessPackageResourceCatalogAssociation("exampleAccessPackageResourceCatalogAssociation", AccessPackageResourceCatalogAssociationArgs.builder()        
 *             .catalogId(exampleCatalog.id())
 *             .resourceOriginId(exampleGroup.objectId())
 *             .resourceOriginSystem("AadGroup")
 *             .build());
 * 
 *         var exampleAccessPackage = new AccessPackage("exampleAccessPackage", AccessPackageArgs.builder()        
 *             .displayName("example-package")
 *             .description("Example Package")
 *             .catalogId(exampleCatalog.id())
 *             .build());
 * 
 *         var exampleAccessPackageResourcePackageAssociation = new AccessPackageResourcePackageAssociation("exampleAccessPackageResourcePackageAssociation", AccessPackageResourcePackageAssociationArgs.builder()        
 *             .accessPackageId(exampleAccessPackage.id())
 *             .catalogResourceAssociationId(exampleAccessPackageResourceCatalogAssociation.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * The resource and catalog association can be imported using the access package ID, the access package ResourceRoleScope, the resource origin ID, and the access type, e.g.
 * 
 * ```sh
 * $ pulumi import azuread:index/accessPackageResourcePackageAssociation:AccessPackageResourcePackageAssociation example 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111_22222222-2222-2222-2222-22222222/33333333-3333-3333-3333-33333333/Member
 * ```
 * 
 * -&gt; This ID format is unique to Terraform and is composed of the Access Package ID, the access package ResourceRoleScope (in the format Role_Scope), the Resource Origin ID, and the Access Type, in the format `{AccessPackageID}/{ResourceRoleScope}/{ResourceOriginID}/{AccessType}`.
 * 
 */
@ResourceType(type="azuread:index/accessPackageResourcePackageAssociation:AccessPackageResourcePackageAssociation")
public class AccessPackageResourcePackageAssociation extends com.pulumi.resources.CustomResource {
    /**
     * The ID of access package this resource association is configured to. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="accessPackageId", refs={String.class}, tree="[0]")
    private Output<String> accessPackageId;

    /**
     * @return The ID of access package this resource association is configured to. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> accessPackageId() {
        return this.accessPackageId;
    }
    /**
     * The role of access type to the specified resource. Valid values are `Member`, or `Owner` The default is `Member`. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="accessType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accessType;

    /**
     * @return The role of access type to the specified resource. Valid values are `Member`, or `Owner` The default is `Member`. Changing this forces a new resource to be created.
     * 
     */
    public Output<Optional<String>> accessType() {
        return Codegen.optional(this.accessType);
    }
    /**
     * The ID of the catalog association from the `azuread.AccessPackageResourceCatalogAssociation` resource. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="catalogResourceAssociationId", refs={String.class}, tree="[0]")
    private Output<String> catalogResourceAssociationId;

    /**
     * @return The ID of the catalog association from the `azuread.AccessPackageResourceCatalogAssociation` resource. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> catalogResourceAssociationId() {
        return this.catalogResourceAssociationId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AccessPackageResourcePackageAssociation(String name) {
        this(name, AccessPackageResourcePackageAssociationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AccessPackageResourcePackageAssociation(String name, AccessPackageResourcePackageAssociationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AccessPackageResourcePackageAssociation(String name, AccessPackageResourcePackageAssociationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/accessPackageResourcePackageAssociation:AccessPackageResourcePackageAssociation", name, args == null ? AccessPackageResourcePackageAssociationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AccessPackageResourcePackageAssociation(String name, Output<String> id, @Nullable AccessPackageResourcePackageAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/accessPackageResourcePackageAssociation:AccessPackageResourcePackageAssociation", name, state, makeResourceOptions(options, id));
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
    public static AccessPackageResourcePackageAssociation get(String name, Output<String> id, @Nullable AccessPackageResourcePackageAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AccessPackageResourcePackageAssociation(name, id, state, options);
    }
}

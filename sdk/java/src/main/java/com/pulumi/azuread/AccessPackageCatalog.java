// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.AccessPackageCatalogArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.AccessPackageCatalogState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an access package catalog within Identity Governance in Azure Active Directory.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.
 * 
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner`, `Catalog creator` or `Global Administrator`
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
 * import com.pulumi.azuread.AccessPackageCatalog;
 * import com.pulumi.azuread.AccessPackageCatalogArgs;
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
 *         var example = new AccessPackageCatalog("example", AccessPackageCatalogArgs.builder()
 *             .displayName("example-access-package-catalog")
 *             .description("Example access package catalog")
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
 * An Access Package Catalog can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import azuread:index/accessPackageCatalog:AccessPackageCatalog example 00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuread:index/accessPackageCatalog:AccessPackageCatalog")
public class AccessPackageCatalog extends com.pulumi.resources.CustomResource {
    /**
     * The description of the access package catalog.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return The description of the access package catalog.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The display name of the access package catalog.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The display name of the access package catalog.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * Whether the access packages in this catalog can be requested by users outside the tenant.
     * 
     */
    @Export(name="externallyVisible", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> externallyVisible;

    /**
     * @return Whether the access packages in this catalog can be requested by users outside the tenant.
     * 
     */
    public Output<Optional<Boolean>> externallyVisible() {
        return Codegen.optional(this.externallyVisible);
    }
    /**
     * Whether the access packages in this catalog are available for management.
     * 
     */
    @Export(name="published", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> published;

    /**
     * @return Whether the access packages in this catalog are available for management.
     * 
     */
    public Output<Optional<Boolean>> published() {
        return Codegen.optional(this.published);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AccessPackageCatalog(String name) {
        this(name, AccessPackageCatalogArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AccessPackageCatalog(String name, AccessPackageCatalogArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AccessPackageCatalog(String name, AccessPackageCatalogArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/accessPackageCatalog:AccessPackageCatalog", name, args == null ? AccessPackageCatalogArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AccessPackageCatalog(String name, Output<String> id, @Nullable AccessPackageCatalogState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/accessPackageCatalog:AccessPackageCatalog", name, state, makeResourceOptions(options, id));
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
    public static AccessPackageCatalog get(String name, Output<String> id, @Nullable AccessPackageCatalogState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AccessPackageCatalog(name, id, state, options);
    }
}

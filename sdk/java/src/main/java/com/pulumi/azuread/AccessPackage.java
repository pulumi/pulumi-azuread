// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.AccessPackageArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.AccessPackageState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an Access Package within Identity Governance in Azure Active Directory.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.
 * 
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner`, `Access package manager` or `Global Administrator`
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
 * import com.pulumi.azuread.AccessPackage;
 * import com.pulumi.azuread.AccessPackageArgs;
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
 *             .displayName("example-catalog")
 *             .description("Example catalog")
 *             .build());
 * 
 *         var exampleAccessPackage = new AccessPackage("exampleAccessPackage", AccessPackageArgs.builder()
 *             .catalogId(example.id())
 *             .displayName("access-package")
 *             .description("Access Package")
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
 * Access Packages can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import azuread:index/accessPackage:AccessPackage example_package 00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuread:index/accessPackage:AccessPackage")
public class AccessPackage extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the Catalog this access package will be created in.
     * 
     */
    @Export(name="catalogId", refs={String.class}, tree="[0]")
    private Output<String> catalogId;

    /**
     * @return The ID of the Catalog this access package will be created in.
     * 
     */
    public Output<String> catalogId() {
        return this.catalogId;
    }
    /**
     * The description of the access package.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return The description of the access package.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The display name of the access package.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The display name of the access package.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * Whether the access package is hidden from the requestor.
     * 
     */
    @Export(name="hidden", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> hidden;

    /**
     * @return Whether the access package is hidden from the requestor.
     * 
     */
    public Output<Optional<Boolean>> hidden() {
        return Codegen.optional(this.hidden);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AccessPackage(String name) {
        this(name, AccessPackageArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AccessPackage(String name, AccessPackageArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AccessPackage(String name, AccessPackageArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/accessPackage:AccessPackage", name, args == null ? AccessPackageArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AccessPackage(String name, Output<String> id, @Nullable AccessPackageState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/accessPackage:AccessPackage", name, state, makeResourceOptions(options, id));
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
    public static AccessPackage get(String name, Output<String> id, @Nullable AccessPackageState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AccessPackage(name, id, state, options);
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.AccessPackageCatalogRoleAssignmentArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.AccessPackageCatalogRoleAssignmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages a single catalog role assignment within Azure Active Directory.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires one of the following application roles: `EntitlementManagement.ReadWrite.All` or `Directory.ReadWrite.All`
 * 
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Identity Governance administrator` or `Global Administrator`
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
 * import com.pulumi.azuread.AzureadFunctions;
 * import com.pulumi.azuread.inputs.GetUserArgs;
 * import com.pulumi.azuread.inputs.GetAccessPackageCatalogRoleArgs;
 * import com.pulumi.azuread.AccessPackageCatalog;
 * import com.pulumi.azuread.AccessPackageCatalogArgs;
 * import com.pulumi.azuread.AccessPackageCatalogRoleAssignment;
 * import com.pulumi.azuread.AccessPackageCatalogRoleAssignmentArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         final var example = AzureadFunctions.getUser(GetUserArgs.builder()
 *             .userPrincipalName("jdoe}{@literal @}{@code example.com")
 *             .build());
 * 
 *         final var exampleGetAccessPackageCatalogRole = AzureadFunctions.getAccessPackageCatalogRole(GetAccessPackageCatalogRoleArgs.builder()
 *             .displayName("Catalog owner")
 *             .build());
 * 
 *         var exampleAccessPackageCatalog = new AccessPackageCatalog("exampleAccessPackageCatalog", AccessPackageCatalogArgs.builder()
 *             .displayName("example-access-package-catalog")
 *             .description("Example access package catalog")
 *             .build());
 * 
 *         var exampleAccessPackageCatalogRoleAssignment = new AccessPackageCatalogRoleAssignment("exampleAccessPackageCatalogRoleAssignment", AccessPackageCatalogRoleAssignmentArgs.builder()
 *             .roleId(exampleGetAccessPackageCatalogRole.objectId())
 *             .principalObjectId(example.objectId())
 *             .catalogId(exampleAccessPackageCatalog.id())
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Catalog role assignments can be imported using the ID of the assignment, e.g.
 * 
 * ```sh
 * $ pulumi import azuread:index/accessPackageCatalogRoleAssignment:AccessPackageCatalogRoleAssignment example 00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuread:index/accessPackageCatalogRoleAssignment:AccessPackageCatalogRoleAssignment")
public class AccessPackageCatalogRoleAssignment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the Catalog this role assignment will be scoped to. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="catalogId", refs={String.class}, tree="[0]")
    private Output<String> catalogId;

    /**
     * @return The ID of the Catalog this role assignment will be scoped to. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> catalogId() {
        return this.catalogId;
    }
    /**
     * The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="principalObjectId", refs={String.class}, tree="[0]")
    private Output<String> principalObjectId;

    /**
     * @return The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> principalObjectId() {
        return this.principalObjectId;
    }
    /**
     * The object ID of the catalog role you want to assign. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="roleId", refs={String.class}, tree="[0]")
    private Output<String> roleId;

    /**
     * @return The object ID of the catalog role you want to assign. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> roleId() {
        return this.roleId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AccessPackageCatalogRoleAssignment(java.lang.String name) {
        this(name, AccessPackageCatalogRoleAssignmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AccessPackageCatalogRoleAssignment(java.lang.String name, AccessPackageCatalogRoleAssignmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AccessPackageCatalogRoleAssignment(java.lang.String name, AccessPackageCatalogRoleAssignmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/accessPackageCatalogRoleAssignment:AccessPackageCatalogRoleAssignment", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AccessPackageCatalogRoleAssignment(java.lang.String name, Output<java.lang.String> id, @Nullable AccessPackageCatalogRoleAssignmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/accessPackageCatalogRoleAssignment:AccessPackageCatalogRoleAssignment", name, state, makeResourceOptions(options, id), false);
    }

    private static AccessPackageCatalogRoleAssignmentArgs makeArgs(AccessPackageCatalogRoleAssignmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AccessPackageCatalogRoleAssignmentArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static AccessPackageCatalogRoleAssignment get(java.lang.String name, Output<java.lang.String> id, @Nullable AccessPackageCatalogRoleAssignmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AccessPackageCatalogRoleAssignment(name, id, state, options);
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.DirectoryRoleAssignmentArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.DirectoryRoleAssignmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages a single directory role assignment within Azure Active Directory.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`
 * 
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`
 * 
 * ## Example Usage
 * 
 * *Assignment for a built-in role*
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
 * import com.pulumi.azuread.DirectoryRole;
 * import com.pulumi.azuread.DirectoryRoleArgs;
 * import com.pulumi.azuread.DirectoryRoleAssignment;
 * import com.pulumi.azuread.DirectoryRoleAssignmentArgs;
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
 *         final var example = AzureadFunctions.getUser(GetUserArgs.builder()
 *             .userPrincipalName("jdoe{@literal @}example.com")
 *             .build());
 * 
 *         var exampleDirectoryRole = new DirectoryRole("exampleDirectoryRole", DirectoryRoleArgs.builder()
 *             .displayName("Security administrator")
 *             .build());
 * 
 *         var exampleDirectoryRoleAssignment = new DirectoryRoleAssignment("exampleDirectoryRoleAssignment", DirectoryRoleAssignmentArgs.builder()
 *             .roleId(exampleDirectoryRole.templateId())
 *             .principalObjectId(example.applyValue(getUserResult -> getUserResult.objectId()))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &gt; Note the use of the `template_id` attribute when referencing built-in roles.
 * 
 * *Assignment for a custom role*
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
 * import com.pulumi.azuread.CustomDirectoryRole;
 * import com.pulumi.azuread.CustomDirectoryRoleArgs;
 * import com.pulumi.azuread.inputs.CustomDirectoryRolePermissionArgs;
 * import com.pulumi.azuread.DirectoryRoleAssignment;
 * import com.pulumi.azuread.DirectoryRoleAssignmentArgs;
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
 *         final var example = AzureadFunctions.getUser(GetUserArgs.builder()
 *             .userPrincipalName("jdoe{@literal @}example.com")
 *             .build());
 * 
 *         var exampleCustomDirectoryRole = new CustomDirectoryRole("exampleCustomDirectoryRole", CustomDirectoryRoleArgs.builder()
 *             .displayName("My Custom Role")
 *             .enabled(true)
 *             .version("1.0")
 *             .permissions(CustomDirectoryRolePermissionArgs.builder()
 *                 .allowedResourceActions(                
 *                     "microsoft.directory/applications/basic/update",
 *                     "microsoft.directory/applications/standard/read")
 *                 .build())
 *             .build());
 * 
 *         var exampleDirectoryRoleAssignment = new DirectoryRoleAssignment("exampleDirectoryRoleAssignment", DirectoryRoleAssignmentArgs.builder()
 *             .roleId(exampleCustomDirectoryRole.objectId())
 *             .principalObjectId(example.applyValue(getUserResult -> getUserResult.objectId()))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * *Scoped assignment for an application*
 * 
 * ## Import
 * 
 * Directory role assignments can be imported using the ID of the assignment, e.g.
 * 
 * ```sh
 * $ pulumi import azuread:index/directoryRoleAssignment:DirectoryRoleAssignment example ePROZI_iKE653D_d6aoLHyr-lKgHI8ZGiIdz8CLVcng-1
 * ```
 * 
 */
@ResourceType(type="azuread:index/directoryRoleAssignment:DirectoryRoleAssignment")
public class DirectoryRoleAssignment extends com.pulumi.resources.CustomResource {
    /**
     * Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with `directory_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&amp;tabs=http) for example usage. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="appScopeId", refs={String.class}, tree="[0]")
    private Output<String> appScopeId;

    /**
     * @return Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with `directory_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&amp;tabs=http) for example usage. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> appScopeId() {
        return this.appScopeId;
    }
    /**
     * Identifier of the app-specific scope when the assignment scope is app-specific
     * 
     * @deprecated
     * `app_scope_object_id` has been renamed to `app_scope_id` and will be removed in version 3.0 or the AzureAD Provider
     * 
     */
    @Deprecated /* `app_scope_object_id` has been renamed to `app_scope_id` and will be removed in version 3.0 or the AzureAD Provider */
    @Export(name="appScopeObjectId", refs={String.class}, tree="[0]")
    private Output<String> appScopeObjectId;

    /**
     * @return Identifier of the app-specific scope when the assignment scope is app-specific
     * 
     */
    public Output<String> appScopeObjectId() {
        return this.appScopeObjectId;
    }
    /**
     * Identifier of the directory object representing the scope of the assignment. Cannot be used with `app_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&amp;tabs=http) for example usage. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="directoryScopeId", refs={String.class}, tree="[0]")
    private Output<String> directoryScopeId;

    /**
     * @return Identifier of the directory object representing the scope of the assignment. Cannot be used with `app_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&amp;tabs=http) for example usage. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> directoryScopeId() {
        return this.directoryScopeId;
    }
    /**
     * Identifier of the directory object representing the scope of the assignment
     * 
     */
    @Export(name="directoryScopeObjectId", refs={String.class}, tree="[0]")
    private Output<String> directoryScopeObjectId;

    /**
     * @return Identifier of the directory object representing the scope of the assignment
     * 
     */
    public Output<String> directoryScopeObjectId() {
        return this.directoryScopeObjectId;
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
     * The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="roleId", refs={String.class}, tree="[0]")
    private Output<String> roleId;

    /**
     * @return The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> roleId() {
        return this.roleId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DirectoryRoleAssignment(String name) {
        this(name, DirectoryRoleAssignmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DirectoryRoleAssignment(String name, DirectoryRoleAssignmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DirectoryRoleAssignment(String name, DirectoryRoleAssignmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/directoryRoleAssignment:DirectoryRoleAssignment", name, args == null ? DirectoryRoleAssignmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DirectoryRoleAssignment(String name, Output<String> id, @Nullable DirectoryRoleAssignmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/directoryRoleAssignment:DirectoryRoleAssignment", name, state, makeResourceOptions(options, id));
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
    public static DirectoryRoleAssignment get(String name, Output<String> id, @Nullable DirectoryRoleAssignmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DirectoryRoleAssignment(name, id, state, options);
    }
}

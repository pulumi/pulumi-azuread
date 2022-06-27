// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.DirectoryRoleArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.DirectoryRoleState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages a Directory Role within Azure Active Directory. Directory Roles are also known as Administrator Roles.
 * 
 * Directory Roles are built-in to Azure Active Directory and are immutable. However, by default they are not activated in a tenant (except for the Global Administrator role). This resource ensures a directory role is activated from its associated role template, and exports the object ID of the role, so that role assignments can be made for it.
 * 
 * Once activated, directory roles cannot be deactivated and so this resource does not perform any actions on destroy.
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
 * *Activate a directory role by its template ID*
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new DirectoryRole(&#34;example&#34;, DirectoryRoleArgs.builder()        
 *             .templateId(&#34;00000000-0000-0000-0000-000000000000&#34;)
 *             .build());
 * 
 *         }
 * }
 * ```
 * 
 * *Activate a directory role by display name*
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new DirectoryRole(&#34;example&#34;, DirectoryRoleArgs.builder()        
 *             .displayName(&#34;Printer administrator&#34;)
 *             .build());
 * 
 *         }
 * }
 * ```
 * 
 * ## Import
 * 
 * This resource does not support importing.
 * 
 */
@ResourceType(type="azuread:index/directoryRole:DirectoryRole")
public class DirectoryRole extends com.pulumi.resources.CustomResource {
    /**
     * The description of the directory role.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output<String> description;

    /**
     * @return The description of the directory role.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The display name of the directory role to activate. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="displayName", type=String.class, parameters={})
    private Output<String> displayName;

    /**
     * @return The display name of the directory role to activate. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * The object ID of the directory role.
     * 
     */
    @Export(name="objectId", type=String.class, parameters={})
    private Output<String> objectId;

    /**
     * @return The object ID of the directory role.
     * 
     */
    public Output<String> objectId() {
        return this.objectId;
    }
    /**
     * The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="templateId", type=String.class, parameters={})
    private Output<String> templateId;

    /**
     * @return The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> templateId() {
        return this.templateId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DirectoryRole(String name) {
        this(name, DirectoryRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DirectoryRole(String name, @Nullable DirectoryRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DirectoryRole(String name, @Nullable DirectoryRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/directoryRole:DirectoryRole", name, args == null ? DirectoryRoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DirectoryRole(String name, Output<String> id, @Nullable DirectoryRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/directoryRole:DirectoryRole", name, state, makeResourceOptions(options, id));
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
    public static DirectoryRole get(String name, Output<String> id, @Nullable DirectoryRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DirectoryRole(name, id, state, options);
    }
}

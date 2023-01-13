// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.CustomDirectoryRoleArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.CustomDirectoryRoleState;
import com.pulumi.azuread.outputs.CustomDirectoryRolePermission;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.CustomDirectoryRole;
 * import com.pulumi.azuread.CustomDirectoryRoleArgs;
 * import com.pulumi.azuread.inputs.CustomDirectoryRolePermissionArgs;
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
 *         var example = new CustomDirectoryRole(&#34;example&#34;, CustomDirectoryRoleArgs.builder()        
 *             .description(&#34;Allows reading applications and updating groups&#34;)
 *             .displayName(&#34;My Custom Role&#34;)
 *             .enabled(true)
 *             .permissions(            
 *                 CustomDirectoryRolePermissionArgs.builder()
 *                     .allowedResourceActions(                    
 *                         &#34;microsoft.directory/applications/basic/update&#34;,
 *                         &#34;microsoft.directory/applications/create&#34;,
 *                         &#34;microsoft.directory/applications/standard/read&#34;)
 *                     .build(),
 *                 CustomDirectoryRolePermissionArgs.builder()
 *                     .allowedResourceActions(                    
 *                         &#34;microsoft.directory/groups/allProperties/read&#34;,
 *                         &#34;microsoft.directory/groups/allProperties/read&#34;,
 *                         &#34;microsoft.directory/groups/basic/update&#34;,
 *                         &#34;microsoft.directory/groups/create&#34;,
 *                         &#34;microsoft.directory/groups/delete&#34;)
 *                     .build())
 *             .version(&#34;1.0&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * This resource does not support importing.
 * 
 */
@ResourceType(type="azuread:index/customDirectoryRole:CustomDirectoryRole")
public class CustomDirectoryRole extends com.pulumi.resources.CustomResource {
    /**
     * The description of the custom directory role.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the custom directory role.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The display name of the custom directory role.
     * 
     */
    @Export(name="displayName", type=String.class, parameters={})
    private Output<String> displayName;

    /**
     * @return The display name of the custom directory role.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * Indicates whether the role is enabled for assignment.
     * 
     */
    @Export(name="enabled", type=Boolean.class, parameters={})
    private Output<Boolean> enabled;

    /**
     * @return Indicates whether the role is enabled for assignment.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }
    /**
     * The object ID of the custom directory role.
     * 
     */
    @Export(name="objectId", type=String.class, parameters={})
    private Output<String> objectId;

    /**
     * @return The object ID of the custom directory role.
     * 
     */
    public Output<String> objectId() {
        return this.objectId;
    }
    /**
     * A collection of `permissions` blocks as documented below.
     * 
     */
    @Export(name="permissions", type=List.class, parameters={CustomDirectoryRolePermission.class})
    private Output<List<CustomDirectoryRolePermission>> permissions;

    /**
     * @return A collection of `permissions` blocks as documented below.
     * 
     */
    public Output<List<CustomDirectoryRolePermission>> permissions() {
        return this.permissions;
    }
    /**
     * Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="templateId", type=String.class, parameters={})
    private Output<String> templateId;

    /**
     * @return Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> templateId() {
        return this.templateId;
    }
    /**
     * The version of the role definition. This can be any arbitrary string between 1-128 characters.
     * 
     */
    @Export(name="version", type=String.class, parameters={})
    private Output<String> version;

    /**
     * @return The version of the role definition. This can be any arbitrary string between 1-128 characters.
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CustomDirectoryRole(String name) {
        this(name, CustomDirectoryRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CustomDirectoryRole(String name, CustomDirectoryRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CustomDirectoryRole(String name, CustomDirectoryRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/customDirectoryRole:CustomDirectoryRole", name, args == null ? CustomDirectoryRoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CustomDirectoryRole(String name, Output<String> id, @Nullable CustomDirectoryRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/customDirectoryRole:CustomDirectoryRole", name, state, makeResourceOptions(options, id));
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
    public static CustomDirectoryRole get(String name, Output<String> id, @Nullable CustomDirectoryRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CustomDirectoryRole(name, id, state, options);
    }
}

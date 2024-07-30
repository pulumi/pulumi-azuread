// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.ApplicationPermissionScopeArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.ApplicationPermissionScopeState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
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
 * import com.pulumi.azuread.ApplicationRegistration;
 * import com.pulumi.azuread.ApplicationRegistrationArgs;
 * import com.pulumi.random.RandomUuid;
 * import com.pulumi.azuread.ApplicationPermissionScope;
 * import com.pulumi.azuread.ApplicationPermissionScopeArgs;
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
 *         var example = new ApplicationRegistration("example", ApplicationRegistrationArgs.builder()
 *             .displayName("example")
 *             .build());
 * 
 *         var exampleAdminister = new RandomUuid("exampleAdminister");
 * 
 *         var exampleApplicationPermissionScope = new ApplicationPermissionScope("exampleApplicationPermissionScope", ApplicationPermissionScopeArgs.builder()
 *             .applicationId(test.id())
 *             .scopeId(exampleAdminister.id())
 *             .value("administer")
 *             .adminConsentDescription("Administer the application")
 *             .adminConsentDisplayName("Administer")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &gt; **Tip** For managing more permissions scopes, create additional instances of this resource
 * 
 * *Usage with azuread.Application resource*
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.Application;
 * import com.pulumi.azuread.ApplicationArgs;
 * import com.pulumi.azuread.ApplicationPermissionScope;
 * import com.pulumi.azuread.ApplicationPermissionScopeArgs;
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
 *         var example = new Application("example", ApplicationArgs.builder()
 *             .displayName("example")
 *             .build());
 * 
 *         var exampleApplicationPermissionScope = new ApplicationPermissionScope("exampleApplicationPermissionScope", ApplicationPermissionScopeArgs.builder()
 *             .applicationId(example.id())
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
 * Application App Roles can be imported using the object ID of the application and the ID of the permission scope, in the following format.
 * 
 * ```sh
 * $ pulumi import azuread:index/applicationPermissionScope:ApplicationPermissionScope example /applications/00000000-0000-0000-0000-000000000000/permissionScopes/11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="azuread:index/applicationPermissionScope:ApplicationPermissionScope")
public class ApplicationPermissionScope extends com.pulumi.resources.CustomResource {
    /**
     * Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
     * 
     */
    @Export(name="adminConsentDescription", refs={String.class}, tree="[0]")
    private Output<String> adminConsentDescription;

    /**
     * @return Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
     * 
     */
    public Output<String> adminConsentDescription() {
        return this.adminConsentDescription;
    }
    /**
     * Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
     * 
     */
    @Export(name="adminConsentDisplayName", refs={String.class}, tree="[0]")
    private Output<String> adminConsentDisplayName;

    /**
     * @return Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
     * 
     */
    public Output<String> adminConsentDisplayName() {
        return this.adminConsentDisplayName;
    }
    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="applicationId", refs={String.class}, tree="[0]")
    private Output<String> applicationId;

    /**
     * @return The resource ID of the application registration. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> applicationId() {
        return this.applicationId;
    }
    /**
     * The unique identifier of the permission scope. Must be a valid UUID. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="scopeId", refs={String.class}, tree="[0]")
    private Output<String> scopeId;

    /**
     * @return The unique identifier of the permission scope. Must be a valid UUID. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> scopeId() {
        return this.scopeId;
    }
    /**
     * Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }
    /**
     * Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
     * 
     */
    @Export(name="userConsentDescription", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userConsentDescription;

    /**
     * @return Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
     * 
     */
    public Output<Optional<String>> userConsentDescription() {
        return Codegen.optional(this.userConsentDescription);
    }
    /**
     * Display name for the delegated permission that appears in the end user consent experience
     * 
     */
    @Export(name="userConsentDisplayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userConsentDisplayName;

    /**
     * @return Display name for the delegated permission that appears in the end user consent experience
     * 
     */
    public Output<Optional<String>> userConsentDisplayName() {
        return Codegen.optional(this.userConsentDisplayName);
    }
    /**
     * The value that is used for the `scp` claim in OAuth access tokens.
     * 
     * &gt; **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
     * 
     */
    @Export(name="value", refs={String.class}, tree="[0]")
    private Output<String> value;

    /**
     * @return The value that is used for the `scp` claim in OAuth access tokens.
     * 
     * &gt; **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ApplicationPermissionScope(String name) {
        this(name, ApplicationPermissionScopeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ApplicationPermissionScope(String name, ApplicationPermissionScopeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ApplicationPermissionScope(String name, ApplicationPermissionScopeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/applicationPermissionScope:ApplicationPermissionScope", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private ApplicationPermissionScope(String name, Output<String> id, @Nullable ApplicationPermissionScopeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/applicationPermissionScope:ApplicationPermissionScope", name, state, makeResourceOptions(options, id));
    }

    private static ApplicationPermissionScopeArgs makeArgs(ApplicationPermissionScopeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ApplicationPermissionScopeArgs.Empty : args;
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
    public static ApplicationPermissionScope get(String name, Output<String> id, @Nullable ApplicationPermissionScopeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ApplicationPermissionScope(name, id, state, options);
    }
}

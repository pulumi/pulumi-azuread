// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.ApplicationApiAccessArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.ApplicationApiAccessState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
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
 * import com.pulumi.azuread.AzureadFunctions;
 * import com.pulumi.azuread.inputs.GetServicePrincipalArgs;
 * import com.pulumi.azuread.ApplicationRegistration;
 * import com.pulumi.azuread.ApplicationRegistrationArgs;
 * import com.pulumi.azuread.ApplicationApiAccess;
 * import com.pulumi.azuread.ApplicationApiAccessArgs;
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
 *         final var wellKnown = AzureadFunctions.getApplicationPublishedAppIds();
 * 
 *         final var msgraph = AzureadFunctions.getServicePrincipal(GetServicePrincipalArgs.builder()
 *             .clientId(wellKnown.applyValue(getApplicationPublishedAppIdsResult -&gt; getApplicationPublishedAppIdsResult.result().MicrosoftGraph()))
 *             .build());
 * 
 *         var example = new ApplicationRegistration(&#34;example&#34;, ApplicationRegistrationArgs.builder()        
 *             .displayName(&#34;example&#34;)
 *             .build());
 * 
 *         var exampleMsgraph = new ApplicationApiAccess(&#34;exampleMsgraph&#34;, ApplicationApiAccessArgs.builder()        
 *             .applicationId(example.id())
 *             .apiClientId(wellKnown.applyValue(getApplicationPublishedAppIdsResult -&gt; getApplicationPublishedAppIdsResult.result().MicrosoftGraph()))
 *             .roleIds(            
 *                 msgraph.applyValue(getServicePrincipalResult -&gt; getServicePrincipalResult.appRoleIds().Group.Read.All()),
 *                 msgraph.applyValue(getServicePrincipalResult -&gt; getServicePrincipalResult.appRoleIds().User.Read.All()))
 *             .scopeIds(msgraph.applyValue(getServicePrincipalResult -&gt; getServicePrincipalResult.oauth2PermissionScopeIds().User.ReadWrite()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * &gt; **Tip** For managing permissions for an additional API, create another instance of this resource
 * 
 * *Usage with azuread.Application resource*
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.Application;
 * import com.pulumi.azuread.ApplicationArgs;
 * import com.pulumi.azuread.ApplicationApiAccess;
 * import com.pulumi.azuread.ApplicationApiAccessArgs;
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
 *         var example = new Application(&#34;example&#34;, ApplicationArgs.builder()        
 *             .displayName(&#34;example&#34;)
 *             .build());
 * 
 *         var exampleApplicationApiAccess = new ApplicationApiAccess(&#34;exampleApplicationApiAccess&#34;, ApplicationApiAccessArgs.builder()        
 *             .applicationId(example.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Application API Access can be imported using the object ID of the application and the client ID of the API, in the following format.
 * 
 * ```sh
 * $ pulumi import azuread:index/applicationApiAccess:ApplicationApiAccess example /applications/00000000-0000-0000-0000-000000000000/apiAccess/11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="azuread:index/applicationApiAccess:ApplicationApiAccess")
public class ApplicationApiAccess extends com.pulumi.resources.CustomResource {
    /**
     * The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="apiClientId", refs={String.class}, tree="[0]")
    private Output<String> apiClientId;

    /**
     * @return The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> apiClientId() {
        return this.apiClientId;
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
     * A set of role IDs to be granted to the application, as published by the API.
     * 
     */
    @Export(name="roleIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> roleIds;

    /**
     * @return A set of role IDs to be granted to the application, as published by the API.
     * 
     */
    public Output<Optional<List<String>>> roleIds() {
        return Codegen.optional(this.roleIds);
    }
    /**
     * A set of scope IDs to be granted to the application, as published by the API.
     * 
     * &gt; At least one of `role_ids` or `scope_ids` must be specified.
     * 
     */
    @Export(name="scopeIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> scopeIds;

    /**
     * @return A set of scope IDs to be granted to the application, as published by the API.
     * 
     * &gt; At least one of `role_ids` or `scope_ids` must be specified.
     * 
     */
    public Output<Optional<List<String>>> scopeIds() {
        return Codegen.optional(this.scopeIds);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ApplicationApiAccess(String name) {
        this(name, ApplicationApiAccessArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ApplicationApiAccess(String name, ApplicationApiAccessArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ApplicationApiAccess(String name, ApplicationApiAccessArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/applicationApiAccess:ApplicationApiAccess", name, args == null ? ApplicationApiAccessArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ApplicationApiAccess(String name, Output<String> id, @Nullable ApplicationApiAccessState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/applicationApiAccess:ApplicationApiAccess", name, state, makeResourceOptions(options, id));
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
    public static ApplicationApiAccess get(String name, Output<String> id, @Nullable ApplicationApiAccessState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ApplicationApiAccess(name, id, state, options);
    }
}

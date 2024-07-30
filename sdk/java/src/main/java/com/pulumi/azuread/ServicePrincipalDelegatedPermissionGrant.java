// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.ServicePrincipalDelegatedPermissionGrantArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.ServicePrincipalDelegatedPermissionGrantState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a delegated permission grant for a service principal, on behalf of a single user, or all users.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires the following application role: `Directory.ReadWrite.All`
 * 
 * When authenticated with a user principal, this resource requires one the following directory role: `Global Administrator`
 * 
 * ## Example Usage
 * 
 * *Delegated permission grant for all users*
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
 * import com.pulumi.azuread.ServicePrincipal;
 * import com.pulumi.azuread.ServicePrincipalArgs;
 * import com.pulumi.azuread.Application;
 * import com.pulumi.azuread.ApplicationArgs;
 * import com.pulumi.azuread.inputs.ApplicationRequiredResourceAccessArgs;
 * import com.pulumi.azuread.ServicePrincipalDelegatedPermissionGrant;
 * import com.pulumi.azuread.ServicePrincipalDelegatedPermissionGrantArgs;
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
 *         var msgraph = new ServicePrincipal("msgraph", ServicePrincipalArgs.builder()
 *             .clientId(wellKnown.applyValue(getApplicationPublishedAppIdsResult -> getApplicationPublishedAppIdsResult.result().microsoftGraph()))
 *             .useExisting(true)
 *             .build());
 * 
 *         var example = new Application("example", ApplicationArgs.builder()
 *             .displayName("example")
 *             .requiredResourceAccesses(ApplicationRequiredResourceAccessArgs.builder()
 *                 .resourceAppId(wellKnown.applyValue(getApplicationPublishedAppIdsResult -> getApplicationPublishedAppIdsResult.result().microsoftGraph()))
 *                 .resourceAccesses(                
 *                     ApplicationRequiredResourceAccessResourceAccessArgs.builder()
 *                         .id(msgraph.oauth2PermissionScopeIds().applyValue(oauth2PermissionScopeIds -> oauth2PermissionScopeIds.openid()))
 *                         .type("Scope")
 *                         .build(),
 *                     ApplicationRequiredResourceAccessResourceAccessArgs.builder()
 *                         .id(msgraph.oauth2PermissionScopeIds().applyValue(oauth2PermissionScopeIds -> oauth2PermissionScopeIds.User.Read()))
 *                         .type("Scope")
 *                         .build())
 *                 .build())
 *             .build());
 * 
 *         var exampleServicePrincipal = new ServicePrincipal("exampleServicePrincipal", ServicePrincipalArgs.builder()
 *             .clientId(example.applicationId())
 *             .build());
 * 
 *         var exampleServicePrincipalDelegatedPermissionGrant = new ServicePrincipalDelegatedPermissionGrant("exampleServicePrincipalDelegatedPermissionGrant", ServicePrincipalDelegatedPermissionGrantArgs.builder()
 *             .servicePrincipalObjectId(exampleServicePrincipal.objectId())
 *             .resourceServicePrincipalObjectId(msgraph.objectId())
 *             .claimValues(            
 *                 "openid",
 *                 "User.Read.All")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * *Delegated permission grant for a single user*
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
 * import com.pulumi.azuread.ServicePrincipal;
 * import com.pulumi.azuread.ServicePrincipalArgs;
 * import com.pulumi.azuread.Application;
 * import com.pulumi.azuread.ApplicationArgs;
 * import com.pulumi.azuread.inputs.ApplicationRequiredResourceAccessArgs;
 * import com.pulumi.azuread.User;
 * import com.pulumi.azuread.UserArgs;
 * import com.pulumi.azuread.ServicePrincipalDelegatedPermissionGrant;
 * import com.pulumi.azuread.ServicePrincipalDelegatedPermissionGrantArgs;
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
 *         var msgraph = new ServicePrincipal("msgraph", ServicePrincipalArgs.builder()
 *             .clientId(wellKnown.applyValue(getApplicationPublishedAppIdsResult -> getApplicationPublishedAppIdsResult.result().microsoftGraph()))
 *             .useExisting(true)
 *             .build());
 * 
 *         var example = new Application("example", ApplicationArgs.builder()
 *             .displayName("example")
 *             .requiredResourceAccesses(ApplicationRequiredResourceAccessArgs.builder()
 *                 .resourceAppId(wellKnown.applyValue(getApplicationPublishedAppIdsResult -> getApplicationPublishedAppIdsResult.result().microsoftGraph()))
 *                 .resourceAccesses(                
 *                     ApplicationRequiredResourceAccessResourceAccessArgs.builder()
 *                         .id(msgraph.oauth2PermissionScopeIds().applyValue(oauth2PermissionScopeIds -> oauth2PermissionScopeIds.openid()))
 *                         .type("Scope")
 *                         .build(),
 *                     ApplicationRequiredResourceAccessResourceAccessArgs.builder()
 *                         .id(msgraph.oauth2PermissionScopeIds().applyValue(oauth2PermissionScopeIds -> oauth2PermissionScopeIds.User.Read()))
 *                         .type("Scope")
 *                         .build())
 *                 .build())
 *             .build());
 * 
 *         var exampleServicePrincipal = new ServicePrincipal("exampleServicePrincipal", ServicePrincipalArgs.builder()
 *             .clientId(example.applicationId())
 *             .build());
 * 
 *         var exampleUser = new User("exampleUser", UserArgs.builder()
 *             .displayName("J. Doe")
 *             .userPrincipalName("jdoe{@literal @}example.com")
 *             .mailNickname("jdoe")
 *             .password("SecretP{@literal @}sswd99!")
 *             .build());
 * 
 *         var exampleServicePrincipalDelegatedPermissionGrant = new ServicePrincipalDelegatedPermissionGrant("exampleServicePrincipalDelegatedPermissionGrant", ServicePrincipalDelegatedPermissionGrantArgs.builder()
 *             .servicePrincipalObjectId(exampleServicePrincipal.objectId())
 *             .resourceServicePrincipalObjectId(msgraph.objectId())
 *             .claimValues(            
 *                 "openid",
 *                 "User.Read.All")
 *             .userObjectId(exampleUser.objectId())
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
 * Delegated permission grants can be imported using their ID, e.g.
 * 
 * ```sh
 * $ pulumi import azuread:index/servicePrincipalDelegatedPermissionGrant:ServicePrincipalDelegatedPermissionGrant example aaBBcDDeFG6h5JKLMN2PQrrssTTUUvWWxxxxxyyyzzz
 * ```
 * 
 */
@ResourceType(type="azuread:index/servicePrincipalDelegatedPermissionGrant:ServicePrincipalDelegatedPermissionGrant")
public class ServicePrincipalDelegatedPermissionGrant extends com.pulumi.resources.CustomResource {
    /**
     * A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
     * 
     */
    @Export(name="claimValues", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> claimValues;

    /**
     * @return A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
     * 
     */
    public Output<List<String>> claimValues() {
        return this.claimValues;
    }
    /**
     * The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="resourceServicePrincipalObjectId", refs={String.class}, tree="[0]")
    private Output<String> resourceServicePrincipalObjectId;

    /**
     * @return The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> resourceServicePrincipalObjectId() {
        return this.resourceServicePrincipalObjectId;
    }
    /**
     * The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="servicePrincipalObjectId", refs={String.class}, tree="[0]")
    private Output<String> servicePrincipalObjectId;

    /**
     * @return The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> servicePrincipalObjectId() {
        return this.servicePrincipalObjectId;
    }
    /**
     * The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
     * 
     * &gt; **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `user_object_id` property.
     * 
     */
    @Export(name="userObjectId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userObjectId;

    /**
     * @return The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
     * 
     * &gt; **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `user_object_id` property.
     * 
     */
    public Output<Optional<String>> userObjectId() {
        return Codegen.optional(this.userObjectId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServicePrincipalDelegatedPermissionGrant(String name) {
        this(name, ServicePrincipalDelegatedPermissionGrantArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServicePrincipalDelegatedPermissionGrant(String name, ServicePrincipalDelegatedPermissionGrantArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServicePrincipalDelegatedPermissionGrant(String name, ServicePrincipalDelegatedPermissionGrantArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/servicePrincipalDelegatedPermissionGrant:ServicePrincipalDelegatedPermissionGrant", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private ServicePrincipalDelegatedPermissionGrant(String name, Output<String> id, @Nullable ServicePrincipalDelegatedPermissionGrantState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/servicePrincipalDelegatedPermissionGrant:ServicePrincipalDelegatedPermissionGrant", name, state, makeResourceOptions(options, id));
    }

    private static ServicePrincipalDelegatedPermissionGrantArgs makeArgs(ServicePrincipalDelegatedPermissionGrantArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServicePrincipalDelegatedPermissionGrantArgs.Empty : args;
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
    public static ServicePrincipalDelegatedPermissionGrant get(String name, Output<String> id, @Nullable ServicePrincipalDelegatedPermissionGrantState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServicePrincipalDelegatedPermissionGrant(name, id, state, options);
    }
}

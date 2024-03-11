// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.ApplicationRedirectUrisArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.ApplicationRedirectUrisState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.ApplicationRegistration;
 * import com.pulumi.azuread.ApplicationRegistrationArgs;
 * import com.pulumi.azuread.ApplicationRedirectUris;
 * import com.pulumi.azuread.ApplicationRedirectUrisArgs;
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
 *         var example = new ApplicationRegistration(&#34;example&#34;, ApplicationRegistrationArgs.builder()        
 *             .displayName(&#34;example&#34;)
 *             .build());
 * 
 *         var examplePublic = new ApplicationRedirectUris(&#34;examplePublic&#34;, ApplicationRedirectUrisArgs.builder()        
 *             .applicationId(example.id())
 *             .type(&#34;PublicClient&#34;)
 *             .redirectUris(            
 *                 &#34;myapp://auth&#34;,
 *                 &#34;sample.mobile.app.bundie.id://auth&#34;,
 *                 &#34;https://login.microsoftonline.com/common/oauth2/nativeclient&#34;,
 *                 &#34;https://login.live.com/oauth20_desktop.srf&#34;,
 *                 &#34;ms-appx-web://Microsoft.AAD.BrokerPlugin/00000000-1111-1111-1111-222222222222&#34;,
 *                 &#34;urn:ietf:wg:oauth:2.0:foo&#34;)
 *             .build());
 * 
 *         var exampleSpa = new ApplicationRedirectUris(&#34;exampleSpa&#34;, ApplicationRedirectUrisArgs.builder()        
 *             .applicationId(example.id())
 *             .type(&#34;SPA&#34;)
 *             .redirectUris(            
 *                 &#34;https://mobile.hashitown.com/&#34;,
 *                 &#34;https://beta.hashitown.com/&#34;)
 *             .build());
 * 
 *         var exampleWeb = new ApplicationRedirectUris(&#34;exampleWeb&#34;, ApplicationRedirectUrisArgs.builder()        
 *             .applicationId(example.id())
 *             .type(&#34;Web&#34;)
 *             .redirectUris(            
 *                 &#34;https://app.hashitown.com/&#34;,
 *                 &#34;https://classic.hashitown.com/&#34;,
 *                 &#34;urn:ietf:wg:oauth:2.0:oob&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Application API Access can be imported using the object ID of the application and the URI type, in the following format.
 * 
 * ```sh
 * $ pulumi import azuread:index/applicationRedirectUris:ApplicationRedirectUris example /applications/00000000-0000-0000-0000-000000000000/redirectUris/Web
 * ```
 * 
 */
@ResourceType(type="azuread:index/applicationRedirectUris:ApplicationRedirectUris")
public class ApplicationRedirectUris extends com.pulumi.resources.CustomResource {
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
     * A set of redirect URIs to assign to the application.
     * 
     */
    @Export(name="redirectUris", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> redirectUris;

    /**
     * @return A set of redirect URIs to assign to the application.
     * 
     */
    public Output<List<String>> redirectUris() {
        return this.redirectUris;
    }
    /**
     * The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ApplicationRedirectUris(String name) {
        this(name, ApplicationRedirectUrisArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ApplicationRedirectUris(String name, ApplicationRedirectUrisArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ApplicationRedirectUris(String name, ApplicationRedirectUrisArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/applicationRedirectUris:ApplicationRedirectUris", name, args == null ? ApplicationRedirectUrisArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ApplicationRedirectUris(String name, Output<String> id, @Nullable ApplicationRedirectUrisState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/applicationRedirectUris:ApplicationRedirectUris", name, state, makeResourceOptions(options, id));
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
    public static ApplicationRedirectUris get(String name, Output<String> id, @Nullable ApplicationRedirectUrisState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ApplicationRedirectUris(name, id, state, options);
    }
}

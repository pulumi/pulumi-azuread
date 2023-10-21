// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.ApplicationKnownClientsArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.ApplicationKnownClientsState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.ApplicationRegistration;
 * import com.pulumi.azuread.ApplicationRegistrationArgs;
 * import com.pulumi.azuread.ApplicationKnownClients;
 * import com.pulumi.azuread.ApplicationKnownClientsArgs;
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
 *         var exampleApplicationRegistration = new ApplicationRegistration(&#34;exampleApplicationRegistration&#34;, ApplicationRegistrationArgs.builder()        
 *             .displayName(&#34;example&#34;)
 *             .build());
 * 
 *         var client = new ApplicationRegistration(&#34;client&#34;, ApplicationRegistrationArgs.builder()        
 *             .displayName(&#34;example client&#34;)
 *             .build());
 * 
 *         var exampleApplicationKnownClients = new ApplicationKnownClients(&#34;exampleApplicationKnownClients&#34;, ApplicationKnownClientsArgs.builder()        
 *             .applicationId(exampleApplicationRegistration.id())
 *             .knownClientIds(client.clientId())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Application Known Clients can be imported using the object ID of the application in the following format.
 * 
 * ```sh
 *  $ pulumi import azuread:index/applicationKnownClients:ApplicationKnownClients example /applications/00000000-0000-0000-0000-000000000000/knownClients
 * ```
 * 
 */
@ResourceType(type="azuread:index/applicationKnownClients:ApplicationKnownClients")
public class ApplicationKnownClients extends com.pulumi.resources.CustomResource {
    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="applicationId", type=String.class, parameters={})
    private Output<String> applicationId;

    /**
     * @return The resource ID of the application registration. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> applicationId() {
        return this.applicationId;
    }
    /**
     * A set of client IDs for the known applications.
     * 
     */
    @Export(name="knownClientIds", type=List.class, parameters={String.class})
    private Output<List<String>> knownClientIds;

    /**
     * @return A set of client IDs for the known applications.
     * 
     */
    public Output<List<String>> knownClientIds() {
        return this.knownClientIds;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ApplicationKnownClients(String name) {
        this(name, ApplicationKnownClientsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ApplicationKnownClients(String name, ApplicationKnownClientsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ApplicationKnownClients(String name, ApplicationKnownClientsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/applicationKnownClients:ApplicationKnownClients", name, args == null ? ApplicationKnownClientsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ApplicationKnownClients(String name, Output<String> id, @Nullable ApplicationKnownClientsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/applicationKnownClients:ApplicationKnownClients", name, state, makeResourceOptions(options, id));
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
    public static ApplicationKnownClients get(String name, Output<String> id, @Nullable ApplicationKnownClientsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ApplicationKnownClients(name, id, state, options);
    }
}

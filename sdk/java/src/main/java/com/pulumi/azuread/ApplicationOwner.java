// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.ApplicationOwnerArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.ApplicationOwnerState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
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
 * import com.pulumi.azuread.User;
 * import com.pulumi.azuread.UserArgs;
 * import com.pulumi.azuread.ApplicationOwner;
 * import com.pulumi.azuread.ApplicationOwnerArgs;
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
 *         var example = new ApplicationRegistration("example", ApplicationRegistrationArgs.builder()
 *             .displayName("example")
 *             .build());
 * 
 *         var jane = new User("jane", UserArgs.builder()
 *             .userPrincipalName("jane.fischer}{@literal @}{@code example.com")
 *             .displayName("Jane Fischer")
 *             .password("Ch}{@literal @}{@code ngeMe")
 *             .build());
 * 
 *         var exampleJane = new ApplicationOwner("exampleJane", ApplicationOwnerArgs.builder()
 *             .applicationId(example.id())
 *             .ownerObjectId(jane.objectId())
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &gt; **Tip** For managing more application owners, create additional instances of this resource
 * 
 * ## Import
 * 
 * Application Owners can be imported using the object ID of the application and the object ID of the owner, in the following format.
 * 
 * ```sh
 * $ pulumi import azuread:index/applicationOwner:ApplicationOwner example /applications/00000000-0000-0000-0000-000000000000/owners/11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="azuread:index/applicationOwner:ApplicationOwner")
public class ApplicationOwner extends com.pulumi.resources.CustomResource {
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
     * The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="ownerObjectId", refs={String.class}, tree="[0]")
    private Output<String> ownerObjectId;

    /**
     * @return The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> ownerObjectId() {
        return this.ownerObjectId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ApplicationOwner(java.lang.String name) {
        this(name, ApplicationOwnerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ApplicationOwner(java.lang.String name, ApplicationOwnerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ApplicationOwner(java.lang.String name, ApplicationOwnerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/applicationOwner:ApplicationOwner", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ApplicationOwner(java.lang.String name, Output<java.lang.String> id, @Nullable ApplicationOwnerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/applicationOwner:ApplicationOwner", name, state, makeResourceOptions(options, id), false);
    }

    private static ApplicationOwnerArgs makeArgs(ApplicationOwnerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ApplicationOwnerArgs.Empty : args;
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
    public static ApplicationOwner get(java.lang.String name, Output<java.lang.String> id, @Nullable ApplicationOwnerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ApplicationOwner(name, id, state, options);
    }
}

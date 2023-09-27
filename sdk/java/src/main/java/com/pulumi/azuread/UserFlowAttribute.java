// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.UserFlowAttributeArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.UserFlowAttributeState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages user flow attributes in an Azure Active Directory (Azure AD) tenant.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires the following application role: `IdentityUserFlow.ReadWrite.All`
 * 
 * ## Example Usage
 * 
 * *Basic example*
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.UserFlowAttribute;
 * import com.pulumi.azuread.UserFlowAttributeArgs;
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
 *         var example = new UserFlowAttribute(&#34;example&#34;, UserFlowAttributeArgs.builder()        
 *             .dataType(&#34;string&#34;)
 *             .description(&#34;Your hobby&#34;)
 *             .displayName(&#34;Hobby&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * User flow attributes can be imported using the `id`, e.g.
 * 
 * ```sh
 *  $ pulumi import azuread:index/userFlowAttribute:UserFlowAttribute example extension_ecc9f88db2924942b8a96f44873616fe_Hobbyjkorv
 * ```
 * 
 *  -&gt; This ID can be queried using the [User Flow Attributes API](https://learn.microsoft.com/en-us/graph/api/identityuserflowattribute-list?view=graph-rest-1.0&amp;tabs=http).
 * 
 */
@ResourceType(type="azuread:index/userFlowAttribute:UserFlowAttribute")
public class UserFlowAttribute extends com.pulumi.resources.CustomResource {
    /**
     * The type of the user flow attribute. Values include `builtIn`, `custom` or `required`.
     * 
     */
    @Export(name="attributeType", refs={String.class}, tree="[0]")
    private Output<String> attributeType;

    /**
     * @return The type of the user flow attribute. Values include `builtIn`, `custom` or `required`.
     * 
     */
    public Output<String> attributeType() {
        return this.attributeType;
    }
    /**
     * The data type of the user flow attribute. Possible values are `boolean`, `dateTime`, `int64`, `string` or `stringCollection`. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="dataType", refs={String.class}, tree="[0]")
    private Output<String> dataType;

    /**
     * @return The data type of the user flow attribute. Possible values are `boolean`, `dateTime`, `int64`, `string` or `stringCollection`. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> dataType() {
        return this.dataType;
    }
    /**
     * The description of the user flow attribute that is shown to the user at the time of sign-up.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return The description of the user flow attribute that is shown to the user at the time of sign-up.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The display name of the user flow attribute. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The display name of the user flow attribute. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UserFlowAttribute(String name) {
        this(name, UserFlowAttributeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UserFlowAttribute(String name, UserFlowAttributeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UserFlowAttribute(String name, UserFlowAttributeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/userFlowAttribute:UserFlowAttribute", name, args == null ? UserFlowAttributeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private UserFlowAttribute(String name, Output<String> id, @Nullable UserFlowAttributeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/userFlowAttribute:UserFlowAttribute", name, state, makeResourceOptions(options, id));
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
    public static UserFlowAttribute get(String name, Output<String> id, @Nullable UserFlowAttributeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UserFlowAttribute(name, id, state, options);
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.ServicePrincipalPasswordArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.ServicePrincipalPasswordState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * *Basic example*
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
 * import com.pulumi.azuread.ServicePrincipal;
 * import com.pulumi.azuread.ServicePrincipalArgs;
 * import com.pulumi.azuread.ServicePrincipalPassword;
 * import com.pulumi.azuread.ServicePrincipalPasswordArgs;
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
 *         var exampleServicePrincipal = new ServicePrincipal("exampleServicePrincipal", ServicePrincipalArgs.builder()        
 *             .clientId(example.clientId())
 *             .build());
 * 
 *         var exampleServicePrincipalPassword = new ServicePrincipalPassword("exampleServicePrincipalPassword", ServicePrincipalPasswordArgs.builder()        
 *             .servicePrincipalId(exampleServicePrincipal.objectId())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * *Time-based rotation*
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
 * import com.pulumi.azuread.ServicePrincipal;
 * import com.pulumi.azuread.ServicePrincipalArgs;
 * import com.pulumi.time.Rotating;
 * import com.pulumi.time.RotatingArgs;
 * import com.pulumi.azuread.ServicePrincipalPassword;
 * import com.pulumi.azuread.ServicePrincipalPasswordArgs;
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
 *         var exampleServicePrincipal = new ServicePrincipal("exampleServicePrincipal", ServicePrincipalArgs.builder()        
 *             .clientId(example.clientId())
 *             .build());
 * 
 *         var exampleRotating = new Rotating("exampleRotating", RotatingArgs.builder()        
 *             .rotationDays(7)
 *             .build());
 * 
 *         var exampleServicePrincipalPassword = new ServicePrincipalPassword("exampleServicePrincipalPassword", ServicePrincipalPasswordArgs.builder()        
 *             .servicePrincipalId(exampleServicePrincipal.objectId())
 *             .rotateWhenChanged(Map.of("rotation", exampleRotating.id()))
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
 * This resource does not support importing.
 * 
 */
@ResourceType(type="azuread:index/servicePrincipalPassword:ServicePrincipalPassword")
public class ServicePrincipalPassword extends com.pulumi.resources.CustomResource {
    /**
     * A display name for the password.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return A display name for the password.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
     * 
     */
    @Export(name="endDate", refs={String.class}, tree="[0]")
    private Output<String> endDate;

    /**
     * @return The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
     * 
     */
    public Output<String> endDate() {
        return this.endDate;
    }
    /**
     * A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
     * 
     */
    @Export(name="endDateRelative", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> endDateRelative;

    /**
     * @return A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
     * 
     */
    public Output<Optional<String>> endDateRelative() {
        return Codegen.optional(this.endDateRelative);
    }
    /**
     * A UUID used to uniquely identify this password credential.
     * 
     */
    @Export(name="keyId", refs={String.class}, tree="[0]")
    private Output<String> keyId;

    /**
     * @return A UUID used to uniquely identify this password credential.
     * 
     */
    public Output<String> keyId() {
        return this.keyId;
    }
    /**
     * A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="rotateWhenChanged", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> rotateWhenChanged;

    /**
     * @return A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
     * 
     */
    public Output<Optional<Map<String,String>>> rotateWhenChanged() {
        return Codegen.optional(this.rotateWhenChanged);
    }
    /**
     * The object ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
     * 
     */
    @Export(name="servicePrincipalId", refs={String.class}, tree="[0]")
    private Output<String> servicePrincipalId;

    /**
     * @return The object ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
     * 
     */
    public Output<String> servicePrincipalId() {
        return this.servicePrincipalId;
    }
    /**
     * The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn&#39;t specified, the current date is used.  Changing this field forces a new resource to be created.
     * 
     */
    @Export(name="startDate", refs={String.class}, tree="[0]")
    private Output<String> startDate;

    /**
     * @return The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn&#39;t specified, the current date is used.  Changing this field forces a new resource to be created.
     * 
     */
    public Output<String> startDate() {
        return this.startDate;
    }
    /**
     * The password for this service principal, which is generated by Azure Active Directory.
     * 
     */
    @Export(name="value", refs={String.class}, tree="[0]")
    private Output<String> value;

    /**
     * @return The password for this service principal, which is generated by Azure Active Directory.
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServicePrincipalPassword(String name) {
        this(name, ServicePrincipalPasswordArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServicePrincipalPassword(String name, ServicePrincipalPasswordArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServicePrincipalPassword(String name, ServicePrincipalPasswordArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/servicePrincipalPassword:ServicePrincipalPassword", name, args == null ? ServicePrincipalPasswordArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServicePrincipalPassword(String name, Output<String> id, @Nullable ServicePrincipalPasswordState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/servicePrincipalPassword:ServicePrincipalPassword", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "value"
            ))
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
    public static ServicePrincipalPassword get(String name, Output<String> id, @Nullable ServicePrincipalPasswordState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServicePrincipalPassword(name, id, state, options);
    }
}

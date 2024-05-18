// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.ApplicationPasswordArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.ApplicationPasswordState;
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
 * import com.pulumi.azuread.ApplicationRegistration;
 * import com.pulumi.azuread.ApplicationRegistrationArgs;
 * import com.pulumi.azuread.ApplicationPassword;
 * import com.pulumi.azuread.ApplicationPasswordArgs;
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
 *         var exampleApplicationPassword = new ApplicationPassword("exampleApplicationPassword", ApplicationPasswordArgs.builder()
 *             .applicationId(example.id())
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
 * import com.pulumi.azuread.ApplicationRegistration;
 * import com.pulumi.azuread.ApplicationRegistrationArgs;
 * import com.pulumi.time.Rotating;
 * import com.pulumi.time.RotatingArgs;
 * import com.pulumi.azuread.ApplicationPassword;
 * import com.pulumi.azuread.ApplicationPasswordArgs;
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
 *         var exampleRotating = new Rotating("exampleRotating", RotatingArgs.builder()
 *             .rotationDays(7)
 *             .build());
 * 
 *         var exampleApplicationPassword = new ApplicationPassword("exampleApplicationPassword", ApplicationPasswordArgs.builder()
 *             .applicationId(example.id())
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
@ResourceType(type="azuread:index/applicationPassword:ApplicationPassword")
public class ApplicationPassword extends com.pulumi.resources.CustomResource {
    /**
     * The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
     * 
     */
    @Export(name="applicationId", refs={String.class}, tree="[0]")
    private Output<String> applicationId;

    /**
     * @return The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
     * 
     */
    public Output<String> applicationId() {
        return this.applicationId;
    }
    /**
     * The object ID of the application for which this password should be created
     * 
     * @deprecated
     * The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
     * 
     */
    @Deprecated /* The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider */
    @Export(name="applicationObjectId", refs={String.class}, tree="[0]")
    private Output<String> applicationObjectId;

    /**
     * @return The object ID of the application for which this password should be created
     * 
     */
    public Output<String> applicationObjectId() {
        return this.applicationObjectId;
    }
    /**
     * A display name for the password. Changing this field forces a new resource to be created.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return A display name for the password. Changing this field forces a new resource to be created.
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
     * The password for this application, which is generated by Azure Active Directory.
     * 
     */
    @Export(name="value", refs={String.class}, tree="[0]")
    private Output<String> value;

    /**
     * @return The password for this application, which is generated by Azure Active Directory.
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ApplicationPassword(String name) {
        this(name, ApplicationPasswordArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ApplicationPassword(String name, @Nullable ApplicationPasswordArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ApplicationPassword(String name, @Nullable ApplicationPasswordArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/applicationPassword:ApplicationPassword", name, args == null ? ApplicationPasswordArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ApplicationPassword(String name, Output<String> id, @Nullable ApplicationPasswordState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/applicationPassword:ApplicationPassword", name, state, makeResourceOptions(options, id));
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
    public static ApplicationPassword get(String name, Output<String> id, @Nullable ApplicationPasswordState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ApplicationPassword(name, id, state, options);
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.ConditionalAccessPolicyArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.ConditionalAccessPolicyState;
import com.pulumi.azuread.outputs.ConditionalAccessPolicyConditions;
import com.pulumi.azuread.outputs.ConditionalAccessPolicyGrantControls;
import com.pulumi.azuread.outputs.ConditionalAccessPolicySessionControls;
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
 * ### All users except guests or external users
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.ConditionalAccessPolicy;
 * import com.pulumi.azuread.ConditionalAccessPolicyArgs;
 * import com.pulumi.azuread.inputs.ConditionalAccessPolicyConditionsArgs;
 * import com.pulumi.azuread.inputs.ConditionalAccessPolicyConditionsApplicationsArgs;
 * import com.pulumi.azuread.inputs.ConditionalAccessPolicyConditionsDevicesArgs;
 * import com.pulumi.azuread.inputs.ConditionalAccessPolicyConditionsDevicesFilterArgs;
 * import com.pulumi.azuread.inputs.ConditionalAccessPolicyConditionsLocationsArgs;
 * import com.pulumi.azuread.inputs.ConditionalAccessPolicyConditionsPlatformsArgs;
 * import com.pulumi.azuread.inputs.ConditionalAccessPolicyConditionsUsersArgs;
 * import com.pulumi.azuread.inputs.ConditionalAccessPolicyGrantControlsArgs;
 * import com.pulumi.azuread.inputs.ConditionalAccessPolicySessionControlsArgs;
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
 *         var example = new ConditionalAccessPolicy("example", ConditionalAccessPolicyArgs.builder()
 *             .displayName("example policy")
 *             .state("disabled")
 *             .conditions(ConditionalAccessPolicyConditionsArgs.builder()
 *                 .clientAppTypes("all")
 *                 .signInRiskLevels("medium")
 *                 .userRiskLevels("medium")
 *                 .applications(ConditionalAccessPolicyConditionsApplicationsArgs.builder()
 *                     .includedApplications("All")
 *                     .excludedApplications()
 *                     .build())
 *                 .devices(ConditionalAccessPolicyConditionsDevicesArgs.builder()
 *                     .filter(ConditionalAccessPolicyConditionsDevicesFilterArgs.builder()
 *                         .mode("exclude")
 *                         .rule("device.operatingSystem eq \"Doors\"")
 *                         .build())
 *                     .build())
 *                 .locations(ConditionalAccessPolicyConditionsLocationsArgs.builder()
 *                     .includedLocations("All")
 *                     .excludedLocations("AllTrusted")
 *                     .build())
 *                 .platforms(ConditionalAccessPolicyConditionsPlatformsArgs.builder()
 *                     .includedPlatforms("android")
 *                     .excludedPlatforms("iOS")
 *                     .build())
 *                 .users(ConditionalAccessPolicyConditionsUsersArgs.builder()
 *                     .includedUsers("All")
 *                     .excludedUsers("GuestsOrExternalUsers")
 *                     .build())
 *                 .build())
 *             .grantControls(ConditionalAccessPolicyGrantControlsArgs.builder()
 *                 .operator("OR")
 *                 .builtInControls("mfa")
 *                 .build())
 *             .sessionControls(ConditionalAccessPolicySessionControlsArgs.builder()
 *                 .applicationEnforcedRestrictionsEnabled(true)
 *                 .disableResilienceDefaults(false)
 *                 .signInFrequency(10)
 *                 .signInFrequencyPeriod("hours")
 *                 .cloudAppSecurityPolicy("monitorOnly")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Included client applications / service principals
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
 * import com.pulumi.azuread.ConditionalAccessPolicy;
 * import com.pulumi.azuread.ConditionalAccessPolicyArgs;
 * import com.pulumi.azuread.inputs.ConditionalAccessPolicyConditionsArgs;
 * import com.pulumi.azuread.inputs.ConditionalAccessPolicyConditionsApplicationsArgs;
 * import com.pulumi.azuread.inputs.ConditionalAccessPolicyConditionsClientApplicationsArgs;
 * import com.pulumi.azuread.inputs.ConditionalAccessPolicyConditionsUsersArgs;
 * import com.pulumi.azuread.inputs.ConditionalAccessPolicyGrantControlsArgs;
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
 *         final var current = AzureadFunctions.getClientConfig();
 * 
 *         var example = new ConditionalAccessPolicy("example", ConditionalAccessPolicyArgs.builder()
 *             .displayName("example policy")
 *             .state("disabled")
 *             .conditions(ConditionalAccessPolicyConditionsArgs.builder()
 *                 .clientAppTypes("all")
 *                 .applications(ConditionalAccessPolicyConditionsApplicationsArgs.builder()
 *                     .includedApplications("All")
 *                     .build())
 *                 .clientApplications(ConditionalAccessPolicyConditionsClientApplicationsArgs.builder()
 *                     .includedServicePrincipals(current.applyValue(getClientConfigResult -> getClientConfigResult.objectId()))
 *                     .excludedServicePrincipals()
 *                     .build())
 *                 .users(ConditionalAccessPolicyConditionsUsersArgs.builder()
 *                     .includedUsers("None")
 *                     .build())
 *                 .build())
 *             .grantControls(ConditionalAccessPolicyGrantControlsArgs.builder()
 *                 .operator("OR")
 *                 .builtInControls("block")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Excluded client applications / service principals
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
 * import com.pulumi.azuread.ConditionalAccessPolicy;
 * import com.pulumi.azuread.ConditionalAccessPolicyArgs;
 * import com.pulumi.azuread.inputs.ConditionalAccessPolicyConditionsArgs;
 * import com.pulumi.azuread.inputs.ConditionalAccessPolicyConditionsApplicationsArgs;
 * import com.pulumi.azuread.inputs.ConditionalAccessPolicyConditionsClientApplicationsArgs;
 * import com.pulumi.azuread.inputs.ConditionalAccessPolicyConditionsUsersArgs;
 * import com.pulumi.azuread.inputs.ConditionalAccessPolicyGrantControlsArgs;
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
 *         final var current = AzureadFunctions.getClientConfig();
 * 
 *         var example = new ConditionalAccessPolicy("example", ConditionalAccessPolicyArgs.builder()
 *             .displayName("example policy")
 *             .state("disabled")
 *             .conditions(ConditionalAccessPolicyConditionsArgs.builder()
 *                 .clientAppTypes("all")
 *                 .applications(ConditionalAccessPolicyConditionsApplicationsArgs.builder()
 *                     .includedApplications("All")
 *                     .build())
 *                 .clientApplications(ConditionalAccessPolicyConditionsClientApplicationsArgs.builder()
 *                     .includedServicePrincipals("ServicePrincipalsInMyTenant")
 *                     .excludedServicePrincipals(current.applyValue(getClientConfigResult -> getClientConfigResult.objectId()))
 *                     .build())
 *                 .users(ConditionalAccessPolicyConditionsUsersArgs.builder()
 *                     .includedUsers("None")
 *                     .build())
 *                 .build())
 *             .grantControls(ConditionalAccessPolicyGrantControlsArgs.builder()
 *                 .operator("OR")
 *                 .builtInControls("block")
 *                 .build())
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
 * Conditional Access Policies can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy my_location 00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy")
public class ConditionalAccessPolicy extends com.pulumi.resources.CustomResource {
    /**
     * A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
     * 
     */
    @Export(name="conditions", refs={ConditionalAccessPolicyConditions.class}, tree="[0]")
    private Output<ConditionalAccessPolicyConditions> conditions;

    /**
     * @return A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
     * 
     */
    public Output<ConditionalAccessPolicyConditions> conditions() {
        return this.conditions;
    }
    /**
     * The friendly name for this Conditional Access Policy.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The friendly name for this Conditional Access Policy.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
     * 
     */
    @Export(name="grantControls", refs={ConditionalAccessPolicyGrantControls.class}, tree="[0]")
    private Output</* @Nullable */ ConditionalAccessPolicyGrantControls> grantControls;

    /**
     * @return A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
     * 
     */
    public Output<Optional<ConditionalAccessPolicyGrantControls>> grantControls() {
        return Codegen.optional(this.grantControls);
    }
    /**
     * The object ID of the policy
     * 
     */
    @Export(name="objectId", refs={String.class}, tree="[0]")
    private Output<String> objectId;

    /**
     * @return The object ID of the policy
     * 
     */
    public Output<String> objectId() {
        return this.objectId;
    }
    /**
     * A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.
     * 
     * &gt; Note: At least one of `grant_controls` and/or `session_controls` blocks must be specified.
     * 
     */
    @Export(name="sessionControls", refs={ConditionalAccessPolicySessionControls.class}, tree="[0]")
    private Output</* @Nullable */ ConditionalAccessPolicySessionControls> sessionControls;

    /**
     * @return A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.
     * 
     * &gt; Note: At least one of `grant_controls` and/or `session_controls` blocks must be specified.
     * 
     */
    public Output<Optional<ConditionalAccessPolicySessionControls>> sessionControls() {
        return Codegen.optional(this.sessionControls);
    }
    /**
     * Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
     * 
     */
    public Output<String> state() {
        return this.state;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ConditionalAccessPolicy(java.lang.String name) {
        this(name, ConditionalAccessPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ConditionalAccessPolicy(java.lang.String name, ConditionalAccessPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ConditionalAccessPolicy(java.lang.String name, ConditionalAccessPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ConditionalAccessPolicy(java.lang.String name, Output<java.lang.String> id, @Nullable ConditionalAccessPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy", name, state, makeResourceOptions(options, id), false);
    }

    private static ConditionalAccessPolicyArgs makeArgs(ConditionalAccessPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ConditionalAccessPolicyArgs.Empty : args;
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
    public static ConditionalAccessPolicy get(java.lang.String name, Output<java.lang.String> id, @Nullable ConditionalAccessPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ConditionalAccessPolicy(name, id, state, options);
    }
}

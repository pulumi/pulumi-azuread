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
 * Manages a Conditional Access Policy within Azure Active Directory.
 * 
 * &gt; **Licensing Requirements** Specifying `client_applications` property requires the activation of Microsoft Entra on your tenant and the availability of sufficient Workload Identities Premium licences (one per service principal managed by a conditional access).
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ConditionalAccess` and `Policy.Read.All`
 * 
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Conditional Access Administrator` or `Global Administrator`
 * 
 * ## Example Usage
 * ### All users except guests or external users
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.ConditionalAccessPolicy;
 * import com.pulumi.azuread.ConditionalAccessPolicyArgs;
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
 *         var example = new ConditionalAccessPolicy(&#34;example&#34;, ConditionalAccessPolicyArgs.builder()        
 *             .displayName(&#34;example policy&#34;)
 *             .state(&#34;disabled&#34;)
 *             .conditions(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
 *             .grantControls(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
 *             .sessionControls(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
 *             .build());
 * 
 *     }
 * }
 * ```
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
    public ConditionalAccessPolicy(String name) {
        this(name, ConditionalAccessPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ConditionalAccessPolicy(String name, ConditionalAccessPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ConditionalAccessPolicy(String name, ConditionalAccessPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy", name, args == null ? ConditionalAccessPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ConditionalAccessPolicy(String name, Output<String> id, @Nullable ConditionalAccessPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy", name, state, makeResourceOptions(options, id));
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
    public static ConditionalAccessPolicy get(String name, Output<String> id, @Nullable ConditionalAccessPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ConditionalAccessPolicy(name, id, state, options);
    }
}

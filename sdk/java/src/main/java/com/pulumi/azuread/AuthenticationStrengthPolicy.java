// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.AuthenticationStrengthPolicyArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.AuthenticationStrengthPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a Authentication Strength Policy within Azure Active Directory.
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
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.AuthenticationStrengthPolicy;
 * import com.pulumi.azuread.AuthenticationStrengthPolicyArgs;
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
 *         var example = new AuthenticationStrengthPolicy("example", AuthenticationStrengthPolicyArgs.builder()        
 *             .displayName("Example Authentication Strength Policy")
 *             .description("Policy for demo purposes")
 *             .allowedCombinations(            
 *                 "fido2",
 *                 "password")
 *             .build());
 * 
 *         var example2 = new AuthenticationStrengthPolicy("example2", AuthenticationStrengthPolicyArgs.builder()        
 *             .displayName("Example Authentication Strength Policy")
 *             .description("Policy for demo purposes with all possible combinations")
 *             .allowedCombinations(            
 *                 "fido2",
 *                 "password",
 *                 "deviceBasedPush",
 *                 "temporaryAccessPassOneTime",
 *                 "federatedMultiFactor",
 *                 "federatedSingleFactor",
 *                 "hardwareOath,federatedSingleFactor",
 *                 "microsoftAuthenticatorPush,federatedSingleFactor",
 *                 "password,hardwareOath",
 *                 "password,microsoftAuthenticatorPush",
 *                 "password,sms",
 *                 "password,softwareOath",
 *                 "password,voice",
 *                 "sms",
 *                 "sms,federatedSingleFactor",
 *                 "softwareOath,federatedSingleFactor",
 *                 "temporaryAccessPassMultiUse",
 *                 "voice,federatedSingleFactor",
 *                 "windowsHelloForBusiness",
 *                 "x509CertificateMultiFactor",
 *                 "x509CertificateSingleFactor")
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
 * Authentication Strength Policies can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import azuread:index/authenticationStrengthPolicy:AuthenticationStrengthPolicy my_policy 00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuread:index/authenticationStrengthPolicy:AuthenticationStrengthPolicy")
public class AuthenticationStrengthPolicy extends com.pulumi.resources.CustomResource {
    /**
     * List of allowed authentication methods for this authentication strength policy.
     * 
     */
    @Export(name="allowedCombinations", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> allowedCombinations;

    /**
     * @return List of allowed authentication methods for this authentication strength policy.
     * 
     */
    public Output<List<String>> allowedCombinations() {
        return this.allowedCombinations;
    }
    /**
     * The description for this authentication strength policy.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description for this authentication strength policy.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The friendly name for this authentication strength policy.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The friendly name for this authentication strength policy.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthenticationStrengthPolicy(String name) {
        this(name, AuthenticationStrengthPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthenticationStrengthPolicy(String name, AuthenticationStrengthPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthenticationStrengthPolicy(String name, AuthenticationStrengthPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/authenticationStrengthPolicy:AuthenticationStrengthPolicy", name, args == null ? AuthenticationStrengthPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuthenticationStrengthPolicy(String name, Output<String> id, @Nullable AuthenticationStrengthPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/authenticationStrengthPolicy:AuthenticationStrengthPolicy", name, state, makeResourceOptions(options, id));
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
    public static AuthenticationStrengthPolicy get(String name, Output<String> id, @Nullable AuthenticationStrengthPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthenticationStrengthPolicy(name, id, state, options);
    }
}

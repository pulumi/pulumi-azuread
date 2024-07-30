// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.ServicePrincipalClaimsMappingPolicyAssignmentArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.ServicePrincipalClaimsMappingPolicyAssignmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages a Claims Mapping Policy Assignment within Azure Active Directory.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ApplicationConfiguration` and `Policy.Read.All`
 * 
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`
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
 * import com.pulumi.azuread.ServicePrincipalClaimsMappingPolicyAssignment;
 * import com.pulumi.azuread.ServicePrincipalClaimsMappingPolicyAssignmentArgs;
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
 *         var app = new ServicePrincipalClaimsMappingPolicyAssignment("app", ServicePrincipalClaimsMappingPolicyAssignmentArgs.builder()
 *             .claimsMappingPolicyId(myPolicy.id())
 *             .servicePrincipalId(myPrincipal.id())
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
 * Claims Mapping Policy can be imported using the `id`, in the form `service-principal-uuid/claimsMappingPolicy/claims-mapping-policy-uuid`, e.g:
 * 
 * ```sh
 * $ pulumi import azuread:index/servicePrincipalClaimsMappingPolicyAssignment:ServicePrincipalClaimsMappingPolicyAssignment app 00000000-0000-0000-0000-000000000000/claimsMappingPolicy/11111111-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuread:index/servicePrincipalClaimsMappingPolicyAssignment:ServicePrincipalClaimsMappingPolicyAssignment")
public class ServicePrincipalClaimsMappingPolicyAssignment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the claims mapping policy to assign.
     * 
     */
    @Export(name="claimsMappingPolicyId", refs={String.class}, tree="[0]")
    private Output<String> claimsMappingPolicyId;

    /**
     * @return The ID of the claims mapping policy to assign.
     * 
     */
    public Output<String> claimsMappingPolicyId() {
        return this.claimsMappingPolicyId;
    }
    /**
     * The object ID of the service principal for the policy assignment.
     * 
     */
    @Export(name="servicePrincipalId", refs={String.class}, tree="[0]")
    private Output<String> servicePrincipalId;

    /**
     * @return The object ID of the service principal for the policy assignment.
     * 
     */
    public Output<String> servicePrincipalId() {
        return this.servicePrincipalId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServicePrincipalClaimsMappingPolicyAssignment(String name) {
        this(name, ServicePrincipalClaimsMappingPolicyAssignmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServicePrincipalClaimsMappingPolicyAssignment(String name, ServicePrincipalClaimsMappingPolicyAssignmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServicePrincipalClaimsMappingPolicyAssignment(String name, ServicePrincipalClaimsMappingPolicyAssignmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/servicePrincipalClaimsMappingPolicyAssignment:ServicePrincipalClaimsMappingPolicyAssignment", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private ServicePrincipalClaimsMappingPolicyAssignment(String name, Output<String> id, @Nullable ServicePrincipalClaimsMappingPolicyAssignmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/servicePrincipalClaimsMappingPolicyAssignment:ServicePrincipalClaimsMappingPolicyAssignment", name, state, makeResourceOptions(options, id));
    }

    private static ServicePrincipalClaimsMappingPolicyAssignmentArgs makeArgs(ServicePrincipalClaimsMappingPolicyAssignmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServicePrincipalClaimsMappingPolicyAssignmentArgs.Empty : args;
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
    public static ServicePrincipalClaimsMappingPolicyAssignment get(String name, Output<String> id, @Nullable ServicePrincipalClaimsMappingPolicyAssignmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServicePrincipalClaimsMappingPolicyAssignment(name, id, state, options);
    }
}

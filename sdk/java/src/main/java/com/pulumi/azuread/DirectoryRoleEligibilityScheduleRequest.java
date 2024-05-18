// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.DirectoryRoleEligibilityScheduleRequestArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.DirectoryRoleEligibilityScheduleRequestState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages a single directory role eligibility schedule request within Azure Active Directory.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * The calling principal requires one of the following application roles: `RoleEligibilitySchedule.ReadWrite.Directory` or `RoleManagement.ReadWrite.Directory`.
 * 
 * The calling principal requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`.
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
 * import com.pulumi.azuread.AzureadFunctions;
 * import com.pulumi.azuread.inputs.GetUserArgs;
 * import com.pulumi.azuread.DirectoryRole;
 * import com.pulumi.azuread.DirectoryRoleArgs;
 * import com.pulumi.azuread.DirectoryRoleEligibilityScheduleRequest;
 * import com.pulumi.azuread.DirectoryRoleEligibilityScheduleRequestArgs;
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
 *         final var example = AzureadFunctions.getUser(GetUserArgs.builder()
 *             .userPrincipalName("jdoe{@literal @}example.com")
 *             .build());
 * 
 *         var exampleDirectoryRole = new DirectoryRole("exampleDirectoryRole", DirectoryRoleArgs.builder()
 *             .displayName("Application Administrator")
 *             .build());
 * 
 *         var exampleDirectoryRoleEligibilityScheduleRequest = new DirectoryRoleEligibilityScheduleRequest("exampleDirectoryRoleEligibilityScheduleRequest", DirectoryRoleEligibilityScheduleRequestArgs.builder()
 *             .roleDefinitionId(exampleDirectoryRole.templateId())
 *             .principalId(exampleAzureadUser.objectId())
 *             .directoryScopeId("/")
 *             .justification("Example")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &gt; Note the use of the `template_id` attribute when referencing built-in roles.
 * 
 * ## Import
 * 
 * Directory role eligibility schedule requests can be imported using the ID of the assignment, e.g.
 * 
 * ```sh
 * $ pulumi import azuread:index/directoryRoleEligibilityScheduleRequest:DirectoryRoleEligibilityScheduleRequest example 822ec710-4c9f-4f71-a27a-451759cc7522
 * ```
 * 
 */
@ResourceType(type="azuread:index/directoryRoleEligibilityScheduleRequest:DirectoryRoleEligibilityScheduleRequest")
public class DirectoryRoleEligibilityScheduleRequest extends com.pulumi.resources.CustomResource {
    /**
     * Identifier of the directory object representing the scope of the role eligibility. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="directoryScopeId", refs={String.class}, tree="[0]")
    private Output<String> directoryScopeId;

    /**
     * @return Identifier of the directory object representing the scope of the role eligibility. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> directoryScopeId() {
        return this.directoryScopeId;
    }
    /**
     * Justification for why the principal is granted the role eligibility. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="justification", refs={String.class}, tree="[0]")
    private Output<String> justification;

    /**
     * @return Justification for why the principal is granted the role eligibility. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> justification() {
        return this.justification;
    }
    /**
     * The object ID of the principal to granted the role eligibility. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="principalId", refs={String.class}, tree="[0]")
    private Output<String> principalId;

    /**
     * @return The object ID of the principal to granted the role eligibility. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> principalId() {
        return this.principalId;
    }
    /**
     * The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="roleDefinitionId", refs={String.class}, tree="[0]")
    private Output<String> roleDefinitionId;

    /**
     * @return The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> roleDefinitionId() {
        return this.roleDefinitionId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DirectoryRoleEligibilityScheduleRequest(String name) {
        this(name, DirectoryRoleEligibilityScheduleRequestArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DirectoryRoleEligibilityScheduleRequest(String name, DirectoryRoleEligibilityScheduleRequestArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DirectoryRoleEligibilityScheduleRequest(String name, DirectoryRoleEligibilityScheduleRequestArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/directoryRoleEligibilityScheduleRequest:DirectoryRoleEligibilityScheduleRequest", name, args == null ? DirectoryRoleEligibilityScheduleRequestArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DirectoryRoleEligibilityScheduleRequest(String name, Output<String> id, @Nullable DirectoryRoleEligibilityScheduleRequestState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/directoryRoleEligibilityScheduleRequest:DirectoryRoleEligibilityScheduleRequest", name, state, makeResourceOptions(options, id));
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
    public static DirectoryRoleEligibilityScheduleRequest get(String name, Output<String> id, @Nullable DirectoryRoleEligibilityScheduleRequestState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DirectoryRoleEligibilityScheduleRequest(name, id, state, options);
    }
}

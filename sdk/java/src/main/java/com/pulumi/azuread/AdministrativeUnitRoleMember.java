// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.AdministrativeUnitRoleMemberArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.AdministrativeUnitRoleMemberState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages a single directory role assignment scoped to an administrative unit within Azure Active Directory.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires one of the following application roles: `AdministrativeUnit.ReadWrite.All` and `RoleManagement.ReadWrite.Directory`, or `Directory.ReadWrite.All`
 * 
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`
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
 * import com.pulumi.azuread.AdministrativeUnit;
 * import com.pulumi.azuread.AdministrativeUnitArgs;
 * import com.pulumi.azuread.DirectoryRole;
 * import com.pulumi.azuread.DirectoryRoleArgs;
 * import com.pulumi.azuread.AdministrativeUnitRoleMember;
 * import com.pulumi.azuread.AdministrativeUnitRoleMemberArgs;
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
 *         var exampleAdministrativeUnit = new AdministrativeUnit("exampleAdministrativeUnit", AdministrativeUnitArgs.builder()
 *             .displayName("Example-AU")
 *             .build());
 * 
 *         var exampleDirectoryRole = new DirectoryRole("exampleDirectoryRole", DirectoryRoleArgs.builder()
 *             .displayName("Security administrator")
 *             .build());
 * 
 *         var exampleAdministrativeUnitRoleMember = new AdministrativeUnitRoleMember("exampleAdministrativeUnitRoleMember", AdministrativeUnitRoleMemberArgs.builder()
 *             .roleObjectId(exampleDirectoryRole.objectId())
 *             .administrativeUnitObjectId(exampleAdministrativeUnit.id())
 *             .memberObjectId(example.applyValue(getUserResult -> getUserResult.id()))
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
 * Administrative unit role members can be imported using the object ID of the administrative unit and the unique ID of the role assignment, e.g.
 * 
 * ```sh
 * $ pulumi import azuread:index/administrativeUnitRoleMember:AdministrativeUnitRoleMember example 00000000-0000-0000-0000-000000000000/roleMember/zX37MRLyF0uvE-xf2WH4B7x-6CPLfudNnxFGj800htpBXqkxW7bITqGb6Rj4kuTuS
 * ```
 * 
 * -&gt; This ID format is unique to Terraform and is composed of the Administrative Unit Object ID and the role assignment ID in the format `{AdministrativeUnitObjectID}/roleMember/{RoleAssignmentID}`.
 * 
 */
@ResourceType(type="azuread:index/administrativeUnitRoleMember:AdministrativeUnitRoleMember")
public class AdministrativeUnitRoleMember extends com.pulumi.resources.CustomResource {
    /**
     * The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="administrativeUnitObjectId", refs={String.class}, tree="[0]")
    private Output<String> administrativeUnitObjectId;

    /**
     * @return The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> administrativeUnitObjectId() {
        return this.administrativeUnitObjectId;
    }
    /**
     * The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="memberObjectId", refs={String.class}, tree="[0]")
    private Output<String> memberObjectId;

    /**
     * @return The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> memberObjectId() {
        return this.memberObjectId;
    }
    /**
     * The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="roleObjectId", refs={String.class}, tree="[0]")
    private Output<String> roleObjectId;

    /**
     * @return The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> roleObjectId() {
        return this.roleObjectId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AdministrativeUnitRoleMember(String name) {
        this(name, AdministrativeUnitRoleMemberArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AdministrativeUnitRoleMember(String name, AdministrativeUnitRoleMemberArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AdministrativeUnitRoleMember(String name, AdministrativeUnitRoleMemberArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/administrativeUnitRoleMember:AdministrativeUnitRoleMember", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private AdministrativeUnitRoleMember(String name, Output<String> id, @Nullable AdministrativeUnitRoleMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/administrativeUnitRoleMember:AdministrativeUnitRoleMember", name, state, makeResourceOptions(options, id));
    }

    private static AdministrativeUnitRoleMemberArgs makeArgs(AdministrativeUnitRoleMemberArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AdministrativeUnitRoleMemberArgs.Empty : args;
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
    public static AdministrativeUnitRoleMember get(String name, Output<String> id, @Nullable AdministrativeUnitRoleMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AdministrativeUnitRoleMember(name, id, state, options);
    }
}

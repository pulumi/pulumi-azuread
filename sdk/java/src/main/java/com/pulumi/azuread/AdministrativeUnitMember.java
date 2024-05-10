// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.AdministrativeUnitMemberArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.AdministrativeUnitMemberState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a single administrative unit membership within Azure Active Directory.
 * 
 * &gt; **Warning** Do not use this resource at the same time as the `members` property of the `azuread.AdministrativeUnit` resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires one of the following application roles: `AdministrativeUnit.ReadWrite.All` or `Directory.ReadWrite.All`
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
 * import com.pulumi.azuread.AdministrativeUnitMember;
 * import com.pulumi.azuread.AdministrativeUnitMemberArgs;
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
 *         var exampleAdministrativeUnitMember = new AdministrativeUnitMember("exampleAdministrativeUnitMember", AdministrativeUnitMemberArgs.builder()        
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
 * Administrative unit members can be imported using the object ID of the administrative unit and the object ID of the member, e.g.
 * 
 * ```sh
 * $ pulumi import azuread:index/administrativeUnitMember:AdministrativeUnitMember example 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
 * ```
 * 
 * -&gt; This ID format is unique to Terraform and is composed of the Administrative Unit Object ID and the target Member Object ID in the format `{AdministrativeUnitObjectID}/member/{MemberObjectID}`.
 * 
 */
@ResourceType(type="azuread:index/administrativeUnitMember:AdministrativeUnitMember")
public class AdministrativeUnitMember extends com.pulumi.resources.CustomResource {
    /**
     * The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="administrativeUnitObjectId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> administrativeUnitObjectId;

    /**
     * @return The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
     * 
     */
    public Output<Optional<String>> administrativeUnitObjectId() {
        return Codegen.optional(this.administrativeUnitObjectId);
    }
    /**
     * The object ID of the user or group you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="memberObjectId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> memberObjectId;

    /**
     * @return The object ID of the user or group you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
     * 
     */
    public Output<Optional<String>> memberObjectId() {
        return Codegen.optional(this.memberObjectId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AdministrativeUnitMember(String name) {
        this(name, AdministrativeUnitMemberArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AdministrativeUnitMember(String name, @Nullable AdministrativeUnitMemberArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AdministrativeUnitMember(String name, @Nullable AdministrativeUnitMemberArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/administrativeUnitMember:AdministrativeUnitMember", name, args == null ? AdministrativeUnitMemberArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AdministrativeUnitMember(String name, Output<String> id, @Nullable AdministrativeUnitMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/administrativeUnitMember:AdministrativeUnitMember", name, state, makeResourceOptions(options, id));
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
    public static AdministrativeUnitMember get(String name, Output<String> id, @Nullable AdministrativeUnitMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AdministrativeUnitMember(name, id, state, options);
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.DirectoryRoleMemberArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.DirectoryRoleMemberState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a single directory role membership (assignment) within Azure Active Directory.
 * 
 * &gt; **Deprecation Warning:** This resource has been superseded by the azuread.DirectoryRoleAssignment resource and will be removed in version 3.0 of the AzureAD provider
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`
 * 
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.AzureadFunctions;
 * import com.pulumi.azuread.inputs.GetUserArgs;
 * import com.pulumi.azuread.DirectoryRole;
 * import com.pulumi.azuread.DirectoryRoleArgs;
 * import com.pulumi.azuread.DirectoryRoleMember;
 * import com.pulumi.azuread.DirectoryRoleMemberArgs;
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
 *         final var exampleUser = AzureadFunctions.getUser(GetUserArgs.builder()
 *             .userPrincipalName(&#34;jdoe@hashicorp.com&#34;)
 *             .build());
 * 
 *         var exampleDirectoryRole = new DirectoryRole(&#34;exampleDirectoryRole&#34;, DirectoryRoleArgs.builder()        
 *             .displayName(&#34;Security administrator&#34;)
 *             .build());
 * 
 *         var exampleDirectoryRoleMember = new DirectoryRoleMember(&#34;exampleDirectoryRoleMember&#34;, DirectoryRoleMemberArgs.builder()        
 *             .roleObjectId(exampleDirectoryRole.objectId())
 *             .memberObjectId(exampleUser.applyValue(getUserResult -&gt; getUserResult.objectId()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Directory role members can be imported using the object ID of the role and the object ID of the member, e.g.
 * 
 * ```sh
 *  $ pulumi import azuread:index/directoryRoleMember:DirectoryRoleMember example 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
 * ```
 * 
 *  -&gt; This ID format is unique to Terraform and is composed of the Directory Role Object ID and the target Member Object ID in the format `{RoleObjectID}/member/{MemberObjectID}`.
 * 
 */
@ResourceType(type="azuread:index/directoryRoleMember:DirectoryRoleMember")
public class DirectoryRoleMember extends com.pulumi.resources.CustomResource {
    /**
     * The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="memberObjectId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> memberObjectId;

    /**
     * @return The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
     * 
     */
    public Output<Optional<String>> memberObjectId() {
        return Codegen.optional(this.memberObjectId);
    }
    /**
     * The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="roleObjectId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> roleObjectId;

    /**
     * @return The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
     * 
     */
    public Output<Optional<String>> roleObjectId() {
        return Codegen.optional(this.roleObjectId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DirectoryRoleMember(String name) {
        this(name, DirectoryRoleMemberArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DirectoryRoleMember(String name, @Nullable DirectoryRoleMemberArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DirectoryRoleMember(String name, @Nullable DirectoryRoleMemberArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/directoryRoleMember:DirectoryRoleMember", name, args == null ? DirectoryRoleMemberArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DirectoryRoleMember(String name, Output<String> id, @Nullable DirectoryRoleMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/directoryRoleMember:DirectoryRoleMember", name, state, makeResourceOptions(options, id));
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
    public static DirectoryRoleMember get(String name, Output<String> id, @Nullable DirectoryRoleMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DirectoryRoleMember(name, id, state, options);
    }
}

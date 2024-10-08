// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.GroupRoleManagementPolicyArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.GroupRoleManagementPolicyState;
import com.pulumi.azuread.outputs.GroupRoleManagementPolicyActivationRules;
import com.pulumi.azuread.outputs.GroupRoleManagementPolicyActiveAssignmentRules;
import com.pulumi.azuread.outputs.GroupRoleManagementPolicyEligibleAssignmentRules;
import com.pulumi.azuread.outputs.GroupRoleManagementPolicyNotificationRules;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manage a role policy for an Azure AD group.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires the `RoleManagementPolicy.ReadWrite.AzureADGroup` Microsoft Graph API permissions.
 * 
 * When authenticated with a user principal, this resource requires `Global Administrator` directory role, or the `Privileged Role Administrator` role in Identity Governance.
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
 * import com.pulumi.azuread.Group;
 * import com.pulumi.azuread.GroupArgs;
 * import com.pulumi.azuread.User;
 * import com.pulumi.azuread.UserArgs;
 * import com.pulumi.azuread.GroupRoleManagementPolicy;
 * import com.pulumi.azuread.GroupRoleManagementPolicyArgs;
 * import com.pulumi.azuread.inputs.GroupRoleManagementPolicyActiveAssignmentRulesArgs;
 * import com.pulumi.azuread.inputs.GroupRoleManagementPolicyEligibleAssignmentRulesArgs;
 * import com.pulumi.azuread.inputs.GroupRoleManagementPolicyNotificationRulesArgs;
 * import com.pulumi.azuread.inputs.GroupRoleManagementPolicyNotificationRulesEligibleAssignmentsArgs;
 * import com.pulumi.azuread.inputs.GroupRoleManagementPolicyNotificationRulesEligibleAssignmentsApproverNotificationsArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var example = new Group("example", GroupArgs.builder()
 *             .displayName("group-name")
 *             .securityEnabled(true)
 *             .build());
 * 
 *         var member = new User("member", UserArgs.builder()
 *             .userPrincipalName("jdoe}{@literal @}{@code example.com")
 *             .displayName("J. Doe")
 *             .mailNickname("jdoe")
 *             .password("SecretP}{@literal @}{@code sswd99!")
 *             .build());
 * 
 *         var exampleGroupRoleManagementPolicy = new GroupRoleManagementPolicy("exampleGroupRoleManagementPolicy", GroupRoleManagementPolicyArgs.builder()
 *             .groupId(example.id())
 *             .roleId("member")
 *             .activeAssignmentRules(GroupRoleManagementPolicyActiveAssignmentRulesArgs.builder()
 *                 .expireAfter("P365D")
 *                 .build())
 *             .eligibleAssignmentRules(GroupRoleManagementPolicyEligibleAssignmentRulesArgs.builder()
 *                 .expirationRequired(false)
 *                 .build())
 *             .notificationRules(GroupRoleManagementPolicyNotificationRulesArgs.builder()
 *                 .eligibleAssignments(GroupRoleManagementPolicyNotificationRulesEligibleAssignmentsArgs.builder()
 *                     .approverNotifications(GroupRoleManagementPolicyNotificationRulesEligibleAssignmentsApproverNotificationsArgs.builder()
 *                         .notificationLevel("Critical")
 *                         .defaultRecipients(false)
 *                         .additionalRecipients(                        
 *                             "someone}{@literal @}{@code example.com",
 *                             "someone.else}{@literal @}{@code example.com")
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Because these policies are created automatically by Entra ID, they will auto-import on first use.
 * 
 */
@ResourceType(type="azuread:index/groupRoleManagementPolicy:GroupRoleManagementPolicy")
public class GroupRoleManagementPolicy extends com.pulumi.resources.CustomResource {
    /**
     * An `activation_rules` block as defined below.
     * 
     */
    @Export(name="activationRules", refs={GroupRoleManagementPolicyActivationRules.class}, tree="[0]")
    private Output<GroupRoleManagementPolicyActivationRules> activationRules;

    /**
     * @return An `activation_rules` block as defined below.
     * 
     */
    public Output<GroupRoleManagementPolicyActivationRules> activationRules() {
        return this.activationRules;
    }
    /**
     * An `active_assignment_rules` block as defined below.
     * 
     */
    @Export(name="activeAssignmentRules", refs={GroupRoleManagementPolicyActiveAssignmentRules.class}, tree="[0]")
    private Output<GroupRoleManagementPolicyActiveAssignmentRules> activeAssignmentRules;

    /**
     * @return An `active_assignment_rules` block as defined below.
     * 
     */
    public Output<GroupRoleManagementPolicyActiveAssignmentRules> activeAssignmentRules() {
        return this.activeAssignmentRules;
    }
    /**
     * (String) The description of this policy.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return (String) The description of this policy.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * (String) The display name of this policy.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return (String) The display name of this policy.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * An `eligible_assignment_rules` block as defined below.
     * 
     */
    @Export(name="eligibleAssignmentRules", refs={GroupRoleManagementPolicyEligibleAssignmentRules.class}, tree="[0]")
    private Output<GroupRoleManagementPolicyEligibleAssignmentRules> eligibleAssignmentRules;

    /**
     * @return An `eligible_assignment_rules` block as defined below.
     * 
     */
    public Output<GroupRoleManagementPolicyEligibleAssignmentRules> eligibleAssignmentRules() {
        return this.eligibleAssignmentRules;
    }
    /**
     * The ID of the Azure AD group for which the policy applies.
     * 
     */
    @Export(name="groupId", refs={String.class}, tree="[0]")
    private Output<String> groupId;

    /**
     * @return The ID of the Azure AD group for which the policy applies.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }
    /**
     * A `notification_rules` block as defined below.
     * 
     */
    @Export(name="notificationRules", refs={GroupRoleManagementPolicyNotificationRules.class}, tree="[0]")
    private Output<GroupRoleManagementPolicyNotificationRules> notificationRules;

    /**
     * @return A `notification_rules` block as defined below.
     * 
     */
    public Output<GroupRoleManagementPolicyNotificationRules> notificationRules() {
        return this.notificationRules;
    }
    /**
     * The type of assignment this policy coveres. Can be either `member` or `owner`.
     * 
     */
    @Export(name="roleId", refs={String.class}, tree="[0]")
    private Output<String> roleId;

    /**
     * @return The type of assignment this policy coveres. Can be either `member` or `owner`.
     * 
     */
    public Output<String> roleId() {
        return this.roleId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupRoleManagementPolicy(java.lang.String name) {
        this(name, GroupRoleManagementPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupRoleManagementPolicy(java.lang.String name, GroupRoleManagementPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupRoleManagementPolicy(java.lang.String name, GroupRoleManagementPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/groupRoleManagementPolicy:GroupRoleManagementPolicy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private GroupRoleManagementPolicy(java.lang.String name, Output<java.lang.String> id, @Nullable GroupRoleManagementPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/groupRoleManagementPolicy:GroupRoleManagementPolicy", name, state, makeResourceOptions(options, id), false);
    }

    private static GroupRoleManagementPolicyArgs makeArgs(GroupRoleManagementPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? GroupRoleManagementPolicyArgs.Empty : args;
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
    public static GroupRoleManagementPolicy get(java.lang.String name, Output<java.lang.String> id, @Nullable GroupRoleManagementPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupRoleManagementPolicy(name, id, state, options);
    }
}

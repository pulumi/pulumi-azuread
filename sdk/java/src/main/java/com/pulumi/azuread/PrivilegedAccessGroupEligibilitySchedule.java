// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.PrivilegedAccessGroupEligibilityScheduleArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.PrivilegedAccessGroupEligibilityScheduleState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an eligible assignment to a privileged access group.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires the `PrivilegedEligibilitySchedule.ReadWrite.AzureADGroup` Microsoft Graph API permissions.
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
 * import com.pulumi.azuread.PrivilegedAccessGroupEligibilitySchedule;
 * import com.pulumi.azuread.PrivilegedAccessGroupEligibilityScheduleArgs;
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
 *         var example = new Group("example", GroupArgs.builder()        
 *             .displayName("group-name")
 *             .securityEnabled(true)
 *             .build());
 * 
 *         var member = new User("member", UserArgs.builder()        
 *             .userPrincipalName("jdoe{@literal @}example.com")
 *             .displayName("J. Doe")
 *             .mailNickname("jdoe")
 *             .password("SecretP{@literal @}sswd99!")
 *             .build());
 * 
 *         var examplePrivilegedAccessGroupEligibilitySchedule = new PrivilegedAccessGroupEligibilitySchedule("examplePrivilegedAccessGroupEligibilitySchedule", PrivilegedAccessGroupEligibilityScheduleArgs.builder()        
 *             .groupId(pim.id())
 *             .principalId(member.id())
 *             .assignmentType("member")
 *             .duration("P30D")
 *             .justification("as requested")
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
 * An assignment schedule can be imported using the schedule ID, e.g.
 * 
 * ```sh
 * $ pulumi import azuread:index/privilegedAccessGroupEligibilitySchedule:PrivilegedAccessGroupEligibilitySchedule example 00000000-0000-0000-0000-000000000000_member_00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuread:index/privilegedAccessGroupEligibilitySchedule:PrivilegedAccessGroupEligibilitySchedule")
public class PrivilegedAccessGroupEligibilitySchedule extends com.pulumi.resources.CustomResource {
    /**
     * The type of assignment to the group. Can be either `member` or `owner`.
     * 
     */
    @Export(name="assignmentType", refs={String.class}, tree="[0]")
    private Output<String> assignmentType;

    /**
     * @return The type of assignment to the group. Can be either `member` or `owner`.
     * 
     */
    public Output<String> assignmentType() {
        return this.assignmentType;
    }
    /**
     * The duration that this assignment is valid for, formatted as an ISO8601 duration (e.g. P30D for 30 days, PT3H for three hours).
     * 
     */
    @Export(name="duration", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> duration;

    /**
     * @return The duration that this assignment is valid for, formatted as an ISO8601 duration (e.g. P30D for 30 days, PT3H for three hours).
     * 
     */
    public Output<Optional<String>> duration() {
        return Codegen.optional(this.duration);
    }
    /**
     * The date that this assignment expires, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z).
     * 
     */
    @Export(name="expirationDate", refs={String.class}, tree="[0]")
    private Output<String> expirationDate;

    /**
     * @return The date that this assignment expires, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z).
     * 
     */
    public Output<String> expirationDate() {
        return this.expirationDate;
    }
    /**
     * The Object ID of the Azure AD group to which the principal will be assigned.
     * 
     */
    @Export(name="groupId", refs={String.class}, tree="[0]")
    private Output<String> groupId;

    /**
     * @return The Object ID of the Azure AD group to which the principal will be assigned.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }
    /**
     * The justification for this assignment. May be required by the role policy.
     * 
     */
    @Export(name="justification", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> justification;

    /**
     * @return The justification for this assignment. May be required by the role policy.
     * 
     */
    public Output<Optional<String>> justification() {
        return Codegen.optional(this.justification);
    }
    /**
     * Is this assigment permanently valid.
     * 
     * At least one of `expiration_date`, `duration`, or `permanent_assignment` must be supplied. The role policy may limit the maximum duration which can be supplied.
     * 
     */
    @Export(name="permanentAssignment", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> permanentAssignment;

    /**
     * @return Is this assigment permanently valid.
     * 
     * At least one of `expiration_date`, `duration`, or `permanent_assignment` must be supplied. The role policy may limit the maximum duration which can be supplied.
     * 
     */
    public Output<Boolean> permanentAssignment() {
        return this.permanentAssignment;
    }
    /**
     * The Object ID of the principal to be assigned to the above group. Can be either a user or a group.
     * 
     */
    @Export(name="principalId", refs={String.class}, tree="[0]")
    private Output<String> principalId;

    /**
     * @return The Object ID of the principal to be assigned to the above group. Can be either a user or a group.
     * 
     */
    public Output<String> principalId() {
        return this.principalId;
    }
    /**
     * The date from which this assignment is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). If not provided, the assignment is immediately valid.
     * 
     */
    @Export(name="startDate", refs={String.class}, tree="[0]")
    private Output<String> startDate;

    /**
     * @return The date from which this assignment is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). If not provided, the assignment is immediately valid.
     * 
     */
    public Output<String> startDate() {
        return this.startDate;
    }
    /**
     * (String) The provisioning status of this request.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return (String) The provisioning status of this request.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The ticket number in the ticket system approving this assignment. May be required by the role policy.
     * 
     */
    @Export(name="ticketNumber", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ticketNumber;

    /**
     * @return The ticket number in the ticket system approving this assignment. May be required by the role policy.
     * 
     */
    public Output<Optional<String>> ticketNumber() {
        return Codegen.optional(this.ticketNumber);
    }
    /**
     * The ticket system containing the ticket number approving this assignment. May be required by the role policy.
     * 
     */
    @Export(name="ticketSystem", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ticketSystem;

    /**
     * @return The ticket system containing the ticket number approving this assignment. May be required by the role policy.
     * 
     */
    public Output<Optional<String>> ticketSystem() {
        return Codegen.optional(this.ticketSystem);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PrivilegedAccessGroupEligibilitySchedule(String name) {
        this(name, PrivilegedAccessGroupEligibilityScheduleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PrivilegedAccessGroupEligibilitySchedule(String name, PrivilegedAccessGroupEligibilityScheduleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PrivilegedAccessGroupEligibilitySchedule(String name, PrivilegedAccessGroupEligibilityScheduleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/privilegedAccessGroupEligibilitySchedule:PrivilegedAccessGroupEligibilitySchedule", name, args == null ? PrivilegedAccessGroupEligibilityScheduleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PrivilegedAccessGroupEligibilitySchedule(String name, Output<String> id, @Nullable PrivilegedAccessGroupEligibilityScheduleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/privilegedAccessGroupEligibilitySchedule:PrivilegedAccessGroupEligibilitySchedule", name, state, makeResourceOptions(options, id));
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
    public static PrivilegedAccessGroupEligibilitySchedule get(String name, Output<String> id, @Nullable PrivilegedAccessGroupEligibilityScheduleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PrivilegedAccessGroupEligibilitySchedule(name, id, state, options);
    }
}
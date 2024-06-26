// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetGroupRoleManagementPolicyArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGroupRoleManagementPolicyArgs Empty = new GetGroupRoleManagementPolicyArgs();

    /**
     * The ID of the Azure AD group for which the policy applies.
     * 
     */
    @Import(name="groupId", required=true)
    private Output<String> groupId;

    /**
     * @return The ID of the Azure AD group for which the policy applies.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }

    /**
     * The type of assignment this policy coveres. Can be either `member` or `owner`.
     * 
     */
    @Import(name="roleId", required=true)
    private Output<String> roleId;

    /**
     * @return The type of assignment this policy coveres. Can be either `member` or `owner`.
     * 
     */
    public Output<String> roleId() {
        return this.roleId;
    }

    private GetGroupRoleManagementPolicyArgs() {}

    private GetGroupRoleManagementPolicyArgs(GetGroupRoleManagementPolicyArgs $) {
        this.groupId = $.groupId;
        this.roleId = $.roleId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGroupRoleManagementPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGroupRoleManagementPolicyArgs $;

        public Builder() {
            $ = new GetGroupRoleManagementPolicyArgs();
        }

        public Builder(GetGroupRoleManagementPolicyArgs defaults) {
            $ = new GetGroupRoleManagementPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param groupId The ID of the Azure AD group for which the policy applies.
         * 
         * @return builder
         * 
         */
        public Builder groupId(Output<String> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId The ID of the Azure AD group for which the policy applies.
         * 
         * @return builder
         * 
         */
        public Builder groupId(String groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param roleId The type of assignment this policy coveres. Can be either `member` or `owner`.
         * 
         * @return builder
         * 
         */
        public Builder roleId(Output<String> roleId) {
            $.roleId = roleId;
            return this;
        }

        /**
         * @param roleId The type of assignment this policy coveres. Can be either `member` or `owner`.
         * 
         * @return builder
         * 
         */
        public Builder roleId(String roleId) {
            return roleId(Output.of(roleId));
        }

        public GetGroupRoleManagementPolicyArgs build() {
            if ($.groupId == null) {
                throw new MissingRequiredPropertyException("GetGroupRoleManagementPolicyArgs", "groupId");
            }
            if ($.roleId == null) {
                throw new MissingRequiredPropertyException("GetGroupRoleManagementPolicyArgs", "roleId");
            }
            return $;
        }
    }

}

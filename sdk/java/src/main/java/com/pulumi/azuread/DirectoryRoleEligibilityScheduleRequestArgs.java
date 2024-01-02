// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class DirectoryRoleEligibilityScheduleRequestArgs extends com.pulumi.resources.ResourceArgs {

    public static final DirectoryRoleEligibilityScheduleRequestArgs Empty = new DirectoryRoleEligibilityScheduleRequestArgs();

    /**
     * Identifier of the directory object representing the scope of the role eligibility. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="directoryScopeId", required=true)
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
    @Import(name="justification", required=true)
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
    @Import(name="principalId", required=true)
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
    @Import(name="roleDefinitionId", required=true)
    private Output<String> roleDefinitionId;

    /**
     * @return The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> roleDefinitionId() {
        return this.roleDefinitionId;
    }

    private DirectoryRoleEligibilityScheduleRequestArgs() {}

    private DirectoryRoleEligibilityScheduleRequestArgs(DirectoryRoleEligibilityScheduleRequestArgs $) {
        this.directoryScopeId = $.directoryScopeId;
        this.justification = $.justification;
        this.principalId = $.principalId;
        this.roleDefinitionId = $.roleDefinitionId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DirectoryRoleEligibilityScheduleRequestArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DirectoryRoleEligibilityScheduleRequestArgs $;

        public Builder() {
            $ = new DirectoryRoleEligibilityScheduleRequestArgs();
        }

        public Builder(DirectoryRoleEligibilityScheduleRequestArgs defaults) {
            $ = new DirectoryRoleEligibilityScheduleRequestArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param directoryScopeId Identifier of the directory object representing the scope of the role eligibility. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder directoryScopeId(Output<String> directoryScopeId) {
            $.directoryScopeId = directoryScopeId;
            return this;
        }

        /**
         * @param directoryScopeId Identifier of the directory object representing the scope of the role eligibility. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder directoryScopeId(String directoryScopeId) {
            return directoryScopeId(Output.of(directoryScopeId));
        }

        /**
         * @param justification Justification for why the principal is granted the role eligibility. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder justification(Output<String> justification) {
            $.justification = justification;
            return this;
        }

        /**
         * @param justification Justification for why the principal is granted the role eligibility. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder justification(String justification) {
            return justification(Output.of(justification));
        }

        /**
         * @param principalId The object ID of the principal to granted the role eligibility. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder principalId(Output<String> principalId) {
            $.principalId = principalId;
            return this;
        }

        /**
         * @param principalId The object ID of the principal to granted the role eligibility. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder principalId(String principalId) {
            return principalId(Output.of(principalId));
        }

        /**
         * @param roleDefinitionId The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder roleDefinitionId(Output<String> roleDefinitionId) {
            $.roleDefinitionId = roleDefinitionId;
            return this;
        }

        /**
         * @param roleDefinitionId The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder roleDefinitionId(String roleDefinitionId) {
            return roleDefinitionId(Output.of(roleDefinitionId));
        }

        public DirectoryRoleEligibilityScheduleRequestArgs build() {
            if ($.directoryScopeId == null) {
                throw new MissingRequiredPropertyException("DirectoryRoleEligibilityScheduleRequestArgs", "directoryScopeId");
            }
            if ($.justification == null) {
                throw new MissingRequiredPropertyException("DirectoryRoleEligibilityScheduleRequestArgs", "justification");
            }
            if ($.principalId == null) {
                throw new MissingRequiredPropertyException("DirectoryRoleEligibilityScheduleRequestArgs", "principalId");
            }
            if ($.roleDefinitionId == null) {
                throw new MissingRequiredPropertyException("DirectoryRoleEligibilityScheduleRequestArgs", "roleDefinitionId");
            }
            return $;
        }
    }

}

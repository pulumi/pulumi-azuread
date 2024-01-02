// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class ServicePrincipalClaimsMappingPolicyAssignmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServicePrincipalClaimsMappingPolicyAssignmentArgs Empty = new ServicePrincipalClaimsMappingPolicyAssignmentArgs();

    /**
     * The ID of the claims mapping policy to assign.
     * 
     */
    @Import(name="claimsMappingPolicyId", required=true)
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
    @Import(name="servicePrincipalId", required=true)
    private Output<String> servicePrincipalId;

    /**
     * @return The object ID of the service principal for the policy assignment.
     * 
     */
    public Output<String> servicePrincipalId() {
        return this.servicePrincipalId;
    }

    private ServicePrincipalClaimsMappingPolicyAssignmentArgs() {}

    private ServicePrincipalClaimsMappingPolicyAssignmentArgs(ServicePrincipalClaimsMappingPolicyAssignmentArgs $) {
        this.claimsMappingPolicyId = $.claimsMappingPolicyId;
        this.servicePrincipalId = $.servicePrincipalId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServicePrincipalClaimsMappingPolicyAssignmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServicePrincipalClaimsMappingPolicyAssignmentArgs $;

        public Builder() {
            $ = new ServicePrincipalClaimsMappingPolicyAssignmentArgs();
        }

        public Builder(ServicePrincipalClaimsMappingPolicyAssignmentArgs defaults) {
            $ = new ServicePrincipalClaimsMappingPolicyAssignmentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param claimsMappingPolicyId The ID of the claims mapping policy to assign.
         * 
         * @return builder
         * 
         */
        public Builder claimsMappingPolicyId(Output<String> claimsMappingPolicyId) {
            $.claimsMappingPolicyId = claimsMappingPolicyId;
            return this;
        }

        /**
         * @param claimsMappingPolicyId The ID of the claims mapping policy to assign.
         * 
         * @return builder
         * 
         */
        public Builder claimsMappingPolicyId(String claimsMappingPolicyId) {
            return claimsMappingPolicyId(Output.of(claimsMappingPolicyId));
        }

        /**
         * @param servicePrincipalId The object ID of the service principal for the policy assignment.
         * 
         * @return builder
         * 
         */
        public Builder servicePrincipalId(Output<String> servicePrincipalId) {
            $.servicePrincipalId = servicePrincipalId;
            return this;
        }

        /**
         * @param servicePrincipalId The object ID of the service principal for the policy assignment.
         * 
         * @return builder
         * 
         */
        public Builder servicePrincipalId(String servicePrincipalId) {
            return servicePrincipalId(Output.of(servicePrincipalId));
        }

        public ServicePrincipalClaimsMappingPolicyAssignmentArgs build() {
            if ($.claimsMappingPolicyId == null) {
                throw new MissingRequiredPropertyException("ServicePrincipalClaimsMappingPolicyAssignmentArgs", "claimsMappingPolicyId");
            }
            if ($.servicePrincipalId == null) {
                throw new MissingRequiredPropertyException("ServicePrincipalClaimsMappingPolicyAssignmentArgs", "servicePrincipalId");
            }
            return $;
        }
    }

}

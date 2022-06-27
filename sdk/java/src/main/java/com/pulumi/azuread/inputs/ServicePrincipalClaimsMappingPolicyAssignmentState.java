// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServicePrincipalClaimsMappingPolicyAssignmentState extends com.pulumi.resources.ResourceArgs {

    public static final ServicePrincipalClaimsMappingPolicyAssignmentState Empty = new ServicePrincipalClaimsMappingPolicyAssignmentState();

    /**
     * The ID of the claims mapping policy to assign.
     * 
     */
    @Import(name="claimsMappingPolicyId")
    private @Nullable Output<String> claimsMappingPolicyId;

    /**
     * @return The ID of the claims mapping policy to assign.
     * 
     */
    public Optional<Output<String>> claimsMappingPolicyId() {
        return Optional.ofNullable(this.claimsMappingPolicyId);
    }

    /**
     * The object ID of the service principal for the policy assignment.
     * 
     */
    @Import(name="servicePrincipalId")
    private @Nullable Output<String> servicePrincipalId;

    /**
     * @return The object ID of the service principal for the policy assignment.
     * 
     */
    public Optional<Output<String>> servicePrincipalId() {
        return Optional.ofNullable(this.servicePrincipalId);
    }

    private ServicePrincipalClaimsMappingPolicyAssignmentState() {}

    private ServicePrincipalClaimsMappingPolicyAssignmentState(ServicePrincipalClaimsMappingPolicyAssignmentState $) {
        this.claimsMappingPolicyId = $.claimsMappingPolicyId;
        this.servicePrincipalId = $.servicePrincipalId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServicePrincipalClaimsMappingPolicyAssignmentState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServicePrincipalClaimsMappingPolicyAssignmentState $;

        public Builder() {
            $ = new ServicePrincipalClaimsMappingPolicyAssignmentState();
        }

        public Builder(ServicePrincipalClaimsMappingPolicyAssignmentState defaults) {
            $ = new ServicePrincipalClaimsMappingPolicyAssignmentState(Objects.requireNonNull(defaults));
        }

        /**
         * @param claimsMappingPolicyId The ID of the claims mapping policy to assign.
         * 
         * @return builder
         * 
         */
        public Builder claimsMappingPolicyId(@Nullable Output<String> claimsMappingPolicyId) {
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
        public Builder servicePrincipalId(@Nullable Output<String> servicePrincipalId) {
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

        public ServicePrincipalClaimsMappingPolicyAssignmentState build() {
            return $;
        }
    }

}

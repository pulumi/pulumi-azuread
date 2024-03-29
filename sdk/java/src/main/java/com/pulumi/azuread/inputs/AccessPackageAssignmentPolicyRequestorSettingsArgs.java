// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.azuread.inputs.AccessPackageAssignmentPolicyRequestorSettingsRequestorArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AccessPackageAssignmentPolicyRequestorSettingsArgs extends com.pulumi.resources.ResourceArgs {

    public static final AccessPackageAssignmentPolicyRequestorSettingsArgs Empty = new AccessPackageAssignmentPolicyRequestorSettingsArgs();

    /**
     * A block specifying the users who are allowed to request on this policy, as documented below.
     * 
     */
    @Import(name="requestors")
    private @Nullable Output<List<AccessPackageAssignmentPolicyRequestorSettingsRequestorArgs>> requestors;

    /**
     * @return A block specifying the users who are allowed to request on this policy, as documented below.
     * 
     */
    public Optional<Output<List<AccessPackageAssignmentPolicyRequestorSettingsRequestorArgs>>> requestors() {
        return Optional.ofNullable(this.requestors);
    }

    /**
     * Whether to accept requests using this policy. When `false`, no new requests can be made using this policy.
     * 
     */
    @Import(name="requestsAccepted")
    private @Nullable Output<Boolean> requestsAccepted;

    /**
     * @return Whether to accept requests using this policy. When `false`, no new requests can be made using this policy.
     * 
     */
    public Optional<Output<Boolean>> requestsAccepted() {
        return Optional.ofNullable(this.requestsAccepted);
    }

    /**
     * Specifies the scopes of the requestors. Valid values are `AllConfiguredConnectedOrganizationSubjects`, `AllExistingConnectedOrganizationSubjects`, `AllExistingDirectoryMemberUsers`, `AllExistingDirectorySubjects`, `AllExternalSubjects`, `NoSubjects`, `SpecificConnectedOrganizationSubjects`, or `SpecificDirectorySubjects`.
     * 
     */
    @Import(name="scopeType")
    private @Nullable Output<String> scopeType;

    /**
     * @return Specifies the scopes of the requestors. Valid values are `AllConfiguredConnectedOrganizationSubjects`, `AllExistingConnectedOrganizationSubjects`, `AllExistingDirectoryMemberUsers`, `AllExistingDirectorySubjects`, `AllExternalSubjects`, `NoSubjects`, `SpecificConnectedOrganizationSubjects`, or `SpecificDirectorySubjects`.
     * 
     */
    public Optional<Output<String>> scopeType() {
        return Optional.ofNullable(this.scopeType);
    }

    private AccessPackageAssignmentPolicyRequestorSettingsArgs() {}

    private AccessPackageAssignmentPolicyRequestorSettingsArgs(AccessPackageAssignmentPolicyRequestorSettingsArgs $) {
        this.requestors = $.requestors;
        this.requestsAccepted = $.requestsAccepted;
        this.scopeType = $.scopeType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccessPackageAssignmentPolicyRequestorSettingsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccessPackageAssignmentPolicyRequestorSettingsArgs $;

        public Builder() {
            $ = new AccessPackageAssignmentPolicyRequestorSettingsArgs();
        }

        public Builder(AccessPackageAssignmentPolicyRequestorSettingsArgs defaults) {
            $ = new AccessPackageAssignmentPolicyRequestorSettingsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param requestors A block specifying the users who are allowed to request on this policy, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder requestors(@Nullable Output<List<AccessPackageAssignmentPolicyRequestorSettingsRequestorArgs>> requestors) {
            $.requestors = requestors;
            return this;
        }

        /**
         * @param requestors A block specifying the users who are allowed to request on this policy, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder requestors(List<AccessPackageAssignmentPolicyRequestorSettingsRequestorArgs> requestors) {
            return requestors(Output.of(requestors));
        }

        /**
         * @param requestors A block specifying the users who are allowed to request on this policy, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder requestors(AccessPackageAssignmentPolicyRequestorSettingsRequestorArgs... requestors) {
            return requestors(List.of(requestors));
        }

        /**
         * @param requestsAccepted Whether to accept requests using this policy. When `false`, no new requests can be made using this policy.
         * 
         * @return builder
         * 
         */
        public Builder requestsAccepted(@Nullable Output<Boolean> requestsAccepted) {
            $.requestsAccepted = requestsAccepted;
            return this;
        }

        /**
         * @param requestsAccepted Whether to accept requests using this policy. When `false`, no new requests can be made using this policy.
         * 
         * @return builder
         * 
         */
        public Builder requestsAccepted(Boolean requestsAccepted) {
            return requestsAccepted(Output.of(requestsAccepted));
        }

        /**
         * @param scopeType Specifies the scopes of the requestors. Valid values are `AllConfiguredConnectedOrganizationSubjects`, `AllExistingConnectedOrganizationSubjects`, `AllExistingDirectoryMemberUsers`, `AllExistingDirectorySubjects`, `AllExternalSubjects`, `NoSubjects`, `SpecificConnectedOrganizationSubjects`, or `SpecificDirectorySubjects`.
         * 
         * @return builder
         * 
         */
        public Builder scopeType(@Nullable Output<String> scopeType) {
            $.scopeType = scopeType;
            return this;
        }

        /**
         * @param scopeType Specifies the scopes of the requestors. Valid values are `AllConfiguredConnectedOrganizationSubjects`, `AllExistingConnectedOrganizationSubjects`, `AllExistingDirectoryMemberUsers`, `AllExistingDirectorySubjects`, `AllExternalSubjects`, `NoSubjects`, `SpecificConnectedOrganizationSubjects`, or `SpecificDirectorySubjects`.
         * 
         * @return builder
         * 
         */
        public Builder scopeType(String scopeType) {
            return scopeType(Output.of(scopeType));
        }

        public AccessPackageAssignmentPolicyRequestorSettingsArgs build() {
            return $;
        }
    }

}

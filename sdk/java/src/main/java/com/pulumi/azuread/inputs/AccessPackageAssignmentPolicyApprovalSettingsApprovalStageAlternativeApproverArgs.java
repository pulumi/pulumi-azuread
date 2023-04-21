// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverArgs extends com.pulumi.resources.ResourceArgs {

    public static final AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverArgs Empty = new AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverArgs();

    /**
     * For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
     * 
     */
    @Import(name="backup")
    private @Nullable Output<Boolean> backup;

    /**
     * @return For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
     * 
     */
    public Optional<Output<Boolean>> backup() {
        return Optional.ofNullable(this.backup);
    }

    /**
     * The ID of the subject.
     * 
     */
    @Import(name="objectId")
    private @Nullable Output<String> objectId;

    /**
     * @return The ID of the subject.
     * 
     */
    public Optional<Output<String>> objectId() {
        return Optional.ofNullable(this.objectId);
    }

    /**
     * Specifies the type of users. Valid values are `singleUser`, `groupMembers`, `connectedOrganizationMembers`, `requestorManager`, `internalSponsors`, or `externalSponsors`.
     * 
     */
    @Import(name="subjectType", required=true)
    private Output<String> subjectType;

    /**
     * @return Specifies the type of users. Valid values are `singleUser`, `groupMembers`, `connectedOrganizationMembers`, `requestorManager`, `internalSponsors`, or `externalSponsors`.
     * 
     */
    public Output<String> subjectType() {
        return this.subjectType;
    }

    private AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverArgs() {}

    private AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverArgs(AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverArgs $) {
        this.backup = $.backup;
        this.objectId = $.objectId;
        this.subjectType = $.subjectType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverArgs $;

        public Builder() {
            $ = new AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverArgs();
        }

        public Builder(AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverArgs defaults) {
            $ = new AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backup For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
         * 
         * @return builder
         * 
         */
        public Builder backup(@Nullable Output<Boolean> backup) {
            $.backup = backup;
            return this;
        }

        /**
         * @param backup For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
         * 
         * @return builder
         * 
         */
        public Builder backup(Boolean backup) {
            return backup(Output.of(backup));
        }

        /**
         * @param objectId The ID of the subject.
         * 
         * @return builder
         * 
         */
        public Builder objectId(@Nullable Output<String> objectId) {
            $.objectId = objectId;
            return this;
        }

        /**
         * @param objectId The ID of the subject.
         * 
         * @return builder
         * 
         */
        public Builder objectId(String objectId) {
            return objectId(Output.of(objectId));
        }

        /**
         * @param subjectType Specifies the type of users. Valid values are `singleUser`, `groupMembers`, `connectedOrganizationMembers`, `requestorManager`, `internalSponsors`, or `externalSponsors`.
         * 
         * @return builder
         * 
         */
        public Builder subjectType(Output<String> subjectType) {
            $.subjectType = subjectType;
            return this;
        }

        /**
         * @param subjectType Specifies the type of users. Valid values are `singleUser`, `groupMembers`, `connectedOrganizationMembers`, `requestorManager`, `internalSponsors`, or `externalSponsors`.
         * 
         * @return builder
         * 
         */
        public Builder subjectType(String subjectType) {
            return subjectType(Output.of(subjectType));
        }

        public AccessPackageAssignmentPolicyApprovalSettingsApprovalStageAlternativeApproverArgs build() {
            $.subjectType = Objects.requireNonNull($.subjectType, "expected parameter 'subjectType' to be non-null");
            return $;
        }
    }

}
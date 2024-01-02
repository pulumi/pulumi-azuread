// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewer {
    /**
     * @return For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
     * 
     */
    private @Nullable Boolean backup;
    /**
     * @return The ID of the subject.
     * 
     */
    private @Nullable String objectId;
    /**
     * @return Specifies the type of users. Valid values are `singleUser`, `groupMembers`, `connectedOrganizationMembers`, `requestorManager`, `internalSponsors`, or `externalSponsors`.
     * 
     */
    private String subjectType;

    private AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewer() {}
    /**
     * @return For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
     * 
     */
    public Optional<Boolean> backup() {
        return Optional.ofNullable(this.backup);
    }
    /**
     * @return The ID of the subject.
     * 
     */
    public Optional<String> objectId() {
        return Optional.ofNullable(this.objectId);
    }
    /**
     * @return Specifies the type of users. Valid values are `singleUser`, `groupMembers`, `connectedOrganizationMembers`, `requestorManager`, `internalSponsors`, or `externalSponsors`.
     * 
     */
    public String subjectType() {
        return this.subjectType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewer defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean backup;
        private @Nullable String objectId;
        private String subjectType;
        public Builder() {}
        public Builder(AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewer defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.backup = defaults.backup;
    	      this.objectId = defaults.objectId;
    	      this.subjectType = defaults.subjectType;
        }

        @CustomType.Setter
        public Builder backup(@Nullable Boolean backup) {

            this.backup = backup;
            return this;
        }
        @CustomType.Setter
        public Builder objectId(@Nullable String objectId) {

            this.objectId = objectId;
            return this;
        }
        @CustomType.Setter
        public Builder subjectType(String subjectType) {
            if (subjectType == null) {
              throw new MissingRequiredPropertyException("AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewer", "subjectType");
            }
            this.subjectType = subjectType;
            return this;
        }
        public AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewer build() {
            final var _resultValue = new AccessPackageAssignmentPolicyAssignmentReviewSettingsReviewer();
            _resultValue.backup = backup;
            _resultValue.objectId = objectId;
            _resultValue.subjectType = subjectType;
            return _resultValue;
        }
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetGroupsResult {
    private String displayNamePrefix;
    /**
     * @return The display names of the groups.
     * 
     */
    private List<String> displayNames;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable Boolean ignoreMissing;
    private Boolean mailEnabled;
    /**
     * @return The object IDs of the groups.
     * 
     */
    private List<String> objectIds;
    private @Nullable Boolean returnAll;
    private Boolean securityEnabled;

    private GetGroupsResult() {}
    public String displayNamePrefix() {
        return this.displayNamePrefix;
    }
    /**
     * @return The display names of the groups.
     * 
     */
    public List<String> displayNames() {
        return this.displayNames;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<Boolean> ignoreMissing() {
        return Optional.ofNullable(this.ignoreMissing);
    }
    public Boolean mailEnabled() {
        return this.mailEnabled;
    }
    /**
     * @return The object IDs of the groups.
     * 
     */
    public List<String> objectIds() {
        return this.objectIds;
    }
    public Optional<Boolean> returnAll() {
        return Optional.ofNullable(this.returnAll);
    }
    public Boolean securityEnabled() {
        return this.securityEnabled;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGroupsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String displayNamePrefix;
        private List<String> displayNames;
        private String id;
        private @Nullable Boolean ignoreMissing;
        private Boolean mailEnabled;
        private List<String> objectIds;
        private @Nullable Boolean returnAll;
        private Boolean securityEnabled;
        public Builder() {}
        public Builder(GetGroupsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.displayNamePrefix = defaults.displayNamePrefix;
    	      this.displayNames = defaults.displayNames;
    	      this.id = defaults.id;
    	      this.ignoreMissing = defaults.ignoreMissing;
    	      this.mailEnabled = defaults.mailEnabled;
    	      this.objectIds = defaults.objectIds;
    	      this.returnAll = defaults.returnAll;
    	      this.securityEnabled = defaults.securityEnabled;
        }

        @CustomType.Setter
        public Builder displayNamePrefix(String displayNamePrefix) {
            if (displayNamePrefix == null) {
              throw new MissingRequiredPropertyException("GetGroupsResult", "displayNamePrefix");
            }
            this.displayNamePrefix = displayNamePrefix;
            return this;
        }
        @CustomType.Setter
        public Builder displayNames(List<String> displayNames) {
            if (displayNames == null) {
              throw new MissingRequiredPropertyException("GetGroupsResult", "displayNames");
            }
            this.displayNames = displayNames;
            return this;
        }
        public Builder displayNames(String... displayNames) {
            return displayNames(List.of(displayNames));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetGroupsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ignoreMissing(@Nullable Boolean ignoreMissing) {

            this.ignoreMissing = ignoreMissing;
            return this;
        }
        @CustomType.Setter
        public Builder mailEnabled(Boolean mailEnabled) {
            if (mailEnabled == null) {
              throw new MissingRequiredPropertyException("GetGroupsResult", "mailEnabled");
            }
            this.mailEnabled = mailEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder objectIds(List<String> objectIds) {
            if (objectIds == null) {
              throw new MissingRequiredPropertyException("GetGroupsResult", "objectIds");
            }
            this.objectIds = objectIds;
            return this;
        }
        public Builder objectIds(String... objectIds) {
            return objectIds(List.of(objectIds));
        }
        @CustomType.Setter
        public Builder returnAll(@Nullable Boolean returnAll) {

            this.returnAll = returnAll;
            return this;
        }
        @CustomType.Setter
        public Builder securityEnabled(Boolean securityEnabled) {
            if (securityEnabled == null) {
              throw new MissingRequiredPropertyException("GetGroupsResult", "securityEnabled");
            }
            this.securityEnabled = securityEnabled;
            return this;
        }
        public GetGroupsResult build() {
            final var _resultValue = new GetGroupsResult();
            _resultValue.displayNamePrefix = displayNamePrefix;
            _resultValue.displayNames = displayNames;
            _resultValue.id = id;
            _resultValue.ignoreMissing = ignoreMissing;
            _resultValue.mailEnabled = mailEnabled;
            _resultValue.objectIds = objectIds;
            _resultValue.returnAll = returnAll;
            _resultValue.securityEnabled = securityEnabled;
            return _resultValue;
        }
    }
}

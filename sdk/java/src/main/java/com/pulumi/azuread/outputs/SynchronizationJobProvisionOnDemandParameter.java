// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.azuread.outputs.SynchronizationJobProvisionOnDemandParameterSubject;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class SynchronizationJobProvisionOnDemandParameter {
    /**
     * @return The identifier of the synchronization rule to be applied. This rule ID is defined in the schema for a given synchronization job or template.
     * 
     */
    private String ruleId;
    /**
     * @return One or more `subject` blocks as documented below.
     * 
     */
    private List<SynchronizationJobProvisionOnDemandParameterSubject> subjects;

    private SynchronizationJobProvisionOnDemandParameter() {}
    /**
     * @return The identifier of the synchronization rule to be applied. This rule ID is defined in the schema for a given synchronization job or template.
     * 
     */
    public String ruleId() {
        return this.ruleId;
    }
    /**
     * @return One or more `subject` blocks as documented below.
     * 
     */
    public List<SynchronizationJobProvisionOnDemandParameterSubject> subjects() {
        return this.subjects;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SynchronizationJobProvisionOnDemandParameter defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String ruleId;
        private List<SynchronizationJobProvisionOnDemandParameterSubject> subjects;
        public Builder() {}
        public Builder(SynchronizationJobProvisionOnDemandParameter defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ruleId = defaults.ruleId;
    	      this.subjects = defaults.subjects;
        }

        @CustomType.Setter
        public Builder ruleId(String ruleId) {
            if (ruleId == null) {
              throw new MissingRequiredPropertyException("SynchronizationJobProvisionOnDemandParameter", "ruleId");
            }
            this.ruleId = ruleId;
            return this;
        }
        @CustomType.Setter
        public Builder subjects(List<SynchronizationJobProvisionOnDemandParameterSubject> subjects) {
            if (subjects == null) {
              throw new MissingRequiredPropertyException("SynchronizationJobProvisionOnDemandParameter", "subjects");
            }
            this.subjects = subjects;
            return this;
        }
        public Builder subjects(SynchronizationJobProvisionOnDemandParameterSubject... subjects) {
            return subjects(List.of(subjects));
        }
        public SynchronizationJobProvisionOnDemandParameter build() {
            final var _resultValue = new SynchronizationJobProvisionOnDemandParameter();
            _resultValue.ruleId = ruleId;
            _resultValue.subjects = subjects;
            return _resultValue;
        }
    }
}

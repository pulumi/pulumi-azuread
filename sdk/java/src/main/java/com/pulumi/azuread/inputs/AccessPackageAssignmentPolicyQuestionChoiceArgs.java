// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.azuread.inputs.AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class AccessPackageAssignmentPolicyQuestionChoiceArgs extends com.pulumi.resources.ResourceArgs {

    public static final AccessPackageAssignmentPolicyQuestionChoiceArgs Empty = new AccessPackageAssignmentPolicyQuestionChoiceArgs();

    /**
     * The actual value of this choice.
     * 
     */
    @Import(name="actualValue", required=true)
    private Output<String> actualValue;

    /**
     * @return The actual value of this choice.
     * 
     */
    public Output<String> actualValue() {
        return this.actualValue;
    }

    /**
     * A block describing the display text of this choice, as documented below.
     * 
     */
    @Import(name="displayValue", required=true)
    private Output<AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs> displayValue;

    /**
     * @return A block describing the display text of this choice, as documented below.
     * 
     */
    public Output<AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs> displayValue() {
        return this.displayValue;
    }

    private AccessPackageAssignmentPolicyQuestionChoiceArgs() {}

    private AccessPackageAssignmentPolicyQuestionChoiceArgs(AccessPackageAssignmentPolicyQuestionChoiceArgs $) {
        this.actualValue = $.actualValue;
        this.displayValue = $.displayValue;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccessPackageAssignmentPolicyQuestionChoiceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccessPackageAssignmentPolicyQuestionChoiceArgs $;

        public Builder() {
            $ = new AccessPackageAssignmentPolicyQuestionChoiceArgs();
        }

        public Builder(AccessPackageAssignmentPolicyQuestionChoiceArgs defaults) {
            $ = new AccessPackageAssignmentPolicyQuestionChoiceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param actualValue The actual value of this choice.
         * 
         * @return builder
         * 
         */
        public Builder actualValue(Output<String> actualValue) {
            $.actualValue = actualValue;
            return this;
        }

        /**
         * @param actualValue The actual value of this choice.
         * 
         * @return builder
         * 
         */
        public Builder actualValue(String actualValue) {
            return actualValue(Output.of(actualValue));
        }

        /**
         * @param displayValue A block describing the display text of this choice, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder displayValue(Output<AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs> displayValue) {
            $.displayValue = displayValue;
            return this;
        }

        /**
         * @param displayValue A block describing the display text of this choice, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder displayValue(AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs displayValue) {
            return displayValue(Output.of(displayValue));
        }

        public AccessPackageAssignmentPolicyQuestionChoiceArgs build() {
            $.actualValue = Objects.requireNonNull($.actualValue, "expected parameter 'actualValue' to be non-null");
            $.displayValue = Objects.requireNonNull($.displayValue, "expected parameter 'displayValue' to be non-null");
            return $;
        }
    }

}
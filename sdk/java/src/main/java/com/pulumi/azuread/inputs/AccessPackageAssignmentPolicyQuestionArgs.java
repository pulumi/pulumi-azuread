// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.azuread.inputs.AccessPackageAssignmentPolicyQuestionChoiceArgs;
import com.pulumi.azuread.inputs.AccessPackageAssignmentPolicyQuestionTextArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AccessPackageAssignmentPolicyQuestionArgs extends com.pulumi.resources.ResourceArgs {

    public static final AccessPackageAssignmentPolicyQuestionArgs Empty = new AccessPackageAssignmentPolicyQuestionArgs();

    /**
     * One or more blocks configuring a choice to the question, as documented below.
     * 
     */
    @Import(name="choices")
    private @Nullable Output<List<AccessPackageAssignmentPolicyQuestionChoiceArgs>> choices;

    /**
     * @return One or more blocks configuring a choice to the question, as documented below.
     * 
     */
    public Optional<Output<List<AccessPackageAssignmentPolicyQuestionChoiceArgs>>> choices() {
        return Optional.ofNullable(this.choices);
    }

    /**
     * Whether this question is required.
     * 
     */
    @Import(name="required")
    private @Nullable Output<Boolean> required;

    /**
     * @return Whether this question is required.
     * 
     */
    public Optional<Output<Boolean>> required() {
        return Optional.ofNullable(this.required);
    }

    /**
     * The sequence number of this question.
     * 
     */
    @Import(name="sequence")
    private @Nullable Output<Integer> sequence;

    /**
     * @return The sequence number of this question.
     * 
     */
    public Optional<Output<Integer>> sequence() {
        return Optional.ofNullable(this.sequence);
    }

    /**
     * A block describing the content of this question, as documented below.
     * 
     */
    @Import(name="text", required=true)
    private Output<AccessPackageAssignmentPolicyQuestionTextArgs> text;

    /**
     * @return A block describing the content of this question, as documented below.
     * 
     */
    public Output<AccessPackageAssignmentPolicyQuestionTextArgs> text() {
        return this.text;
    }

    private AccessPackageAssignmentPolicyQuestionArgs() {}

    private AccessPackageAssignmentPolicyQuestionArgs(AccessPackageAssignmentPolicyQuestionArgs $) {
        this.choices = $.choices;
        this.required = $.required;
        this.sequence = $.sequence;
        this.text = $.text;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccessPackageAssignmentPolicyQuestionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccessPackageAssignmentPolicyQuestionArgs $;

        public Builder() {
            $ = new AccessPackageAssignmentPolicyQuestionArgs();
        }

        public Builder(AccessPackageAssignmentPolicyQuestionArgs defaults) {
            $ = new AccessPackageAssignmentPolicyQuestionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param choices One or more blocks configuring a choice to the question, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder choices(@Nullable Output<List<AccessPackageAssignmentPolicyQuestionChoiceArgs>> choices) {
            $.choices = choices;
            return this;
        }

        /**
         * @param choices One or more blocks configuring a choice to the question, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder choices(List<AccessPackageAssignmentPolicyQuestionChoiceArgs> choices) {
            return choices(Output.of(choices));
        }

        /**
         * @param choices One or more blocks configuring a choice to the question, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder choices(AccessPackageAssignmentPolicyQuestionChoiceArgs... choices) {
            return choices(List.of(choices));
        }

        /**
         * @param required Whether this question is required.
         * 
         * @return builder
         * 
         */
        public Builder required(@Nullable Output<Boolean> required) {
            $.required = required;
            return this;
        }

        /**
         * @param required Whether this question is required.
         * 
         * @return builder
         * 
         */
        public Builder required(Boolean required) {
            return required(Output.of(required));
        }

        /**
         * @param sequence The sequence number of this question.
         * 
         * @return builder
         * 
         */
        public Builder sequence(@Nullable Output<Integer> sequence) {
            $.sequence = sequence;
            return this;
        }

        /**
         * @param sequence The sequence number of this question.
         * 
         * @return builder
         * 
         */
        public Builder sequence(Integer sequence) {
            return sequence(Output.of(sequence));
        }

        /**
         * @param text A block describing the content of this question, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder text(Output<AccessPackageAssignmentPolicyQuestionTextArgs> text) {
            $.text = text;
            return this;
        }

        /**
         * @param text A block describing the content of this question, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder text(AccessPackageAssignmentPolicyQuestionTextArgs text) {
            return text(Output.of(text));
        }

        public AccessPackageAssignmentPolicyQuestionArgs build() {
            if ($.text == null) {
                throw new MissingRequiredPropertyException("AccessPackageAssignmentPolicyQuestionArgs", "text");
            }
            return $;
        }
    }

}

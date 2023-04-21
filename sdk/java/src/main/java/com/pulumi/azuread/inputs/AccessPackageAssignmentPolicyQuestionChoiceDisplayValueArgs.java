// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.azuread.inputs.AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs extends com.pulumi.resources.ResourceArgs {

    public static final AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs Empty = new AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs();

    /**
     * The default text of this question.
     * 
     */
    @Import(name="defaultText", required=true)
    private Output<String> defaultText;

    /**
     * @return The default text of this question.
     * 
     */
    public Output<String> defaultText() {
        return this.defaultText;
    }

    /**
     * One or more blocks describing localized text of this question, as documented below.
     * 
     */
    @Import(name="localizedTexts")
    private @Nullable Output<List<AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs>> localizedTexts;

    /**
     * @return One or more blocks describing localized text of this question, as documented below.
     * 
     */
    public Optional<Output<List<AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs>>> localizedTexts() {
        return Optional.ofNullable(this.localizedTexts);
    }

    private AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs() {}

    private AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs(AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs $) {
        this.defaultText = $.defaultText;
        this.localizedTexts = $.localizedTexts;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs $;

        public Builder() {
            $ = new AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs();
        }

        public Builder(AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs defaults) {
            $ = new AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param defaultText The default text of this question.
         * 
         * @return builder
         * 
         */
        public Builder defaultText(Output<String> defaultText) {
            $.defaultText = defaultText;
            return this;
        }

        /**
         * @param defaultText The default text of this question.
         * 
         * @return builder
         * 
         */
        public Builder defaultText(String defaultText) {
            return defaultText(Output.of(defaultText));
        }

        /**
         * @param localizedTexts One or more blocks describing localized text of this question, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder localizedTexts(@Nullable Output<List<AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs>> localizedTexts) {
            $.localizedTexts = localizedTexts;
            return this;
        }

        /**
         * @param localizedTexts One or more blocks describing localized text of this question, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder localizedTexts(List<AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs> localizedTexts) {
            return localizedTexts(Output.of(localizedTexts));
        }

        /**
         * @param localizedTexts One or more blocks describing localized text of this question, as documented below.
         * 
         * @return builder
         * 
         */
        public Builder localizedTexts(AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs... localizedTexts) {
            return localizedTexts(List.of(localizedTexts));
        }

        public AccessPackageAssignmentPolicyQuestionChoiceDisplayValueArgs build() {
            $.defaultText = Objects.requireNonNull($.defaultText, "expected parameter 'defaultText' to be non-null");
            return $;
        }
    }

}
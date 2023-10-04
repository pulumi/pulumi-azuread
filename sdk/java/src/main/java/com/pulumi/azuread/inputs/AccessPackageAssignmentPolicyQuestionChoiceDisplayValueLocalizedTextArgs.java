// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs extends com.pulumi.resources.ResourceArgs {

    public static final AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs Empty = new AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs();

    /**
     * The localized content of this question choice.
     * 
     */
    @Import(name="content", required=true)
    private Output<String> content;

    /**
     * @return The localized content of this question choice.
     * 
     */
    public Output<String> content() {
        return this.content;
    }

    /**
     * The ISO 639 language code for this question choice content.
     * 
     */
    @Import(name="languageCode", required=true)
    private Output<String> languageCode;

    /**
     * @return The ISO 639 language code for this question choice content.
     * 
     */
    public Output<String> languageCode() {
        return this.languageCode;
    }

    private AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs() {}

    private AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs(AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs $) {
        this.content = $.content;
        this.languageCode = $.languageCode;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs $;

        public Builder() {
            $ = new AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs();
        }

        public Builder(AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs defaults) {
            $ = new AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param content The localized content of this question choice.
         * 
         * @return builder
         * 
         */
        public Builder content(Output<String> content) {
            $.content = content;
            return this;
        }

        /**
         * @param content The localized content of this question choice.
         * 
         * @return builder
         * 
         */
        public Builder content(String content) {
            return content(Output.of(content));
        }

        /**
         * @param languageCode The ISO 639 language code for this question choice content.
         * 
         * @return builder
         * 
         */
        public Builder languageCode(Output<String> languageCode) {
            $.languageCode = languageCode;
            return this;
        }

        /**
         * @param languageCode The ISO 639 language code for this question choice content.
         * 
         * @return builder
         * 
         */
        public Builder languageCode(String languageCode) {
            return languageCode(Output.of(languageCode));
        }

        public AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedTextArgs build() {
            $.content = Objects.requireNonNull($.content, "expected parameter 'content' to be non-null");
            $.languageCode = Objects.requireNonNull($.languageCode, "expected parameter 'languageCode' to be non-null");
            return $;
        }
    }

}

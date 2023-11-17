// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedText {
    /**
     * @return The localized content of this question.`content` (Required) The localized content of this question choice.
     * 
     */
    private String content;
    /**
     * @return The ISO 639 language code for this question content.
     * 
     * `language_code` (Required) The ISO 639 language code for this question choice content.
     * 
     */
    private String languageCode;

    private AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedText() {}
    /**
     * @return The localized content of this question.`content` (Required) The localized content of this question choice.
     * 
     */
    public String content() {
        return this.content;
    }
    /**
     * @return The ISO 639 language code for this question content.
     * 
     * `language_code` (Required) The ISO 639 language code for this question choice content.
     * 
     */
    public String languageCode() {
        return this.languageCode;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedText defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String content;
        private String languageCode;
        public Builder() {}
        public Builder(AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedText defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.content = defaults.content;
    	      this.languageCode = defaults.languageCode;
        }

        @CustomType.Setter
        public Builder content(String content) {
            this.content = Objects.requireNonNull(content);
            return this;
        }
        @CustomType.Setter
        public Builder languageCode(String languageCode) {
            this.languageCode = Objects.requireNonNull(languageCode);
            return this;
        }
        public AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedText build() {
            final var o = new AccessPackageAssignmentPolicyQuestionChoiceDisplayValueLocalizedText();
            o.content = content;
            o.languageCode = languageCode;
            return o;
        }
    }
}

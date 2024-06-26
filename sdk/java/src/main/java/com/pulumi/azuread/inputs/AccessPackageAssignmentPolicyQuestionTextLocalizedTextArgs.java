// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class AccessPackageAssignmentPolicyQuestionTextLocalizedTextArgs extends com.pulumi.resources.ResourceArgs {

    public static final AccessPackageAssignmentPolicyQuestionTextLocalizedTextArgs Empty = new AccessPackageAssignmentPolicyQuestionTextLocalizedTextArgs();

    /**
     * The localized content of this question
     * 
     */
    @Import(name="content", required=true)
    private Output<String> content;

    /**
     * @return The localized content of this question
     * 
     */
    public Output<String> content() {
        return this.content;
    }

    /**
     * The language code of this question content
     * 
     */
    @Import(name="languageCode", required=true)
    private Output<String> languageCode;

    /**
     * @return The language code of this question content
     * 
     */
    public Output<String> languageCode() {
        return this.languageCode;
    }

    private AccessPackageAssignmentPolicyQuestionTextLocalizedTextArgs() {}

    private AccessPackageAssignmentPolicyQuestionTextLocalizedTextArgs(AccessPackageAssignmentPolicyQuestionTextLocalizedTextArgs $) {
        this.content = $.content;
        this.languageCode = $.languageCode;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccessPackageAssignmentPolicyQuestionTextLocalizedTextArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccessPackageAssignmentPolicyQuestionTextLocalizedTextArgs $;

        public Builder() {
            $ = new AccessPackageAssignmentPolicyQuestionTextLocalizedTextArgs();
        }

        public Builder(AccessPackageAssignmentPolicyQuestionTextLocalizedTextArgs defaults) {
            $ = new AccessPackageAssignmentPolicyQuestionTextLocalizedTextArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param content The localized content of this question
         * 
         * @return builder
         * 
         */
        public Builder content(Output<String> content) {
            $.content = content;
            return this;
        }

        /**
         * @param content The localized content of this question
         * 
         * @return builder
         * 
         */
        public Builder content(String content) {
            return content(Output.of(content));
        }

        /**
         * @param languageCode The language code of this question content
         * 
         * @return builder
         * 
         */
        public Builder languageCode(Output<String> languageCode) {
            $.languageCode = languageCode;
            return this;
        }

        /**
         * @param languageCode The language code of this question content
         * 
         * @return builder
         * 
         */
        public Builder languageCode(String languageCode) {
            return languageCode(Output.of(languageCode));
        }

        public AccessPackageAssignmentPolicyQuestionTextLocalizedTextArgs build() {
            if ($.content == null) {
                throw new MissingRequiredPropertyException("AccessPackageAssignmentPolicyQuestionTextLocalizedTextArgs", "content");
            }
            if ($.languageCode == null) {
                throw new MissingRequiredPropertyException("AccessPackageAssignmentPolicyQuestionTextLocalizedTextArgs", "languageCode");
            }
            return $;
        }
    }

}

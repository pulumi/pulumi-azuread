// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InvitationMessageArgs extends com.pulumi.resources.ResourceArgs {

    public static final InvitationMessageArgs Empty = new InvitationMessageArgs();

    /**
     * Email addresses of additional recipients the invitation message should be sent to. Only 1 additional recipient is currently supported by Azure.
     * 
     */
    @Import(name="additionalRecipients")
    private @Nullable Output<String> additionalRecipients;

    /**
     * @return Email addresses of additional recipients the invitation message should be sent to. Only 1 additional recipient is currently supported by Azure.
     * 
     */
    public Optional<Output<String>> additionalRecipients() {
        return Optional.ofNullable(this.additionalRecipients);
    }

    /**
     * Customized message body you want to send if you don&#39;t want to send the default message. Cannot be specified with `language`.
     * 
     */
    @Import(name="body")
    private @Nullable Output<String> body;

    /**
     * @return Customized message body you want to send if you don&#39;t want to send the default message. Cannot be specified with `language`.
     * 
     */
    public Optional<Output<String>> body() {
        return Optional.ofNullable(this.body);
    }

    /**
     * The language you want to send the default message in. The value specified must be in ISO 639 format. Defaults to `en-US`. Cannot be specified with `body`.
     * 
     */
    @Import(name="language")
    private @Nullable Output<String> language;

    /**
     * @return The language you want to send the default message in. The value specified must be in ISO 639 format. Defaults to `en-US`. Cannot be specified with `body`.
     * 
     */
    public Optional<Output<String>> language() {
        return Optional.ofNullable(this.language);
    }

    private InvitationMessageArgs() {}

    private InvitationMessageArgs(InvitationMessageArgs $) {
        this.additionalRecipients = $.additionalRecipients;
        this.body = $.body;
        this.language = $.language;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InvitationMessageArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InvitationMessageArgs $;

        public Builder() {
            $ = new InvitationMessageArgs();
        }

        public Builder(InvitationMessageArgs defaults) {
            $ = new InvitationMessageArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param additionalRecipients Email addresses of additional recipients the invitation message should be sent to. Only 1 additional recipient is currently supported by Azure.
         * 
         * @return builder
         * 
         */
        public Builder additionalRecipients(@Nullable Output<String> additionalRecipients) {
            $.additionalRecipients = additionalRecipients;
            return this;
        }

        /**
         * @param additionalRecipients Email addresses of additional recipients the invitation message should be sent to. Only 1 additional recipient is currently supported by Azure.
         * 
         * @return builder
         * 
         */
        public Builder additionalRecipients(String additionalRecipients) {
            return additionalRecipients(Output.of(additionalRecipients));
        }

        /**
         * @param body Customized message body you want to send if you don&#39;t want to send the default message. Cannot be specified with `language`.
         * 
         * @return builder
         * 
         */
        public Builder body(@Nullable Output<String> body) {
            $.body = body;
            return this;
        }

        /**
         * @param body Customized message body you want to send if you don&#39;t want to send the default message. Cannot be specified with `language`.
         * 
         * @return builder
         * 
         */
        public Builder body(String body) {
            return body(Output.of(body));
        }

        /**
         * @param language The language you want to send the default message in. The value specified must be in ISO 639 format. Defaults to `en-US`. Cannot be specified with `body`.
         * 
         * @return builder
         * 
         */
        public Builder language(@Nullable Output<String> language) {
            $.language = language;
            return this;
        }

        /**
         * @param language The language you want to send the default message in. The value specified must be in ISO 639 format. Defaults to `en-US`. Cannot be specified with `body`.
         * 
         * @return builder
         * 
         */
        public Builder language(String language) {
            return language(Output.of(language));
        }

        public InvitationMessageArgs build() {
            return $;
        }
    }

}

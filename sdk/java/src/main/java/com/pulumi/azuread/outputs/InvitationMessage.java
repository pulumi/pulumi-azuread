// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InvitationMessage {
    /**
     * @return Email addresses of additional recipients the invitation message should be sent to. Only 1 additional recipient is currently supported by Azure.
     * 
     */
    private @Nullable String additionalRecipients;
    /**
     * @return Customized message body you want to send if you don&#39;t want to send the default message. Cannot be specified with `language`.
     * 
     */
    private @Nullable String body;
    /**
     * @return The language you want to send the default message in. The value specified must be in ISO 639 format. Defaults to `en-US`. Cannot be specified with `body`.
     * 
     */
    private @Nullable String language;

    private InvitationMessage() {}
    /**
     * @return Email addresses of additional recipients the invitation message should be sent to. Only 1 additional recipient is currently supported by Azure.
     * 
     */
    public Optional<String> additionalRecipients() {
        return Optional.ofNullable(this.additionalRecipients);
    }
    /**
     * @return Customized message body you want to send if you don&#39;t want to send the default message. Cannot be specified with `language`.
     * 
     */
    public Optional<String> body() {
        return Optional.ofNullable(this.body);
    }
    /**
     * @return The language you want to send the default message in. The value specified must be in ISO 639 format. Defaults to `en-US`. Cannot be specified with `body`.
     * 
     */
    public Optional<String> language() {
        return Optional.ofNullable(this.language);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InvitationMessage defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String additionalRecipients;
        private @Nullable String body;
        private @Nullable String language;
        public Builder() {}
        public Builder(InvitationMessage defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.additionalRecipients = defaults.additionalRecipients;
    	      this.body = defaults.body;
    	      this.language = defaults.language;
        }

        @CustomType.Setter
        public Builder additionalRecipients(@Nullable String additionalRecipients) {
            this.additionalRecipients = additionalRecipients;
            return this;
        }
        @CustomType.Setter
        public Builder body(@Nullable String body) {
            this.body = body;
            return this;
        }
        @CustomType.Setter
        public Builder language(@Nullable String language) {
            this.language = language;
            return this;
        }
        public InvitationMessage build() {
            final var o = new InvitationMessage();
            o.additionalRecipients = additionalRecipients;
            o.body = body;
            o.language = language;
            return o;
        }
    }
}

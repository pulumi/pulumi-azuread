// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ApplicationFallbackPublicClientState extends com.pulumi.resources.ResourceArgs {

    public static final ApplicationFallbackPublicClientState Empty = new ApplicationFallbackPublicClientState();

    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="applicationId")
    private @Nullable Output<String> applicationId;

    /**
     * @return The resource ID of the application registration. Changing this forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> applicationId() {
        return Optional.ofNullable(this.applicationId);
    }

    /**
     * Whether to enable the application as a fallback public client.
     * 
     * &gt; Some configurations may require the Fallback Public Client setting to be `null`, for this case simply destroy this resource (or don&#39;t use it)
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Whether to enable the application as a fallback public client.
     * 
     * &gt; Some configurations may require the Fallback Public Client setting to be `null`, for this case simply destroy this resource (or don&#39;t use it)
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    private ApplicationFallbackPublicClientState() {}

    private ApplicationFallbackPublicClientState(ApplicationFallbackPublicClientState $) {
        this.applicationId = $.applicationId;
        this.enabled = $.enabled;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ApplicationFallbackPublicClientState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ApplicationFallbackPublicClientState $;

        public Builder() {
            $ = new ApplicationFallbackPublicClientState();
        }

        public Builder(ApplicationFallbackPublicClientState defaults) {
            $ = new ApplicationFallbackPublicClientState(Objects.requireNonNull(defaults));
        }

        /**
         * @param applicationId The resource ID of the application registration. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder applicationId(@Nullable Output<String> applicationId) {
            $.applicationId = applicationId;
            return this;
        }

        /**
         * @param applicationId The resource ID of the application registration. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder applicationId(String applicationId) {
            return applicationId(Output.of(applicationId));
        }

        /**
         * @param enabled Whether to enable the application as a fallback public client.
         * 
         * &gt; Some configurations may require the Fallback Public Client setting to be `null`, for this case simply destroy this resource (or don&#39;t use it)
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Whether to enable the application as a fallback public client.
         * 
         * &gt; Some configurations may require the Fallback Public Client setting to be `null`, for this case simply destroy this resource (or don&#39;t use it)
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        public ApplicationFallbackPublicClientState build() {
            return $;
        }
    }

}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.azuread.inputs.SynchronizationSecretCredentialArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SynchronizationSecretState extends com.pulumi.resources.ResourceArgs {

    public static final SynchronizationSecretState Empty = new SynchronizationSecretState();

    /**
     * One or more `credential` blocks as documented below.
     * 
     */
    @Import(name="credentials")
    private @Nullable Output<List<SynchronizationSecretCredentialArgs>> credentials;

    /**
     * @return One or more `credential` blocks as documented below.
     * 
     */
    public Optional<Output<List<SynchronizationSecretCredentialArgs>>> credentials() {
        return Optional.ofNullable(this.credentials);
    }

    /**
     * The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
     * 
     */
    @Import(name="servicePrincipalId")
    private @Nullable Output<String> servicePrincipalId;

    /**
     * @return The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> servicePrincipalId() {
        return Optional.ofNullable(this.servicePrincipalId);
    }

    private SynchronizationSecretState() {}

    private SynchronizationSecretState(SynchronizationSecretState $) {
        this.credentials = $.credentials;
        this.servicePrincipalId = $.servicePrincipalId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SynchronizationSecretState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SynchronizationSecretState $;

        public Builder() {
            $ = new SynchronizationSecretState();
        }

        public Builder(SynchronizationSecretState defaults) {
            $ = new SynchronizationSecretState(Objects.requireNonNull(defaults));
        }

        /**
         * @param credentials One or more `credential` blocks as documented below.
         * 
         * @return builder
         * 
         */
        public Builder credentials(@Nullable Output<List<SynchronizationSecretCredentialArgs>> credentials) {
            $.credentials = credentials;
            return this;
        }

        /**
         * @param credentials One or more `credential` blocks as documented below.
         * 
         * @return builder
         * 
         */
        public Builder credentials(List<SynchronizationSecretCredentialArgs> credentials) {
            return credentials(Output.of(credentials));
        }

        /**
         * @param credentials One or more `credential` blocks as documented below.
         * 
         * @return builder
         * 
         */
        public Builder credentials(SynchronizationSecretCredentialArgs... credentials) {
            return credentials(List.of(credentials));
        }

        /**
         * @param servicePrincipalId The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder servicePrincipalId(@Nullable Output<String> servicePrincipalId) {
            $.servicePrincipalId = servicePrincipalId;
            return this;
        }

        /**
         * @param servicePrincipalId The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder servicePrincipalId(String servicePrincipalId) {
            return servicePrincipalId(Output.of(servicePrincipalId));
        }

        public SynchronizationSecretState build() {
            return $;
        }
    }

}

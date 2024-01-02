// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class SynchronizationSecretCredentialArgs extends com.pulumi.resources.ResourceArgs {

    public static final SynchronizationSecretCredentialArgs Empty = new SynchronizationSecretCredentialArgs();

    /**
     * The key of the secret.
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return The key of the secret.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    /**
     * The value of the secret.
     * 
     */
    @Import(name="value", required=true)
    private Output<String> value;

    /**
     * @return The value of the secret.
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    private SynchronizationSecretCredentialArgs() {}

    private SynchronizationSecretCredentialArgs(SynchronizationSecretCredentialArgs $) {
        this.key = $.key;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SynchronizationSecretCredentialArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SynchronizationSecretCredentialArgs $;

        public Builder() {
            $ = new SynchronizationSecretCredentialArgs();
        }

        public Builder(SynchronizationSecretCredentialArgs defaults) {
            $ = new SynchronizationSecretCredentialArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param key The key of the secret.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The key of the secret.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param value The value of the secret.
         * 
         * @return builder
         * 
         */
        public Builder value(Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value The value of the secret.
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public SynchronizationSecretCredentialArgs build() {
            if ($.key == null) {
                throw new MissingRequiredPropertyException("SynchronizationSecretCredentialArgs", "key");
            }
            if ($.value == null) {
                throw new MissingRequiredPropertyException("SynchronizationSecretCredentialArgs", "value");
            }
            return $;
        }
    }

}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ApplicationPasswordArgs extends com.pulumi.resources.ResourceArgs {

    public static final ApplicationPasswordArgs Empty = new ApplicationPasswordArgs();

    /**
     * A display name for the password. Changing this field forces a new resource to be created.
     * 
     */
    @Import(name="displayName", required=true)
    private Output<String> displayName;

    /**
     * @return A display name for the password. Changing this field forces a new resource to be created.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }

    /**
     * The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
     * 
     */
    @Import(name="endDate")
    private @Nullable Output<String> endDate;

    /**
     * @return The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> endDate() {
        return Optional.ofNullable(this.endDate);
    }

    /**
     * (Required) The unique key ID for the generated password.
     * 
     */
    @Import(name="keyId")
    private @Nullable Output<String> keyId;

    /**
     * @return (Required) The unique key ID for the generated password.
     * 
     */
    public Optional<Output<String>> keyId() {
        return Optional.ofNullable(this.keyId);
    }

    /**
     * The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn&#39;t specified, the current date is used.  Changing this field forces a new resource to be created.
     * 
     */
    @Import(name="startDate")
    private @Nullable Output<String> startDate;

    /**
     * @return The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn&#39;t specified, the current date is used.  Changing this field forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> startDate() {
        return Optional.ofNullable(this.startDate);
    }

    /**
     * (Required) The generated password for the application.
     * 
     */
    @Import(name="value")
    private @Nullable Output<String> value;

    /**
     * @return (Required) The generated password for the application.
     * 
     */
    public Optional<Output<String>> value() {
        return Optional.ofNullable(this.value);
    }

    private ApplicationPasswordArgs() {}

    private ApplicationPasswordArgs(ApplicationPasswordArgs $) {
        this.displayName = $.displayName;
        this.endDate = $.endDate;
        this.keyId = $.keyId;
        this.startDate = $.startDate;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ApplicationPasswordArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ApplicationPasswordArgs $;

        public Builder() {
            $ = new ApplicationPasswordArgs();
        }

        public Builder(ApplicationPasswordArgs defaults) {
            $ = new ApplicationPasswordArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param displayName A display name for the password. Changing this field forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder displayName(Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName A display name for the password. Changing this field forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param endDate The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder endDate(@Nullable Output<String> endDate) {
            $.endDate = endDate;
            return this;
        }

        /**
         * @param endDate The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder endDate(String endDate) {
            return endDate(Output.of(endDate));
        }

        /**
         * @param keyId (Required) The unique key ID for the generated password.
         * 
         * @return builder
         * 
         */
        public Builder keyId(@Nullable Output<String> keyId) {
            $.keyId = keyId;
            return this;
        }

        /**
         * @param keyId (Required) The unique key ID for the generated password.
         * 
         * @return builder
         * 
         */
        public Builder keyId(String keyId) {
            return keyId(Output.of(keyId));
        }

        /**
         * @param startDate The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn&#39;t specified, the current date is used.  Changing this field forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder startDate(@Nullable Output<String> startDate) {
            $.startDate = startDate;
            return this;
        }

        /**
         * @param startDate The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn&#39;t specified, the current date is used.  Changing this field forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder startDate(String startDate) {
            return startDate(Output.of(startDate));
        }

        /**
         * @param value (Required) The generated password for the application.
         * 
         * @return builder
         * 
         */
        public Builder value(@Nullable Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value (Required) The generated password for the application.
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public ApplicationPasswordArgs build() {
            if ($.displayName == null) {
                throw new MissingRequiredPropertyException("ApplicationPasswordArgs", "displayName");
            }
            return $;
        }
    }

}

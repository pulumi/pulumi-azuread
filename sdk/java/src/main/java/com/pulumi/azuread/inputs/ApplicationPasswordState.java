// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ApplicationPasswordState extends com.pulumi.resources.ResourceArgs {

    public static final ApplicationPasswordState Empty = new ApplicationPasswordState();

    /**
     * The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
     * 
     */
    @Import(name="applicationId")
    private @Nullable Output<String> applicationId;

    /**
     * @return The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> applicationId() {
        return Optional.ofNullable(this.applicationId);
    }

    /**
     * The object ID of the application for which this password should be created
     * 
     * @deprecated
     * The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
     * 
     */
    @Deprecated /* The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider */
    @Import(name="applicationObjectId")
    private @Nullable Output<String> applicationObjectId;

    /**
     * @return The object ID of the application for which this password should be created
     * 
     * @deprecated
     * The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
     * 
     */
    @Deprecated /* The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider */
    public Optional<Output<String>> applicationObjectId() {
        return Optional.ofNullable(this.applicationObjectId);
    }

    /**
     * A display name for the password. Changing this field forces a new resource to be created.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return A display name for the password. Changing this field forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
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
     * A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
     * 
     */
    @Import(name="endDateRelative")
    private @Nullable Output<String> endDateRelative;

    /**
     * @return A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> endDateRelative() {
        return Optional.ofNullable(this.endDateRelative);
    }

    /**
     * A UUID used to uniquely identify this password credential.
     * 
     */
    @Import(name="keyId")
    private @Nullable Output<String> keyId;

    /**
     * @return A UUID used to uniquely identify this password credential.
     * 
     */
    public Optional<Output<String>> keyId() {
        return Optional.ofNullable(this.keyId);
    }

    /**
     * A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="rotateWhenChanged")
    private @Nullable Output<Map<String,String>> rotateWhenChanged;

    /**
     * @return A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
     * 
     */
    public Optional<Output<Map<String,String>>> rotateWhenChanged() {
        return Optional.ofNullable(this.rotateWhenChanged);
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
     * The password for this application, which is generated by Azure Active Directory.
     * 
     */
    @Import(name="value")
    private @Nullable Output<String> value;

    /**
     * @return The password for this application, which is generated by Azure Active Directory.
     * 
     */
    public Optional<Output<String>> value() {
        return Optional.ofNullable(this.value);
    }

    private ApplicationPasswordState() {}

    private ApplicationPasswordState(ApplicationPasswordState $) {
        this.applicationId = $.applicationId;
        this.applicationObjectId = $.applicationObjectId;
        this.displayName = $.displayName;
        this.endDate = $.endDate;
        this.endDateRelative = $.endDateRelative;
        this.keyId = $.keyId;
        this.rotateWhenChanged = $.rotateWhenChanged;
        this.startDate = $.startDate;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ApplicationPasswordState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ApplicationPasswordState $;

        public Builder() {
            $ = new ApplicationPasswordState();
        }

        public Builder(ApplicationPasswordState defaults) {
            $ = new ApplicationPasswordState(Objects.requireNonNull(defaults));
        }

        /**
         * @param applicationId The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder applicationId(@Nullable Output<String> applicationId) {
            $.applicationId = applicationId;
            return this;
        }

        /**
         * @param applicationId The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder applicationId(String applicationId) {
            return applicationId(Output.of(applicationId));
        }

        /**
         * @param applicationObjectId The object ID of the application for which this password should be created
         * 
         * @return builder
         * 
         * @deprecated
         * The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
         * 
         */
        @Deprecated /* The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider */
        public Builder applicationObjectId(@Nullable Output<String> applicationObjectId) {
            $.applicationObjectId = applicationObjectId;
            return this;
        }

        /**
         * @param applicationObjectId The object ID of the application for which this password should be created
         * 
         * @return builder
         * 
         * @deprecated
         * The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider
         * 
         */
        @Deprecated /* The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider */
        public Builder applicationObjectId(String applicationObjectId) {
            return applicationObjectId(Output.of(applicationObjectId));
        }

        /**
         * @param displayName A display name for the password. Changing this field forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
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
         * @param endDateRelative A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder endDateRelative(@Nullable Output<String> endDateRelative) {
            $.endDateRelative = endDateRelative;
            return this;
        }

        /**
         * @param endDateRelative A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder endDateRelative(String endDateRelative) {
            return endDateRelative(Output.of(endDateRelative));
        }

        /**
         * @param keyId A UUID used to uniquely identify this password credential.
         * 
         * @return builder
         * 
         */
        public Builder keyId(@Nullable Output<String> keyId) {
            $.keyId = keyId;
            return this;
        }

        /**
         * @param keyId A UUID used to uniquely identify this password credential.
         * 
         * @return builder
         * 
         */
        public Builder keyId(String keyId) {
            return keyId(Output.of(keyId));
        }

        /**
         * @param rotateWhenChanged A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder rotateWhenChanged(@Nullable Output<Map<String,String>> rotateWhenChanged) {
            $.rotateWhenChanged = rotateWhenChanged;
            return this;
        }

        /**
         * @param rotateWhenChanged A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder rotateWhenChanged(Map<String,String> rotateWhenChanged) {
            return rotateWhenChanged(Output.of(rotateWhenChanged));
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
         * @param value The password for this application, which is generated by Azure Active Directory.
         * 
         * @return builder
         * 
         */
        public Builder value(@Nullable Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value The password for this application, which is generated by Azure Active Directory.
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public ApplicationPasswordState build() {
            return $;
        }
    }

}

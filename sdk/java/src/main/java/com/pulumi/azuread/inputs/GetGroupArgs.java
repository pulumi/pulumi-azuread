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


public final class GetGroupArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGroupArgs Empty = new GetGroupArgs();

    /**
     * The display name for the group.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return The display name for the group.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * Whether the group is mail-enabled.
     * 
     */
    @Import(name="mailEnabled")
    private @Nullable Output<Boolean> mailEnabled;

    /**
     * @return Whether the group is mail-enabled.
     * 
     */
    public Optional<Output<Boolean>> mailEnabled() {
        return Optional.ofNullable(this.mailEnabled);
    }

    /**
     * The mail alias for the group, unique in the organisation.
     * 
     */
    @Import(name="mailNickname")
    private @Nullable Output<String> mailNickname;

    /**
     * @return The mail alias for the group, unique in the organisation.
     * 
     */
    public Optional<Output<String>> mailNickname() {
        return Optional.ofNullable(this.mailNickname);
    }

    /**
     * Specifies the object ID of the group.
     * 
     */
    @Import(name="objectId")
    private @Nullable Output<String> objectId;

    /**
     * @return Specifies the object ID of the group.
     * 
     */
    public Optional<Output<String>> objectId() {
        return Optional.ofNullable(this.objectId);
    }

    /**
     * Whether the group is a security group.
     * 
     * &gt; One of `display_name`, `object_id` or `mail_nickname` must be specified.
     * 
     */
    @Import(name="securityEnabled")
    private @Nullable Output<Boolean> securityEnabled;

    /**
     * @return Whether the group is a security group.
     * 
     * &gt; One of `display_name`, `object_id` or `mail_nickname` must be specified.
     * 
     */
    public Optional<Output<Boolean>> securityEnabled() {
        return Optional.ofNullable(this.securityEnabled);
    }

    private GetGroupArgs() {}

    private GetGroupArgs(GetGroupArgs $) {
        this.displayName = $.displayName;
        this.mailEnabled = $.mailEnabled;
        this.mailNickname = $.mailNickname;
        this.objectId = $.objectId;
        this.securityEnabled = $.securityEnabled;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGroupArgs $;

        public Builder() {
            $ = new GetGroupArgs();
        }

        public Builder(GetGroupArgs defaults) {
            $ = new GetGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param displayName The display name for the group.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName The display name for the group.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param mailEnabled Whether the group is mail-enabled.
         * 
         * @return builder
         * 
         */
        public Builder mailEnabled(@Nullable Output<Boolean> mailEnabled) {
            $.mailEnabled = mailEnabled;
            return this;
        }

        /**
         * @param mailEnabled Whether the group is mail-enabled.
         * 
         * @return builder
         * 
         */
        public Builder mailEnabled(Boolean mailEnabled) {
            return mailEnabled(Output.of(mailEnabled));
        }

        /**
         * @param mailNickname The mail alias for the group, unique in the organisation.
         * 
         * @return builder
         * 
         */
        public Builder mailNickname(@Nullable Output<String> mailNickname) {
            $.mailNickname = mailNickname;
            return this;
        }

        /**
         * @param mailNickname The mail alias for the group, unique in the organisation.
         * 
         * @return builder
         * 
         */
        public Builder mailNickname(String mailNickname) {
            return mailNickname(Output.of(mailNickname));
        }

        /**
         * @param objectId Specifies the object ID of the group.
         * 
         * @return builder
         * 
         */
        public Builder objectId(@Nullable Output<String> objectId) {
            $.objectId = objectId;
            return this;
        }

        /**
         * @param objectId Specifies the object ID of the group.
         * 
         * @return builder
         * 
         */
        public Builder objectId(String objectId) {
            return objectId(Output.of(objectId));
        }

        /**
         * @param securityEnabled Whether the group is a security group.
         * 
         * &gt; One of `display_name`, `object_id` or `mail_nickname` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder securityEnabled(@Nullable Output<Boolean> securityEnabled) {
            $.securityEnabled = securityEnabled;
            return this;
        }

        /**
         * @param securityEnabled Whether the group is a security group.
         * 
         * &gt; One of `display_name`, `object_id` or `mail_nickname` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder securityEnabled(Boolean securityEnabled) {
            return securityEnabled(Output.of(securityEnabled));
        }

        public GetGroupArgs build() {
            return $;
        }
    }

}

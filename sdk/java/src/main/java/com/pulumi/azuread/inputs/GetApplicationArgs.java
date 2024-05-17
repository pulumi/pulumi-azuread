// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetApplicationArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetApplicationArgs Empty = new GetApplicationArgs();

    /**
     * @deprecated
     * The `application_id` property has been replaced with the `client_id` property and will be removed in version 3.0 of the AzureAD provider
     * 
     */
    @Deprecated /* The `application_id` property has been replaced with the `client_id` property and will be removed in version 3.0 of the AzureAD provider */
    @Import(name="applicationId")
    private @Nullable Output<String> applicationId;

    /**
     * @deprecated
     * The `application_id` property has been replaced with the `client_id` property and will be removed in version 3.0 of the AzureAD provider
     * 
     */
    @Deprecated /* The `application_id` property has been replaced with the `client_id` property and will be removed in version 3.0 of the AzureAD provider */
    public Optional<Output<String>> applicationId() {
        return Optional.ofNullable(this.applicationId);
    }

    /**
     * Specifies the Client ID of the application.
     * 
     */
    @Import(name="clientId")
    private @Nullable Output<String> clientId;

    /**
     * @return Specifies the Client ID of the application.
     * 
     */
    public Optional<Output<String>> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    /**
     * Specifies the display name of the application.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return Specifies the display name of the application.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * Specifies any identifier URI of the application. See also the `identifier_uris` attribute which contains a list of all identifier URIs for the application.
     * 
     * &gt; One of `client_id`, `display_name`, `object_id`, or `identifier_uri` must be specified.
     * 
     */
    @Import(name="identifierUri")
    private @Nullable Output<String> identifierUri;

    /**
     * @return Specifies any identifier URI of the application. See also the `identifier_uris` attribute which contains a list of all identifier URIs for the application.
     * 
     * &gt; One of `client_id`, `display_name`, `object_id`, or `identifier_uri` must be specified.
     * 
     */
    public Optional<Output<String>> identifierUri() {
        return Optional.ofNullable(this.identifierUri);
    }

    /**
     * Specifies the Object ID of the application.
     * 
     */
    @Import(name="objectId")
    private @Nullable Output<String> objectId;

    /**
     * @return Specifies the Object ID of the application.
     * 
     */
    public Optional<Output<String>> objectId() {
        return Optional.ofNullable(this.objectId);
    }

    private GetApplicationArgs() {}

    private GetApplicationArgs(GetApplicationArgs $) {
        this.applicationId = $.applicationId;
        this.clientId = $.clientId;
        this.displayName = $.displayName;
        this.identifierUri = $.identifierUri;
        this.objectId = $.objectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetApplicationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetApplicationArgs $;

        public Builder() {
            $ = new GetApplicationArgs();
        }

        public Builder(GetApplicationArgs defaults) {
            $ = new GetApplicationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @return builder
         * 
         * @deprecated
         * The `application_id` property has been replaced with the `client_id` property and will be removed in version 3.0 of the AzureAD provider
         * 
         */
        @Deprecated /* The `application_id` property has been replaced with the `client_id` property and will be removed in version 3.0 of the AzureAD provider */
        public Builder applicationId(@Nullable Output<String> applicationId) {
            $.applicationId = applicationId;
            return this;
        }

        /**
         * @return builder
         * 
         * @deprecated
         * The `application_id` property has been replaced with the `client_id` property and will be removed in version 3.0 of the AzureAD provider
         * 
         */
        @Deprecated /* The `application_id` property has been replaced with the `client_id` property and will be removed in version 3.0 of the AzureAD provider */
        public Builder applicationId(String applicationId) {
            return applicationId(Output.of(applicationId));
        }

        /**
         * @param clientId Specifies the Client ID of the application.
         * 
         * @return builder
         * 
         */
        public Builder clientId(@Nullable Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        /**
         * @param clientId Specifies the Client ID of the application.
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        /**
         * @param displayName Specifies the display name of the application.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName Specifies the display name of the application.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param identifierUri Specifies any identifier URI of the application. See also the `identifier_uris` attribute which contains a list of all identifier URIs for the application.
         * 
         * &gt; One of `client_id`, `display_name`, `object_id`, or `identifier_uri` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder identifierUri(@Nullable Output<String> identifierUri) {
            $.identifierUri = identifierUri;
            return this;
        }

        /**
         * @param identifierUri Specifies any identifier URI of the application. See also the `identifier_uris` attribute which contains a list of all identifier URIs for the application.
         * 
         * &gt; One of `client_id`, `display_name`, `object_id`, or `identifier_uri` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder identifierUri(String identifierUri) {
            return identifierUri(Output.of(identifierUri));
        }

        /**
         * @param objectId Specifies the Object ID of the application.
         * 
         * @return builder
         * 
         */
        public Builder objectId(@Nullable Output<String> objectId) {
            $.objectId = objectId;
            return this;
        }

        /**
         * @param objectId Specifies the Object ID of the application.
         * 
         * @return builder
         * 
         */
        public Builder objectId(String objectId) {
            return objectId(Output.of(objectId));
        }

        public GetApplicationArgs build() {
            return $;
        }
    }

}

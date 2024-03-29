// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAccessPackageCatalogArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAccessPackageCatalogArgs Empty = new GetAccessPackageCatalogArgs();

    /**
     * The display name of the access package catalog.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return The display name of the access package catalog.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * The ID of this access package catalog.
     * 
     * &gt; One of `display_name` or `object_id` must be specified.
     * 
     */
    @Import(name="objectId")
    private @Nullable Output<String> objectId;

    /**
     * @return The ID of this access package catalog.
     * 
     * &gt; One of `display_name` or `object_id` must be specified.
     * 
     */
    public Optional<Output<String>> objectId() {
        return Optional.ofNullable(this.objectId);
    }

    private GetAccessPackageCatalogArgs() {}

    private GetAccessPackageCatalogArgs(GetAccessPackageCatalogArgs $) {
        this.displayName = $.displayName;
        this.objectId = $.objectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAccessPackageCatalogArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAccessPackageCatalogArgs $;

        public Builder() {
            $ = new GetAccessPackageCatalogArgs();
        }

        public Builder(GetAccessPackageCatalogArgs defaults) {
            $ = new GetAccessPackageCatalogArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param displayName The display name of the access package catalog.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName The display name of the access package catalog.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param objectId The ID of this access package catalog.
         * 
         * &gt; One of `display_name` or `object_id` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder objectId(@Nullable Output<String> objectId) {
            $.objectId = objectId;
            return this;
        }

        /**
         * @param objectId The ID of this access package catalog.
         * 
         * &gt; One of `display_name` or `object_id` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder objectId(String objectId) {
            return objectId(Output.of(objectId));
        }

        public GetAccessPackageCatalogArgs build() {
            return $;
        }
    }

}

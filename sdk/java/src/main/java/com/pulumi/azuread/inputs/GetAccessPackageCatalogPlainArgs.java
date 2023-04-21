// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAccessPackageCatalogPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAccessPackageCatalogPlainArgs Empty = new GetAccessPackageCatalogPlainArgs();

    /**
     * The display name of the access package catalog.
     * 
     */
    @Import(name="displayName")
    private @Nullable String displayName;

    /**
     * @return The display name of the access package catalog.
     * 
     */
    public Optional<String> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * The ID of this access package catalog.
     * 
     */
    @Import(name="objectId")
    private @Nullable String objectId;

    /**
     * @return The ID of this access package catalog.
     * 
     */
    public Optional<String> objectId() {
        return Optional.ofNullable(this.objectId);
    }

    private GetAccessPackageCatalogPlainArgs() {}

    private GetAccessPackageCatalogPlainArgs(GetAccessPackageCatalogPlainArgs $) {
        this.displayName = $.displayName;
        this.objectId = $.objectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAccessPackageCatalogPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAccessPackageCatalogPlainArgs $;

        public Builder() {
            $ = new GetAccessPackageCatalogPlainArgs();
        }

        public Builder(GetAccessPackageCatalogPlainArgs defaults) {
            $ = new GetAccessPackageCatalogPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param displayName The display name of the access package catalog.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable String displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param objectId The ID of this access package catalog.
         * 
         * @return builder
         * 
         */
        public Builder objectId(@Nullable String objectId) {
            $.objectId = objectId;
            return this;
        }

        public GetAccessPackageCatalogPlainArgs build() {
            return $;
        }
    }

}

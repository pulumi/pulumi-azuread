// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAccessPackagePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAccessPackagePlainArgs Empty = new GetAccessPackagePlainArgs();

    /**
     * The ID of the Catalog this access package is in.
     * 
     */
    @Import(name="catalogId")
    private @Nullable String catalogId;

    /**
     * @return The ID of the Catalog this access package is in.
     * 
     */
    public Optional<String> catalogId() {
        return Optional.ofNullable(this.catalogId);
    }

    /**
     * The display name of the access package.
     * 
     */
    @Import(name="displayName")
    private @Nullable String displayName;

    /**
     * @return The display name of the access package.
     * 
     */
    public Optional<String> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * The ID of this access package.
     * 
     * &gt; Either `object_id`, or both `catalog_id` and `display_name`, must be specified.
     * 
     */
    @Import(name="objectId")
    private @Nullable String objectId;

    /**
     * @return The ID of this access package.
     * 
     * &gt; Either `object_id`, or both `catalog_id` and `display_name`, must be specified.
     * 
     */
    public Optional<String> objectId() {
        return Optional.ofNullable(this.objectId);
    }

    private GetAccessPackagePlainArgs() {}

    private GetAccessPackagePlainArgs(GetAccessPackagePlainArgs $) {
        this.catalogId = $.catalogId;
        this.displayName = $.displayName;
        this.objectId = $.objectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAccessPackagePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAccessPackagePlainArgs $;

        public Builder() {
            $ = new GetAccessPackagePlainArgs();
        }

        public Builder(GetAccessPackagePlainArgs defaults) {
            $ = new GetAccessPackagePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param catalogId The ID of the Catalog this access package is in.
         * 
         * @return builder
         * 
         */
        public Builder catalogId(@Nullable String catalogId) {
            $.catalogId = catalogId;
            return this;
        }

        /**
         * @param displayName The display name of the access package.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable String displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param objectId The ID of this access package.
         * 
         * &gt; Either `object_id`, or both `catalog_id` and `display_name`, must be specified.
         * 
         * @return builder
         * 
         */
        public Builder objectId(@Nullable String objectId) {
            $.objectId = objectId;
            return this;
        }

        public GetAccessPackagePlainArgs build() {
            return $;
        }
    }

}

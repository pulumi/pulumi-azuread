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


public final class AccessPackageCatalogState extends com.pulumi.resources.ResourceArgs {

    public static final AccessPackageCatalogState Empty = new AccessPackageCatalogState();

    /**
     * The description of the access package catalog.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the access package catalog.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

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
     * Whether the access packages in this catalog can be requested by users outside the tenant.
     * 
     */
    @Import(name="externallyVisible")
    private @Nullable Output<Boolean> externallyVisible;

    /**
     * @return Whether the access packages in this catalog can be requested by users outside the tenant.
     * 
     */
    public Optional<Output<Boolean>> externallyVisible() {
        return Optional.ofNullable(this.externallyVisible);
    }

    /**
     * Whether the access packages in this catalog are available for management.
     * 
     */
    @Import(name="published")
    private @Nullable Output<Boolean> published;

    /**
     * @return Whether the access packages in this catalog are available for management.
     * 
     */
    public Optional<Output<Boolean>> published() {
        return Optional.ofNullable(this.published);
    }

    private AccessPackageCatalogState() {}

    private AccessPackageCatalogState(AccessPackageCatalogState $) {
        this.description = $.description;
        this.displayName = $.displayName;
        this.externallyVisible = $.externallyVisible;
        this.published = $.published;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccessPackageCatalogState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccessPackageCatalogState $;

        public Builder() {
            $ = new AccessPackageCatalogState();
        }

        public Builder(AccessPackageCatalogState defaults) {
            $ = new AccessPackageCatalogState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the access package catalog.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the access package catalog.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
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
         * @param externallyVisible Whether the access packages in this catalog can be requested by users outside the tenant.
         * 
         * @return builder
         * 
         */
        public Builder externallyVisible(@Nullable Output<Boolean> externallyVisible) {
            $.externallyVisible = externallyVisible;
            return this;
        }

        /**
         * @param externallyVisible Whether the access packages in this catalog can be requested by users outside the tenant.
         * 
         * @return builder
         * 
         */
        public Builder externallyVisible(Boolean externallyVisible) {
            return externallyVisible(Output.of(externallyVisible));
        }

        /**
         * @param published Whether the access packages in this catalog are available for management.
         * 
         * @return builder
         * 
         */
        public Builder published(@Nullable Output<Boolean> published) {
            $.published = published;
            return this;
        }

        /**
         * @param published Whether the access packages in this catalog are available for management.
         * 
         * @return builder
         * 
         */
        public Builder published(Boolean published) {
            return published(Output.of(published));
        }

        public AccessPackageCatalogState build() {
            return $;
        }
    }

}
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


public final class AccessPackageState extends com.pulumi.resources.ResourceArgs {

    public static final AccessPackageState Empty = new AccessPackageState();

    /**
     * The ID of the Catalog this access package will be created in.
     * 
     */
    @Import(name="catalogId")
    private @Nullable Output<String> catalogId;

    /**
     * @return The ID of the Catalog this access package will be created in.
     * 
     */
    public Optional<Output<String>> catalogId() {
        return Optional.ofNullable(this.catalogId);
    }

    /**
     * The description of the access package.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the access package.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The display name of the access package.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return The display name of the access package.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * Whether the access package is hidden from the requestor.
     * 
     */
    @Import(name="hidden")
    private @Nullable Output<Boolean> hidden;

    /**
     * @return Whether the access package is hidden from the requestor.
     * 
     */
    public Optional<Output<Boolean>> hidden() {
        return Optional.ofNullable(this.hidden);
    }

    private AccessPackageState() {}

    private AccessPackageState(AccessPackageState $) {
        this.catalogId = $.catalogId;
        this.description = $.description;
        this.displayName = $.displayName;
        this.hidden = $.hidden;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccessPackageState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccessPackageState $;

        public Builder() {
            $ = new AccessPackageState();
        }

        public Builder(AccessPackageState defaults) {
            $ = new AccessPackageState(Objects.requireNonNull(defaults));
        }

        /**
         * @param catalogId The ID of the Catalog this access package will be created in.
         * 
         * @return builder
         * 
         */
        public Builder catalogId(@Nullable Output<String> catalogId) {
            $.catalogId = catalogId;
            return this;
        }

        /**
         * @param catalogId The ID of the Catalog this access package will be created in.
         * 
         * @return builder
         * 
         */
        public Builder catalogId(String catalogId) {
            return catalogId(Output.of(catalogId));
        }

        /**
         * @param description The description of the access package.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the access package.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param displayName The display name of the access package.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName The display name of the access package.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param hidden Whether the access package is hidden from the requestor.
         * 
         * @return builder
         * 
         */
        public Builder hidden(@Nullable Output<Boolean> hidden) {
            $.hidden = hidden;
            return this;
        }

        /**
         * @param hidden Whether the access package is hidden from the requestor.
         * 
         * @return builder
         * 
         */
        public Builder hidden(Boolean hidden) {
            return hidden(Output.of(hidden));
        }

        public AccessPackageState build() {
            return $;
        }
    }

}

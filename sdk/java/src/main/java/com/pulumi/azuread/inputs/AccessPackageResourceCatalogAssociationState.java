// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AccessPackageResourceCatalogAssociationState extends com.pulumi.resources.ResourceArgs {

    public static final AccessPackageResourceCatalogAssociationState Empty = new AccessPackageResourceCatalogAssociationState();

    /**
     * The unique ID of the access package catalog. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="catalogId")
    private @Nullable Output<String> catalogId;

    /**
     * @return The unique ID of the access package catalog. Changing this forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> catalogId() {
        return Optional.ofNullable(this.catalogId);
    }

    /**
     * The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="resourceOriginId")
    private @Nullable Output<String> resourceOriginId;

    /**
     * @return The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> resourceOriginId() {
        return Optional.ofNullable(this.resourceOriginId);
    }

    /**
     * The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="resourceOriginSystem")
    private @Nullable Output<String> resourceOriginSystem;

    /**
     * @return The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> resourceOriginSystem() {
        return Optional.ofNullable(this.resourceOriginSystem);
    }

    private AccessPackageResourceCatalogAssociationState() {}

    private AccessPackageResourceCatalogAssociationState(AccessPackageResourceCatalogAssociationState $) {
        this.catalogId = $.catalogId;
        this.resourceOriginId = $.resourceOriginId;
        this.resourceOriginSystem = $.resourceOriginSystem;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccessPackageResourceCatalogAssociationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccessPackageResourceCatalogAssociationState $;

        public Builder() {
            $ = new AccessPackageResourceCatalogAssociationState();
        }

        public Builder(AccessPackageResourceCatalogAssociationState defaults) {
            $ = new AccessPackageResourceCatalogAssociationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param catalogId The unique ID of the access package catalog. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder catalogId(@Nullable Output<String> catalogId) {
            $.catalogId = catalogId;
            return this;
        }

        /**
         * @param catalogId The unique ID of the access package catalog. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder catalogId(String catalogId) {
            return catalogId(Output.of(catalogId));
        }

        /**
         * @param resourceOriginId The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder resourceOriginId(@Nullable Output<String> resourceOriginId) {
            $.resourceOriginId = resourceOriginId;
            return this;
        }

        /**
         * @param resourceOriginId The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder resourceOriginId(String resourceOriginId) {
            return resourceOriginId(Output.of(resourceOriginId));
        }

        /**
         * @param resourceOriginSystem The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder resourceOriginSystem(@Nullable Output<String> resourceOriginSystem) {
            $.resourceOriginSystem = resourceOriginSystem;
            return this;
        }

        /**
         * @param resourceOriginSystem The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder resourceOriginSystem(String resourceOriginSystem) {
            return resourceOriginSystem(Output.of(resourceOriginSystem));
        }

        public AccessPackageResourceCatalogAssociationState build() {
            return $;
        }
    }

}

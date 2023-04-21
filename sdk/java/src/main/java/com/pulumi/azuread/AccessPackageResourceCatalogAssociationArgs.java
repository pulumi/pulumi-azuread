// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class AccessPackageResourceCatalogAssociationArgs extends com.pulumi.resources.ResourceArgs {

    public static final AccessPackageResourceCatalogAssociationArgs Empty = new AccessPackageResourceCatalogAssociationArgs();

    /**
     * The unique ID of the access package catalog. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="catalogId", required=true)
    private Output<String> catalogId;

    /**
     * @return The unique ID of the access package catalog. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> catalogId() {
        return this.catalogId;
    }

    /**
     * The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="resourceOriginId", required=true)
    private Output<String> resourceOriginId;

    /**
     * @return The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> resourceOriginId() {
        return this.resourceOriginId;
    }

    /**
     * The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="resourceOriginSystem", required=true)
    private Output<String> resourceOriginSystem;

    /**
     * @return The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> resourceOriginSystem() {
        return this.resourceOriginSystem;
    }

    private AccessPackageResourceCatalogAssociationArgs() {}

    private AccessPackageResourceCatalogAssociationArgs(AccessPackageResourceCatalogAssociationArgs $) {
        this.catalogId = $.catalogId;
        this.resourceOriginId = $.resourceOriginId;
        this.resourceOriginSystem = $.resourceOriginSystem;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccessPackageResourceCatalogAssociationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccessPackageResourceCatalogAssociationArgs $;

        public Builder() {
            $ = new AccessPackageResourceCatalogAssociationArgs();
        }

        public Builder(AccessPackageResourceCatalogAssociationArgs defaults) {
            $ = new AccessPackageResourceCatalogAssociationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param catalogId The unique ID of the access package catalog. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder catalogId(Output<String> catalogId) {
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
        public Builder resourceOriginId(Output<String> resourceOriginId) {
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
        public Builder resourceOriginSystem(Output<String> resourceOriginSystem) {
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

        public AccessPackageResourceCatalogAssociationArgs build() {
            $.catalogId = Objects.requireNonNull($.catalogId, "expected parameter 'catalogId' to be non-null");
            $.resourceOriginId = Objects.requireNonNull($.resourceOriginId, "expected parameter 'resourceOriginId' to be non-null");
            $.resourceOriginSystem = Objects.requireNonNull($.resourceOriginSystem, "expected parameter 'resourceOriginSystem' to be non-null");
            return $;
        }
    }

}

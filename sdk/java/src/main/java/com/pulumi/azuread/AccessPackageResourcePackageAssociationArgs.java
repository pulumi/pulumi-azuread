// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AccessPackageResourcePackageAssociationArgs extends com.pulumi.resources.ResourceArgs {

    public static final AccessPackageResourcePackageAssociationArgs Empty = new AccessPackageResourcePackageAssociationArgs();

    /**
     * The ID of access package this resource association is configured to. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="accessPackageId", required=true)
    private Output<String> accessPackageId;

    /**
     * @return The ID of access package this resource association is configured to. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> accessPackageId() {
        return this.accessPackageId;
    }

    /**
     * The role of access type to the specified resource. Valid values are `Member`, or `Owner` The default is `Member`. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="accessType")
    private @Nullable Output<String> accessType;

    /**
     * @return The role of access type to the specified resource. Valid values are `Member`, or `Owner` The default is `Member`. Changing this forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> accessType() {
        return Optional.ofNullable(this.accessType);
    }

    /**
     * The ID of the catalog association from the `azuread.AccessPackageResourceCatalogAssociation` resource. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="catalogResourceAssociationId", required=true)
    private Output<String> catalogResourceAssociationId;

    /**
     * @return The ID of the catalog association from the `azuread.AccessPackageResourceCatalogAssociation` resource. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> catalogResourceAssociationId() {
        return this.catalogResourceAssociationId;
    }

    private AccessPackageResourcePackageAssociationArgs() {}

    private AccessPackageResourcePackageAssociationArgs(AccessPackageResourcePackageAssociationArgs $) {
        this.accessPackageId = $.accessPackageId;
        this.accessType = $.accessType;
        this.catalogResourceAssociationId = $.catalogResourceAssociationId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccessPackageResourcePackageAssociationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccessPackageResourcePackageAssociationArgs $;

        public Builder() {
            $ = new AccessPackageResourcePackageAssociationArgs();
        }

        public Builder(AccessPackageResourcePackageAssociationArgs defaults) {
            $ = new AccessPackageResourcePackageAssociationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessPackageId The ID of access package this resource association is configured to. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder accessPackageId(Output<String> accessPackageId) {
            $.accessPackageId = accessPackageId;
            return this;
        }

        /**
         * @param accessPackageId The ID of access package this resource association is configured to. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder accessPackageId(String accessPackageId) {
            return accessPackageId(Output.of(accessPackageId));
        }

        /**
         * @param accessType The role of access type to the specified resource. Valid values are `Member`, or `Owner` The default is `Member`. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder accessType(@Nullable Output<String> accessType) {
            $.accessType = accessType;
            return this;
        }

        /**
         * @param accessType The role of access type to the specified resource. Valid values are `Member`, or `Owner` The default is `Member`. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder accessType(String accessType) {
            return accessType(Output.of(accessType));
        }

        /**
         * @param catalogResourceAssociationId The ID of the catalog association from the `azuread.AccessPackageResourceCatalogAssociation` resource. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder catalogResourceAssociationId(Output<String> catalogResourceAssociationId) {
            $.catalogResourceAssociationId = catalogResourceAssociationId;
            return this;
        }

        /**
         * @param catalogResourceAssociationId The ID of the catalog association from the `azuread.AccessPackageResourceCatalogAssociation` resource. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder catalogResourceAssociationId(String catalogResourceAssociationId) {
            return catalogResourceAssociationId(Output.of(catalogResourceAssociationId));
        }

        public AccessPackageResourcePackageAssociationArgs build() {
            if ($.accessPackageId == null) {
                throw new MissingRequiredPropertyException("AccessPackageResourcePackageAssociationArgs", "accessPackageId");
            }
            if ($.catalogResourceAssociationId == null) {
                throw new MissingRequiredPropertyException("AccessPackageResourcePackageAssociationArgs", "catalogResourceAssociationId");
            }
            return $;
        }
    }

}

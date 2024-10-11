// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.azuread.outputs.GetServicePrincipalsServicePrincipal;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetServicePrincipalsResult {
    /**
     * @return The client ID of the application associated with this service principal.
     * 
     */
    private List<String> clientIds;
    /**
     * @return A list of display names of the applications associated with the service principals.
     * 
     */
    private List<String> displayNames;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable Boolean ignoreMissing;
    /**
     * @return The object IDs of the service principals.
     * 
     */
    private List<String> objectIds;
    private @Nullable Boolean returnAll;
    /**
     * @return A list of service principals. Each `service_principal` object provides the attributes documented below.
     * 
     */
    private List<GetServicePrincipalsServicePrincipal> servicePrincipals;

    private GetServicePrincipalsResult() {}
    /**
     * @return The client ID of the application associated with this service principal.
     * 
     */
    public List<String> clientIds() {
        return this.clientIds;
    }
    /**
     * @return A list of display names of the applications associated with the service principals.
     * 
     */
    public List<String> displayNames() {
        return this.displayNames;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<Boolean> ignoreMissing() {
        return Optional.ofNullable(this.ignoreMissing);
    }
    /**
     * @return The object IDs of the service principals.
     * 
     */
    public List<String> objectIds() {
        return this.objectIds;
    }
    public Optional<Boolean> returnAll() {
        return Optional.ofNullable(this.returnAll);
    }
    /**
     * @return A list of service principals. Each `service_principal` object provides the attributes documented below.
     * 
     */
    public List<GetServicePrincipalsServicePrincipal> servicePrincipals() {
        return this.servicePrincipals;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServicePrincipalsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> clientIds;
        private List<String> displayNames;
        private String id;
        private @Nullable Boolean ignoreMissing;
        private List<String> objectIds;
        private @Nullable Boolean returnAll;
        private List<GetServicePrincipalsServicePrincipal> servicePrincipals;
        public Builder() {}
        public Builder(GetServicePrincipalsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clientIds = defaults.clientIds;
    	      this.displayNames = defaults.displayNames;
    	      this.id = defaults.id;
    	      this.ignoreMissing = defaults.ignoreMissing;
    	      this.objectIds = defaults.objectIds;
    	      this.returnAll = defaults.returnAll;
    	      this.servicePrincipals = defaults.servicePrincipals;
        }

        @CustomType.Setter
        public Builder clientIds(List<String> clientIds) {
            if (clientIds == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalsResult", "clientIds");
            }
            this.clientIds = clientIds;
            return this;
        }
        public Builder clientIds(String... clientIds) {
            return clientIds(List.of(clientIds));
        }
        @CustomType.Setter
        public Builder displayNames(List<String> displayNames) {
            if (displayNames == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalsResult", "displayNames");
            }
            this.displayNames = displayNames;
            return this;
        }
        public Builder displayNames(String... displayNames) {
            return displayNames(List.of(displayNames));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ignoreMissing(@Nullable Boolean ignoreMissing) {

            this.ignoreMissing = ignoreMissing;
            return this;
        }
        @CustomType.Setter
        public Builder objectIds(List<String> objectIds) {
            if (objectIds == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalsResult", "objectIds");
            }
            this.objectIds = objectIds;
            return this;
        }
        public Builder objectIds(String... objectIds) {
            return objectIds(List.of(objectIds));
        }
        @CustomType.Setter
        public Builder returnAll(@Nullable Boolean returnAll) {

            this.returnAll = returnAll;
            return this;
        }
        @CustomType.Setter
        public Builder servicePrincipals(List<GetServicePrincipalsServicePrincipal> servicePrincipals) {
            if (servicePrincipals == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalsResult", "servicePrincipals");
            }
            this.servicePrincipals = servicePrincipals;
            return this;
        }
        public Builder servicePrincipals(GetServicePrincipalsServicePrincipal... servicePrincipals) {
            return servicePrincipals(List.of(servicePrincipals));
        }
        public GetServicePrincipalsResult build() {
            final var _resultValue = new GetServicePrincipalsResult();
            _resultValue.clientIds = clientIds;
            _resultValue.displayNames = displayNames;
            _resultValue.id = id;
            _resultValue.ignoreMissing = ignoreMissing;
            _resultValue.objectIds = objectIds;
            _resultValue.returnAll = returnAll;
            _resultValue.servicePrincipals = servicePrincipals;
            return _resultValue;
        }
    }
}

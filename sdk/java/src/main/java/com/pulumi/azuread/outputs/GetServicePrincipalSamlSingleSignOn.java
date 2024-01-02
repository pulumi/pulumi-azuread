// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetServicePrincipalSamlSingleSignOn {
    /**
     * @return The relative URI the service provider would redirect to after completion of the single sign-on flow.
     * 
     */
    private String relayState;

    private GetServicePrincipalSamlSingleSignOn() {}
    /**
     * @return The relative URI the service provider would redirect to after completion of the single sign-on flow.
     * 
     */
    public String relayState() {
        return this.relayState;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServicePrincipalSamlSingleSignOn defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String relayState;
        public Builder() {}
        public Builder(GetServicePrincipalSamlSingleSignOn defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.relayState = defaults.relayState;
        }

        @CustomType.Setter
        public Builder relayState(String relayState) {
            if (relayState == null) {
              throw new MissingRequiredPropertyException("GetServicePrincipalSamlSingleSignOn", "relayState");
            }
            this.relayState = relayState;
            return this;
        }
        public GetServicePrincipalSamlSingleSignOn build() {
            final var _resultValue = new GetServicePrincipalSamlSingleSignOn();
            _resultValue.relayState = relayState;
            return _resultValue;
        }
    }
}

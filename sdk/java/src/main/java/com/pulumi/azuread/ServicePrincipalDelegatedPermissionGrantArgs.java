// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServicePrincipalDelegatedPermissionGrantArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServicePrincipalDelegatedPermissionGrantArgs Empty = new ServicePrincipalDelegatedPermissionGrantArgs();

    /**
     * A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
     * 
     */
    @Import(name="claimValues", required=true)
    private Output<List<String>> claimValues;

    /**
     * @return A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
     * 
     */
    public Output<List<String>> claimValues() {
        return this.claimValues;
    }

    /**
     * The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="resourceServicePrincipalObjectId", required=true)
    private Output<String> resourceServicePrincipalObjectId;

    /**
     * @return The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> resourceServicePrincipalObjectId() {
        return this.resourceServicePrincipalObjectId;
    }

    /**
     * The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="servicePrincipalObjectId", required=true)
    private Output<String> servicePrincipalObjectId;

    /**
     * @return The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> servicePrincipalObjectId() {
        return this.servicePrincipalObjectId;
    }

    /**
     * The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
     * 
     * &gt; **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `user_object_id` property.
     * 
     */
    @Import(name="userObjectId")
    private @Nullable Output<String> userObjectId;

    /**
     * @return The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
     * 
     * &gt; **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `user_object_id` property.
     * 
     */
    public Optional<Output<String>> userObjectId() {
        return Optional.ofNullable(this.userObjectId);
    }

    private ServicePrincipalDelegatedPermissionGrantArgs() {}

    private ServicePrincipalDelegatedPermissionGrantArgs(ServicePrincipalDelegatedPermissionGrantArgs $) {
        this.claimValues = $.claimValues;
        this.resourceServicePrincipalObjectId = $.resourceServicePrincipalObjectId;
        this.servicePrincipalObjectId = $.servicePrincipalObjectId;
        this.userObjectId = $.userObjectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServicePrincipalDelegatedPermissionGrantArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServicePrincipalDelegatedPermissionGrantArgs $;

        public Builder() {
            $ = new ServicePrincipalDelegatedPermissionGrantArgs();
        }

        public Builder(ServicePrincipalDelegatedPermissionGrantArgs defaults) {
            $ = new ServicePrincipalDelegatedPermissionGrantArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param claimValues A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
         * 
         * @return builder
         * 
         */
        public Builder claimValues(Output<List<String>> claimValues) {
            $.claimValues = claimValues;
            return this;
        }

        /**
         * @param claimValues A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
         * 
         * @return builder
         * 
         */
        public Builder claimValues(List<String> claimValues) {
            return claimValues(Output.of(claimValues));
        }

        /**
         * @param claimValues A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
         * 
         * @return builder
         * 
         */
        public Builder claimValues(String... claimValues) {
            return claimValues(List.of(claimValues));
        }

        /**
         * @param resourceServicePrincipalObjectId The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder resourceServicePrincipalObjectId(Output<String> resourceServicePrincipalObjectId) {
            $.resourceServicePrincipalObjectId = resourceServicePrincipalObjectId;
            return this;
        }

        /**
         * @param resourceServicePrincipalObjectId The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder resourceServicePrincipalObjectId(String resourceServicePrincipalObjectId) {
            return resourceServicePrincipalObjectId(Output.of(resourceServicePrincipalObjectId));
        }

        /**
         * @param servicePrincipalObjectId The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder servicePrincipalObjectId(Output<String> servicePrincipalObjectId) {
            $.servicePrincipalObjectId = servicePrincipalObjectId;
            return this;
        }

        /**
         * @param servicePrincipalObjectId The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder servicePrincipalObjectId(String servicePrincipalObjectId) {
            return servicePrincipalObjectId(Output.of(servicePrincipalObjectId));
        }

        /**
         * @param userObjectId The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
         * 
         * &gt; **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `user_object_id` property.
         * 
         * @return builder
         * 
         */
        public Builder userObjectId(@Nullable Output<String> userObjectId) {
            $.userObjectId = userObjectId;
            return this;
        }

        /**
         * @param userObjectId The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
         * 
         * &gt; **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `user_object_id` property.
         * 
         * @return builder
         * 
         */
        public Builder userObjectId(String userObjectId) {
            return userObjectId(Output.of(userObjectId));
        }

        public ServicePrincipalDelegatedPermissionGrantArgs build() {
            if ($.claimValues == null) {
                throw new MissingRequiredPropertyException("ServicePrincipalDelegatedPermissionGrantArgs", "claimValues");
            }
            if ($.resourceServicePrincipalObjectId == null) {
                throw new MissingRequiredPropertyException("ServicePrincipalDelegatedPermissionGrantArgs", "resourceServicePrincipalObjectId");
            }
            if ($.servicePrincipalObjectId == null) {
                throw new MissingRequiredPropertyException("ServicePrincipalDelegatedPermissionGrantArgs", "servicePrincipalObjectId");
            }
            return $;
        }
    }

}

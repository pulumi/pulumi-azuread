// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SynchronizationJobArgs extends com.pulumi.resources.ResourceArgs {

    public static final SynchronizationJobArgs Empty = new SynchronizationJobArgs();

    /**
     * Whether the provisioning job is enabled. Default state is `true`.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Whether the provisioning job is enabled. Default state is `true`.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * The ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
     * 
     */
    @Import(name="servicePrincipalId", required=true)
    private Output<String> servicePrincipalId;

    /**
     * @return The ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
     * 
     */
    public Output<String> servicePrincipalId() {
        return this.servicePrincipalId;
    }

    /**
     * Identifier of the synchronization template this job is based on.
     * 
     */
    @Import(name="templateId", required=true)
    private Output<String> templateId;

    /**
     * @return Identifier of the synchronization template this job is based on.
     * 
     */
    public Output<String> templateId() {
        return this.templateId;
    }

    private SynchronizationJobArgs() {}

    private SynchronizationJobArgs(SynchronizationJobArgs $) {
        this.enabled = $.enabled;
        this.servicePrincipalId = $.servicePrincipalId;
        this.templateId = $.templateId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SynchronizationJobArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SynchronizationJobArgs $;

        public Builder() {
            $ = new SynchronizationJobArgs();
        }

        public Builder(SynchronizationJobArgs defaults) {
            $ = new SynchronizationJobArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enabled Whether the provisioning job is enabled. Default state is `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Whether the provisioning job is enabled. Default state is `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param servicePrincipalId The ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder servicePrincipalId(Output<String> servicePrincipalId) {
            $.servicePrincipalId = servicePrincipalId;
            return this;
        }

        /**
         * @param servicePrincipalId The ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder servicePrincipalId(String servicePrincipalId) {
            return servicePrincipalId(Output.of(servicePrincipalId));
        }

        /**
         * @param templateId Identifier of the synchronization template this job is based on.
         * 
         * @return builder
         * 
         */
        public Builder templateId(Output<String> templateId) {
            $.templateId = templateId;
            return this;
        }

        /**
         * @param templateId Identifier of the synchronization template this job is based on.
         * 
         * @return builder
         * 
         */
        public Builder templateId(String templateId) {
            return templateId(Output.of(templateId));
        }

        public SynchronizationJobArgs build() {
            if ($.servicePrincipalId == null) {
                throw new MissingRequiredPropertyException("SynchronizationJobArgs", "servicePrincipalId");
            }
            if ($.templateId == null) {
                throw new MissingRequiredPropertyException("SynchronizationJobArgs", "templateId");
            }
            return $;
        }
    }

}

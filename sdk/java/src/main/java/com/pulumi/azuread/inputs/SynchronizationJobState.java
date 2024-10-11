// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.azuread.inputs.SynchronizationJobScheduleArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SynchronizationJobState extends com.pulumi.resources.ResourceArgs {

    public static final SynchronizationJobState Empty = new SynchronizationJobState();

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
     * A `schedule` list as documented below.
     * 
     */
    @Import(name="schedules")
    private @Nullable Output<List<SynchronizationJobScheduleArgs>> schedules;

    /**
     * @return A `schedule` list as documented below.
     * 
     */
    public Optional<Output<List<SynchronizationJobScheduleArgs>>> schedules() {
        return Optional.ofNullable(this.schedules);
    }

    /**
     * The ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
     * 
     */
    @Import(name="servicePrincipalId")
    private @Nullable Output<String> servicePrincipalId;

    /**
     * @return The ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> servicePrincipalId() {
        return Optional.ofNullable(this.servicePrincipalId);
    }

    /**
     * Identifier of the synchronization template this job is based on.
     * 
     */
    @Import(name="templateId")
    private @Nullable Output<String> templateId;

    /**
     * @return Identifier of the synchronization template this job is based on.
     * 
     */
    public Optional<Output<String>> templateId() {
        return Optional.ofNullable(this.templateId);
    }

    private SynchronizationJobState() {}

    private SynchronizationJobState(SynchronizationJobState $) {
        this.enabled = $.enabled;
        this.schedules = $.schedules;
        this.servicePrincipalId = $.servicePrincipalId;
        this.templateId = $.templateId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SynchronizationJobState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SynchronizationJobState $;

        public Builder() {
            $ = new SynchronizationJobState();
        }

        public Builder(SynchronizationJobState defaults) {
            $ = new SynchronizationJobState(Objects.requireNonNull(defaults));
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
         * @param schedules A `schedule` list as documented below.
         * 
         * @return builder
         * 
         */
        public Builder schedules(@Nullable Output<List<SynchronizationJobScheduleArgs>> schedules) {
            $.schedules = schedules;
            return this;
        }

        /**
         * @param schedules A `schedule` list as documented below.
         * 
         * @return builder
         * 
         */
        public Builder schedules(List<SynchronizationJobScheduleArgs> schedules) {
            return schedules(Output.of(schedules));
        }

        /**
         * @param schedules A `schedule` list as documented below.
         * 
         * @return builder
         * 
         */
        public Builder schedules(SynchronizationJobScheduleArgs... schedules) {
            return schedules(List.of(schedules));
        }

        /**
         * @param servicePrincipalId The ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder servicePrincipalId(@Nullable Output<String> servicePrincipalId) {
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
        public Builder templateId(@Nullable Output<String> templateId) {
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

        public SynchronizationJobState build() {
            return $;
        }
    }

}

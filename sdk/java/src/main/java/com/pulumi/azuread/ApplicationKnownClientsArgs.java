// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class ApplicationKnownClientsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ApplicationKnownClientsArgs Empty = new ApplicationKnownClientsArgs();

    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="applicationId", required=true)
    private Output<String> applicationId;

    /**
     * @return The resource ID of the application registration. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> applicationId() {
        return this.applicationId;
    }

    /**
     * A set of client IDs for the known applications.
     * 
     */
    @Import(name="knownClientIds", required=true)
    private Output<List<String>> knownClientIds;

    /**
     * @return A set of client IDs for the known applications.
     * 
     */
    public Output<List<String>> knownClientIds() {
        return this.knownClientIds;
    }

    private ApplicationKnownClientsArgs() {}

    private ApplicationKnownClientsArgs(ApplicationKnownClientsArgs $) {
        this.applicationId = $.applicationId;
        this.knownClientIds = $.knownClientIds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ApplicationKnownClientsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ApplicationKnownClientsArgs $;

        public Builder() {
            $ = new ApplicationKnownClientsArgs();
        }

        public Builder(ApplicationKnownClientsArgs defaults) {
            $ = new ApplicationKnownClientsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param applicationId The resource ID of the application registration. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder applicationId(Output<String> applicationId) {
            $.applicationId = applicationId;
            return this;
        }

        /**
         * @param applicationId The resource ID of the application registration. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder applicationId(String applicationId) {
            return applicationId(Output.of(applicationId));
        }

        /**
         * @param knownClientIds A set of client IDs for the known applications.
         * 
         * @return builder
         * 
         */
        public Builder knownClientIds(Output<List<String>> knownClientIds) {
            $.knownClientIds = knownClientIds;
            return this;
        }

        /**
         * @param knownClientIds A set of client IDs for the known applications.
         * 
         * @return builder
         * 
         */
        public Builder knownClientIds(List<String> knownClientIds) {
            return knownClientIds(Output.of(knownClientIds));
        }

        /**
         * @param knownClientIds A set of client IDs for the known applications.
         * 
         * @return builder
         * 
         */
        public Builder knownClientIds(String... knownClientIds) {
            return knownClientIds(List.of(knownClientIds));
        }

        public ApplicationKnownClientsArgs build() {
            $.applicationId = Objects.requireNonNull($.applicationId, "expected parameter 'applicationId' to be non-null");
            $.knownClientIds = Objects.requireNonNull($.knownClientIds, "expected parameter 'knownClientIds' to be non-null");
            return $;
        }
    }

}

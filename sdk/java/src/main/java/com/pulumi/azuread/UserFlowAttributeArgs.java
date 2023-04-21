// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class UserFlowAttributeArgs extends com.pulumi.resources.ResourceArgs {

    public static final UserFlowAttributeArgs Empty = new UserFlowAttributeArgs();

    /**
     * The data type of the user flow attribute. Possible values are `boolean`, `dateTime`, `int64`, `string` or `stringCollection`. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="dataType", required=true)
    private Output<String> dataType;

    /**
     * @return The data type of the user flow attribute. Possible values are `boolean`, `dateTime`, `int64`, `string` or `stringCollection`. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> dataType() {
        return this.dataType;
    }

    /**
     * The description of the user flow attribute that is shown to the user at the time of sign-up.
     * 
     */
    @Import(name="description", required=true)
    private Output<String> description;

    /**
     * @return The description of the user flow attribute that is shown to the user at the time of sign-up.
     * 
     */
    public Output<String> description() {
        return this.description;
    }

    /**
     * The display name of the user flow attribute. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="displayName", required=true)
    private Output<String> displayName;

    /**
     * @return The display name of the user flow attribute. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }

    private UserFlowAttributeArgs() {}

    private UserFlowAttributeArgs(UserFlowAttributeArgs $) {
        this.dataType = $.dataType;
        this.description = $.description;
        this.displayName = $.displayName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserFlowAttributeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserFlowAttributeArgs $;

        public Builder() {
            $ = new UserFlowAttributeArgs();
        }

        public Builder(UserFlowAttributeArgs defaults) {
            $ = new UserFlowAttributeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dataType The data type of the user flow attribute. Possible values are `boolean`, `dateTime`, `int64`, `string` or `stringCollection`. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder dataType(Output<String> dataType) {
            $.dataType = dataType;
            return this;
        }

        /**
         * @param dataType The data type of the user flow attribute. Possible values are `boolean`, `dateTime`, `int64`, `string` or `stringCollection`. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder dataType(String dataType) {
            return dataType(Output.of(dataType));
        }

        /**
         * @param description The description of the user flow attribute that is shown to the user at the time of sign-up.
         * 
         * @return builder
         * 
         */
        public Builder description(Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the user flow attribute that is shown to the user at the time of sign-up.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param displayName The display name of the user flow attribute. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder displayName(Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName The display name of the user flow attribute. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        public UserFlowAttributeArgs build() {
            $.dataType = Objects.requireNonNull($.dataType, "expected parameter 'dataType' to be non-null");
            $.description = Objects.requireNonNull($.description, "expected parameter 'description' to be non-null");
            $.displayName = Objects.requireNonNull($.displayName, "expected parameter 'displayName' to be non-null");
            return $;
        }
    }

}
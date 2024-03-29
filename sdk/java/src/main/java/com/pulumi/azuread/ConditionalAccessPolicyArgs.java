// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.inputs.ConditionalAccessPolicyConditionsArgs;
import com.pulumi.azuread.inputs.ConditionalAccessPolicyGrantControlsArgs;
import com.pulumi.azuread.inputs.ConditionalAccessPolicySessionControlsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConditionalAccessPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConditionalAccessPolicyArgs Empty = new ConditionalAccessPolicyArgs();

    /**
     * A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
     * 
     */
    @Import(name="conditions", required=true)
    private Output<ConditionalAccessPolicyConditionsArgs> conditions;

    /**
     * @return A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
     * 
     */
    public Output<ConditionalAccessPolicyConditionsArgs> conditions() {
        return this.conditions;
    }

    /**
     * The friendly name for this Conditional Access Policy.
     * 
     */
    @Import(name="displayName", required=true)
    private Output<String> displayName;

    /**
     * @return The friendly name for this Conditional Access Policy.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }

    /**
     * A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
     * 
     */
    @Import(name="grantControls")
    private @Nullable Output<ConditionalAccessPolicyGrantControlsArgs> grantControls;

    /**
     * @return A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
     * 
     */
    public Optional<Output<ConditionalAccessPolicyGrantControlsArgs>> grantControls() {
        return Optional.ofNullable(this.grantControls);
    }

    /**
     * A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.
     * 
     * &gt; Note: At least one of `grant_controls` and/or `session_controls` blocks must be specified.
     * 
     */
    @Import(name="sessionControls")
    private @Nullable Output<ConditionalAccessPolicySessionControlsArgs> sessionControls;

    /**
     * @return A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.
     * 
     * &gt; Note: At least one of `grant_controls` and/or `session_controls` blocks must be specified.
     * 
     */
    public Optional<Output<ConditionalAccessPolicySessionControlsArgs>> sessionControls() {
        return Optional.ofNullable(this.sessionControls);
    }

    /**
     * Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
     * 
     */
    @Import(name="state", required=true)
    private Output<String> state;

    /**
     * @return Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
     * 
     */
    public Output<String> state() {
        return this.state;
    }

    private ConditionalAccessPolicyArgs() {}

    private ConditionalAccessPolicyArgs(ConditionalAccessPolicyArgs $) {
        this.conditions = $.conditions;
        this.displayName = $.displayName;
        this.grantControls = $.grantControls;
        this.sessionControls = $.sessionControls;
        this.state = $.state;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConditionalAccessPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConditionalAccessPolicyArgs $;

        public Builder() {
            $ = new ConditionalAccessPolicyArgs();
        }

        public Builder(ConditionalAccessPolicyArgs defaults) {
            $ = new ConditionalAccessPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param conditions A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
         * 
         * @return builder
         * 
         */
        public Builder conditions(Output<ConditionalAccessPolicyConditionsArgs> conditions) {
            $.conditions = conditions;
            return this;
        }

        /**
         * @param conditions A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
         * 
         * @return builder
         * 
         */
        public Builder conditions(ConditionalAccessPolicyConditionsArgs conditions) {
            return conditions(Output.of(conditions));
        }

        /**
         * @param displayName The friendly name for this Conditional Access Policy.
         * 
         * @return builder
         * 
         */
        public Builder displayName(Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName The friendly name for this Conditional Access Policy.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param grantControls A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
         * 
         * @return builder
         * 
         */
        public Builder grantControls(@Nullable Output<ConditionalAccessPolicyGrantControlsArgs> grantControls) {
            $.grantControls = grantControls;
            return this;
        }

        /**
         * @param grantControls A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
         * 
         * @return builder
         * 
         */
        public Builder grantControls(ConditionalAccessPolicyGrantControlsArgs grantControls) {
            return grantControls(Output.of(grantControls));
        }

        /**
         * @param sessionControls A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.
         * 
         * &gt; Note: At least one of `grant_controls` and/or `session_controls` blocks must be specified.
         * 
         * @return builder
         * 
         */
        public Builder sessionControls(@Nullable Output<ConditionalAccessPolicySessionControlsArgs> sessionControls) {
            $.sessionControls = sessionControls;
            return this;
        }

        /**
         * @param sessionControls A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.
         * 
         * &gt; Note: At least one of `grant_controls` and/or `session_controls` blocks must be specified.
         * 
         * @return builder
         * 
         */
        public Builder sessionControls(ConditionalAccessPolicySessionControlsArgs sessionControls) {
            return sessionControls(Output.of(sessionControls));
        }

        /**
         * @param state Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
         * 
         * @return builder
         * 
         */
        public Builder state(Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        public ConditionalAccessPolicyArgs build() {
            if ($.conditions == null) {
                throw new MissingRequiredPropertyException("ConditionalAccessPolicyArgs", "conditions");
            }
            if ($.displayName == null) {
                throw new MissingRequiredPropertyException("ConditionalAccessPolicyArgs", "displayName");
            }
            if ($.state == null) {
                throw new MissingRequiredPropertyException("ConditionalAccessPolicyArgs", "state");
            }
            return $;
        }
    }

}

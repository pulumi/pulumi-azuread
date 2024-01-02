// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUserExternalTenantArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUserExternalTenantArgs Empty = new ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUserExternalTenantArgs();

    /**
     * A list tenant IDs. Can only be specified if `membership_kind` is `enumerated`.
     * 
     */
    @Import(name="members")
    private @Nullable Output<List<String>> members;

    /**
     * @return A list tenant IDs. Can only be specified if `membership_kind` is `enumerated`.
     * 
     */
    public Optional<Output<List<String>>> members() {
        return Optional.ofNullable(this.members);
    }

    /**
     * The external tenant membership kind. Possible values are: `all`, `enumerated`, `unknownFutureValue`.
     * 
     */
    @Import(name="membershipKind", required=true)
    private Output<String> membershipKind;

    /**
     * @return The external tenant membership kind. Possible values are: `all`, `enumerated`, `unknownFutureValue`.
     * 
     */
    public Output<String> membershipKind() {
        return this.membershipKind;
    }

    private ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUserExternalTenantArgs() {}

    private ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUserExternalTenantArgs(ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUserExternalTenantArgs $) {
        this.members = $.members;
        this.membershipKind = $.membershipKind;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUserExternalTenantArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUserExternalTenantArgs $;

        public Builder() {
            $ = new ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUserExternalTenantArgs();
        }

        public Builder(ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUserExternalTenantArgs defaults) {
            $ = new ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUserExternalTenantArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param members A list tenant IDs. Can only be specified if `membership_kind` is `enumerated`.
         * 
         * @return builder
         * 
         */
        public Builder members(@Nullable Output<List<String>> members) {
            $.members = members;
            return this;
        }

        /**
         * @param members A list tenant IDs. Can only be specified if `membership_kind` is `enumerated`.
         * 
         * @return builder
         * 
         */
        public Builder members(List<String> members) {
            return members(Output.of(members));
        }

        /**
         * @param members A list tenant IDs. Can only be specified if `membership_kind` is `enumerated`.
         * 
         * @return builder
         * 
         */
        public Builder members(String... members) {
            return members(List.of(members));
        }

        /**
         * @param membershipKind The external tenant membership kind. Possible values are: `all`, `enumerated`, `unknownFutureValue`.
         * 
         * @return builder
         * 
         */
        public Builder membershipKind(Output<String> membershipKind) {
            $.membershipKind = membershipKind;
            return this;
        }

        /**
         * @param membershipKind The external tenant membership kind. Possible values are: `all`, `enumerated`, `unknownFutureValue`.
         * 
         * @return builder
         * 
         */
        public Builder membershipKind(String membershipKind) {
            return membershipKind(Output.of(membershipKind));
        }

        public ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUserExternalTenantArgs build() {
            if ($.membershipKind == null) {
                throw new MissingRequiredPropertyException("ConditionalAccessPolicyConditionsUsersExcludedGuestsOrExternalUserExternalTenantArgs", "membershipKind");
            }
            return $;
        }
    }

}

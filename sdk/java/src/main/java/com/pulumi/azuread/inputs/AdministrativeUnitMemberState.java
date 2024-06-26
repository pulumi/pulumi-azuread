// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AdministrativeUnitMemberState extends com.pulumi.resources.ResourceArgs {

    public static final AdministrativeUnitMemberState Empty = new AdministrativeUnitMemberState();

    /**
     * The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="administrativeUnitObjectId")
    private @Nullable Output<String> administrativeUnitObjectId;

    /**
     * @return The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> administrativeUnitObjectId() {
        return Optional.ofNullable(this.administrativeUnitObjectId);
    }

    /**
     * The object ID of the user or group you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
     * 
     * &gt; **Caution** When using the azuread.AdministrativeUnitMember resource to manage Administrative Unit membership for a group, you will need to use an `ignore_changes = [administrative_unit_ids]` lifecycle meta argument for the `azuread.Group` resource, in order to avoid a persistent diff.
     * 
     */
    @Import(name="memberObjectId")
    private @Nullable Output<String> memberObjectId;

    /**
     * @return The object ID of the user or group you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
     * 
     * &gt; **Caution** When using the azuread.AdministrativeUnitMember resource to manage Administrative Unit membership for a group, you will need to use an `ignore_changes = [administrative_unit_ids]` lifecycle meta argument for the `azuread.Group` resource, in order to avoid a persistent diff.
     * 
     */
    public Optional<Output<String>> memberObjectId() {
        return Optional.ofNullable(this.memberObjectId);
    }

    private AdministrativeUnitMemberState() {}

    private AdministrativeUnitMemberState(AdministrativeUnitMemberState $) {
        this.administrativeUnitObjectId = $.administrativeUnitObjectId;
        this.memberObjectId = $.memberObjectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AdministrativeUnitMemberState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AdministrativeUnitMemberState $;

        public Builder() {
            $ = new AdministrativeUnitMemberState();
        }

        public Builder(AdministrativeUnitMemberState defaults) {
            $ = new AdministrativeUnitMemberState(Objects.requireNonNull(defaults));
        }

        /**
         * @param administrativeUnitObjectId The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder administrativeUnitObjectId(@Nullable Output<String> administrativeUnitObjectId) {
            $.administrativeUnitObjectId = administrativeUnitObjectId;
            return this;
        }

        /**
         * @param administrativeUnitObjectId The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder administrativeUnitObjectId(String administrativeUnitObjectId) {
            return administrativeUnitObjectId(Output.of(administrativeUnitObjectId));
        }

        /**
         * @param memberObjectId The object ID of the user or group you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
         * 
         * &gt; **Caution** When using the azuread.AdministrativeUnitMember resource to manage Administrative Unit membership for a group, you will need to use an `ignore_changes = [administrative_unit_ids]` lifecycle meta argument for the `azuread.Group` resource, in order to avoid a persistent diff.
         * 
         * @return builder
         * 
         */
        public Builder memberObjectId(@Nullable Output<String> memberObjectId) {
            $.memberObjectId = memberObjectId;
            return this;
        }

        /**
         * @param memberObjectId The object ID of the user or group you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
         * 
         * &gt; **Caution** When using the azuread.AdministrativeUnitMember resource to manage Administrative Unit membership for a group, you will need to use an `ignore_changes = [administrative_unit_ids]` lifecycle meta argument for the `azuread.Group` resource, in order to avoid a persistent diff.
         * 
         * @return builder
         * 
         */
        public Builder memberObjectId(String memberObjectId) {
            return memberObjectId(Output.of(memberObjectId));
        }

        public AdministrativeUnitMemberState build() {
            return $;
        }
    }

}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DirectoryRoleMemberArgs extends com.pulumi.resources.ResourceArgs {

    public static final DirectoryRoleMemberArgs Empty = new DirectoryRoleMemberArgs();

    /**
     * The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="memberObjectId")
    private @Nullable Output<String> memberObjectId;

    /**
     * @return The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> memberObjectId() {
        return Optional.ofNullable(this.memberObjectId);
    }

    /**
     * The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="roleObjectId")
    private @Nullable Output<String> roleObjectId;

    /**
     * @return The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> roleObjectId() {
        return Optional.ofNullable(this.roleObjectId);
    }

    private DirectoryRoleMemberArgs() {}

    private DirectoryRoleMemberArgs(DirectoryRoleMemberArgs $) {
        this.memberObjectId = $.memberObjectId;
        this.roleObjectId = $.roleObjectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DirectoryRoleMemberArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DirectoryRoleMemberArgs $;

        public Builder() {
            $ = new DirectoryRoleMemberArgs();
        }

        public Builder(DirectoryRoleMemberArgs defaults) {
            $ = new DirectoryRoleMemberArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param memberObjectId The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder memberObjectId(@Nullable Output<String> memberObjectId) {
            $.memberObjectId = memberObjectId;
            return this;
        }

        /**
         * @param memberObjectId The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder memberObjectId(String memberObjectId) {
            return memberObjectId(Output.of(memberObjectId));
        }

        /**
         * @param roleObjectId The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder roleObjectId(@Nullable Output<String> roleObjectId) {
            $.roleObjectId = roleObjectId;
            return this;
        }

        /**
         * @param roleObjectId The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder roleObjectId(String roleObjectId) {
            return roleObjectId(Output.of(roleObjectId));
        }

        public DirectoryRoleMemberArgs build() {
            return $;
        }
    }

}

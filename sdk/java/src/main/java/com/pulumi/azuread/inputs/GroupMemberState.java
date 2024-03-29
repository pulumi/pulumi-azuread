// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupMemberState extends com.pulumi.resources.ResourceArgs {

    public static final GroupMemberState Empty = new GroupMemberState();

    /**
     * The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="groupObjectId")
    private @Nullable Output<String> groupObjectId;

    /**
     * @return The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> groupObjectId() {
        return Optional.ofNullable(this.groupObjectId);
    }

    /**
     * The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="memberObjectId")
    private @Nullable Output<String> memberObjectId;

    /**
     * @return The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> memberObjectId() {
        return Optional.ofNullable(this.memberObjectId);
    }

    private GroupMemberState() {}

    private GroupMemberState(GroupMemberState $) {
        this.groupObjectId = $.groupObjectId;
        this.memberObjectId = $.memberObjectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupMemberState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupMemberState $;

        public Builder() {
            $ = new GroupMemberState();
        }

        public Builder(GroupMemberState defaults) {
            $ = new GroupMemberState(Objects.requireNonNull(defaults));
        }

        /**
         * @param groupObjectId The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder groupObjectId(@Nullable Output<String> groupObjectId) {
            $.groupObjectId = groupObjectId;
            return this;
        }

        /**
         * @param groupObjectId The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder groupObjectId(String groupObjectId) {
            return groupObjectId(Output.of(groupObjectId));
        }

        /**
         * @param memberObjectId The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder memberObjectId(@Nullable Output<String> memberObjectId) {
            $.memberObjectId = memberObjectId;
            return this;
        }

        /**
         * @param memberObjectId The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder memberObjectId(String memberObjectId) {
            return memberObjectId(Output.of(memberObjectId));
        }

        public GroupMemberState build() {
            return $;
        }
    }

}

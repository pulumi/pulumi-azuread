// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.azuread.outputs.GetUsersUser;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetUsersResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    private final @Nullable Boolean ignoreMissing;
    /**
     * @return The email aliases of the users.
     * 
     */
    private final List<String> mailNicknames;
    /**
     * @return The object IDs of the users.
     * 
     */
    private final List<String> objectIds;
    private final @Nullable Boolean returnAll;
    /**
     * @return The user principal names (UPNs) of the users.
     * 
     */
    private final List<String> userPrincipalNames;
    /**
     * @return A list of users. Each `user` object provides the attributes documented below.
     * 
     */
    private final List<GetUsersUser> users;

    @CustomType.Constructor
    private GetUsersResult(
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("ignoreMissing") @Nullable Boolean ignoreMissing,
        @CustomType.Parameter("mailNicknames") List<String> mailNicknames,
        @CustomType.Parameter("objectIds") List<String> objectIds,
        @CustomType.Parameter("returnAll") @Nullable Boolean returnAll,
        @CustomType.Parameter("userPrincipalNames") List<String> userPrincipalNames,
        @CustomType.Parameter("users") List<GetUsersUser> users) {
        this.id = id;
        this.ignoreMissing = ignoreMissing;
        this.mailNicknames = mailNicknames;
        this.objectIds = objectIds;
        this.returnAll = returnAll;
        this.userPrincipalNames = userPrincipalNames;
        this.users = users;
    }

    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<Boolean> ignoreMissing() {
        return Optional.ofNullable(this.ignoreMissing);
    }
    /**
     * @return The email aliases of the users.
     * 
     */
    public List<String> mailNicknames() {
        return this.mailNicknames;
    }
    /**
     * @return The object IDs of the users.
     * 
     */
    public List<String> objectIds() {
        return this.objectIds;
    }
    public Optional<Boolean> returnAll() {
        return Optional.ofNullable(this.returnAll);
    }
    /**
     * @return The user principal names (UPNs) of the users.
     * 
     */
    public List<String> userPrincipalNames() {
        return this.userPrincipalNames;
    }
    /**
     * @return A list of users. Each `user` object provides the attributes documented below.
     * 
     */
    public List<GetUsersUser> users() {
        return this.users;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetUsersResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String id;
        private @Nullable Boolean ignoreMissing;
        private List<String> mailNicknames;
        private List<String> objectIds;
        private @Nullable Boolean returnAll;
        private List<String> userPrincipalNames;
        private List<GetUsersUser> users;

        public Builder() {
    	      // Empty
        }

        public Builder(GetUsersResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ignoreMissing = defaults.ignoreMissing;
    	      this.mailNicknames = defaults.mailNicknames;
    	      this.objectIds = defaults.objectIds;
    	      this.returnAll = defaults.returnAll;
    	      this.userPrincipalNames = defaults.userPrincipalNames;
    	      this.users = defaults.users;
        }

        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder ignoreMissing(@Nullable Boolean ignoreMissing) {
            this.ignoreMissing = ignoreMissing;
            return this;
        }
        public Builder mailNicknames(List<String> mailNicknames) {
            this.mailNicknames = Objects.requireNonNull(mailNicknames);
            return this;
        }
        public Builder mailNicknames(String... mailNicknames) {
            return mailNicknames(List.of(mailNicknames));
        }
        public Builder objectIds(List<String> objectIds) {
            this.objectIds = Objects.requireNonNull(objectIds);
            return this;
        }
        public Builder objectIds(String... objectIds) {
            return objectIds(List.of(objectIds));
        }
        public Builder returnAll(@Nullable Boolean returnAll) {
            this.returnAll = returnAll;
            return this;
        }
        public Builder userPrincipalNames(List<String> userPrincipalNames) {
            this.userPrincipalNames = Objects.requireNonNull(userPrincipalNames);
            return this;
        }
        public Builder userPrincipalNames(String... userPrincipalNames) {
            return userPrincipalNames(List.of(userPrincipalNames));
        }
        public Builder users(List<GetUsersUser> users) {
            this.users = Objects.requireNonNull(users);
            return this;
        }
        public Builder users(GetUsersUser... users) {
            return users(List.of(users));
        }        public GetUsersResult build() {
            return new GetUsersResult(id, ignoreMissing, mailNicknames, objectIds, returnAll, userPrincipalNames, users);
        }
    }
}

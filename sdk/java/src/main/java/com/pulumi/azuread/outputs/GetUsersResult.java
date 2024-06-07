// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.azuread.outputs.GetUsersUser;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetUsersResult {
    /**
     * @return The employee identifiers assigned to the users by the organisation.
     * 
     */
    private List<String> employeeIds;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable Boolean ignoreMissing;
    /**
     * @return The email aliases of the users.
     * 
     */
    private List<String> mailNicknames;
    /**
     * @return The SMTP email addresses of the users.
     * 
     */
    private List<String> mails;
    /**
     * @return The object IDs of the users.
     * 
     */
    private List<String> objectIds;
    private @Nullable Boolean returnAll;
    /**
     * @return The user principal names (UPNs) of the users.
     * 
     */
    private List<String> userPrincipalNames;
    /**
     * @return A list of users. Each `user` object provides the attributes documented below.
     * 
     */
    private List<GetUsersUser> users;

    private GetUsersResult() {}
    /**
     * @return The employee identifiers assigned to the users by the organisation.
     * 
     */
    public List<String> employeeIds() {
        return this.employeeIds;
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
     * @return The SMTP email addresses of the users.
     * 
     */
    public List<String> mails() {
        return this.mails;
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
    @CustomType.Builder
    public static final class Builder {
        private List<String> employeeIds;
        private String id;
        private @Nullable Boolean ignoreMissing;
        private List<String> mailNicknames;
        private List<String> mails;
        private List<String> objectIds;
        private @Nullable Boolean returnAll;
        private List<String> userPrincipalNames;
        private List<GetUsersUser> users;
        public Builder() {}
        public Builder(GetUsersResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.employeeIds = defaults.employeeIds;
    	      this.id = defaults.id;
    	      this.ignoreMissing = defaults.ignoreMissing;
    	      this.mailNicknames = defaults.mailNicknames;
    	      this.mails = defaults.mails;
    	      this.objectIds = defaults.objectIds;
    	      this.returnAll = defaults.returnAll;
    	      this.userPrincipalNames = defaults.userPrincipalNames;
    	      this.users = defaults.users;
        }

        @CustomType.Setter
        public Builder employeeIds(List<String> employeeIds) {
            if (employeeIds == null) {
              throw new MissingRequiredPropertyException("GetUsersResult", "employeeIds");
            }
            this.employeeIds = employeeIds;
            return this;
        }
        public Builder employeeIds(String... employeeIds) {
            return employeeIds(List.of(employeeIds));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetUsersResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ignoreMissing(@Nullable Boolean ignoreMissing) {

            this.ignoreMissing = ignoreMissing;
            return this;
        }
        @CustomType.Setter
        public Builder mailNicknames(List<String> mailNicknames) {
            if (mailNicknames == null) {
              throw new MissingRequiredPropertyException("GetUsersResult", "mailNicknames");
            }
            this.mailNicknames = mailNicknames;
            return this;
        }
        public Builder mailNicknames(String... mailNicknames) {
            return mailNicknames(List.of(mailNicknames));
        }
        @CustomType.Setter
        public Builder mails(List<String> mails) {
            if (mails == null) {
              throw new MissingRequiredPropertyException("GetUsersResult", "mails");
            }
            this.mails = mails;
            return this;
        }
        public Builder mails(String... mails) {
            return mails(List.of(mails));
        }
        @CustomType.Setter
        public Builder objectIds(List<String> objectIds) {
            if (objectIds == null) {
              throw new MissingRequiredPropertyException("GetUsersResult", "objectIds");
            }
            this.objectIds = objectIds;
            return this;
        }
        public Builder objectIds(String... objectIds) {
            return objectIds(List.of(objectIds));
        }
        @CustomType.Setter
        public Builder returnAll(@Nullable Boolean returnAll) {

            this.returnAll = returnAll;
            return this;
        }
        @CustomType.Setter
        public Builder userPrincipalNames(List<String> userPrincipalNames) {
            if (userPrincipalNames == null) {
              throw new MissingRequiredPropertyException("GetUsersResult", "userPrincipalNames");
            }
            this.userPrincipalNames = userPrincipalNames;
            return this;
        }
        public Builder userPrincipalNames(String... userPrincipalNames) {
            return userPrincipalNames(List.of(userPrincipalNames));
        }
        @CustomType.Setter
        public Builder users(List<GetUsersUser> users) {
            if (users == null) {
              throw new MissingRequiredPropertyException("GetUsersResult", "users");
            }
            this.users = users;
            return this;
        }
        public Builder users(GetUsersUser... users) {
            return users(List.of(users));
        }
        public GetUsersResult build() {
            final var _resultValue = new GetUsersResult();
            _resultValue.employeeIds = employeeIds;
            _resultValue.id = id;
            _resultValue.ignoreMissing = ignoreMissing;
            _resultValue.mailNicknames = mailNicknames;
            _resultValue.mails = mails;
            _resultValue.objectIds = objectIds;
            _resultValue.returnAll = returnAll;
            _resultValue.userPrincipalNames = userPrincipalNames;
            _resultValue.users = users;
            return _resultValue;
        }
    }
}

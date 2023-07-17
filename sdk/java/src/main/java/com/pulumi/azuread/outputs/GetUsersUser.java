// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetUsersUser {
    /**
     * @return Whether or not the account is enabled.
     * 
     */
    private Boolean accountEnabled;
    /**
     * @return The display name of the user.
     * 
     */
    private String displayName;
    /**
     * @return The employee identifier assigned to the user by the organisation.
     * 
     */
    private String employeeId;
    /**
     * @return The primary email address of the user.
     * 
     */
    private String mail;
    /**
     * @return The email alias of the user.
     * 
     */
    private String mailNickname;
    /**
     * @return The object ID of the user.
     * 
     */
    private String objectId;
    /**
     * @return The value used to associate an on-premises Active Directory user account with their Azure AD user object.
     * 
     */
    private String onpremisesImmutableId;
    /**
     * @return The on-premise SAM account name of the user.
     * 
     */
    private String onpremisesSamAccountName;
    /**
     * @return The on-premise user principal name of the user.
     * 
     */
    private String onpremisesUserPrincipalName;
    /**
     * @return The usage location of the user.
     * 
     */
    private String usageLocation;
    /**
     * @return The user principal name (UPN) of the user.
     * 
     */
    private String userPrincipalName;

    private GetUsersUser() {}
    /**
     * @return Whether or not the account is enabled.
     * 
     */
    public Boolean accountEnabled() {
        return this.accountEnabled;
    }
    /**
     * @return The display name of the user.
     * 
     */
    public String displayName() {
        return this.displayName;
    }
    /**
     * @return The employee identifier assigned to the user by the organisation.
     * 
     */
    public String employeeId() {
        return this.employeeId;
    }
    /**
     * @return The primary email address of the user.
     * 
     */
    public String mail() {
        return this.mail;
    }
    /**
     * @return The email alias of the user.
     * 
     */
    public String mailNickname() {
        return this.mailNickname;
    }
    /**
     * @return The object ID of the user.
     * 
     */
    public String objectId() {
        return this.objectId;
    }
    /**
     * @return The value used to associate an on-premises Active Directory user account with their Azure AD user object.
     * 
     */
    public String onpremisesImmutableId() {
        return this.onpremisesImmutableId;
    }
    /**
     * @return The on-premise SAM account name of the user.
     * 
     */
    public String onpremisesSamAccountName() {
        return this.onpremisesSamAccountName;
    }
    /**
     * @return The on-premise user principal name of the user.
     * 
     */
    public String onpremisesUserPrincipalName() {
        return this.onpremisesUserPrincipalName;
    }
    /**
     * @return The usage location of the user.
     * 
     */
    public String usageLocation() {
        return this.usageLocation;
    }
    /**
     * @return The user principal name (UPN) of the user.
     * 
     */
    public String userPrincipalName() {
        return this.userPrincipalName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetUsersUser defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean accountEnabled;
        private String displayName;
        private String employeeId;
        private String mail;
        private String mailNickname;
        private String objectId;
        private String onpremisesImmutableId;
        private String onpremisesSamAccountName;
        private String onpremisesUserPrincipalName;
        private String usageLocation;
        private String userPrincipalName;
        public Builder() {}
        public Builder(GetUsersUser defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accountEnabled = defaults.accountEnabled;
    	      this.displayName = defaults.displayName;
    	      this.employeeId = defaults.employeeId;
    	      this.mail = defaults.mail;
    	      this.mailNickname = defaults.mailNickname;
    	      this.objectId = defaults.objectId;
    	      this.onpremisesImmutableId = defaults.onpremisesImmutableId;
    	      this.onpremisesSamAccountName = defaults.onpremisesSamAccountName;
    	      this.onpremisesUserPrincipalName = defaults.onpremisesUserPrincipalName;
    	      this.usageLocation = defaults.usageLocation;
    	      this.userPrincipalName = defaults.userPrincipalName;
        }

        @CustomType.Setter
        public Builder accountEnabled(Boolean accountEnabled) {
            this.accountEnabled = Objects.requireNonNull(accountEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder displayName(String displayName) {
            this.displayName = Objects.requireNonNull(displayName);
            return this;
        }
        @CustomType.Setter
        public Builder employeeId(String employeeId) {
            this.employeeId = Objects.requireNonNull(employeeId);
            return this;
        }
        @CustomType.Setter
        public Builder mail(String mail) {
            this.mail = Objects.requireNonNull(mail);
            return this;
        }
        @CustomType.Setter
        public Builder mailNickname(String mailNickname) {
            this.mailNickname = Objects.requireNonNull(mailNickname);
            return this;
        }
        @CustomType.Setter
        public Builder objectId(String objectId) {
            this.objectId = Objects.requireNonNull(objectId);
            return this;
        }
        @CustomType.Setter
        public Builder onpremisesImmutableId(String onpremisesImmutableId) {
            this.onpremisesImmutableId = Objects.requireNonNull(onpremisesImmutableId);
            return this;
        }
        @CustomType.Setter
        public Builder onpremisesSamAccountName(String onpremisesSamAccountName) {
            this.onpremisesSamAccountName = Objects.requireNonNull(onpremisesSamAccountName);
            return this;
        }
        @CustomType.Setter
        public Builder onpremisesUserPrincipalName(String onpremisesUserPrincipalName) {
            this.onpremisesUserPrincipalName = Objects.requireNonNull(onpremisesUserPrincipalName);
            return this;
        }
        @CustomType.Setter
        public Builder usageLocation(String usageLocation) {
            this.usageLocation = Objects.requireNonNull(usageLocation);
            return this;
        }
        @CustomType.Setter
        public Builder userPrincipalName(String userPrincipalName) {
            this.userPrincipalName = Objects.requireNonNull(userPrincipalName);
            return this;
        }
        public GetUsersUser build() {
            final var o = new GetUsersUser();
            o.accountEnabled = accountEnabled;
            o.displayName = displayName;
            o.employeeId = employeeId;
            o.mail = mail;
            o.mailNickname = mailNickname;
            o.objectId = objectId;
            o.onpremisesImmutableId = onpremisesImmutableId;
            o.onpremisesSamAccountName = onpremisesSamAccountName;
            o.onpremisesUserPrincipalName = onpremisesUserPrincipalName;
            o.usageLocation = usageLocation;
            o.userPrincipalName = userPrincipalName;
            return o;
        }
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetUserResult {
    /**
     * @return Whether or not the account is enabled.
     * 
     */
    private Boolean accountEnabled;
    /**
     * @return The age group of the user. Supported values are `Adult`, `NotAdult` and `Minor`.
     * 
     */
    private String ageGroup;
    /**
     * @return A list of telephone numbers for the user.
     * 
     */
    private List<String> businessPhones;
    /**
     * @return The city in which the user is located.
     * 
     */
    private String city;
    /**
     * @return The company name which the user is associated. This property can be useful for describing the company that an external user comes from.
     * 
     */
    private String companyName;
    /**
     * @return Whether consent has been obtained for minors. Supported values are `Granted`, `Denied` and `NotRequired`.
     * 
     */
    private String consentProvidedForMinor;
    /**
     * @return The cost center associated with the user.
     * 
     */
    private String costCenter;
    /**
     * @return The country/region in which the user is located, e.g. `US` or `UK`.
     * 
     */
    private String country;
    /**
     * @return Indicates whether the user account was created as a regular school or work account (`null`), an external account (`Invitation`), a local account for an Azure Active Directory B2C tenant (`LocalAccount`) or self-service sign-up using email verification (`EmailVerified`).
     * 
     */
    private String creationType;
    /**
     * @return The name for the department in which the user works.
     * 
     */
    private String department;
    /**
     * @return The display name of the user.
     * 
     */
    private String displayName;
    /**
     * @return The name of the division in which the user works.
     * 
     */
    private String division;
    /**
     * @return The employee identifier assigned to the user by the organisation.
     * 
     */
    private String employeeId;
    /**
     * @return Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.
     * 
     */
    private String employeeType;
    /**
     * @return For an external user invited to the tenant, this property represents the invited user&#39;s invitation status. Possible values are `PendingAcceptance` or `Accepted`.
     * 
     */
    private String externalUserState;
    /**
     * @return The fax number of the user.
     * 
     */
    private String faxNumber;
    /**
     * @return The given name (first name) of the user.
     * 
     */
    private String givenName;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A list of instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.
     * 
     */
    private List<String> imAddresses;
    /**
     * @return The user’s job title.
     * 
     */
    private String jobTitle;
    /**
     * @return The SMTP address for the user.
     * 
     */
    private String mail;
    /**
     * @return The email alias of the user.
     * 
     */
    private String mailNickname;
    /**
     * @return The object ID of the user&#39;s manager.
     * 
     */
    private String managerId;
    /**
     * @return The primary cellular telephone number for the user.
     * 
     */
    private String mobilePhone;
    /**
     * @return The object ID of the user.
     * 
     */
    private String objectId;
    /**
     * @return The office location in the user&#39;s place of business.
     * 
     */
    private String officeLocation;
    /**
     * @return The on-premises distinguished name (DN) of the user, synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    private String onpremisesDistinguishedName;
    /**
     * @return The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    private String onpremisesDomainName;
    /**
     * @return The value used to associate an on-premise Active Directory user account with their Azure AD user object.
     * 
     */
    private String onpremisesImmutableId;
    /**
     * @return The on-premise SAM account name of the user.
     * 
     */
    private String onpremisesSamAccountName;
    /**
     * @return The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    private String onpremisesSecurityIdentifier;
    /**
     * @return Whether this user is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
     * 
     */
    private Boolean onpremisesSyncEnabled;
    /**
     * @return The on-premise user principal name of the user.
     * 
     */
    private String onpremisesUserPrincipalName;
    /**
     * @return A list of additional email addresses for the user.
     * 
     */
    private List<String> otherMails;
    /**
     * @return The postal code for the user&#39;s postal address. The postal code is specific to the user&#39;s country/region. In the United States of America, this attribute contains the ZIP code.
     * 
     */
    private String postalCode;
    /**
     * @return The user&#39;s preferred language, in ISO 639-1 notation.
     * 
     */
    private String preferredLanguage;
    /**
     * @return List of email addresses for the user that direct to the same mailbox.
     * 
     */
    private List<String> proxyAddresses;
    /**
     * @return Whether or not the Outlook global address list should include this user.
     * 
     */
    private Boolean showInAddressList;
    /**
     * @return The state or province in the user&#39;s address.
     * 
     */
    private String state;
    /**
     * @return The street address of the user&#39;s place of business.
     * 
     */
    private String streetAddress;
    /**
     * @return The user&#39;s surname (family name or last name).
     * 
     */
    private String surname;
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
    /**
     * @return The user type in the directory. Possible values are `Guest` or `Member`.
     * 
     */
    private String userType;

    private GetUserResult() {}
    /**
     * @return Whether or not the account is enabled.
     * 
     */
    public Boolean accountEnabled() {
        return this.accountEnabled;
    }
    /**
     * @return The age group of the user. Supported values are `Adult`, `NotAdult` and `Minor`.
     * 
     */
    public String ageGroup() {
        return this.ageGroup;
    }
    /**
     * @return A list of telephone numbers for the user.
     * 
     */
    public List<String> businessPhones() {
        return this.businessPhones;
    }
    /**
     * @return The city in which the user is located.
     * 
     */
    public String city() {
        return this.city;
    }
    /**
     * @return The company name which the user is associated. This property can be useful for describing the company that an external user comes from.
     * 
     */
    public String companyName() {
        return this.companyName;
    }
    /**
     * @return Whether consent has been obtained for minors. Supported values are `Granted`, `Denied` and `NotRequired`.
     * 
     */
    public String consentProvidedForMinor() {
        return this.consentProvidedForMinor;
    }
    /**
     * @return The cost center associated with the user.
     * 
     */
    public String costCenter() {
        return this.costCenter;
    }
    /**
     * @return The country/region in which the user is located, e.g. `US` or `UK`.
     * 
     */
    public String country() {
        return this.country;
    }
    /**
     * @return Indicates whether the user account was created as a regular school or work account (`null`), an external account (`Invitation`), a local account for an Azure Active Directory B2C tenant (`LocalAccount`) or self-service sign-up using email verification (`EmailVerified`).
     * 
     */
    public String creationType() {
        return this.creationType;
    }
    /**
     * @return The name for the department in which the user works.
     * 
     */
    public String department() {
        return this.department;
    }
    /**
     * @return The display name of the user.
     * 
     */
    public String displayName() {
        return this.displayName;
    }
    /**
     * @return The name of the division in which the user works.
     * 
     */
    public String division() {
        return this.division;
    }
    /**
     * @return The employee identifier assigned to the user by the organisation.
     * 
     */
    public String employeeId() {
        return this.employeeId;
    }
    /**
     * @return Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.
     * 
     */
    public String employeeType() {
        return this.employeeType;
    }
    /**
     * @return For an external user invited to the tenant, this property represents the invited user&#39;s invitation status. Possible values are `PendingAcceptance` or `Accepted`.
     * 
     */
    public String externalUserState() {
        return this.externalUserState;
    }
    /**
     * @return The fax number of the user.
     * 
     */
    public String faxNumber() {
        return this.faxNumber;
    }
    /**
     * @return The given name (first name) of the user.
     * 
     */
    public String givenName() {
        return this.givenName;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A list of instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.
     * 
     */
    public List<String> imAddresses() {
        return this.imAddresses;
    }
    /**
     * @return The user’s job title.
     * 
     */
    public String jobTitle() {
        return this.jobTitle;
    }
    /**
     * @return The SMTP address for the user.
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
     * @return The object ID of the user&#39;s manager.
     * 
     */
    public String managerId() {
        return this.managerId;
    }
    /**
     * @return The primary cellular telephone number for the user.
     * 
     */
    public String mobilePhone() {
        return this.mobilePhone;
    }
    /**
     * @return The object ID of the user.
     * 
     */
    public String objectId() {
        return this.objectId;
    }
    /**
     * @return The office location in the user&#39;s place of business.
     * 
     */
    public String officeLocation() {
        return this.officeLocation;
    }
    /**
     * @return The on-premises distinguished name (DN) of the user, synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    public String onpremisesDistinguishedName() {
        return this.onpremisesDistinguishedName;
    }
    /**
     * @return The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    public String onpremisesDomainName() {
        return this.onpremisesDomainName;
    }
    /**
     * @return The value used to associate an on-premise Active Directory user account with their Azure AD user object.
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
     * @return The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    public String onpremisesSecurityIdentifier() {
        return this.onpremisesSecurityIdentifier;
    }
    /**
     * @return Whether this user is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
     * 
     */
    public Boolean onpremisesSyncEnabled() {
        return this.onpremisesSyncEnabled;
    }
    /**
     * @return The on-premise user principal name of the user.
     * 
     */
    public String onpremisesUserPrincipalName() {
        return this.onpremisesUserPrincipalName;
    }
    /**
     * @return A list of additional email addresses for the user.
     * 
     */
    public List<String> otherMails() {
        return this.otherMails;
    }
    /**
     * @return The postal code for the user&#39;s postal address. The postal code is specific to the user&#39;s country/region. In the United States of America, this attribute contains the ZIP code.
     * 
     */
    public String postalCode() {
        return this.postalCode;
    }
    /**
     * @return The user&#39;s preferred language, in ISO 639-1 notation.
     * 
     */
    public String preferredLanguage() {
        return this.preferredLanguage;
    }
    /**
     * @return List of email addresses for the user that direct to the same mailbox.
     * 
     */
    public List<String> proxyAddresses() {
        return this.proxyAddresses;
    }
    /**
     * @return Whether or not the Outlook global address list should include this user.
     * 
     */
    public Boolean showInAddressList() {
        return this.showInAddressList;
    }
    /**
     * @return The state or province in the user&#39;s address.
     * 
     */
    public String state() {
        return this.state;
    }
    /**
     * @return The street address of the user&#39;s place of business.
     * 
     */
    public String streetAddress() {
        return this.streetAddress;
    }
    /**
     * @return The user&#39;s surname (family name or last name).
     * 
     */
    public String surname() {
        return this.surname;
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
    /**
     * @return The user type in the directory. Possible values are `Guest` or `Member`.
     * 
     */
    public String userType() {
        return this.userType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetUserResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean accountEnabled;
        private String ageGroup;
        private List<String> businessPhones;
        private String city;
        private String companyName;
        private String consentProvidedForMinor;
        private String costCenter;
        private String country;
        private String creationType;
        private String department;
        private String displayName;
        private String division;
        private String employeeId;
        private String employeeType;
        private String externalUserState;
        private String faxNumber;
        private String givenName;
        private String id;
        private List<String> imAddresses;
        private String jobTitle;
        private String mail;
        private String mailNickname;
        private String managerId;
        private String mobilePhone;
        private String objectId;
        private String officeLocation;
        private String onpremisesDistinguishedName;
        private String onpremisesDomainName;
        private String onpremisesImmutableId;
        private String onpremisesSamAccountName;
        private String onpremisesSecurityIdentifier;
        private Boolean onpremisesSyncEnabled;
        private String onpremisesUserPrincipalName;
        private List<String> otherMails;
        private String postalCode;
        private String preferredLanguage;
        private List<String> proxyAddresses;
        private Boolean showInAddressList;
        private String state;
        private String streetAddress;
        private String surname;
        private String usageLocation;
        private String userPrincipalName;
        private String userType;
        public Builder() {}
        public Builder(GetUserResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accountEnabled = defaults.accountEnabled;
    	      this.ageGroup = defaults.ageGroup;
    	      this.businessPhones = defaults.businessPhones;
    	      this.city = defaults.city;
    	      this.companyName = defaults.companyName;
    	      this.consentProvidedForMinor = defaults.consentProvidedForMinor;
    	      this.costCenter = defaults.costCenter;
    	      this.country = defaults.country;
    	      this.creationType = defaults.creationType;
    	      this.department = defaults.department;
    	      this.displayName = defaults.displayName;
    	      this.division = defaults.division;
    	      this.employeeId = defaults.employeeId;
    	      this.employeeType = defaults.employeeType;
    	      this.externalUserState = defaults.externalUserState;
    	      this.faxNumber = defaults.faxNumber;
    	      this.givenName = defaults.givenName;
    	      this.id = defaults.id;
    	      this.imAddresses = defaults.imAddresses;
    	      this.jobTitle = defaults.jobTitle;
    	      this.mail = defaults.mail;
    	      this.mailNickname = defaults.mailNickname;
    	      this.managerId = defaults.managerId;
    	      this.mobilePhone = defaults.mobilePhone;
    	      this.objectId = defaults.objectId;
    	      this.officeLocation = defaults.officeLocation;
    	      this.onpremisesDistinguishedName = defaults.onpremisesDistinguishedName;
    	      this.onpremisesDomainName = defaults.onpremisesDomainName;
    	      this.onpremisesImmutableId = defaults.onpremisesImmutableId;
    	      this.onpremisesSamAccountName = defaults.onpremisesSamAccountName;
    	      this.onpremisesSecurityIdentifier = defaults.onpremisesSecurityIdentifier;
    	      this.onpremisesSyncEnabled = defaults.onpremisesSyncEnabled;
    	      this.onpremisesUserPrincipalName = defaults.onpremisesUserPrincipalName;
    	      this.otherMails = defaults.otherMails;
    	      this.postalCode = defaults.postalCode;
    	      this.preferredLanguage = defaults.preferredLanguage;
    	      this.proxyAddresses = defaults.proxyAddresses;
    	      this.showInAddressList = defaults.showInAddressList;
    	      this.state = defaults.state;
    	      this.streetAddress = defaults.streetAddress;
    	      this.surname = defaults.surname;
    	      this.usageLocation = defaults.usageLocation;
    	      this.userPrincipalName = defaults.userPrincipalName;
    	      this.userType = defaults.userType;
        }

        @CustomType.Setter
        public Builder accountEnabled(Boolean accountEnabled) {
            this.accountEnabled = Objects.requireNonNull(accountEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder ageGroup(String ageGroup) {
            this.ageGroup = Objects.requireNonNull(ageGroup);
            return this;
        }
        @CustomType.Setter
        public Builder businessPhones(List<String> businessPhones) {
            this.businessPhones = Objects.requireNonNull(businessPhones);
            return this;
        }
        public Builder businessPhones(String... businessPhones) {
            return businessPhones(List.of(businessPhones));
        }
        @CustomType.Setter
        public Builder city(String city) {
            this.city = Objects.requireNonNull(city);
            return this;
        }
        @CustomType.Setter
        public Builder companyName(String companyName) {
            this.companyName = Objects.requireNonNull(companyName);
            return this;
        }
        @CustomType.Setter
        public Builder consentProvidedForMinor(String consentProvidedForMinor) {
            this.consentProvidedForMinor = Objects.requireNonNull(consentProvidedForMinor);
            return this;
        }
        @CustomType.Setter
        public Builder costCenter(String costCenter) {
            this.costCenter = Objects.requireNonNull(costCenter);
            return this;
        }
        @CustomType.Setter
        public Builder country(String country) {
            this.country = Objects.requireNonNull(country);
            return this;
        }
        @CustomType.Setter
        public Builder creationType(String creationType) {
            this.creationType = Objects.requireNonNull(creationType);
            return this;
        }
        @CustomType.Setter
        public Builder department(String department) {
            this.department = Objects.requireNonNull(department);
            return this;
        }
        @CustomType.Setter
        public Builder displayName(String displayName) {
            this.displayName = Objects.requireNonNull(displayName);
            return this;
        }
        @CustomType.Setter
        public Builder division(String division) {
            this.division = Objects.requireNonNull(division);
            return this;
        }
        @CustomType.Setter
        public Builder employeeId(String employeeId) {
            this.employeeId = Objects.requireNonNull(employeeId);
            return this;
        }
        @CustomType.Setter
        public Builder employeeType(String employeeType) {
            this.employeeType = Objects.requireNonNull(employeeType);
            return this;
        }
        @CustomType.Setter
        public Builder externalUserState(String externalUserState) {
            this.externalUserState = Objects.requireNonNull(externalUserState);
            return this;
        }
        @CustomType.Setter
        public Builder faxNumber(String faxNumber) {
            this.faxNumber = Objects.requireNonNull(faxNumber);
            return this;
        }
        @CustomType.Setter
        public Builder givenName(String givenName) {
            this.givenName = Objects.requireNonNull(givenName);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder imAddresses(List<String> imAddresses) {
            this.imAddresses = Objects.requireNonNull(imAddresses);
            return this;
        }
        public Builder imAddresses(String... imAddresses) {
            return imAddresses(List.of(imAddresses));
        }
        @CustomType.Setter
        public Builder jobTitle(String jobTitle) {
            this.jobTitle = Objects.requireNonNull(jobTitle);
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
        public Builder managerId(String managerId) {
            this.managerId = Objects.requireNonNull(managerId);
            return this;
        }
        @CustomType.Setter
        public Builder mobilePhone(String mobilePhone) {
            this.mobilePhone = Objects.requireNonNull(mobilePhone);
            return this;
        }
        @CustomType.Setter
        public Builder objectId(String objectId) {
            this.objectId = Objects.requireNonNull(objectId);
            return this;
        }
        @CustomType.Setter
        public Builder officeLocation(String officeLocation) {
            this.officeLocation = Objects.requireNonNull(officeLocation);
            return this;
        }
        @CustomType.Setter
        public Builder onpremisesDistinguishedName(String onpremisesDistinguishedName) {
            this.onpremisesDistinguishedName = Objects.requireNonNull(onpremisesDistinguishedName);
            return this;
        }
        @CustomType.Setter
        public Builder onpremisesDomainName(String onpremisesDomainName) {
            this.onpremisesDomainName = Objects.requireNonNull(onpremisesDomainName);
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
        public Builder onpremisesSecurityIdentifier(String onpremisesSecurityIdentifier) {
            this.onpremisesSecurityIdentifier = Objects.requireNonNull(onpremisesSecurityIdentifier);
            return this;
        }
        @CustomType.Setter
        public Builder onpremisesSyncEnabled(Boolean onpremisesSyncEnabled) {
            this.onpremisesSyncEnabled = Objects.requireNonNull(onpremisesSyncEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder onpremisesUserPrincipalName(String onpremisesUserPrincipalName) {
            this.onpremisesUserPrincipalName = Objects.requireNonNull(onpremisesUserPrincipalName);
            return this;
        }
        @CustomType.Setter
        public Builder otherMails(List<String> otherMails) {
            this.otherMails = Objects.requireNonNull(otherMails);
            return this;
        }
        public Builder otherMails(String... otherMails) {
            return otherMails(List.of(otherMails));
        }
        @CustomType.Setter
        public Builder postalCode(String postalCode) {
            this.postalCode = Objects.requireNonNull(postalCode);
            return this;
        }
        @CustomType.Setter
        public Builder preferredLanguage(String preferredLanguage) {
            this.preferredLanguage = Objects.requireNonNull(preferredLanguage);
            return this;
        }
        @CustomType.Setter
        public Builder proxyAddresses(List<String> proxyAddresses) {
            this.proxyAddresses = Objects.requireNonNull(proxyAddresses);
            return this;
        }
        public Builder proxyAddresses(String... proxyAddresses) {
            return proxyAddresses(List.of(proxyAddresses));
        }
        @CustomType.Setter
        public Builder showInAddressList(Boolean showInAddressList) {
            this.showInAddressList = Objects.requireNonNull(showInAddressList);
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            this.state = Objects.requireNonNull(state);
            return this;
        }
        @CustomType.Setter
        public Builder streetAddress(String streetAddress) {
            this.streetAddress = Objects.requireNonNull(streetAddress);
            return this;
        }
        @CustomType.Setter
        public Builder surname(String surname) {
            this.surname = Objects.requireNonNull(surname);
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
        @CustomType.Setter
        public Builder userType(String userType) {
            this.userType = Objects.requireNonNull(userType);
            return this;
        }
        public GetUserResult build() {
            final var _resultValue = new GetUserResult();
            _resultValue.accountEnabled = accountEnabled;
            _resultValue.ageGroup = ageGroup;
            _resultValue.businessPhones = businessPhones;
            _resultValue.city = city;
            _resultValue.companyName = companyName;
            _resultValue.consentProvidedForMinor = consentProvidedForMinor;
            _resultValue.costCenter = costCenter;
            _resultValue.country = country;
            _resultValue.creationType = creationType;
            _resultValue.department = department;
            _resultValue.displayName = displayName;
            _resultValue.division = division;
            _resultValue.employeeId = employeeId;
            _resultValue.employeeType = employeeType;
            _resultValue.externalUserState = externalUserState;
            _resultValue.faxNumber = faxNumber;
            _resultValue.givenName = givenName;
            _resultValue.id = id;
            _resultValue.imAddresses = imAddresses;
            _resultValue.jobTitle = jobTitle;
            _resultValue.mail = mail;
            _resultValue.mailNickname = mailNickname;
            _resultValue.managerId = managerId;
            _resultValue.mobilePhone = mobilePhone;
            _resultValue.objectId = objectId;
            _resultValue.officeLocation = officeLocation;
            _resultValue.onpremisesDistinguishedName = onpremisesDistinguishedName;
            _resultValue.onpremisesDomainName = onpremisesDomainName;
            _resultValue.onpremisesImmutableId = onpremisesImmutableId;
            _resultValue.onpremisesSamAccountName = onpremisesSamAccountName;
            _resultValue.onpremisesSecurityIdentifier = onpremisesSecurityIdentifier;
            _resultValue.onpremisesSyncEnabled = onpremisesSyncEnabled;
            _resultValue.onpremisesUserPrincipalName = onpremisesUserPrincipalName;
            _resultValue.otherMails = otherMails;
            _resultValue.postalCode = postalCode;
            _resultValue.preferredLanguage = preferredLanguage;
            _resultValue.proxyAddresses = proxyAddresses;
            _resultValue.showInAddressList = showInAddressList;
            _resultValue.state = state;
            _resultValue.streetAddress = streetAddress;
            _resultValue.surname = surname;
            _resultValue.usageLocation = usageLocation;
            _resultValue.userPrincipalName = userPrincipalName;
            _resultValue.userType = userType;
            return _resultValue;
        }
    }
}

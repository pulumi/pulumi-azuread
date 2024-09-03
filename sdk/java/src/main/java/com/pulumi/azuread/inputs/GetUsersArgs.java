// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetUsersArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetUsersArgs Empty = new GetUsersArgs();

    /**
     * The employee identifiers assigned to the users by the organisation.
     * 
     */
    @Import(name="employeeIds")
    private @Nullable Output<List<String>> employeeIds;

    /**
     * @return The employee identifiers assigned to the users by the organisation.
     * 
     */
    public Optional<Output<List<String>>> employeeIds() {
        return Optional.ofNullable(this.employeeIds);
    }

    /**
     * Ignore missing users and return users that were found. The data source will still fail if no users are found. Cannot be specified with `return_all`. Defaults to `false`.
     * 
     */
    @Import(name="ignoreMissing")
    private @Nullable Output<Boolean> ignoreMissing;

    /**
     * @return Ignore missing users and return users that were found. The data source will still fail if no users are found. Cannot be specified with `return_all`. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> ignoreMissing() {
        return Optional.ofNullable(this.ignoreMissing);
    }

    /**
     * The email aliases of the users.
     * 
     */
    @Import(name="mailNicknames")
    private @Nullable Output<List<String>> mailNicknames;

    /**
     * @return The email aliases of the users.
     * 
     */
    public Optional<Output<List<String>>> mailNicknames() {
        return Optional.ofNullable(this.mailNicknames);
    }

    /**
     * The SMTP email addresses of the users.
     * 
     */
    @Import(name="mails")
    private @Nullable Output<List<String>> mails;

    /**
     * @return The SMTP email addresses of the users.
     * 
     */
    public Optional<Output<List<String>>> mails() {
        return Optional.ofNullable(this.mails);
    }

    /**
     * The object IDs of the users.
     * 
     */
    @Import(name="objectIds")
    private @Nullable Output<List<String>> objectIds;

    /**
     * @return The object IDs of the users.
     * 
     */
    public Optional<Output<List<String>>> objectIds() {
        return Optional.ofNullable(this.objectIds);
    }

    /**
     * When `true`, the data source will return all users. Cannot be used with `ignore_missing`. Defaults to `false`.
     * 
     */
    @Import(name="returnAll")
    private @Nullable Output<Boolean> returnAll;

    /**
     * @return When `true`, the data source will return all users. Cannot be used with `ignore_missing`. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> returnAll() {
        return Optional.ofNullable(this.returnAll);
    }

    /**
     * The user principal names (UPNs) of the users.
     * 
     * &gt; Either `return_all`, or one of `user_principal_names`, `object_ids`, `mail_nicknames`, `mails`, or `employee_ids` must be specified. These *may* be specified as an empty list, in which case no results will be returned.
     * 
     */
    @Import(name="userPrincipalNames")
    private @Nullable Output<List<String>> userPrincipalNames;

    /**
     * @return The user principal names (UPNs) of the users.
     * 
     * &gt; Either `return_all`, or one of `user_principal_names`, `object_ids`, `mail_nicknames`, `mails`, or `employee_ids` must be specified. These *may* be specified as an empty list, in which case no results will be returned.
     * 
     */
    public Optional<Output<List<String>>> userPrincipalNames() {
        return Optional.ofNullable(this.userPrincipalNames);
    }

    private GetUsersArgs() {}

    private GetUsersArgs(GetUsersArgs $) {
        this.employeeIds = $.employeeIds;
        this.ignoreMissing = $.ignoreMissing;
        this.mailNicknames = $.mailNicknames;
        this.mails = $.mails;
        this.objectIds = $.objectIds;
        this.returnAll = $.returnAll;
        this.userPrincipalNames = $.userPrincipalNames;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetUsersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetUsersArgs $;

        public Builder() {
            $ = new GetUsersArgs();
        }

        public Builder(GetUsersArgs defaults) {
            $ = new GetUsersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param employeeIds The employee identifiers assigned to the users by the organisation.
         * 
         * @return builder
         * 
         */
        public Builder employeeIds(@Nullable Output<List<String>> employeeIds) {
            $.employeeIds = employeeIds;
            return this;
        }

        /**
         * @param employeeIds The employee identifiers assigned to the users by the organisation.
         * 
         * @return builder
         * 
         */
        public Builder employeeIds(List<String> employeeIds) {
            return employeeIds(Output.of(employeeIds));
        }

        /**
         * @param employeeIds The employee identifiers assigned to the users by the organisation.
         * 
         * @return builder
         * 
         */
        public Builder employeeIds(String... employeeIds) {
            return employeeIds(List.of(employeeIds));
        }

        /**
         * @param ignoreMissing Ignore missing users and return users that were found. The data source will still fail if no users are found. Cannot be specified with `return_all`. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder ignoreMissing(@Nullable Output<Boolean> ignoreMissing) {
            $.ignoreMissing = ignoreMissing;
            return this;
        }

        /**
         * @param ignoreMissing Ignore missing users and return users that were found. The data source will still fail if no users are found. Cannot be specified with `return_all`. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder ignoreMissing(Boolean ignoreMissing) {
            return ignoreMissing(Output.of(ignoreMissing));
        }

        /**
         * @param mailNicknames The email aliases of the users.
         * 
         * @return builder
         * 
         */
        public Builder mailNicknames(@Nullable Output<List<String>> mailNicknames) {
            $.mailNicknames = mailNicknames;
            return this;
        }

        /**
         * @param mailNicknames The email aliases of the users.
         * 
         * @return builder
         * 
         */
        public Builder mailNicknames(List<String> mailNicknames) {
            return mailNicknames(Output.of(mailNicknames));
        }

        /**
         * @param mailNicknames The email aliases of the users.
         * 
         * @return builder
         * 
         */
        public Builder mailNicknames(String... mailNicknames) {
            return mailNicknames(List.of(mailNicknames));
        }

        /**
         * @param mails The SMTP email addresses of the users.
         * 
         * @return builder
         * 
         */
        public Builder mails(@Nullable Output<List<String>> mails) {
            $.mails = mails;
            return this;
        }

        /**
         * @param mails The SMTP email addresses of the users.
         * 
         * @return builder
         * 
         */
        public Builder mails(List<String> mails) {
            return mails(Output.of(mails));
        }

        /**
         * @param mails The SMTP email addresses of the users.
         * 
         * @return builder
         * 
         */
        public Builder mails(String... mails) {
            return mails(List.of(mails));
        }

        /**
         * @param objectIds The object IDs of the users.
         * 
         * @return builder
         * 
         */
        public Builder objectIds(@Nullable Output<List<String>> objectIds) {
            $.objectIds = objectIds;
            return this;
        }

        /**
         * @param objectIds The object IDs of the users.
         * 
         * @return builder
         * 
         */
        public Builder objectIds(List<String> objectIds) {
            return objectIds(Output.of(objectIds));
        }

        /**
         * @param objectIds The object IDs of the users.
         * 
         * @return builder
         * 
         */
        public Builder objectIds(String... objectIds) {
            return objectIds(List.of(objectIds));
        }

        /**
         * @param returnAll When `true`, the data source will return all users. Cannot be used with `ignore_missing`. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder returnAll(@Nullable Output<Boolean> returnAll) {
            $.returnAll = returnAll;
            return this;
        }

        /**
         * @param returnAll When `true`, the data source will return all users. Cannot be used with `ignore_missing`. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder returnAll(Boolean returnAll) {
            return returnAll(Output.of(returnAll));
        }

        /**
         * @param userPrincipalNames The user principal names (UPNs) of the users.
         * 
         * &gt; Either `return_all`, or one of `user_principal_names`, `object_ids`, `mail_nicknames`, `mails`, or `employee_ids` must be specified. These *may* be specified as an empty list, in which case no results will be returned.
         * 
         * @return builder
         * 
         */
        public Builder userPrincipalNames(@Nullable Output<List<String>> userPrincipalNames) {
            $.userPrincipalNames = userPrincipalNames;
            return this;
        }

        /**
         * @param userPrincipalNames The user principal names (UPNs) of the users.
         * 
         * &gt; Either `return_all`, or one of `user_principal_names`, `object_ids`, `mail_nicknames`, `mails`, or `employee_ids` must be specified. These *may* be specified as an empty list, in which case no results will be returned.
         * 
         * @return builder
         * 
         */
        public Builder userPrincipalNames(List<String> userPrincipalNames) {
            return userPrincipalNames(Output.of(userPrincipalNames));
        }

        /**
         * @param userPrincipalNames The user principal names (UPNs) of the users.
         * 
         * &gt; Either `return_all`, or one of `user_principal_names`, `object_ids`, `mail_nicknames`, `mails`, or `employee_ids` must be specified. These *may* be specified as an empty list, in which case no results will be returned.
         * 
         * @return builder
         * 
         */
        public Builder userPrincipalNames(String... userPrincipalNames) {
            return userPrincipalNames(List.of(userPrincipalNames));
        }

        public GetUsersArgs build() {
            return $;
        }
    }

}

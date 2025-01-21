// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.ApplicationRegistrationArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.ApplicationRegistrationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an application registration within Azure Active Directory.
 * 
 * For a more comprehensive alternative, please see the azuread.Application resource. Please note that this resource should not be used together with the `azuread.Application` resource when managing the same application.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`
 * 
 * When authenticated with a user principal, this resource may require one of the following directory roles: `Application Administrator` or `Global Administrator`
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.ApplicationRegistration;
 * import com.pulumi.azuread.ApplicationRegistrationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new ApplicationRegistration("example", ApplicationRegistrationArgs.builder()
 *             .displayName("Example Application")
 *             .description("My example application")
 *             .signInAudience("AzureADMyOrg")
 *             .homepageUrl("https://app.example.com/")
 *             .logoutUrl("https://app.example.com/logout")
 *             .marketingUrl("https://example.com/")
 *             .privacyStatementUrl("https://example.com/privacy")
 *             .supportUrl("https://support.example.com/")
 *             .termsOfServiceUrl("https://example.com/terms")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Application Registrations can be imported using the object ID of the application, in the following format.
 * 
 * ```sh
 * $ pulumi import azuread:index/applicationRegistration:ApplicationRegistration example /applications/00000000-0000-0000-0000-000000000000
 * ```
 * 
 */
@ResourceType(type="azuread:index/applicationRegistration:ApplicationRegistration")
public class ApplicationRegistration extends com.pulumi.resources.CustomResource {
    /**
     * The Client ID for the application, which is globally unique.
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output<String> clientId;

    /**
     * @return The Client ID for the application, which is globally unique.
     * 
     */
    public Output<String> clientId() {
        return this.clientId;
    }
    /**
     * A description of the application, as shown to end users.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the application, as shown to end users.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
     * 
     */
    @Export(name="disabledByMicrosoft", refs={String.class}, tree="[0]")
    private Output<String> disabledByMicrosoft;

    /**
     * @return Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
     * 
     */
    public Output<String> disabledByMicrosoft() {
        return this.disabledByMicrosoft;
    }
    /**
     * The display name for the application.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The display name for the application.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * Configures the `groups` claim issued in a user or OAuth access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
     * 
     */
    @Export(name="groupMembershipClaims", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> groupMembershipClaims;

    /**
     * @return Configures the `groups` claim issued in a user or OAuth access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
     * 
     */
    public Output<Optional<List<String>>> groupMembershipClaims() {
        return Codegen.optional(this.groupMembershipClaims);
    }
    /**
     * Home page or landing page of the application.
     * 
     */
    @Export(name="homepageUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> homepageUrl;

    /**
     * @return Home page or landing page of the application.
     * 
     */
    public Output<Optional<String>> homepageUrl() {
        return Codegen.optional(this.homepageUrl);
    }
    /**
     * Whether this web application can request an access token using OAuth implicit flow.
     * 
     */
    @Export(name="implicitAccessTokenIssuanceEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> implicitAccessTokenIssuanceEnabled;

    /**
     * @return Whether this web application can request an access token using OAuth implicit flow.
     * 
     */
    public Output<Optional<Boolean>> implicitAccessTokenIssuanceEnabled() {
        return Codegen.optional(this.implicitAccessTokenIssuanceEnabled);
    }
    /**
     * Whether this web application can request an ID token using OAuth implicit flow.
     * 
     */
    @Export(name="implicitIdTokenIssuanceEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> implicitIdTokenIssuanceEnabled;

    /**
     * @return Whether this web application can request an ID token using OAuth implicit flow.
     * 
     */
    public Output<Optional<Boolean>> implicitIdTokenIssuanceEnabled() {
        return Codegen.optional(this.implicitIdTokenIssuanceEnabled);
    }
    /**
     * The URL that will be used by Microsoft&#39;s authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
     * 
     */
    @Export(name="logoutUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> logoutUrl;

    /**
     * @return The URL that will be used by Microsoft&#39;s authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
     * 
     */
    public Output<Optional<String>> logoutUrl() {
        return Codegen.optional(this.logoutUrl);
    }
    /**
     * URL of the marketing page for the application.
     * 
     */
    @Export(name="marketingUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> marketingUrl;

    /**
     * @return URL of the marketing page for the application.
     * 
     */
    public Output<Optional<String>> marketingUrl() {
        return Codegen.optional(this.marketingUrl);
    }
    /**
     * User-specified notes relevant for the management of the application.
     * 
     */
    @Export(name="notes", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> notes;

    /**
     * @return User-specified notes relevant for the management of the application.
     * 
     */
    public Output<Optional<String>> notes() {
        return Codegen.optional(this.notes);
    }
    /**
     * The object ID of the application within the tenant.
     * 
     */
    @Export(name="objectId", refs={String.class}, tree="[0]")
    private Output<String> objectId;

    /**
     * @return The object ID of the application within the tenant.
     * 
     */
    public Output<String> objectId() {
        return this.objectId;
    }
    /**
     * URL of the privacy statement for the application.
     * 
     */
    @Export(name="privacyStatementUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> privacyStatementUrl;

    /**
     * @return URL of the privacy statement for the application.
     * 
     */
    public Output<Optional<String>> privacyStatementUrl() {
        return Codegen.optional(this.privacyStatementUrl);
    }
    /**
     * The verified publisher domain for the application.
     * 
     */
    @Export(name="publisherDomain", refs={String.class}, tree="[0]")
    private Output<String> publisherDomain;

    /**
     * @return The verified publisher domain for the application.
     * 
     */
    public Output<String> publisherDomain() {
        return this.publisherDomain;
    }
    /**
     * The access token version expected by this resource. Must be one of `1` or `2`, and must be `2` when `sign_in_audience` is either `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount` Defaults to `2`.
     * 
     */
    @Export(name="requestedAccessTokenVersion", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> requestedAccessTokenVersion;

    /**
     * @return The access token version expected by this resource. Must be one of `1` or `2`, and must be `2` when `sign_in_audience` is either `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount` Defaults to `2`.
     * 
     */
    public Output<Optional<Integer>> requestedAccessTokenVersion() {
        return Codegen.optional(this.requestedAccessTokenVersion);
    }
    /**
     * References application context information from a Service or Asset Management database.
     * 
     */
    @Export(name="serviceManagementReference", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> serviceManagementReference;

    /**
     * @return References application context information from a Service or Asset Management database.
     * 
     */
    public Output<Optional<String>> serviceManagementReference() {
        return Codegen.optional(this.serviceManagementReference);
    }
    /**
     * The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
     * 
     */
    @Export(name="signInAudience", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> signInAudience;

    /**
     * @return The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
     * 
     */
    public Output<Optional<String>> signInAudience() {
        return Codegen.optional(this.signInAudience);
    }
    /**
     * URL of the support page for the application.
     * 
     */
    @Export(name="supportUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> supportUrl;

    /**
     * @return URL of the support page for the application.
     * 
     */
    public Output<Optional<String>> supportUrl() {
        return Codegen.optional(this.supportUrl);
    }
    /**
     * URL of the terms of service statement for the application.
     * 
     */
    @Export(name="termsOfServiceUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> termsOfServiceUrl;

    /**
     * @return URL of the terms of service statement for the application.
     * 
     */
    public Output<Optional<String>> termsOfServiceUrl() {
        return Codegen.optional(this.termsOfServiceUrl);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ApplicationRegistration(java.lang.String name) {
        this(name, ApplicationRegistrationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ApplicationRegistration(java.lang.String name, ApplicationRegistrationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ApplicationRegistration(java.lang.String name, ApplicationRegistrationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/applicationRegistration:ApplicationRegistration", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ApplicationRegistration(java.lang.String name, Output<java.lang.String> id, @Nullable ApplicationRegistrationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/applicationRegistration:ApplicationRegistration", name, state, makeResourceOptions(options, id), false);
    }

    private static ApplicationRegistrationArgs makeArgs(ApplicationRegistrationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ApplicationRegistrationArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static ApplicationRegistration get(java.lang.String name, Output<java.lang.String> id, @Nullable ApplicationRegistrationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ApplicationRegistration(name, id, state, options);
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.azuread.InvitationArgs;
import com.pulumi.azuread.Utilities;
import com.pulumi.azuread.inputs.InvitationState;
import com.pulumi.azuread.outputs.InvitationMessage;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an invitation of a guest user within Azure Active Directory.
 * 
 * ## API Permissions
 * 
 * The following API permissions are required in order to use this resource.
 * 
 * When authenticated with a service principal, this resource requires one of the following application roles: `User.Invite.All`, `User.ReadWrite.All` or `Directory.ReadWrite.All`
 * 
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Guest Inviter`, `User Administrator` or `Global Administrator`
 * 
 * ## Example Usage
 * 
 * *Basic example*
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.Invitation;
 * import com.pulumi.azuread.InvitationArgs;
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
 *         var example = new Invitation(&#34;example&#34;, InvitationArgs.builder()        
 *             .userEmailAddress(&#34;jdoe@example.com&#34;)
 *             .redirectUrl(&#34;https://portal.azure.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * *Invitation with standard message*
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.Invitation;
 * import com.pulumi.azuread.InvitationArgs;
 * import com.pulumi.azuread.inputs.InvitationMessageArgs;
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
 *         var example = new Invitation(&#34;example&#34;, InvitationArgs.builder()        
 *             .userEmailAddress(&#34;jdoe@example.com&#34;)
 *             .redirectUrl(&#34;https://portal.azure.com&#34;)
 *             .message(InvitationMessageArgs.builder()
 *                 .language(&#34;en-US&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * *Invitation with custom message body and an additional recipient*
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.azuread.Invitation;
 * import com.pulumi.azuread.InvitationArgs;
 * import com.pulumi.azuread.inputs.InvitationMessageArgs;
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
 *         var example = new Invitation(&#34;example&#34;, InvitationArgs.builder()        
 *             .userDisplayName(&#34;Bob Bobson&#34;)
 *             .userEmailAddress(&#34;bbobson@example.com&#34;)
 *             .redirectUrl(&#34;https://portal.azure.com&#34;)
 *             .message(InvitationMessageArgs.builder()
 *                 .additionalRecipients(&#34;aaliceberg@example.com&#34;)
 *                 .body(&#34;Hello there! You are invited to join my Azure tenant!&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * This resource does not support importing.
 * 
 */
@ResourceType(type="azuread:index/invitation:Invitation")
public class Invitation extends com.pulumi.resources.CustomResource {
    /**
     * A `message` block as documented below, which configures the message being sent to the invited user. If this block is omitted, no message will be sent.
     * 
     */
    @Export(name="message", refs={InvitationMessage.class}, tree="[0]")
    private Output</* @Nullable */ InvitationMessage> message;

    /**
     * @return A `message` block as documented below, which configures the message being sent to the invited user. If this block is omitted, no message will be sent.
     * 
     */
    public Output<Optional<InvitationMessage>> message() {
        return Codegen.optional(this.message);
    }
    /**
     * The URL the user can use to redeem their invitation.
     * 
     */
    @Export(name="redeemUrl", refs={String.class}, tree="[0]")
    private Output<String> redeemUrl;

    /**
     * @return The URL the user can use to redeem their invitation.
     * 
     */
    public Output<String> redeemUrl() {
        return this.redeemUrl;
    }
    /**
     * The URL that the user should be redirected to once the invitation is redeemed.
     * 
     */
    @Export(name="redirectUrl", refs={String.class}, tree="[0]")
    private Output<String> redirectUrl;

    /**
     * @return The URL that the user should be redirected to once the invitation is redeemed.
     * 
     */
    public Output<String> redirectUrl() {
        return this.redirectUrl;
    }
    /**
     * The display name of the user being invited.
     * 
     */
    @Export(name="userDisplayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userDisplayName;

    /**
     * @return The display name of the user being invited.
     * 
     */
    public Output<Optional<String>> userDisplayName() {
        return Codegen.optional(this.userDisplayName);
    }
    /**
     * The email address of the user being invited.
     * 
     */
    @Export(name="userEmailAddress", refs={String.class}, tree="[0]")
    private Output<String> userEmailAddress;

    /**
     * @return The email address of the user being invited.
     * 
     */
    public Output<String> userEmailAddress() {
        return this.userEmailAddress;
    }
    /**
     * Object ID of the invited user.
     * 
     */
    @Export(name="userId", refs={String.class}, tree="[0]")
    private Output<String> userId;

    /**
     * @return Object ID of the invited user.
     * 
     */
    public Output<String> userId() {
        return this.userId;
    }
    /**
     * The user type of the user being invited. Must be one of `Guest` or `Member`. Only Global Administrators can invite users as members. Defaults to `Guest`.
     * 
     */
    @Export(name="userType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userType;

    /**
     * @return The user type of the user being invited. Must be one of `Guest` or `Member`. Only Global Administrators can invite users as members. Defaults to `Guest`.
     * 
     */
    public Output<Optional<String>> userType() {
        return Codegen.optional(this.userType);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Invitation(String name) {
        this(name, InvitationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Invitation(String name, InvitationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Invitation(String name, InvitationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/invitation:Invitation", name, args == null ? InvitationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Invitation(String name, Output<String> id, @Nullable InvitationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("azuread:index/invitation:Invitation", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static Invitation get(String name, Output<String> id, @Nullable InvitationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Invitation(name, id, state, options);
    }
}

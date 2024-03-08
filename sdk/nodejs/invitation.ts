// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

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
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = new azuread.Invitation("example", {
 *     userEmailAddress: "jdoe@example.com",
 *     redirectUrl: "https://portal.azure.com",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * *Invitation with standard message*
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = new azuread.Invitation("example", {
 *     userEmailAddress: "jdoe@example.com",
 *     redirectUrl: "https://portal.azure.com",
 *     message: {
 *         language: "en-US",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * *Invitation with custom message body and an additional recipient*
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = new azuread.Invitation("example", {
 *     userDisplayName: "Bob Bobson",
 *     userEmailAddress: "bbobson@example.com",
 *     redirectUrl: "https://portal.azure.com",
 *     message: {
 *         additionalRecipients: "aaliceberg@example.com",
 *         body: "Hello there! You are invited to join my Azure tenant!",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * This resource does not support importing.
 */
export class Invitation extends pulumi.CustomResource {
    /**
     * Get an existing Invitation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InvitationState, opts?: pulumi.CustomResourceOptions): Invitation {
        return new Invitation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/invitation:Invitation';

    /**
     * Returns true if the given object is an instance of Invitation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Invitation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Invitation.__pulumiType;
    }

    /**
     * A `message` block as documented below, which configures the message being sent to the invited user. If this block is omitted, no message will be sent.
     */
    public readonly message!: pulumi.Output<outputs.InvitationMessage | undefined>;
    /**
     * The URL the user can use to redeem their invitation.
     */
    public /*out*/ readonly redeemUrl!: pulumi.Output<string>;
    /**
     * The URL that the user should be redirected to once the invitation is redeemed.
     */
    public readonly redirectUrl!: pulumi.Output<string>;
    /**
     * The display name of the user being invited.
     */
    public readonly userDisplayName!: pulumi.Output<string | undefined>;
    /**
     * The email address of the user being invited.
     */
    public readonly userEmailAddress!: pulumi.Output<string>;
    /**
     * Object ID of the invited user.
     */
    public /*out*/ readonly userId!: pulumi.Output<string>;
    /**
     * The user type of the user being invited. Must be one of `Guest` or `Member`. Only Global Administrators can invite users as members. Defaults to `Guest`.
     */
    public readonly userType!: pulumi.Output<string | undefined>;

    /**
     * Create a Invitation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InvitationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InvitationArgs | InvitationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InvitationState | undefined;
            resourceInputs["message"] = state ? state.message : undefined;
            resourceInputs["redeemUrl"] = state ? state.redeemUrl : undefined;
            resourceInputs["redirectUrl"] = state ? state.redirectUrl : undefined;
            resourceInputs["userDisplayName"] = state ? state.userDisplayName : undefined;
            resourceInputs["userEmailAddress"] = state ? state.userEmailAddress : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
            resourceInputs["userType"] = state ? state.userType : undefined;
        } else {
            const args = argsOrState as InvitationArgs | undefined;
            if ((!args || args.redirectUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'redirectUrl'");
            }
            if ((!args || args.userEmailAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userEmailAddress'");
            }
            resourceInputs["message"] = args ? args.message : undefined;
            resourceInputs["redirectUrl"] = args ? args.redirectUrl : undefined;
            resourceInputs["userDisplayName"] = args ? args.userDisplayName : undefined;
            resourceInputs["userEmailAddress"] = args ? args.userEmailAddress : undefined;
            resourceInputs["userType"] = args ? args.userType : undefined;
            resourceInputs["redeemUrl"] = undefined /*out*/;
            resourceInputs["userId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Invitation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Invitation resources.
 */
export interface InvitationState {
    /**
     * A `message` block as documented below, which configures the message being sent to the invited user. If this block is omitted, no message will be sent.
     */
    message?: pulumi.Input<inputs.InvitationMessage>;
    /**
     * The URL the user can use to redeem their invitation.
     */
    redeemUrl?: pulumi.Input<string>;
    /**
     * The URL that the user should be redirected to once the invitation is redeemed.
     */
    redirectUrl?: pulumi.Input<string>;
    /**
     * The display name of the user being invited.
     */
    userDisplayName?: pulumi.Input<string>;
    /**
     * The email address of the user being invited.
     */
    userEmailAddress?: pulumi.Input<string>;
    /**
     * Object ID of the invited user.
     */
    userId?: pulumi.Input<string>;
    /**
     * The user type of the user being invited. Must be one of `Guest` or `Member`. Only Global Administrators can invite users as members. Defaults to `Guest`.
     */
    userType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Invitation resource.
 */
export interface InvitationArgs {
    /**
     * A `message` block as documented below, which configures the message being sent to the invited user. If this block is omitted, no message will be sent.
     */
    message?: pulumi.Input<inputs.InvitationMessage>;
    /**
     * The URL that the user should be redirected to once the invitation is redeemed.
     */
    redirectUrl: pulumi.Input<string>;
    /**
     * The display name of the user being invited.
     */
    userDisplayName?: pulumi.Input<string>;
    /**
     * The email address of the user being invited.
     */
    userEmailAddress: pulumi.Input<string>;
    /**
     * The user type of the user being invited. Must be one of `Guest` or `Member`. Only Global Administrators can invite users as members. Defaults to `Guest`.
     */
    userType?: pulumi.Input<string>;
}

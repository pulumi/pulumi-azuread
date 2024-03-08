// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a Authentication Strength Policy within Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ConditionalAccess` and `Policy.Read.All`
 *
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Conditional Access Administrator` or `Global Administrator`
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = new azuread.AuthenticationStrengthPolicy("example", {
 *     displayName: "Example Authentication Strength Policy",
 *     description: "Policy for demo purposes",
 *     allowedCombinations: [
 *         "fido2",
 *         "password",
 *     ],
 * });
 * const example2 = new azuread.AuthenticationStrengthPolicy("example2", {
 *     displayName: "Example Authentication Strength Policy",
 *     description: "Policy for demo purposes with all possible combinations",
 *     allowedCombinations: [
 *         "fido2",
 *         "password",
 *         "deviceBasedPush",
 *         "temporaryAccessPassOneTime",
 *         "federatedMultiFactor",
 *         "federatedSingleFactor",
 *         "hardwareOath,federatedSingleFactor",
 *         "microsoftAuthenticatorPush,federatedSingleFactor",
 *         "password,hardwareOath",
 *         "password,microsoftAuthenticatorPush",
 *         "password,sms",
 *         "password,softwareOath",
 *         "password,voice",
 *         "sms",
 *         "sms,federatedSingleFactor",
 *         "softwareOath,federatedSingleFactor",
 *         "temporaryAccessPassMultiUse",
 *         "voice,federatedSingleFactor",
 *         "windowsHelloForBusiness",
 *         "x509CertificateMultiFactor",
 *         "x509CertificateSingleFactor",
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Authentication Strength Policies can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import azuread:index/authenticationStrengthPolicy:AuthenticationStrengthPolicy my_policy 00000000-0000-0000-0000-000000000000
 * ```
 */
export class AuthenticationStrengthPolicy extends pulumi.CustomResource {
    /**
     * Get an existing AuthenticationStrengthPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthenticationStrengthPolicyState, opts?: pulumi.CustomResourceOptions): AuthenticationStrengthPolicy {
        return new AuthenticationStrengthPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/authenticationStrengthPolicy:AuthenticationStrengthPolicy';

    /**
     * Returns true if the given object is an instance of AuthenticationStrengthPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthenticationStrengthPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthenticationStrengthPolicy.__pulumiType;
    }

    /**
     * List of allowed authentication methods for this authentication strength policy.
     */
    public readonly allowedCombinations!: pulumi.Output<string[]>;
    /**
     * The description for this authentication strength policy.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The friendly name for this authentication strength policy.
     */
    public readonly displayName!: pulumi.Output<string>;

    /**
     * Create a AuthenticationStrengthPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthenticationStrengthPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthenticationStrengthPolicyArgs | AuthenticationStrengthPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthenticationStrengthPolicyState | undefined;
            resourceInputs["allowedCombinations"] = state ? state.allowedCombinations : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
        } else {
            const args = argsOrState as AuthenticationStrengthPolicyArgs | undefined;
            if ((!args || args.allowedCombinations === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allowedCombinations'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["allowedCombinations"] = args ? args.allowedCombinations : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AuthenticationStrengthPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthenticationStrengthPolicy resources.
 */
export interface AuthenticationStrengthPolicyState {
    /**
     * List of allowed authentication methods for this authentication strength policy.
     */
    allowedCombinations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description for this authentication strength policy.
     */
    description?: pulumi.Input<string>;
    /**
     * The friendly name for this authentication strength policy.
     */
    displayName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthenticationStrengthPolicy resource.
 */
export interface AuthenticationStrengthPolicyArgs {
    /**
     * List of allowed authentication methods for this authentication strength policy.
     */
    allowedCombinations: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description for this authentication strength policy.
     */
    description?: pulumi.Input<string>;
    /**
     * The friendly name for this authentication strength policy.
     */
    displayName: pulumi.Input<string>;
}

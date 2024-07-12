// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a Claims Mapping Policy within Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ApplicationConfiguration` and `Policy.Read.All`
 *
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const myPolicy = new azuread.ClaimsMappingPolicy("my_policy", {
 *     definitions: [JSON.stringify({
 *         ClaimsMappingPolicy: {
 *             ClaimsSchema: [
 *                 {
 *                     ID: "employeeid",
 *                     JwtClaimType: "name",
 *                     SamlClaimType: "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/name",
 *                     Source: "user",
 *                 },
 *                 {
 *                     ID: "tenantcountry",
 *                     JwtClaimType: "country",
 *                     SamlClaimType: "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/country",
 *                     Source: "company",
 *                 },
 *             ],
 *             IncludeBasicClaimSet: "true",
 *             Version: 1,
 *         },
 *     })],
 *     displayName: "My Policy",
 * });
 * ```
 *
 * ## Import
 *
 * Claims Mapping Policy can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import azuread:index/claimsMappingPolicy:ClaimsMappingPolicy my_policy 00000000-0000-0000-0000-000000000000
 * ```
 */
export class ClaimsMappingPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ClaimsMappingPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClaimsMappingPolicyState, opts?: pulumi.CustomResourceOptions): ClaimsMappingPolicy {
        return new ClaimsMappingPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/claimsMappingPolicy:ClaimsMappingPolicy';

    /**
     * Returns true if the given object is an instance of ClaimsMappingPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClaimsMappingPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClaimsMappingPolicy.__pulumiType;
    }

    /**
     * The claims mapping policy. This is a JSON formatted string, for which the `jsonencode()` function can be used.
     */
    public readonly definitions!: pulumi.Output<string[]>;
    /**
     * The display name for this Claims Mapping Policy.
     */
    public readonly displayName!: pulumi.Output<string>;

    /**
     * Create a ClaimsMappingPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClaimsMappingPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClaimsMappingPolicyArgs | ClaimsMappingPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClaimsMappingPolicyState | undefined;
            resourceInputs["definitions"] = state ? state.definitions : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
        } else {
            const args = argsOrState as ClaimsMappingPolicyArgs | undefined;
            if ((!args || args.definitions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'definitions'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["definitions"] = args ? args.definitions : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ClaimsMappingPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClaimsMappingPolicy resources.
 */
export interface ClaimsMappingPolicyState {
    /**
     * The claims mapping policy. This is a JSON formatted string, for which the `jsonencode()` function can be used.
     */
    definitions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The display name for this Claims Mapping Policy.
     */
    displayName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClaimsMappingPolicy resource.
 */
export interface ClaimsMappingPolicyArgs {
    /**
     * The claims mapping policy. This is a JSON formatted string, for which the `jsonencode()` function can be used.
     */
    definitions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The display name for this Claims Mapping Policy.
     */
    displayName: pulumi.Input<string>;
}

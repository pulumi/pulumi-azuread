// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ### All users except guests or external users
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = new azuread.ConditionalAccessPolicy("example", {
 *     displayName: "example policy",
 *     state: "disabled",
 *     conditions: {
 *         clientAppTypes: ["all"],
 *         signInRiskLevels: ["medium"],
 *         userRiskLevels: ["medium"],
 *         applications: {
 *             includedApplications: ["All"],
 *             excludedApplications: [],
 *         },
 *         devices: {
 *             filter: {
 *                 mode: "exclude",
 *                 rule: "device.operatingSystem eq \"Doors\"",
 *             },
 *         },
 *         locations: {
 *             includedLocations: ["All"],
 *             excludedLocations: ["AllTrusted"],
 *         },
 *         platforms: {
 *             includedPlatforms: ["android"],
 *             excludedPlatforms: ["iOS"],
 *         },
 *         users: {
 *             includedUsers: ["All"],
 *             excludedUsers: ["GuestsOrExternalUsers"],
 *         },
 *     },
 *     grantControls: {
 *         operator: "OR",
 *         builtInControls: ["mfa"],
 *     },
 *     sessionControls: {
 *         applicationEnforcedRestrictionsEnabled: true,
 *         disableResilienceDefaults: false,
 *         signInFrequency: 10,
 *         signInFrequencyPeriod: "hours",
 *         cloudAppSecurityPolicy: "monitorOnly",
 *     },
 * });
 * ```
 *
 * ### Included client applications / service principals
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const current = azuread.getClientConfig({});
 * const example = new azuread.ConditionalAccessPolicy("example", {
 *     displayName: "example policy",
 *     state: "disabled",
 *     conditions: {
 *         clientAppTypes: ["all"],
 *         applications: {
 *             includedApplications: ["All"],
 *         },
 *         clientApplications: {
 *             includedServicePrincipals: [current.then(current => current.objectId)],
 *             excludedServicePrincipals: [],
 *         },
 *         users: {
 *             includedUsers: ["None"],
 *         },
 *     },
 *     grantControls: {
 *         operator: "OR",
 *         builtInControls: ["block"],
 *     },
 * });
 * ```
 *
 * ### Excluded client applications / service principals
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const current = azuread.getClientConfig({});
 * const example = new azuread.ConditionalAccessPolicy("example", {
 *     displayName: "example policy",
 *     state: "disabled",
 *     conditions: {
 *         clientAppTypes: ["all"],
 *         applications: {
 *             includedApplications: ["All"],
 *         },
 *         clientApplications: {
 *             includedServicePrincipals: ["ServicePrincipalsInMyTenant"],
 *             excludedServicePrincipals: [current.then(current => current.objectId)],
 *         },
 *         users: {
 *             includedUsers: ["None"],
 *         },
 *     },
 *     grantControls: {
 *         operator: "OR",
 *         builtInControls: ["block"],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Conditional Access Policies can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy my_location 00000000-0000-0000-0000-000000000000
 * ```
 */
export class ConditionalAccessPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ConditionalAccessPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConditionalAccessPolicyState, opts?: pulumi.CustomResourceOptions): ConditionalAccessPolicy {
        return new ConditionalAccessPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy';

    /**
     * Returns true if the given object is an instance of ConditionalAccessPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConditionalAccessPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConditionalAccessPolicy.__pulumiType;
    }

    /**
     * A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
     */
    public readonly conditions!: pulumi.Output<outputs.ConditionalAccessPolicyConditions>;
    /**
     * The friendly name for this Conditional Access Policy.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * A `grantControls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
     */
    public readonly grantControls!: pulumi.Output<outputs.ConditionalAccessPolicyGrantControls | undefined>;
    /**
     * A `sessionControls` block as documented below, which specifies the session controls that are enforced after sign-in.
     *
     * > Note: At least one of `grantControls` and/or `sessionControls` blocks must be specified.
     */
    public readonly sessionControls!: pulumi.Output<outputs.ConditionalAccessPolicySessionControls | undefined>;
    /**
     * Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
     */
    public readonly state!: pulumi.Output<string>;

    /**
     * Create a ConditionalAccessPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConditionalAccessPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConditionalAccessPolicyArgs | ConditionalAccessPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConditionalAccessPolicyState | undefined;
            resourceInputs["conditions"] = state ? state.conditions : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["grantControls"] = state ? state.grantControls : undefined;
            resourceInputs["sessionControls"] = state ? state.sessionControls : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as ConditionalAccessPolicyArgs | undefined;
            if ((!args || args.conditions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'conditions'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.state === undefined) && !opts.urn) {
                throw new Error("Missing required property 'state'");
            }
            resourceInputs["conditions"] = args ? args.conditions : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["grantControls"] = args ? args.grantControls : undefined;
            resourceInputs["sessionControls"] = args ? args.sessionControls : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConditionalAccessPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConditionalAccessPolicy resources.
 */
export interface ConditionalAccessPolicyState {
    /**
     * A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
     */
    conditions?: pulumi.Input<inputs.ConditionalAccessPolicyConditions>;
    /**
     * The friendly name for this Conditional Access Policy.
     */
    displayName?: pulumi.Input<string>;
    /**
     * A `grantControls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
     */
    grantControls?: pulumi.Input<inputs.ConditionalAccessPolicyGrantControls>;
    /**
     * A `sessionControls` block as documented below, which specifies the session controls that are enforced after sign-in.
     *
     * > Note: At least one of `grantControls` and/or `sessionControls` blocks must be specified.
     */
    sessionControls?: pulumi.Input<inputs.ConditionalAccessPolicySessionControls>;
    /**
     * Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
     */
    state?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConditionalAccessPolicy resource.
 */
export interface ConditionalAccessPolicyArgs {
    /**
     * A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
     */
    conditions: pulumi.Input<inputs.ConditionalAccessPolicyConditions>;
    /**
     * The friendly name for this Conditional Access Policy.
     */
    displayName: pulumi.Input<string>;
    /**
     * A `grantControls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
     */
    grantControls?: pulumi.Input<inputs.ConditionalAccessPolicyGrantControls>;
    /**
     * A `sessionControls` block as documented below, which specifies the session controls that are enforced after sign-in.
     *
     * > Note: At least one of `grantControls` and/or `sessionControls` blocks must be specified.
     */
    sessionControls?: pulumi.Input<inputs.ConditionalAccessPolicySessionControls>;
    /**
     * Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
     */
    state: pulumi.Input<string>;
}

// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = new azuread.ApplicationRegistration("example", {displayName: "example"});
 * const exampleApplicationOptionalClaims = new azuread.ApplicationOptionalClaims("example", {
 *     applicationId: example.id,
 *     accessTokens: [
 *         {
 *             name: "myclaim",
 *         },
 *         {
 *             name: "otherclaim",
 *         },
 *     ],
 *     idTokens: [{
 *         name: "userclaim",
 *         source: "user",
 *         essential: true,
 *         additionalProperties: ["emit_as_roles"],
 *     }],
 *     saml2Tokens: [{
 *         name: "samlexample",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Application Optional Claims can be imported using the object ID of the application, in the following format.
 *
 * ```sh
 * $ pulumi import azuread:index/applicationOptionalClaims:ApplicationOptionalClaims example /applications/00000000-0000-0000-0000-000000000000
 * ```
 */
export class ApplicationOptionalClaims extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationOptionalClaims resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationOptionalClaimsState, opts?: pulumi.CustomResourceOptions): ApplicationOptionalClaims {
        return new ApplicationOptionalClaims(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/applicationOptionalClaims:ApplicationOptionalClaims';

    /**
     * Returns true if the given object is an instance of ApplicationOptionalClaims.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationOptionalClaims {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationOptionalClaims.__pulumiType;
    }

    /**
     * One or more `accessToken` blocks as documented below.
     */
    public readonly accessTokens!: pulumi.Output<outputs.ApplicationOptionalClaimsAccessToken[] | undefined>;
    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    public readonly applicationId!: pulumi.Output<string>;
    /**
     * One or more `idToken` blocks as documented below.
     */
    public readonly idTokens!: pulumi.Output<outputs.ApplicationOptionalClaimsIdToken[] | undefined>;
    /**
     * One or more `saml2Token` blocks as documented below.
     *
     * > At least one of `accessToken`, `idToken` or `saml2Token` must be specified
     */
    public readonly saml2Tokens!: pulumi.Output<outputs.ApplicationOptionalClaimsSaml2Token[] | undefined>;

    /**
     * Create a ApplicationOptionalClaims resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationOptionalClaimsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationOptionalClaimsArgs | ApplicationOptionalClaimsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationOptionalClaimsState | undefined;
            resourceInputs["accessTokens"] = state ? state.accessTokens : undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["idTokens"] = state ? state.idTokens : undefined;
            resourceInputs["saml2Tokens"] = state ? state.saml2Tokens : undefined;
        } else {
            const args = argsOrState as ApplicationOptionalClaimsArgs | undefined;
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            resourceInputs["accessTokens"] = args ? args.accessTokens : undefined;
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["idTokens"] = args ? args.idTokens : undefined;
            resourceInputs["saml2Tokens"] = args ? args.saml2Tokens : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApplicationOptionalClaims.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationOptionalClaims resources.
 */
export interface ApplicationOptionalClaimsState {
    /**
     * One or more `accessToken` blocks as documented below.
     */
    accessTokens?: pulumi.Input<pulumi.Input<inputs.ApplicationOptionalClaimsAccessToken>[]>;
    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * One or more `idToken` blocks as documented below.
     */
    idTokens?: pulumi.Input<pulumi.Input<inputs.ApplicationOptionalClaimsIdToken>[]>;
    /**
     * One or more `saml2Token` blocks as documented below.
     *
     * > At least one of `accessToken`, `idToken` or `saml2Token` must be specified
     */
    saml2Tokens?: pulumi.Input<pulumi.Input<inputs.ApplicationOptionalClaimsSaml2Token>[]>;
}

/**
 * The set of arguments for constructing a ApplicationOptionalClaims resource.
 */
export interface ApplicationOptionalClaimsArgs {
    /**
     * One or more `accessToken` blocks as documented below.
     */
    accessTokens?: pulumi.Input<pulumi.Input<inputs.ApplicationOptionalClaimsAccessToken>[]>;
    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    applicationId: pulumi.Input<string>;
    /**
     * One or more `idToken` blocks as documented below.
     */
    idTokens?: pulumi.Input<pulumi.Input<inputs.ApplicationOptionalClaimsIdToken>[]>;
    /**
     * One or more `saml2Token` blocks as documented below.
     *
     * > At least one of `accessToken`, `idToken` or `saml2Token` must be specified
     */
    saml2Tokens?: pulumi.Input<pulumi.Input<inputs.ApplicationOptionalClaimsSaml2Token>[]>;
}

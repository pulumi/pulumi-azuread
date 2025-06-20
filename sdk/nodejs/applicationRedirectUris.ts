// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = new azuread.ApplicationRegistration("example", {displayName: "example"});
 * const examplePublic = new azuread.ApplicationRedirectUris("example_public", {
 *     applicationId: example.id,
 *     type: "PublicClient",
 *     redirectUris: [
 *         "myapp://auth",
 *         "sample.mobile.app.bundie.id://auth",
 *         "https://login.microsoftonline.com/common/oauth2/nativeclient",
 *         "https://login.live.com/oauth20_desktop.srf",
 *         "ms-appx-web://Microsoft.AAD.BrokerPlugin/00000000-1111-1111-1111-222222222222",
 *         "urn:ietf:wg:oauth:2.0:foo",
 *     ],
 * });
 * const exampleSpa = new azuread.ApplicationRedirectUris("example_spa", {
 *     applicationId: example.id,
 *     type: "SPA",
 *     redirectUris: [
 *         "https://mobile.example.com/",
 *         "https://beta.example.com/",
 *     ],
 * });
 * const exampleWeb = new azuread.ApplicationRedirectUris("example_web", {
 *     applicationId: example.id,
 *     type: "Web",
 *     redirectUris: [
 *         "https://app.example.com/",
 *         "https://classic.example.com/",
 *         "urn:ietf:wg:oauth:2.0:oob",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Application API Access can be imported using the object ID of the application and the URI type, in the following format.
 *
 * ```sh
 * $ pulumi import azuread:index/applicationRedirectUris:ApplicationRedirectUris example /applications/00000000-0000-0000-0000-000000000000/redirectUris/Web
 * ```
 */
export class ApplicationRedirectUris extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationRedirectUris resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationRedirectUrisState, opts?: pulumi.CustomResourceOptions): ApplicationRedirectUris {
        return new ApplicationRedirectUris(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/applicationRedirectUris:ApplicationRedirectUris';

    /**
     * Returns true if the given object is an instance of ApplicationRedirectUris.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationRedirectUris {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationRedirectUris.__pulumiType;
    }

    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    public readonly applicationId!: pulumi.Output<string>;
    /**
     * A set of redirect URIs to assign to the application.
     */
    public readonly redirectUris!: pulumi.Output<string[]>;
    /**
     * The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a ApplicationRedirectUris resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationRedirectUrisArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationRedirectUrisArgs | ApplicationRedirectUrisState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationRedirectUrisState | undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["redirectUris"] = state ? state.redirectUris : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ApplicationRedirectUrisArgs | undefined;
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            if ((!args || args.redirectUris === undefined) && !opts.urn) {
                throw new Error("Missing required property 'redirectUris'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["redirectUris"] = args ? args.redirectUris : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApplicationRedirectUris.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationRedirectUris resources.
 */
export interface ApplicationRedirectUrisState {
    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * A set of redirect URIs to assign to the application.
     */
    redirectUris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApplicationRedirectUris resource.
 */
export interface ApplicationRedirectUrisArgs {
    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    applicationId: pulumi.Input<string>;
    /**
     * A set of redirect URIs to assign to the application.
     */
    redirectUris: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
     */
    type: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * Application Identifier URIs can be imported using the object ID of the application and the base64-encoded identifier URI, in the following format.
 *
 * ```sh
 *  $ pulumi import azuread:index/applicationIdentifierUri:ApplicationIdentifierUri example /applications/00000000-0000-0000-0000-000000000000/identifierUris/aHR0cHM6Ly9leGFtcGxlLm5ldC8=
 * ```
 */
export class ApplicationIdentifierUri extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationIdentifierUri resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationIdentifierUriState, opts?: pulumi.CustomResourceOptions): ApplicationIdentifierUri {
        return new ApplicationIdentifierUri(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/applicationIdentifierUri:ApplicationIdentifierUri';

    /**
     * Returns true if the given object is an instance of ApplicationIdentifierUri.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationIdentifierUri {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationIdentifierUri.__pulumiType;
    }

    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    public readonly applicationId!: pulumi.Output<string>;
    /**
     * The user-defined URI that uniquely identifies an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant. Changing this forces a new resource to be created.
     */
    public readonly identifierUri!: pulumi.Output<string>;

    /**
     * Create a ApplicationIdentifierUri resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationIdentifierUriArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationIdentifierUriArgs | ApplicationIdentifierUriState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationIdentifierUriState | undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["identifierUri"] = state ? state.identifierUri : undefined;
        } else {
            const args = argsOrState as ApplicationIdentifierUriArgs | undefined;
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            if ((!args || args.identifierUri === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identifierUri'");
            }
            resourceInputs["applicationId"] = args ? args.applicationId : undefined;
            resourceInputs["identifierUri"] = args ? args.identifierUri : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApplicationIdentifierUri.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationIdentifierUri resources.
 */
export interface ApplicationIdentifierUriState {
    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * The user-defined URI that uniquely identifies an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant. Changing this forces a new resource to be created.
     */
    identifierUri?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApplicationIdentifierUri resource.
 */
export interface ApplicationIdentifierUriArgs {
    /**
     * The resource ID of the application registration. Changing this forces a new resource to be created.
     */
    applicationId: pulumi.Input<string>;
    /**
     * The user-defined URI that uniquely identifies an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant. Changing this forces a new resource to be created.
     */
    identifierUri: pulumi.Input<string>;
}

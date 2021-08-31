// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * This resource does not support importing.
 */
export class ApplicationPassword extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationPassword resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationPasswordState, opts?: pulumi.CustomResourceOptions): ApplicationPassword {
        return new ApplicationPassword(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/applicationPassword:ApplicationPassword';

    /**
     * Returns true if the given object is an instance of ApplicationPassword.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationPassword {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationPassword.__pulumiType;
    }

    /**
     * The object ID of the application for which this password should be created. Changing this field forces a new resource to be created.
     */
    public readonly applicationObjectId!: pulumi.Output<string>;
    /**
     * A display name for the password.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
     */
    public readonly endDate!: pulumi.Output<string>;
    /**
     * A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
     */
    public readonly endDateRelative!: pulumi.Output<string | undefined>;
    /**
     * Arbitrary map of values that, when changed, will trigger rotation of the password
     */
    public readonly keepers!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A UUID used to uniquely identify this password credential.
     */
    public /*out*/ readonly keyId!: pulumi.Output<string>;
    /**
     * The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
     */
    public readonly startDate!: pulumi.Output<string>;
    /**
     * The password for this application, which is generated by Azure Active Directory.
     */
    public /*out*/ readonly value!: pulumi.Output<string>;

    /**
     * Create a ApplicationPassword resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationPasswordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationPasswordArgs | ApplicationPasswordState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationPasswordState | undefined;
            inputs["applicationObjectId"] = state ? state.applicationObjectId : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["endDate"] = state ? state.endDate : undefined;
            inputs["endDateRelative"] = state ? state.endDateRelative : undefined;
            inputs["keepers"] = state ? state.keepers : undefined;
            inputs["keyId"] = state ? state.keyId : undefined;
            inputs["startDate"] = state ? state.startDate : undefined;
            inputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as ApplicationPasswordArgs | undefined;
            if ((!args || args.applicationObjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationObjectId'");
            }
            inputs["applicationObjectId"] = args ? args.applicationObjectId : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["endDate"] = args ? args.endDate : undefined;
            inputs["endDateRelative"] = args ? args.endDateRelative : undefined;
            inputs["keepers"] = args ? args.keepers : undefined;
            inputs["startDate"] = args ? args.startDate : undefined;
            inputs["keyId"] = undefined /*out*/;
            inputs["value"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ApplicationPassword.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationPassword resources.
 */
export interface ApplicationPasswordState {
    /**
     * The object ID of the application for which this password should be created. Changing this field forces a new resource to be created.
     */
    applicationObjectId?: pulumi.Input<string>;
    /**
     * A display name for the password.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
     */
    endDate?: pulumi.Input<string>;
    /**
     * A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
     */
    endDateRelative?: pulumi.Input<string>;
    /**
     * Arbitrary map of values that, when changed, will trigger rotation of the password
     */
    keepers?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A UUID used to uniquely identify this password credential.
     */
    keyId?: pulumi.Input<string>;
    /**
     * The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
     */
    startDate?: pulumi.Input<string>;
    /**
     * The password for this application, which is generated by Azure Active Directory.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApplicationPassword resource.
 */
export interface ApplicationPasswordArgs {
    /**
     * The object ID of the application for which this password should be created. Changing this field forces a new resource to be created.
     */
    applicationObjectId: pulumi.Input<string>;
    /**
     * A display name for the password.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
     */
    endDate?: pulumi.Input<string>;
    /**
     * A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
     */
    endDateRelative?: pulumi.Input<string>;
    /**
     * Arbitrary map of values that, when changed, will trigger rotation of the password
     */
    keepers?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
     */
    startDate?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a password credential associated with a service principal within Azure Active Directory. See also the azuread.ApplicationPassword resource.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.All` or `Directory.ReadWrite.All`
 *
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`
 *
 * ## Import
 *
 * This resource does not support importing.
 */
export class ServicePrincipalPassword extends pulumi.CustomResource {
    /**
     * Get an existing ServicePrincipalPassword resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServicePrincipalPasswordState, opts?: pulumi.CustomResourceOptions): ServicePrincipalPassword {
        return new ServicePrincipalPassword(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/servicePrincipalPassword:ServicePrincipalPassword';

    /**
     * Returns true if the given object is an instance of ServicePrincipalPassword.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServicePrincipalPassword {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServicePrincipalPassword.__pulumiType;
    }

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
     * A UUID used to uniquely identify this password credential.
     */
    public /*out*/ readonly keyId!: pulumi.Output<string>;
    /**
     * A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
     */
    public readonly rotateWhenChanged!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The object ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
     */
    public readonly servicePrincipalId!: pulumi.Output<string>;
    /**
     * The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
     */
    public readonly startDate!: pulumi.Output<string>;
    /**
     * The password for this service principal, which is generated by Azure Active Directory.
     */
    public /*out*/ readonly value!: pulumi.Output<string>;

    /**
     * Create a ServicePrincipalPassword resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServicePrincipalPasswordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServicePrincipalPasswordArgs | ServicePrincipalPasswordState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServicePrincipalPasswordState | undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["endDate"] = state ? state.endDate : undefined;
            resourceInputs["endDateRelative"] = state ? state.endDateRelative : undefined;
            resourceInputs["keyId"] = state ? state.keyId : undefined;
            resourceInputs["rotateWhenChanged"] = state ? state.rotateWhenChanged : undefined;
            resourceInputs["servicePrincipalId"] = state ? state.servicePrincipalId : undefined;
            resourceInputs["startDate"] = state ? state.startDate : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as ServicePrincipalPasswordArgs | undefined;
            if ((!args || args.servicePrincipalId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'servicePrincipalId'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["endDate"] = args ? args.endDate : undefined;
            resourceInputs["endDateRelative"] = args ? args.endDateRelative : undefined;
            resourceInputs["rotateWhenChanged"] = args ? args.rotateWhenChanged : undefined;
            resourceInputs["servicePrincipalId"] = args ? args.servicePrincipalId : undefined;
            resourceInputs["startDate"] = args ? args.startDate : undefined;
            resourceInputs["keyId"] = undefined /*out*/;
            resourceInputs["value"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServicePrincipalPassword.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServicePrincipalPassword resources.
 */
export interface ServicePrincipalPasswordState {
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
     * A UUID used to uniquely identify this password credential.
     */
    keyId?: pulumi.Input<string>;
    /**
     * A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
     */
    rotateWhenChanged?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The object ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
     */
    servicePrincipalId?: pulumi.Input<string>;
    /**
     * The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
     */
    startDate?: pulumi.Input<string>;
    /**
     * The password for this service principal, which is generated by Azure Active Directory.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServicePrincipalPassword resource.
 */
export interface ServicePrincipalPasswordArgs {
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
     * A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
     */
    rotateWhenChanged?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The object ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
     */
    servicePrincipalId: pulumi.Input<string>;
    /**
     * The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
     */
    startDate?: pulumi.Input<string>;
}

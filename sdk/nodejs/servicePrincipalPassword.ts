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
     * The display name for the password.
     */
    public /*out*/ readonly displayName!: pulumi.Output<string>;
    /**
     * The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
     */
    public /*out*/ readonly endDate!: pulumi.Output<string>;
    /**
     * Arbitrary map of values that, when changed, will trigger rotation of the password
     */
    public readonly keepers!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A UUID used to uniquely identify this password credential.
     */
    public /*out*/ readonly keyId!: pulumi.Output<string>;
    /**
     * The object ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
     */
    public readonly servicePrincipalId!: pulumi.Output<string>;
    /**
     * The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
     */
    public /*out*/ readonly startDate!: pulumi.Output<string>;
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServicePrincipalPasswordState | undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["endDate"] = state ? state.endDate : undefined;
            inputs["keepers"] = state ? state.keepers : undefined;
            inputs["keyId"] = state ? state.keyId : undefined;
            inputs["servicePrincipalId"] = state ? state.servicePrincipalId : undefined;
            inputs["startDate"] = state ? state.startDate : undefined;
            inputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as ServicePrincipalPasswordArgs | undefined;
            if ((!args || args.servicePrincipalId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'servicePrincipalId'");
            }
            inputs["keepers"] = args ? args.keepers : undefined;
            inputs["servicePrincipalId"] = args ? args.servicePrincipalId : undefined;
            inputs["displayName"] = undefined /*out*/;
            inputs["endDate"] = undefined /*out*/;
            inputs["keyId"] = undefined /*out*/;
            inputs["startDate"] = undefined /*out*/;
            inputs["value"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ServicePrincipalPassword.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServicePrincipalPassword resources.
 */
export interface ServicePrincipalPasswordState {
    /**
     * The display name for the password.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
     */
    endDate?: pulumi.Input<string>;
    /**
     * Arbitrary map of values that, when changed, will trigger rotation of the password
     */
    keepers?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A UUID used to uniquely identify this password credential.
     */
    keyId?: pulumi.Input<string>;
    /**
     * The object ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
     */
    servicePrincipalId?: pulumi.Input<string>;
    /**
     * The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
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
     * Arbitrary map of values that, when changed, will trigger rotation of the password
     */
    keepers?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The object ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
     */
    servicePrincipalId: pulumi.Input<string>;
}

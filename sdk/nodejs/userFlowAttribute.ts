// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages user flow attributes in an Azure Active Directory (Azure AD) tenant.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires the following application role: `IdentityUserFlow.ReadWrite.All`
 *
 * ## Example Usage
 *
 * *Basic example*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = new azuread.UserFlowAttribute("example", {
 *     displayName: "Hobby",
 *     description: "Your hobby",
 *     dataType: "string",
 * });
 * ```
 *
 * ## Import
 *
 * User flow attributes can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import azuread:index/userFlowAttribute:UserFlowAttribute example extension_ecc9f88db2924942b8a96f44873616fe_Hobbyjkorv
 * ```
 *
 *  -> This ID can be queried using the [User Flow Attributes API](https://learn.microsoft.com/en-us/graph/api/identityuserflowattribute-list?view=graph-rest-1.0&tabs=http).
 */
export class UserFlowAttribute extends pulumi.CustomResource {
    /**
     * Get an existing UserFlowAttribute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserFlowAttributeState, opts?: pulumi.CustomResourceOptions): UserFlowAttribute {
        return new UserFlowAttribute(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/userFlowAttribute:UserFlowAttribute';

    /**
     * Returns true if the given object is an instance of UserFlowAttribute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserFlowAttribute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserFlowAttribute.__pulumiType;
    }

    /**
     * The type of the user flow attribute. Values include `builtIn`, `custom` or `required`.
     */
    public /*out*/ readonly attributeType!: pulumi.Output<string>;
    /**
     * The data type of the user flow attribute. Possible values are `boolean`, `dateTime`, `int64`, `string` or `stringCollection`. Changing this forces a new resource to be created.
     */
    public readonly dataType!: pulumi.Output<string>;
    /**
     * The description of the user flow attribute that is shown to the user at the time of sign-up.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The display name of the user flow attribute. Changing this forces a new resource to be created.
     */
    public readonly displayName!: pulumi.Output<string>;

    /**
     * Create a UserFlowAttribute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserFlowAttributeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserFlowAttributeArgs | UserFlowAttributeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserFlowAttributeState | undefined;
            resourceInputs["attributeType"] = state ? state.attributeType : undefined;
            resourceInputs["dataType"] = state ? state.dataType : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
        } else {
            const args = argsOrState as UserFlowAttributeArgs | undefined;
            if ((!args || args.dataType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataType'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["dataType"] = args ? args.dataType : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["attributeType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserFlowAttribute.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserFlowAttribute resources.
 */
export interface UserFlowAttributeState {
    /**
     * The type of the user flow attribute. Values include `builtIn`, `custom` or `required`.
     */
    attributeType?: pulumi.Input<string>;
    /**
     * The data type of the user flow attribute. Possible values are `boolean`, `dateTime`, `int64`, `string` or `stringCollection`. Changing this forces a new resource to be created.
     */
    dataType?: pulumi.Input<string>;
    /**
     * The description of the user flow attribute that is shown to the user at the time of sign-up.
     */
    description?: pulumi.Input<string>;
    /**
     * The display name of the user flow attribute. Changing this forces a new resource to be created.
     */
    displayName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserFlowAttribute resource.
 */
export interface UserFlowAttributeArgs {
    /**
     * The data type of the user flow attribute. Possible values are `boolean`, `dateTime`, `int64`, `string` or `stringCollection`. Changing this forces a new resource to be created.
     */
    dataType: pulumi.Input<string>;
    /**
     * The description of the user flow attribute that is shown to the user at the time of sign-up.
     */
    description: pulumi.Input<string>;
    /**
     * The display name of the user flow attribute. Changing this forces a new resource to be created.
     */
    displayName: pulumi.Input<string>;
}

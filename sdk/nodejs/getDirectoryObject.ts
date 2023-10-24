// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Retrieves the OData type for a generic directory object having the provided object ID.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this data source.
 *
 * When authenticated with a service principal, this data source requires either `User.Read.All`, `Group.Read.All` or `Directory.Read.All`, depending on the type of object being queried.
 *
 * When authenticated with a user principal, this data source does not require any additional roles.
 *
 * ## Attributes Reference
 *
 * The following attributes are exported:
 *
 * *`objectId` - The object ID of the directory object.
 * *`type` - The shortened OData type of the directory object. Possible values include: `Group`, `User` or `ServicePrincipal`.
 */
export function getDirectoryObject(args: GetDirectoryObjectArgs, opts?: pulumi.InvokeOptions): Promise<GetDirectoryObjectResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuread:index/getDirectoryObject:getDirectoryObject", {
        "objectId": args.objectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDirectoryObject.
 */
export interface GetDirectoryObjectArgs {
    /**
     * Specifies the Object ID of the directory object to look up.
     */
    objectId: string;
}

/**
 * A collection of values returned by getDirectoryObject.
 */
export interface GetDirectoryObjectResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly objectId: string;
    readonly type: string;
}
/**
 * Retrieves the OData type for a generic directory object having the provided object ID.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this data source.
 *
 * When authenticated with a service principal, this data source requires either `User.Read.All`, `Group.Read.All` or `Directory.Read.All`, depending on the type of object being queried.
 *
 * When authenticated with a user principal, this data source does not require any additional roles.
 *
 * ## Attributes Reference
 *
 * The following attributes are exported:
 *
 * *`objectId` - The object ID of the directory object.
 * *`type` - The shortened OData type of the directory object. Possible values include: `Group`, `User` or `ServicePrincipal`.
 */
export function getDirectoryObjectOutput(args: GetDirectoryObjectOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDirectoryObjectResult> {
    return pulumi.output(args).apply((a: any) => getDirectoryObject(a, opts))
}

/**
 * A collection of arguments for invoking getDirectoryObject.
 */
export interface GetDirectoryObjectOutputArgs {
    /**
     * Specifies the Object ID of the directory object to look up.
     */
    objectId: pulumi.Input<string>;
}

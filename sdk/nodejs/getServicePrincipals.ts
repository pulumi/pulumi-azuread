// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets basic information for multiple Azure Active Directory service principals.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this data source.
 *
 * When authenticated with a service principal, this data source requires one of the following application roles: `Application.Read.All` or `Directory.Read.All`
 *
 * When authenticated with a user principal, this data source does not require any additional roles.
 *
 * ## Example Usage
 *
 * *Look up by application display names*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getServicePrincipals({
 *     displayNames: [
 *         "example-app",
 *         "another-app",
 *     ],
 * });
 * ```
 *
 * *Look up by application IDs (client IDs*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getServicePrincipals({
 *     clientIds: [
 *         "11111111-0000-0000-0000-000000000000",
 *         "22222222-0000-0000-0000-000000000000",
 *         "33333333-0000-0000-0000-000000000000",
 *     ],
 * });
 * ```
 *
 * *Look up by service principal object IDs*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getServicePrincipals({
 *     objectIds: [
 *         "00000000-0000-0000-0000-000000000000",
 *         "00000000-0000-0000-0000-111111111111",
 *         "00000000-0000-0000-0000-222222222222",
 *     ],
 * });
 * ```
 */
export function getServicePrincipals(args?: GetServicePrincipalsArgs, opts?: pulumi.InvokeOptions): Promise<GetServicePrincipalsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuread:index/getServicePrincipals:getServicePrincipals", {
        "applicationIds": args.applicationIds,
        "clientIds": args.clientIds,
        "displayNames": args.displayNames,
        "ignoreMissing": args.ignoreMissing,
        "objectIds": args.objectIds,
        "returnAll": args.returnAll,
    }, opts);
}

/**
 * A collection of arguments for invoking getServicePrincipals.
 */
export interface GetServicePrincipalsArgs {
    /**
     * A list of client IDs of the applications associated with the service principals.
     *
     * @deprecated The `application_ids` property has been replaced with the `client_ids` property and will be removed in version 3.0 of the AzureAD provider
     */
    applicationIds?: string[];
    /**
     * A list of client IDs of the applications associated with the service principals.
     */
    clientIds?: string[];
    /**
     * A list of display names of the applications associated with the service principals.
     */
    displayNames?: string[];
    /**
     * Ignore missing service principals and return all service principals that are found. The data source will still fail if no service principals are found. Defaults to false.
     */
    ignoreMissing?: boolean;
    /**
     * The object IDs of the service principals.
     */
    objectIds?: string[];
    /**
     * When `true`, the data source will return all service principals. Cannot be used with `ignoreMissing`. Defaults to false.
     *
     * > Either `returnAll`, or one of `clientIds`, `displayNames` or `objectIds` must be specified. These _may_ be specified as an empty list, in which case no results will be returned.
     */
    returnAll?: boolean;
}

/**
 * A collection of values returned by getServicePrincipals.
 */
export interface GetServicePrincipalsResult {
    /**
     * A list of client IDs of the applications associated with the service principals.
     *
     * @deprecated The `application_ids` property has been replaced with the `client_ids` property and will be removed in version 3.0 of the AzureAD provider
     */
    readonly applicationIds: string[];
    /**
     * The client ID of the application associated with this service principal.
     */
    readonly clientIds: string[];
    /**
     * A list of display names of the applications associated with the service principals.
     */
    readonly displayNames: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ignoreMissing?: boolean;
    /**
     * The object IDs of the service principals.
     */
    readonly objectIds: string[];
    readonly returnAll?: boolean;
    /**
     * A list of service principals. Each `servicePrincipal` object provides the attributes documented below.
     */
    readonly servicePrincipals: outputs.GetServicePrincipalsServicePrincipal[];
}
/**
 * Gets basic information for multiple Azure Active Directory service principals.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this data source.
 *
 * When authenticated with a service principal, this data source requires one of the following application roles: `Application.Read.All` or `Directory.Read.All`
 *
 * When authenticated with a user principal, this data source does not require any additional roles.
 *
 * ## Example Usage
 *
 * *Look up by application display names*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getServicePrincipals({
 *     displayNames: [
 *         "example-app",
 *         "another-app",
 *     ],
 * });
 * ```
 *
 * *Look up by application IDs (client IDs*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getServicePrincipals({
 *     clientIds: [
 *         "11111111-0000-0000-0000-000000000000",
 *         "22222222-0000-0000-0000-000000000000",
 *         "33333333-0000-0000-0000-000000000000",
 *     ],
 * });
 * ```
 *
 * *Look up by service principal object IDs*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getServicePrincipals({
 *     objectIds: [
 *         "00000000-0000-0000-0000-000000000000",
 *         "00000000-0000-0000-0000-111111111111",
 *         "00000000-0000-0000-0000-222222222222",
 *     ],
 * });
 * ```
 */
export function getServicePrincipalsOutput(args?: GetServicePrincipalsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServicePrincipalsResult> {
    return pulumi.output(args).apply((a: any) => getServicePrincipals(a, opts))
}

/**
 * A collection of arguments for invoking getServicePrincipals.
 */
export interface GetServicePrincipalsOutputArgs {
    /**
     * A list of client IDs of the applications associated with the service principals.
     *
     * @deprecated The `application_ids` property has been replaced with the `client_ids` property and will be removed in version 3.0 of the AzureAD provider
     */
    applicationIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of client IDs of the applications associated with the service principals.
     */
    clientIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of display names of the applications associated with the service principals.
     */
    displayNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Ignore missing service principals and return all service principals that are found. The data source will still fail if no service principals are found. Defaults to false.
     */
    ignoreMissing?: pulumi.Input<boolean>;
    /**
     * The object IDs of the service principals.
     */
    objectIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When `true`, the data source will return all service principals. Cannot be used with `ignoreMissing`. Defaults to false.
     *
     * > Either `returnAll`, or one of `clientIds`, `displayNames` or `objectIds` must be specified. These _may_ be specified as an empty list, in which case no results will be returned.
     */
    returnAll?: pulumi.Input<boolean>;
}

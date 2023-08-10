// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about a Named Location within Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this data source.
 *
 * When authenticated with a service principal, this resource requires the following application roles: `Policy.Read.All`
 *
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Conditional Access Administrator` or `Global Reader`
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getNamedLocation({
 *     displayName: "My Named Location",
 * });
 * ```
 * ## Attributes Reference
 *
 * The following attributes are exported:
 *
 * * `country` - A `country` block as documented below, which describes a country-based named location.
 * * `id` - The ID of the named location.
 * * `ip` - An `ip` block as documented below, which describes an IP-based named location.
 * * 
 * ***
 *
 * `country` block exports the following:
 *
 * * `countriesAndRegions` - List of countries and/or regions in two-letter format specified by ISO 3166-2.
 * * `includeUnknownCountriesAndRegions` - Whether IP addresses that don't map to a country or region are included in the named location.
 *
 * ***
 *
 * `ip` block exports the following:
 *
 * * `ipRanges` - List of IP address ranges in IPv4 CIDR format (e.g. `1.2.3.4/32`) or any allowable IPv6 format from IETF RFC596.
 * * `trusted` - Whether the named location is trusted.
 */
export function getNamedLocation(args: GetNamedLocationArgs, opts?: pulumi.InvokeOptions): Promise<GetNamedLocationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuread:index/getNamedLocation:getNamedLocation", {
        "displayName": args.displayName,
    }, opts);
}

/**
 * A collection of arguments for invoking getNamedLocation.
 */
export interface GetNamedLocationArgs {
    /**
     * Specifies the display named of the named location to look up.
     */
    displayName: string;
}

/**
 * A collection of values returned by getNamedLocation.
 */
export interface GetNamedLocationResult {
    readonly countries: outputs.GetNamedLocationCountry[];
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ips: outputs.GetNamedLocationIp[];
}
/**
 * Gets information about a Named Location within Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this data source.
 *
 * When authenticated with a service principal, this resource requires the following application roles: `Policy.Read.All`
 *
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Conditional Access Administrator` or `Global Reader`
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getNamedLocation({
 *     displayName: "My Named Location",
 * });
 * ```
 * ## Attributes Reference
 *
 * The following attributes are exported:
 *
 * * `country` - A `country` block as documented below, which describes a country-based named location.
 * * `id` - The ID of the named location.
 * * `ip` - An `ip` block as documented below, which describes an IP-based named location.
 * * 
 * ***
 *
 * `country` block exports the following:
 *
 * * `countriesAndRegions` - List of countries and/or regions in two-letter format specified by ISO 3166-2.
 * * `includeUnknownCountriesAndRegions` - Whether IP addresses that don't map to a country or region are included in the named location.
 *
 * ***
 *
 * `ip` block exports the following:
 *
 * * `ipRanges` - List of IP address ranges in IPv4 CIDR format (e.g. `1.2.3.4/32`) or any allowable IPv6 format from IETF RFC596.
 * * `trusted` - Whether the named location is trusted.
 */
export function getNamedLocationOutput(args: GetNamedLocationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNamedLocationResult> {
    return pulumi.output(args).apply((a: any) => getNamedLocation(a, opts))
}

/**
 * A collection of arguments for invoking getNamedLocation.
 */
export interface GetNamedLocationOutputArgs {
    /**
     * Specifies the display named of the named location to look up.
     */
    displayName: pulumi.Input<string>;
}
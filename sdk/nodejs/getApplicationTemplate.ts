// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about an Application Template from the [Azure AD App Gallery](https://azuremarketplace.microsoft.com/en-US/marketplace/apps/category/azure-active-directory-apps).
 *
 * ## API Permissions
 *
 * This data source does not require any additional roles.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getApplicationTemplate({
 *     displayName: "Marketo",
 * });
 * export const applicationTemplateId = example.then(example => example.templateId);
 * ```
 */
export function getApplicationTemplate(args?: GetApplicationTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationTemplateResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("azuread:index/getApplicationTemplate:getApplicationTemplate", {
        "displayName": args.displayName,
        "templateId": args.templateId,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplicationTemplate.
 */
export interface GetApplicationTemplateArgs {
    /**
     * Specifies the display name of the templated application.
     */
    displayName?: string;
    /**
     * Specifies the ID of the templated application.
     *
     * > One of `templateId` or `displayName` must be specified.
     */
    templateId?: string;
}

/**
 * A collection of values returned by getApplicationTemplate.
 */
export interface GetApplicationTemplateResult {
    /**
     * List of categories for this templated application.
     */
    readonly categories: string[];
    /**
     * The display name for the templated application.
     */
    readonly displayName: string;
    /**
     * Home page URL of the templated application.
     */
    readonly homepageUrl: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * URL to retrieve the logo for this templated application.
     */
    readonly logoUrl: string;
    /**
     * Name of the publisher for this templated application.
     */
    readonly publisher: string;
    /**
     * List of provisioning modes supported by this templated application.
     */
    readonly supportedProvisioningTypes: string[];
    /**
     * List of single sign on modes supported by this templated application.
     */
    readonly supportedSingleSignOnModes: string[];
    /**
     * The ID of the templated application.
     */
    readonly templateId: string;
}
/**
 * Use this data source to access information about an Application Template from the [Azure AD App Gallery](https://azuremarketplace.microsoft.com/en-US/marketplace/apps/category/azure-active-directory-apps).
 *
 * ## API Permissions
 *
 * This data source does not require any additional roles.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example = azuread.getApplicationTemplate({
 *     displayName: "Marketo",
 * });
 * export const applicationTemplateId = example.then(example => example.templateId);
 * ```
 */
export function getApplicationTemplateOutput(args?: GetApplicationTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationTemplateResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("azuread:index/getApplicationTemplate:getApplicationTemplate", {
        "displayName": args.displayName,
        "templateId": args.templateId,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplicationTemplate.
 */
export interface GetApplicationTemplateOutputArgs {
    /**
     * Specifies the display name of the templated application.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Specifies the ID of the templated application.
     *
     * > One of `templateId` or `displayName` must be specified.
     */
    templateId?: pulumi.Input<string>;
}

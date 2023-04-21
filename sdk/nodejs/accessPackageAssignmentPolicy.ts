// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manages an assignment policy for an access package within Identity Governance in Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.
 *
 * When authenticated with a user principal, this resource requires `Global Administrator` directory role, or one of the `Catalog Owner` and `Access Package Manager` role in Idneity Governance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const exampleGroup = new azuread.Group("exampleGroup", {
 *     displayName: "group-name",
 *     securityEnabled: true,
 * });
 * const exampleAccessPackageCatalog = new azuread.AccessPackageCatalog("exampleAccessPackageCatalog", {
 *     displayName: "example-catalog",
 *     description: "Example catalog",
 * });
 * const exampleAccessPackage = new azuread.AccessPackage("exampleAccessPackage", {
 *     catalogId: exampleAccessPackageCatalog.id,
 *     displayName: "access-package",
 *     description: "Access Package",
 * });
 * const test = new azuread.AccessPackageAssignmentPolicy("test", {
 *     accessPackageId: azuread_access_package.test.id,
 *     displayName: "assignment-policy",
 *     description: "My assignment policy",
 *     durationInDays: 90,
 *     requestorSettings: {
 *         scopeType: "AllExistingDirectoryMemberUsers",
 *     },
 *     approvalSettings: {
 *         approvalRequired: true,
 *         approvalStages: [{
 *             approvalTimeoutInDays: 14,
 *             primaryApprovers: [{
 *                 objectId: azuread_group.test.object_id,
 *                 subjectType: "groupMembers",
 *             }],
 *         }],
 *     },
 *     assignmentReviewSettings: {
 *         enabled: true,
 *         reviewFrequency: "weekly",
 *         durationInDays: 3,
 *         reviewType: "Self",
 *         accessReviewTimeoutBehavior: "keepAccess",
 *     },
 *     questions: [{
 *         text: {
 *             defaultText: "hello, how are you?",
 *         },
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * An access package assignment policy can be imported using the ID, e.g.
 *
 * ```sh
 *  $ pulumi import azuread:index/accessPackageAssignmentPolicy:AccessPackageAssignmentPolicy example 00000000-0000-0000-0000-000000000000
 * ```
 */
export class AccessPackageAssignmentPolicy extends pulumi.CustomResource {
    /**
     * Get an existing AccessPackageAssignmentPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessPackageAssignmentPolicyState, opts?: pulumi.CustomResourceOptions): AccessPackageAssignmentPolicy {
        return new AccessPackageAssignmentPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/accessPackageAssignmentPolicy:AccessPackageAssignmentPolicy';

    /**
     * Returns true if the given object is an instance of AccessPackageAssignmentPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessPackageAssignmentPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessPackageAssignmentPolicy.__pulumiType;
    }

    /**
     * The ID of the access package that will contain the policy.
     */
    public readonly accessPackageId!: pulumi.Output<string>;
    /**
     * An `approvalSettings` block to specify whether approvals are required and how they are obtained, as documented below.
     */
    public readonly approvalSettings!: pulumi.Output<outputs.AccessPackageAssignmentPolicyApprovalSettings | undefined>;
    /**
     * An `assignmentReviewSettings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
     */
    public readonly assignmentReviewSettings!: pulumi.Output<outputs.AccessPackageAssignmentPolicyAssignmentReviewSettings | undefined>;
    /**
     * The description of the policy.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The display name of the policy.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * How many days this assignment is valid for.
     */
    public readonly durationInDays!: pulumi.Output<number | undefined>;
    /**
     * The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
     */
    public readonly expirationDate!: pulumi.Output<string | undefined>;
    /**
     * Whether users will be able to request extension of their access to this package before their access expires.
     */
    public readonly extensionEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * One or more `question` blocks for the requestor, as documented below.
     */
    public readonly questions!: pulumi.Output<outputs.AccessPackageAssignmentPolicyQuestion[] | undefined>;
    /**
     * A `requestorSettings` block to configure the users who can request access, as documented below.
     */
    public readonly requestorSettings!: pulumi.Output<outputs.AccessPackageAssignmentPolicyRequestorSettings | undefined>;

    /**
     * Create a AccessPackageAssignmentPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessPackageAssignmentPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessPackageAssignmentPolicyArgs | AccessPackageAssignmentPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccessPackageAssignmentPolicyState | undefined;
            resourceInputs["accessPackageId"] = state ? state.accessPackageId : undefined;
            resourceInputs["approvalSettings"] = state ? state.approvalSettings : undefined;
            resourceInputs["assignmentReviewSettings"] = state ? state.assignmentReviewSettings : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["durationInDays"] = state ? state.durationInDays : undefined;
            resourceInputs["expirationDate"] = state ? state.expirationDate : undefined;
            resourceInputs["extensionEnabled"] = state ? state.extensionEnabled : undefined;
            resourceInputs["questions"] = state ? state.questions : undefined;
            resourceInputs["requestorSettings"] = state ? state.requestorSettings : undefined;
        } else {
            const args = argsOrState as AccessPackageAssignmentPolicyArgs | undefined;
            if ((!args || args.accessPackageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessPackageId'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["accessPackageId"] = args ? args.accessPackageId : undefined;
            resourceInputs["approvalSettings"] = args ? args.approvalSettings : undefined;
            resourceInputs["assignmentReviewSettings"] = args ? args.assignmentReviewSettings : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["durationInDays"] = args ? args.durationInDays : undefined;
            resourceInputs["expirationDate"] = args ? args.expirationDate : undefined;
            resourceInputs["extensionEnabled"] = args ? args.extensionEnabled : undefined;
            resourceInputs["questions"] = args ? args.questions : undefined;
            resourceInputs["requestorSettings"] = args ? args.requestorSettings : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccessPackageAssignmentPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessPackageAssignmentPolicy resources.
 */
export interface AccessPackageAssignmentPolicyState {
    /**
     * The ID of the access package that will contain the policy.
     */
    accessPackageId?: pulumi.Input<string>;
    /**
     * An `approvalSettings` block to specify whether approvals are required and how they are obtained, as documented below.
     */
    approvalSettings?: pulumi.Input<inputs.AccessPackageAssignmentPolicyApprovalSettings>;
    /**
     * An `assignmentReviewSettings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
     */
    assignmentReviewSettings?: pulumi.Input<inputs.AccessPackageAssignmentPolicyAssignmentReviewSettings>;
    /**
     * The description of the policy.
     */
    description?: pulumi.Input<string>;
    /**
     * The display name of the policy.
     */
    displayName?: pulumi.Input<string>;
    /**
     * How many days this assignment is valid for.
     */
    durationInDays?: pulumi.Input<number>;
    /**
     * The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
     */
    expirationDate?: pulumi.Input<string>;
    /**
     * Whether users will be able to request extension of their access to this package before their access expires.
     */
    extensionEnabled?: pulumi.Input<boolean>;
    /**
     * One or more `question` blocks for the requestor, as documented below.
     */
    questions?: pulumi.Input<pulumi.Input<inputs.AccessPackageAssignmentPolicyQuestion>[]>;
    /**
     * A `requestorSettings` block to configure the users who can request access, as documented below.
     */
    requestorSettings?: pulumi.Input<inputs.AccessPackageAssignmentPolicyRequestorSettings>;
}

/**
 * The set of arguments for constructing a AccessPackageAssignmentPolicy resource.
 */
export interface AccessPackageAssignmentPolicyArgs {
    /**
     * The ID of the access package that will contain the policy.
     */
    accessPackageId: pulumi.Input<string>;
    /**
     * An `approvalSettings` block to specify whether approvals are required and how they are obtained, as documented below.
     */
    approvalSettings?: pulumi.Input<inputs.AccessPackageAssignmentPolicyApprovalSettings>;
    /**
     * An `assignmentReviewSettings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
     */
    assignmentReviewSettings?: pulumi.Input<inputs.AccessPackageAssignmentPolicyAssignmentReviewSettings>;
    /**
     * The description of the policy.
     */
    description: pulumi.Input<string>;
    /**
     * The display name of the policy.
     */
    displayName: pulumi.Input<string>;
    /**
     * How many days this assignment is valid for.
     */
    durationInDays?: pulumi.Input<number>;
    /**
     * The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
     */
    expirationDate?: pulumi.Input<string>;
    /**
     * Whether users will be able to request extension of their access to this package before their access expires.
     */
    extensionEnabled?: pulumi.Input<boolean>;
    /**
     * One or more `question` blocks for the requestor, as documented below.
     */
    questions?: pulumi.Input<pulumi.Input<inputs.AccessPackageAssignmentPolicyQuestion>[]>;
    /**
     * A `requestorSettings` block to configure the users who can request access, as documented below.
     */
    requestorSettings?: pulumi.Input<inputs.AccessPackageAssignmentPolicyRequestorSettings>;
}
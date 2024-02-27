// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    /// <summary>
    /// Manages an assignment policy for an access package within Identity Governance in Azure Active Directory.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.
    /// 
    /// When authenticated with a user principal, this resource requires `Global Administrator` directory role, or one of the `Catalog Owner` and `Access Package Manager` role in Identity Governance.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureAD.Group("example", new()
    ///     {
    ///         DisplayName = "group-name",
    ///         SecurityEnabled = true,
    ///     });
    /// 
    ///     var exampleAccessPackageCatalog = new AzureAD.AccessPackageCatalog("example", new()
    ///     {
    ///         DisplayName = "example-catalog",
    ///         Description = "Example catalog",
    ///     });
    /// 
    ///     var exampleAccessPackage = new AzureAD.AccessPackage("example", new()
    ///     {
    ///         CatalogId = exampleAccessPackageCatalog.Id,
    ///         DisplayName = "access-package",
    ///         Description = "Access Package",
    ///     });
    /// 
    ///     var exampleAccessPackageAssignmentPolicy = new AzureAD.AccessPackageAssignmentPolicy("example", new()
    ///     {
    ///         AccessPackageId = exampleAccessPackage.Id,
    ///         DisplayName = "assignment-policy",
    ///         Description = "My assignment policy",
    ///         DurationInDays = 90,
    ///         RequestorSettings = new AzureAD.Inputs.AccessPackageAssignmentPolicyRequestorSettingsArgs
    ///         {
    ///             ScopeType = "AllExistingDirectoryMemberUsers",
    ///         },
    ///         ApprovalSettings = new AzureAD.Inputs.AccessPackageAssignmentPolicyApprovalSettingsArgs
    ///         {
    ///             ApprovalRequired = true,
    ///             ApprovalStages = new[]
    ///             {
    ///                 new AzureAD.Inputs.AccessPackageAssignmentPolicyApprovalSettingsApprovalStageArgs
    ///                 {
    ///                     ApprovalTimeoutInDays = 14,
    ///                     PrimaryApprovers = new[]
    ///                     {
    ///                         new AzureAD.Inputs.AccessPackageAssignmentPolicyApprovalSettingsApprovalStagePrimaryApproverArgs
    ///                         {
    ///                             ObjectId = example.ObjectId,
    ///                             SubjectType = "groupMembers",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         AssignmentReviewSettings = new AzureAD.Inputs.AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs
    ///         {
    ///             Enabled = true,
    ///             ReviewFrequency = "weekly",
    ///             DurationInDays = 3,
    ///             ReviewType = "Self",
    ///             AccessReviewTimeoutBehavior = "keepAccess",
    ///         },
    ///         Questions = new[]
    ///         {
    ///             new AzureAD.Inputs.AccessPackageAssignmentPolicyQuestionArgs
    ///             {
    ///                 Text = new AzureAD.Inputs.AccessPackageAssignmentPolicyQuestionTextArgs
    ///                 {
    ///                     DefaultText = "hello, how are you?",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// An access package assignment policy can be imported using the ID, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuread:index/accessPackageAssignmentPolicy:AccessPackageAssignmentPolicy example 00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureADResourceType("azuread:index/accessPackageAssignmentPolicy:AccessPackageAssignmentPolicy")]
    public partial class AccessPackageAssignmentPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the access package that will contain the policy.
        /// </summary>
        [Output("accessPackageId")]
        public Output<string> AccessPackageId { get; private set; } = null!;

        /// <summary>
        /// An `approval_settings` block to specify whether approvals are required and how they are obtained, as documented below.
        /// </summary>
        [Output("approvalSettings")]
        public Output<Outputs.AccessPackageAssignmentPolicyApprovalSettings?> ApprovalSettings { get; private set; } = null!;

        /// <summary>
        /// An `assignment_review_settings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
        /// </summary>
        [Output("assignmentReviewSettings")]
        public Output<Outputs.AccessPackageAssignmentPolicyAssignmentReviewSettings?> AssignmentReviewSettings { get; private set; } = null!;

        /// <summary>
        /// The description of the policy.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the policy.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// How many days this assignment is valid for.
        /// </summary>
        [Output("durationInDays")]
        public Output<int?> DurationInDays { get; private set; } = null!;

        /// <summary>
        /// The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
        /// </summary>
        [Output("expirationDate")]
        public Output<string?> ExpirationDate { get; private set; } = null!;

        /// <summary>
        /// Whether users will be able to request extension of their access to this package before their access expires.
        /// </summary>
        [Output("extensionEnabled")]
        public Output<bool?> ExtensionEnabled { get; private set; } = null!;

        /// <summary>
        /// One or more `question` blocks for the requestor, as documented below.
        /// </summary>
        [Output("questions")]
        public Output<ImmutableArray<Outputs.AccessPackageAssignmentPolicyQuestion>> Questions { get; private set; } = null!;

        /// <summary>
        /// A `requestor_settings` block to configure the users who can request access, as documented below.
        /// </summary>
        [Output("requestorSettings")]
        public Output<Outputs.AccessPackageAssignmentPolicyRequestorSettings?> RequestorSettings { get; private set; } = null!;


        /// <summary>
        /// Create a AccessPackageAssignmentPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessPackageAssignmentPolicy(string name, AccessPackageAssignmentPolicyArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/accessPackageAssignmentPolicy:AccessPackageAssignmentPolicy", name, args ?? new AccessPackageAssignmentPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessPackageAssignmentPolicy(string name, Input<string> id, AccessPackageAssignmentPolicyState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/accessPackageAssignmentPolicy:AccessPackageAssignmentPolicy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AccessPackageAssignmentPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessPackageAssignmentPolicy Get(string name, Input<string> id, AccessPackageAssignmentPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new AccessPackageAssignmentPolicy(name, id, state, options);
        }
    }

    public sealed class AccessPackageAssignmentPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the access package that will contain the policy.
        /// </summary>
        [Input("accessPackageId", required: true)]
        public Input<string> AccessPackageId { get; set; } = null!;

        /// <summary>
        /// An `approval_settings` block to specify whether approvals are required and how they are obtained, as documented below.
        /// </summary>
        [Input("approvalSettings")]
        public Input<Inputs.AccessPackageAssignmentPolicyApprovalSettingsArgs>? ApprovalSettings { get; set; }

        /// <summary>
        /// An `assignment_review_settings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
        /// </summary>
        [Input("assignmentReviewSettings")]
        public Input<Inputs.AccessPackageAssignmentPolicyAssignmentReviewSettingsArgs>? AssignmentReviewSettings { get; set; }

        /// <summary>
        /// The description of the policy.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The display name of the policy.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// How many days this assignment is valid for.
        /// </summary>
        [Input("durationInDays")]
        public Input<int>? DurationInDays { get; set; }

        /// <summary>
        /// The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
        /// </summary>
        [Input("expirationDate")]
        public Input<string>? ExpirationDate { get; set; }

        /// <summary>
        /// Whether users will be able to request extension of their access to this package before their access expires.
        /// </summary>
        [Input("extensionEnabled")]
        public Input<bool>? ExtensionEnabled { get; set; }

        [Input("questions")]
        private InputList<Inputs.AccessPackageAssignmentPolicyQuestionArgs>? _questions;

        /// <summary>
        /// One or more `question` blocks for the requestor, as documented below.
        /// </summary>
        public InputList<Inputs.AccessPackageAssignmentPolicyQuestionArgs> Questions
        {
            get => _questions ?? (_questions = new InputList<Inputs.AccessPackageAssignmentPolicyQuestionArgs>());
            set => _questions = value;
        }

        /// <summary>
        /// A `requestor_settings` block to configure the users who can request access, as documented below.
        /// </summary>
        [Input("requestorSettings")]
        public Input<Inputs.AccessPackageAssignmentPolicyRequestorSettingsArgs>? RequestorSettings { get; set; }

        public AccessPackageAssignmentPolicyArgs()
        {
        }
        public static new AccessPackageAssignmentPolicyArgs Empty => new AccessPackageAssignmentPolicyArgs();
    }

    public sealed class AccessPackageAssignmentPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the access package that will contain the policy.
        /// </summary>
        [Input("accessPackageId")]
        public Input<string>? AccessPackageId { get; set; }

        /// <summary>
        /// An `approval_settings` block to specify whether approvals are required and how they are obtained, as documented below.
        /// </summary>
        [Input("approvalSettings")]
        public Input<Inputs.AccessPackageAssignmentPolicyApprovalSettingsGetArgs>? ApprovalSettings { get; set; }

        /// <summary>
        /// An `assignment_review_settings` block, to specify whether assignment review is needed and how it is conducted, as documented below.
        /// </summary>
        [Input("assignmentReviewSettings")]
        public Input<Inputs.AccessPackageAssignmentPolicyAssignmentReviewSettingsGetArgs>? AssignmentReviewSettings { get; set; }

        /// <summary>
        /// The description of the policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the policy.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// How many days this assignment is valid for.
        /// </summary>
        [Input("durationInDays")]
        public Input<int>? DurationInDays { get; set; }

        /// <summary>
        /// The date that this assignment expires, formatted as an RFC3339 date string in UTC(e.g. 2018-01-01T01:02:03Z).
        /// </summary>
        [Input("expirationDate")]
        public Input<string>? ExpirationDate { get; set; }

        /// <summary>
        /// Whether users will be able to request extension of their access to this package before their access expires.
        /// </summary>
        [Input("extensionEnabled")]
        public Input<bool>? ExtensionEnabled { get; set; }

        [Input("questions")]
        private InputList<Inputs.AccessPackageAssignmentPolicyQuestionGetArgs>? _questions;

        /// <summary>
        /// One or more `question` blocks for the requestor, as documented below.
        /// </summary>
        public InputList<Inputs.AccessPackageAssignmentPolicyQuestionGetArgs> Questions
        {
            get => _questions ?? (_questions = new InputList<Inputs.AccessPackageAssignmentPolicyQuestionGetArgs>());
            set => _questions = value;
        }

        /// <summary>
        /// A `requestor_settings` block to configure the users who can request access, as documented below.
        /// </summary>
        [Input("requestorSettings")]
        public Input<Inputs.AccessPackageAssignmentPolicyRequestorSettingsGetArgs>? RequestorSettings { get; set; }

        public AccessPackageAssignmentPolicyState()
        {
        }
        public static new AccessPackageAssignmentPolicyState Empty => new AccessPackageAssignmentPolicyState();
    }
}

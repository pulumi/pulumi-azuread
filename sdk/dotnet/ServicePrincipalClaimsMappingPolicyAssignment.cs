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
    /// Manages a Claims Mapping Policy Assignment within Azure Active Directory.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ApplicationConfiguration` and `Policy.Read.All`
    /// 
    /// When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var app = new AzureAD.ServicePrincipalClaimsMappingPolicyAssignment("app", new()
    ///     {
    ///         ClaimsMappingPolicyId = azuread_claims_mapping_policy.My_policy.Id,
    ///         ServicePrincipalId = azuread_service_principal.My_principal.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Claims Mapping Policy can be imported using the `id`, in the form `service-principal-uuid/claimsMappingPolicy/claims-mapping-policy-uuid`, e.g
    /// 
    /// ```sh
    ///  $ pulumi import azuread:index/servicePrincipalClaimsMappingPolicyAssignment:ServicePrincipalClaimsMappingPolicyAssignment app 00000000-0000-0000-0000-000000000000/claimsMappingPolicy/11111111-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureADResourceType("azuread:index/servicePrincipalClaimsMappingPolicyAssignment:ServicePrincipalClaimsMappingPolicyAssignment")]
    public partial class ServicePrincipalClaimsMappingPolicyAssignment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the claims mapping policy to assign.
        /// </summary>
        [Output("claimsMappingPolicyId")]
        public Output<string> ClaimsMappingPolicyId { get; private set; } = null!;

        /// <summary>
        /// The object ID of the service principal for the policy assignment.
        /// </summary>
        [Output("servicePrincipalId")]
        public Output<string> ServicePrincipalId { get; private set; } = null!;


        /// <summary>
        /// Create a ServicePrincipalClaimsMappingPolicyAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServicePrincipalClaimsMappingPolicyAssignment(string name, ServicePrincipalClaimsMappingPolicyAssignmentArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/servicePrincipalClaimsMappingPolicyAssignment:ServicePrincipalClaimsMappingPolicyAssignment", name, args ?? new ServicePrincipalClaimsMappingPolicyAssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServicePrincipalClaimsMappingPolicyAssignment(string name, Input<string> id, ServicePrincipalClaimsMappingPolicyAssignmentState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/servicePrincipalClaimsMappingPolicyAssignment:ServicePrincipalClaimsMappingPolicyAssignment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServicePrincipalClaimsMappingPolicyAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServicePrincipalClaimsMappingPolicyAssignment Get(string name, Input<string> id, ServicePrincipalClaimsMappingPolicyAssignmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ServicePrincipalClaimsMappingPolicyAssignment(name, id, state, options);
        }
    }

    public sealed class ServicePrincipalClaimsMappingPolicyAssignmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the claims mapping policy to assign.
        /// </summary>
        [Input("claimsMappingPolicyId", required: true)]
        public Input<string> ClaimsMappingPolicyId { get; set; } = null!;

        /// <summary>
        /// The object ID of the service principal for the policy assignment.
        /// </summary>
        [Input("servicePrincipalId", required: true)]
        public Input<string> ServicePrincipalId { get; set; } = null!;

        public ServicePrincipalClaimsMappingPolicyAssignmentArgs()
        {
        }
        public static new ServicePrincipalClaimsMappingPolicyAssignmentArgs Empty => new ServicePrincipalClaimsMappingPolicyAssignmentArgs();
    }

    public sealed class ServicePrincipalClaimsMappingPolicyAssignmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the claims mapping policy to assign.
        /// </summary>
        [Input("claimsMappingPolicyId")]
        public Input<string>? ClaimsMappingPolicyId { get; set; }

        /// <summary>
        /// The object ID of the service principal for the policy assignment.
        /// </summary>
        [Input("servicePrincipalId")]
        public Input<string>? ServicePrincipalId { get; set; }

        public ServicePrincipalClaimsMappingPolicyAssignmentState()
        {
        }
        public static new ServicePrincipalClaimsMappingPolicyAssignmentState Empty => new ServicePrincipalClaimsMappingPolicyAssignmentState();
    }
}

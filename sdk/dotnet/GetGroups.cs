// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    public static class GetGroups
    {
        /// <summary>
        /// Gets Object IDs or Display Names for multiple Azure Active Directory groups.
        /// 
        /// &gt; **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var groups = Output.Create(AzureAD.GetGroups.InvokeAsync(new AzureAD.GetGroupsArgs
        ///         {
        ///             Names = 
        ///             {
        ///                 "group-a",
        ///                 "group-b",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGroupsResult> InvokeAsync(GetGroupsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupsResult>("azuread:index/getGroups:getGroups", args ?? new GetGroupsArgs(), options.WithVersion());
    }


    public sealed class GetGroupsArgs : Pulumi.InvokeArgs
    {
        [Input("displayNames")]
        private List<string>? _displayNames;
        public List<string> DisplayNames
        {
            get => _displayNames ?? (_displayNames = new List<string>());
            set => _displayNames = value;
        }

        [Input("names")]
        private List<string>? _names;

        /// <summary>
        /// The Display Names of the Azure AD Groups.
        /// </summary>
        [Obsolete(@"This property has been renamed to `display_names` and will be removed in v2.0 of this provider.")]
        public List<string> Names
        {
            get => _names ?? (_names = new List<string>());
            set => _names = value;
        }

        [Input("objectIds")]
        private List<string>? _objectIds;

        /// <summary>
        /// The Object IDs of the Azure AD Groups.
        /// </summary>
        public List<string> ObjectIds
        {
            get => _objectIds ?? (_objectIds = new List<string>());
            set => _objectIds = value;
        }

        public GetGroupsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGroupsResult
    {
        public readonly ImmutableArray<string> DisplayNames;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Display Names of the Azure AD Groups.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        /// <summary>
        /// The Object IDs of the Azure AD Groups.
        /// </summary>
        public readonly ImmutableArray<string> ObjectIds;

        [OutputConstructor]
        private GetGroupsResult(
            ImmutableArray<string> displayNames,

            string id,

            ImmutableArray<string> names,

            ImmutableArray<string> objectIds)
        {
            DisplayNames = displayNames;
            Id = id;
            Names = names;
            ObjectIds = objectIds;
        }
    }
}

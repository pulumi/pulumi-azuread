// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    public static class GetServicePrincipals
    {
        /// <summary>
        /// Gets basic information for multiple Azure Active Directory service principals.
        /// 
        /// ## API Permissions
        /// 
        /// The following API permissions are required in order to use this data source.
        /// 
        /// When authenticated with a service principal, this data source requires one of the following application roles: `Application.Read.All` or `Directory.Read.All`
        /// 
        /// When authenticated with a user principal, this data source does not require any additional roles.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// *Look up by application display names*
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureAD.GetServicePrincipals.Invoke(new()
        ///     {
        ///         DisplayNames = new[]
        ///         {
        ///             "example-app",
        ///             "another-app",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// *Look up by application IDs (client IDs*
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureAD.GetServicePrincipals.Invoke(new()
        ///     {
        ///         ApplicationIds = new[]
        ///         {
        ///             "11111111-0000-0000-0000-000000000000",
        ///             "22222222-0000-0000-0000-000000000000",
        ///             "33333333-0000-0000-0000-000000000000",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// *Look up by service principal object IDs*
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureAD.GetServicePrincipals.Invoke(new()
        ///     {
        ///         ObjectIds = new[]
        ///         {
        ///             "00000000-0000-0000-0000-000000000000",
        ///             "00000000-0000-0000-0000-111111111111",
        ///             "00000000-0000-0000-0000-222222222222",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServicePrincipalsResult> InvokeAsync(GetServicePrincipalsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServicePrincipalsResult>("azuread:index/getServicePrincipals:getServicePrincipals", args ?? new GetServicePrincipalsArgs(), options.WithDefaults());

        /// <summary>
        /// Gets basic information for multiple Azure Active Directory service principals.
        /// 
        /// ## API Permissions
        /// 
        /// The following API permissions are required in order to use this data source.
        /// 
        /// When authenticated with a service principal, this data source requires one of the following application roles: `Application.Read.All` or `Directory.Read.All`
        /// 
        /// When authenticated with a user principal, this data source does not require any additional roles.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// *Look up by application display names*
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureAD.GetServicePrincipals.Invoke(new()
        ///     {
        ///         DisplayNames = new[]
        ///         {
        ///             "example-app",
        ///             "another-app",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// *Look up by application IDs (client IDs*
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureAD.GetServicePrincipals.Invoke(new()
        ///     {
        ///         ApplicationIds = new[]
        ///         {
        ///             "11111111-0000-0000-0000-000000000000",
        ///             "22222222-0000-0000-0000-000000000000",
        ///             "33333333-0000-0000-0000-000000000000",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// *Look up by service principal object IDs*
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureAD.GetServicePrincipals.Invoke(new()
        ///     {
        ///         ObjectIds = new[]
        ///         {
        ///             "00000000-0000-0000-0000-000000000000",
        ///             "00000000-0000-0000-0000-111111111111",
        ///             "00000000-0000-0000-0000-222222222222",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetServicePrincipalsResult> Invoke(GetServicePrincipalsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServicePrincipalsResult>("azuread:index/getServicePrincipals:getServicePrincipals", args ?? new GetServicePrincipalsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServicePrincipalsArgs : global::Pulumi.InvokeArgs
    {
        [Input("applicationIds")]
        private List<string>? _applicationIds;

        /// <summary>
        /// A list of application IDs (client IDs) of the applications associated with the service principals.
        /// </summary>
        public List<string> ApplicationIds
        {
            get => _applicationIds ?? (_applicationIds = new List<string>());
            set => _applicationIds = value;
        }

        [Input("displayNames")]
        private List<string>? _displayNames;

        /// <summary>
        /// A list of display names of the applications associated with the service principals.
        /// </summary>
        public List<string> DisplayNames
        {
            get => _displayNames ?? (_displayNames = new List<string>());
            set => _displayNames = value;
        }

        /// <summary>
        /// Ignore missing service principals and return all service principals that are found. The data source will still fail if no service principals are found. Defaults to false.
        /// </summary>
        [Input("ignoreMissing")]
        public bool? IgnoreMissing { get; set; }

        [Input("objectIds")]
        private List<string>? _objectIds;

        /// <summary>
        /// The object IDs of the service principals.
        /// </summary>
        public List<string> ObjectIds
        {
            get => _objectIds ?? (_objectIds = new List<string>());
            set => _objectIds = value;
        }

        /// <summary>
        /// When `true`, the data source will return all service principals. Cannot be used with `ignore_missing`. Defaults to false.
        /// </summary>
        [Input("returnAll")]
        public bool? ReturnAll { get; set; }

        public GetServicePrincipalsArgs()
        {
        }
        public static new GetServicePrincipalsArgs Empty => new GetServicePrincipalsArgs();
    }

    public sealed class GetServicePrincipalsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("applicationIds")]
        private InputList<string>? _applicationIds;

        /// <summary>
        /// A list of application IDs (client IDs) of the applications associated with the service principals.
        /// </summary>
        public InputList<string> ApplicationIds
        {
            get => _applicationIds ?? (_applicationIds = new InputList<string>());
            set => _applicationIds = value;
        }

        [Input("displayNames")]
        private InputList<string>? _displayNames;

        /// <summary>
        /// A list of display names of the applications associated with the service principals.
        /// </summary>
        public InputList<string> DisplayNames
        {
            get => _displayNames ?? (_displayNames = new InputList<string>());
            set => _displayNames = value;
        }

        /// <summary>
        /// Ignore missing service principals and return all service principals that are found. The data source will still fail if no service principals are found. Defaults to false.
        /// </summary>
        [Input("ignoreMissing")]
        public Input<bool>? IgnoreMissing { get; set; }

        [Input("objectIds")]
        private InputList<string>? _objectIds;

        /// <summary>
        /// The object IDs of the service principals.
        /// </summary>
        public InputList<string> ObjectIds
        {
            get => _objectIds ?? (_objectIds = new InputList<string>());
            set => _objectIds = value;
        }

        /// <summary>
        /// When `true`, the data source will return all service principals. Cannot be used with `ignore_missing`. Defaults to false.
        /// </summary>
        [Input("returnAll")]
        public Input<bool>? ReturnAll { get; set; }

        public GetServicePrincipalsInvokeArgs()
        {
        }
        public static new GetServicePrincipalsInvokeArgs Empty => new GetServicePrincipalsInvokeArgs();
    }


    [OutputType]
    public sealed class GetServicePrincipalsResult
    {
        /// <summary>
        /// A list of application IDs (client IDs) of the applications associated with the service principals.
        /// </summary>
        public readonly ImmutableArray<string> ApplicationIds;
        /// <summary>
        /// A list of display names of the applications associated with the service principals.
        /// </summary>
        public readonly ImmutableArray<string> DisplayNames;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? IgnoreMissing;
        /// <summary>
        /// The object IDs of the service principals.
        /// </summary>
        public readonly ImmutableArray<string> ObjectIds;
        public readonly bool? ReturnAll;
        /// <summary>
        /// A list of service principals. Each `service_principal` object provides the attributes documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServicePrincipalsServicePrincipalResult> ServicePrincipals;

        [OutputConstructor]
        private GetServicePrincipalsResult(
            ImmutableArray<string> applicationIds,

            ImmutableArray<string> displayNames,

            string id,

            bool? ignoreMissing,

            ImmutableArray<string> objectIds,

            bool? returnAll,

            ImmutableArray<Outputs.GetServicePrincipalsServicePrincipalResult> servicePrincipals)
        {
            ApplicationIds = applicationIds;
            DisplayNames = displayNames;
            Id = id;
            IgnoreMissing = ignoreMissing;
            ObjectIds = objectIds;
            ReturnAll = returnAll;
            ServicePrincipals = servicePrincipals;
        }
    }
}

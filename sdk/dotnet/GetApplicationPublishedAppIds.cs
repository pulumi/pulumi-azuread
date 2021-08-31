// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    public static class GetApplicationPublishedAppIds
    {
        /// <summary>
        /// Use this data source to discover application IDs for APIs published by Microsoft.
        /// 
        /// This data source uses an [unofficial source of application IDs](https://github.com/manicminer/hamilton/blob/main/environments/published.go), as there is currently no available official indexed source for applications or APIs published by Microsoft.
        /// 
        /// The app IDs returned by this data source are sourced from the Azure Global (Public) Cloud, however some of them are known to work in government and national clouds.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// *Listing well-known application IDs*
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var wellKnown = Output.Create(AzureAD.GetApplicationPublishedAppIds.InvokeAsync());
        ///         this.PublishedAppIds = wellKnown.Apply(wellKnown =&gt; wellKnown.Result);
        ///     }
        /// 
        ///     [Output("publishedAppIds")]
        ///     public Output&lt;string&gt; PublishedAppIds { get; set; }
        /// }
        /// ```
        /// 
        /// *Granting access to an application*
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var wellKnown = Output.Create(AzureAD.GetApplicationPublishedAppIds.InvokeAsync());
        ///         var msgraph = new AzureAD.ServicePrincipal("msgraph", new AzureAD.ServicePrincipalArgs
        ///         {
        ///             ApplicationId = wellKnown.Apply(wellKnown =&gt; wellKnown.Result?.MicrosoftGraph),
        ///             UseExisting = true,
        ///         });
        ///         var example = new AzureAD.Application("example", new AzureAD.ApplicationArgs
        ///         {
        ///             DisplayName = "example",
        ///             RequiredResourceAccesses = 
        ///             {
        ///                 new AzureAD.Inputs.ApplicationRequiredResourceAccessArgs
        ///                 {
        ///                     ResourceAppId = wellKnown.Apply(wellKnown =&gt; wellKnown.Result?.MicrosoftGraph),
        ///                     ResourceAccesses = 
        ///                     {
        ///                         new AzureAD.Inputs.ApplicationRequiredResourceAccessResourceAccessArgs
        ///                         {
        ///                             Id = msgraph.AppRoleIds.Apply(appRoleIds =&gt; appRoleIds.User_Read_All),
        ///                             Type = "Role",
        ///                         },
        ///                         new AzureAD.Inputs.ApplicationRequiredResourceAccessResourceAccessArgs
        ///                         {
        ///                             Id = msgraph.Oauth2PermissionScopeIds.Apply(oauth2PermissionScopeIds =&gt; oauth2PermissionScopeIds.User_ReadWrite),
        ///                             Type = "Scope",
        ///                         },
        ///                     },
        ///                 },
        ///             },
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetApplicationPublishedAppIdsResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApplicationPublishedAppIdsResult>("azuread:index/getApplicationPublishedAppIds:getApplicationPublishedAppIds", InvokeArgs.Empty, options.WithVersion());
    }


    [OutputType]
    public sealed class GetApplicationPublishedAppIdsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A map of application names to application IDs.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Result;

        [OutputConstructor]
        private GetApplicationPublishedAppIdsResult(
            string id,

            ImmutableDictionary<string, string> result)
        {
            Id = id;
            Result = result;
        }
    }
}

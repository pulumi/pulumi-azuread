// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    public static class GetApplicationTemplate
    {
        /// <summary>
        /// Use this data source to access information about an Application Template from the [Azure AD App Gallery](https://azuremarketplace.microsoft.com/en-US/marketplace/apps/category/azure-active-directory-apps).
        /// 
        /// ## API Permissions
        /// 
        /// This data source does not require any additional roles.
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
        ///     var example = AzureAD.GetApplicationTemplate.Invoke(new()
        ///     {
        ///         DisplayName = "Marketo",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["applicationTemplateId"] = example.Apply(getApplicationTemplateResult =&gt; getApplicationTemplateResult.TemplateId),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetApplicationTemplateResult> InvokeAsync(GetApplicationTemplateArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApplicationTemplateResult>("azuread:index/getApplicationTemplate:getApplicationTemplate", args ?? new GetApplicationTemplateArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an Application Template from the [Azure AD App Gallery](https://azuremarketplace.microsoft.com/en-US/marketplace/apps/category/azure-active-directory-apps).
        /// 
        /// ## API Permissions
        /// 
        /// This data source does not require any additional roles.
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
        ///     var example = AzureAD.GetApplicationTemplate.Invoke(new()
        ///     {
        ///         DisplayName = "Marketo",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["applicationTemplateId"] = example.Apply(getApplicationTemplateResult =&gt; getApplicationTemplateResult.TemplateId),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetApplicationTemplateResult> Invoke(GetApplicationTemplateInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationTemplateResult>("azuread:index/getApplicationTemplate:getApplicationTemplate", args ?? new GetApplicationTemplateInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to access information about an Application Template from the [Azure AD App Gallery](https://azuremarketplace.microsoft.com/en-US/marketplace/apps/category/azure-active-directory-apps).
        /// 
        /// ## API Permissions
        /// 
        /// This data source does not require any additional roles.
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
        ///     var example = AzureAD.GetApplicationTemplate.Invoke(new()
        ///     {
        ///         DisplayName = "Marketo",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["applicationTemplateId"] = example.Apply(getApplicationTemplateResult =&gt; getApplicationTemplateResult.TemplateId),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetApplicationTemplateResult> Invoke(GetApplicationTemplateInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationTemplateResult>("azuread:index/getApplicationTemplate:getApplicationTemplate", args ?? new GetApplicationTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationTemplateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the display name of the templated application.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        /// <summary>
        /// Specifies the ID of the templated application.
        /// 
        /// &gt; One of `template_id` or `display_name` must be specified.
        /// </summary>
        [Input("templateId")]
        public string? TemplateId { get; set; }

        public GetApplicationTemplateArgs()
        {
        }
        public static new GetApplicationTemplateArgs Empty => new GetApplicationTemplateArgs();
    }

    public sealed class GetApplicationTemplateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the display name of the templated application.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Specifies the ID of the templated application.
        /// 
        /// &gt; One of `template_id` or `display_name` must be specified.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        public GetApplicationTemplateInvokeArgs()
        {
        }
        public static new GetApplicationTemplateInvokeArgs Empty => new GetApplicationTemplateInvokeArgs();
    }


    [OutputType]
    public sealed class GetApplicationTemplateResult
    {
        /// <summary>
        /// List of categories for this templated application.
        /// </summary>
        public readonly ImmutableArray<string> Categories;
        /// <summary>
        /// The display name for the templated application.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Home page URL of the templated application.
        /// </summary>
        public readonly string HomepageUrl;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// URL to retrieve the logo for this templated application.
        /// </summary>
        public readonly string LogoUrl;
        /// <summary>
        /// Name of the publisher for this templated application.
        /// </summary>
        public readonly string Publisher;
        /// <summary>
        /// List of provisioning modes supported by this templated application.
        /// </summary>
        public readonly ImmutableArray<string> SupportedProvisioningTypes;
        /// <summary>
        /// List of single sign on modes supported by this templated application.
        /// </summary>
        public readonly ImmutableArray<string> SupportedSingleSignOnModes;
        /// <summary>
        /// The ID of the templated application.
        /// </summary>
        public readonly string TemplateId;

        [OutputConstructor]
        private GetApplicationTemplateResult(
            ImmutableArray<string> categories,

            string displayName,

            string homepageUrl,

            string id,

            string logoUrl,

            string publisher,

            ImmutableArray<string> supportedProvisioningTypes,

            ImmutableArray<string> supportedSingleSignOnModes,

            string templateId)
        {
            Categories = categories;
            DisplayName = displayName;
            HomepageUrl = homepageUrl;
            Id = id;
            LogoUrl = logoUrl;
            Publisher = publisher;
            SupportedProvisioningTypes = supportedProvisioningTypes;
            SupportedSingleSignOnModes = supportedSingleSignOnModes;
            TemplateId = templateId;
        }
    }
}

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
    /// Creates an application registration and associated service principal from a gallery template.
    /// 
    /// &gt; The azuread.Application resource can also be used to instantiate a gallery application, however unlike the `azuread.Application` resource, this resource does not attempt to manage any properties of the resulting application.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`
    /// 
    /// When authenticated with a user principal, this resource may require one of the following directory roles: `Application Administrator` or `Global Administrator`
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
    ///     var exampleApplicationFromTemplate = new AzureAD.ApplicationFromTemplate("example", new()
    ///     {
    ///         DisplayName = "Example Application",
    ///         TemplateId = example.Apply(getApplicationTemplateResult =&gt; getApplicationTemplateResult.TemplateId),
    ///     });
    /// 
    ///     var exampleGetApplication = AzureAD.GetApplication.Invoke(new()
    ///     {
    ///         ObjectId = exampleApplicationFromTemplate.ApplicationObjectId,
    ///     });
    /// 
    ///     var exampleGetServicePrincipal = AzureAD.GetServicePrincipal.Invoke(new()
    ///     {
    ///         ObjectId = exampleApplicationFromTemplate.ServicePrincipalObjectId,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Templated Applications can be imported using the template ID, the object ID of the application, and the object ID of the service principal, in the following format.
    /// 
    /// ```sh
    /// $ pulumi import azuread:index/applicationFromTemplate:ApplicationFromTemplate example /applicationTemplates/00000000-0000-0000-0000-000000000000/instantiate/11111111-1111-1111-1111-111111111111/22222222-2222-2222-2222-222222222222
    /// ```
    /// </summary>
    [AzureADResourceType("azuread:index/applicationFromTemplate:ApplicationFromTemplate")]
    public partial class ApplicationFromTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The resource ID for the application.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// The object ID for the application.
        /// </summary>
        [Output("applicationObjectId")]
        public Output<string> ApplicationObjectId { get; private set; } = null!;

        /// <summary>
        /// The display name for the application.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The resource ID for the service principal.
        /// </summary>
        [Output("servicePrincipalId")]
        public Output<string> ServicePrincipalId { get; private set; } = null!;

        /// <summary>
        /// The object ID for the service principal.
        /// </summary>
        [Output("servicePrincipalObjectId")]
        public Output<string> ServicePrincipalObjectId { get; private set; } = null!;

        /// <summary>
        /// Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
        /// </summary>
        [Output("templateId")]
        public Output<string> TemplateId { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationFromTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationFromTemplate(string name, ApplicationFromTemplateArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/applicationFromTemplate:ApplicationFromTemplate", name, args ?? new ApplicationFromTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationFromTemplate(string name, Input<string> id, ApplicationFromTemplateState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/applicationFromTemplate:ApplicationFromTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationFromTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationFromTemplate Get(string name, Input<string> id, ApplicationFromTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationFromTemplate(name, id, state, options);
        }
    }

    public sealed class ApplicationFromTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The display name for the application.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
        /// </summary>
        [Input("templateId", required: true)]
        public Input<string> TemplateId { get; set; } = null!;

        public ApplicationFromTemplateArgs()
        {
        }
        public static new ApplicationFromTemplateArgs Empty => new ApplicationFromTemplateArgs();
    }

    public sealed class ApplicationFromTemplateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID for the application.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// The object ID for the application.
        /// </summary>
        [Input("applicationObjectId")]
        public Input<string>? ApplicationObjectId { get; set; }

        /// <summary>
        /// The display name for the application.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The resource ID for the service principal.
        /// </summary>
        [Input("servicePrincipalId")]
        public Input<string>? ServicePrincipalId { get; set; }

        /// <summary>
        /// The object ID for the service principal.
        /// </summary>
        [Input("servicePrincipalObjectId")]
        public Input<string>? ServicePrincipalObjectId { get; set; }

        /// <summary>
        /// Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        public ApplicationFromTemplateState()
        {
        }
        public static new ApplicationFromTemplateState Empty => new ApplicationFromTemplateState();
    }
}

// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ServicePrincipalFeatureArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether this service principal represents a custom SAML application
        /// </summary>
        [Input("customSingleSignOnApp")]
        public Input<bool>? CustomSingleSignOnApp { get; set; }

        /// <summary>
        /// Whether this service principal represents an Enterprise Application
        /// </summary>
        [Input("enterpriseApplication")]
        public Input<bool>? EnterpriseApplication { get; set; }

        /// <summary>
        /// Whether this service principal represents a gallery application
        /// </summary>
        [Input("galleryApplication")]
        public Input<bool>? GalleryApplication { get; set; }

        /// <summary>
        /// Whether this app is visible to users in My Apps and Office 365 Launcher
        /// </summary>
        [Input("visibleToUsers")]
        public Input<bool>? VisibleToUsers { get; set; }

        public ServicePrincipalFeatureArgs()
        {
        }
        public static new ServicePrincipalFeatureArgs Empty => new ServicePrincipalFeatureArgs();
    }
}

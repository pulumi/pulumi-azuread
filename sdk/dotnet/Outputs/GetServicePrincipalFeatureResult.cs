// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Outputs
{

    [OutputType]
    public sealed class GetServicePrincipalFeatureResult
    {
        public readonly bool CustomSingleSignOnApp;
        public readonly bool EnterpriseApplication;
        public readonly bool GalleryApplication;
        public readonly bool VisibleToUsers;

        [OutputConstructor]
        private GetServicePrincipalFeatureResult(
            bool customSingleSignOnApp,

            bool enterpriseApplication,

            bool galleryApplication,

            bool visibleToUsers)
        {
            CustomSingleSignOnApp = customSingleSignOnApp;
            EnterpriseApplication = enterpriseApplication;
            GalleryApplication = galleryApplication;
            VisibleToUsers = visibleToUsers;
        }
    }
}

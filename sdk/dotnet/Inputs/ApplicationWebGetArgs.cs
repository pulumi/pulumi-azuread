// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ApplicationWebGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("homepageUrl")]
        public Input<string>? HomepageUrl { get; set; }

        [Input("implicitGrant")]
        public Input<Inputs.ApplicationWebImplicitGrantGetArgs>? ImplicitGrant { get; set; }

        [Input("logoutUrl")]
        public Input<string>? LogoutUrl { get; set; }

        [Input("redirectUris")]
        private InputList<string>? _redirectUris;
        public InputList<string> RedirectUris
        {
            get => _redirectUris ?? (_redirectUris = new InputList<string>());
            set => _redirectUris = value;
        }

        public ApplicationWebGetArgs()
        {
        }
        public static new ApplicationWebGetArgs Empty => new ApplicationWebGetArgs();
    }
}

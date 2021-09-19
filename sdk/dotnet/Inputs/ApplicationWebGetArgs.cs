// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ApplicationWebGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Home page or landing page of the application.
        /// </summary>
        [Input("homepageUrl")]
        public Input<string>? HomepageUrl { get; set; }

        /// <summary>
        /// An `implicit_grant` block as documented above.
        /// </summary>
        [Input("implicitGrant")]
        public Input<Inputs.ApplicationWebImplicitGrantGetArgs>? ImplicitGrant { get; set; }

        /// <summary>
        /// The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
        /// </summary>
        [Input("logoutUrl")]
        public Input<string>? LogoutUrl { get; set; }

        [Input("redirectUris")]
        private InputList<string>? _redirectUris;

        /// <summary>
        /// A set of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. Must be a valid `http` URL or a URN.
        /// </summary>
        public InputList<string> RedirectUris
        {
            get => _redirectUris ?? (_redirectUris = new InputList<string>());
            set => _redirectUris = value;
        }

        public ApplicationWebGetArgs()
        {
        }
    }
}

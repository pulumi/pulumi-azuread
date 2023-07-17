// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ServicePrincipalSamlSingleSignOnArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The relative URI the service provider would redirect to after completion of the single sign-on flow.
        /// </summary>
        [Input("relayState")]
        public Input<string>? RelayState { get; set; }

        public ServicePrincipalSamlSingleSignOnArgs()
        {
        }
        public static new ServicePrincipalSamlSingleSignOnArgs Empty => new ServicePrincipalSamlSingleSignOnArgs();
    }
}

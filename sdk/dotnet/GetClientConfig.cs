// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    public static class GetClientConfig
    {
        public static Task<GetClientConfigResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClientConfigResult>("azuread:index/getClientConfig:getClientConfig", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetClientConfigResult
    {
        public readonly string ClientId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ObjectId;
        public readonly string TenantId;

        [OutputConstructor]
        private GetClientConfigResult(
            string clientId,

            string id,

            string objectId,

            string tenantId)
        {
            ClientId = clientId;
            Id = id;
            ObjectId = objectId;
            TenantId = tenantId;
        }
    }
}

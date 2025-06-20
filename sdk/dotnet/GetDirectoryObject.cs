// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    public static class GetDirectoryObject
    {
        /// <summary>
        /// Retrieves the OData type for a generic directory object having the provided object ID.
        /// 
        /// ## API Permissions
        /// 
        /// The following API permissions are required in order to use this data source.
        /// 
        /// When authenticated with a service principal, this data source requires either `User.Read.All`, `Group.Read.All` or `Directory.Read.All`, depending on the type of object being queried.
        /// 
        /// When authenticated with a user principal, this data source does not require any additional roles.
        /// 
        /// ## Example Usage
        /// 
        /// *Look up and output type of object by ID*
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureAD.GetDirectoryObject.Invoke(new()
        ///     {
        ///         ObjectId = "00000000-0000-0000-0000-000000000000",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["objectType"] = example.Apply(getDirectoryObjectResult =&gt; getDirectoryObjectResult.Type),
        ///     };
        /// });
        /// ```
        /// 
        /// ## Attributes Reference 
        /// 
        /// The following attributes are exported:
        /// 
        /// * `object_id` - The object ID of the directory object.
        /// * `type` - The shortened OData type of the directory object. Possible values include: `Group`, `User` or `ServicePrincipal`.
        /// </summary>
        public static Task<GetDirectoryObjectResult> InvokeAsync(GetDirectoryObjectArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDirectoryObjectResult>("azuread:index/getDirectoryObject:getDirectoryObject", args ?? new GetDirectoryObjectArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the OData type for a generic directory object having the provided object ID.
        /// 
        /// ## API Permissions
        /// 
        /// The following API permissions are required in order to use this data source.
        /// 
        /// When authenticated with a service principal, this data source requires either `User.Read.All`, `Group.Read.All` or `Directory.Read.All`, depending on the type of object being queried.
        /// 
        /// When authenticated with a user principal, this data source does not require any additional roles.
        /// 
        /// ## Example Usage
        /// 
        /// *Look up and output type of object by ID*
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureAD.GetDirectoryObject.Invoke(new()
        ///     {
        ///         ObjectId = "00000000-0000-0000-0000-000000000000",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["objectType"] = example.Apply(getDirectoryObjectResult =&gt; getDirectoryObjectResult.Type),
        ///     };
        /// });
        /// ```
        /// 
        /// ## Attributes Reference 
        /// 
        /// The following attributes are exported:
        /// 
        /// * `object_id` - The object ID of the directory object.
        /// * `type` - The shortened OData type of the directory object. Possible values include: `Group`, `User` or `ServicePrincipal`.
        /// </summary>
        public static Output<GetDirectoryObjectResult> Invoke(GetDirectoryObjectInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDirectoryObjectResult>("azuread:index/getDirectoryObject:getDirectoryObject", args ?? new GetDirectoryObjectInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the OData type for a generic directory object having the provided object ID.
        /// 
        /// ## API Permissions
        /// 
        /// The following API permissions are required in order to use this data source.
        /// 
        /// When authenticated with a service principal, this data source requires either `User.Read.All`, `Group.Read.All` or `Directory.Read.All`, depending on the type of object being queried.
        /// 
        /// When authenticated with a user principal, this data source does not require any additional roles.
        /// 
        /// ## Example Usage
        /// 
        /// *Look up and output type of object by ID*
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = AzureAD.GetDirectoryObject.Invoke(new()
        ///     {
        ///         ObjectId = "00000000-0000-0000-0000-000000000000",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["objectType"] = example.Apply(getDirectoryObjectResult =&gt; getDirectoryObjectResult.Type),
        ///     };
        /// });
        /// ```
        /// 
        /// ## Attributes Reference 
        /// 
        /// The following attributes are exported:
        /// 
        /// * `object_id` - The object ID of the directory object.
        /// * `type` - The shortened OData type of the directory object. Possible values include: `Group`, `User` or `ServicePrincipal`.
        /// </summary>
        public static Output<GetDirectoryObjectResult> Invoke(GetDirectoryObjectInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDirectoryObjectResult>("azuread:index/getDirectoryObject:getDirectoryObject", args ?? new GetDirectoryObjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDirectoryObjectArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the Object ID of the directory object to look up.
        /// </summary>
        [Input("objectId", required: true)]
        public string ObjectId { get; set; } = null!;

        public GetDirectoryObjectArgs()
        {
        }
        public static new GetDirectoryObjectArgs Empty => new GetDirectoryObjectArgs();
    }

    public sealed class GetDirectoryObjectInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the Object ID of the directory object to look up.
        /// </summary>
        [Input("objectId", required: true)]
        public Input<string> ObjectId { get; set; } = null!;

        public GetDirectoryObjectInvokeArgs()
        {
        }
        public static new GetDirectoryObjectInvokeArgs Empty => new GetDirectoryObjectInvokeArgs();
    }


    [OutputType]
    public sealed class GetDirectoryObjectResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ObjectId;
        public readonly string Type;

        [OutputConstructor]
        private GetDirectoryObjectResult(
            string id,

            string objectId,

            string type)
        {
            Id = id;
            ObjectId = objectId;
            Type = type;
        }
    }
}

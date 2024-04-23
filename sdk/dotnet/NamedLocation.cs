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
    /// Manages a Named Location within Azure Active Directory.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ConditionalAccess` and `Policy.Read.All`
    /// 
    /// When authenticated with a user principal, this resource requires one of the following directory roles: `Conditional Access Administrator` or `Global Administrator`
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
    ///     var example_ip = new AzureAD.NamedLocation("example-ip", new()
    ///     {
    ///         DisplayName = "IP Named Location",
    ///         Ip = new AzureAD.Inputs.NamedLocationIpArgs
    ///         {
    ///             IpRanges = new[]
    ///             {
    ///                 "1.1.1.1/32",
    ///                 "2.2.2.2/32",
    ///             },
    ///             Trusted = true,
    ///         },
    ///     });
    /// 
    ///     var example_country = new AzureAD.NamedLocation("example-country", new()
    ///     {
    ///         DisplayName = "Country Named Location",
    ///         Country = new AzureAD.Inputs.NamedLocationCountryArgs
    ///         {
    ///             CountriesAndRegions = new[]
    ///             {
    ///                 "GB",
    ///                 "US",
    ///             },
    ///             IncludeUnknownCountriesAndRegions = false,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Named Locations can be imported using the `id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuread:index/namedLocation:NamedLocation my_location 00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureADResourceType("azuread:index/namedLocation:NamedLocation")]
    public partial class NamedLocation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A `country` block as documented below, which configures a country-based named location.
        /// </summary>
        [Output("country")]
        public Output<Outputs.NamedLocationCountry?> Country { get; private set; } = null!;

        /// <summary>
        /// The friendly name for this named location.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// An `ip` block as documented below, which configures an IP-based named location.
        /// 
        /// &gt; Exactly one of `ip` or `country` must be specified. Changing between these forces a new resource to be created.
        /// </summary>
        [Output("ip")]
        public Output<Outputs.NamedLocationIp?> Ip { get; private set; } = null!;


        /// <summary>
        /// Create a NamedLocation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NamedLocation(string name, NamedLocationArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/namedLocation:NamedLocation", name, args ?? new NamedLocationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NamedLocation(string name, Input<string> id, NamedLocationState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/namedLocation:NamedLocation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NamedLocation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NamedLocation Get(string name, Input<string> id, NamedLocationState? state = null, CustomResourceOptions? options = null)
        {
            return new NamedLocation(name, id, state, options);
        }
    }

    public sealed class NamedLocationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `country` block as documented below, which configures a country-based named location.
        /// </summary>
        [Input("country")]
        public Input<Inputs.NamedLocationCountryArgs>? Country { get; set; }

        /// <summary>
        /// The friendly name for this named location.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// An `ip` block as documented below, which configures an IP-based named location.
        /// 
        /// &gt; Exactly one of `ip` or `country` must be specified. Changing between these forces a new resource to be created.
        /// </summary>
        [Input("ip")]
        public Input<Inputs.NamedLocationIpArgs>? Ip { get; set; }

        public NamedLocationArgs()
        {
        }
        public static new NamedLocationArgs Empty => new NamedLocationArgs();
    }

    public sealed class NamedLocationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `country` block as documented below, which configures a country-based named location.
        /// </summary>
        [Input("country")]
        public Input<Inputs.NamedLocationCountryGetArgs>? Country { get; set; }

        /// <summary>
        /// The friendly name for this named location.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// An `ip` block as documented below, which configures an IP-based named location.
        /// 
        /// &gt; Exactly one of `ip` or `country` must be specified. Changing between these forces a new resource to be created.
        /// </summary>
        [Input("ip")]
        public Input<Inputs.NamedLocationIpGetArgs>? Ip { get; set; }

        public NamedLocationState()
        {
        }
        public static new NamedLocationState Empty => new NamedLocationState();
    }
}

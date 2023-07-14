// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class NamedLocationCountryArgs : global::Pulumi.ResourceArgs
    {
        [Input("countriesAndRegions", required: true)]
        private InputList<string>? _countriesAndRegions;
        public InputList<string> CountriesAndRegions
        {
            get => _countriesAndRegions ?? (_countriesAndRegions = new InputList<string>());
            set => _countriesAndRegions = value;
        }

        [Input("includeUnknownCountriesAndRegions")]
        public Input<bool>? IncludeUnknownCountriesAndRegions { get; set; }

        public NamedLocationCountryArgs()
        {
        }
        public static new NamedLocationCountryArgs Empty => new NamedLocationCountryArgs();
    }
}

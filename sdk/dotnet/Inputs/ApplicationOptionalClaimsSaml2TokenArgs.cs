// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ApplicationOptionalClaimsSaml2TokenArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalProperties")]
        private InputList<string>? _additionalProperties;

        /// <summary>
        /// List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim. Possible values are: `cloud_displayname`, `dns_domain_and_sam_account_name`, `emit_as_roles`, `include_externally_authenticated_upn_without_hash`, `include_externally_authenticated_upn`, `max_size_limit`, `netbios_domain_and_sam_account_name`, `on_premise_security_identifier`, `sam_account_name`, and `use_guid`.
        /// </summary>
        public InputList<string> AdditionalProperties
        {
            get => _additionalProperties ?? (_additionalProperties = new InputList<string>());
            set => _additionalProperties = value;
        }

        /// <summary>
        /// Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
        /// </summary>
        [Input("essential")]
        public Input<bool>? Essential { get; set; }

        /// <summary>
        /// The name of the optional claim.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        public ApplicationOptionalClaimsSaml2TokenArgs()
        {
        }
        public static new ApplicationOptionalClaimsSaml2TokenArgs Empty => new ApplicationOptionalClaimsSaml2TokenArgs();
    }
}

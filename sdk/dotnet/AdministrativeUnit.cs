// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    [AzureADResourceType("azuread:index/administrativeUnit:AdministrativeUnit")]
    public partial class AdministrativeUnit : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description for the administrative unit
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The display name for the administrative unit
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Whether the administrative unit and its members are hidden or publicly viewable in the directory
        /// </summary>
        [Output("hiddenMembershipEnabled")]
        public Output<bool?> HiddenMembershipEnabled { get; private set; } = null!;

        /// <summary>
        /// A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or
        /// Groups
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The object ID of the administrative unit
        /// </summary>
        [Output("objectId")]
        public Output<string> ObjectId { get; private set; } = null!;

        /// <summary>
        /// If `true`, will return an error if an existing administrative unit is found with the same name
        /// </summary>
        [Output("preventDuplicateNames")]
        public Output<bool?> PreventDuplicateNames { get; private set; } = null!;


        /// <summary>
        /// Create a AdministrativeUnit resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AdministrativeUnit(string name, AdministrativeUnitArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/administrativeUnit:AdministrativeUnit", name, args ?? new AdministrativeUnitArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AdministrativeUnit(string name, Input<string> id, AdministrativeUnitState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/administrativeUnit:AdministrativeUnit", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AdministrativeUnit resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AdministrativeUnit Get(string name, Input<string> id, AdministrativeUnitState? state = null, CustomResourceOptions? options = null)
        {
            return new AdministrativeUnit(name, id, state, options);
        }
    }

    public sealed class AdministrativeUnitArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description for the administrative unit
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name for the administrative unit
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Whether the administrative unit and its members are hidden or publicly viewable in the directory
        /// </summary>
        [Input("hiddenMembershipEnabled")]
        public Input<bool>? HiddenMembershipEnabled { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or
        /// Groups
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// If `true`, will return an error if an existing administrative unit is found with the same name
        /// </summary>
        [Input("preventDuplicateNames")]
        public Input<bool>? PreventDuplicateNames { get; set; }

        public AdministrativeUnitArgs()
        {
        }
        public static new AdministrativeUnitArgs Empty => new AdministrativeUnitArgs();
    }

    public sealed class AdministrativeUnitState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description for the administrative unit
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name for the administrative unit
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Whether the administrative unit and its members are hidden or publicly viewable in the directory
        /// </summary>
        [Input("hiddenMembershipEnabled")]
        public Input<bool>? HiddenMembershipEnabled { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or
        /// Groups
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The object ID of the administrative unit
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        /// <summary>
        /// If `true`, will return an error if an existing administrative unit is found with the same name
        /// </summary>
        [Input("preventDuplicateNames")]
        public Input<bool>? PreventDuplicateNames { get; set; }

        public AdministrativeUnitState()
        {
        }
        public static new AdministrativeUnitState Empty => new AdministrativeUnitState();
    }
}

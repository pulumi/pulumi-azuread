// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ConditionalAccessPolicyConditionsUsersGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("excludedGroups")]
        private InputList<string>? _excludedGroups;
        public InputList<string> ExcludedGroups
        {
            get => _excludedGroups ?? (_excludedGroups = new InputList<string>());
            set => _excludedGroups = value;
        }

        [Input("excludedRoles")]
        private InputList<string>? _excludedRoles;
        public InputList<string> ExcludedRoles
        {
            get => _excludedRoles ?? (_excludedRoles = new InputList<string>());
            set => _excludedRoles = value;
        }

        [Input("excludedUsers")]
        private InputList<string>? _excludedUsers;
        public InputList<string> ExcludedUsers
        {
            get => _excludedUsers ?? (_excludedUsers = new InputList<string>());
            set => _excludedUsers = value;
        }

        [Input("includedGroups")]
        private InputList<string>? _includedGroups;
        public InputList<string> IncludedGroups
        {
            get => _includedGroups ?? (_includedGroups = new InputList<string>());
            set => _includedGroups = value;
        }

        [Input("includedRoles")]
        private InputList<string>? _includedRoles;
        public InputList<string> IncludedRoles
        {
            get => _includedRoles ?? (_includedRoles = new InputList<string>());
            set => _includedRoles = value;
        }

        [Input("includedUsers")]
        private InputList<string>? _includedUsers;
        public InputList<string> IncludedUsers
        {
            get => _includedUsers ?? (_includedUsers = new InputList<string>());
            set => _includedUsers = value;
        }

        public ConditionalAccessPolicyConditionsUsersGetArgs()
        {
        }
        public static new ConditionalAccessPolicyConditionsUsersGetArgs Empty => new ConditionalAccessPolicyConditionsUsersGetArgs();
    }
}

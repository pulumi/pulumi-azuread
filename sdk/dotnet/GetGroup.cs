// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    public static class GetGroup
    {
        public static Task<GetGroupResult> InvokeAsync(GetGroupArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("azuread:index/getGroup:getGroup", args ?? new GetGroupArgs(), options.WithDefaults());

        public static Output<GetGroupResult> Invoke(GetGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupResult>("azuread:index/getGroup:getGroup", args ?? new GetGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("displayName")]
        public string? DisplayName { get; set; }

        [Input("mailEnabled")]
        public bool? MailEnabled { get; set; }

        [Input("objectId")]
        public string? ObjectId { get; set; }

        [Input("securityEnabled")]
        public bool? SecurityEnabled { get; set; }

        public GetGroupArgs()
        {
        }
        public static new GetGroupArgs Empty => new GetGroupArgs();
    }

    public sealed class GetGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("mailEnabled")]
        public Input<bool>? MailEnabled { get; set; }

        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        [Input("securityEnabled")]
        public Input<bool>? SecurityEnabled { get; set; }

        public GetGroupInvokeArgs()
        {
        }
        public static new GetGroupInvokeArgs Empty => new GetGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupResult
    {
        public readonly bool AssignableToRole;
        public readonly bool AutoSubscribeNewMembers;
        public readonly ImmutableArray<string> Behaviors;
        public readonly string Description;
        public readonly string DisplayName;
        public readonly ImmutableArray<Outputs.GetGroupDynamicMembershipResult> DynamicMemberships;
        public readonly bool ExternalSendersAllowed;
        public readonly bool HideFromAddressLists;
        public readonly bool HideFromOutlookClients;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Mail;
        public readonly bool MailEnabled;
        public readonly string MailNickname;
        public readonly ImmutableArray<string> Members;
        public readonly string ObjectId;
        public readonly string OnpremisesDomainName;
        public readonly string OnpremisesGroupType;
        public readonly string OnpremisesNetbiosName;
        public readonly string OnpremisesSamAccountName;
        public readonly string OnpremisesSecurityIdentifier;
        public readonly bool OnpremisesSyncEnabled;
        public readonly ImmutableArray<string> Owners;
        public readonly string PreferredLanguage;
        public readonly ImmutableArray<string> ProvisioningOptions;
        public readonly ImmutableArray<string> ProxyAddresses;
        public readonly bool SecurityEnabled;
        public readonly string Theme;
        public readonly ImmutableArray<string> Types;
        public readonly string Visibility;
        public readonly bool WritebackEnabled;

        [OutputConstructor]
        private GetGroupResult(
            bool assignableToRole,

            bool autoSubscribeNewMembers,

            ImmutableArray<string> behaviors,

            string description,

            string displayName,

            ImmutableArray<Outputs.GetGroupDynamicMembershipResult> dynamicMemberships,

            bool externalSendersAllowed,

            bool hideFromAddressLists,

            bool hideFromOutlookClients,

            string id,

            string mail,

            bool mailEnabled,

            string mailNickname,

            ImmutableArray<string> members,

            string objectId,

            string onpremisesDomainName,

            string onpremisesGroupType,

            string onpremisesNetbiosName,

            string onpremisesSamAccountName,

            string onpremisesSecurityIdentifier,

            bool onpremisesSyncEnabled,

            ImmutableArray<string> owners,

            string preferredLanguage,

            ImmutableArray<string> provisioningOptions,

            ImmutableArray<string> proxyAddresses,

            bool securityEnabled,

            string theme,

            ImmutableArray<string> types,

            string visibility,

            bool writebackEnabled)
        {
            AssignableToRole = assignableToRole;
            AutoSubscribeNewMembers = autoSubscribeNewMembers;
            Behaviors = behaviors;
            Description = description;
            DisplayName = displayName;
            DynamicMemberships = dynamicMemberships;
            ExternalSendersAllowed = externalSendersAllowed;
            HideFromAddressLists = hideFromAddressLists;
            HideFromOutlookClients = hideFromOutlookClients;
            Id = id;
            Mail = mail;
            MailEnabled = mailEnabled;
            MailNickname = mailNickname;
            Members = members;
            ObjectId = objectId;
            OnpremisesDomainName = onpremisesDomainName;
            OnpremisesGroupType = onpremisesGroupType;
            OnpremisesNetbiosName = onpremisesNetbiosName;
            OnpremisesSamAccountName = onpremisesSamAccountName;
            OnpremisesSecurityIdentifier = onpremisesSecurityIdentifier;
            OnpremisesSyncEnabled = onpremisesSyncEnabled;
            Owners = owners;
            PreferredLanguage = preferredLanguage;
            ProvisioningOptions = provisioningOptions;
            ProxyAddresses = proxyAddresses;
            SecurityEnabled = securityEnabled;
            Theme = theme;
            Types = types;
            Visibility = visibility;
            WritebackEnabled = writebackEnabled;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    public static class GetGroups
    {
        /// <summary>
        /// Gets Object IDs or Display Names for multiple Azure Active Directory groups.
        /// 
        /// ## API Permissions
        /// 
        /// The following API permissions are required in order to use this data source.
        /// 
        /// When authenticated with a service principal, this data source requires one of the following application roles: `Group.Read.All` or `Directory.Read.All`
        /// 
        /// When authenticated with a user principal, this data source does not require any additional roles.
        /// </summary>
        public static Task<GetGroupsResult> InvokeAsync(GetGroupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupsResult>("azuread:index/getGroups:getGroups", args ?? new GetGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// Gets Object IDs or Display Names for multiple Azure Active Directory groups.
        /// 
        /// ## API Permissions
        /// 
        /// The following API permissions are required in order to use this data source.
        /// 
        /// When authenticated with a service principal, this data source requires one of the following application roles: `Group.Read.All` or `Directory.Read.All`
        /// 
        /// When authenticated with a user principal, this data source does not require any additional roles.
        /// </summary>
        public static Output<GetGroupsResult> Invoke(GetGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupsResult>("azuread:index/getGroups:getGroups", args ?? new GetGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A common display name prefix to match when returning groups.
        /// </summary>
        [Input("displayNamePrefix")]
        public string? DisplayNamePrefix { get; set; }

        [Input("displayNames")]
        private List<string>? _displayNames;

        /// <summary>
        /// The display names of the groups.
        /// </summary>
        public List<string> DisplayNames
        {
            get => _displayNames ?? (_displayNames = new List<string>());
            set => _displayNames = value;
        }

        /// <summary>
        /// Ignore missing groups and return groups that were found. The data source will still fail if no groups are found. Cannot be specified with `return_all`. Defaults to `false`.
        /// </summary>
        [Input("ignoreMissing")]
        public bool? IgnoreMissing { get; set; }

        /// <summary>
        /// Whether the returned groups should be mail-enabled. By itself this does not exclude security-enabled groups. Setting this to `true` ensures all groups are mail-enabled, and setting to `false` ensures that all groups are _not_ mail-enabled. To ignore this filter, omit the property or set it to null. Cannot be specified together with `object_ids`.
        /// </summary>
        [Input("mailEnabled")]
        public bool? MailEnabled { get; set; }

        [Input("objectIds")]
        private List<string>? _objectIds;

        /// <summary>
        /// The object IDs of the groups.
        /// </summary>
        public List<string> ObjectIds
        {
            get => _objectIds ?? (_objectIds = new List<string>());
            set => _objectIds = value;
        }

        /// <summary>
        /// A flag to denote if all groups should be fetched and returned. Cannot be specified wth `ignore_missing`. Defaults to `false`.
        /// </summary>
        [Input("returnAll")]
        public bool? ReturnAll { get; set; }

        /// <summary>
        /// Whether the returned groups should be security-enabled. By itself this does not exclude mail-enabled groups. Setting this to `true` ensures all groups are security-enabled, and setting to `false` ensures that all groups are _not_ security-enabled. To ignore this filter, omit the property or set it to null. Cannot be specified together with `object_ids`.
        /// 
        /// &gt; One of `display_names`, `display_name_prefix`, `object_ids` or `return_all` should be specified. Either `display_name` or `object_ids` _may_ be specified as an empty list, in which case no results will be returned.
        /// </summary>
        [Input("securityEnabled")]
        public bool? SecurityEnabled { get; set; }

        public GetGroupsArgs()
        {
        }
        public static new GetGroupsArgs Empty => new GetGroupsArgs();
    }

    public sealed class GetGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A common display name prefix to match when returning groups.
        /// </summary>
        [Input("displayNamePrefix")]
        public Input<string>? DisplayNamePrefix { get; set; }

        [Input("displayNames")]
        private InputList<string>? _displayNames;

        /// <summary>
        /// The display names of the groups.
        /// </summary>
        public InputList<string> DisplayNames
        {
            get => _displayNames ?? (_displayNames = new InputList<string>());
            set => _displayNames = value;
        }

        /// <summary>
        /// Ignore missing groups and return groups that were found. The data source will still fail if no groups are found. Cannot be specified with `return_all`. Defaults to `false`.
        /// </summary>
        [Input("ignoreMissing")]
        public Input<bool>? IgnoreMissing { get; set; }

        /// <summary>
        /// Whether the returned groups should be mail-enabled. By itself this does not exclude security-enabled groups. Setting this to `true` ensures all groups are mail-enabled, and setting to `false` ensures that all groups are _not_ mail-enabled. To ignore this filter, omit the property or set it to null. Cannot be specified together with `object_ids`.
        /// </summary>
        [Input("mailEnabled")]
        public Input<bool>? MailEnabled { get; set; }

        [Input("objectIds")]
        private InputList<string>? _objectIds;

        /// <summary>
        /// The object IDs of the groups.
        /// </summary>
        public InputList<string> ObjectIds
        {
            get => _objectIds ?? (_objectIds = new InputList<string>());
            set => _objectIds = value;
        }

        /// <summary>
        /// A flag to denote if all groups should be fetched and returned. Cannot be specified wth `ignore_missing`. Defaults to `false`.
        /// </summary>
        [Input("returnAll")]
        public Input<bool>? ReturnAll { get; set; }

        /// <summary>
        /// Whether the returned groups should be security-enabled. By itself this does not exclude mail-enabled groups. Setting this to `true` ensures all groups are security-enabled, and setting to `false` ensures that all groups are _not_ security-enabled. To ignore this filter, omit the property or set it to null. Cannot be specified together with `object_ids`.
        /// 
        /// &gt; One of `display_names`, `display_name_prefix`, `object_ids` or `return_all` should be specified. Either `display_name` or `object_ids` _may_ be specified as an empty list, in which case no results will be returned.
        /// </summary>
        [Input("securityEnabled")]
        public Input<bool>? SecurityEnabled { get; set; }

        public GetGroupsInvokeArgs()
        {
        }
        public static new GetGroupsInvokeArgs Empty => new GetGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupsResult
    {
        public readonly string DisplayNamePrefix;
        /// <summary>
        /// The display names of the groups.
        /// </summary>
        public readonly ImmutableArray<string> DisplayNames;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? IgnoreMissing;
        public readonly bool MailEnabled;
        /// <summary>
        /// The object IDs of the groups.
        /// </summary>
        public readonly ImmutableArray<string> ObjectIds;
        public readonly bool? ReturnAll;
        public readonly bool SecurityEnabled;

        [OutputConstructor]
        private GetGroupsResult(
            string displayNamePrefix,

            ImmutableArray<string> displayNames,

            string id,

            bool? ignoreMissing,

            bool mailEnabled,

            ImmutableArray<string> objectIds,

            bool? returnAll,

            bool securityEnabled)
        {
            DisplayNamePrefix = displayNamePrefix;
            DisplayNames = displayNames;
            Id = id;
            IgnoreMissing = ignoreMissing;
            MailEnabled = mailEnabled;
            ObjectIds = objectIds;
            ReturnAll = returnAll;
            SecurityEnabled = securityEnabled;
        }
    }
}

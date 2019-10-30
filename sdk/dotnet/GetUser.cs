// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azuread
{
    public static partial class Invokes
    {
        /// <summary>
        /// Gets information about an Azure Active Directory user.
        /// 
        /// &gt; **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azuread/blob/master/website/docs/d/user.html.markdown.
        /// </summary>
        public static Task<GetUserResult> GetUser(GetUserArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("azuread:index/getUser:getUser", args, options.WithVersion());
    }

    public sealed class GetUserArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the Object ID of the Application within Azure Active Directory.
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        /// <summary>
        /// The User Principal Name of the Azure AD User.
        /// </summary>
        [Input("userPrincipalName")]
        public Input<string>? UserPrincipalName { get; set; }

        public GetUserArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetUserResult
    {
        /// <summary>
        /// `True` if the account is enabled; otherwise `False`.
        /// </summary>
        public readonly bool AccountEnabled;
        /// <summary>
        /// The Display Name of the Azure AD User.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The primary email address of the Azure AD User.
        /// </summary>
        public readonly string Mail;
        /// <summary>
        /// The email alias of the Azure AD User.
        /// </summary>
        public readonly string MailNickname;
        public readonly string ObjectId;
        /// <summary>
        /// The User Principal Name of the Azure AD User.
        /// </summary>
        public readonly string UserPrincipalName;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetUserResult(
            bool accountEnabled,
            string displayName,
            string mail,
            string mailNickname,
            string objectId,
            string userPrincipalName,
            string id)
        {
            AccountEnabled = accountEnabled;
            DisplayName = displayName;
            Mail = mail;
            MailNickname = mailNickname;
            ObjectId = objectId;
            UserPrincipalName = userPrincipalName;
            Id = id;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the azuread package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'azuread';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    /**
     * The Azure DevOps Pipeline Service Connection ID.
     */
    public readonly adoPipelineServiceConnectionId!: pulumi.Output<string | undefined>;
    /**
     * Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate
     */
    public readonly clientCertificate!: pulumi.Output<string | undefined>;
    /**
     * The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
     * Certificate
     */
    public readonly clientCertificatePassword!: pulumi.Output<string | undefined>;
    /**
     * The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
     * Principal using a Client Certificate
     */
    public readonly clientCertificatePath!: pulumi.Output<string | undefined>;
    /**
     * The Client ID which should be used for service principal authentication
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * The path to a file containing the Client ID which should be used for service principal authentication
     */
    public readonly clientIdFilePath!: pulumi.Output<string | undefined>;
    /**
     * The application password to use when authenticating as a Service Principal using a Client Secret
     */
    public readonly clientSecret!: pulumi.Output<string | undefined>;
    /**
     * The path to a file containing the application password to use when authenticating as a Service Principal using a Client
     * Secret
     */
    public readonly clientSecretFilePath!: pulumi.Output<string | undefined>;
    /**
     * The cloud environment which should be used. Possible values are: `global` (also `public`), `usgovernmentl4` (also
     * `usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`. Not used and should not be specified
     * when `metadataHost` is specified.
     */
    public readonly environment!: pulumi.Output<string | undefined>;
    /**
     * The Hostname which should be used for the Azure Metadata Service.
     */
    public readonly metadataHost!: pulumi.Output<string | undefined>;
    /**
     * The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
     */
    public readonly msiEndpoint!: pulumi.Output<string | undefined>;
    /**
     * The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
     * Connect.
     */
    public readonly oidcRequestToken!: pulumi.Output<string | undefined>;
    /**
     * The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
     * using OpenID Connect.
     */
    public readonly oidcRequestUrl!: pulumi.Output<string | undefined>;
    /**
     * The ID token for use when authenticating as a Service Principal using OpenID Connect.
     */
    public readonly oidcToken!: pulumi.Output<string | undefined>;
    /**
     * The path to a file containing an ID token for use when authenticating as a Service Principal using OpenID Connect.
     */
    public readonly oidcTokenFilePath!: pulumi.Output<string | undefined>;
    /**
     * A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
     */
    public readonly partnerId!: pulumi.Output<string | undefined>;
    /**
     * The Tenant ID which should be used. Works with all authentication methods except Managed Identity
     */
    public readonly tenantId!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["adoPipelineServiceConnectionId"] = args ? args.adoPipelineServiceConnectionId : undefined;
            resourceInputs["clientCertificate"] = args ? args.clientCertificate : undefined;
            resourceInputs["clientCertificatePassword"] = args?.clientCertificatePassword ? pulumi.secret(args.clientCertificatePassword) : undefined;
            resourceInputs["clientCertificatePath"] = args ? args.clientCertificatePath : undefined;
            resourceInputs["clientId"] = args?.clientId ? pulumi.secret(args.clientId) : undefined;
            resourceInputs["clientIdFilePath"] = args ? args.clientIdFilePath : undefined;
            resourceInputs["clientSecret"] = args?.clientSecret ? pulumi.secret(args.clientSecret) : undefined;
            resourceInputs["clientSecretFilePath"] = args ? args.clientSecretFilePath : undefined;
            resourceInputs["disableTerraformPartnerId"] = pulumi.output(args ? args.disableTerraformPartnerId : undefined).apply(JSON.stringify);
            resourceInputs["environment"] = (args ? args.environment : undefined) ?? (utilities.getEnv("ARM_ENVIRONMENT") || "public");
            resourceInputs["metadataHost"] = args ? args.metadataHost : undefined;
            resourceInputs["msiEndpoint"] = (args ? args.msiEndpoint : undefined) ?? utilities.getEnv("ARM_MSI_ENDPOINT");
            resourceInputs["oidcRequestToken"] = args ? args.oidcRequestToken : undefined;
            resourceInputs["oidcRequestUrl"] = args ? args.oidcRequestUrl : undefined;
            resourceInputs["oidcToken"] = args ? args.oidcToken : undefined;
            resourceInputs["oidcTokenFilePath"] = args ? args.oidcTokenFilePath : undefined;
            resourceInputs["partnerId"] = args ? args.partnerId : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["useAksWorkloadIdentity"] = pulumi.output(args ? args.useAksWorkloadIdentity : undefined).apply(JSON.stringify);
            resourceInputs["useCli"] = pulumi.output(args ? args.useCli : undefined).apply(JSON.stringify);
            resourceInputs["useMsi"] = pulumi.output((args ? args.useMsi : undefined) ?? (utilities.getEnvBoolean("ARM_USE_MSI") || false)).apply(JSON.stringify);
            resourceInputs["useOidc"] = pulumi.output(args ? args.useOidc : undefined).apply(JSON.stringify);
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["clientCertificatePassword", "clientId", "clientSecret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * The Azure DevOps Pipeline Service Connection ID.
     */
    adoPipelineServiceConnectionId?: pulumi.Input<string>;
    /**
     * Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate
     */
    clientCertificate?: pulumi.Input<string>;
    /**
     * The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
     * Certificate
     */
    clientCertificatePassword?: pulumi.Input<string>;
    /**
     * The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
     * Principal using a Client Certificate
     */
    clientCertificatePath?: pulumi.Input<string>;
    /**
     * The Client ID which should be used for service principal authentication
     */
    clientId?: pulumi.Input<string>;
    /**
     * The path to a file containing the Client ID which should be used for service principal authentication
     */
    clientIdFilePath?: pulumi.Input<string>;
    /**
     * The application password to use when authenticating as a Service Principal using a Client Secret
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * The path to a file containing the application password to use when authenticating as a Service Principal using a Client
     * Secret
     */
    clientSecretFilePath?: pulumi.Input<string>;
    disableTerraformPartnerId?: pulumi.Input<boolean>;
    /**
     * The cloud environment which should be used. Possible values are: `global` (also `public`), `usgovernmentl4` (also
     * `usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`. Not used and should not be specified
     * when `metadataHost` is specified.
     */
    environment?: pulumi.Input<string>;
    /**
     * The Hostname which should be used for the Azure Metadata Service.
     */
    metadataHost?: pulumi.Input<string>;
    /**
     * The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
     */
    msiEndpoint?: pulumi.Input<string>;
    /**
     * The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
     * Connect.
     */
    oidcRequestToken?: pulumi.Input<string>;
    /**
     * The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
     * using OpenID Connect.
     */
    oidcRequestUrl?: pulumi.Input<string>;
    /**
     * The ID token for use when authenticating as a Service Principal using OpenID Connect.
     */
    oidcToken?: pulumi.Input<string>;
    /**
     * The path to a file containing an ID token for use when authenticating as a Service Principal using OpenID Connect.
     */
    oidcTokenFilePath?: pulumi.Input<string>;
    /**
     * A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
     */
    partnerId?: pulumi.Input<string>;
    /**
     * The Tenant ID which should be used. Works with all authentication methods except Managed Identity
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Allow Azure AKS Workload Identity to be used for Authentication.
     */
    useAksWorkloadIdentity?: pulumi.Input<boolean>;
    /**
     * Allow Azure CLI to be used for Authentication
     */
    useCli?: pulumi.Input<boolean>;
    /**
     * Allow Managed Identity to be used for Authentication
     */
    useMsi?: pulumi.Input<boolean>;
    /**
     * Allow OpenID Connect to be used for authentication
     */
    useOidc?: pulumi.Input<boolean>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("azuread");

/**
 * Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate
 */
export declare const clientCertificate: string | undefined;
Object.defineProperty(exports, "clientCertificate", {
    get() {
        return __config.get("clientCertificate");
    },
    enumerable: true,
});

/**
 * The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
 * Certificate
 */
export declare const clientCertificatePassword: string | undefined;
Object.defineProperty(exports, "clientCertificatePassword", {
    get() {
        return __config.get("clientCertificatePassword");
    },
    enumerable: true,
});

/**
 * The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
 * Principal using a Client Certificate
 */
export declare const clientCertificatePath: string | undefined;
Object.defineProperty(exports, "clientCertificatePath", {
    get() {
        return __config.get("clientCertificatePath");
    },
    enumerable: true,
});

/**
 * The Client ID which should be used for service principal authentication
 */
export declare const clientId: string | undefined;
Object.defineProperty(exports, "clientId", {
    get() {
        return __config.get("clientId");
    },
    enumerable: true,
});

/**
 * The path to a file containing the Client ID which should be used for service principal authentication
 */
export declare const clientIdFilePath: string | undefined;
Object.defineProperty(exports, "clientIdFilePath", {
    get() {
        return __config.get("clientIdFilePath");
    },
    enumerable: true,
});

/**
 * The application password to use when authenticating as a Service Principal using a Client Secret
 */
export declare const clientSecret: string | undefined;
Object.defineProperty(exports, "clientSecret", {
    get() {
        return __config.get("clientSecret");
    },
    enumerable: true,
});

/**
 * The path to a file containing the application password to use when authenticating as a Service Principal using a Client
 * Secret
 */
export declare const clientSecretFilePath: string | undefined;
Object.defineProperty(exports, "clientSecretFilePath", {
    get() {
        return __config.get("clientSecretFilePath");
    },
    enumerable: true,
});

/**
 * Disable the Terraform Partner ID, which is used if a custom `partner_id` isn't specified
 */
export declare const disableTerraformPartnerId: boolean | undefined;
Object.defineProperty(exports, "disableTerraformPartnerId", {
    get() {
        return __config.getObject<boolean>("disableTerraformPartnerId");
    },
    enumerable: true,
});

/**
 * The cloud environment which should be used. Possible values are: `global` (also `public`), `usgovernmentl4` (also
 * `usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`
 */
export declare const environment: string;
Object.defineProperty(exports, "environment", {
    get() {
        return __config.get("environment") ?? (utilities.getEnv("ARM_ENVIRONMENT") || "public");
    },
    enumerable: true,
});

/**
 * The Hostname which should be used for the Azure Metadata Service.
 */
export declare const metadataHost: string | undefined;
Object.defineProperty(exports, "metadataHost", {
    get() {
        return __config.get("metadataHost");
    },
    enumerable: true,
});

/**
 * The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
 */
export declare const msiEndpoint: string | undefined;
Object.defineProperty(exports, "msiEndpoint", {
    get() {
        return __config.get("msiEndpoint") ?? utilities.getEnv("ARM_MSI_ENDPOINT");
    },
    enumerable: true,
});

/**
 * The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
 * Connect.
 */
export declare const oidcRequestToken: string | undefined;
Object.defineProperty(exports, "oidcRequestToken", {
    get() {
        return __config.get("oidcRequestToken");
    },
    enumerable: true,
});

/**
 * The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
 * using OpenID Connect.
 */
export declare const oidcRequestUrl: string | undefined;
Object.defineProperty(exports, "oidcRequestUrl", {
    get() {
        return __config.get("oidcRequestUrl");
    },
    enumerable: true,
});

/**
 * The ID token for use when authenticating as a Service Principal using OpenID Connect.
 */
export declare const oidcToken: string | undefined;
Object.defineProperty(exports, "oidcToken", {
    get() {
        return __config.get("oidcToken");
    },
    enumerable: true,
});

/**
 * The path to a file containing an ID token for use when authenticating as a Service Principal using OpenID Connect.
 */
export declare const oidcTokenFilePath: string | undefined;
Object.defineProperty(exports, "oidcTokenFilePath", {
    get() {
        return __config.get("oidcTokenFilePath");
    },
    enumerable: true,
});

/**
 * A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
 */
export declare const partnerId: string | undefined;
Object.defineProperty(exports, "partnerId", {
    get() {
        return __config.get("partnerId");
    },
    enumerable: true,
});

/**
 * The Tenant ID which should be used. Works with all authentication methods except Managed Identity
 */
export declare const tenantId: string | undefined;
Object.defineProperty(exports, "tenantId", {
    get() {
        return __config.get("tenantId");
    },
    enumerable: true,
});

/**
 * Allow Azure CLI to be used for Authentication
 */
export declare const useCli: boolean | undefined;
Object.defineProperty(exports, "useCli", {
    get() {
        return __config.getObject<boolean>("useCli");
    },
    enumerable: true,
});

/**
 * Allow Managed Identity to be used for Authentication
 */
export declare const useMsi: boolean;
Object.defineProperty(exports, "useMsi", {
    get() {
        return __config.getObject<boolean>("useMsi") ?? (utilities.getEnvBoolean("ARM_USE_MSI") || false);
    },
    enumerable: true,
});

/**
 * Allow OpenID Connect to be used for authentication
 */
export declare const useOidc: boolean | undefined;
Object.defineProperty(exports, "useOidc", {
    get() {
        return __config.getObject<boolean>("useOidc");
    },
    enumerable: true,
});


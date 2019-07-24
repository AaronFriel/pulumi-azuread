// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about an existing Application within Azure Active Directory.
 * 
 * > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 * 
 * const example = pulumi.output(azuread.getApplication({
 *     name: "My First AzureAD Application",
 * }));
 * 
 * export const azureAdObjectId = example.id;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azuread/blob/master/website/docs/d/application.html.markdown.
 */
export function getApplication(args?: GetApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationResult> & GetApplicationResult {
    args = args || {};
    const promise: Promise<GetApplicationResult> = pulumi.runtime.invoke("azuread:index/getApplication:getApplication", {
        "appRoles": args.appRoles,
        "name": args.name,
        "oauth2Permissions": args.oauth2Permissions,
        "objectId": args.objectId,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getApplication.
 */
export interface GetApplicationArgs {
    readonly appRoles?: { allowedMemberTypes?: string[], description?: string, displayName?: string, id?: string, isEnabled?: boolean, value?: string }[];
    /**
     * Specifies the name of the Application within Azure Active Directory.
     */
    readonly name?: string;
    readonly oauth2Permissions?: { adminConsentDescription?: string, adminConsentDisplayName?: string, id?: string, isEnabled?: boolean, type?: string, userConsentDescription?: string, userConsentDisplayName?: string, value?: string }[];
    /**
     * Specifies the Object ID of the Application within Azure Active Directory.
     */
    readonly objectId?: string;
}

/**
 * A collection of values returned by getApplication.
 */
export interface GetApplicationResult {
    /**
     * A collection of `app_role` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
     */
    readonly appRoles: { allowedMemberTypes: string[], description: string, displayName: string, id: string, isEnabled: boolean, value: string }[];
    /**
     * the Application ID of the Azure Active Directory Application.
     */
    readonly applicationId: string;
    /**
     * Is this Azure AD Application available to other tenants?
     */
    readonly availableToOtherTenants: boolean;
    /**
     * The `groups` claim issued in a user or OAuth 2.0 access token that the app expects.
     */
    readonly groupMembershipClaims: string;
    readonly homepage: string;
    /**
     * A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
     */
    readonly identifierUris: string[];
    readonly name: string;
    /**
     * Does this Azure AD Application allow OAuth2.0 implicit flow tokens?
     */
    readonly oauth2AllowImplicitFlow: boolean;
    /**
     * A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a `oauth2_permission` block as documented below.
     */
    readonly oauth2Permissions: { adminConsentDescription: string, adminConsentDisplayName: string, id: string, isEnabled: boolean, type: string, userConsentDescription: string, userConsentDisplayName: string, value: string }[];
    /**
     * the Object ID of the Azure Active Directory Application.
     */
    readonly objectId: string;
    /**
     * A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
     */
    readonly replyUrls: string[];
    /**
     * A collection of `required_resource_access` blocks as documented below.
     */
    readonly requiredResourceAccesses: { resourceAccesses: { id: string, type: string }[], resourceAppId: string }[];
    /**
     * The type of the permission
     */
    readonly type: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}

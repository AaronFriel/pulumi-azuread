// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about an Azure Active Directory user.
 * 
 * > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 * 
 * const example = pulumi.output(azuread.getUser({
 *     userPrincipalName: "user@hashicorp.com",
 * }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azuread/blob/master/website/docs/d/user.html.markdown.
 */
export function getUser(args?: GetUserArgs, opts?: pulumi.InvokeOptions): Promise<GetUserResult> & GetUserResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetUserResult> = pulumi.runtime.invoke("azuread:index/getUser:getUser", {
        "objectId": args.objectId,
        "userPrincipalName": args.userPrincipalName,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserArgs {
    /**
     * Specifies the Object ID of the Application within Azure Active Directory.
     */
    readonly objectId?: string;
    /**
     * The User Principal Name of the Azure AD User.
     */
    readonly userPrincipalName?: string;
}

/**
 * A collection of values returned by getUser.
 */
export interface GetUserResult {
    /**
     * `True` if the account is enabled; otherwise `False`.
     */
    readonly accountEnabled: boolean;
    /**
     * The Display Name of the Azure AD User.
     */
    readonly displayName: string;
    /**
     * The primary email address of the Azure AD User.
     */
    readonly mail: string;
    /**
     * The email alias of the Azure AD User.
     */
    readonly mailNickname: string;
    readonly objectId: string;
    /**
     * The User Principal Name of the Azure AD User.
     */
    readonly userPrincipalName: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}

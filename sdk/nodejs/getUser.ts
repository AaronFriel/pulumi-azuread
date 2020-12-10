// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
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
 * }, { async: true }));
 * ```
 */
export function getUser(args?: GetUserArgs, opts?: pulumi.InvokeOptions): Promise<GetUserResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azuread:index/getUser:getUser", {
        "mailNickname": args.mailNickname,
        "objectId": args.objectId,
        "userPrincipalName": args.userPrincipalName,
    }, opts);
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserArgs {
    /**
     * The email alias of the Azure AD User.
     */
    readonly mailNickname?: string;
    /**
     * Specifies the Object ID of the User within Azure Active Directory.
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
     * The city in which the user is located.
     */
    readonly city: string;
    /**
     * The company name which the user is associated. This property can be useful for describing the company that an external user comes from.
     */
    readonly companyName: string;
    /**
     * The country/region in which the user is located; for example, “US” or “UK”.
     */
    readonly country: string;
    /**
     * The name for the department in which the user works.
     */
    readonly department: string;
    /**
     * The Display Name of the Azure AD User.
     */
    readonly displayName: string;
    /**
     * The given name (first name) of the user.
     */
    readonly givenName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The value used to associate an on-premise Active Directory user account with their Azure AD user object.
     */
    readonly immutableId: string;
    /**
     * The user’s job title.
     */
    readonly jobTitle: string;
    /**
     * The primary email address of the Azure AD User.
     */
    readonly mail: string;
    /**
     * The email alias of the Azure AD User.
     */
    readonly mailNickname: string;
    /**
     * The primary cellular telephone number for the user.
     */
    readonly mobile: string;
    readonly objectId: string;
    /**
     * The on-premise SAM account name of the Azure AD User.
     */
    readonly onpremisesSamAccountName: string;
    /**
     * The on-premise user principal name of the Azure AD User.
     */
    readonly onpremisesUserPrincipalName: string;
    /**
     * The office location in the user's place of business.
     */
    readonly physicalDeliveryOfficeName: string;
    /**
     * The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.
     */
    readonly postalCode: string;
    /**
     * The state or province in the user's address.
     */
    readonly state: string;
    /**
     * The street address of the user's place of business.
     */
    readonly streetAddress: string;
    /**
     * The user's surname (family name or last name).
     */
    readonly surname: string;
    /**
     * The usage location of the Azure AD User.
     */
    readonly usageLocation: string;
    /**
     * The User Principal Name of the Azure AD User.
     */
    readonly userPrincipalName: string;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to access information about an existing Domains within Azure Active Directory.
 * 
 * > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Directory.Read.All` within the `Windows Azure Active Directory` API.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 * 
 * const aadDomains = azuread.getDomains();
 * 
 * export const domains = aadDomains.domains;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azuread/blob/master/website/docs/d/domains.html.markdown.
 */
export function getDomains(args?: GetDomainsArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azuread:index/getDomains:getDomains", {
        "includeUnverified": args.includeUnverified,
        "onlyDefault": args.onlyDefault,
        "onlyInitial": args.onlyInitial,
    }, opts);
}

/**
 * A collection of arguments for invoking getDomains.
 */
export interface GetDomainsArgs {
    /**
     * Set to `true` if unverified Azure AD Domains should be included. Defaults to `false`.
     */
    readonly includeUnverified?: boolean;
    /**
     * Set to `true` to only return the default domain.
     */
    readonly onlyDefault?: boolean;
    /**
     * Set to `true` to only return the initial domain, which is your primary Azure Active Directory tenant domain. Defaults to `false`.
     */
    readonly onlyInitial?: boolean;
}

/**
 * A collection of values returned by getDomains.
 */
export interface GetDomainsResult {
    /**
     * One or more `domain` blocks as defined below.
     */
    readonly domains: outputs.GetDomainsDomain[];
    readonly includeUnverified?: boolean;
    readonly onlyDefault?: boolean;
    readonly onlyInitial?: boolean;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}

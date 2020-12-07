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
        return obj['__pulumiType'] === Provider.__pulumiType;
    }


    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        {
            if ((!args || args.metadataHost === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'metadataHost'");
            }
            inputs["clientCertificatePassword"] = (args ? args.clientCertificatePassword : undefined) || (utilities.getEnv("ARM_CLIENT_CERTIFICATE_PASSWORD") || "");
            inputs["clientCertificatePath"] = (args ? args.clientCertificatePath : undefined) || (utilities.getEnv("ARM_CLIENT_CERTIFICATE_PATH") || "");
            inputs["clientId"] = (args ? args.clientId : undefined) || (utilities.getEnv("ARM_CLIENT_ID") || "");
            inputs["clientSecret"] = (args ? args.clientSecret : undefined) || (utilities.getEnv("ARM_CLIENT_SECRET") || "");
            inputs["environment"] = (args ? args.environment : undefined) || (utilities.getEnv("ARM_ENVIRONMENT") || "public");
            inputs["metadataHost"] = args ? args.metadataHost : undefined;
            inputs["msiEndpoint"] = (args ? args.msiEndpoint : undefined) || (utilities.getEnv("ARM_MSI_ENDPOINT") || "");
            inputs["tenantId"] = (args ? args.tenantId : undefined) || (utilities.getEnv("ARM_TENANT_ID") || "");
            inputs["useMsi"] = pulumi.output((args ? args.useMsi : undefined) || (<any>utilities.getEnvBoolean("ARM_USE_MSI") || false)).apply(JSON.stringify);
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Provider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    readonly clientCertificatePassword?: pulumi.Input<string>;
    readonly clientCertificatePath?: pulumi.Input<string>;
    readonly clientId?: pulumi.Input<string>;
    readonly clientSecret?: pulumi.Input<string>;
    readonly environment?: pulumi.Input<string>;
    /**
     * The Hostname which should be used to fetch environment metadata from.
     */
    readonly metadataHost: pulumi.Input<string>;
    readonly msiEndpoint?: pulumi.Input<string>;
    readonly tenantId?: pulumi.Input<string>;
    readonly useMsi?: pulumi.Input<boolean>;
}

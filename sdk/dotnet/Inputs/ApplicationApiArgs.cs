// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Inputs
{

    public sealed class ApplicationApiArgs : Pulumi.ResourceArgs
    {
        [Input("knownClientApplications")]
        private InputList<string>? _knownClientApplications;

        /// <summary>
        /// A set of application IDs (client IDs), used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app.
        /// </summary>
        public InputList<string> KnownClientApplications
        {
            get => _knownClientApplications ?? (_knownClientApplications = new InputList<string>());
            set => _knownClientApplications = value;
        }

        /// <summary>
        /// Allows an application to use claims mapping without specifying a custom signing key. Defaults to `false`.
        /// </summary>
        [Input("mappedClaimsEnabled")]
        public Input<bool>? MappedClaimsEnabled { get; set; }

        [Input("oauth2PermissionScopes")]
        private InputList<Inputs.ApplicationApiOauth2PermissionScopeArgs>? _oauth2PermissionScopes;

        /// <summary>
        /// One or more `oauth2_permission_scope` blocks as documented below, to describe delegated permissions exposed by the web API represented by this application.
        /// </summary>
        public InputList<Inputs.ApplicationApiOauth2PermissionScopeArgs> Oauth2PermissionScopes
        {
            get => _oauth2PermissionScopes ?? (_oauth2PermissionScopes = new InputList<Inputs.ApplicationApiOauth2PermissionScopeArgs>());
            set => _oauth2PermissionScopes = value;
        }

        /// <summary>
        /// The access token version expected by this resource. Must be one of `1` or `2`, and must be `2` when `sign_in_audience` is either `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount` Defaults to `1`.
        /// </summary>
        [Input("requestedAccessTokenVersion")]
        public Input<int>? RequestedAccessTokenVersion { get; set; }

        public ApplicationApiArgs()
        {
        }
    }
}

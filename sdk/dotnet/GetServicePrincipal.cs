// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    public static class GetServicePrincipal
    {
        /// <summary>
        /// Gets information about an existing service principal associated with an application within Azure Active Directory.
        /// 
        /// ## API Permissions
        /// 
        /// The following API permissions are required in order to use this data source.
        /// 
        /// When authenticated with a service principal, this data source requires one of the following application roles: `Application.Read.All` or `Directory.Read.All`
        /// 
        /// When authenticated with a user principal, this data source does not require any additional roles.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// *Look up by application display name*
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(AzureAD.GetServicePrincipal.InvokeAsync(new AzureAD.GetServicePrincipalArgs
        ///         {
        ///             DisplayName = "my-awesome-application",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// *Look up by application ID (client ID)*
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(AzureAD.GetServicePrincipal.InvokeAsync(new AzureAD.GetServicePrincipalArgs
        ///         {
        ///             ApplicationId = "00000000-0000-0000-0000-000000000000",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// *Look up by service principal object ID*
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(AzureAD.GetServicePrincipal.InvokeAsync(new AzureAD.GetServicePrincipalArgs
        ///         {
        ///             ObjectId = "00000000-0000-0000-0000-000000000000",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServicePrincipalResult> InvokeAsync(GetServicePrincipalArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServicePrincipalResult>("azuread:index/getServicePrincipal:getServicePrincipal", args ?? new GetServicePrincipalArgs(), options.WithVersion());
    }


    public sealed class GetServicePrincipalArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The application ID (client ID) of the application associated with this service principal.
        /// </summary>
        [Input("applicationId")]
        public string? ApplicationId { get; set; }

        /// <summary>
        /// The display name of the application associated with this service principal.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        /// <summary>
        /// The object ID of the service principal.
        /// </summary>
        [Input("objectId")]
        public string? ObjectId { get; set; }

        public GetServicePrincipalArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServicePrincipalResult
    {
        /// <summary>
        /// Whether or not the service principal account is enabled.
        /// </summary>
        public readonly bool AccountEnabled;
        /// <summary>
        /// A list of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
        /// </summary>
        public readonly ImmutableArray<string> AlternativeNames;
        /// <summary>
        /// Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application.
        /// </summary>
        public readonly bool AppRoleAssignmentRequired;
        /// <summary>
        /// A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
        /// </summary>
        public readonly ImmutableDictionary<string, string> AppRoleIds;
        /// <summary>
        /// A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServicePrincipalAppRoleResult> AppRoles;
        /// <summary>
        /// The application ID (client ID) of the application associated with this service principal.
        /// </summary>
        public readonly string ApplicationId;
        /// <summary>
        /// The tenant ID where the associated application is registered.
        /// </summary>
        public readonly string ApplicationTenantId;
        /// <summary>
        /// Permission help text that appears in the admin app assignment and consent experiences.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Display name for the permission that appears in the admin consent and app assignment experiences.
        /// </summary>
        public readonly string DisplayName;
        public readonly ImmutableArray<Outputs.GetServicePrincipalFeatureTagResult> FeatureTags;
        /// <summary>
        /// A `features` block as described below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServicePrincipalFeatureResult> Features;
        /// <summary>
        /// Home page or landing page of the associated application.
        /// </summary>
        public readonly string HomepageUrl;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps.
        /// </summary>
        public readonly string LoginUrl;
        /// <summary>
        /// The URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
        /// </summary>
        public readonly string LogoutUrl;
        /// <summary>
        /// A free text field to capture information about the service principal, typically used for operational purposes.
        /// </summary>
        public readonly string Notes;
        /// <summary>
        /// A list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
        /// </summary>
        public readonly ImmutableArray<string> NotificationEmailAddresses;
        /// <summary>
        /// A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Oauth2PermissionScopeIds;
        /// <summary>
        /// A collection of OAuth 2.0 delegated permissions exposed by the associated application. Each permission is covered by an `oauth2_permission_scopes` block as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServicePrincipalOauth2PermissionScopeResult> Oauth2PermissionScopes;
        /// <summary>
        /// The object ID of the service principal.
        /// </summary>
        public readonly string ObjectId;
        /// <summary>
        /// The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps.
        /// </summary>
        public readonly string PreferredSingleSignOnMode;
        /// <summary>
        /// A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
        /// </summary>
        public readonly ImmutableArray<string> RedirectUris;
        /// <summary>
        /// The URL where the service exposes SAML metadata for federation.
        /// </summary>
        public readonly string SamlMetadataUrl;
        /// <summary>
        /// A `saml_single_sign_on` block as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServicePrincipalSamlSingleSignOnResult> SamlSingleSignOns;
        /// <summary>
        /// A list of identifier URI(s), copied over from the associated application.
        /// </summary>
        public readonly ImmutableArray<string> ServicePrincipalNames;
        /// <summary>
        /// The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
        /// </summary>
        public readonly string SignInAudience;
        /// <summary>
        /// A list of tags applied to the service principal.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetServicePrincipalResult(
            bool accountEnabled,

            ImmutableArray<string> alternativeNames,

            bool appRoleAssignmentRequired,

            ImmutableDictionary<string, string> appRoleIds,

            ImmutableArray<Outputs.GetServicePrincipalAppRoleResult> appRoles,

            string applicationId,

            string applicationTenantId,

            string description,

            string displayName,

            ImmutableArray<Outputs.GetServicePrincipalFeatureTagResult> featureTags,

            ImmutableArray<Outputs.GetServicePrincipalFeatureResult> features,

            string homepageUrl,

            string id,

            string loginUrl,

            string logoutUrl,

            string notes,

            ImmutableArray<string> notificationEmailAddresses,

            ImmutableDictionary<string, string> oauth2PermissionScopeIds,

            ImmutableArray<Outputs.GetServicePrincipalOauth2PermissionScopeResult> oauth2PermissionScopes,

            string objectId,

            string preferredSingleSignOnMode,

            ImmutableArray<string> redirectUris,

            string samlMetadataUrl,

            ImmutableArray<Outputs.GetServicePrincipalSamlSingleSignOnResult> samlSingleSignOns,

            ImmutableArray<string> servicePrincipalNames,

            string signInAudience,

            ImmutableArray<string> tags,

            string type)
        {
            AccountEnabled = accountEnabled;
            AlternativeNames = alternativeNames;
            AppRoleAssignmentRequired = appRoleAssignmentRequired;
            AppRoleIds = appRoleIds;
            AppRoles = appRoles;
            ApplicationId = applicationId;
            ApplicationTenantId = applicationTenantId;
            Description = description;
            DisplayName = displayName;
            FeatureTags = featureTags;
            Features = features;
            HomepageUrl = homepageUrl;
            Id = id;
            LoginUrl = loginUrl;
            LogoutUrl = logoutUrl;
            Notes = notes;
            NotificationEmailAddresses = notificationEmailAddresses;
            Oauth2PermissionScopeIds = oauth2PermissionScopeIds;
            Oauth2PermissionScopes = oauth2PermissionScopes;
            ObjectId = objectId;
            PreferredSingleSignOnMode = preferredSingleSignOnMode;
            RedirectUris = redirectUris;
            SamlMetadataUrl = samlMetadataUrl;
            SamlSingleSignOns = samlSingleSignOns;
            ServicePrincipalNames = servicePrincipalNames;
            SignInAudience = signInAudience;
            Tags = tags;
            Type = type;
        }
    }
}

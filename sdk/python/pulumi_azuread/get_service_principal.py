# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetServicePrincipalResult',
    'AwaitableGetServicePrincipalResult',
    'get_service_principal',
    'get_service_principal_output',
]

@pulumi.output_type
class GetServicePrincipalResult:
    """
    A collection of values returned by getServicePrincipal.
    """
    def __init__(__self__, account_enabled=None, alternative_names=None, app_role_assignment_required=None, app_role_ids=None, app_roles=None, application_id=None, application_tenant_id=None, description=None, display_name=None, feature_tags=None, features=None, homepage_url=None, id=None, login_url=None, logout_url=None, notes=None, notification_email_addresses=None, oauth2_permission_scope_ids=None, oauth2_permission_scopes=None, object_id=None, preferred_single_sign_on_mode=None, redirect_uris=None, saml_metadata_url=None, saml_single_sign_ons=None, service_principal_names=None, sign_in_audience=None, tags=None, type=None):
        if account_enabled and not isinstance(account_enabled, bool):
            raise TypeError("Expected argument 'account_enabled' to be a bool")
        pulumi.set(__self__, "account_enabled", account_enabled)
        if alternative_names and not isinstance(alternative_names, list):
            raise TypeError("Expected argument 'alternative_names' to be a list")
        pulumi.set(__self__, "alternative_names", alternative_names)
        if app_role_assignment_required and not isinstance(app_role_assignment_required, bool):
            raise TypeError("Expected argument 'app_role_assignment_required' to be a bool")
        pulumi.set(__self__, "app_role_assignment_required", app_role_assignment_required)
        if app_role_ids and not isinstance(app_role_ids, dict):
            raise TypeError("Expected argument 'app_role_ids' to be a dict")
        pulumi.set(__self__, "app_role_ids", app_role_ids)
        if app_roles and not isinstance(app_roles, list):
            raise TypeError("Expected argument 'app_roles' to be a list")
        pulumi.set(__self__, "app_roles", app_roles)
        if application_id and not isinstance(application_id, str):
            raise TypeError("Expected argument 'application_id' to be a str")
        pulumi.set(__self__, "application_id", application_id)
        if application_tenant_id and not isinstance(application_tenant_id, str):
            raise TypeError("Expected argument 'application_tenant_id' to be a str")
        pulumi.set(__self__, "application_tenant_id", application_tenant_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if feature_tags and not isinstance(feature_tags, list):
            raise TypeError("Expected argument 'feature_tags' to be a list")
        pulumi.set(__self__, "feature_tags", feature_tags)
        if features and not isinstance(features, list):
            raise TypeError("Expected argument 'features' to be a list")
        if features is not None:
            warnings.warn("""This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider""", DeprecationWarning)
            pulumi.log.warn("""features is deprecated: This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider""")

        pulumi.set(__self__, "features", features)
        if homepage_url and not isinstance(homepage_url, str):
            raise TypeError("Expected argument 'homepage_url' to be a str")
        pulumi.set(__self__, "homepage_url", homepage_url)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if login_url and not isinstance(login_url, str):
            raise TypeError("Expected argument 'login_url' to be a str")
        pulumi.set(__self__, "login_url", login_url)
        if logout_url and not isinstance(logout_url, str):
            raise TypeError("Expected argument 'logout_url' to be a str")
        pulumi.set(__self__, "logout_url", logout_url)
        if notes and not isinstance(notes, str):
            raise TypeError("Expected argument 'notes' to be a str")
        pulumi.set(__self__, "notes", notes)
        if notification_email_addresses and not isinstance(notification_email_addresses, list):
            raise TypeError("Expected argument 'notification_email_addresses' to be a list")
        pulumi.set(__self__, "notification_email_addresses", notification_email_addresses)
        if oauth2_permission_scope_ids and not isinstance(oauth2_permission_scope_ids, dict):
            raise TypeError("Expected argument 'oauth2_permission_scope_ids' to be a dict")
        pulumi.set(__self__, "oauth2_permission_scope_ids", oauth2_permission_scope_ids)
        if oauth2_permission_scopes and not isinstance(oauth2_permission_scopes, list):
            raise TypeError("Expected argument 'oauth2_permission_scopes' to be a list")
        pulumi.set(__self__, "oauth2_permission_scopes", oauth2_permission_scopes)
        if object_id and not isinstance(object_id, str):
            raise TypeError("Expected argument 'object_id' to be a str")
        pulumi.set(__self__, "object_id", object_id)
        if preferred_single_sign_on_mode and not isinstance(preferred_single_sign_on_mode, str):
            raise TypeError("Expected argument 'preferred_single_sign_on_mode' to be a str")
        pulumi.set(__self__, "preferred_single_sign_on_mode", preferred_single_sign_on_mode)
        if redirect_uris and not isinstance(redirect_uris, list):
            raise TypeError("Expected argument 'redirect_uris' to be a list")
        pulumi.set(__self__, "redirect_uris", redirect_uris)
        if saml_metadata_url and not isinstance(saml_metadata_url, str):
            raise TypeError("Expected argument 'saml_metadata_url' to be a str")
        pulumi.set(__self__, "saml_metadata_url", saml_metadata_url)
        if saml_single_sign_ons and not isinstance(saml_single_sign_ons, list):
            raise TypeError("Expected argument 'saml_single_sign_ons' to be a list")
        pulumi.set(__self__, "saml_single_sign_ons", saml_single_sign_ons)
        if service_principal_names and not isinstance(service_principal_names, list):
            raise TypeError("Expected argument 'service_principal_names' to be a list")
        pulumi.set(__self__, "service_principal_names", service_principal_names)
        if sign_in_audience and not isinstance(sign_in_audience, str):
            raise TypeError("Expected argument 'sign_in_audience' to be a str")
        pulumi.set(__self__, "sign_in_audience", sign_in_audience)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="accountEnabled")
    def account_enabled(self) -> bool:
        """
        Whether or not the service principal account is enabled.
        """
        return pulumi.get(self, "account_enabled")

    @property
    @pulumi.getter(name="alternativeNames")
    def alternative_names(self) -> Sequence[str]:
        """
        A list of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
        """
        return pulumi.get(self, "alternative_names")

    @property
    @pulumi.getter(name="appRoleAssignmentRequired")
    def app_role_assignment_required(self) -> bool:
        """
        Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application.
        """
        return pulumi.get(self, "app_role_assignment_required")

    @property
    @pulumi.getter(name="appRoleIds")
    def app_role_ids(self) -> Mapping[str, str]:
        """
        A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
        """
        return pulumi.get(self, "app_role_ids")

    @property
    @pulumi.getter(name="appRoles")
    def app_roles(self) -> Sequence['outputs.GetServicePrincipalAppRoleResult']:
        """
        A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
        """
        return pulumi.get(self, "app_roles")

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> str:
        """
        The application ID (client ID) of the application associated with this service principal.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="applicationTenantId")
    def application_tenant_id(self) -> str:
        """
        The tenant ID where the associated application is registered.
        """
        return pulumi.get(self, "application_tenant_id")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Permission help text that appears in the admin app assignment and consent experiences.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Display name for the permission that appears in the admin consent and app assignment experiences.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="featureTags")
    def feature_tags(self) -> Sequence['outputs.GetServicePrincipalFeatureTagResult']:
        return pulumi.get(self, "feature_tags")

    @property
    @pulumi.getter
    def features(self) -> Sequence['outputs.GetServicePrincipalFeatureResult']:
        """
        A `features` block as described below.
        """
        return pulumi.get(self, "features")

    @property
    @pulumi.getter(name="homepageUrl")
    def homepage_url(self) -> str:
        """
        Home page or landing page of the associated application.
        """
        return pulumi.get(self, "homepage_url")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="loginUrl")
    def login_url(self) -> str:
        """
        The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps.
        """
        return pulumi.get(self, "login_url")

    @property
    @pulumi.getter(name="logoutUrl")
    def logout_url(self) -> str:
        """
        The URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
        """
        return pulumi.get(self, "logout_url")

    @property
    @pulumi.getter
    def notes(self) -> str:
        """
        A free text field to capture information about the service principal, typically used for operational purposes.
        """
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter(name="notificationEmailAddresses")
    def notification_email_addresses(self) -> Sequence[str]:
        """
        A list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
        """
        return pulumi.get(self, "notification_email_addresses")

    @property
    @pulumi.getter(name="oauth2PermissionScopeIds")
    def oauth2_permission_scope_ids(self) -> Mapping[str, str]:
        """
        A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.
        """
        return pulumi.get(self, "oauth2_permission_scope_ids")

    @property
    @pulumi.getter(name="oauth2PermissionScopes")
    def oauth2_permission_scopes(self) -> Sequence['outputs.GetServicePrincipalOauth2PermissionScopeResult']:
        """
        A collection of OAuth 2.0 delegated permissions exposed by the associated application. Each permission is covered by an `oauth2_permission_scopes` block as documented below.
        """
        return pulumi.get(self, "oauth2_permission_scopes")

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> str:
        """
        The object ID of the service principal.
        """
        return pulumi.get(self, "object_id")

    @property
    @pulumi.getter(name="preferredSingleSignOnMode")
    def preferred_single_sign_on_mode(self) -> str:
        """
        The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps.
        """
        return pulumi.get(self, "preferred_single_sign_on_mode")

    @property
    @pulumi.getter(name="redirectUris")
    def redirect_uris(self) -> Sequence[str]:
        """
        A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
        """
        return pulumi.get(self, "redirect_uris")

    @property
    @pulumi.getter(name="samlMetadataUrl")
    def saml_metadata_url(self) -> str:
        """
        The URL where the service exposes SAML metadata for federation.
        """
        return pulumi.get(self, "saml_metadata_url")

    @property
    @pulumi.getter(name="samlSingleSignOns")
    def saml_single_sign_ons(self) -> Sequence['outputs.GetServicePrincipalSamlSingleSignOnResult']:
        """
        A `saml_single_sign_on` block as documented below.
        """
        return pulumi.get(self, "saml_single_sign_ons")

    @property
    @pulumi.getter(name="servicePrincipalNames")
    def service_principal_names(self) -> Sequence[str]:
        """
        A list of identifier URI(s), copied over from the associated application.
        """
        return pulumi.get(self, "service_principal_names")

    @property
    @pulumi.getter(name="signInAudience")
    def sign_in_audience(self) -> str:
        """
        The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
        """
        return pulumi.get(self, "sign_in_audience")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        A list of tags applied to the service principal.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
        """
        return pulumi.get(self, "type")


class AwaitableGetServicePrincipalResult(GetServicePrincipalResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServicePrincipalResult(
            account_enabled=self.account_enabled,
            alternative_names=self.alternative_names,
            app_role_assignment_required=self.app_role_assignment_required,
            app_role_ids=self.app_role_ids,
            app_roles=self.app_roles,
            application_id=self.application_id,
            application_tenant_id=self.application_tenant_id,
            description=self.description,
            display_name=self.display_name,
            feature_tags=self.feature_tags,
            features=self.features,
            homepage_url=self.homepage_url,
            id=self.id,
            login_url=self.login_url,
            logout_url=self.logout_url,
            notes=self.notes,
            notification_email_addresses=self.notification_email_addresses,
            oauth2_permission_scope_ids=self.oauth2_permission_scope_ids,
            oauth2_permission_scopes=self.oauth2_permission_scopes,
            object_id=self.object_id,
            preferred_single_sign_on_mode=self.preferred_single_sign_on_mode,
            redirect_uris=self.redirect_uris,
            saml_metadata_url=self.saml_metadata_url,
            saml_single_sign_ons=self.saml_single_sign_ons,
            service_principal_names=self.service_principal_names,
            sign_in_audience=self.sign_in_audience,
            tags=self.tags,
            type=self.type)


def get_service_principal(application_id: Optional[str] = None,
                          display_name: Optional[str] = None,
                          object_id: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServicePrincipalResult:
    """
    Gets information about an existing service principal associated with an application within Azure Active Directory.

    ## API Permissions

    The following API permissions are required in order to use this data source.

    When authenticated with a service principal, this data source requires one of the following application roles: `Application.Read.All` or `Directory.Read.All`

    When authenticated with a user principal, this data source does not require any additional roles.

    ## Example Usage

    *Look up by application display name*

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_service_principal(display_name="my-awesome-application")
    ```

    *Look up by application ID (client ID)*

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_service_principal(application_id="00000000-0000-0000-0000-000000000000")
    ```

    *Look up by service principal object ID*

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_service_principal(object_id="00000000-0000-0000-0000-000000000000")
    ```


    :param str application_id: The application ID (client ID) of the application associated with this service principal.
    :param str display_name: The display name of the application associated with this service principal.
    :param str object_id: The object ID of the service principal.
    """
    __args__ = dict()
    __args__['applicationId'] = application_id
    __args__['displayName'] = display_name
    __args__['objectId'] = object_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azuread:index/getServicePrincipal:getServicePrincipal', __args__, opts=opts, typ=GetServicePrincipalResult).value

    return AwaitableGetServicePrincipalResult(
        account_enabled=__ret__.account_enabled,
        alternative_names=__ret__.alternative_names,
        app_role_assignment_required=__ret__.app_role_assignment_required,
        app_role_ids=__ret__.app_role_ids,
        app_roles=__ret__.app_roles,
        application_id=__ret__.application_id,
        application_tenant_id=__ret__.application_tenant_id,
        description=__ret__.description,
        display_name=__ret__.display_name,
        feature_tags=__ret__.feature_tags,
        features=__ret__.features,
        homepage_url=__ret__.homepage_url,
        id=__ret__.id,
        login_url=__ret__.login_url,
        logout_url=__ret__.logout_url,
        notes=__ret__.notes,
        notification_email_addresses=__ret__.notification_email_addresses,
        oauth2_permission_scope_ids=__ret__.oauth2_permission_scope_ids,
        oauth2_permission_scopes=__ret__.oauth2_permission_scopes,
        object_id=__ret__.object_id,
        preferred_single_sign_on_mode=__ret__.preferred_single_sign_on_mode,
        redirect_uris=__ret__.redirect_uris,
        saml_metadata_url=__ret__.saml_metadata_url,
        saml_single_sign_ons=__ret__.saml_single_sign_ons,
        service_principal_names=__ret__.service_principal_names,
        sign_in_audience=__ret__.sign_in_audience,
        tags=__ret__.tags,
        type=__ret__.type)


@_utilities.lift_output_func(get_service_principal)
def get_service_principal_output(application_id: Optional[pulumi.Input[Optional[str]]] = None,
                                 display_name: Optional[pulumi.Input[Optional[str]]] = None,
                                 object_id: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServicePrincipalResult]:
    """
    Gets information about an existing service principal associated with an application within Azure Active Directory.

    ## API Permissions

    The following API permissions are required in order to use this data source.

    When authenticated with a service principal, this data source requires one of the following application roles: `Application.Read.All` or `Directory.Read.All`

    When authenticated with a user principal, this data source does not require any additional roles.

    ## Example Usage

    *Look up by application display name*

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_service_principal(display_name="my-awesome-application")
    ```

    *Look up by application ID (client ID)*

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_service_principal(application_id="00000000-0000-0000-0000-000000000000")
    ```

    *Look up by service principal object ID*

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_service_principal(object_id="00000000-0000-0000-0000-000000000000")
    ```


    :param str application_id: The application ID (client ID) of the application associated with this service principal.
    :param str display_name: The display name of the application associated with this service principal.
    :param str object_id: The object ID of the service principal.
    """
    ...

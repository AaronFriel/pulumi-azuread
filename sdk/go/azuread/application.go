// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Application within Azure Active Directory.
//
// > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write owned by applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuread/sdk/v2/go/azuread"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := azuread.NewApplication(ctx, "example", &azuread.ApplicationArgs{
// 			AppRoles: azuread.ApplicationAppRoleArray{
// 				&azuread.ApplicationAppRoleArgs{
// 					AllowedMemberTypes: pulumi.StringArray{
// 						pulumi.String("User"),
// 						pulumi.String("Application"),
// 					},
// 					Description: pulumi.String("Admins can manage roles and perform all task actions"),
// 					DisplayName: pulumi.String("Admin"),
// 					IsEnabled:   pulumi.Bool(true),
// 					Value:       pulumi.String("Admin"),
// 				},
// 			},
// 			AvailableToOtherTenants: pulumi.Bool(false),
// 			Homepage:                pulumi.String("https://homepage"),
// 			IdentifierUris: pulumi.StringArray{
// 				pulumi.String("https://uri"),
// 			},
// 			Oauth2AllowImplicitFlow: pulumi.Bool(true),
// 			Oauth2Permissions: azuread.ApplicationOauth2PermissionArray{
// 				&azuread.ApplicationOauth2PermissionArgs{
// 					AdminConsentDescription: pulumi.String("Allow the application to access example on behalf of the signed-in user."),
// 					AdminConsentDisplayName: pulumi.String("Access example"),
// 					IsEnabled:               pulumi.Bool(true),
// 					Type:                    pulumi.String("User"),
// 					UserConsentDescription:  pulumi.String("Allow the application to access example on your behalf."),
// 					UserConsentDisplayName:  pulumi.String("Access example"),
// 					Value:                   pulumi.String("user_impersonation"),
// 				},
// 				&azuread.ApplicationOauth2PermissionArgs{
// 					AdminConsentDescription: pulumi.String("Administer the example application"),
// 					AdminConsentDisplayName: pulumi.String("Administer"),
// 					IsEnabled:               pulumi.Bool(true),
// 					Type:                    pulumi.String("Admin"),
// 					Value:                   pulumi.String("administer"),
// 				},
// 			},
// 			OptionalClaims: &azuread.ApplicationOptionalClaimsArgs{
// 				AccessTokens: azuread.ApplicationOptionalClaimsAccessTokenArray{
// 					&azuread.ApplicationOptionalClaimsAccessTokenArgs{
// 						Name: pulumi.String("myclaim"),
// 					},
// 					&azuread.ApplicationOptionalClaimsAccessTokenArgs{
// 						Name: pulumi.String("otherclaim"),
// 					},
// 				},
// 				IdTokens: azuread.ApplicationOptionalClaimsIdTokenArray{
// 					&azuread.ApplicationOptionalClaimsIdTokenArgs{
// 						AdditionalProperties: pulumi.StringArray{
// 							pulumi.String("emit_as_roles"),
// 						},
// 						Essential: pulumi.Bool(true),
// 						Name:      pulumi.String("userclaim"),
// 						Source:    pulumi.String("user"),
// 					},
// 				},
// 			},
// 			Owners: pulumi.StringArray{
// 				pulumi.String("00000004-0000-0000-c000-000000000000"),
// 			},
// 			ReplyUrls: pulumi.StringArray{
// 				pulumi.String("https://replyurl"),
// 			},
// 			RequiredResourceAccesses: azuread.ApplicationRequiredResourceAccessArray{
// 				&azuread.ApplicationRequiredResourceAccessArgs{
// 					ResourceAccesses: azuread.ApplicationRequiredResourceAccessResourceAccessArray{
// 						&azuread.ApplicationRequiredResourceAccessResourceAccessArgs{
// 							Id:   pulumi.String("..."),
// 							Type: pulumi.String("Role"),
// 						},
// 						&azuread.ApplicationRequiredResourceAccessResourceAccessArgs{
// 							Id:   pulumi.String("..."),
// 							Type: pulumi.String("Scope"),
// 						},
// 						&azuread.ApplicationRequiredResourceAccessResourceAccessArgs{
// 							Id:   pulumi.String("..."),
// 							Type: pulumi.String("Scope"),
// 						},
// 					},
// 					ResourceAppId: pulumi.String("00000003-0000-0000-c000-000000000000"),
// 				},
// 				&azuread.ApplicationRequiredResourceAccessArgs{
// 					ResourceAccesses: azuread.ApplicationRequiredResourceAccessResourceAccessArray{
// 						&azuread.ApplicationRequiredResourceAccessResourceAccessArgs{
// 							Id:   pulumi.String("..."),
// 							Type: pulumi.String("Scope"),
// 						},
// 					},
// 					ResourceAppId: pulumi.String("00000002-0000-0000-c000-000000000000"),
// 				},
// 			},
// 			Type: pulumi.String("webapp/api"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Azure Active Directory Applications can be imported using the `object id`, e.g.
//
// ```sh
//  $ pulumi import azuread:index/application:Application test 00000000-0000-0000-0000-000000000000
// ```
type Application struct {
	pulumi.CustomResourceState

	// A collection of `appRole` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
	AppRoles ApplicationAppRoleTypeArrayOutput `pulumi:"appRoles"`
	// The Application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// Is this Azure AD Application available to other tenants? Defaults to `false`.
	AvailableToOtherTenants pulumi.BoolPtrOutput `pulumi:"availableToOtherTenants"`
	// Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
	GroupMembershipClaims pulumi.StringPtrOutput `pulumi:"groupMembershipClaims"`
	// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
	Homepage pulumi.StringOutput `pulumi:"homepage"`
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris pulumi.StringArrayOutput `pulumi:"identifierUris"`
	// The URL of the logout page.
	LogoutUrl pulumi.StringPtrOutput `pulumi:"logoutUrl"`
	// The display name for the application.
	Name pulumi.StringOutput `pulumi:"name"`
	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
	Oauth2AllowImplicitFlow pulumi.BoolPtrOutput `pulumi:"oauth2AllowImplicitFlow"`
	// A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by `oauth2Permissions` blocks as documented below.
	Oauth2Permissions ApplicationOauth2PermissionArrayOutput `pulumi:"oauth2Permissions"`
	// The Application's Object ID.
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// A collection of `accessToken` or `idToken` blocks as documented below which list the optional claims configured for each token type. For more information see https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-optional-claims
	OptionalClaims ApplicationOptionalClaimsPtrOutput `pulumi:"optionalClaims"`
	// A list of Azure AD Object IDs that will be granted ownership of the application. Defaults to the Object ID of the caller creating the application. If a list is specified the caller Object ID will no longer be included unless explicitly added to the list.
	Owners pulumi.StringArrayOutput `pulumi:"owners"`
	// If `true`, will return an error when an existing Application is found with the same name. Defaults to `false`.
	PreventDuplicateNames pulumi.BoolPtrOutput `pulumi:"preventDuplicateNames"`
	// Is this Azure AD Application a public client? Defaults to `false`.
	PublicClient pulumi.BoolOutput `pulumi:"publicClient"`
	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls pulumi.StringArrayOutput `pulumi:"replyUrls"`
	// A collection of `requiredResourceAccess` blocks as documented below.
	RequiredResourceAccesses ApplicationRequiredResourceAccessArrayOutput `pulumi:"requiredResourceAccesses"`
	// Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifierUris` property can not not be set.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		args = &ApplicationArgs{}
	}

	var resource Application
	err := ctx.RegisterResource("azuread:index/application:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("azuread:index/application:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// A collection of `appRole` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
	AppRoles []ApplicationAppRoleType `pulumi:"appRoles"`
	// The Application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// Is this Azure AD Application available to other tenants? Defaults to `false`.
	AvailableToOtherTenants *bool `pulumi:"availableToOtherTenants"`
	// Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
	GroupMembershipClaims *string `pulumi:"groupMembershipClaims"`
	// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
	Homepage *string `pulumi:"homepage"`
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris []string `pulumi:"identifierUris"`
	// The URL of the logout page.
	LogoutUrl *string `pulumi:"logoutUrl"`
	// The display name for the application.
	Name *string `pulumi:"name"`
	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
	Oauth2AllowImplicitFlow *bool `pulumi:"oauth2AllowImplicitFlow"`
	// A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by `oauth2Permissions` blocks as documented below.
	Oauth2Permissions []ApplicationOauth2Permission `pulumi:"oauth2Permissions"`
	// The Application's Object ID.
	ObjectId *string `pulumi:"objectId"`
	// A collection of `accessToken` or `idToken` blocks as documented below which list the optional claims configured for each token type. For more information see https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-optional-claims
	OptionalClaims *ApplicationOptionalClaims `pulumi:"optionalClaims"`
	// A list of Azure AD Object IDs that will be granted ownership of the application. Defaults to the Object ID of the caller creating the application. If a list is specified the caller Object ID will no longer be included unless explicitly added to the list.
	Owners []string `pulumi:"owners"`
	// If `true`, will return an error when an existing Application is found with the same name. Defaults to `false`.
	PreventDuplicateNames *bool `pulumi:"preventDuplicateNames"`
	// Is this Azure AD Application a public client? Defaults to `false`.
	PublicClient *bool `pulumi:"publicClient"`
	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls []string `pulumi:"replyUrls"`
	// A collection of `requiredResourceAccess` blocks as documented below.
	RequiredResourceAccesses []ApplicationRequiredResourceAccess `pulumi:"requiredResourceAccesses"`
	// Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifierUris` property can not not be set.
	Type *string `pulumi:"type"`
}

type ApplicationState struct {
	// A collection of `appRole` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
	AppRoles ApplicationAppRoleTypeArrayInput
	// The Application ID.
	ApplicationId pulumi.StringPtrInput
	// Is this Azure AD Application available to other tenants? Defaults to `false`.
	AvailableToOtherTenants pulumi.BoolPtrInput
	// Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
	GroupMembershipClaims pulumi.StringPtrInput
	// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
	Homepage pulumi.StringPtrInput
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris pulumi.StringArrayInput
	// The URL of the logout page.
	LogoutUrl pulumi.StringPtrInput
	// The display name for the application.
	Name pulumi.StringPtrInput
	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
	Oauth2AllowImplicitFlow pulumi.BoolPtrInput
	// A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by `oauth2Permissions` blocks as documented below.
	Oauth2Permissions ApplicationOauth2PermissionArrayInput
	// The Application's Object ID.
	ObjectId pulumi.StringPtrInput
	// A collection of `accessToken` or `idToken` blocks as documented below which list the optional claims configured for each token type. For more information see https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-optional-claims
	OptionalClaims ApplicationOptionalClaimsPtrInput
	// A list of Azure AD Object IDs that will be granted ownership of the application. Defaults to the Object ID of the caller creating the application. If a list is specified the caller Object ID will no longer be included unless explicitly added to the list.
	Owners pulumi.StringArrayInput
	// If `true`, will return an error when an existing Application is found with the same name. Defaults to `false`.
	PreventDuplicateNames pulumi.BoolPtrInput
	// Is this Azure AD Application a public client? Defaults to `false`.
	PublicClient pulumi.BoolPtrInput
	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls pulumi.StringArrayInput
	// A collection of `requiredResourceAccess` blocks as documented below.
	RequiredResourceAccesses ApplicationRequiredResourceAccessArrayInput
	// Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifierUris` property can not not be set.
	Type pulumi.StringPtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// A collection of `appRole` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
	AppRoles []ApplicationAppRoleType `pulumi:"appRoles"`
	// Is this Azure AD Application available to other tenants? Defaults to `false`.
	AvailableToOtherTenants *bool `pulumi:"availableToOtherTenants"`
	// Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
	GroupMembershipClaims *string `pulumi:"groupMembershipClaims"`
	// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
	Homepage *string `pulumi:"homepage"`
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris []string `pulumi:"identifierUris"`
	// The URL of the logout page.
	LogoutUrl *string `pulumi:"logoutUrl"`
	// The display name for the application.
	Name *string `pulumi:"name"`
	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
	Oauth2AllowImplicitFlow *bool `pulumi:"oauth2AllowImplicitFlow"`
	// A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by `oauth2Permissions` blocks as documented below.
	Oauth2Permissions []ApplicationOauth2Permission `pulumi:"oauth2Permissions"`
	// A collection of `accessToken` or `idToken` blocks as documented below which list the optional claims configured for each token type. For more information see https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-optional-claims
	OptionalClaims *ApplicationOptionalClaims `pulumi:"optionalClaims"`
	// A list of Azure AD Object IDs that will be granted ownership of the application. Defaults to the Object ID of the caller creating the application. If a list is specified the caller Object ID will no longer be included unless explicitly added to the list.
	Owners []string `pulumi:"owners"`
	// If `true`, will return an error when an existing Application is found with the same name. Defaults to `false`.
	PreventDuplicateNames *bool `pulumi:"preventDuplicateNames"`
	// Is this Azure AD Application a public client? Defaults to `false`.
	PublicClient *bool `pulumi:"publicClient"`
	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls []string `pulumi:"replyUrls"`
	// A collection of `requiredResourceAccess` blocks as documented below.
	RequiredResourceAccesses []ApplicationRequiredResourceAccess `pulumi:"requiredResourceAccesses"`
	// Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifierUris` property can not not be set.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// A collection of `appRole` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
	AppRoles ApplicationAppRoleTypeArrayInput
	// Is this Azure AD Application available to other tenants? Defaults to `false`.
	AvailableToOtherTenants pulumi.BoolPtrInput
	// Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
	GroupMembershipClaims pulumi.StringPtrInput
	// The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
	Homepage pulumi.StringPtrInput
	// A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris pulumi.StringArrayInput
	// The URL of the logout page.
	LogoutUrl pulumi.StringPtrInput
	// The display name for the application.
	Name pulumi.StringPtrInput
	// Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
	Oauth2AllowImplicitFlow pulumi.BoolPtrInput
	// A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by `oauth2Permissions` blocks as documented below.
	Oauth2Permissions ApplicationOauth2PermissionArrayInput
	// A collection of `accessToken` or `idToken` blocks as documented below which list the optional claims configured for each token type. For more information see https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-optional-claims
	OptionalClaims ApplicationOptionalClaimsPtrInput
	// A list of Azure AD Object IDs that will be granted ownership of the application. Defaults to the Object ID of the caller creating the application. If a list is specified the caller Object ID will no longer be included unless explicitly added to the list.
	Owners pulumi.StringArrayInput
	// If `true`, will return an error when an existing Application is found with the same name. Defaults to `false`.
	PreventDuplicateNames pulumi.BoolPtrInput
	// Is this Azure AD Application a public client? Defaults to `false`.
	PublicClient pulumi.BoolPtrInput
	// A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
	ReplyUrls pulumi.StringArrayInput
	// A collection of `requiredResourceAccess` blocks as documented below.
	RequiredResourceAccesses ApplicationRequiredResourceAccessArrayInput
	// Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifierUris` property can not not be set.
	Type pulumi.StringPtrInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (Application) ElementType() reflect.Type {
	return reflect.TypeOf((*Application)(nil)).Elem()
}

func (i Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

type ApplicationOutput struct {
	*pulumi.OutputState
}

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationOutput)(nil)).Elem()
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApplicationOutput{})
}

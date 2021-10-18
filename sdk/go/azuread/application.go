// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Applications can be imported using their object ID, e.g.
//
// ```sh
//  $ pulumi import azuread:index/application:Application test 00000000-0000-0000-0000-000000000000
// ```
type Application struct {
	pulumi.CustomResourceState

	// An `api` block as documented below, which configures API related settings for this application.
	Api ApplicationApiPtrOutput `pulumi:"api"`
	// A mapping of app role values to app role IDs, intended to be useful when referencing app roles in other resources in your configuration.
	AppRoleIds pulumi.StringMapOutput `pulumi:"appRoleIds"`
	// A collection of `appRole` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
	AppRoles ApplicationAppRoleArrayOutput `pulumi:"appRoles"`
	// The Application ID (also called Client ID).
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// Specifies whether this application supports device authentication without a user. Defaults to `false`.
	DeviceOnlyAuthEnabled pulumi.BoolPtrOutput `pulumi:"deviceOnlyAuthEnabled"`
	// Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
	DisabledByMicrosoft pulumi.StringOutput `pulumi:"disabledByMicrosoft"`
	// The display name for the application.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Specifies whether the application is a public client. Appropriate for apps using token grant flows that don't use a redirect URI. Defaults to `false`.
	FallbackPublicClientEnabled pulumi.BoolPtrOutput `pulumi:"fallbackPublicClientEnabled"`
	// A `featureTags` block as described below. Cannot be used together with the `tags` property.
	FeatureTags ApplicationFeatureTagArrayOutput `pulumi:"featureTags"`
	// Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
	GroupMembershipClaims pulumi.StringArrayOutput `pulumi:"groupMembershipClaims"`
	// A set of user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris pulumi.StringArrayOutput `pulumi:"identifierUris"`
	// A logo image to upload for the application, as a raw base64-encoded string. The image should be in gif, jpeg or png format. Note that once an image has been uploaded, it is not possible to remove it without replacing it with another image.
	LogoImage pulumi.StringPtrOutput `pulumi:"logoImage"`
	// CDN URL to the application's logo, as uploaded with the `logoImage` property.
	LogoUrl pulumi.StringOutput `pulumi:"logoUrl"`
	// URL of the application's marketing page.
	MarketingUrl pulumi.StringPtrOutput `pulumi:"marketingUrl"`
	// A mapping of OAuth2.0 permission scope values to scope IDs, intended to be useful when referencing permission scopes in other resources in your configuration.
	Oauth2PermissionScopeIds pulumi.StringMapOutput `pulumi:"oauth2PermissionScopeIds"`
	// Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. Defaults to `false`, which specifies that only GET requests are allowed.
	Oauth2PostResponseRequired pulumi.BoolPtrOutput `pulumi:"oauth2PostResponseRequired"`
	// The application's object ID.
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// An `optionalClaims` block as documented below.
	OptionalClaims ApplicationOptionalClaimsPtrOutput `pulumi:"optionalClaims"`
	// A set of object IDs of principals that will be granted ownership of the application. Supported object types are users or service principals. By default, no owners are assigned.
	Owners pulumi.StringArrayOutput `pulumi:"owners"`
	// If `true`, will return an error if an existing application is found with the same name. Defaults to `false`.
	PreventDuplicateNames pulumi.BoolPtrOutput `pulumi:"preventDuplicateNames"`
	// URL of the application's privacy statement.
	PrivacyStatementUrl pulumi.StringPtrOutput `pulumi:"privacyStatementUrl"`
	// A `publicClient` block as documented below, which configures non-web app or non-web API application settings, for example mobile or other public clients such as an installed application running on a desktop device.
	PublicClient ApplicationPublicClientPtrOutput `pulumi:"publicClient"`
	// The verified publisher domain for the application.
	PublisherDomain pulumi.StringOutput `pulumi:"publisherDomain"`
	// A collection of `requiredResourceAccess` blocks as documented below.
	RequiredResourceAccesses ApplicationRequiredResourceAccessArrayOutput `pulumi:"requiredResourceAccesses"`
	// The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
	SignInAudience pulumi.StringPtrOutput `pulumi:"signInAudience"`
	// A `singlePageApplication` block as documented below, which configures single-page application (SPA) related settings for this application.
	SinglePageApplication ApplicationSinglePageApplicationPtrOutput `pulumi:"singlePageApplication"`
	// URL of the application's support page.
	SupportUrl pulumi.StringPtrOutput `pulumi:"supportUrl"`
	// A set of tags to apply to the application. Cannot be used together with the `featureTags` block.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
	TemplateId pulumi.StringOutput `pulumi:"templateId"`
	// URL of the application's terms of service statement.
	TermsOfServiceUrl pulumi.StringPtrOutput `pulumi:"termsOfServiceUrl"`
	// A `web` block as documented below, which configures web related settings for this application.
	Web ApplicationWebPtrOutput `pulumi:"web"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
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
	// An `api` block as documented below, which configures API related settings for this application.
	Api *ApplicationApi `pulumi:"api"`
	// A mapping of app role values to app role IDs, intended to be useful when referencing app roles in other resources in your configuration.
	AppRoleIds map[string]string `pulumi:"appRoleIds"`
	// A collection of `appRole` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
	AppRoles []ApplicationAppRole `pulumi:"appRoles"`
	// The Application ID (also called Client ID).
	ApplicationId *string `pulumi:"applicationId"`
	// Specifies whether this application supports device authentication without a user. Defaults to `false`.
	DeviceOnlyAuthEnabled *bool `pulumi:"deviceOnlyAuthEnabled"`
	// Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
	DisabledByMicrosoft *string `pulumi:"disabledByMicrosoft"`
	// The display name for the application.
	DisplayName *string `pulumi:"displayName"`
	// Specifies whether the application is a public client. Appropriate for apps using token grant flows that don't use a redirect URI. Defaults to `false`.
	FallbackPublicClientEnabled *bool `pulumi:"fallbackPublicClientEnabled"`
	// A `featureTags` block as described below. Cannot be used together with the `tags` property.
	FeatureTags []ApplicationFeatureTag `pulumi:"featureTags"`
	// Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
	GroupMembershipClaims []string `pulumi:"groupMembershipClaims"`
	// A set of user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris []string `pulumi:"identifierUris"`
	// A logo image to upload for the application, as a raw base64-encoded string. The image should be in gif, jpeg or png format. Note that once an image has been uploaded, it is not possible to remove it without replacing it with another image.
	LogoImage *string `pulumi:"logoImage"`
	// CDN URL to the application's logo, as uploaded with the `logoImage` property.
	LogoUrl *string `pulumi:"logoUrl"`
	// URL of the application's marketing page.
	MarketingUrl *string `pulumi:"marketingUrl"`
	// A mapping of OAuth2.0 permission scope values to scope IDs, intended to be useful when referencing permission scopes in other resources in your configuration.
	Oauth2PermissionScopeIds map[string]string `pulumi:"oauth2PermissionScopeIds"`
	// Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. Defaults to `false`, which specifies that only GET requests are allowed.
	Oauth2PostResponseRequired *bool `pulumi:"oauth2PostResponseRequired"`
	// The application's object ID.
	ObjectId *string `pulumi:"objectId"`
	// An `optionalClaims` block as documented below.
	OptionalClaims *ApplicationOptionalClaims `pulumi:"optionalClaims"`
	// A set of object IDs of principals that will be granted ownership of the application. Supported object types are users or service principals. By default, no owners are assigned.
	Owners []string `pulumi:"owners"`
	// If `true`, will return an error if an existing application is found with the same name. Defaults to `false`.
	PreventDuplicateNames *bool `pulumi:"preventDuplicateNames"`
	// URL of the application's privacy statement.
	PrivacyStatementUrl *string `pulumi:"privacyStatementUrl"`
	// A `publicClient` block as documented below, which configures non-web app or non-web API application settings, for example mobile or other public clients such as an installed application running on a desktop device.
	PublicClient *ApplicationPublicClient `pulumi:"publicClient"`
	// The verified publisher domain for the application.
	PublisherDomain *string `pulumi:"publisherDomain"`
	// A collection of `requiredResourceAccess` blocks as documented below.
	RequiredResourceAccesses []ApplicationRequiredResourceAccess `pulumi:"requiredResourceAccesses"`
	// The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
	SignInAudience *string `pulumi:"signInAudience"`
	// A `singlePageApplication` block as documented below, which configures single-page application (SPA) related settings for this application.
	SinglePageApplication *ApplicationSinglePageApplication `pulumi:"singlePageApplication"`
	// URL of the application's support page.
	SupportUrl *string `pulumi:"supportUrl"`
	// A set of tags to apply to the application. Cannot be used together with the `featureTags` block.
	Tags []string `pulumi:"tags"`
	// Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
	TemplateId *string `pulumi:"templateId"`
	// URL of the application's terms of service statement.
	TermsOfServiceUrl *string `pulumi:"termsOfServiceUrl"`
	// A `web` block as documented below, which configures web related settings for this application.
	Web *ApplicationWeb `pulumi:"web"`
}

type ApplicationState struct {
	// An `api` block as documented below, which configures API related settings for this application.
	Api ApplicationApiPtrInput
	// A mapping of app role values to app role IDs, intended to be useful when referencing app roles in other resources in your configuration.
	AppRoleIds pulumi.StringMapInput
	// A collection of `appRole` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
	AppRoles ApplicationAppRoleArrayInput
	// The Application ID (also called Client ID).
	ApplicationId pulumi.StringPtrInput
	// Specifies whether this application supports device authentication without a user. Defaults to `false`.
	DeviceOnlyAuthEnabled pulumi.BoolPtrInput
	// Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
	DisabledByMicrosoft pulumi.StringPtrInput
	// The display name for the application.
	DisplayName pulumi.StringPtrInput
	// Specifies whether the application is a public client. Appropriate for apps using token grant flows that don't use a redirect URI. Defaults to `false`.
	FallbackPublicClientEnabled pulumi.BoolPtrInput
	// A `featureTags` block as described below. Cannot be used together with the `tags` property.
	FeatureTags ApplicationFeatureTagArrayInput
	// Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
	GroupMembershipClaims pulumi.StringArrayInput
	// A set of user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris pulumi.StringArrayInput
	// A logo image to upload for the application, as a raw base64-encoded string. The image should be in gif, jpeg or png format. Note that once an image has been uploaded, it is not possible to remove it without replacing it with another image.
	LogoImage pulumi.StringPtrInput
	// CDN URL to the application's logo, as uploaded with the `logoImage` property.
	LogoUrl pulumi.StringPtrInput
	// URL of the application's marketing page.
	MarketingUrl pulumi.StringPtrInput
	// A mapping of OAuth2.0 permission scope values to scope IDs, intended to be useful when referencing permission scopes in other resources in your configuration.
	Oauth2PermissionScopeIds pulumi.StringMapInput
	// Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. Defaults to `false`, which specifies that only GET requests are allowed.
	Oauth2PostResponseRequired pulumi.BoolPtrInput
	// The application's object ID.
	ObjectId pulumi.StringPtrInput
	// An `optionalClaims` block as documented below.
	OptionalClaims ApplicationOptionalClaimsPtrInput
	// A set of object IDs of principals that will be granted ownership of the application. Supported object types are users or service principals. By default, no owners are assigned.
	Owners pulumi.StringArrayInput
	// If `true`, will return an error if an existing application is found with the same name. Defaults to `false`.
	PreventDuplicateNames pulumi.BoolPtrInput
	// URL of the application's privacy statement.
	PrivacyStatementUrl pulumi.StringPtrInput
	// A `publicClient` block as documented below, which configures non-web app or non-web API application settings, for example mobile or other public clients such as an installed application running on a desktop device.
	PublicClient ApplicationPublicClientPtrInput
	// The verified publisher domain for the application.
	PublisherDomain pulumi.StringPtrInput
	// A collection of `requiredResourceAccess` blocks as documented below.
	RequiredResourceAccesses ApplicationRequiredResourceAccessArrayInput
	// The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
	SignInAudience pulumi.StringPtrInput
	// A `singlePageApplication` block as documented below, which configures single-page application (SPA) related settings for this application.
	SinglePageApplication ApplicationSinglePageApplicationPtrInput
	// URL of the application's support page.
	SupportUrl pulumi.StringPtrInput
	// A set of tags to apply to the application. Cannot be used together with the `featureTags` block.
	Tags pulumi.StringArrayInput
	// Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
	TemplateId pulumi.StringPtrInput
	// URL of the application's terms of service statement.
	TermsOfServiceUrl pulumi.StringPtrInput
	// A `web` block as documented below, which configures web related settings for this application.
	Web ApplicationWebPtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// An `api` block as documented below, which configures API related settings for this application.
	Api *ApplicationApi `pulumi:"api"`
	// A collection of `appRole` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
	AppRoles []ApplicationAppRole `pulumi:"appRoles"`
	// Specifies whether this application supports device authentication without a user. Defaults to `false`.
	DeviceOnlyAuthEnabled *bool `pulumi:"deviceOnlyAuthEnabled"`
	// The display name for the application.
	DisplayName string `pulumi:"displayName"`
	// Specifies whether the application is a public client. Appropriate for apps using token grant flows that don't use a redirect URI. Defaults to `false`.
	FallbackPublicClientEnabled *bool `pulumi:"fallbackPublicClientEnabled"`
	// A `featureTags` block as described below. Cannot be used together with the `tags` property.
	FeatureTags []ApplicationFeatureTag `pulumi:"featureTags"`
	// Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
	GroupMembershipClaims []string `pulumi:"groupMembershipClaims"`
	// A set of user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris []string `pulumi:"identifierUris"`
	// A logo image to upload for the application, as a raw base64-encoded string. The image should be in gif, jpeg or png format. Note that once an image has been uploaded, it is not possible to remove it without replacing it with another image.
	LogoImage *string `pulumi:"logoImage"`
	// URL of the application's marketing page.
	MarketingUrl *string `pulumi:"marketingUrl"`
	// Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. Defaults to `false`, which specifies that only GET requests are allowed.
	Oauth2PostResponseRequired *bool `pulumi:"oauth2PostResponseRequired"`
	// An `optionalClaims` block as documented below.
	OptionalClaims *ApplicationOptionalClaims `pulumi:"optionalClaims"`
	// A set of object IDs of principals that will be granted ownership of the application. Supported object types are users or service principals. By default, no owners are assigned.
	Owners []string `pulumi:"owners"`
	// If `true`, will return an error if an existing application is found with the same name. Defaults to `false`.
	PreventDuplicateNames *bool `pulumi:"preventDuplicateNames"`
	// URL of the application's privacy statement.
	PrivacyStatementUrl *string `pulumi:"privacyStatementUrl"`
	// A `publicClient` block as documented below, which configures non-web app or non-web API application settings, for example mobile or other public clients such as an installed application running on a desktop device.
	PublicClient *ApplicationPublicClient `pulumi:"publicClient"`
	// A collection of `requiredResourceAccess` blocks as documented below.
	RequiredResourceAccesses []ApplicationRequiredResourceAccess `pulumi:"requiredResourceAccesses"`
	// The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
	SignInAudience *string `pulumi:"signInAudience"`
	// A `singlePageApplication` block as documented below, which configures single-page application (SPA) related settings for this application.
	SinglePageApplication *ApplicationSinglePageApplication `pulumi:"singlePageApplication"`
	// URL of the application's support page.
	SupportUrl *string `pulumi:"supportUrl"`
	// A set of tags to apply to the application. Cannot be used together with the `featureTags` block.
	Tags []string `pulumi:"tags"`
	// Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
	TemplateId *string `pulumi:"templateId"`
	// URL of the application's terms of service statement.
	TermsOfServiceUrl *string `pulumi:"termsOfServiceUrl"`
	// A `web` block as documented below, which configures web related settings for this application.
	Web *ApplicationWeb `pulumi:"web"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// An `api` block as documented below, which configures API related settings for this application.
	Api ApplicationApiPtrInput
	// A collection of `appRole` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
	AppRoles ApplicationAppRoleArrayInput
	// Specifies whether this application supports device authentication without a user. Defaults to `false`.
	DeviceOnlyAuthEnabled pulumi.BoolPtrInput
	// The display name for the application.
	DisplayName pulumi.StringInput
	// Specifies whether the application is a public client. Appropriate for apps using token grant flows that don't use a redirect URI. Defaults to `false`.
	FallbackPublicClientEnabled pulumi.BoolPtrInput
	// A `featureTags` block as described below. Cannot be used together with the `tags` property.
	FeatureTags ApplicationFeatureTagArrayInput
	// Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
	GroupMembershipClaims pulumi.StringArrayInput
	// A set of user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	IdentifierUris pulumi.StringArrayInput
	// A logo image to upload for the application, as a raw base64-encoded string. The image should be in gif, jpeg or png format. Note that once an image has been uploaded, it is not possible to remove it without replacing it with another image.
	LogoImage pulumi.StringPtrInput
	// URL of the application's marketing page.
	MarketingUrl pulumi.StringPtrInput
	// Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. Defaults to `false`, which specifies that only GET requests are allowed.
	Oauth2PostResponseRequired pulumi.BoolPtrInput
	// An `optionalClaims` block as documented below.
	OptionalClaims ApplicationOptionalClaimsPtrInput
	// A set of object IDs of principals that will be granted ownership of the application. Supported object types are users or service principals. By default, no owners are assigned.
	Owners pulumi.StringArrayInput
	// If `true`, will return an error if an existing application is found with the same name. Defaults to `false`.
	PreventDuplicateNames pulumi.BoolPtrInput
	// URL of the application's privacy statement.
	PrivacyStatementUrl pulumi.StringPtrInput
	// A `publicClient` block as documented below, which configures non-web app or non-web API application settings, for example mobile or other public clients such as an installed application running on a desktop device.
	PublicClient ApplicationPublicClientPtrInput
	// A collection of `requiredResourceAccess` blocks as documented below.
	RequiredResourceAccesses ApplicationRequiredResourceAccessArrayInput
	// The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
	SignInAudience pulumi.StringPtrInput
	// A `singlePageApplication` block as documented below, which configures single-page application (SPA) related settings for this application.
	SinglePageApplication ApplicationSinglePageApplicationPtrInput
	// URL of the application's support page.
	SupportUrl pulumi.StringPtrInput
	// A set of tags to apply to the application. Cannot be used together with the `featureTags` block.
	Tags pulumi.StringArrayInput
	// Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
	TemplateId pulumi.StringPtrInput
	// URL of the application's terms of service statement.
	TermsOfServiceUrl pulumi.StringPtrInput
	// A `web` block as documented below, which configures web related settings for this application.
	Web ApplicationWebPtrInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((*Application)(nil))
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

func (i *Application) ToApplicationPtrOutput() ApplicationPtrOutput {
	return i.ToApplicationPtrOutputWithContext(context.Background())
}

func (i *Application) ToApplicationPtrOutputWithContext(ctx context.Context) ApplicationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPtrOutput)
}

type ApplicationPtrInput interface {
	pulumi.Input

	ToApplicationPtrOutput() ApplicationPtrOutput
	ToApplicationPtrOutputWithContext(ctx context.Context) ApplicationPtrOutput
}

type applicationPtrType ApplicationArgs

func (*applicationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil))
}

func (i *applicationPtrType) ToApplicationPtrOutput() ApplicationPtrOutput {
	return i.ToApplicationPtrOutputWithContext(context.Background())
}

func (i *applicationPtrType) ToApplicationPtrOutputWithContext(ctx context.Context) ApplicationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationPtrOutput)
}

// ApplicationArrayInput is an input type that accepts ApplicationArray and ApplicationArrayOutput values.
// You can construct a concrete instance of `ApplicationArrayInput` via:
//
//          ApplicationArray{ ApplicationArgs{...} }
type ApplicationArrayInput interface {
	pulumi.Input

	ToApplicationArrayOutput() ApplicationArrayOutput
	ToApplicationArrayOutputWithContext(context.Context) ApplicationArrayOutput
}

type ApplicationArray []ApplicationInput

func (ApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Application)(nil)).Elem()
}

func (i ApplicationArray) ToApplicationArrayOutput() ApplicationArrayOutput {
	return i.ToApplicationArrayOutputWithContext(context.Background())
}

func (i ApplicationArray) ToApplicationArrayOutputWithContext(ctx context.Context) ApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationArrayOutput)
}

// ApplicationMapInput is an input type that accepts ApplicationMap and ApplicationMapOutput values.
// You can construct a concrete instance of `ApplicationMapInput` via:
//
//          ApplicationMap{ "key": ApplicationArgs{...} }
type ApplicationMapInput interface {
	pulumi.Input

	ToApplicationMapOutput() ApplicationMapOutput
	ToApplicationMapOutputWithContext(context.Context) ApplicationMapOutput
}

type ApplicationMap map[string]ApplicationInput

func (ApplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Application)(nil)).Elem()
}

func (i ApplicationMap) ToApplicationMapOutput() ApplicationMapOutput {
	return i.ToApplicationMapOutputWithContext(context.Background())
}

func (i ApplicationMap) ToApplicationMapOutputWithContext(ctx context.Context) ApplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationMapOutput)
}

type ApplicationOutput struct{ *pulumi.OutputState }

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Application)(nil))
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationPtrOutput() ApplicationPtrOutput {
	return o.ToApplicationPtrOutputWithContext(context.Background())
}

func (o ApplicationOutput) ToApplicationPtrOutputWithContext(ctx context.Context) ApplicationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Application) *Application {
		return &v
	}).(ApplicationPtrOutput)
}

type ApplicationPtrOutput struct{ *pulumi.OutputState }

func (ApplicationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil))
}

func (o ApplicationPtrOutput) ToApplicationPtrOutput() ApplicationPtrOutput {
	return o
}

func (o ApplicationPtrOutput) ToApplicationPtrOutputWithContext(ctx context.Context) ApplicationPtrOutput {
	return o
}

func (o ApplicationPtrOutput) Elem() ApplicationOutput {
	return o.ApplyT(func(v *Application) Application {
		if v != nil {
			return *v
		}
		var ret Application
		return ret
	}).(ApplicationOutput)
}

type ApplicationArrayOutput struct{ *pulumi.OutputState }

func (ApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Application)(nil))
}

func (o ApplicationArrayOutput) ToApplicationArrayOutput() ApplicationArrayOutput {
	return o
}

func (o ApplicationArrayOutput) ToApplicationArrayOutputWithContext(ctx context.Context) ApplicationArrayOutput {
	return o
}

func (o ApplicationArrayOutput) Index(i pulumi.IntInput) ApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Application {
		return vs[0].([]Application)[vs[1].(int)]
	}).(ApplicationOutput)
}

type ApplicationMapOutput struct{ *pulumi.OutputState }

func (ApplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Application)(nil))
}

func (o ApplicationMapOutput) ToApplicationMapOutput() ApplicationMapOutput {
	return o
}

func (o ApplicationMapOutput) ToApplicationMapOutputWithContext(ctx context.Context) ApplicationMapOutput {
	return o
}

func (o ApplicationMapOutput) MapIndex(k pulumi.StringInput) ApplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Application {
		return vs[0].(map[string]Application)[vs[1].(string)]
	}).(ApplicationOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationOutput{})
	pulumi.RegisterOutputType(ApplicationPtrOutput{})
	pulumi.RegisterOutputType(ApplicationArrayOutput{})
	pulumi.RegisterOutputType(ApplicationMapOutput{})
}

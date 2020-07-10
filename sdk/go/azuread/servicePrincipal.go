// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Service Principal associated with an Application within Azure Active Directory.
//
// > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API. Please see The Granting a Service Principal permission to manage AAD for the required steps.
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
// 		exampleApplication, err := azuread.NewApplication(ctx, "exampleApplication", &azuread.ApplicationArgs{
// 			Homepage: pulumi.String("http://homepage"),
// 			IdentifierUris: pulumi.StringArray{
// 				pulumi.String("http://uri"),
// 			},
// 			ReplyUrls: pulumi.StringArray{
// 				pulumi.String("http://replyurl"),
// 			},
// 			AvailableToOtherTenants: pulumi.Bool(false),
// 			Oauth2AllowImplicitFlow: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuread.NewServicePrincipal(ctx, "exampleServicePrincipal", &azuread.ServicePrincipalArgs{
// 			ApplicationId:             exampleApplication.ApplicationId,
// 			AppRoleAssignmentRequired: pulumi.Bool(false),
// 			Tags: pulumi.StringArray{
// 				pulumi.String("example"),
// 				pulumi.String("tags"),
// 				pulumi.String("here"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ServicePrincipal struct {
	pulumi.CustomResourceState

	// Does this Service Principal require an AppRoleAssignment to a user or group before Azure AD will issue a user or access token to the application? Defaults to `false`.
	AppRoleAssignmentRequired pulumi.BoolPtrOutput `pulumi:"appRoleAssignmentRequired"`
	// The ID of the Azure AD Application for which to create a Service Principal.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// The Display Name of the Azure Active Directory Application associated with this Service Principal.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// A collection of OAuth 2.0 permissions exposed by the associated application. Each permission is covered by a `oauth2Permission` block as documented below.
	Oauth2Permissions ServicePrincipalOauth2PermissionArrayOutput `pulumi:"oauth2Permissions"`
	// The Service Principal's Object ID.
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// A list of tags to apply to the Service Principal.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewServicePrincipal registers a new resource with the given unique name, arguments, and options.
func NewServicePrincipal(ctx *pulumi.Context,
	name string, args *ServicePrincipalArgs, opts ...pulumi.ResourceOption) (*ServicePrincipal, error) {
	if args == nil || args.ApplicationId == nil {
		return nil, errors.New("missing required argument 'ApplicationId'")
	}
	if args == nil {
		args = &ServicePrincipalArgs{}
	}
	var resource ServicePrincipal
	err := ctx.RegisterResource("azuread:index/servicePrincipal:ServicePrincipal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicePrincipal gets an existing ServicePrincipal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicePrincipal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicePrincipalState, opts ...pulumi.ResourceOption) (*ServicePrincipal, error) {
	var resource ServicePrincipal
	err := ctx.ReadResource("azuread:index/servicePrincipal:ServicePrincipal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicePrincipal resources.
type servicePrincipalState struct {
	// Does this Service Principal require an AppRoleAssignment to a user or group before Azure AD will issue a user or access token to the application? Defaults to `false`.
	AppRoleAssignmentRequired *bool `pulumi:"appRoleAssignmentRequired"`
	// The ID of the Azure AD Application for which to create a Service Principal.
	ApplicationId *string `pulumi:"applicationId"`
	// The Display Name of the Azure Active Directory Application associated with this Service Principal.
	DisplayName *string `pulumi:"displayName"`
	// A collection of OAuth 2.0 permissions exposed by the associated application. Each permission is covered by a `oauth2Permission` block as documented below.
	Oauth2Permissions []ServicePrincipalOauth2Permission `pulumi:"oauth2Permissions"`
	// The Service Principal's Object ID.
	ObjectId *string `pulumi:"objectId"`
	// A list of tags to apply to the Service Principal.
	Tags []string `pulumi:"tags"`
}

type ServicePrincipalState struct {
	// Does this Service Principal require an AppRoleAssignment to a user or group before Azure AD will issue a user or access token to the application? Defaults to `false`.
	AppRoleAssignmentRequired pulumi.BoolPtrInput
	// The ID of the Azure AD Application for which to create a Service Principal.
	ApplicationId pulumi.StringPtrInput
	// The Display Name of the Azure Active Directory Application associated with this Service Principal.
	DisplayName pulumi.StringPtrInput
	// A collection of OAuth 2.0 permissions exposed by the associated application. Each permission is covered by a `oauth2Permission` block as documented below.
	Oauth2Permissions ServicePrincipalOauth2PermissionArrayInput
	// The Service Principal's Object ID.
	ObjectId pulumi.StringPtrInput
	// A list of tags to apply to the Service Principal.
	Tags pulumi.StringArrayInput
}

func (ServicePrincipalState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePrincipalState)(nil)).Elem()
}

type servicePrincipalArgs struct {
	// Does this Service Principal require an AppRoleAssignment to a user or group before Azure AD will issue a user or access token to the application? Defaults to `false`.
	AppRoleAssignmentRequired *bool `pulumi:"appRoleAssignmentRequired"`
	// The ID of the Azure AD Application for which to create a Service Principal.
	ApplicationId string `pulumi:"applicationId"`
	// A collection of OAuth 2.0 permissions exposed by the associated application. Each permission is covered by a `oauth2Permission` block as documented below.
	Oauth2Permissions []ServicePrincipalOauth2Permission `pulumi:"oauth2Permissions"`
	// A list of tags to apply to the Service Principal.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a ServicePrincipal resource.
type ServicePrincipalArgs struct {
	// Does this Service Principal require an AppRoleAssignment to a user or group before Azure AD will issue a user or access token to the application? Defaults to `false`.
	AppRoleAssignmentRequired pulumi.BoolPtrInput
	// The ID of the Azure AD Application for which to create a Service Principal.
	ApplicationId pulumi.StringInput
	// A collection of OAuth 2.0 permissions exposed by the associated application. Each permission is covered by a `oauth2Permission` block as documented below.
	Oauth2Permissions ServicePrincipalOauth2PermissionArrayInput
	// A list of tags to apply to the Service Principal.
	Tags pulumi.StringArrayInput
}

func (ServicePrincipalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePrincipalArgs)(nil)).Elem()
}

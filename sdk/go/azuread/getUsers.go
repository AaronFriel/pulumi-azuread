// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package azuread

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Gets Object IDs or UPNs for multiple Azure Active Directory users.
// 
// > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azuread/blob/master/website/docs/d/users.html.markdown.
func GetUsers(ctx *pulumi.Context, args *GetUsersArgs, opts ...pulumi.InvokeOption) (*GetUsersResult, error) {
	var rv GetUsersResult
	err := ctx.Invoke("azuread:index/getUsers:getUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUsers.
type GetUsersArgs struct {
	// The email aliases of the Azure AD Users.
	MailNicknames []string `pulumi:"mailNicknames"`
	// The Object IDs of the Azure AD Users.
	ObjectIds []string `pulumi:"objectIds"`
	// The User Principal Names of the Azure AD Users.
	UserPrincipalNames []string `pulumi:"userPrincipalNames"`
}


// A collection of values returned by getUsers.
type GetUsersResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The email aliases of the Azure AD Users.
	MailNicknames []string `pulumi:"mailNicknames"`
	// The Object IDs of the Azure AD Users.
	ObjectIds []string `pulumi:"objectIds"`
	// The User Principal Names of the Azure AD Users.
	UserPrincipalNames []string `pulumi:"userPrincipalNames"`
}


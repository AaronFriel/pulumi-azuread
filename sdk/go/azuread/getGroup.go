// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about an Azure Active Directory group.
//
// > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.
//
// ## Example Usage
// ### By Group Display Name)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuread/sdk/v4/go/azuread"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "MyGroupName"
// 		opt1 := true
// 		_, err := azuread.LookupGroup(ctx, &GetGroupArgs{
// 			DisplayName:     &opt0,
// 			SecurityEnabled: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	var rv LookupGroupResult
	err := ctx.Invoke("azuread:index/getGroup:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroup.
type LookupGroupArgs struct {
	// The display name for the Group.
	DisplayName *string `pulumi:"displayName"`
	// Whether the group is mail-enabled.
	MailEnabled *bool `pulumi:"mailEnabled"`
	// Deprecated: This property has been renamed to `display_name` and will be removed in version 2.0 of the AzureAD provider.
	Name *string `pulumi:"name"`
	// Specifies the Object ID of the Group.
	ObjectId *string `pulumi:"objectId"`
	// Whether the group is a security group.
	SecurityEnabled *bool `pulumi:"securityEnabled"`
}

// A collection of values returned by getGroup.
type LookupGroupResult struct {
	// The optional description of the Group.
	Description string `pulumi:"description"`
	// The display name for the Group.
	DisplayName string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Whether the group is mail-enabled.
	MailEnabled bool `pulumi:"mailEnabled"`
	// The Object IDs of the Group members.
	Members []string `pulumi:"members"`
	// Deprecated: This property has been renamed to `display_name` and will be removed in version 2.0 of the AzureAD provider.
	Name     string `pulumi:"name"`
	ObjectId string `pulumi:"objectId"`
	// The Object IDs of the Group owners.
	Owners []string `pulumi:"owners"`
	// Whether the group is a security group.
	SecurityEnabled bool `pulumi:"securityEnabled"`
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about an Azure Active Directory group.
//
// ## API Permissions
//
// The following API permissions are required in order to use this data source.
//
// When authenticated with a service principal, this data source requires one of the following application roles: `Group.Read.All` or `Directory.Read.All`
//
// When authenticated with a user principal, this data source does not require any additional roles.
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
	// The display name for the group.
	DisplayName *string `pulumi:"displayName"`
	// Whether the group is mail-enabled.
	MailEnabled *bool `pulumi:"mailEnabled"`
	// Specifies the object ID of the group.
	ObjectId *string `pulumi:"objectId"`
	// Whether the group is a security group.
	SecurityEnabled *bool `pulumi:"securityEnabled"`
}

// A collection of values returned by getGroup.
type LookupGroupResult struct {
	// Indicates whether this group can be assigned to an Azure Active Directory role.
	AssignableToRole bool `pulumi:"assignableToRole"`
	// A list of behaviors for a Microsoft 365 group, such as `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details.
	Behaviors []string `pulumi:"behaviors"`
	// The optional description of the group.
	Description string `pulumi:"description"`
	// The display name for the group.
	DisplayName string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The SMTP address for the group.
	Mail string `pulumi:"mail"`
	// Whether the group is mail-enabled.
	MailEnabled bool `pulumi:"mailEnabled"`
	// The mail alias for the group, unique in the organisation.
	MailNickname string `pulumi:"mailNickname"`
	// List of object IDs of the group members.
	Members []string `pulumi:"members"`
	// The object ID of the group.
	ObjectId string `pulumi:"objectId"`
	// The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesDomainName string `pulumi:"onpremisesDomainName"`
	// The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesNetbiosName string `pulumi:"onpremisesNetbiosName"`
	// The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesSamAccountName string `pulumi:"onpremisesSamAccountName"`
	// The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
	OnpremisesSecurityIdentifier string `pulumi:"onpremisesSecurityIdentifier"`
	// Whether this group is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
	OnpremisesSyncEnabled bool `pulumi:"onpremisesSyncEnabled"`
	// List of object IDs of the group owners.
	Owners []string `pulumi:"owners"`
	// The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
	PreferredLanguage string `pulumi:"preferredLanguage"`
	// A list of provisioning options for a Microsoft 365 group, such as `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details.
	ProvisioningOptions []string `pulumi:"provisioningOptions"`
	// List of email addresses for the group that direct to the same group mailbox.
	ProxyAddresses []string `pulumi:"proxyAddresses"`
	// Whether the group is a security group.
	SecurityEnabled bool `pulumi:"securityEnabled"`
	// The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. When no theme is set, the value is `null`.
	Theme string `pulumi:"theme"`
	// A list of group types configured for the group. The only supported type is `Unified`, which specifies a Microsoft 365 group.
	Types []string `pulumi:"types"`
	// The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility.
	Visibility string `pulumi:"visibility"`
}

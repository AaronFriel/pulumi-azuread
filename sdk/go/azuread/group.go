// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Group within Azure Active Directory.
//
// > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read and write all groups` within the `Windows Azure Active Directory` API. In addition it must also have either the `Company Administrator` or `User Account Administrator` Azure Active Directory roles assigned in order to be able to delete groups. You can assign one of the required Azure Active Directory Roles with the **AzureAD PowerShell Module**, which is available for Windows PowerShell or in the Azure Cloud Shell. Please refer to [this documentation](https://docs.microsoft.com/en-us/powershell/module/azuread/add-azureaddirectoryrolemember) for more details.
//
// ## Example Usage
//
// *Basic example*
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuread/sdk/v3/go/azuread/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := azuread.NewGroup(ctx, "example", &azuread.GroupArgs{
// 			DisplayName: pulumi.String("A-AD-Group"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// *A group with members*
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuread/sdk/v3/go/azuread/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleUser, err := azuread.NewUser(ctx, "exampleUser", &azuread.UserArgs{
// 			DisplayName:       pulumi.String("J Doe"),
// 			Password:          pulumi.String("notSecure123"),
// 			UserPrincipalName: pulumi.String("jdoe@hashicorp.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuread.NewGroup(ctx, "exampleGroup", &azuread.GroupArgs{
// 			DisplayName: pulumi.String("MyGroup"),
// 			Members: pulumi.StringArray{
// 				exampleUser.ObjectId,
// 			},
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
// Azure Active Directory Groups can be imported using the `object id`, e.g.
//
// ```sh
//  $ pulumi import azuread:index/group:Group my_group 00000000-0000-0000-0000-000000000000
// ```
type Group struct {
	pulumi.CustomResourceState

	// The description for the Group.  Changing this forces a new resource to be created.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name for the Group. Changing this forces a new resource to be created.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// A set of members who should be present in this Group. Supported Object types are Users, Groups or Service Principals.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// Deprecated: This property has been renamed to `display_name` and will be removed in v2.0 of this provider.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Object ID of the Group.
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// A set of owners who own this Group. Supported Object types are Users or Service Principals.
	Owners pulumi.StringArrayOutput `pulumi:"owners"`
	// If `true`, will return an error when an existing Group is found with the same name. Defaults to `false`.
	PreventDuplicateNames pulumi.BoolPtrOutput `pulumi:"preventDuplicateNames"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		args = &GroupArgs{}
	}

	var resource Group
	err := ctx.RegisterResource("azuread:index/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("azuread:index/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// The description for the Group.  Changing this forces a new resource to be created.
	Description *string `pulumi:"description"`
	// The display name for the Group. Changing this forces a new resource to be created.
	DisplayName *string `pulumi:"displayName"`
	// A set of members who should be present in this Group. Supported Object types are Users, Groups or Service Principals.
	Members []string `pulumi:"members"`
	// Deprecated: This property has been renamed to `display_name` and will be removed in v2.0 of this provider.
	Name *string `pulumi:"name"`
	// The Object ID of the Group.
	ObjectId *string `pulumi:"objectId"`
	// A set of owners who own this Group. Supported Object types are Users or Service Principals.
	Owners []string `pulumi:"owners"`
	// If `true`, will return an error when an existing Group is found with the same name. Defaults to `false`.
	PreventDuplicateNames *bool `pulumi:"preventDuplicateNames"`
}

type GroupState struct {
	// The description for the Group.  Changing this forces a new resource to be created.
	Description pulumi.StringPtrInput
	// The display name for the Group. Changing this forces a new resource to be created.
	DisplayName pulumi.StringPtrInput
	// A set of members who should be present in this Group. Supported Object types are Users, Groups or Service Principals.
	Members pulumi.StringArrayInput
	// Deprecated: This property has been renamed to `display_name` and will be removed in v2.0 of this provider.
	Name pulumi.StringPtrInput
	// The Object ID of the Group.
	ObjectId pulumi.StringPtrInput
	// A set of owners who own this Group. Supported Object types are Users or Service Principals.
	Owners pulumi.StringArrayInput
	// If `true`, will return an error when an existing Group is found with the same name. Defaults to `false`.
	PreventDuplicateNames pulumi.BoolPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// The description for the Group.  Changing this forces a new resource to be created.
	Description *string `pulumi:"description"`
	// The display name for the Group. Changing this forces a new resource to be created.
	DisplayName *string `pulumi:"displayName"`
	// A set of members who should be present in this Group. Supported Object types are Users, Groups or Service Principals.
	Members []string `pulumi:"members"`
	// Deprecated: This property has been renamed to `display_name` and will be removed in v2.0 of this provider.
	Name *string `pulumi:"name"`
	// A set of owners who own this Group. Supported Object types are Users or Service Principals.
	Owners []string `pulumi:"owners"`
	// If `true`, will return an error when an existing Group is found with the same name. Defaults to `false`.
	PreventDuplicateNames *bool `pulumi:"preventDuplicateNames"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// The description for the Group.  Changing this forces a new resource to be created.
	Description pulumi.StringPtrInput
	// The display name for the Group. Changing this forces a new resource to be created.
	DisplayName pulumi.StringPtrInput
	// A set of members who should be present in this Group. Supported Object types are Users, Groups or Service Principals.
	Members pulumi.StringArrayInput
	// Deprecated: This property has been renamed to `display_name` and will be removed in v2.0 of this provider.
	Name pulumi.StringPtrInput
	// A set of owners who own this Group. Supported Object types are Users or Service Principals.
	Owners pulumi.StringArrayInput
	// If `true`, will return an error when an existing Group is found with the same name. Defaults to `false`.
	PreventDuplicateNames pulumi.BoolPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((*Group)(nil))
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

type GroupOutput struct {
	*pulumi.OutputState
}

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Group)(nil))
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GroupOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package platform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource for looking up secert file type secret.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-harness/sdk/go/harness/Platform"
//	"github.com/pulumi/pulumi-harness/sdk/go/harness/Platform"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Platform.GetSecretFile(ctx, &platform.GetSecretFileArgs{
//				Identifier: pulumi.StringRef("identifier"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupSecretFile(ctx *pulumi.Context, args *LookupSecretFileArgs, opts ...pulumi.InvokeOption) (*LookupSecretFileResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSecretFileResult
	err := ctx.Invoke("harness:Platform/getSecretFile:getSecretFile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecretFile.
type LookupSecretFileArgs struct {
	// Unique identifier of the resource.
	Identifier *string `pulumi:"identifier"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId *string `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getSecretFile.
type LookupSecretFileResult struct {
	// Description of the resource.
	Description string `pulumi:"description"`
	// Path of the file containing secret value
	FilePath string `pulumi:"filePath"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Unique identifier of the resource.
	Identifier *string `pulumi:"identifier"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId *string `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
	// Identifier of the Secret Manager used to manage the secret.
	SecretManagerIdentifier string `pulumi:"secretManagerIdentifier"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags []string `pulumi:"tags"`
}

func LookupSecretFileOutput(ctx *pulumi.Context, args LookupSecretFileOutputArgs, opts ...pulumi.InvokeOption) LookupSecretFileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecretFileResult, error) {
			args := v.(LookupSecretFileArgs)
			r, err := LookupSecretFile(ctx, &args, opts...)
			var s LookupSecretFileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecretFileResultOutput)
}

// A collection of arguments for invoking getSecretFile.
type LookupSecretFileOutputArgs struct {
	// Unique identifier of the resource.
	Identifier pulumi.StringPtrInput `pulumi:"identifier"`
	// Name of the resource.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (LookupSecretFileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretFileArgs)(nil)).Elem()
}

// A collection of values returned by getSecretFile.
type LookupSecretFileResultOutput struct{ *pulumi.OutputState }

func (LookupSecretFileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretFileResult)(nil)).Elem()
}

func (o LookupSecretFileResultOutput) ToLookupSecretFileResultOutput() LookupSecretFileResultOutput {
	return o
}

func (o LookupSecretFileResultOutput) ToLookupSecretFileResultOutputWithContext(ctx context.Context) LookupSecretFileResultOutput {
	return o
}

// Description of the resource.
func (o LookupSecretFileResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretFileResult) string { return v.Description }).(pulumi.StringOutput)
}

// Path of the file containing secret value
func (o LookupSecretFileResultOutput) FilePath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretFileResult) string { return v.FilePath }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSecretFileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretFileResult) string { return v.Id }).(pulumi.StringOutput)
}

// Unique identifier of the resource.
func (o LookupSecretFileResultOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretFileResult) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o LookupSecretFileResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretFileResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Unique identifier of the organization.
func (o LookupSecretFileResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretFileResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Unique identifier of the project.
func (o LookupSecretFileResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretFileResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// Identifier of the Secret Manager used to manage the secret.
func (o LookupSecretFileResultOutput) SecretManagerIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretFileResult) string { return v.SecretManagerIdentifier }).(pulumi.StringOutput)
}

// Tags to associate with the resource. Tags should be in the form `name:value`.
func (o LookupSecretFileResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSecretFileResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecretFileResultOutput{})
}

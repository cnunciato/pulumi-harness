// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package platformconnector

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource for looking up a Sumologic connector.
func LookupSumologic(ctx *pulumi.Context, args *LookupSumologicArgs, opts ...pulumi.InvokeOption) (*LookupSumologicResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSumologicResult
	err := ctx.Invoke("harness:PlatformConnector/getSumologic:getSumologic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSumologic.
type LookupSumologicArgs struct {
	// Unique identifier of the resource.
	Identifier *string `pulumi:"identifier"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId *string `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getSumologic.
type LookupSumologicResult struct {
	// Reference to the Harness secret containing the access id.
	AccessIdRef string `pulumi:"accessIdRef"`
	// Reference to the Harness secret containing the access key.
	AccessKeyRef string `pulumi:"accessKeyRef"`
	// Connect using only the delegates which have these tags.
	DelegateSelectors []string `pulumi:"delegateSelectors"`
	// Description of the resource.
	Description string `pulumi:"description"`
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
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags []string `pulumi:"tags"`
	// Url of the SumoLogic server.
	Url string `pulumi:"url"`
}

func LookupSumologicOutput(ctx *pulumi.Context, args LookupSumologicOutputArgs, opts ...pulumi.InvokeOption) LookupSumologicResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSumologicResult, error) {
			args := v.(LookupSumologicArgs)
			r, err := LookupSumologic(ctx, &args, opts...)
			var s LookupSumologicResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSumologicResultOutput)
}

// A collection of arguments for invoking getSumologic.
type LookupSumologicOutputArgs struct {
	// Unique identifier of the resource.
	Identifier pulumi.StringPtrInput `pulumi:"identifier"`
	// Name of the resource.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (LookupSumologicOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSumologicArgs)(nil)).Elem()
}

// A collection of values returned by getSumologic.
type LookupSumologicResultOutput struct{ *pulumi.OutputState }

func (LookupSumologicResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSumologicResult)(nil)).Elem()
}

func (o LookupSumologicResultOutput) ToLookupSumologicResultOutput() LookupSumologicResultOutput {
	return o
}

func (o LookupSumologicResultOutput) ToLookupSumologicResultOutputWithContext(ctx context.Context) LookupSumologicResultOutput {
	return o
}

// Reference to the Harness secret containing the access id.
func (o LookupSumologicResultOutput) AccessIdRef() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSumologicResult) string { return v.AccessIdRef }).(pulumi.StringOutput)
}

// Reference to the Harness secret containing the access key.
func (o LookupSumologicResultOutput) AccessKeyRef() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSumologicResult) string { return v.AccessKeyRef }).(pulumi.StringOutput)
}

// Connect using only the delegates which have these tags.
func (o LookupSumologicResultOutput) DelegateSelectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSumologicResult) []string { return v.DelegateSelectors }).(pulumi.StringArrayOutput)
}

// Description of the resource.
func (o LookupSumologicResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSumologicResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSumologicResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSumologicResult) string { return v.Id }).(pulumi.StringOutput)
}

// Unique identifier of the resource.
func (o LookupSumologicResultOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSumologicResult) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o LookupSumologicResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSumologicResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Unique identifier of the organization.
func (o LookupSumologicResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSumologicResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Unique identifier of the project.
func (o LookupSumologicResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSumologicResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// Tags to associate with the resource. Tags should be in the form `name:value`.
func (o LookupSumologicResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSumologicResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// Url of the SumoLogic server.
func (o LookupSumologicResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSumologicResult) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSumologicResultOutput{})
}

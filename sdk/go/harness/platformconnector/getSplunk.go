// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package platformconnector

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource for looking up a Splunk connector.
func LookupSplunk(ctx *pulumi.Context, args *LookupSplunkArgs, opts ...pulumi.InvokeOption) (*LookupSplunkResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSplunkResult
	err := ctx.Invoke("harness:PlatformConnector/getSplunk:getSplunk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSplunk.
type LookupSplunkArgs struct {
	// Unique identifier of the resource.
	Identifier *string `pulumi:"identifier"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId *string `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getSplunk.
type LookupSplunkResult struct {
	// Splunk account id.
	AccountId string `pulumi:"accountId"`
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
	// The reference to the Harness secret containing the Splunk password.
	PasswordRef string `pulumi:"passwordRef"`
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags []string `pulumi:"tags"`
	// Url of the Splunk server.
	Url string `pulumi:"url"`
	// The username used for connecting to Splunk.
	Username string `pulumi:"username"`
}

func LookupSplunkOutput(ctx *pulumi.Context, args LookupSplunkOutputArgs, opts ...pulumi.InvokeOption) LookupSplunkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSplunkResult, error) {
			args := v.(LookupSplunkArgs)
			r, err := LookupSplunk(ctx, &args, opts...)
			var s LookupSplunkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSplunkResultOutput)
}

// A collection of arguments for invoking getSplunk.
type LookupSplunkOutputArgs struct {
	// Unique identifier of the resource.
	Identifier pulumi.StringPtrInput `pulumi:"identifier"`
	// Name of the resource.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (LookupSplunkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSplunkArgs)(nil)).Elem()
}

// A collection of values returned by getSplunk.
type LookupSplunkResultOutput struct{ *pulumi.OutputState }

func (LookupSplunkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSplunkResult)(nil)).Elem()
}

func (o LookupSplunkResultOutput) ToLookupSplunkResultOutput() LookupSplunkResultOutput {
	return o
}

func (o LookupSplunkResultOutput) ToLookupSplunkResultOutputWithContext(ctx context.Context) LookupSplunkResultOutput {
	return o
}

// Splunk account id.
func (o LookupSplunkResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSplunkResult) string { return v.AccountId }).(pulumi.StringOutput)
}

// Connect using only the delegates which have these tags.
func (o LookupSplunkResultOutput) DelegateSelectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSplunkResult) []string { return v.DelegateSelectors }).(pulumi.StringArrayOutput)
}

// Description of the resource.
func (o LookupSplunkResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSplunkResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSplunkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSplunkResult) string { return v.Id }).(pulumi.StringOutput)
}

// Unique identifier of the resource.
func (o LookupSplunkResultOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSplunkResult) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o LookupSplunkResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSplunkResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Unique identifier of the organization.
func (o LookupSplunkResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSplunkResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// The reference to the Harness secret containing the Splunk password.
func (o LookupSplunkResultOutput) PasswordRef() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSplunkResult) string { return v.PasswordRef }).(pulumi.StringOutput)
}

// Unique identifier of the project.
func (o LookupSplunkResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSplunkResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// Tags to associate with the resource. Tags should be in the form `name:value`.
func (o LookupSplunkResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSplunkResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// Url of the Splunk server.
func (o LookupSplunkResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSplunkResult) string { return v.Url }).(pulumi.StringOutput)
}

// The username used for connecting to Splunk.
func (o LookupSplunkResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSplunkResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSplunkResultOutput{})
}

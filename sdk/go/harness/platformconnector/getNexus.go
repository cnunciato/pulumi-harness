// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package platformconnector

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource for looking up a Nexus connector.
func LookupNexus(ctx *pulumi.Context, args *LookupNexusArgs, opts ...pulumi.InvokeOption) (*LookupNexusResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupNexusResult
	err := ctx.Invoke("harness:PlatformConnector/getNexus:getNexus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNexus.
type LookupNexusArgs struct {
	// Unique identifier of the resource.
	Identifier *string `pulumi:"identifier"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId *string `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getNexus.
type LookupNexusResult struct {
	// Credentials to use for authentication.
	Credentials []GetNexusCredential `pulumi:"credentials"`
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
	// URL of the Nexus server.
	Url string `pulumi:"url"`
	// Version of the Nexus server. Valid values are 2.x, 3.x
	Version string `pulumi:"version"`
}

func LookupNexusOutput(ctx *pulumi.Context, args LookupNexusOutputArgs, opts ...pulumi.InvokeOption) LookupNexusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNexusResult, error) {
			args := v.(LookupNexusArgs)
			r, err := LookupNexus(ctx, &args, opts...)
			var s LookupNexusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNexusResultOutput)
}

// A collection of arguments for invoking getNexus.
type LookupNexusOutputArgs struct {
	// Unique identifier of the resource.
	Identifier pulumi.StringPtrInput `pulumi:"identifier"`
	// Name of the resource.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (LookupNexusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNexusArgs)(nil)).Elem()
}

// A collection of values returned by getNexus.
type LookupNexusResultOutput struct{ *pulumi.OutputState }

func (LookupNexusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNexusResult)(nil)).Elem()
}

func (o LookupNexusResultOutput) ToLookupNexusResultOutput() LookupNexusResultOutput {
	return o
}

func (o LookupNexusResultOutput) ToLookupNexusResultOutputWithContext(ctx context.Context) LookupNexusResultOutput {
	return o
}

// Credentials to use for authentication.
func (o LookupNexusResultOutput) Credentials() GetNexusCredentialArrayOutput {
	return o.ApplyT(func(v LookupNexusResult) []GetNexusCredential { return v.Credentials }).(GetNexusCredentialArrayOutput)
}

// Connect using only the delegates which have these tags.
func (o LookupNexusResultOutput) DelegateSelectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNexusResult) []string { return v.DelegateSelectors }).(pulumi.StringArrayOutput)
}

// Description of the resource.
func (o LookupNexusResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNexusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusResult) string { return v.Id }).(pulumi.StringOutput)
}

// Unique identifier of the resource.
func (o LookupNexusResultOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNexusResult) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o LookupNexusResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNexusResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Unique identifier of the organization.
func (o LookupNexusResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNexusResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Unique identifier of the project.
func (o LookupNexusResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNexusResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// Tags to associate with the resource. Tags should be in the form `name:value`.
func (o LookupNexusResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNexusResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// URL of the Nexus server.
func (o LookupNexusResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusResult) string { return v.Url }).(pulumi.StringOutput)
}

// Version of the Nexus server. Valid values are 2.x, 3.x
func (o LookupNexusResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNexusResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNexusResultOutput{})
}

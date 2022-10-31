// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package platform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetGitopsRepository(ctx *pulumi.Context, args *GetGitopsRepositoryArgs, opts ...pulumi.InvokeOption) (*GetGitopsRepositoryResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetGitopsRepositoryResult
	err := ctx.Invoke("harness:platform/getGitopsRepository:getGitopsRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGitopsRepository.
type GetGitopsRepositoryArgs struct {
	AccountId         string                          `pulumi:"accountId"`
	AgentId           *string                         `pulumi:"agentId"`
	CredsOnly         *bool                           `pulumi:"credsOnly"`
	Identifier        string                          `pulumi:"identifier"`
	OrgId             *string                         `pulumi:"orgId"`
	ProjectId         string                          `pulumi:"projectId"`
	QueryForceRefresh *bool                           `pulumi:"queryForceRefresh"`
	QueryProject      *string                         `pulumi:"queryProject"`
	QueryRepo         *string                         `pulumi:"queryRepo"`
	Repos             []GetGitopsRepositoryRepo       `pulumi:"repos"`
	UpdateMasks       []GetGitopsRepositoryUpdateMask `pulumi:"updateMasks"`
	Upsert            *bool                           `pulumi:"upsert"`
}

// A collection of values returned by getGitopsRepository.
type GetGitopsRepositoryResult struct {
	AccountId string  `pulumi:"accountId"`
	AgentId   *string `pulumi:"agentId"`
	CredsOnly *bool   `pulumi:"credsOnly"`
	// The provider-assigned unique ID for this managed resource.
	Id                string                          `pulumi:"id"`
	Identifier        string                          `pulumi:"identifier"`
	OrgId             *string                         `pulumi:"orgId"`
	ProjectId         string                          `pulumi:"projectId"`
	QueryForceRefresh *bool                           `pulumi:"queryForceRefresh"`
	QueryProject      *string                         `pulumi:"queryProject"`
	QueryRepo         *string                         `pulumi:"queryRepo"`
	Repos             []GetGitopsRepositoryRepo       `pulumi:"repos"`
	UpdateMasks       []GetGitopsRepositoryUpdateMask `pulumi:"updateMasks"`
	Upsert            *bool                           `pulumi:"upsert"`
}

func GetGitopsRepositoryOutput(ctx *pulumi.Context, args GetGitopsRepositoryOutputArgs, opts ...pulumi.InvokeOption) GetGitopsRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGitopsRepositoryResult, error) {
			args := v.(GetGitopsRepositoryArgs)
			r, err := GetGitopsRepository(ctx, &args, opts...)
			var s GetGitopsRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGitopsRepositoryResultOutput)
}

// A collection of arguments for invoking getGitopsRepository.
type GetGitopsRepositoryOutputArgs struct {
	AccountId         pulumi.StringInput                      `pulumi:"accountId"`
	AgentId           pulumi.StringPtrInput                   `pulumi:"agentId"`
	CredsOnly         pulumi.BoolPtrInput                     `pulumi:"credsOnly"`
	Identifier        pulumi.StringInput                      `pulumi:"identifier"`
	OrgId             pulumi.StringPtrInput                   `pulumi:"orgId"`
	ProjectId         pulumi.StringInput                      `pulumi:"projectId"`
	QueryForceRefresh pulumi.BoolPtrInput                     `pulumi:"queryForceRefresh"`
	QueryProject      pulumi.StringPtrInput                   `pulumi:"queryProject"`
	QueryRepo         pulumi.StringPtrInput                   `pulumi:"queryRepo"`
	Repos             GetGitopsRepositoryRepoArrayInput       `pulumi:"repos"`
	UpdateMasks       GetGitopsRepositoryUpdateMaskArrayInput `pulumi:"updateMasks"`
	Upsert            pulumi.BoolPtrInput                     `pulumi:"upsert"`
}

func (GetGitopsRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGitopsRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getGitopsRepository.
type GetGitopsRepositoryResultOutput struct{ *pulumi.OutputState }

func (GetGitopsRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGitopsRepositoryResult)(nil)).Elem()
}

func (o GetGitopsRepositoryResultOutput) ToGetGitopsRepositoryResultOutput() GetGitopsRepositoryResultOutput {
	return o
}

func (o GetGitopsRepositoryResultOutput) ToGetGitopsRepositoryResultOutputWithContext(ctx context.Context) GetGitopsRepositoryResultOutput {
	return o
}

func (o GetGitopsRepositoryResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v GetGitopsRepositoryResult) string { return v.AccountId }).(pulumi.StringOutput)
}

func (o GetGitopsRepositoryResultOutput) AgentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGitopsRepositoryResult) *string { return v.AgentId }).(pulumi.StringPtrOutput)
}

func (o GetGitopsRepositoryResultOutput) CredsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetGitopsRepositoryResult) *bool { return v.CredsOnly }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGitopsRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGitopsRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetGitopsRepositoryResultOutput) Identifier() pulumi.StringOutput {
	return o.ApplyT(func(v GetGitopsRepositoryResult) string { return v.Identifier }).(pulumi.StringOutput)
}

func (o GetGitopsRepositoryResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGitopsRepositoryResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

func (o GetGitopsRepositoryResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetGitopsRepositoryResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GetGitopsRepositoryResultOutput) QueryForceRefresh() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetGitopsRepositoryResult) *bool { return v.QueryForceRefresh }).(pulumi.BoolPtrOutput)
}

func (o GetGitopsRepositoryResultOutput) QueryProject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGitopsRepositoryResult) *string { return v.QueryProject }).(pulumi.StringPtrOutput)
}

func (o GetGitopsRepositoryResultOutput) QueryRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGitopsRepositoryResult) *string { return v.QueryRepo }).(pulumi.StringPtrOutput)
}

func (o GetGitopsRepositoryResultOutput) Repos() GetGitopsRepositoryRepoArrayOutput {
	return o.ApplyT(func(v GetGitopsRepositoryResult) []GetGitopsRepositoryRepo { return v.Repos }).(GetGitopsRepositoryRepoArrayOutput)
}

func (o GetGitopsRepositoryResultOutput) UpdateMasks() GetGitopsRepositoryUpdateMaskArrayOutput {
	return o.ApplyT(func(v GetGitopsRepositoryResult) []GetGitopsRepositoryUpdateMask { return v.UpdateMasks }).(GetGitopsRepositoryUpdateMaskArrayOutput)
}

func (o GetGitopsRepositoryResultOutput) Upsert() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetGitopsRepositoryResult) *bool { return v.Upsert }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGitopsRepositoryResultOutput{})
}
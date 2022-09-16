// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprovider

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for creating a GCP cloud provider. This resource uses the config-as-code API's. When updating the `name` or `path` of this resource you should typically also set the `createBeforeDestroy = true` lifecycle setting.
type Gcp struct {
	pulumi.CustomResourceState

	// Delegate selectors to use for this provider.
	DelegateSelectors pulumi.StringArrayOutput `pulumi:"delegateSelectors"`
	// The name of the cloud provider.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the secret containing the GCP credentials
	SecretFileId pulumi.StringPtrOutput `pulumi:"secretFileId"`
	// Skip validation of GCP configuration.
	SkipValidation pulumi.BoolPtrOutput `pulumi:"skipValidation"`
	// This block is used for scoping the resource to a specific set of applications or environments.
	UsageScopes GcpUsageScopeArrayOutput `pulumi:"usageScopes"`
}

// NewGcp registers a new resource with the given unique name, arguments, and options.
func NewGcp(ctx *pulumi.Context,
	name string, args *GcpArgs, opts ...pulumi.ResourceOption) (*Gcp, error) {
	if args == nil {
		args = &GcpArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Gcp
	err := ctx.RegisterResource("harness:Cloudprovider/gcp:Gcp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGcp gets an existing Gcp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGcp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GcpState, opts ...pulumi.ResourceOption) (*Gcp, error) {
	var resource Gcp
	err := ctx.ReadResource("harness:Cloudprovider/gcp:Gcp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Gcp resources.
type gcpState struct {
	// Delegate selectors to use for this provider.
	DelegateSelectors []string `pulumi:"delegateSelectors"`
	// The name of the cloud provider.
	Name *string `pulumi:"name"`
	// The id of the secret containing the GCP credentials
	SecretFileId *string `pulumi:"secretFileId"`
	// Skip validation of GCP configuration.
	SkipValidation *bool `pulumi:"skipValidation"`
	// This block is used for scoping the resource to a specific set of applications or environments.
	UsageScopes []GcpUsageScope `pulumi:"usageScopes"`
}

type GcpState struct {
	// Delegate selectors to use for this provider.
	DelegateSelectors pulumi.StringArrayInput
	// The name of the cloud provider.
	Name pulumi.StringPtrInput
	// The id of the secret containing the GCP credentials
	SecretFileId pulumi.StringPtrInput
	// Skip validation of GCP configuration.
	SkipValidation pulumi.BoolPtrInput
	// This block is used for scoping the resource to a specific set of applications or environments.
	UsageScopes GcpUsageScopeArrayInput
}

func (GcpState) ElementType() reflect.Type {
	return reflect.TypeOf((*gcpState)(nil)).Elem()
}

type gcpArgs struct {
	// Delegate selectors to use for this provider.
	DelegateSelectors []string `pulumi:"delegateSelectors"`
	// The name of the cloud provider.
	Name *string `pulumi:"name"`
	// The id of the secret containing the GCP credentials
	SecretFileId *string `pulumi:"secretFileId"`
	// Skip validation of GCP configuration.
	SkipValidation *bool `pulumi:"skipValidation"`
	// This block is used for scoping the resource to a specific set of applications or environments.
	UsageScopes []GcpUsageScope `pulumi:"usageScopes"`
}

// The set of arguments for constructing a Gcp resource.
type GcpArgs struct {
	// Delegate selectors to use for this provider.
	DelegateSelectors pulumi.StringArrayInput
	// The name of the cloud provider.
	Name pulumi.StringPtrInput
	// The id of the secret containing the GCP credentials
	SecretFileId pulumi.StringPtrInput
	// Skip validation of GCP configuration.
	SkipValidation pulumi.BoolPtrInput
	// This block is used for scoping the resource to a specific set of applications or environments.
	UsageScopes GcpUsageScopeArrayInput
}

func (GcpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gcpArgs)(nil)).Elem()
}

type GcpInput interface {
	pulumi.Input

	ToGcpOutput() GcpOutput
	ToGcpOutputWithContext(ctx context.Context) GcpOutput
}

func (*Gcp) ElementType() reflect.Type {
	return reflect.TypeOf((**Gcp)(nil)).Elem()
}

func (i *Gcp) ToGcpOutput() GcpOutput {
	return i.ToGcpOutputWithContext(context.Background())
}

func (i *Gcp) ToGcpOutputWithContext(ctx context.Context) GcpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcpOutput)
}

// GcpArrayInput is an input type that accepts GcpArray and GcpArrayOutput values.
// You can construct a concrete instance of `GcpArrayInput` via:
//
//	GcpArray{ GcpArgs{...} }
type GcpArrayInput interface {
	pulumi.Input

	ToGcpArrayOutput() GcpArrayOutput
	ToGcpArrayOutputWithContext(context.Context) GcpArrayOutput
}

type GcpArray []GcpInput

func (GcpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Gcp)(nil)).Elem()
}

func (i GcpArray) ToGcpArrayOutput() GcpArrayOutput {
	return i.ToGcpArrayOutputWithContext(context.Background())
}

func (i GcpArray) ToGcpArrayOutputWithContext(ctx context.Context) GcpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcpArrayOutput)
}

// GcpMapInput is an input type that accepts GcpMap and GcpMapOutput values.
// You can construct a concrete instance of `GcpMapInput` via:
//
//	GcpMap{ "key": GcpArgs{...} }
type GcpMapInput interface {
	pulumi.Input

	ToGcpMapOutput() GcpMapOutput
	ToGcpMapOutputWithContext(context.Context) GcpMapOutput
}

type GcpMap map[string]GcpInput

func (GcpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Gcp)(nil)).Elem()
}

func (i GcpMap) ToGcpMapOutput() GcpMapOutput {
	return i.ToGcpMapOutputWithContext(context.Background())
}

func (i GcpMap) ToGcpMapOutputWithContext(ctx context.Context) GcpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcpMapOutput)
}

type GcpOutput struct{ *pulumi.OutputState }

func (GcpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Gcp)(nil)).Elem()
}

func (o GcpOutput) ToGcpOutput() GcpOutput {
	return o
}

func (o GcpOutput) ToGcpOutputWithContext(ctx context.Context) GcpOutput {
	return o
}

// Delegate selectors to use for this provider.
func (o GcpOutput) DelegateSelectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Gcp) pulumi.StringArrayOutput { return v.DelegateSelectors }).(pulumi.StringArrayOutput)
}

// The name of the cloud provider.
func (o GcpOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Gcp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The id of the secret containing the GCP credentials
func (o GcpOutput) SecretFileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Gcp) pulumi.StringPtrOutput { return v.SecretFileId }).(pulumi.StringPtrOutput)
}

// Skip validation of GCP configuration.
func (o GcpOutput) SkipValidation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Gcp) pulumi.BoolPtrOutput { return v.SkipValidation }).(pulumi.BoolPtrOutput)
}

// This block is used for scoping the resource to a specific set of applications or environments.
func (o GcpOutput) UsageScopes() GcpUsageScopeArrayOutput {
	return o.ApplyT(func(v *Gcp) GcpUsageScopeArrayOutput { return v.UsageScopes }).(GcpUsageScopeArrayOutput)
}

type GcpArrayOutput struct{ *pulumi.OutputState }

func (GcpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Gcp)(nil)).Elem()
}

func (o GcpArrayOutput) ToGcpArrayOutput() GcpArrayOutput {
	return o
}

func (o GcpArrayOutput) ToGcpArrayOutputWithContext(ctx context.Context) GcpArrayOutput {
	return o
}

func (o GcpArrayOutput) Index(i pulumi.IntInput) GcpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Gcp {
		return vs[0].([]*Gcp)[vs[1].(int)]
	}).(GcpOutput)
}

type GcpMapOutput struct{ *pulumi.OutputState }

func (GcpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Gcp)(nil)).Elem()
}

func (o GcpMapOutput) ToGcpMapOutput() GcpMapOutput {
	return o
}

func (o GcpMapOutput) ToGcpMapOutputWithContext(ctx context.Context) GcpMapOutput {
	return o
}

func (o GcpMapOutput) MapIndex(k pulumi.StringInput) GcpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Gcp {
		return vs[0].(map[string]*Gcp)[vs[1].(string)]
	}).(GcpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GcpInput)(nil)).Elem(), &Gcp{})
	pulumi.RegisterInputType(reflect.TypeOf((*GcpArrayInput)(nil)).Elem(), GcpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GcpMapInput)(nil)).Elem(), GcpMap{})
	pulumi.RegisterOutputType(GcpOutput{})
	pulumi.RegisterOutputType(GcpArrayOutput{})
	pulumi.RegisterOutputType(GcpMapOutput{})
}

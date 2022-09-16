// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package platformconnector

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for creating a Splunk connector.
type Splunk struct {
	pulumi.CustomResourceState

	// Splunk account id.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// Connect using only the delegates which have these tags.
	DelegateSelectors pulumi.StringArrayOutput `pulumi:"delegateSelectors"`
	// Description of the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Unique identifier of the resource.
	Identifier pulumi.StringOutput `pulumi:"identifier"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// The reference to the Harness secret containing the Splunk password.
	PasswordRef pulumi.StringOutput `pulumi:"passwordRef"`
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Url of the Splunk server.
	Url pulumi.StringOutput `pulumi:"url"`
	// The username used for connecting to Splunk.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewSplunk registers a new resource with the given unique name, arguments, and options.
func NewSplunk(ctx *pulumi.Context,
	name string, args *SplunkArgs, opts ...pulumi.ResourceOption) (*Splunk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	if args.Identifier == nil {
		return nil, errors.New("invalid value for required argument 'Identifier'")
	}
	if args.PasswordRef == nil {
		return nil, errors.New("invalid value for required argument 'PasswordRef'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Splunk
	err := ctx.RegisterResource("harness:PlatformConnector/splunk:Splunk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSplunk gets an existing Splunk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSplunk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SplunkState, opts ...pulumi.ResourceOption) (*Splunk, error) {
	var resource Splunk
	err := ctx.ReadResource("harness:PlatformConnector/splunk:Splunk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Splunk resources.
type splunkState struct {
	// Splunk account id.
	AccountId *string `pulumi:"accountId"`
	// Connect using only the delegates which have these tags.
	DelegateSelectors []string `pulumi:"delegateSelectors"`
	// Description of the resource.
	Description *string `pulumi:"description"`
	// Unique identifier of the resource.
	Identifier *string `pulumi:"identifier"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId *string `pulumi:"orgId"`
	// The reference to the Harness secret containing the Splunk password.
	PasswordRef *string `pulumi:"passwordRef"`
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags []string `pulumi:"tags"`
	// Url of the Splunk server.
	Url *string `pulumi:"url"`
	// The username used for connecting to Splunk.
	Username *string `pulumi:"username"`
}

type SplunkState struct {
	// Splunk account id.
	AccountId pulumi.StringPtrInput
	// Connect using only the delegates which have these tags.
	DelegateSelectors pulumi.StringArrayInput
	// Description of the resource.
	Description pulumi.StringPtrInput
	// Unique identifier of the resource.
	Identifier pulumi.StringPtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrInput
	// The reference to the Harness secret containing the Splunk password.
	PasswordRef pulumi.StringPtrInput
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrInput
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags pulumi.StringArrayInput
	// Url of the Splunk server.
	Url pulumi.StringPtrInput
	// The username used for connecting to Splunk.
	Username pulumi.StringPtrInput
}

func (SplunkState) ElementType() reflect.Type {
	return reflect.TypeOf((*splunkState)(nil)).Elem()
}

type splunkArgs struct {
	// Splunk account id.
	AccountId string `pulumi:"accountId"`
	// Connect using only the delegates which have these tags.
	DelegateSelectors []string `pulumi:"delegateSelectors"`
	// Description of the resource.
	Description *string `pulumi:"description"`
	// Unique identifier of the resource.
	Identifier string `pulumi:"identifier"`
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

// The set of arguments for constructing a Splunk resource.
type SplunkArgs struct {
	// Splunk account id.
	AccountId pulumi.StringInput
	// Connect using only the delegates which have these tags.
	DelegateSelectors pulumi.StringArrayInput
	// Description of the resource.
	Description pulumi.StringPtrInput
	// Unique identifier of the resource.
	Identifier pulumi.StringInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrInput
	// The reference to the Harness secret containing the Splunk password.
	PasswordRef pulumi.StringInput
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrInput
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags pulumi.StringArrayInput
	// Url of the Splunk server.
	Url pulumi.StringInput
	// The username used for connecting to Splunk.
	Username pulumi.StringInput
}

func (SplunkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*splunkArgs)(nil)).Elem()
}

type SplunkInput interface {
	pulumi.Input

	ToSplunkOutput() SplunkOutput
	ToSplunkOutputWithContext(ctx context.Context) SplunkOutput
}

func (*Splunk) ElementType() reflect.Type {
	return reflect.TypeOf((**Splunk)(nil)).Elem()
}

func (i *Splunk) ToSplunkOutput() SplunkOutput {
	return i.ToSplunkOutputWithContext(context.Background())
}

func (i *Splunk) ToSplunkOutputWithContext(ctx context.Context) SplunkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SplunkOutput)
}

// SplunkArrayInput is an input type that accepts SplunkArray and SplunkArrayOutput values.
// You can construct a concrete instance of `SplunkArrayInput` via:
//
//	SplunkArray{ SplunkArgs{...} }
type SplunkArrayInput interface {
	pulumi.Input

	ToSplunkArrayOutput() SplunkArrayOutput
	ToSplunkArrayOutputWithContext(context.Context) SplunkArrayOutput
}

type SplunkArray []SplunkInput

func (SplunkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Splunk)(nil)).Elem()
}

func (i SplunkArray) ToSplunkArrayOutput() SplunkArrayOutput {
	return i.ToSplunkArrayOutputWithContext(context.Background())
}

func (i SplunkArray) ToSplunkArrayOutputWithContext(ctx context.Context) SplunkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SplunkArrayOutput)
}

// SplunkMapInput is an input type that accepts SplunkMap and SplunkMapOutput values.
// You can construct a concrete instance of `SplunkMapInput` via:
//
//	SplunkMap{ "key": SplunkArgs{...} }
type SplunkMapInput interface {
	pulumi.Input

	ToSplunkMapOutput() SplunkMapOutput
	ToSplunkMapOutputWithContext(context.Context) SplunkMapOutput
}

type SplunkMap map[string]SplunkInput

func (SplunkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Splunk)(nil)).Elem()
}

func (i SplunkMap) ToSplunkMapOutput() SplunkMapOutput {
	return i.ToSplunkMapOutputWithContext(context.Background())
}

func (i SplunkMap) ToSplunkMapOutputWithContext(ctx context.Context) SplunkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SplunkMapOutput)
}

type SplunkOutput struct{ *pulumi.OutputState }

func (SplunkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Splunk)(nil)).Elem()
}

func (o SplunkOutput) ToSplunkOutput() SplunkOutput {
	return o
}

func (o SplunkOutput) ToSplunkOutputWithContext(ctx context.Context) SplunkOutput {
	return o
}

// Splunk account id.
func (o SplunkOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Splunk) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// Connect using only the delegates which have these tags.
func (o SplunkOutput) DelegateSelectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Splunk) pulumi.StringArrayOutput { return v.DelegateSelectors }).(pulumi.StringArrayOutput)
}

// Description of the resource.
func (o SplunkOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Splunk) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Unique identifier of the resource.
func (o SplunkOutput) Identifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Splunk) pulumi.StringOutput { return v.Identifier }).(pulumi.StringOutput)
}

// Name of the resource.
func (o SplunkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Splunk) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Unique identifier of the organization.
func (o SplunkOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Splunk) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// The reference to the Harness secret containing the Splunk password.
func (o SplunkOutput) PasswordRef() pulumi.StringOutput {
	return o.ApplyT(func(v *Splunk) pulumi.StringOutput { return v.PasswordRef }).(pulumi.StringOutput)
}

// Unique identifier of the project.
func (o SplunkOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Splunk) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// Tags to associate with the resource. Tags should be in the form `name:value`.
func (o SplunkOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Splunk) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Url of the Splunk server.
func (o SplunkOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Splunk) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The username used for connecting to Splunk.
func (o SplunkOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *Splunk) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type SplunkArrayOutput struct{ *pulumi.OutputState }

func (SplunkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Splunk)(nil)).Elem()
}

func (o SplunkArrayOutput) ToSplunkArrayOutput() SplunkArrayOutput {
	return o
}

func (o SplunkArrayOutput) ToSplunkArrayOutputWithContext(ctx context.Context) SplunkArrayOutput {
	return o
}

func (o SplunkArrayOutput) Index(i pulumi.IntInput) SplunkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Splunk {
		return vs[0].([]*Splunk)[vs[1].(int)]
	}).(SplunkOutput)
}

type SplunkMapOutput struct{ *pulumi.OutputState }

func (SplunkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Splunk)(nil)).Elem()
}

func (o SplunkMapOutput) ToSplunkMapOutput() SplunkMapOutput {
	return o
}

func (o SplunkMapOutput) ToSplunkMapOutputWithContext(ctx context.Context) SplunkMapOutput {
	return o
}

func (o SplunkMapOutput) MapIndex(k pulumi.StringInput) SplunkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Splunk {
		return vs[0].(map[string]*Splunk)[vs[1].(string)]
	}).(SplunkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SplunkInput)(nil)).Elem(), &Splunk{})
	pulumi.RegisterInputType(reflect.TypeOf((*SplunkArrayInput)(nil)).Elem(), SplunkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SplunkMapInput)(nil)).Elem(), SplunkMap{})
	pulumi.RegisterOutputType(SplunkOutput{})
	pulumi.RegisterOutputType(SplunkArrayOutput{})
	pulumi.RegisterOutputType(SplunkMapOutput{})
}

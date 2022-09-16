// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package platformconnector

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for creating a Datadog connector.
type Datadog struct {
	pulumi.CustomResourceState

	// Reference to the Harness secret containing the api key.
	ApiKeyRef pulumi.StringOutput `pulumi:"apiKeyRef"`
	// Reference to the Harness secret containing the application key.
	ApplicationKeyRef pulumi.StringOutput `pulumi:"applicationKeyRef"`
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
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Url of the Datadog server.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewDatadog registers a new resource with the given unique name, arguments, and options.
func NewDatadog(ctx *pulumi.Context,
	name string, args *DatadogArgs, opts ...pulumi.ResourceOption) (*Datadog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiKeyRef == nil {
		return nil, errors.New("invalid value for required argument 'ApiKeyRef'")
	}
	if args.ApplicationKeyRef == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationKeyRef'")
	}
	if args.Identifier == nil {
		return nil, errors.New("invalid value for required argument 'Identifier'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Datadog
	err := ctx.RegisterResource("harness:PlatformConnector/datadog:Datadog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatadog gets an existing Datadog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatadog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatadogState, opts ...pulumi.ResourceOption) (*Datadog, error) {
	var resource Datadog
	err := ctx.ReadResource("harness:PlatformConnector/datadog:Datadog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Datadog resources.
type datadogState struct {
	// Reference to the Harness secret containing the api key.
	ApiKeyRef *string `pulumi:"apiKeyRef"`
	// Reference to the Harness secret containing the application key.
	ApplicationKeyRef *string `pulumi:"applicationKeyRef"`
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
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags []string `pulumi:"tags"`
	// Url of the Datadog server.
	Url *string `pulumi:"url"`
}

type DatadogState struct {
	// Reference to the Harness secret containing the api key.
	ApiKeyRef pulumi.StringPtrInput
	// Reference to the Harness secret containing the application key.
	ApplicationKeyRef pulumi.StringPtrInput
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
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrInput
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags pulumi.StringArrayInput
	// Url of the Datadog server.
	Url pulumi.StringPtrInput
}

func (DatadogState) ElementType() reflect.Type {
	return reflect.TypeOf((*datadogState)(nil)).Elem()
}

type datadogArgs struct {
	// Reference to the Harness secret containing the api key.
	ApiKeyRef string `pulumi:"apiKeyRef"`
	// Reference to the Harness secret containing the application key.
	ApplicationKeyRef string `pulumi:"applicationKeyRef"`
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
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags []string `pulumi:"tags"`
	// Url of the Datadog server.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a Datadog resource.
type DatadogArgs struct {
	// Reference to the Harness secret containing the api key.
	ApiKeyRef pulumi.StringInput
	// Reference to the Harness secret containing the application key.
	ApplicationKeyRef pulumi.StringInput
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
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrInput
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags pulumi.StringArrayInput
	// Url of the Datadog server.
	Url pulumi.StringInput
}

func (DatadogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datadogArgs)(nil)).Elem()
}

type DatadogInput interface {
	pulumi.Input

	ToDatadogOutput() DatadogOutput
	ToDatadogOutputWithContext(ctx context.Context) DatadogOutput
}

func (*Datadog) ElementType() reflect.Type {
	return reflect.TypeOf((**Datadog)(nil)).Elem()
}

func (i *Datadog) ToDatadogOutput() DatadogOutput {
	return i.ToDatadogOutputWithContext(context.Background())
}

func (i *Datadog) ToDatadogOutputWithContext(ctx context.Context) DatadogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogOutput)
}

// DatadogArrayInput is an input type that accepts DatadogArray and DatadogArrayOutput values.
// You can construct a concrete instance of `DatadogArrayInput` via:
//
//	DatadogArray{ DatadogArgs{...} }
type DatadogArrayInput interface {
	pulumi.Input

	ToDatadogArrayOutput() DatadogArrayOutput
	ToDatadogArrayOutputWithContext(context.Context) DatadogArrayOutput
}

type DatadogArray []DatadogInput

func (DatadogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Datadog)(nil)).Elem()
}

func (i DatadogArray) ToDatadogArrayOutput() DatadogArrayOutput {
	return i.ToDatadogArrayOutputWithContext(context.Background())
}

func (i DatadogArray) ToDatadogArrayOutputWithContext(ctx context.Context) DatadogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogArrayOutput)
}

// DatadogMapInput is an input type that accepts DatadogMap and DatadogMapOutput values.
// You can construct a concrete instance of `DatadogMapInput` via:
//
//	DatadogMap{ "key": DatadogArgs{...} }
type DatadogMapInput interface {
	pulumi.Input

	ToDatadogMapOutput() DatadogMapOutput
	ToDatadogMapOutputWithContext(context.Context) DatadogMapOutput
}

type DatadogMap map[string]DatadogInput

func (DatadogMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Datadog)(nil)).Elem()
}

func (i DatadogMap) ToDatadogMapOutput() DatadogMapOutput {
	return i.ToDatadogMapOutputWithContext(context.Background())
}

func (i DatadogMap) ToDatadogMapOutputWithContext(ctx context.Context) DatadogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatadogMapOutput)
}

type DatadogOutput struct{ *pulumi.OutputState }

func (DatadogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Datadog)(nil)).Elem()
}

func (o DatadogOutput) ToDatadogOutput() DatadogOutput {
	return o
}

func (o DatadogOutput) ToDatadogOutputWithContext(ctx context.Context) DatadogOutput {
	return o
}

// Reference to the Harness secret containing the api key.
func (o DatadogOutput) ApiKeyRef() pulumi.StringOutput {
	return o.ApplyT(func(v *Datadog) pulumi.StringOutput { return v.ApiKeyRef }).(pulumi.StringOutput)
}

// Reference to the Harness secret containing the application key.
func (o DatadogOutput) ApplicationKeyRef() pulumi.StringOutput {
	return o.ApplyT(func(v *Datadog) pulumi.StringOutput { return v.ApplicationKeyRef }).(pulumi.StringOutput)
}

// Connect using only the delegates which have these tags.
func (o DatadogOutput) DelegateSelectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Datadog) pulumi.StringArrayOutput { return v.DelegateSelectors }).(pulumi.StringArrayOutput)
}

// Description of the resource.
func (o DatadogOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datadog) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Unique identifier of the resource.
func (o DatadogOutput) Identifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Datadog) pulumi.StringOutput { return v.Identifier }).(pulumi.StringOutput)
}

// Name of the resource.
func (o DatadogOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Datadog) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Unique identifier of the organization.
func (o DatadogOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datadog) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Unique identifier of the project.
func (o DatadogOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datadog) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// Tags to associate with the resource. Tags should be in the form `name:value`.
func (o DatadogOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Datadog) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Url of the Datadog server.
func (o DatadogOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Datadog) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type DatadogArrayOutput struct{ *pulumi.OutputState }

func (DatadogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Datadog)(nil)).Elem()
}

func (o DatadogArrayOutput) ToDatadogArrayOutput() DatadogArrayOutput {
	return o
}

func (o DatadogArrayOutput) ToDatadogArrayOutputWithContext(ctx context.Context) DatadogArrayOutput {
	return o
}

func (o DatadogArrayOutput) Index(i pulumi.IntInput) DatadogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Datadog {
		return vs[0].([]*Datadog)[vs[1].(int)]
	}).(DatadogOutput)
}

type DatadogMapOutput struct{ *pulumi.OutputState }

func (DatadogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Datadog)(nil)).Elem()
}

func (o DatadogMapOutput) ToDatadogMapOutput() DatadogMapOutput {
	return o
}

func (o DatadogMapOutput) ToDatadogMapOutputWithContext(ctx context.Context) DatadogMapOutput {
	return o
}

func (o DatadogMapOutput) MapIndex(k pulumi.StringInput) DatadogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Datadog {
		return vs[0].(map[string]*Datadog)[vs[1].(string)]
	}).(DatadogOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatadogInput)(nil)).Elem(), &Datadog{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatadogArrayInput)(nil)).Elem(), DatadogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatadogMapInput)(nil)).Elem(), DatadogMap{})
	pulumi.RegisterOutputType(DatadogOutput{})
	pulumi.RegisterOutputType(DatadogArrayOutput{})
	pulumi.RegisterOutputType(DatadogMapOutput{})
}

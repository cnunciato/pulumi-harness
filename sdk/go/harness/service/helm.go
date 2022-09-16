// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package service

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for creating a Kubernetes Helm service. This resource uses the config-as-code API's. When updating the `name` or `path` of this resource you should typically also set the `createBeforeDestroy = true` lifecycle setting.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-harness/sdk/go/harness"
//	"github.com/lbrlabs/pulumi-harness/sdk/go/harness/Service"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleApplication, err := harness.NewApplication(ctx, "exampleApplication", nil)
//			if err != nil {
//				return err
//			}
//			_, err = Service.NewHelm(ctx, "exampleHelm", &Service.HelmArgs{
//				AppId:       exampleApplication.ID(),
//				Description: pulumi.String("Service for deploying native Helm application.s"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # Import using the Harness application id and service id
//
// ```sh
//
//	$ pulumi import harness:Service/helm:Helm example <app_id>/<svc_id>
//
// ```
type Helm struct {
	pulumi.CustomResourceState

	// The id of the application the service belongs to
	AppId pulumi.StringOutput `pulumi:"appId"`
	// Description of th service
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the service
	Name pulumi.StringOutput `pulumi:"name"`
	// Variables to be used in the service
	Variables HelmVariableArrayOutput `pulumi:"variables"`
}

// NewHelm registers a new resource with the given unique name, arguments, and options.
func NewHelm(ctx *pulumi.Context,
	name string, args *HelmArgs, opts ...pulumi.ResourceOption) (*Helm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Helm
	err := ctx.RegisterResource("harness:Service/helm:Helm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHelm gets an existing Helm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHelm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HelmState, opts ...pulumi.ResourceOption) (*Helm, error) {
	var resource Helm
	err := ctx.ReadResource("harness:Service/helm:Helm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Helm resources.
type helmState struct {
	// The id of the application the service belongs to
	AppId *string `pulumi:"appId"`
	// Description of th service
	Description *string `pulumi:"description"`
	// Name of the service
	Name *string `pulumi:"name"`
	// Variables to be used in the service
	Variables []HelmVariable `pulumi:"variables"`
}

type HelmState struct {
	// The id of the application the service belongs to
	AppId pulumi.StringPtrInput
	// Description of th service
	Description pulumi.StringPtrInput
	// Name of the service
	Name pulumi.StringPtrInput
	// Variables to be used in the service
	Variables HelmVariableArrayInput
}

func (HelmState) ElementType() reflect.Type {
	return reflect.TypeOf((*helmState)(nil)).Elem()
}

type helmArgs struct {
	// The id of the application the service belongs to
	AppId string `pulumi:"appId"`
	// Description of th service
	Description *string `pulumi:"description"`
	// Name of the service
	Name *string `pulumi:"name"`
	// Variables to be used in the service
	Variables []HelmVariable `pulumi:"variables"`
}

// The set of arguments for constructing a Helm resource.
type HelmArgs struct {
	// The id of the application the service belongs to
	AppId pulumi.StringInput
	// Description of th service
	Description pulumi.StringPtrInput
	// Name of the service
	Name pulumi.StringPtrInput
	// Variables to be used in the service
	Variables HelmVariableArrayInput
}

func (HelmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*helmArgs)(nil)).Elem()
}

type HelmInput interface {
	pulumi.Input

	ToHelmOutput() HelmOutput
	ToHelmOutputWithContext(ctx context.Context) HelmOutput
}

func (*Helm) ElementType() reflect.Type {
	return reflect.TypeOf((**Helm)(nil)).Elem()
}

func (i *Helm) ToHelmOutput() HelmOutput {
	return i.ToHelmOutputWithContext(context.Background())
}

func (i *Helm) ToHelmOutputWithContext(ctx context.Context) HelmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmOutput)
}

// HelmArrayInput is an input type that accepts HelmArray and HelmArrayOutput values.
// You can construct a concrete instance of `HelmArrayInput` via:
//
//	HelmArray{ HelmArgs{...} }
type HelmArrayInput interface {
	pulumi.Input

	ToHelmArrayOutput() HelmArrayOutput
	ToHelmArrayOutputWithContext(context.Context) HelmArrayOutput
}

type HelmArray []HelmInput

func (HelmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Helm)(nil)).Elem()
}

func (i HelmArray) ToHelmArrayOutput() HelmArrayOutput {
	return i.ToHelmArrayOutputWithContext(context.Background())
}

func (i HelmArray) ToHelmArrayOutputWithContext(ctx context.Context) HelmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmArrayOutput)
}

// HelmMapInput is an input type that accepts HelmMap and HelmMapOutput values.
// You can construct a concrete instance of `HelmMapInput` via:
//
//	HelmMap{ "key": HelmArgs{...} }
type HelmMapInput interface {
	pulumi.Input

	ToHelmMapOutput() HelmMapOutput
	ToHelmMapOutputWithContext(context.Context) HelmMapOutput
}

type HelmMap map[string]HelmInput

func (HelmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Helm)(nil)).Elem()
}

func (i HelmMap) ToHelmMapOutput() HelmMapOutput {
	return i.ToHelmMapOutputWithContext(context.Background())
}

func (i HelmMap) ToHelmMapOutputWithContext(ctx context.Context) HelmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmMapOutput)
}

type HelmOutput struct{ *pulumi.OutputState }

func (HelmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Helm)(nil)).Elem()
}

func (o HelmOutput) ToHelmOutput() HelmOutput {
	return o
}

func (o HelmOutput) ToHelmOutputWithContext(ctx context.Context) HelmOutput {
	return o
}

// The id of the application the service belongs to
func (o HelmOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *Helm) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

// Description of th service
func (o HelmOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Helm) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of the service
func (o HelmOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Helm) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Variables to be used in the service
func (o HelmOutput) Variables() HelmVariableArrayOutput {
	return o.ApplyT(func(v *Helm) HelmVariableArrayOutput { return v.Variables }).(HelmVariableArrayOutput)
}

type HelmArrayOutput struct{ *pulumi.OutputState }

func (HelmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Helm)(nil)).Elem()
}

func (o HelmArrayOutput) ToHelmArrayOutput() HelmArrayOutput {
	return o
}

func (o HelmArrayOutput) ToHelmArrayOutputWithContext(ctx context.Context) HelmArrayOutput {
	return o
}

func (o HelmArrayOutput) Index(i pulumi.IntInput) HelmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Helm {
		return vs[0].([]*Helm)[vs[1].(int)]
	}).(HelmOutput)
}

type HelmMapOutput struct{ *pulumi.OutputState }

func (HelmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Helm)(nil)).Elem()
}

func (o HelmMapOutput) ToHelmMapOutput() HelmMapOutput {
	return o
}

func (o HelmMapOutput) ToHelmMapOutputWithContext(ctx context.Context) HelmMapOutput {
	return o
}

func (o HelmMapOutput) MapIndex(k pulumi.StringInput) HelmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Helm {
		return vs[0].(map[string]*Helm)[vs[1].(string)]
	}).(HelmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HelmInput)(nil)).Elem(), &Helm{})
	pulumi.RegisterInputType(reflect.TypeOf((*HelmArrayInput)(nil)).Elem(), HelmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HelmMapInput)(nil)).Elem(), HelmMap{})
	pulumi.RegisterOutputType(HelmOutput{})
	pulumi.RegisterOutputType(HelmArrayOutput{})
	pulumi.RegisterOutputType(HelmMapOutput{})
}

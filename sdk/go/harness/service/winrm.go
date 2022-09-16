// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package service

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for creating an WinRM service. This resource uses the config-as-code API's. When updating the `name` or `path` of this resource you should typically also set the `createBeforeDestroy = true` lifecycle setting.
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
//			_, err = Service.NewWinrm(ctx, "exampleWinrm", &Service.WinrmArgs{
//				AppId:        exampleApplication.ID(),
//				ArtifactType: pulumi.String("IIS_APP"),
//				Description:  pulumi.String("Service for deploying IIS appliactions using winrm."),
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
//	$ pulumi import harness:Service/winrm:Winrm example <app_id>/<svc_id>
//
// ```
type Winrm struct {
	pulumi.CustomResourceState

	// The id of the application the service belongs to
	AppId pulumi.StringOutput `pulumi:"appId"`
	// The type of artifact to deploy.
	ArtifactType pulumi.StringOutput `pulumi:"artifactType"`
	// Description of th service
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the service
	Name pulumi.StringOutput `pulumi:"name"`
	// Variables to be used in the service
	Variables WinrmVariableArrayOutput `pulumi:"variables"`
}

// NewWinrm registers a new resource with the given unique name, arguments, and options.
func NewWinrm(ctx *pulumi.Context,
	name string, args *WinrmArgs, opts ...pulumi.ResourceOption) (*Winrm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.ArtifactType == nil {
		return nil, errors.New("invalid value for required argument 'ArtifactType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Winrm
	err := ctx.RegisterResource("harness:Service/winrm:Winrm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWinrm gets an existing Winrm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWinrm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WinrmState, opts ...pulumi.ResourceOption) (*Winrm, error) {
	var resource Winrm
	err := ctx.ReadResource("harness:Service/winrm:Winrm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Winrm resources.
type winrmState struct {
	// The id of the application the service belongs to
	AppId *string `pulumi:"appId"`
	// The type of artifact to deploy.
	ArtifactType *string `pulumi:"artifactType"`
	// Description of th service
	Description *string `pulumi:"description"`
	// Name of the service
	Name *string `pulumi:"name"`
	// Variables to be used in the service
	Variables []WinrmVariable `pulumi:"variables"`
}

type WinrmState struct {
	// The id of the application the service belongs to
	AppId pulumi.StringPtrInput
	// The type of artifact to deploy.
	ArtifactType pulumi.StringPtrInput
	// Description of th service
	Description pulumi.StringPtrInput
	// Name of the service
	Name pulumi.StringPtrInput
	// Variables to be used in the service
	Variables WinrmVariableArrayInput
}

func (WinrmState) ElementType() reflect.Type {
	return reflect.TypeOf((*winrmState)(nil)).Elem()
}

type winrmArgs struct {
	// The id of the application the service belongs to
	AppId string `pulumi:"appId"`
	// The type of artifact to deploy.
	ArtifactType string `pulumi:"artifactType"`
	// Description of th service
	Description *string `pulumi:"description"`
	// Name of the service
	Name *string `pulumi:"name"`
	// Variables to be used in the service
	Variables []WinrmVariable `pulumi:"variables"`
}

// The set of arguments for constructing a Winrm resource.
type WinrmArgs struct {
	// The id of the application the service belongs to
	AppId pulumi.StringInput
	// The type of artifact to deploy.
	ArtifactType pulumi.StringInput
	// Description of th service
	Description pulumi.StringPtrInput
	// Name of the service
	Name pulumi.StringPtrInput
	// Variables to be used in the service
	Variables WinrmVariableArrayInput
}

func (WinrmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*winrmArgs)(nil)).Elem()
}

type WinrmInput interface {
	pulumi.Input

	ToWinrmOutput() WinrmOutput
	ToWinrmOutputWithContext(ctx context.Context) WinrmOutput
}

func (*Winrm) ElementType() reflect.Type {
	return reflect.TypeOf((**Winrm)(nil)).Elem()
}

func (i *Winrm) ToWinrmOutput() WinrmOutput {
	return i.ToWinrmOutputWithContext(context.Background())
}

func (i *Winrm) ToWinrmOutputWithContext(ctx context.Context) WinrmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WinrmOutput)
}

// WinrmArrayInput is an input type that accepts WinrmArray and WinrmArrayOutput values.
// You can construct a concrete instance of `WinrmArrayInput` via:
//
//	WinrmArray{ WinrmArgs{...} }
type WinrmArrayInput interface {
	pulumi.Input

	ToWinrmArrayOutput() WinrmArrayOutput
	ToWinrmArrayOutputWithContext(context.Context) WinrmArrayOutput
}

type WinrmArray []WinrmInput

func (WinrmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Winrm)(nil)).Elem()
}

func (i WinrmArray) ToWinrmArrayOutput() WinrmArrayOutput {
	return i.ToWinrmArrayOutputWithContext(context.Background())
}

func (i WinrmArray) ToWinrmArrayOutputWithContext(ctx context.Context) WinrmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WinrmArrayOutput)
}

// WinrmMapInput is an input type that accepts WinrmMap and WinrmMapOutput values.
// You can construct a concrete instance of `WinrmMapInput` via:
//
//	WinrmMap{ "key": WinrmArgs{...} }
type WinrmMapInput interface {
	pulumi.Input

	ToWinrmMapOutput() WinrmMapOutput
	ToWinrmMapOutputWithContext(context.Context) WinrmMapOutput
}

type WinrmMap map[string]WinrmInput

func (WinrmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Winrm)(nil)).Elem()
}

func (i WinrmMap) ToWinrmMapOutput() WinrmMapOutput {
	return i.ToWinrmMapOutputWithContext(context.Background())
}

func (i WinrmMap) ToWinrmMapOutputWithContext(ctx context.Context) WinrmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WinrmMapOutput)
}

type WinrmOutput struct{ *pulumi.OutputState }

func (WinrmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Winrm)(nil)).Elem()
}

func (o WinrmOutput) ToWinrmOutput() WinrmOutput {
	return o
}

func (o WinrmOutput) ToWinrmOutputWithContext(ctx context.Context) WinrmOutput {
	return o
}

// The id of the application the service belongs to
func (o WinrmOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *Winrm) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

// The type of artifact to deploy.
func (o WinrmOutput) ArtifactType() pulumi.StringOutput {
	return o.ApplyT(func(v *Winrm) pulumi.StringOutput { return v.ArtifactType }).(pulumi.StringOutput)
}

// Description of th service
func (o WinrmOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Winrm) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of the service
func (o WinrmOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Winrm) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Variables to be used in the service
func (o WinrmOutput) Variables() WinrmVariableArrayOutput {
	return o.ApplyT(func(v *Winrm) WinrmVariableArrayOutput { return v.Variables }).(WinrmVariableArrayOutput)
}

type WinrmArrayOutput struct{ *pulumi.OutputState }

func (WinrmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Winrm)(nil)).Elem()
}

func (o WinrmArrayOutput) ToWinrmArrayOutput() WinrmArrayOutput {
	return o
}

func (o WinrmArrayOutput) ToWinrmArrayOutputWithContext(ctx context.Context) WinrmArrayOutput {
	return o
}

func (o WinrmArrayOutput) Index(i pulumi.IntInput) WinrmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Winrm {
		return vs[0].([]*Winrm)[vs[1].(int)]
	}).(WinrmOutput)
}

type WinrmMapOutput struct{ *pulumi.OutputState }

func (WinrmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Winrm)(nil)).Elem()
}

func (o WinrmMapOutput) ToWinrmMapOutput() WinrmMapOutput {
	return o
}

func (o WinrmMapOutput) ToWinrmMapOutputWithContext(ctx context.Context) WinrmMapOutput {
	return o
}

func (o WinrmMapOutput) MapIndex(k pulumi.StringInput) WinrmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Winrm {
		return vs[0].(map[string]*Winrm)[vs[1].(string)]
	}).(WinrmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WinrmInput)(nil)).Elem(), &Winrm{})
	pulumi.RegisterInputType(reflect.TypeOf((*WinrmArrayInput)(nil)).Elem(), WinrmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WinrmMapInput)(nil)).Elem(), WinrmMap{})
	pulumi.RegisterOutputType(WinrmOutput{})
	pulumi.RegisterOutputType(WinrmArrayOutput{})
	pulumi.RegisterOutputType(WinrmMapOutput{})
}

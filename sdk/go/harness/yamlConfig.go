// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package harness

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for creating a raw YAML configuration in Harness. Note: This works for all objects EXCEPT application objects. This resource uses the config-as-code API's. When updating the `name` or `path` of this resource you should typically also set the `createBeforeDestroy = true` lifecycle setting.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/lbrlabs/pulumi-harness/sdk/go/harness"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := harness.NewYamlConfig(ctx, "test", &harness.YamlConfigArgs{
//				Content: pulumi.String(fmt.Sprintf(`harnessApiVersion: '1.0'
//
// type: KUBERNETES_CLUSTER
// delegateSelectors:
// - k8s
// skipValidation: true
// useKubernetesDelegate: true
//
// `)),
//
//				Path: pulumi.String("Setup/Cloud Providers/Kubernetes.yaml"),
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
// # Importing a global config only using the yaml path
//
// ```sh
//
//	$ pulumi import harness:index/yamlConfig:YamlConfig k8s_cloudprovider "Setup/Cloud Providers/kubernetes.yaml"
//
// ```
//
// # Importing a service which requires both the application id and the yaml path.
//
// ```sh
//
//	$ pulumi import harness:index/yamlConfig:YamlConfig k8s_cloudprovider "Setup/Applications/MyApp/Services/MyService/Index.yaml:<APPLICATION_ID>"
//
// ```
type YamlConfig struct {
	pulumi.CustomResourceState

	// The id of the application. This is required for all resources except global ones.
	AppId pulumi.StringPtrOutput `pulumi:"appId"`
	// The raw YAML configuration.
	Content pulumi.StringOutput `pulumi:"content"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The path of the resource.
	Path pulumi.StringOutput `pulumi:"path"`
}

// NewYamlConfig registers a new resource with the given unique name, arguments, and options.
func NewYamlConfig(ctx *pulumi.Context,
	name string, args *YamlConfigArgs, opts ...pulumi.ResourceOption) (*YamlConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource YamlConfig
	err := ctx.RegisterResource("harness:index/yamlConfig:YamlConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetYamlConfig gets an existing YamlConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetYamlConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *YamlConfigState, opts ...pulumi.ResourceOption) (*YamlConfig, error) {
	var resource YamlConfig
	err := ctx.ReadResource("harness:index/yamlConfig:YamlConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering YamlConfig resources.
type yamlConfigState struct {
	// The id of the application. This is required for all resources except global ones.
	AppId *string `pulumi:"appId"`
	// The raw YAML configuration.
	Content *string `pulumi:"content"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The path of the resource.
	Path *string `pulumi:"path"`
}

type YamlConfigState struct {
	// The id of the application. This is required for all resources except global ones.
	AppId pulumi.StringPtrInput
	// The raw YAML configuration.
	Content pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The path of the resource.
	Path pulumi.StringPtrInput
}

func (YamlConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*yamlConfigState)(nil)).Elem()
}

type yamlConfigArgs struct {
	// The id of the application. This is required for all resources except global ones.
	AppId *string `pulumi:"appId"`
	// The raw YAML configuration.
	Content string `pulumi:"content"`
	// The path of the resource.
	Path string `pulumi:"path"`
}

// The set of arguments for constructing a YamlConfig resource.
type YamlConfigArgs struct {
	// The id of the application. This is required for all resources except global ones.
	AppId pulumi.StringPtrInput
	// The raw YAML configuration.
	Content pulumi.StringInput
	// The path of the resource.
	Path pulumi.StringInput
}

func (YamlConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*yamlConfigArgs)(nil)).Elem()
}

type YamlConfigInput interface {
	pulumi.Input

	ToYamlConfigOutput() YamlConfigOutput
	ToYamlConfigOutputWithContext(ctx context.Context) YamlConfigOutput
}

func (*YamlConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**YamlConfig)(nil)).Elem()
}

func (i *YamlConfig) ToYamlConfigOutput() YamlConfigOutput {
	return i.ToYamlConfigOutputWithContext(context.Background())
}

func (i *YamlConfig) ToYamlConfigOutputWithContext(ctx context.Context) YamlConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YamlConfigOutput)
}

// YamlConfigArrayInput is an input type that accepts YamlConfigArray and YamlConfigArrayOutput values.
// You can construct a concrete instance of `YamlConfigArrayInput` via:
//
//	YamlConfigArray{ YamlConfigArgs{...} }
type YamlConfigArrayInput interface {
	pulumi.Input

	ToYamlConfigArrayOutput() YamlConfigArrayOutput
	ToYamlConfigArrayOutputWithContext(context.Context) YamlConfigArrayOutput
}

type YamlConfigArray []YamlConfigInput

func (YamlConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*YamlConfig)(nil)).Elem()
}

func (i YamlConfigArray) ToYamlConfigArrayOutput() YamlConfigArrayOutput {
	return i.ToYamlConfigArrayOutputWithContext(context.Background())
}

func (i YamlConfigArray) ToYamlConfigArrayOutputWithContext(ctx context.Context) YamlConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YamlConfigArrayOutput)
}

// YamlConfigMapInput is an input type that accepts YamlConfigMap and YamlConfigMapOutput values.
// You can construct a concrete instance of `YamlConfigMapInput` via:
//
//	YamlConfigMap{ "key": YamlConfigArgs{...} }
type YamlConfigMapInput interface {
	pulumi.Input

	ToYamlConfigMapOutput() YamlConfigMapOutput
	ToYamlConfigMapOutputWithContext(context.Context) YamlConfigMapOutput
}

type YamlConfigMap map[string]YamlConfigInput

func (YamlConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*YamlConfig)(nil)).Elem()
}

func (i YamlConfigMap) ToYamlConfigMapOutput() YamlConfigMapOutput {
	return i.ToYamlConfigMapOutputWithContext(context.Background())
}

func (i YamlConfigMap) ToYamlConfigMapOutputWithContext(ctx context.Context) YamlConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YamlConfigMapOutput)
}

type YamlConfigOutput struct{ *pulumi.OutputState }

func (YamlConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**YamlConfig)(nil)).Elem()
}

func (o YamlConfigOutput) ToYamlConfigOutput() YamlConfigOutput {
	return o
}

func (o YamlConfigOutput) ToYamlConfigOutputWithContext(ctx context.Context) YamlConfigOutput {
	return o
}

// The id of the application. This is required for all resources except global ones.
func (o YamlConfigOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *YamlConfig) pulumi.StringPtrOutput { return v.AppId }).(pulumi.StringPtrOutput)
}

// The raw YAML configuration.
func (o YamlConfigOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *YamlConfig) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// The name of the resource.
func (o YamlConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *YamlConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The path of the resource.
func (o YamlConfigOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *YamlConfig) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

type YamlConfigArrayOutput struct{ *pulumi.OutputState }

func (YamlConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*YamlConfig)(nil)).Elem()
}

func (o YamlConfigArrayOutput) ToYamlConfigArrayOutput() YamlConfigArrayOutput {
	return o
}

func (o YamlConfigArrayOutput) ToYamlConfigArrayOutputWithContext(ctx context.Context) YamlConfigArrayOutput {
	return o
}

func (o YamlConfigArrayOutput) Index(i pulumi.IntInput) YamlConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *YamlConfig {
		return vs[0].([]*YamlConfig)[vs[1].(int)]
	}).(YamlConfigOutput)
}

type YamlConfigMapOutput struct{ *pulumi.OutputState }

func (YamlConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*YamlConfig)(nil)).Elem()
}

func (o YamlConfigMapOutput) ToYamlConfigMapOutput() YamlConfigMapOutput {
	return o
}

func (o YamlConfigMapOutput) ToYamlConfigMapOutputWithContext(ctx context.Context) YamlConfigMapOutput {
	return o
}

func (o YamlConfigMapOutput) MapIndex(k pulumi.StringInput) YamlConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *YamlConfig {
		return vs[0].(map[string]*YamlConfig)[vs[1].(string)]
	}).(YamlConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*YamlConfigInput)(nil)).Elem(), &YamlConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*YamlConfigArrayInput)(nil)).Elem(), YamlConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*YamlConfigMapInput)(nil)).Elem(), YamlConfigMap{})
	pulumi.RegisterOutputType(YamlConfigOutput{})
	pulumi.RegisterOutputType(YamlConfigArrayOutput{})
	pulumi.RegisterOutputType(YamlConfigMapOutput{})
}

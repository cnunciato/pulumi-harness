// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package platformconnector

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for creating a K8s connector.
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
//	"github.com/lbrlabs/pulumi-harness/sdk/go/harness/PlatformConnector"
//	"github.com/pulumi/pulumi-harness/sdk/go/harness/PlatformConnector"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := PlatformConnector.NewKubernetes(ctx, "clientKeyCert", &PlatformConnector.KubernetesArgs{
//				ClientKeyCert: &platformconnector.KubernetesClientKeyCertArgs{
//					CaCertRef:              pulumi.String("account.TEST_k8ss_client_stuff"),
//					ClientCertRef:          pulumi.String("account.test_k8s_client_cert"),
//					ClientKeyAlgorithm:     pulumi.String("RSA"),
//					ClientKeyPassphraseRef: pulumi.String("account.TEST_k8s_client_test"),
//					ClientKeyRef:           pulumi.String("account.TEST_k8s_client_key"),
//					MasterUrl:              pulumi.String("https://kubernetes.example.com"),
//				},
//				DelegateSelectors: pulumi.StringArray{
//					pulumi.String("harness-delegate"),
//				},
//				Description: pulumi.String("description"),
//				Identifier:  pulumi.String("identifier"),
//				Tags: pulumi.StringArray{
//					pulumi.String("foo:bar"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = PlatformConnector.NewKubernetes(ctx, "usernamePassword", &PlatformConnector.KubernetesArgs{
//				DelegateSelectors: pulumi.StringArray{
//					pulumi.String("harness-delegate"),
//				},
//				Description: pulumi.String("description"),
//				Identifier:  pulumi.String("identifier"),
//				Tags: pulumi.StringArray{
//					pulumi.String("foo:bar"),
//				},
//				UsernamePassword: &platformconnector.KubernetesUsernamePasswordArgs{
//					MasterUrl:   pulumi.String("https://kubernetes.example.com"),
//					PasswordRef: pulumi.String("account.TEST_k8s_client_test"),
//					Username:    pulumi.String("admin"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = PlatformConnector.NewKubernetes(ctx, "serviceAccount", &PlatformConnector.KubernetesArgs{
//				DelegateSelectors: pulumi.StringArray{
//					pulumi.String("harness-delegate"),
//				},
//				Description: pulumi.String("description"),
//				Identifier:  pulumi.String("identifier"),
//				ServiceAccount: &platformconnector.KubernetesServiceAccountArgs{
//					MasterUrl:              pulumi.String("https://kubernetes.example.com"),
//					ServiceAccountTokenRef: pulumi.String("account.TEST_k8s_client_test"),
//				},
//				Tags: pulumi.StringArray{
//					pulumi.String("foo:bar"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = PlatformConnector.NewKubernetes(ctx, "openIDConnect", &PlatformConnector.KubernetesArgs{
//				DelegateSelectors: pulumi.StringArray{
//					pulumi.String("harness-delegate"),
//				},
//				Description: pulumi.String("description"),
//				Identifier:  pulumi.String(fmt.Sprintf("%v[1]s", "%")),
//				OpenidConnect: &platformconnector.KubernetesOpenidConnectArgs{
//					ClientIdRef: pulumi.String("account.TEST_k8s_client_test"),
//					IssuerUrl:   pulumi.String("https://oidc.example.com"),
//					MasterUrl:   pulumi.String("https://kubernetes.example.com"),
//					PasswordRef: pulumi.String("account.TEST_k8s_client_test"),
//					Scopes: pulumi.StringArray{
//						pulumi.String("scope1"),
//						pulumi.String("scope2"),
//					},
//					SecretRef:   pulumi.String("account.TEST_k8s_client_test"),
//					UsernameRef: pulumi.String("account.TEST_k8s_client_test"),
//				},
//				Tags: pulumi.StringArray{
//					pulumi.String("foo:bar"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = PlatformConnector.NewKubernetes(ctx, "inheritFromDelegate", &PlatformConnector.KubernetesArgs{
//				Description: pulumi.String("description"),
//				Identifier:  pulumi.String("identifier"),
//				InheritFromDelegate: &platformconnector.KubernetesInheritFromDelegateArgs{
//					DelegateSelectors: pulumi.StringArray{
//						pulumi.String("harness-delegate"),
//					},
//				},
//				Tags: pulumi.StringArray{
//					pulumi.String("foo:bar"),
//				},
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
// # Import using kubernetes connector id
//
// ```sh
//
//	$ pulumi import harness:PlatformConnector/kubernetes:Kubernetes example <connector_id>
//
// ```
type Kubernetes struct {
	pulumi.CustomResourceState

	// Client key and certificate config for the connector.
	ClientKeyCert KubernetesClientKeyCertPtrOutput `pulumi:"clientKeyCert"`
	// Selectors to use for the delegate.
	DelegateSelectors pulumi.StringArrayOutput `pulumi:"delegateSelectors"`
	// Description of the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Unique identifier of the resource.
	Identifier pulumi.StringOutput `pulumi:"identifier"`
	// Credentials are inherited from the delegate.
	InheritFromDelegate KubernetesInheritFromDelegatePtrOutput `pulumi:"inheritFromDelegate"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// OpenID configuration for the connector.
	OpenidConnect KubernetesOpenidConnectPtrOutput `pulumi:"openidConnect"`
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	// Service account for the connector.
	ServiceAccount KubernetesServiceAccountPtrOutput `pulumi:"serviceAccount"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Username and password for the connector.
	UsernamePassword KubernetesUsernamePasswordPtrOutput `pulumi:"usernamePassword"`
}

// NewKubernetes registers a new resource with the given unique name, arguments, and options.
func NewKubernetes(ctx *pulumi.Context,
	name string, args *KubernetesArgs, opts ...pulumi.ResourceOption) (*Kubernetes, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Identifier == nil {
		return nil, errors.New("invalid value for required argument 'Identifier'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Kubernetes
	err := ctx.RegisterResource("harness:PlatformConnector/kubernetes:Kubernetes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubernetes gets an existing Kubernetes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesState, opts ...pulumi.ResourceOption) (*Kubernetes, error) {
	var resource Kubernetes
	err := ctx.ReadResource("harness:PlatformConnector/kubernetes:Kubernetes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Kubernetes resources.
type kubernetesState struct {
	// Client key and certificate config for the connector.
	ClientKeyCert *KubernetesClientKeyCert `pulumi:"clientKeyCert"`
	// Selectors to use for the delegate.
	DelegateSelectors []string `pulumi:"delegateSelectors"`
	// Description of the resource.
	Description *string `pulumi:"description"`
	// Unique identifier of the resource.
	Identifier *string `pulumi:"identifier"`
	// Credentials are inherited from the delegate.
	InheritFromDelegate *KubernetesInheritFromDelegate `pulumi:"inheritFromDelegate"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// OpenID configuration for the connector.
	OpenidConnect *KubernetesOpenidConnect `pulumi:"openidConnect"`
	// Unique identifier of the organization.
	OrgId *string `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
	// Service account for the connector.
	ServiceAccount *KubernetesServiceAccount `pulumi:"serviceAccount"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags []string `pulumi:"tags"`
	// Username and password for the connector.
	UsernamePassword *KubernetesUsernamePassword `pulumi:"usernamePassword"`
}

type KubernetesState struct {
	// Client key and certificate config for the connector.
	ClientKeyCert KubernetesClientKeyCertPtrInput
	// Selectors to use for the delegate.
	DelegateSelectors pulumi.StringArrayInput
	// Description of the resource.
	Description pulumi.StringPtrInput
	// Unique identifier of the resource.
	Identifier pulumi.StringPtrInput
	// Credentials are inherited from the delegate.
	InheritFromDelegate KubernetesInheritFromDelegatePtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// OpenID configuration for the connector.
	OpenidConnect KubernetesOpenidConnectPtrInput
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrInput
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrInput
	// Service account for the connector.
	ServiceAccount KubernetesServiceAccountPtrInput
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags pulumi.StringArrayInput
	// Username and password for the connector.
	UsernamePassword KubernetesUsernamePasswordPtrInput
}

func (KubernetesState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesState)(nil)).Elem()
}

type kubernetesArgs struct {
	// Client key and certificate config for the connector.
	ClientKeyCert *KubernetesClientKeyCert `pulumi:"clientKeyCert"`
	// Selectors to use for the delegate.
	DelegateSelectors []string `pulumi:"delegateSelectors"`
	// Description of the resource.
	Description *string `pulumi:"description"`
	// Unique identifier of the resource.
	Identifier string `pulumi:"identifier"`
	// Credentials are inherited from the delegate.
	InheritFromDelegate *KubernetesInheritFromDelegate `pulumi:"inheritFromDelegate"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// OpenID configuration for the connector.
	OpenidConnect *KubernetesOpenidConnect `pulumi:"openidConnect"`
	// Unique identifier of the organization.
	OrgId *string `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
	// Service account for the connector.
	ServiceAccount *KubernetesServiceAccount `pulumi:"serviceAccount"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags []string `pulumi:"tags"`
	// Username and password for the connector.
	UsernamePassword *KubernetesUsernamePassword `pulumi:"usernamePassword"`
}

// The set of arguments for constructing a Kubernetes resource.
type KubernetesArgs struct {
	// Client key and certificate config for the connector.
	ClientKeyCert KubernetesClientKeyCertPtrInput
	// Selectors to use for the delegate.
	DelegateSelectors pulumi.StringArrayInput
	// Description of the resource.
	Description pulumi.StringPtrInput
	// Unique identifier of the resource.
	Identifier pulumi.StringInput
	// Credentials are inherited from the delegate.
	InheritFromDelegate KubernetesInheritFromDelegatePtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// OpenID configuration for the connector.
	OpenidConnect KubernetesOpenidConnectPtrInput
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrInput
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrInput
	// Service account for the connector.
	ServiceAccount KubernetesServiceAccountPtrInput
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags pulumi.StringArrayInput
	// Username and password for the connector.
	UsernamePassword KubernetesUsernamePasswordPtrInput
}

func (KubernetesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesArgs)(nil)).Elem()
}

type KubernetesInput interface {
	pulumi.Input

	ToKubernetesOutput() KubernetesOutput
	ToKubernetesOutputWithContext(ctx context.Context) KubernetesOutput
}

func (*Kubernetes) ElementType() reflect.Type {
	return reflect.TypeOf((**Kubernetes)(nil)).Elem()
}

func (i *Kubernetes) ToKubernetesOutput() KubernetesOutput {
	return i.ToKubernetesOutputWithContext(context.Background())
}

func (i *Kubernetes) ToKubernetesOutputWithContext(ctx context.Context) KubernetesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesOutput)
}

// KubernetesArrayInput is an input type that accepts KubernetesArray and KubernetesArrayOutput values.
// You can construct a concrete instance of `KubernetesArrayInput` via:
//
//	KubernetesArray{ KubernetesArgs{...} }
type KubernetesArrayInput interface {
	pulumi.Input

	ToKubernetesArrayOutput() KubernetesArrayOutput
	ToKubernetesArrayOutputWithContext(context.Context) KubernetesArrayOutput
}

type KubernetesArray []KubernetesInput

func (KubernetesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Kubernetes)(nil)).Elem()
}

func (i KubernetesArray) ToKubernetesArrayOutput() KubernetesArrayOutput {
	return i.ToKubernetesArrayOutputWithContext(context.Background())
}

func (i KubernetesArray) ToKubernetesArrayOutputWithContext(ctx context.Context) KubernetesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesArrayOutput)
}

// KubernetesMapInput is an input type that accepts KubernetesMap and KubernetesMapOutput values.
// You can construct a concrete instance of `KubernetesMapInput` via:
//
//	KubernetesMap{ "key": KubernetesArgs{...} }
type KubernetesMapInput interface {
	pulumi.Input

	ToKubernetesMapOutput() KubernetesMapOutput
	ToKubernetesMapOutputWithContext(context.Context) KubernetesMapOutput
}

type KubernetesMap map[string]KubernetesInput

func (KubernetesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Kubernetes)(nil)).Elem()
}

func (i KubernetesMap) ToKubernetesMapOutput() KubernetesMapOutput {
	return i.ToKubernetesMapOutputWithContext(context.Background())
}

func (i KubernetesMap) ToKubernetesMapOutputWithContext(ctx context.Context) KubernetesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesMapOutput)
}

type KubernetesOutput struct{ *pulumi.OutputState }

func (KubernetesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Kubernetes)(nil)).Elem()
}

func (o KubernetesOutput) ToKubernetesOutput() KubernetesOutput {
	return o
}

func (o KubernetesOutput) ToKubernetesOutputWithContext(ctx context.Context) KubernetesOutput {
	return o
}

// Client key and certificate config for the connector.
func (o KubernetesOutput) ClientKeyCert() KubernetesClientKeyCertPtrOutput {
	return o.ApplyT(func(v *Kubernetes) KubernetesClientKeyCertPtrOutput { return v.ClientKeyCert }).(KubernetesClientKeyCertPtrOutput)
}

// Selectors to use for the delegate.
func (o KubernetesOutput) DelegateSelectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringArrayOutput { return v.DelegateSelectors }).(pulumi.StringArrayOutput)
}

// Description of the resource.
func (o KubernetesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Unique identifier of the resource.
func (o KubernetesOutput) Identifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringOutput { return v.Identifier }).(pulumi.StringOutput)
}

// Credentials are inherited from the delegate.
func (o KubernetesOutput) InheritFromDelegate() KubernetesInheritFromDelegatePtrOutput {
	return o.ApplyT(func(v *Kubernetes) KubernetesInheritFromDelegatePtrOutput { return v.InheritFromDelegate }).(KubernetesInheritFromDelegatePtrOutput)
}

// Name of the resource.
func (o KubernetesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// OpenID configuration for the connector.
func (o KubernetesOutput) OpenidConnect() KubernetesOpenidConnectPtrOutput {
	return o.ApplyT(func(v *Kubernetes) KubernetesOpenidConnectPtrOutput { return v.OpenidConnect }).(KubernetesOpenidConnectPtrOutput)
}

// Unique identifier of the organization.
func (o KubernetesOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Unique identifier of the project.
func (o KubernetesOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// Service account for the connector.
func (o KubernetesOutput) ServiceAccount() KubernetesServiceAccountPtrOutput {
	return o.ApplyT(func(v *Kubernetes) KubernetesServiceAccountPtrOutput { return v.ServiceAccount }).(KubernetesServiceAccountPtrOutput)
}

// Tags to associate with the resource. Tags should be in the form `name:value`.
func (o KubernetesOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Username and password for the connector.
func (o KubernetesOutput) UsernamePassword() KubernetesUsernamePasswordPtrOutput {
	return o.ApplyT(func(v *Kubernetes) KubernetesUsernamePasswordPtrOutput { return v.UsernamePassword }).(KubernetesUsernamePasswordPtrOutput)
}

type KubernetesArrayOutput struct{ *pulumi.OutputState }

func (KubernetesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Kubernetes)(nil)).Elem()
}

func (o KubernetesArrayOutput) ToKubernetesArrayOutput() KubernetesArrayOutput {
	return o
}

func (o KubernetesArrayOutput) ToKubernetesArrayOutputWithContext(ctx context.Context) KubernetesArrayOutput {
	return o
}

func (o KubernetesArrayOutput) Index(i pulumi.IntInput) KubernetesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Kubernetes {
		return vs[0].([]*Kubernetes)[vs[1].(int)]
	}).(KubernetesOutput)
}

type KubernetesMapOutput struct{ *pulumi.OutputState }

func (KubernetesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Kubernetes)(nil)).Elem()
}

func (o KubernetesMapOutput) ToKubernetesMapOutput() KubernetesMapOutput {
	return o
}

func (o KubernetesMapOutput) ToKubernetesMapOutputWithContext(ctx context.Context) KubernetesMapOutput {
	return o
}

func (o KubernetesMapOutput) MapIndex(k pulumi.StringInput) KubernetesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Kubernetes {
		return vs[0].(map[string]*Kubernetes)[vs[1].(string)]
	}).(KubernetesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesInput)(nil)).Elem(), &Kubernetes{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesArrayInput)(nil)).Elem(), KubernetesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesMapInput)(nil)).Elem(), KubernetesMap{})
	pulumi.RegisterOutputType(KubernetesOutput{})
	pulumi.RegisterOutputType(KubernetesArrayOutput{})
	pulumi.RegisterOutputType(KubernetesMapOutput{})
}

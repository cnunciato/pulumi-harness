// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package platformconnector

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for creating an AWS Cloud Cost connector.
type AwsCC struct {
	pulumi.CustomResourceState

	// The AWS account id.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// Harness uses the secure cross-account role to access your AWS account. The role includes a restricted policy to access the cost and usage reports and resources for the sole purpose of cost analysis and cost optimization.
	CrossAccountAccess AwsCCCrossAccountAccessOutput `pulumi:"crossAccountAccess"`
	// Description of the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The features enabled for the connector. Valid values are BILLING, OPTIMIZATION, VISIBILITY.
	FeaturesEnableds pulumi.StringArrayOutput `pulumi:"featuresEnableds"`
	// Unique identifier of the resource.
	Identifier pulumi.StringOutput `pulumi:"identifier"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	// The cost and usage report name. Provided in the delivery options when the template is opened in the AWS console.
	ReportName pulumi.StringOutput `pulumi:"reportName"`
	// The name of s3 bucket.
	S3Bucket pulumi.StringOutput `pulumi:"s3Bucket"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewAwsCC registers a new resource with the given unique name, arguments, and options.
func NewAwsCC(ctx *pulumi.Context,
	name string, args *AwsCCArgs, opts ...pulumi.ResourceOption) (*AwsCC, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	if args.CrossAccountAccess == nil {
		return nil, errors.New("invalid value for required argument 'CrossAccountAccess'")
	}
	if args.FeaturesEnableds == nil {
		return nil, errors.New("invalid value for required argument 'FeaturesEnableds'")
	}
	if args.Identifier == nil {
		return nil, errors.New("invalid value for required argument 'Identifier'")
	}
	if args.ReportName == nil {
		return nil, errors.New("invalid value for required argument 'ReportName'")
	}
	if args.S3Bucket == nil {
		return nil, errors.New("invalid value for required argument 'S3Bucket'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AwsCC
	err := ctx.RegisterResource("harness:PlatformConnector/awsCC:AwsCC", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAwsCC gets an existing AwsCC resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAwsCC(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AwsCCState, opts ...pulumi.ResourceOption) (*AwsCC, error) {
	var resource AwsCC
	err := ctx.ReadResource("harness:PlatformConnector/awsCC:AwsCC", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AwsCC resources.
type awsCCState struct {
	// The AWS account id.
	AccountId *string `pulumi:"accountId"`
	// Harness uses the secure cross-account role to access your AWS account. The role includes a restricted policy to access the cost and usage reports and resources for the sole purpose of cost analysis and cost optimization.
	CrossAccountAccess *AwsCCCrossAccountAccess `pulumi:"crossAccountAccess"`
	// Description of the resource.
	Description *string `pulumi:"description"`
	// The features enabled for the connector. Valid values are BILLING, OPTIMIZATION, VISIBILITY.
	FeaturesEnableds []string `pulumi:"featuresEnableds"`
	// Unique identifier of the resource.
	Identifier *string `pulumi:"identifier"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId *string `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
	// The cost and usage report name. Provided in the delivery options when the template is opened in the AWS console.
	ReportName *string `pulumi:"reportName"`
	// The name of s3 bucket.
	S3Bucket *string `pulumi:"s3Bucket"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags []string `pulumi:"tags"`
}

type AwsCCState struct {
	// The AWS account id.
	AccountId pulumi.StringPtrInput
	// Harness uses the secure cross-account role to access your AWS account. The role includes a restricted policy to access the cost and usage reports and resources for the sole purpose of cost analysis and cost optimization.
	CrossAccountAccess AwsCCCrossAccountAccessPtrInput
	// Description of the resource.
	Description pulumi.StringPtrInput
	// The features enabled for the connector. Valid values are BILLING, OPTIMIZATION, VISIBILITY.
	FeaturesEnableds pulumi.StringArrayInput
	// Unique identifier of the resource.
	Identifier pulumi.StringPtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrInput
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrInput
	// The cost and usage report name. Provided in the delivery options when the template is opened in the AWS console.
	ReportName pulumi.StringPtrInput
	// The name of s3 bucket.
	S3Bucket pulumi.StringPtrInput
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags pulumi.StringArrayInput
}

func (AwsCCState) ElementType() reflect.Type {
	return reflect.TypeOf((*awsCCState)(nil)).Elem()
}

type awsCCArgs struct {
	// The AWS account id.
	AccountId string `pulumi:"accountId"`
	// Harness uses the secure cross-account role to access your AWS account. The role includes a restricted policy to access the cost and usage reports and resources for the sole purpose of cost analysis and cost optimization.
	CrossAccountAccess AwsCCCrossAccountAccess `pulumi:"crossAccountAccess"`
	// Description of the resource.
	Description *string `pulumi:"description"`
	// The features enabled for the connector. Valid values are BILLING, OPTIMIZATION, VISIBILITY.
	FeaturesEnableds []string `pulumi:"featuresEnableds"`
	// Unique identifier of the resource.
	Identifier string `pulumi:"identifier"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId *string `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
	// The cost and usage report name. Provided in the delivery options when the template is opened in the AWS console.
	ReportName string `pulumi:"reportName"`
	// The name of s3 bucket.
	S3Bucket string `pulumi:"s3Bucket"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a AwsCC resource.
type AwsCCArgs struct {
	// The AWS account id.
	AccountId pulumi.StringInput
	// Harness uses the secure cross-account role to access your AWS account. The role includes a restricted policy to access the cost and usage reports and resources for the sole purpose of cost analysis and cost optimization.
	CrossAccountAccess AwsCCCrossAccountAccessInput
	// Description of the resource.
	Description pulumi.StringPtrInput
	// The features enabled for the connector. Valid values are BILLING, OPTIMIZATION, VISIBILITY.
	FeaturesEnableds pulumi.StringArrayInput
	// Unique identifier of the resource.
	Identifier pulumi.StringInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrInput
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrInput
	// The cost and usage report name. Provided in the delivery options when the template is opened in the AWS console.
	ReportName pulumi.StringInput
	// The name of s3 bucket.
	S3Bucket pulumi.StringInput
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags pulumi.StringArrayInput
}

func (AwsCCArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*awsCCArgs)(nil)).Elem()
}

type AwsCCInput interface {
	pulumi.Input

	ToAwsCCOutput() AwsCCOutput
	ToAwsCCOutputWithContext(ctx context.Context) AwsCCOutput
}

func (*AwsCC) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsCC)(nil)).Elem()
}

func (i *AwsCC) ToAwsCCOutput() AwsCCOutput {
	return i.ToAwsCCOutputWithContext(context.Background())
}

func (i *AwsCC) ToAwsCCOutputWithContext(ctx context.Context) AwsCCOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsCCOutput)
}

// AwsCCArrayInput is an input type that accepts AwsCCArray and AwsCCArrayOutput values.
// You can construct a concrete instance of `AwsCCArrayInput` via:
//
//	AwsCCArray{ AwsCCArgs{...} }
type AwsCCArrayInput interface {
	pulumi.Input

	ToAwsCCArrayOutput() AwsCCArrayOutput
	ToAwsCCArrayOutputWithContext(context.Context) AwsCCArrayOutput
}

type AwsCCArray []AwsCCInput

func (AwsCCArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AwsCC)(nil)).Elem()
}

func (i AwsCCArray) ToAwsCCArrayOutput() AwsCCArrayOutput {
	return i.ToAwsCCArrayOutputWithContext(context.Background())
}

func (i AwsCCArray) ToAwsCCArrayOutputWithContext(ctx context.Context) AwsCCArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsCCArrayOutput)
}

// AwsCCMapInput is an input type that accepts AwsCCMap and AwsCCMapOutput values.
// You can construct a concrete instance of `AwsCCMapInput` via:
//
//	AwsCCMap{ "key": AwsCCArgs{...} }
type AwsCCMapInput interface {
	pulumi.Input

	ToAwsCCMapOutput() AwsCCMapOutput
	ToAwsCCMapOutputWithContext(context.Context) AwsCCMapOutput
}

type AwsCCMap map[string]AwsCCInput

func (AwsCCMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AwsCC)(nil)).Elem()
}

func (i AwsCCMap) ToAwsCCMapOutput() AwsCCMapOutput {
	return i.ToAwsCCMapOutputWithContext(context.Background())
}

func (i AwsCCMap) ToAwsCCMapOutputWithContext(ctx context.Context) AwsCCMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsCCMapOutput)
}

type AwsCCOutput struct{ *pulumi.OutputState }

func (AwsCCOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsCC)(nil)).Elem()
}

func (o AwsCCOutput) ToAwsCCOutput() AwsCCOutput {
	return o
}

func (o AwsCCOutput) ToAwsCCOutputWithContext(ctx context.Context) AwsCCOutput {
	return o
}

// The AWS account id.
func (o AwsCCOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsCC) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// Harness uses the secure cross-account role to access your AWS account. The role includes a restricted policy to access the cost and usage reports and resources for the sole purpose of cost analysis and cost optimization.
func (o AwsCCOutput) CrossAccountAccess() AwsCCCrossAccountAccessOutput {
	return o.ApplyT(func(v *AwsCC) AwsCCCrossAccountAccessOutput { return v.CrossAccountAccess }).(AwsCCCrossAccountAccessOutput)
}

// Description of the resource.
func (o AwsCCOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AwsCC) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The features enabled for the connector. Valid values are BILLING, OPTIMIZATION, VISIBILITY.
func (o AwsCCOutput) FeaturesEnableds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AwsCC) pulumi.StringArrayOutput { return v.FeaturesEnableds }).(pulumi.StringArrayOutput)
}

// Unique identifier of the resource.
func (o AwsCCOutput) Identifier() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsCC) pulumi.StringOutput { return v.Identifier }).(pulumi.StringOutput)
}

// Name of the resource.
func (o AwsCCOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsCC) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Unique identifier of the organization.
func (o AwsCCOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AwsCC) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Unique identifier of the project.
func (o AwsCCOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AwsCC) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// The cost and usage report name. Provided in the delivery options when the template is opened in the AWS console.
func (o AwsCCOutput) ReportName() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsCC) pulumi.StringOutput { return v.ReportName }).(pulumi.StringOutput)
}

// The name of s3 bucket.
func (o AwsCCOutput) S3Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsCC) pulumi.StringOutput { return v.S3Bucket }).(pulumi.StringOutput)
}

// Tags to associate with the resource. Tags should be in the form `name:value`.
func (o AwsCCOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AwsCC) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

type AwsCCArrayOutput struct{ *pulumi.OutputState }

func (AwsCCArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AwsCC)(nil)).Elem()
}

func (o AwsCCArrayOutput) ToAwsCCArrayOutput() AwsCCArrayOutput {
	return o
}

func (o AwsCCArrayOutput) ToAwsCCArrayOutputWithContext(ctx context.Context) AwsCCArrayOutput {
	return o
}

func (o AwsCCArrayOutput) Index(i pulumi.IntInput) AwsCCOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AwsCC {
		return vs[0].([]*AwsCC)[vs[1].(int)]
	}).(AwsCCOutput)
}

type AwsCCMapOutput struct{ *pulumi.OutputState }

func (AwsCCMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AwsCC)(nil)).Elem()
}

func (o AwsCCMapOutput) ToAwsCCMapOutput() AwsCCMapOutput {
	return o
}

func (o AwsCCMapOutput) ToAwsCCMapOutputWithContext(ctx context.Context) AwsCCMapOutput {
	return o
}

func (o AwsCCMapOutput) MapIndex(k pulumi.StringInput) AwsCCOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AwsCC {
		return vs[0].(map[string]*AwsCC)[vs[1].(string)]
	}).(AwsCCOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AwsCCInput)(nil)).Elem(), &AwsCC{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsCCArrayInput)(nil)).Elem(), AwsCCArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsCCMapInput)(nil)).Elem(), AwsCCMap{})
	pulumi.RegisterOutputType(AwsCCOutput{})
	pulumi.RegisterOutputType(AwsCCArrayOutput{})
	pulumi.RegisterOutputType(AwsCCMapOutput{})
}

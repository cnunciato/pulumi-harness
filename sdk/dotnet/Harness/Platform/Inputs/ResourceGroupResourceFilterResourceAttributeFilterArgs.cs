// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Harness.Platform.Inputs
{

    public sealed class ResourceGroupResourceFilterResourceAttributeFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("attributeName")]
        public Input<string>? AttributeName { get; set; }

        [Input("attributeValues")]
        private InputList<string>? _attributeValues;
        public InputList<string> AttributeValues
        {
            get => _attributeValues ?? (_attributeValues = new InputList<string>());
            set => _attributeValues = value;
        }

        public ResourceGroupResourceFilterResourceAttributeFilterArgs()
        {
        }
        public static new ResourceGroupResourceFilterResourceAttributeFilterArgs Empty => new ResourceGroupResourceFilterResourceAttributeFilterArgs();
    }
}
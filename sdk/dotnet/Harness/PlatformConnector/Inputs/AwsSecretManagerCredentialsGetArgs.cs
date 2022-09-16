// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Harness.PlatformConnector.Inputs
{

    public sealed class AwsSecretManagerCredentialsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Connect using STS assume role.
        /// </summary>
        [Input("assumeRole")]
        public Input<Inputs.AwsSecretManagerCredentialsAssumeRoleGetArgs>? AssumeRole { get; set; }

        /// <summary>
        /// Inherit the credentials from from the delegate.
        /// </summary>
        [Input("inheritFromDelegate")]
        public Input<bool>? InheritFromDelegate { get; set; }

        /// <summary>
        /// Specify the AWS key and secret used for authenticating.
        /// </summary>
        [Input("manual")]
        public Input<Inputs.AwsSecretManagerCredentialsManualGetArgs>? Manual { get; set; }

        public AwsSecretManagerCredentialsGetArgs()
        {
        }
        public static new AwsSecretManagerCredentialsGetArgs Empty => new AwsSecretManagerCredentialsGetArgs();
    }
}
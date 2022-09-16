// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Harness.Cloudprovider.Inputs
{

    public sealed class KubernetesAuthenticationServiceAccountGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("caCertificateSecretName")]
        public Input<string>? CaCertificateSecretName { get; set; }

        [Input("masterUrl", required: true)]
        public Input<string> MasterUrl { get; set; } = null!;

        [Input("serviceAccountTokenSecretName", required: true)]
        public Input<string> ServiceAccountTokenSecretName { get; set; } = null!;

        public KubernetesAuthenticationServiceAccountGetArgs()
        {
        }
        public static new KubernetesAuthenticationServiceAccountGetArgs Empty => new KubernetesAuthenticationServiceAccountGetArgs();
    }
}
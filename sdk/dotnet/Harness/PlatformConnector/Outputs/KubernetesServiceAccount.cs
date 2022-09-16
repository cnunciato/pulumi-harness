// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Harness.PlatformConnector.Outputs
{

    [OutputType]
    public sealed class KubernetesServiceAccount
    {
        /// <summary>
        /// The URL of the Kubernetes cluster.
        /// </summary>
        public readonly string MasterUrl;
        /// <summary>
        /// Reference to the secret containing the service account token for the connector.
        /// </summary>
        public readonly string ServiceAccountTokenRef;

        [OutputConstructor]
        private KubernetesServiceAccount(
            string masterUrl,

            string serviceAccountTokenRef)
        {
            MasterUrl = masterUrl;
            ServiceAccountTokenRef = serviceAccountTokenRef;
        }
    }
}
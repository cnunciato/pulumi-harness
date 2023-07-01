// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Harness.Platform.Outputs
{

    [OutputType]
    public sealed class GetGitopsApplicationsApplicationResult
    {
        /// <summary>
        /// Metadata corresponding to the resources. This includes all the objects a user must create.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGitopsApplicationsApplicationMetadataResult> Metadatas;
        /// <summary>
        /// Specifications of the GitOps application. This includes the repository URL, application definition, source, destination and sync policy.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGitopsApplicationsApplicationSpecResult> Specs;

        [OutputConstructor]
        private GetGitopsApplicationsApplicationResult(
            ImmutableArray<Outputs.GetGitopsApplicationsApplicationMetadataResult> metadatas,

            ImmutableArray<Outputs.GetGitopsApplicationsApplicationSpecResult> specs)
        {
            Metadatas = metadatas;
            Specs = specs;
        }
    }
}
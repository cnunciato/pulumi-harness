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

    public sealed class GitCredentialsSshArgs : global::Pulumi.ResourceArgs
    {
        [Input("sshKeyRef", required: true)]
        public Input<string> SshKeyRef { get; set; } = null!;

        public GitCredentialsSshArgs()
        {
        }
        public static new GitCredentialsSshArgs Empty => new GitCredentialsSshArgs();
    }
}
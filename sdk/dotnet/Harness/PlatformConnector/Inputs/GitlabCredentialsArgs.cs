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

    public sealed class GitlabCredentialsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authenticate using Username and password over http(s) for the connection.
        /// </summary>
        [Input("http")]
        public Input<Inputs.GitlabCredentialsHttpArgs>? Http { get; set; }

        /// <summary>
        /// Authenticate using SSH for the connection.
        /// </summary>
        [Input("ssh")]
        public Input<Inputs.GitlabCredentialsSshArgs>? Ssh { get; set; }

        public GitlabCredentialsArgs()
        {
        }
        public static new GitlabCredentialsArgs Empty => new GitlabCredentialsArgs();
    }
}
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
    public sealed class GithubCredentials
    {
        /// <summary>
        /// Authenticate using Username and password over http(s) for the connection.
        /// </summary>
        public readonly Outputs.GithubCredentialsHttp? Http;
        /// <summary>
        /// Authenticate using SSH for the connection.
        /// </summary>
        public readonly Outputs.GithubCredentialsSsh? Ssh;

        [OutputConstructor]
        private GithubCredentials(
            Outputs.GithubCredentialsHttp? http,

            Outputs.GithubCredentialsSsh? ssh)
        {
            Http = http;
            Ssh = ssh;
        }
    }
}
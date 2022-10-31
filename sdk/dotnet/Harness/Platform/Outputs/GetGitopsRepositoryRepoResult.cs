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
    public sealed class GetGitopsRepositoryRepoResult
    {
        public readonly string ConnectionType;
        public readonly bool? EnableLfs;
        public readonly bool? EnableOci;
        public readonly string? GithubAppEnterpriseBaseUrl;
        public readonly string? GithubAppId;
        public readonly string? GithubAppInstallationId;
        public readonly string? GithubAppPrivateKey;
        public readonly bool? InheritedCreds;
        public readonly bool? Insecure;
        public readonly bool? InsecureIgnoreHostKey;
        public readonly string? Name;
        public readonly string? Password;
        public readonly string Project;
        public readonly string? Proxy;
        public readonly string Repo;
        public readonly string? SshPrivateKey;
        public readonly string? TlsClientCertData;
        public readonly string? TlsClientCertKey;
        public readonly string Type_;
        public readonly string? Username;

        [OutputConstructor]
        private GetGitopsRepositoryRepoResult(
            string connectionType,

            bool? enableLfs,

            bool? enableOci,

            string? githubAppEnterpriseBaseUrl,

            string? githubAppId,

            string? githubAppInstallationId,

            string? githubAppPrivateKey,

            bool? inheritedCreds,

            bool? insecure,

            bool? insecureIgnoreHostKey,

            string? name,

            string? password,

            string project,

            string? proxy,

            string repo,

            string? sshPrivateKey,

            string? tlsClientCertData,

            string? tlsClientCertKey,

            string type_,

            string? username)
        {
            ConnectionType = connectionType;
            EnableLfs = enableLfs;
            EnableOci = enableOci;
            GithubAppEnterpriseBaseUrl = githubAppEnterpriseBaseUrl;
            GithubAppId = githubAppId;
            GithubAppInstallationId = githubAppInstallationId;
            GithubAppPrivateKey = githubAppPrivateKey;
            InheritedCreds = inheritedCreds;
            Insecure = insecure;
            InsecureIgnoreHostKey = insecureIgnoreHostKey;
            Name = name;
            Password = password;
            Project = project;
            Proxy = proxy;
            Repo = repo;
            SshPrivateKey = sshPrivateKey;
            TlsClientCertData = tlsClientCertData;
            TlsClientCertKey = tlsClientCertKey;
            Type_ = type_;
            Username = username;
        }
    }
}
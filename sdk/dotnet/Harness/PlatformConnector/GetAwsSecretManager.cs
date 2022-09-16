// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Harness.PlatformConnector
{
    public static class GetAwsSecretManager
    {
        /// <summary>
        /// Datasource for looking up an AWS Secret Manager connector.
        /// </summary>
        public static Task<GetAwsSecretManagerResult> InvokeAsync(GetAwsSecretManagerArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAwsSecretManagerResult>("harness:PlatformConnector/getAwsSecretManager:getAwsSecretManager", args ?? new GetAwsSecretManagerArgs(), options.WithDefaults());

        /// <summary>
        /// Datasource for looking up an AWS Secret Manager connector.
        /// </summary>
        public static Output<GetAwsSecretManagerResult> Invoke(GetAwsSecretManagerInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAwsSecretManagerResult>("harness:PlatformConnector/getAwsSecretManager:getAwsSecretManager", args ?? new GetAwsSecretManagerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAwsSecretManagerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier of the resource.
        /// </summary>
        [Input("identifier")]
        public string? Identifier { get; set; }

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Unique identifier of the organization.
        /// </summary>
        [Input("orgId")]
        public string? OrgId { get; set; }

        /// <summary>
        /// Unique identifier of the project.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        public GetAwsSecretManagerArgs()
        {
        }
        public static new GetAwsSecretManagerArgs Empty => new GetAwsSecretManagerArgs();
    }

    public sealed class GetAwsSecretManagerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier of the resource.
        /// </summary>
        [Input("identifier")]
        public Input<string>? Identifier { get; set; }

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Unique identifier of the organization.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// Unique identifier of the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public GetAwsSecretManagerInvokeArgs()
        {
        }
        public static new GetAwsSecretManagerInvokeArgs Empty => new GetAwsSecretManagerInvokeArgs();
    }


    [OutputType]
    public sealed class GetAwsSecretManagerResult
    {
        /// <summary>
        /// The credentials to use for connecting to aws.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAwsSecretManagerCredentialResult> Credentials;
        /// <summary>
        /// Connect using only the delegates which have these tags.
        /// </summary>
        public readonly ImmutableArray<string> DelegateSelectors;
        /// <summary>
        /// Description of the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Unique identifier of the resource.
        /// </summary>
        public readonly string? Identifier;
        /// <summary>
        /// Name of the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Unique identifier of the organization.
        /// </summary>
        public readonly string? OrgId;
        /// <summary>
        /// Unique identifier of the project.
        /// </summary>
        public readonly string? ProjectId;
        /// <summary>
        /// The AWS region where the AWS Secret Manager is.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// A prefix to be added to all secrets.
        /// </summary>
        public readonly string SecretNamePrefix;
        /// <summary>
        /// Tags to associate with the resource. Tags should be in the form `name:value`.
        /// </summary>
        public readonly ImmutableArray<string> Tags;

        [OutputConstructor]
        private GetAwsSecretManagerResult(
            ImmutableArray<Outputs.GetAwsSecretManagerCredentialResult> credentials,

            ImmutableArray<string> delegateSelectors,

            string description,

            string id,

            string? identifier,

            string? name,

            string? orgId,

            string? projectId,

            string region,

            string secretNamePrefix,

            ImmutableArray<string> tags)
        {
            Credentials = credentials;
            DelegateSelectors = delegateSelectors;
            Description = description;
            Id = id;
            Identifier = identifier;
            Name = name;
            OrgId = orgId;
            ProjectId = projectId;
            Region = region;
            SecretNamePrefix = secretNamePrefix;
            Tags = tags;
        }
    }
}
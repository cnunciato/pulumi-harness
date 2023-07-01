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

    public sealed class GetTemplateGitDetailsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the branch.
        /// </summary>
        [Input("branchName")]
        public string? BranchName { get; set; }

        /// <summary>
        /// File path of the Entity in the repository.
        /// </summary>
        [Input("filePath")]
        public string? FilePath { get; set; }

        /// <summary>
        /// File url of the Entity in the repository.
        /// </summary>
        [Input("fileUrl")]
        public string? FileUrl { get; set; }

        /// <summary>
        /// Last commit identifier (for Git Repositories other than Github). To be provided only when updating Pipeline.
        /// </summary>
        [Input("lastCommitId", required: true)]
        public string LastCommitId { get; set; } = null!;

        /// <summary>
        /// Last object identifier (for Github). To be provided only when updating Pipeline.
        /// </summary>
        [Input("lastObjectId", required: true)]
        public string LastObjectId { get; set; } = null!;

        /// <summary>
        /// Name of the repository.
        /// </summary>
        [Input("repoName")]
        public string? RepoName { get; set; }

        /// <summary>
        /// Repo url of the Entity in the repository.
        /// </summary>
        [Input("repoUrl")]
        public string? RepoUrl { get; set; }

        public GetTemplateGitDetailsArgs()
        {
        }
        public static new GetTemplateGitDetailsArgs Empty => new GetTemplateGitDetailsArgs();
    }
}
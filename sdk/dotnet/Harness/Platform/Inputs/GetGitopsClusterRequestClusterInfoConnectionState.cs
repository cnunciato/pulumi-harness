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

    public sealed class GetGitopsClusterRequestClusterInfoConnectionStateArgs : global::Pulumi.InvokeArgs
    {
        [Input("attemptedAts")]
        private List<Inputs.GetGitopsClusterRequestClusterInfoConnectionStateAttemptedAtArgs>? _attemptedAts;
        public List<Inputs.GetGitopsClusterRequestClusterInfoConnectionStateAttemptedAtArgs> AttemptedAts
        {
            get => _attemptedAts ?? (_attemptedAts = new List<Inputs.GetGitopsClusterRequestClusterInfoConnectionStateAttemptedAtArgs>());
            set => _attemptedAts = value;
        }

        [Input("message")]
        public string? Message { get; set; }

        [Input("status")]
        public string? Status { get; set; }

        public GetGitopsClusterRequestClusterInfoConnectionStateArgs()
        {
        }
        public static new GetGitopsClusterRequestClusterInfoConnectionStateArgs Empty => new GetGitopsClusterRequestClusterInfoConnectionStateArgs();
    }
}
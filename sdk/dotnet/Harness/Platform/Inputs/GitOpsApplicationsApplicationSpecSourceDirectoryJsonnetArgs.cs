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

    public sealed class GitOpsApplicationsApplicationSpecSourceDirectoryJsonnetArgs : global::Pulumi.ResourceArgs
    {
        [Input("extVars")]
        private InputList<Inputs.GitOpsApplicationsApplicationSpecSourceDirectoryJsonnetExtVarArgs>? _extVars;
        public InputList<Inputs.GitOpsApplicationsApplicationSpecSourceDirectoryJsonnetExtVarArgs> ExtVars
        {
            get => _extVars ?? (_extVars = new InputList<Inputs.GitOpsApplicationsApplicationSpecSourceDirectoryJsonnetExtVarArgs>());
            set => _extVars = value;
        }

        [Input("libs")]
        private InputList<string>? _libs;
        public InputList<string> Libs
        {
            get => _libs ?? (_libs = new InputList<string>());
            set => _libs = value;
        }

        [Input("tlas")]
        private InputList<Inputs.GitOpsApplicationsApplicationSpecSourceDirectoryJsonnetTlaArgs>? _tlas;
        public InputList<Inputs.GitOpsApplicationsApplicationSpecSourceDirectoryJsonnetTlaArgs> Tlas
        {
            get => _tlas ?? (_tlas = new InputList<Inputs.GitOpsApplicationsApplicationSpecSourceDirectoryJsonnetTlaArgs>());
            set => _tlas = value;
        }

        public GitOpsApplicationsApplicationSpecSourceDirectoryJsonnetArgs()
        {
        }
        public static new GitOpsApplicationsApplicationSpecSourceDirectoryJsonnetArgs Empty => new GitOpsApplicationsApplicationSpecSourceDirectoryJsonnetArgs();
    }
}
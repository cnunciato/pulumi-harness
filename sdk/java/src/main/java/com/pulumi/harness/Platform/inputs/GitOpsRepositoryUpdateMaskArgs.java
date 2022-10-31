// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GitOpsRepositoryUpdateMaskArgs extends com.pulumi.resources.ResourceArgs {

    public static final GitOpsRepositoryUpdateMaskArgs Empty = new GitOpsRepositoryUpdateMaskArgs();

    @Import(name="paths")
    private @Nullable Output<List<String>> paths;

    public Optional<Output<List<String>>> paths() {
        return Optional.ofNullable(this.paths);
    }

    private GitOpsRepositoryUpdateMaskArgs() {}

    private GitOpsRepositoryUpdateMaskArgs(GitOpsRepositoryUpdateMaskArgs $) {
        this.paths = $.paths;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GitOpsRepositoryUpdateMaskArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GitOpsRepositoryUpdateMaskArgs $;

        public Builder() {
            $ = new GitOpsRepositoryUpdateMaskArgs();
        }

        public Builder(GitOpsRepositoryUpdateMaskArgs defaults) {
            $ = new GitOpsRepositoryUpdateMaskArgs(Objects.requireNonNull(defaults));
        }

        public Builder paths(@Nullable Output<List<String>> paths) {
            $.paths = paths;
            return this;
        }

        public Builder paths(List<String> paths) {
            return paths(Output.of(paths));
        }

        public Builder paths(String... paths) {
            return paths(List.of(paths));
        }

        public GitOpsRepositoryUpdateMaskArgs build() {
            return $;
        }
    }

}
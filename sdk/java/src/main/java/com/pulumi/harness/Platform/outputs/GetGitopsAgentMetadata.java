// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetGitopsAgentMetadata {
    private @Nullable Boolean highAvailability;
    private String namespace;

    private GetGitopsAgentMetadata() {}
    public Optional<Boolean> highAvailability() {
        return Optional.ofNullable(this.highAvailability);
    }
    public String namespace() {
        return this.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGitopsAgentMetadata defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean highAvailability;
        private String namespace;
        public Builder() {}
        public Builder(GetGitopsAgentMetadata defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.highAvailability = defaults.highAvailability;
    	      this.namespace = defaults.namespace;
        }

        @CustomType.Setter
        public Builder highAvailability(@Nullable Boolean highAvailability) {
            this.highAvailability = highAvailability;
            return this;
        }
        @CustomType.Setter
        public Builder namespace(String namespace) {
            this.namespace = Objects.requireNonNull(namespace);
            return this;
        }
        public GetGitopsAgentMetadata build() {
            final var o = new GetGitopsAgentMetadata();
            o.highAvailability = highAvailability;
            o.namespace = namespace;
            return o;
        }
    }
}
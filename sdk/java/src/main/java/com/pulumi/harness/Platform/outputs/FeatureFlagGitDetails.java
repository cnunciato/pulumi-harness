// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class FeatureFlagGitDetails {
    /**
     * @return The commit message to use as part of a gitsync operation
     * 
     */
    private String commitMsg;

    private FeatureFlagGitDetails() {}
    /**
     * @return The commit message to use as part of a gitsync operation
     * 
     */
    public String commitMsg() {
        return this.commitMsg;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FeatureFlagGitDetails defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String commitMsg;
        public Builder() {}
        public Builder(FeatureFlagGitDetails defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.commitMsg = defaults.commitMsg;
        }

        @CustomType.Setter
        public Builder commitMsg(String commitMsg) {
            this.commitMsg = Objects.requireNonNull(commitMsg);
            return this;
        }
        public FeatureFlagGitDetails build() {
            final var o = new FeatureFlagGitDetails();
            o.commitMsg = commitMsg;
            return o;
        }
    }
}
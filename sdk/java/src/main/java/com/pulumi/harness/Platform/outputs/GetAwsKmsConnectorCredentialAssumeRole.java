// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetAwsKmsConnectorCredentialAssumeRole {
    private Integer duration;
    private String externalId;
    private String roleArn;

    private GetAwsKmsConnectorCredentialAssumeRole() {}
    public Integer duration() {
        return this.duration;
    }
    public String externalId() {
        return this.externalId;
    }
    public String roleArn() {
        return this.roleArn;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAwsKmsConnectorCredentialAssumeRole defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer duration;
        private String externalId;
        private String roleArn;
        public Builder() {}
        public Builder(GetAwsKmsConnectorCredentialAssumeRole defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.duration = defaults.duration;
    	      this.externalId = defaults.externalId;
    	      this.roleArn = defaults.roleArn;
        }

        @CustomType.Setter
        public Builder duration(Integer duration) {
            this.duration = Objects.requireNonNull(duration);
            return this;
        }
        @CustomType.Setter
        public Builder externalId(String externalId) {
            this.externalId = Objects.requireNonNull(externalId);
            return this;
        }
        @CustomType.Setter
        public Builder roleArn(String roleArn) {
            this.roleArn = Objects.requireNonNull(roleArn);
            return this;
        }
        public GetAwsKmsConnectorCredentialAssumeRole build() {
            final var o = new GetAwsKmsConnectorCredentialAssumeRole();
            o.duration = duration;
            o.externalId = externalId;
            o.roleArn = roleArn;
            return o;
        }
    }
}
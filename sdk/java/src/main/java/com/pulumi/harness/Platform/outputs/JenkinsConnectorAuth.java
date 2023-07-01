// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.harness.platform.outputs.JenkinsConnectorAuthJenkinsBearerToken;
import com.pulumi.harness.platform.outputs.JenkinsConnectorAuthJenkinsUserNamePassword;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class JenkinsConnectorAuth {
    /**
     * @return Authenticate to App Dynamics using bearer token.
     * 
     */
    private @Nullable JenkinsConnectorAuthJenkinsBearerToken jenkinsBearerToken;
    /**
     * @return Authenticate to App Dynamics using user name and password.
     * 
     */
    private @Nullable JenkinsConnectorAuthJenkinsUserNamePassword jenkinsUserNamePassword;
    /**
     * @return Can be one of UsernamePassword, Anonymous, Bearer Token(HTTP Header)
     * 
     */
    private String type;

    private JenkinsConnectorAuth() {}
    /**
     * @return Authenticate to App Dynamics using bearer token.
     * 
     */
    public Optional<JenkinsConnectorAuthJenkinsBearerToken> jenkinsBearerToken() {
        return Optional.ofNullable(this.jenkinsBearerToken);
    }
    /**
     * @return Authenticate to App Dynamics using user name and password.
     * 
     */
    public Optional<JenkinsConnectorAuthJenkinsUserNamePassword> jenkinsUserNamePassword() {
        return Optional.ofNullable(this.jenkinsUserNamePassword);
    }
    /**
     * @return Can be one of UsernamePassword, Anonymous, Bearer Token(HTTP Header)
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(JenkinsConnectorAuth defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable JenkinsConnectorAuthJenkinsBearerToken jenkinsBearerToken;
        private @Nullable JenkinsConnectorAuthJenkinsUserNamePassword jenkinsUserNamePassword;
        private String type;
        public Builder() {}
        public Builder(JenkinsConnectorAuth defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.jenkinsBearerToken = defaults.jenkinsBearerToken;
    	      this.jenkinsUserNamePassword = defaults.jenkinsUserNamePassword;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder jenkinsBearerToken(@Nullable JenkinsConnectorAuthJenkinsBearerToken jenkinsBearerToken) {
            this.jenkinsBearerToken = jenkinsBearerToken;
            return this;
        }
        @CustomType.Setter
        public Builder jenkinsUserNamePassword(@Nullable JenkinsConnectorAuthJenkinsUserNamePassword jenkinsUserNamePassword) {
            this.jenkinsUserNamePassword = jenkinsUserNamePassword;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public JenkinsConnectorAuth build() {
            final var o = new JenkinsConnectorAuth();
            o.jenkinsBearerToken = jenkinsBearerToken;
            o.jenkinsUserNamePassword = jenkinsUserNamePassword;
            o.type = type;
            return o;
        }
    }
}
// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FeatureFlagApiKeyArgs extends com.pulumi.resources.ResourceArgs {

    public static final FeatureFlagApiKeyArgs Empty = new FeatureFlagApiKeyArgs();

    /**
     * Description of the SDK API Key
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the SDK API Key
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Environment Identifier
     * 
     */
    @Import(name="envId", required=true)
    private Output<String> envId;

    /**
     * @return Environment Identifier
     * 
     */
    public Output<String> envId() {
        return this.envId;
    }

    /**
     * Expiration datetime of the SDK API Key
     * 
     */
    @Import(name="expiredAt")
    private @Nullable Output<Integer> expiredAt;

    /**
     * @return Expiration datetime of the SDK API Key
     * 
     */
    public Optional<Output<Integer>> expiredAt() {
        return Optional.ofNullable(this.expiredAt);
    }

    /**
     * Identifier of the SDK API Key
     * 
     */
    @Import(name="identifier", required=true)
    private Output<String> identifier;

    /**
     * @return Identifier of the SDK API Key
     * 
     */
    public Output<String> identifier() {
        return this.identifier;
    }

    /**
     * Name of the SDK API Key
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the SDK API Key
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Organization Identifier
     * 
     */
    @Import(name="orgId", required=true)
    private Output<String> orgId;

    /**
     * @return Organization Identifier
     * 
     */
    public Output<String> orgId() {
        return this.orgId;
    }

    /**
     * Project Identifier
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return Project Identifier
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * Type of SDK. Valid values are `Server` or `Client`.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Type of SDK. Valid values are `Server` or `Client`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private FeatureFlagApiKeyArgs() {}

    private FeatureFlagApiKeyArgs(FeatureFlagApiKeyArgs $) {
        this.description = $.description;
        this.envId = $.envId;
        this.expiredAt = $.expiredAt;
        this.identifier = $.identifier;
        this.name = $.name;
        this.orgId = $.orgId;
        this.projectId = $.projectId;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FeatureFlagApiKeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FeatureFlagApiKeyArgs $;

        public Builder() {
            $ = new FeatureFlagApiKeyArgs();
        }

        public Builder(FeatureFlagApiKeyArgs defaults) {
            $ = new FeatureFlagApiKeyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Description of the SDK API Key
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the SDK API Key
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param envId Environment Identifier
         * 
         * @return builder
         * 
         */
        public Builder envId(Output<String> envId) {
            $.envId = envId;
            return this;
        }

        /**
         * @param envId Environment Identifier
         * 
         * @return builder
         * 
         */
        public Builder envId(String envId) {
            return envId(Output.of(envId));
        }

        /**
         * @param expiredAt Expiration datetime of the SDK API Key
         * 
         * @return builder
         * 
         */
        public Builder expiredAt(@Nullable Output<Integer> expiredAt) {
            $.expiredAt = expiredAt;
            return this;
        }

        /**
         * @param expiredAt Expiration datetime of the SDK API Key
         * 
         * @return builder
         * 
         */
        public Builder expiredAt(Integer expiredAt) {
            return expiredAt(Output.of(expiredAt));
        }

        /**
         * @param identifier Identifier of the SDK API Key
         * 
         * @return builder
         * 
         */
        public Builder identifier(Output<String> identifier) {
            $.identifier = identifier;
            return this;
        }

        /**
         * @param identifier Identifier of the SDK API Key
         * 
         * @return builder
         * 
         */
        public Builder identifier(String identifier) {
            return identifier(Output.of(identifier));
        }

        /**
         * @param name Name of the SDK API Key
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the SDK API Key
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param orgId Organization Identifier
         * 
         * @return builder
         * 
         */
        public Builder orgId(Output<String> orgId) {
            $.orgId = orgId;
            return this;
        }

        /**
         * @param orgId Organization Identifier
         * 
         * @return builder
         * 
         */
        public Builder orgId(String orgId) {
            return orgId(Output.of(orgId));
        }

        /**
         * @param projectId Project Identifier
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId Project Identifier
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param type Type of SDK. Valid values are `Server` or `Client`.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of SDK. Valid values are `Server` or `Client`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public FeatureFlagApiKeyArgs build() {
            $.envId = Objects.requireNonNull($.envId, "expected parameter 'envId' to be non-null");
            $.identifier = Objects.requireNonNull($.identifier, "expected parameter 'identifier' to be non-null");
            $.orgId = Objects.requireNonNull($.orgId, "expected parameter 'orgId' to be non-null");
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.harness.Utilities;
import com.pulumi.harness.platform.VaultConnectorArgs;
import com.pulumi.harness.platform.inputs.VaultConnectorState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Resource for creating a HashiCorp Vault Secret Manager connector.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.harness.platform.VaultConnector;
 * import com.pulumi.harness.platform.VaultConnectorArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var awsAuth = new VaultConnector(&#34;awsAuth&#34;, VaultConnectorArgs.builder()        
 *             .accessType(&#34;AWS_IAM&#34;)
 *             .awsRegion(&#34;aws_region&#34;)
 *             .basePath(&#34;base_path&#34;)
 *             .default_(false)
 *             .delegateSelectors(&#34;harness-delegate&#34;)
 *             .description(&#34;test&#34;)
 *             .identifier(&#34;identifier&#34;)
 *             .readOnly(true)
 *             .renewalIntervalMinutes(60)
 *             .secretEngineManuallyConfigured(true)
 *             .secretEngineName(&#34;secret_engine_name&#34;)
 *             .secretEngineVersion(2)
 *             .tags(&#34;foo:bar&#34;)
 *             .useAwsIam(true)
 *             .useK8sAuth(false)
 *             .useVaultAgent(false)
 *             .vaultAwsIamRole(&#34;vault_aws_iam_role&#34;)
 *             .vaultUrl(&#34;https://vault_url.com&#34;)
 *             .xvaultAwsIamServerId(String.format(&#34;account.%s&#34;, harness_platform_secret_text.test().id()))
 *             .build());
 * 
 *         var appRole = new VaultConnector(&#34;appRole&#34;, VaultConnectorArgs.builder()        
 *             .accessType(&#34;APP_ROLE&#34;)
 *             .appRoleId(&#34;app_role_id&#34;)
 *             .basePath(&#34;base_path&#34;)
 *             .default_(false)
 *             .delegateSelectors(&#34;harness-delegate&#34;)
 *             .description(&#34;test&#34;)
 *             .identifier(&#34;identifier&#34;)
 *             .readOnly(true)
 *             .renewAppRoleToken(true)
 *             .renewalIntervalMinutes(60)
 *             .secretEngineManuallyConfigured(true)
 *             .secretEngineName(&#34;secret_engine_name&#34;)
 *             .secretEngineVersion(2)
 *             .secretId(String.format(&#34;account.%s&#34;, harness_platform_secret_text.test().id()))
 *             .tags(&#34;foo:bar&#34;)
 *             .useAwsIam(false)
 *             .useK8sAuth(false)
 *             .useVaultAgent(false)
 *             .vaultUrl(&#34;https://vault_url.com&#34;)
 *             .build());
 * 
 *         var k8sAuth = new VaultConnector(&#34;k8sAuth&#34;, VaultConnectorArgs.builder()        
 *             .accessType(&#34;K8s_AUTH&#34;)
 *             .authToken(String.format(&#34;account.%s&#34;, harness_platform_secret_text.test().id()))
 *             .basePath(&#34;base_path&#34;)
 *             .default_(false)
 *             .delegateSelectors(&#34;harness-delegate&#34;)
 *             .description(&#34;test&#34;)
 *             .identifier(&#34;identifier&#34;)
 *             .k8sAuthEndpoint(&#34;k8s_auth_endpoint&#34;)
 *             .namespace(&#34;namespace&#34;)
 *             .readOnly(true)
 *             .renewalIntervalMinutes(10)
 *             .secretEngineManuallyConfigured(true)
 *             .secretEngineName(&#34;secret_engine_name&#34;)
 *             .secretEngineVersion(2)
 *             .serviceAccountTokenPath(&#34;service_account_token_path&#34;)
 *             .tags(&#34;foo:bar&#34;)
 *             .useAwsIam(false)
 *             .useK8sAuth(true)
 *             .useVaultAgent(false)
 *             .vaultAwsIamRole(&#34;vault_aws_iam_role&#34;)
 *             .vaultK8sAuthRole(&#34;vault_k8s_auth_role&#34;)
 *             .vaultUrl(&#34;https://vault_url.com&#34;)
 *             .build());
 * 
 *         var vaultAgent = new VaultConnector(&#34;vaultAgent&#34;, VaultConnectorArgs.builder()        
 *             .accessType(&#34;VAULT_AGENT&#34;)
 *             .authToken(String.format(&#34;account.%s&#34;, harness_platform_secret_text.test().id()))
 *             .basePath(&#34;base_path&#34;)
 *             .default_(false)
 *             .delegateSelectors(&#34;harness-delegate&#34;)
 *             .description(&#34;test&#34;)
 *             .identifier(&#34;identifier&#34;)
 *             .namespace(&#34;namespace&#34;)
 *             .readOnly(true)
 *             .renewalIntervalMinutes(10)
 *             .secretEngineManuallyConfigured(true)
 *             .secretEngineName(&#34;secret_engine_name&#34;)
 *             .secretEngineVersion(2)
 *             .sinkPath(&#34;sink_path&#34;)
 *             .tags(&#34;foo:bar&#34;)
 *             .useAwsIam(false)
 *             .useK8sAuth(false)
 *             .useVaultAgent(true)
 *             .vaultUrl(&#34;https://vault_url.com&#34;)
 *             .build());
 * 
 *         var token = new VaultConnector(&#34;token&#34;, VaultConnectorArgs.builder()        
 *             .accessType(&#34;TOKEN&#34;)
 *             .authToken(String.format(&#34;account.%s&#34;, harness_platform_secret_text.test().id()))
 *             .basePath(&#34;base_path&#34;)
 *             .default_(false)
 *             .description(&#34;test&#34;)
 *             .identifier(&#34;identifier&#34;)
 *             .namespace(&#34;namespace&#34;)
 *             .readOnly(true)
 *             .renewalIntervalMinutes(10)
 *             .secretEngineManuallyConfigured(true)
 *             .secretEngineName(&#34;secret_engine_name&#34;)
 *             .secretEngineVersion(2)
 *             .tags(&#34;foo:bar&#34;)
 *             .useAwsIam(false)
 *             .useK8sAuth(false)
 *             .vaultUrl(&#34;https://vault_url.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Import using vault connector id
 * 
 * ```sh
 *  $ pulumi import harness:platform/vaultConnector:VaultConnector example &lt;connector_id&gt;
 * ```
 * 
 */
@ResourceType(type="harness:platform/vaultConnector:VaultConnector")
public class VaultConnector extends com.pulumi.resources.CustomResource {
    /**
     * Access type.
     * 
     */
    @Export(name="accessType", type=String.class, parameters={})
    private Output</* @Nullable */ String> accessType;

    /**
     * @return Access type.
     * 
     */
    public Output<Optional<String>> accessType() {
        return Codegen.optional(this.accessType);
    }
    /**
     * ID of App Role.
     * 
     */
    @Export(name="appRoleId", type=String.class, parameters={})
    private Output</* @Nullable */ String> appRoleId;

    /**
     * @return ID of App Role.
     * 
     */
    public Output<Optional<String>> appRoleId() {
        return Codegen.optional(this.appRoleId);
    }
    /**
     * Authentication token for Vault.
     * 
     */
    @Export(name="authToken", type=String.class, parameters={})
    private Output<String> authToken;

    /**
     * @return Authentication token for Vault.
     * 
     */
    public Output<String> authToken() {
        return this.authToken;
    }
    /**
     * AWS region where the AWS IAM authentication will happen.
     * 
     */
    @Export(name="awsRegion", type=String.class, parameters={})
    private Output</* @Nullable */ String> awsRegion;

    /**
     * @return AWS region where the AWS IAM authentication will happen.
     * 
     */
    public Output<Optional<String>> awsRegion() {
        return Codegen.optional(this.awsRegion);
    }
    /**
     * Location of the Vault directory where the secret will be stored.
     * 
     */
    @Export(name="basePath", type=String.class, parameters={})
    private Output</* @Nullable */ String> basePath;

    /**
     * @return Location of the Vault directory where the secret will be stored.
     * 
     */
    public Output<Optional<String>> basePath() {
        return Codegen.optional(this.basePath);
    }
    /**
     * Is default or not.
     * 
     */
    @Export(name="default", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> default_;

    /**
     * @return Is default or not.
     * 
     */
    public Output<Optional<Boolean>> default_() {
        return Codegen.optional(this.default_);
    }
    /**
     * List of Delegate Selectors that belong to the same Delegate and are used to connect to the Secret Manager.
     * 
     */
    @Export(name="delegateSelectors", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> delegateSelectors;

    /**
     * @return List of Delegate Selectors that belong to the same Delegate and are used to connect to the Secret Manager.
     * 
     */
    public Output<Optional<List<String>>> delegateSelectors() {
        return Codegen.optional(this.delegateSelectors);
    }
    /**
     * Description of the resource.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Unique identifier of the resource.
     * 
     */
    @Export(name="identifier", type=String.class, parameters={})
    private Output<String> identifier;

    /**
     * @return Unique identifier of the resource.
     * 
     */
    public Output<String> identifier() {
        return this.identifier;
    }
    /**
     * Is default or not.
     * 
     */
    @Export(name="isDefault", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> isDefault;

    /**
     * @return Is default or not.
     * 
     */
    public Output<Optional<Boolean>> isDefault() {
        return Codegen.optional(this.isDefault);
    }
    /**
     * Read only or not.
     * 
     */
    @Export(name="isReadOnly", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> isReadOnly;

    /**
     * @return Read only or not.
     * 
     */
    public Output<Optional<Boolean>> isReadOnly() {
        return Codegen.optional(this.isReadOnly);
    }
    /**
     * The path where Kubernetes Auth is enabled in Vault.
     * 
     */
    @Export(name="k8sAuthEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> k8sAuthEndpoint;

    /**
     * @return The path where Kubernetes Auth is enabled in Vault.
     * 
     */
    public Output<Optional<String>> k8sAuthEndpoint() {
        return Codegen.optional(this.k8sAuthEndpoint);
    }
    /**
     * Name of the resource.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the resource.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Vault namespace where the Secret will be created.
     * 
     */
    @Export(name="namespace", type=String.class, parameters={})
    private Output</* @Nullable */ String> namespace;

    /**
     * @return Vault namespace where the Secret will be created.
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }
    /**
     * Unique identifier of the Organization.
     * 
     */
    @Export(name="orgId", type=String.class, parameters={})
    private Output</* @Nullable */ String> orgId;

    /**
     * @return Unique identifier of the Organization.
     * 
     */
    public Output<Optional<String>> orgId() {
        return Codegen.optional(this.orgId);
    }
    /**
     * Unique identifier of the Project.
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
    private Output</* @Nullable */ String> projectId;

    /**
     * @return Unique identifier of the Project.
     * 
     */
    public Output<Optional<String>> projectId() {
        return Codegen.optional(this.projectId);
    }
    /**
     * Read only.
     * 
     */
    @Export(name="readOnly", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> readOnly;

    /**
     * @return Read only.
     * 
     */
    public Output<Optional<Boolean>> readOnly() {
        return Codegen.optional(this.readOnly);
    }
    /**
     * Boolean value to indicate if AppRole token renewal is enabled or not.
     * 
     */
    @Export(name="renewAppRoleToken", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> renewAppRoleToken;

    /**
     * @return Boolean value to indicate if AppRole token renewal is enabled or not.
     * 
     */
    public Output<Optional<Boolean>> renewAppRoleToken() {
        return Codegen.optional(this.renewAppRoleToken);
    }
    /**
     * The time interval for the token renewal.
     * 
     */
    @Export(name="renewalIntervalMinutes", type=Integer.class, parameters={})
    private Output<Integer> renewalIntervalMinutes;

    /**
     * @return The time interval for the token renewal.
     * 
     */
    public Output<Integer> renewalIntervalMinutes() {
        return this.renewalIntervalMinutes;
    }
    /**
     * Manually entered Secret Engine.
     * 
     */
    @Export(name="secretEngineManuallyConfigured", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> secretEngineManuallyConfigured;

    /**
     * @return Manually entered Secret Engine.
     * 
     */
    public Output<Optional<Boolean>> secretEngineManuallyConfigured() {
        return Codegen.optional(this.secretEngineManuallyConfigured);
    }
    /**
     * Name of the Secret Engine.
     * 
     */
    @Export(name="secretEngineName", type=String.class, parameters={})
    private Output</* @Nullable */ String> secretEngineName;

    /**
     * @return Name of the Secret Engine.
     * 
     */
    public Output<Optional<String>> secretEngineName() {
        return Codegen.optional(this.secretEngineName);
    }
    /**
     * Version of Secret Engine.
     * 
     */
    @Export(name="secretEngineVersion", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> secretEngineVersion;

    /**
     * @return Version of Secret Engine.
     * 
     */
    public Output<Optional<Integer>> secretEngineVersion() {
        return Codegen.optional(this.secretEngineVersion);
    }
    /**
     * ID of the Secret.
     * 
     */
    @Export(name="secretId", type=String.class, parameters={})
    private Output</* @Nullable */ String> secretId;

    /**
     * @return ID of the Secret.
     * 
     */
    public Output<Optional<String>> secretId() {
        return Codegen.optional(this.secretId);
    }
    /**
     * The Service Account token path in the K8s pod where the token is mounted.
     * 
     */
    @Export(name="serviceAccountTokenPath", type=String.class, parameters={})
    private Output</* @Nullable */ String> serviceAccountTokenPath;

    /**
     * @return The Service Account token path in the K8s pod where the token is mounted.
     * 
     */
    public Output<Optional<String>> serviceAccountTokenPath() {
        return Codegen.optional(this.serviceAccountTokenPath);
    }
    /**
     * The location from which the authentication token should be read.
     * 
     */
    @Export(name="sinkPath", type=String.class, parameters={})
    private Output</* @Nullable */ String> sinkPath;

    /**
     * @return The location from which the authentication token should be read.
     * 
     */
    public Output<Optional<String>> sinkPath() {
        return Codegen.optional(this.sinkPath);
    }
    /**
     * Tags to associate with the resource. Tags should be in the form `name:value`.
     * 
     */
    @Export(name="tags", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return Tags to associate with the resource. Tags should be in the form `name:value`.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Boolean value to indicate if AWS IAM is used for authentication.
     * 
     */
    @Export(name="useAwsIam", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> useAwsIam;

    /**
     * @return Boolean value to indicate if AWS IAM is used for authentication.
     * 
     */
    public Output<Optional<Boolean>> useAwsIam() {
        return Codegen.optional(this.useAwsIam);
    }
    /**
     * Boolean value to indicate if K8s Auth is used for authentication.
     * 
     */
    @Export(name="useK8sAuth", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> useK8sAuth;

    /**
     * @return Boolean value to indicate if K8s Auth is used for authentication.
     * 
     */
    public Output<Optional<Boolean>> useK8sAuth() {
        return Codegen.optional(this.useK8sAuth);
    }
    /**
     * Boolean value to indicate if Vault Agent is used for authentication.
     * 
     */
    @Export(name="useVaultAgent", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> useVaultAgent;

    /**
     * @return Boolean value to indicate if Vault Agent is used for authentication.
     * 
     */
    public Output<Optional<Boolean>> useVaultAgent() {
        return Codegen.optional(this.useVaultAgent);
    }
    /**
     * The Vault role defined to bind to aws iam account/role being accessed.
     * 
     */
    @Export(name="vaultAwsIamRole", type=String.class, parameters={})
    private Output</* @Nullable */ String> vaultAwsIamRole;

    /**
     * @return The Vault role defined to bind to aws iam account/role being accessed.
     * 
     */
    public Output<Optional<String>> vaultAwsIamRole() {
        return Codegen.optional(this.vaultAwsIamRole);
    }
    /**
     * The role where K8s Auth will happen.
     * 
     */
    @Export(name="vaultK8sAuthRole", type=String.class, parameters={})
    private Output</* @Nullable */ String> vaultK8sAuthRole;

    /**
     * @return The role where K8s Auth will happen.
     * 
     */
    public Output<Optional<String>> vaultK8sAuthRole() {
        return Codegen.optional(this.vaultK8sAuthRole);
    }
    /**
     * URL of the HashiCorp Vault.
     * 
     */
    @Export(name="vaultUrl", type=String.class, parameters={})
    private Output<String> vaultUrl;

    /**
     * @return URL of the HashiCorp Vault.
     * 
     */
    public Output<String> vaultUrl() {
        return this.vaultUrl;
    }
    /**
     * The AWS IAM Header Server ID that has been configured for this AWS IAM instance.
     * 
     */
    @Export(name="xvaultAwsIamServerId", type=String.class, parameters={})
    private Output</* @Nullable */ String> xvaultAwsIamServerId;

    /**
     * @return The AWS IAM Header Server ID that has been configured for this AWS IAM instance.
     * 
     */
    public Output<Optional<String>> xvaultAwsIamServerId() {
        return Codegen.optional(this.xvaultAwsIamServerId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VaultConnector(String name) {
        this(name, VaultConnectorArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VaultConnector(String name, VaultConnectorArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VaultConnector(String name, VaultConnectorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("harness:platform/vaultConnector:VaultConnector", name, args == null ? VaultConnectorArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VaultConnector(String name, Output<String> id, @Nullable VaultConnectorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("harness:platform/vaultConnector:VaultConnector", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static VaultConnector get(String name, Output<String> id, @Nullable VaultConnectorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VaultConnector(name, id, state, options);
    }
}
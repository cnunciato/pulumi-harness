// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.harness.Utilities;
import com.pulumi.harness.platform.MonitoredServiceArgs;
import com.pulumi.harness.platform.inputs.MonitoredServiceState;
import com.pulumi.harness.platform.outputs.MonitoredServiceRequest;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Resource for creating a monitored service.
 * 
 * ## Example Usage
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.harness.platform.MonitoredService;
 * import com.pulumi.harness.platform.MonitoredServiceArgs;
 * import com.pulumi.harness.platform.inputs.MonitoredServiceRequestArgs;
 * import static com.pulumi.codegen.internal.Serialization.*;
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
 *         var example = new MonitoredService(&#34;example&#34;, MonitoredServiceArgs.builder()        
 *             .accountId(&#34;account_id&#34;)
 *             .orgId(&#34;default&#34;)
 *             .projectId(&#34;default_project&#34;)
 *             .identifier(&#34;Terraform&#34;)
 *             .request(MonitoredServiceRequestArgs.builder()
 *                 .name(&#34;name&#34;)
 *                 .type(&#34;Application&#34;)
 *                 .description(&#34;description&#34;)
 *                 .serviceRef(&#34;service_ref&#34;)
 *                 .environmentRef(&#34;environment_ref&#34;)
 *                 .tags(                
 *                     &#34;foo:bar&#34;,
 *                     &#34;bar:foo&#34;)
 *                 .healthSources(MonitoredServiceRequestHealthSourceArgs.builder()
 *                     .name(&#34;name&#34;)
 *                     .identifier(&#34;identifier&#34;)
 *                     .type(&#34;ElasticSearch&#34;)
 *                     .spec(serializeJson(
 *                         jsonObject(
 *                             jsonProperty(&#34;connectorRef&#34;, &#34;connectorRef&#34;),
 *                             jsonProperty(&#34;feature&#34;, &#34;feature&#34;),
 *                             jsonProperty(&#34;queries&#34;, jsonArray(
 *                                 jsonObject(
 *                                     jsonProperty(&#34;name&#34;, &#34;name&#34;),
 *                                     jsonProperty(&#34;query&#34;, &#34;query&#34;),
 *                                     jsonProperty(&#34;index&#34;, &#34;index&#34;),
 *                                     jsonProperty(&#34;serviceInstanceIdentifier&#34;, &#34;serviceInstanceIdentifier&#34;),
 *                                     jsonProperty(&#34;timeStampIdentifier&#34;, &#34;timeStampIdentifier&#34;),
 *                                     jsonProperty(&#34;timeStampFormat&#34;, &#34;timeStampFormat&#34;),
 *                                     jsonProperty(&#34;messageIdentifier&#34;, &#34;messageIdentifier&#34;)
 *                                 ), 
 *                                 jsonObject(
 *                                     jsonProperty(&#34;name&#34;, &#34;name2&#34;),
 *                                     jsonProperty(&#34;query&#34;, &#34;query2&#34;),
 *                                     jsonProperty(&#34;index&#34;, &#34;index2&#34;),
 *                                     jsonProperty(&#34;serviceInstanceIdentifier&#34;, &#34;serviceInstanceIdentifier2&#34;),
 *                                     jsonProperty(&#34;timeStampIdentifier&#34;, &#34;timeStampIdentifier2&#34;),
 *                                     jsonProperty(&#34;timeStampFormat&#34;, &#34;timeStampFormat2&#34;),
 *                                     jsonProperty(&#34;messageIdentifier&#34;, &#34;messageIdentifier2&#34;)
 *                                 )
 *                             ))
 *                         )))
 *                     .build())
 *                 .changeSources(MonitoredServiceRequestChangeSourceArgs.builder()
 *                     .name(&#34;csName1&#34;)
 *                     .identifier(&#34;harness_cd_next_gen&#34;)
 *                     .type(&#34;HarnessCDNextGen&#34;)
 *                     .enabled(true)
 *                     .spec(serializeJson(
 *                         jsonObject(
 * 
 *                         )))
 *                     .category(&#34;Deployment&#34;)
 *                     .build())
 *                 .notificationRuleRefs(                
 *                     MonitoredServiceRequestNotificationRuleRefArgs.builder()
 *                         .notificationRuleRef(&#34;notification_rule_ref&#34;)
 *                         .enabled(true)
 *                         .build(),
 *                     MonitoredServiceRequestNotificationRuleRefArgs.builder()
 *                         .notificationRuleRef(&#34;notification_rule_ref1&#34;)
 *                         .enabled(false)
 *                         .build())
 *                 .enabled(true)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Import account level monitored_service
 * 
 * ```sh
 *  $ pulumi import harness:platform/monitoredService:MonitoredService example &lt;monitored_service_id&gt;
 * ```
 * 
 *  Import organization level monitored_service
 * 
 * ```sh
 *  $ pulumi import harness:platform/monitoredService:MonitoredService example &lt;org_id&gt;/&lt;monitored_service_id&gt;
 * ```
 * 
 *  Import project level monitored_service
 * 
 * ```sh
 *  $ pulumi import harness:platform/monitoredService:MonitoredService example &lt;org_id&gt;/&lt;project_id&gt;/&lt;monitored_service_id&gt;
 * ```
 * 
 */
@ResourceType(type="harness:platform/monitoredService:MonitoredService")
public class MonitoredService extends com.pulumi.resources.CustomResource {
    /**
     * Identifier of the monitored service.
     * 
     */
    @Export(name="identifier", refs={String.class}, tree="[0]")
    private Output<String> identifier;

    /**
     * @return Identifier of the monitored service.
     * 
     */
    public Output<String> identifier() {
        return this.identifier;
    }
    /**
     * Identifier of the organization in which the monitored service is configured.
     * 
     */
    @Export(name="orgId", refs={String.class}, tree="[0]")
    private Output<String> orgId;

    /**
     * @return Identifier of the organization in which the monitored service is configured.
     * 
     */
    public Output<String> orgId() {
        return this.orgId;
    }
    /**
     * Identifier of the project in which the monitored service is configured.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return Identifier of the project in which the monitored service is configured.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * Request for creating or updating a monitored service.
     * 
     */
    @Export(name="request", refs={MonitoredServiceRequest.class}, tree="[0]")
    private Output</* @Nullable */ MonitoredServiceRequest> request;

    /**
     * @return Request for creating or updating a monitored service.
     * 
     */
    public Output<Optional<MonitoredServiceRequest>> request() {
        return Codegen.optional(this.request);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MonitoredService(String name) {
        this(name, MonitoredServiceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MonitoredService(String name, MonitoredServiceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MonitoredService(String name, MonitoredServiceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("harness:platform/monitoredService:MonitoredService", name, args == null ? MonitoredServiceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MonitoredService(String name, Output<String> id, @Nullable MonitoredServiceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("harness:platform/monitoredService:MonitoredService", name, state, makeResourceOptions(options, id));
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
    public static MonitoredService get(String name, Output<String> id, @Nullable MonitoredServiceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MonitoredService(name, id, state, options);
    }
}
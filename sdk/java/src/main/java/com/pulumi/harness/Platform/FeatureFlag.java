// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.harness.Utilities;
import com.pulumi.harness.platform.FeatureFlagArgs;
import com.pulumi.harness.platform.inputs.FeatureFlagState;
import com.pulumi.harness.platform.outputs.FeatureFlagGitDetails;
import com.pulumi.harness.platform.outputs.FeatureFlagVariation;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Resource for managing Feature Flags.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.harness.platform.FeatureFlag;
 * import com.pulumi.harness.platform.FeatureFlagArgs;
 * import com.pulumi.harness.platform.inputs.FeatureFlagVariationArgs;
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
 *         var mybooleanflag = new FeatureFlag(&#34;mybooleanflag&#34;, FeatureFlagArgs.builder()        
 *             .defaultOffVariation(&#34;Disabled&#34;)
 *             .defaultOnVariation(&#34;Enabled&#34;)
 *             .identifier(&#34;MY_FEATURE&#34;)
 *             .kind(&#34;boolean&#34;)
 *             .orgId(&#34;test&#34;)
 *             .permanent(false)
 *             .projectId(&#34;testff&#34;)
 *             .variations(            
 *                 FeatureFlagVariationArgs.builder()
 *                     .description(&#34;The feature is enabled&#34;)
 *                     .identifier(&#34;Enabled&#34;)
 *                     .name(&#34;Enabled&#34;)
 *                     .value(&#34;true&#34;)
 *                     .build(),
 *                 FeatureFlagVariationArgs.builder()
 *                     .description(&#34;The feature is disabled&#34;)
 *                     .identifier(&#34;Disabled&#34;)
 *                     .name(&#34;Disabled&#34;)
 *                     .value(&#34;false&#34;)
 *                     .build())
 *             .build());
 * 
 *         var mymultivariateflag = new FeatureFlag(&#34;mymultivariateflag&#34;, FeatureFlagArgs.builder()        
 *             .defaultOffVariation(&#34;trial20&#34;)
 *             .defaultOnVariation(&#34;trial7&#34;)
 *             .identifier(&#34;FREE_TRIAL_DURATION&#34;)
 *             .kind(&#34;int&#34;)
 *             .orgId(&#34;test&#34;)
 *             .permanent(false)
 *             .projectId(&#34;testff&#34;)
 *             .variations(            
 *                 FeatureFlagVariationArgs.builder()
 *                     .description(&#34;Free trial period 7 days&#34;)
 *                     .identifier(&#34;trial7&#34;)
 *                     .name(&#34;7 days trial&#34;)
 *                     .value(&#34;7&#34;)
 *                     .build(),
 *                 FeatureFlagVariationArgs.builder()
 *                     .description(&#34;Free trial period 14 days&#34;)
 *                     .identifier(&#34;trial14&#34;)
 *                     .name(&#34;14 days trial&#34;)
 *                     .value(&#34;14&#34;)
 *                     .build(),
 *                 FeatureFlagVariationArgs.builder()
 *                     .description(&#34;Free trial period 20 days&#34;)
 *                     .identifier(&#34;trial20&#34;)
 *                     .name(&#34;20 days trial&#34;)
 *                     .value(&#34;20&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="harness:platform/featureFlag:FeatureFlag")
public class FeatureFlag extends com.pulumi.resources.CustomResource {
    /**
     * Whether or not the flag is archived
     * 
     */
    @Export(name="archived", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> archived;

    /**
     * @return Whether or not the flag is archived
     * 
     */
    public Output<Optional<Boolean>> archived() {
        return Codegen.optional(this.archived);
    }
    /**
     * Which of the variations to use when the flag is toggled to off state
     * 
     */
    @Export(name="defaultOffVariation", refs={String.class}, tree="[0]")
    private Output<String> defaultOffVariation;

    /**
     * @return Which of the variations to use when the flag is toggled to off state
     * 
     */
    public Output<String> defaultOffVariation() {
        return this.defaultOffVariation;
    }
    /**
     * Which of the variations to use when the flag is toggled to on state
     * 
     */
    @Export(name="defaultOnVariation", refs={String.class}, tree="[0]")
    private Output<String> defaultOnVariation;

    /**
     * @return Which of the variations to use when the flag is toggled to on state
     * 
     */
    public Output<String> defaultOnVariation() {
        return this.defaultOnVariation;
    }
    @Export(name="gitDetails", refs={FeatureFlagGitDetails.class}, tree="[0]")
    private Output</* @Nullable */ FeatureFlagGitDetails> gitDetails;

    public Output<Optional<FeatureFlagGitDetails>> gitDetails() {
        return Codegen.optional(this.gitDetails);
    }
    /**
     * Identifier of the Feature Flag
     * 
     */
    @Export(name="identifier", refs={String.class}, tree="[0]")
    private Output<String> identifier;

    /**
     * @return Identifier of the Feature Flag
     * 
     */
    public Output<String> identifier() {
        return this.identifier;
    }
    /**
     * The type of data the flag represents. Valid values are `boolean`, `int`, `string`, `json`
     * 
     */
    @Export(name="kind", refs={String.class}, tree="[0]")
    private Output<String> kind;

    /**
     * @return The type of data the flag represents. Valid values are `boolean`, `int`, `string`, `json`
     * 
     */
    public Output<String> kind() {
        return this.kind;
    }
    /**
     * Name of the Feature Flag
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the Feature Flag
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Organization Identifier
     * 
     */
    @Export(name="orgId", refs={String.class}, tree="[0]")
    private Output<String> orgId;

    /**
     * @return Organization Identifier
     * 
     */
    public Output<String> orgId() {
        return this.orgId;
    }
    /**
     * The owner of the flag
     * 
     */
    @Export(name="owner", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> owner;

    /**
     * @return The owner of the flag
     * 
     */
    public Output<Optional<String>> owner() {
        return Codegen.optional(this.owner);
    }
    /**
     * Whether or not the flag is permanent. If it is, it will never be flagged as stale
     * 
     */
    @Export(name="permanent", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> permanent;

    /**
     * @return Whether or not the flag is permanent. If it is, it will never be flagged as stale
     * 
     */
    public Output<Boolean> permanent() {
        return this.permanent;
    }
    /**
     * Project Identifier
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return Project Identifier
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The options available for your flag
     * 
     */
    @Export(name="variations", refs={List.class,FeatureFlagVariation.class}, tree="[0,1]")
    private Output<List<FeatureFlagVariation>> variations;

    /**
     * @return The options available for your flag
     * 
     */
    public Output<List<FeatureFlagVariation>> variations() {
        return this.variations;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FeatureFlag(String name) {
        this(name, FeatureFlagArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FeatureFlag(String name, FeatureFlagArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FeatureFlag(String name, FeatureFlagArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("harness:platform/featureFlag:FeatureFlag", name, args == null ? FeatureFlagArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private FeatureFlag(String name, Output<String> id, @Nullable FeatureFlagState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("harness:platform/featureFlag:FeatureFlag", name, state, makeResourceOptions(options, id));
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
    public static FeatureFlag get(String name, Output<String> id, @Nullable FeatureFlagState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FeatureFlag(name, id, state, options);
    }
}
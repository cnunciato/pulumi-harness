// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Resource for creating an SLO.
 *
 * ## Import
 *
 * Import account level SLO
 *
 * ```sh
 *  $ pulumi import harness:platform/slo:Slo example <slo_id>
 * ```
 *
 *  Import organization level SLO
 *
 * ```sh
 *  $ pulumi import harness:platform/slo:Slo example <org_id>/<slo_id>
 * ```
 *
 *  Import project level SLO
 *
 * ```sh
 *  $ pulumi import harness:platform/slo:Slo example <org_id>/<project_id>/<slo_id>
 * ```
 */
export class Slo extends pulumi.CustomResource {
    /**
     * Get an existing Slo resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SloState, opts?: pulumi.CustomResourceOptions): Slo {
        return new Slo(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'harness:platform/slo:Slo';

    /**
     * Returns true if the given object is an instance of Slo.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Slo {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Slo.__pulumiType;
    }

    /**
     * Identifier of the SLO.
     */
    public readonly identifier!: pulumi.Output<string>;
    /**
     * Identifier of the organization in which the SLO is configured.
     */
    public readonly orgId!: pulumi.Output<string>;
    /**
     * Identifier of the project in which the SLO is configured.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Request for creating or updating SLO.
     */
    public readonly request!: pulumi.Output<outputs.platform.SloRequest | undefined>;

    /**
     * Create a Slo resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SloArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SloArgs | SloState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SloState | undefined;
            resourceInputs["identifier"] = state ? state.identifier : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["request"] = state ? state.request : undefined;
        } else {
            const args = argsOrState as SloArgs | undefined;
            if ((!args || args.identifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identifier'");
            }
            if ((!args || args.orgId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'orgId'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["identifier"] = args ? args.identifier : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["request"] = args ? args.request : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Slo.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Slo resources.
 */
export interface SloState {
    /**
     * Identifier of the SLO.
     */
    identifier?: pulumi.Input<string>;
    /**
     * Identifier of the organization in which the SLO is configured.
     */
    orgId?: pulumi.Input<string>;
    /**
     * Identifier of the project in which the SLO is configured.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Request for creating or updating SLO.
     */
    request?: pulumi.Input<inputs.platform.SloRequest>;
}

/**
 * The set of arguments for constructing a Slo resource.
 */
export interface SloArgs {
    /**
     * Identifier of the SLO.
     */
    identifier: pulumi.Input<string>;
    /**
     * Identifier of the organization in which the SLO is configured.
     */
    orgId: pulumi.Input<string>;
    /**
     * Identifier of the project in which the SLO is configured.
     */
    projectId: pulumi.Input<string>;
    /**
     * Request for creating or updating SLO.
     */
    request?: pulumi.Input<inputs.platform.SloRequest>;
}
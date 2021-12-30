// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Identical to random.RandomString with the exception that the result is treated as sensitive and, thus, _not_ displayed in console output.
 *
 * This resource *does* use a cryptographic random number generator.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as random from "@pulumi/random";
 *
 * const password = new random.RandomPassword("password", {
 *     length: 16,
 *     special: true,
 *     overrideSpecial: `_%@`,
 * });
 * const example = new aws.rds.Instance("example", {
 *     instanceClass: "db.t3.micro",
 *     allocatedStorage: 64,
 *     engine: "mysql",
 *     username: "someone",
 *     password: password.result,
 * });
 * ```
 *
 * ## Import
 *
 * # Random Password can be imported by specifying the value of the string
 *
 * ```sh
 *  $ pulumi import random:index/randomPassword:RandomPassword password securepassword
 * ```
 */
export class RandomPassword extends pulumi.CustomResource {
    /**
     * Get an existing RandomPassword resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RandomPasswordState, opts?: pulumi.CustomResourceOptions): RandomPassword {
        return new RandomPassword(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'random:index/randomPassword:RandomPassword';

    /**
     * Returns true if the given object is an instance of RandomPassword.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RandomPassword {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RandomPassword.__pulumiType;
    }

    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     */
    public readonly keepers!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The length of the string desired.
     */
    public readonly length!: pulumi.Output<number>;
    /**
     * Include lowercase alphabet characters in the result.
     */
    public readonly lower!: pulumi.Output<boolean | undefined>;
    /**
     * Minimum number of lowercase alphabet characters in the result.
     */
    public readonly minLower!: pulumi.Output<number | undefined>;
    /**
     * Minimum number of numeric characters in the result.
     */
    public readonly minNumeric!: pulumi.Output<number | undefined>;
    /**
     * Minimum number of special characters in the result.
     */
    public readonly minSpecial!: pulumi.Output<number | undefined>;
    /**
     * Minimum number of uppercase alphabet characters in the result.
     */
    public readonly minUpper!: pulumi.Output<number | undefined>;
    /**
     * Include numeric characters in the result.
     */
    public readonly number!: pulumi.Output<boolean | undefined>;
    /**
     * Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
     */
    public readonly overrideSpecial!: pulumi.Output<string | undefined>;
    /**
     * The generated random string.
     */
    public /*out*/ readonly result!: pulumi.Output<string>;
    /**
     * Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`
     */
    public readonly special!: pulumi.Output<boolean | undefined>;
    /**
     * Include uppercase alphabet characters in the result.
     */
    public readonly upper!: pulumi.Output<boolean | undefined>;

    /**
     * Create a RandomPassword resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RandomPasswordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RandomPasswordArgs | RandomPasswordState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RandomPasswordState | undefined;
            resourceInputs["keepers"] = state ? state.keepers : undefined;
            resourceInputs["length"] = state ? state.length : undefined;
            resourceInputs["lower"] = state ? state.lower : undefined;
            resourceInputs["minLower"] = state ? state.minLower : undefined;
            resourceInputs["minNumeric"] = state ? state.minNumeric : undefined;
            resourceInputs["minSpecial"] = state ? state.minSpecial : undefined;
            resourceInputs["minUpper"] = state ? state.minUpper : undefined;
            resourceInputs["number"] = state ? state.number : undefined;
            resourceInputs["overrideSpecial"] = state ? state.overrideSpecial : undefined;
            resourceInputs["result"] = state ? state.result : undefined;
            resourceInputs["special"] = state ? state.special : undefined;
            resourceInputs["upper"] = state ? state.upper : undefined;
        } else {
            const args = argsOrState as RandomPasswordArgs | undefined;
            if ((!args || args.length === undefined) && !opts.urn) {
                throw new Error("Missing required property 'length'");
            }
            resourceInputs["keepers"] = args ? args.keepers : undefined;
            resourceInputs["length"] = args ? args.length : undefined;
            resourceInputs["lower"] = args ? args.lower : undefined;
            resourceInputs["minLower"] = args ? args.minLower : undefined;
            resourceInputs["minNumeric"] = args ? args.minNumeric : undefined;
            resourceInputs["minSpecial"] = args ? args.minSpecial : undefined;
            resourceInputs["minUpper"] = args ? args.minUpper : undefined;
            resourceInputs["number"] = args ? args.number : undefined;
            resourceInputs["overrideSpecial"] = args ? args.overrideSpecial : undefined;
            resourceInputs["special"] = args ? args.special : undefined;
            resourceInputs["upper"] = args ? args.upper : undefined;
            resourceInputs["result"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RandomPassword.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RandomPassword resources.
 */
export interface RandomPasswordState {
    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     */
    keepers?: pulumi.Input<{[key: string]: any}>;
    /**
     * The length of the string desired.
     */
    length?: pulumi.Input<number>;
    /**
     * Include lowercase alphabet characters in the result.
     */
    lower?: pulumi.Input<boolean>;
    /**
     * Minimum number of lowercase alphabet characters in the result.
     */
    minLower?: pulumi.Input<number>;
    /**
     * Minimum number of numeric characters in the result.
     */
    minNumeric?: pulumi.Input<number>;
    /**
     * Minimum number of special characters in the result.
     */
    minSpecial?: pulumi.Input<number>;
    /**
     * Minimum number of uppercase alphabet characters in the result.
     */
    minUpper?: pulumi.Input<number>;
    /**
     * Include numeric characters in the result.
     */
    number?: pulumi.Input<boolean>;
    /**
     * Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
     */
    overrideSpecial?: pulumi.Input<string>;
    /**
     * The generated random string.
     */
    result?: pulumi.Input<string>;
    /**
     * Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`
     */
    special?: pulumi.Input<boolean>;
    /**
     * Include uppercase alphabet characters in the result.
     */
    upper?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a RandomPassword resource.
 */
export interface RandomPasswordArgs {
    /**
     * Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
     */
    keepers?: pulumi.Input<{[key: string]: any}>;
    /**
     * The length of the string desired.
     */
    length: pulumi.Input<number>;
    /**
     * Include lowercase alphabet characters in the result.
     */
    lower?: pulumi.Input<boolean>;
    /**
     * Minimum number of lowercase alphabet characters in the result.
     */
    minLower?: pulumi.Input<number>;
    /**
     * Minimum number of numeric characters in the result.
     */
    minNumeric?: pulumi.Input<number>;
    /**
     * Minimum number of special characters in the result.
     */
    minSpecial?: pulumi.Input<number>;
    /**
     * Minimum number of uppercase alphabet characters in the result.
     */
    minUpper?: pulumi.Input<number>;
    /**
     * Include numeric characters in the result.
     */
    number?: pulumi.Input<boolean>;
    /**
     * Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
     */
    overrideSpecial?: pulumi.Input<string>;
    /**
     * Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`
     */
    special?: pulumi.Input<boolean>;
    /**
     * Include uppercase alphabet characters in the result.
     */
    upper?: pulumi.Input<boolean>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package random

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > **Note:** Requires random provider version >= 2.2.0
// 
// Identical to .RandomString with the exception that the
// result is treated as sensitive and, thus, _not_ displayed in console output.
// 
// > **Note:** All attributes including the generated password will be stored in
// the raw state as plain-text. [Read more about sensitive data in
// state](https://www.terraform.io/docs/state/sensitive-data.html).
// 
// This resource *does* use a cryptographic random number generator.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-random/blob/master/website/docs/r/password.html.markdown.
type RandomPassword struct {
	s *pulumi.ResourceState
}

// NewRandomPassword registers a new resource with the given unique name, arguments, and options.
func NewRandomPassword(ctx *pulumi.Context,
	name string, args *RandomPasswordArgs, opts ...pulumi.ResourceOpt) (*RandomPassword, error) {
	if args == nil || args.Length == nil {
		return nil, errors.New("missing required argument 'Length'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["keepers"] = nil
		inputs["length"] = nil
		inputs["lower"] = nil
		inputs["minLower"] = nil
		inputs["minNumeric"] = nil
		inputs["minSpecial"] = nil
		inputs["minUpper"] = nil
		inputs["number"] = nil
		inputs["overrideSpecial"] = nil
		inputs["special"] = nil
		inputs["upper"] = nil
	} else {
		inputs["keepers"] = args.Keepers
		inputs["length"] = args.Length
		inputs["lower"] = args.Lower
		inputs["minLower"] = args.MinLower
		inputs["minNumeric"] = args.MinNumeric
		inputs["minSpecial"] = args.MinSpecial
		inputs["minUpper"] = args.MinUpper
		inputs["number"] = args.Number
		inputs["overrideSpecial"] = args.OverrideSpecial
		inputs["special"] = args.Special
		inputs["upper"] = args.Upper
	}
	inputs["result"] = nil
	s, err := ctx.RegisterResource("random:index/randomPassword:RandomPassword", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RandomPassword{s: s}, nil
}

// GetRandomPassword gets an existing RandomPassword resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandomPassword(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RandomPasswordState, opts ...pulumi.ResourceOpt) (*RandomPassword, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["keepers"] = state.Keepers
		inputs["length"] = state.Length
		inputs["lower"] = state.Lower
		inputs["minLower"] = state.MinLower
		inputs["minNumeric"] = state.MinNumeric
		inputs["minSpecial"] = state.MinSpecial
		inputs["minUpper"] = state.MinUpper
		inputs["number"] = state.Number
		inputs["overrideSpecial"] = state.OverrideSpecial
		inputs["result"] = state.Result
		inputs["special"] = state.Special
		inputs["upper"] = state.Upper
	}
	s, err := ctx.ReadResource("random:index/randomPassword:RandomPassword", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RandomPassword{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RandomPassword) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RandomPassword) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *RandomPassword) Keepers() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["keepers"])
}

func (r *RandomPassword) Length() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["length"])
}

func (r *RandomPassword) Lower() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["lower"])
}

func (r *RandomPassword) MinLower() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["minLower"])
}

func (r *RandomPassword) MinNumeric() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["minNumeric"])
}

func (r *RandomPassword) MinSpecial() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["minSpecial"])
}

func (r *RandomPassword) MinUpper() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["minUpper"])
}

func (r *RandomPassword) Number() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["number"])
}

func (r *RandomPassword) OverrideSpecial() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["overrideSpecial"])
}

func (r *RandomPassword) Result() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["result"])
}

func (r *RandomPassword) Special() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["special"])
}

func (r *RandomPassword) Upper() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["upper"])
}

// Input properties used for looking up and filtering RandomPassword resources.
type RandomPasswordState struct {
	Keepers interface{}
	Length interface{}
	Lower interface{}
	MinLower interface{}
	MinNumeric interface{}
	MinSpecial interface{}
	MinUpper interface{}
	Number interface{}
	OverrideSpecial interface{}
	Result interface{}
	Special interface{}
	Upper interface{}
}

// The set of arguments for constructing a RandomPassword resource.
type RandomPasswordArgs struct {
	Keepers interface{}
	Length interface{}
	Lower interface{}
	MinLower interface{}
	MinNumeric interface{}
	MinSpecial interface{}
	MinUpper interface{}
	Number interface{}
	OverrideSpecial interface{}
	Special interface{}
	Upper interface{}
}

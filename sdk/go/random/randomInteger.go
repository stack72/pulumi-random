// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package random

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The resource `RandomInteger` generates random values from a given range, described by the `min` and `max` attributes of a given resource.
//
// This resource can be used in conjunction with resources that have
// the `createBeforeDestroy` lifecycle flag set, to avoid conflicts with
// unique names during the brief period where both the old and new resources
// exist concurrently.
//
// ## Example Usage
//
// The following example shows how to generate a random priority between 1 and 50000 for
// a `awsAlbListenerRule` resource:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/alb"
// 	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		priority, err := random.NewRandomInteger(ctx, "priority", &random.RandomIntegerArgs{
// 			Keepers: pulumi.StringMap{
// 				"listener_arn": pulumi.String(_var.Listener_arn),
// 			},
// 			Max: pulumi.Int(50000),
// 			Min: pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = alb.NewListenerRule(ctx, "main", &alb.ListenerRuleArgs{
// 			Actions: alb.ListenerRuleActionArray{
// 				&alb.ListenerRuleActionArgs{
// 					TargetGroupArn: pulumi.String(_var.Target_group_arn),
// 					Type:           pulumi.String("forward"),
// 				},
// 			},
// 			ListenerArn: pulumi.String(_var.Listener_arn),
// 			Priority:    priority.Result,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// The result of the above will set a random priority.
type RandomInteger struct {
	pulumi.CustomResourceState

	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated.
	Keepers pulumi.MapOutput `pulumi:"keepers"`
	// The maximum inclusive value of the range.
	Max pulumi.IntOutput `pulumi:"max"`
	// The minimum inclusive value of the range.
	Min pulumi.IntOutput `pulumi:"min"`
	// (int) The random Integer result.
	Result pulumi.IntOutput `pulumi:"result"`
	// A custom seed to always produce the same value.
	Seed pulumi.StringPtrOutput `pulumi:"seed"`
}

// NewRandomInteger registers a new resource with the given unique name, arguments, and options.
func NewRandomInteger(ctx *pulumi.Context,
	name string, args *RandomIntegerArgs, opts ...pulumi.ResourceOption) (*RandomInteger, error) {
	if args == nil || args.Max == nil {
		return nil, errors.New("missing required argument 'Max'")
	}
	if args == nil || args.Min == nil {
		return nil, errors.New("missing required argument 'Min'")
	}
	if args == nil {
		args = &RandomIntegerArgs{}
	}
	var resource RandomInteger
	err := ctx.RegisterResource("random:index/randomInteger:RandomInteger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRandomInteger gets an existing RandomInteger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandomInteger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RandomIntegerState, opts ...pulumi.ResourceOption) (*RandomInteger, error) {
	var resource RandomInteger
	err := ctx.ReadResource("random:index/randomInteger:RandomInteger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RandomInteger resources.
type randomIntegerState struct {
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// The maximum inclusive value of the range.
	Max *int `pulumi:"max"`
	// The minimum inclusive value of the range.
	Min *int `pulumi:"min"`
	// (int) The random Integer result.
	Result *int `pulumi:"result"`
	// A custom seed to always produce the same value.
	Seed *string `pulumi:"seed"`
}

type RandomIntegerState struct {
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated.
	Keepers pulumi.MapInput
	// The maximum inclusive value of the range.
	Max pulumi.IntPtrInput
	// The minimum inclusive value of the range.
	Min pulumi.IntPtrInput
	// (int) The random Integer result.
	Result pulumi.IntPtrInput
	// A custom seed to always produce the same value.
	Seed pulumi.StringPtrInput
}

func (RandomIntegerState) ElementType() reflect.Type {
	return reflect.TypeOf((*randomIntegerState)(nil)).Elem()
}

type randomIntegerArgs struct {
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// The maximum inclusive value of the range.
	Max int `pulumi:"max"`
	// The minimum inclusive value of the range.
	Min int `pulumi:"min"`
	// A custom seed to always produce the same value.
	Seed *string `pulumi:"seed"`
}

// The set of arguments for constructing a RandomInteger resource.
type RandomIntegerArgs struct {
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated.
	Keepers pulumi.MapInput
	// The maximum inclusive value of the range.
	Max pulumi.IntInput
	// The minimum inclusive value of the range.
	Min pulumi.IntInput
	// A custom seed to always produce the same value.
	Seed pulumi.StringPtrInput
}

func (RandomIntegerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomIntegerArgs)(nil)).Elem()
}

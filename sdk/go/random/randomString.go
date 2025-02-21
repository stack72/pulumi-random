// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package random

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The resource `RandomString` generates a random permutation of alphanumeric characters and optionally special characters.
//
// This resource *does* use a cryptographic random number generator.
//
// Historically this resource's intended usage has been ambiguous as the original example used it in a password. For backwards compatibility it will continue to exist. For unique ids please use random_id, for sensitive random values please use random_password.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := random.NewRandomString(ctx, "random", &random.RandomStringArgs{
// 			Length:          pulumi.Int(16),
// 			OverrideSpecial: pulumi.String(fmt.Sprintf("%v%v", "/@£", "$")),
// 			Special:         pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// # Strings can be imported by just specifying the value of the string
//
// ```sh
//  $ pulumi import random:index/randomString:RandomString test test
// ```
type RandomString struct {
	pulumi.CustomResourceState

	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers pulumi.MapOutput `pulumi:"keepers"`
	// The length of the string desired.
	Length pulumi.IntOutput `pulumi:"length"`
	// Include lowercase alphabet characters in the result.
	Lower pulumi.BoolPtrOutput `pulumi:"lower"`
	// Minimum number of lowercase alphabet characters in the result.
	MinLower pulumi.IntPtrOutput `pulumi:"minLower"`
	// Minimum number of numeric characters in the result.
	MinNumeric pulumi.IntPtrOutput `pulumi:"minNumeric"`
	// Minimum number of special characters in the result.
	MinSpecial pulumi.IntPtrOutput `pulumi:"minSpecial"`
	// Minimum number of uppercase alphabet characters in the result.
	MinUpper pulumi.IntPtrOutput `pulumi:"minUpper"`
	// Include numeric characters in the result.
	Number pulumi.BoolPtrOutput `pulumi:"number"`
	// Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
	OverrideSpecial pulumi.StringPtrOutput `pulumi:"overrideSpecial"`
	// The generated random string.
	Result pulumi.StringOutput `pulumi:"result"`
	// Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`
	Special pulumi.BoolPtrOutput `pulumi:"special"`
	// Include uppercase alphabet characters in the result.
	Upper pulumi.BoolPtrOutput `pulumi:"upper"`
}

// NewRandomString registers a new resource with the given unique name, arguments, and options.
func NewRandomString(ctx *pulumi.Context,
	name string, args *RandomStringArgs, opts ...pulumi.ResourceOption) (*RandomString, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Length == nil {
		return nil, errors.New("invalid value for required argument 'Length'")
	}
	var resource RandomString
	err := ctx.RegisterResource("random:index/randomString:RandomString", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRandomString gets an existing RandomString resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandomString(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RandomStringState, opts ...pulumi.ResourceOption) (*RandomString, error) {
	var resource RandomString
	err := ctx.ReadResource("random:index/randomString:RandomString", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RandomString resources.
type randomStringState struct {
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// The length of the string desired.
	Length *int `pulumi:"length"`
	// Include lowercase alphabet characters in the result.
	Lower *bool `pulumi:"lower"`
	// Minimum number of lowercase alphabet characters in the result.
	MinLower *int `pulumi:"minLower"`
	// Minimum number of numeric characters in the result.
	MinNumeric *int `pulumi:"minNumeric"`
	// Minimum number of special characters in the result.
	MinSpecial *int `pulumi:"minSpecial"`
	// Minimum number of uppercase alphabet characters in the result.
	MinUpper *int `pulumi:"minUpper"`
	// Include numeric characters in the result.
	Number *bool `pulumi:"number"`
	// Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
	OverrideSpecial *string `pulumi:"overrideSpecial"`
	// The generated random string.
	Result *string `pulumi:"result"`
	// Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`
	Special *bool `pulumi:"special"`
	// Include uppercase alphabet characters in the result.
	Upper *bool `pulumi:"upper"`
}

type RandomStringState struct {
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers pulumi.MapInput
	// The length of the string desired.
	Length pulumi.IntPtrInput
	// Include lowercase alphabet characters in the result.
	Lower pulumi.BoolPtrInput
	// Minimum number of lowercase alphabet characters in the result.
	MinLower pulumi.IntPtrInput
	// Minimum number of numeric characters in the result.
	MinNumeric pulumi.IntPtrInput
	// Minimum number of special characters in the result.
	MinSpecial pulumi.IntPtrInput
	// Minimum number of uppercase alphabet characters in the result.
	MinUpper pulumi.IntPtrInput
	// Include numeric characters in the result.
	Number pulumi.BoolPtrInput
	// Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
	OverrideSpecial pulumi.StringPtrInput
	// The generated random string.
	Result pulumi.StringPtrInput
	// Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`
	Special pulumi.BoolPtrInput
	// Include uppercase alphabet characters in the result.
	Upper pulumi.BoolPtrInput
}

func (RandomStringState) ElementType() reflect.Type {
	return reflect.TypeOf((*randomStringState)(nil)).Elem()
}

type randomStringArgs struct {
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// The length of the string desired.
	Length int `pulumi:"length"`
	// Include lowercase alphabet characters in the result.
	Lower *bool `pulumi:"lower"`
	// Minimum number of lowercase alphabet characters in the result.
	MinLower *int `pulumi:"minLower"`
	// Minimum number of numeric characters in the result.
	MinNumeric *int `pulumi:"minNumeric"`
	// Minimum number of special characters in the result.
	MinSpecial *int `pulumi:"minSpecial"`
	// Minimum number of uppercase alphabet characters in the result.
	MinUpper *int `pulumi:"minUpper"`
	// Include numeric characters in the result.
	Number *bool `pulumi:"number"`
	// Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
	OverrideSpecial *string `pulumi:"overrideSpecial"`
	// Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`
	Special *bool `pulumi:"special"`
	// Include uppercase alphabet characters in the result.
	Upper *bool `pulumi:"upper"`
}

// The set of arguments for constructing a RandomString resource.
type RandomStringArgs struct {
	// Arbitrary map of values that, when changed, will trigger recreation of resource. See the main provider documentation for more information.
	Keepers pulumi.MapInput
	// The length of the string desired.
	Length pulumi.IntInput
	// Include lowercase alphabet characters in the result.
	Lower pulumi.BoolPtrInput
	// Minimum number of lowercase alphabet characters in the result.
	MinLower pulumi.IntPtrInput
	// Minimum number of numeric characters in the result.
	MinNumeric pulumi.IntPtrInput
	// Minimum number of special characters in the result.
	MinSpecial pulumi.IntPtrInput
	// Minimum number of uppercase alphabet characters in the result.
	MinUpper pulumi.IntPtrInput
	// Include numeric characters in the result.
	Number pulumi.BoolPtrInput
	// Supply your own list of special characters to use for string generation.  This overrides the default character list in the special argument.  The `special` argument must still be set to true for any overwritten characters to be used in generation.
	OverrideSpecial pulumi.StringPtrInput
	// Include special characters in the result. These are `!@#$%&*()-_=+[]{}<>:?`
	Special pulumi.BoolPtrInput
	// Include uppercase alphabet characters in the result.
	Upper pulumi.BoolPtrInput
}

func (RandomStringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomStringArgs)(nil)).Elem()
}

type RandomStringInput interface {
	pulumi.Input

	ToRandomStringOutput() RandomStringOutput
	ToRandomStringOutputWithContext(ctx context.Context) RandomStringOutput
}

func (*RandomString) ElementType() reflect.Type {
	return reflect.TypeOf((*RandomString)(nil))
}

func (i *RandomString) ToRandomStringOutput() RandomStringOutput {
	return i.ToRandomStringOutputWithContext(context.Background())
}

func (i *RandomString) ToRandomStringOutputWithContext(ctx context.Context) RandomStringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomStringOutput)
}

func (i *RandomString) ToRandomStringPtrOutput() RandomStringPtrOutput {
	return i.ToRandomStringPtrOutputWithContext(context.Background())
}

func (i *RandomString) ToRandomStringPtrOutputWithContext(ctx context.Context) RandomStringPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomStringPtrOutput)
}

type RandomStringPtrInput interface {
	pulumi.Input

	ToRandomStringPtrOutput() RandomStringPtrOutput
	ToRandomStringPtrOutputWithContext(ctx context.Context) RandomStringPtrOutput
}

type randomStringPtrType RandomStringArgs

func (*randomStringPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomString)(nil))
}

func (i *randomStringPtrType) ToRandomStringPtrOutput() RandomStringPtrOutput {
	return i.ToRandomStringPtrOutputWithContext(context.Background())
}

func (i *randomStringPtrType) ToRandomStringPtrOutputWithContext(ctx context.Context) RandomStringPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomStringPtrOutput)
}

// RandomStringArrayInput is an input type that accepts RandomStringArray and RandomStringArrayOutput values.
// You can construct a concrete instance of `RandomStringArrayInput` via:
//
//          RandomStringArray{ RandomStringArgs{...} }
type RandomStringArrayInput interface {
	pulumi.Input

	ToRandomStringArrayOutput() RandomStringArrayOutput
	ToRandomStringArrayOutputWithContext(context.Context) RandomStringArrayOutput
}

type RandomStringArray []RandomStringInput

func (RandomStringArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RandomString)(nil)).Elem()
}

func (i RandomStringArray) ToRandomStringArrayOutput() RandomStringArrayOutput {
	return i.ToRandomStringArrayOutputWithContext(context.Background())
}

func (i RandomStringArray) ToRandomStringArrayOutputWithContext(ctx context.Context) RandomStringArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomStringArrayOutput)
}

// RandomStringMapInput is an input type that accepts RandomStringMap and RandomStringMapOutput values.
// You can construct a concrete instance of `RandomStringMapInput` via:
//
//          RandomStringMap{ "key": RandomStringArgs{...} }
type RandomStringMapInput interface {
	pulumi.Input

	ToRandomStringMapOutput() RandomStringMapOutput
	ToRandomStringMapOutputWithContext(context.Context) RandomStringMapOutput
}

type RandomStringMap map[string]RandomStringInput

func (RandomStringMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RandomString)(nil)).Elem()
}

func (i RandomStringMap) ToRandomStringMapOutput() RandomStringMapOutput {
	return i.ToRandomStringMapOutputWithContext(context.Background())
}

func (i RandomStringMap) ToRandomStringMapOutputWithContext(ctx context.Context) RandomStringMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomStringMapOutput)
}

type RandomStringOutput struct{ *pulumi.OutputState }

func (RandomStringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RandomString)(nil))
}

func (o RandomStringOutput) ToRandomStringOutput() RandomStringOutput {
	return o
}

func (o RandomStringOutput) ToRandomStringOutputWithContext(ctx context.Context) RandomStringOutput {
	return o
}

func (o RandomStringOutput) ToRandomStringPtrOutput() RandomStringPtrOutput {
	return o.ToRandomStringPtrOutputWithContext(context.Background())
}

func (o RandomStringOutput) ToRandomStringPtrOutputWithContext(ctx context.Context) RandomStringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RandomString) *RandomString {
		return &v
	}).(RandomStringPtrOutput)
}

type RandomStringPtrOutput struct{ *pulumi.OutputState }

func (RandomStringPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomString)(nil))
}

func (o RandomStringPtrOutput) ToRandomStringPtrOutput() RandomStringPtrOutput {
	return o
}

func (o RandomStringPtrOutput) ToRandomStringPtrOutputWithContext(ctx context.Context) RandomStringPtrOutput {
	return o
}

func (o RandomStringPtrOutput) Elem() RandomStringOutput {
	return o.ApplyT(func(v *RandomString) RandomString {
		if v != nil {
			return *v
		}
		var ret RandomString
		return ret
	}).(RandomStringOutput)
}

type RandomStringArrayOutput struct{ *pulumi.OutputState }

func (RandomStringArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RandomString)(nil))
}

func (o RandomStringArrayOutput) ToRandomStringArrayOutput() RandomStringArrayOutput {
	return o
}

func (o RandomStringArrayOutput) ToRandomStringArrayOutputWithContext(ctx context.Context) RandomStringArrayOutput {
	return o
}

func (o RandomStringArrayOutput) Index(i pulumi.IntInput) RandomStringOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RandomString {
		return vs[0].([]RandomString)[vs[1].(int)]
	}).(RandomStringOutput)
}

type RandomStringMapOutput struct{ *pulumi.OutputState }

func (RandomStringMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RandomString)(nil))
}

func (o RandomStringMapOutput) ToRandomStringMapOutput() RandomStringMapOutput {
	return o
}

func (o RandomStringMapOutput) ToRandomStringMapOutputWithContext(ctx context.Context) RandomStringMapOutput {
	return o
}

func (o RandomStringMapOutput) MapIndex(k pulumi.StringInput) RandomStringOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RandomString {
		return vs[0].(map[string]RandomString)[vs[1].(string)]
	}).(RandomStringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RandomStringInput)(nil)).Elem(), &RandomString{})
	pulumi.RegisterInputType(reflect.TypeOf((*RandomStringPtrInput)(nil)).Elem(), &RandomString{})
	pulumi.RegisterInputType(reflect.TypeOf((*RandomStringArrayInput)(nil)).Elem(), RandomStringArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RandomStringMapInput)(nil)).Elem(), RandomStringMap{})
	pulumi.RegisterOutputType(RandomStringOutput{})
	pulumi.RegisterOutputType(RandomStringPtrOutput{})
	pulumi.RegisterOutputType(RandomStringArrayOutput{})
	pulumi.RegisterOutputType(RandomStringMapOutput{})
}

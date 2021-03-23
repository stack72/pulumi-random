// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package random

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "random:index/randomId:RandomId":
		r, err = NewRandomId(ctx, name, nil, pulumi.URN_(urn))
	case "random:index/randomInteger:RandomInteger":
		r, err = NewRandomInteger(ctx, name, nil, pulumi.URN_(urn))
	case "random:index/randomPassword:RandomPassword":
		r, err = NewRandomPassword(ctx, name, nil, pulumi.URN_(urn))
	case "random:index/randomPet:RandomPet":
		r, err = NewRandomPet(ctx, name, nil, pulumi.URN_(urn))
	case "random:index/randomShuffle:RandomShuffle":
		r, err = NewRandomShuffle(ctx, name, nil, pulumi.URN_(urn))
	case "random:index/randomString:RandomString":
		r, err = NewRandomString(ctx, name, nil, pulumi.URN_(urn))
	case "random:index/randomUuid:RandomUuid":
		r, err = NewRandomUuid(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:random" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	return NewProvider(ctx, name, nil, pulumi.URN_(urn))
}

func init() {
	version, err := PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"random",
		"index/randomId",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"random",
		"index/randomInteger",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"random",
		"index/randomPassword",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"random",
		"index/randomPet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"random",
		"index/randomShuffle",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"random",
		"index/randomString",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"random",
		"index/randomUuid",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"random",
		&pkg{version},
	)
}

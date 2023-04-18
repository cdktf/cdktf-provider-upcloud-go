package loadbalancer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v6/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v6/loadbalancer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerNodesNetworksIpAddressesList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) LoadbalancerNodesNetworksIpAddressesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LoadbalancerNodesNetworksIpAddressesList
type jsiiProxy_LoadbalancerNodesNetworksIpAddressesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_LoadbalancerNodesNetworksIpAddressesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerNodesNetworksIpAddressesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerNodesNetworksIpAddressesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerNodesNetworksIpAddressesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerNodesNetworksIpAddressesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewLoadbalancerNodesNetworksIpAddressesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) LoadbalancerNodesNetworksIpAddressesList {
	_init_.Initialize()

	if err := validateNewLoadbalancerNodesNetworksIpAddressesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LoadbalancerNodesNetworksIpAddressesList{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.loadbalancer.LoadbalancerNodesNetworksIpAddressesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewLoadbalancerNodesNetworksIpAddressesList_Override(l LoadbalancerNodesNetworksIpAddressesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.loadbalancer.LoadbalancerNodesNetworksIpAddressesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		l,
	)
}

func (j *jsiiProxy_LoadbalancerNodesNetworksIpAddressesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerNodesNetworksIpAddressesList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerNodesNetworksIpAddressesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (l *jsiiProxy_LoadbalancerNodesNetworksIpAddressesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerNodesNetworksIpAddressesList) Get(index *float64) LoadbalancerNodesNetworksIpAddressesOutputReference {
	if err := l.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns LoadbalancerNodesNetworksIpAddressesOutputReference

	_jsii_.Invoke(
		l,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerNodesNetworksIpAddressesList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerNodesNetworksIpAddressesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


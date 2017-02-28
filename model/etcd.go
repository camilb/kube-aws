package model

import "fmt"

type Etcd struct {
	Subnets []Subnet `yaml:"subnets,omitempty"`
}

type EtcdInstance interface {
	SubnetRef() string
	DependencyExists() bool
	DependencyRef() (string, error)
}

type etcdInstanceImpl struct {
	subnet     Subnet
	natGateway NATGateway
}

func NewEtcdInstanceDependsOnNewlyCreatedNGW(s Subnet, ngw NATGateway) EtcdInstance {
	return etcdInstanceImpl{
		subnet:     s,
		natGateway: ngw,
	}
}

func NewEtcdInstance(s Subnet) EtcdInstance {
	return etcdInstanceImpl{
		subnet: s,
	}
}

func (i etcdInstanceImpl) SubnetRef() string {
	return i.subnet.Ref()
}

func (i etcdInstanceImpl) SubnetAvailabilityZone() string {
	return i.subnet.AvailabilityZone
}

func (i etcdInstanceImpl) DependencyExists() bool {
	return i.subnet.Private && i.subnet.ManageRouteToNATGateway()
}

func (i etcdInstanceImpl) DependencyRef() (string, error) {
	// We have to wait until the route to the NAT gateway if it doesn't exist yet(hence ManageRoute=true) or the etcd node fails due to inability to connect internet
	if i.DependencyExists() {
		name := i.subnet.NATGatewayRouteLogicalName()
		return fmt.Sprintf(`"%s"`, name), nil
	}
	return "", nil
}

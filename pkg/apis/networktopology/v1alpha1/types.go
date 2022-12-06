package v1alpha1

import (
	_ "github.com/gogo/protobuf/gogoproto"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TopologyKey is the key of a OriginList in a NetworkTopology.
type TopologyKey string

// Constants for Network Topology
const (
	// NetworkTopologyRegion corresponds to "topology.kubernetes.io/region"
	NetworkTopologyRegion TopologyKey = v1.LabelTopologyRegion

	// NetworkTopologyRegion corresponds to "topology.kubernetes.io/zone"
	NetworkTopologyZone TopologyKey = v1.LabelTopologyZone

	// NetworkTopologyNetperfCosts corresponds to costs defined with measurements via the Netperf Component: "NetperfCosts"
	NetworkTopologyNetperfCosts string = "NetperfCosts"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Namespaced,shortName=nt

// NetworkTopology defines network costs in the cluster between regions and zones
type NetworkTopology struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// NetworkTopologySpec defines the zones and regions of the cluster.
	Spec NetworkTopologySpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// NetworkTopologyStatus defines the observed use.
	// +optional
	Status NetworkTopologyStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// NetworkTopologySpec represents the template of a NetworkTopology.
// +protobuf=true
type NetworkTopologySpec struct {
	// The manual defined weights of the cluster
	Weights WeightList `json:"weights,omitempty" protobuf:"bytes,1,opt,name=weights,casttype=WeightList"`

	// ConfigmapName to be used for cost calculation
	ConfigmapName string `json:"configmapName,omitempty" protobuf:"bytes,2,opt,name=configmapName"`
}

// NetworkTopologyStatus represents the current state of a Network Topology.
// +protobuf=true
type NetworkTopologyStatus struct {
	// The total number of nodes in the cluster
	NodeCount int64 `json:"nodeCount,omitempty" protobuf:"bytes,1,opt,name=nodeCount"`

	// The calculation time for the weights in the network topology CRD
	WeightCalculationTime metav1.Time `json:"weightCalculationTime,omitempty" protobuf:"bytes,2,opt,name=weightCalculationTime"`
}

// WeightList contains an array of WeightInfo objects.
// +protobuf=true
type WeightList []WeightInfo

// TopologyList contains an array of OriginInfo objects.
// +protobuf=true
type TopologyList []TopologyInfo

// WeightInfo contains information about all network costs for a given algorithm.
// +protobuf=true
type WeightInfo struct {
	// Algorithm Name for network cost calculation (e.g., userDefined)
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`

	// TopologyList owns Costs between origins
	TopologyList TopologyList `json:"topologyList,omitempty" protobuf:"bytes,2,opt,name=topologyList,casttype=TopologyList"`
}

// TopologyInfo contains information about network costs for a particular Topology Key.
// +protobuf=true
type TopologyInfo struct {
	// Topology key (e.g., "topology.kubernetes.io/region", "topology.kubernetes.io/zone").
	TopologyKey TopologyKey `json:"topologyKey,omitempty" protobuf:"bytes,1,opt,name=topologyKey"` // add as enum instead of string

	// OriginList for a particular origin.
	OriginList OriginList `json:"originList,omitempty" protobuf:"bytes,2,rep,name=originList,casttype=OriginList"`
}

// OriginList contains an array of OriginInfo objects.
// +protobuf=true
type OriginList []OriginInfo

// OriginInfo contains information about network costs for a particular Origin.
// +protobuf=true
type OriginInfo struct {
	// Name of the origin (e.g., Region Name, Zone Name).
	Origin string `json:"origin,omitempty" protobuf:"bytes,1,opt,name=origin"`

	// Costs for the particular origin.
	CostList CostList `json:"costList,omitempty" protobuf:"bytes,2,rep,name=costList,casttype=CostList"`
}

// CostList contains an array of CostInfo objects.
// +protobuf=true
type CostList []CostInfo

// CostInfo contains information about networkCosts.
// +protobuf=true
type CostInfo struct {
	// Name of the destination (e.g., Region Name, Zone Name).
	Destination string `json:"destination,omitempty" protobuf:"bytes,1,opt,name=destination"`

	// Bandwidth capacity between origin and destination.
	// +optional
	BandwidthCapacity resource.Quantity `json:"bandwidthCapacity,omitempty" protobuf:"bytes,2,opt,name=bandwidthCapacity"`

	// Bandwidth allocated between origin and destination.
	// +optional
	BandwidthAllocated resource.Quantity `json:"bandwidthAllocated,omitempty" protobuf:"bytes,3,opt,name=bandwidthAllocated"`

	// Network Cost between origin and destination (e.g., Dijkstra shortest path, etc)
	NetworkCost int64 `json:"networkCost,omitempty" protobuf:"bytes,4,opt,name=networkCost"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkTopologyList is a collection of netTopologies.
type NetworkTopologyList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is the list of NetworkTopology
	Items []NetworkTopology `json:"items"`
}

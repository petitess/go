package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

var base_url string = "https://management.azure.com/"
var sub_id string = "123"
var url string = base_url + "subscriptions/" + sub_id

func main() {

	cred := GetAzToken()

	CreateUpdateAzVnet(cred, "/resourceGroups/rg-go/providers/Microsoft.Network/virtualNetworks/vnet-dev-02?api-version=2024-05-01")

}

func GetAzToken() string {
	cred, _ := azidentity.NewDefaultAzureCredential(nil)
	var ctx = context.Background()
	policy := policy.TokenRequestOptions{
		Scopes: []string{"https://management.azure.com/.default"},
	}
	token, err := cred.GetToken(ctx, policy)
	if err != nil {
		errors.New("could not generate token. " + err.Error())
	}
	// fmt.Println(token.Token[1:15])
	return token.Token
}

type VnetModel struct {
	Location   string            `json:"location"`
	Properties VnetProperties    `json:"properties"`
	Tags       map[string]string `json:"tags"`
}

type VnetProperties struct {
	AddressSpace        AddressSpaceModel `json:"addressSpace"`
	DdosProtectionPlan  any               `json:"ddosProtectionPlan"`
	DhcpOptions         DhcpOptions       `json:"dhcpOptions"`
	EableDdosProtection bool              `json:"enableDdosProtection"`
	Subnets             []SnetModel       `json:"subnets"`
}

type DhcpOptions struct {
	DnsServers []string `json:"dnsServers"`
}

type AddressSpaceModel struct {
	AddressPrefixes []string `json:"addressPrefixes"`
}

type SnetModel struct {
	SubnetName string         `json:"name"`
	Properties SnetProperties `json:"properties"`
}

type SnetProperties struct {
	AddressPrefixes                   []string           `json:"addressPrefixes"`
	DefaultOutboundAccess             bool               `json:"defaultOutboundAccess"`
	Delegations                       []Delegation       `json:"delegations"`
	NatGateway                        any                `json:"natGateway"`
	NetworkSecurityGroup              any                `json:"networkSecurityGroup"`
	PrivateEndpointNetworkPolicies    string             `json:"privateEndpointNetworkPolicies"`
	PrivateLinkServiceNetworkPolicies string             `json:"privateLinkServiceNetworkPolicies"`
	RouteTable                        any                `json:"routeTable"`
	ServiceEndpoints                  []ServiceEndpoints `json:"serviceEndpoints"`
}

type Delegation struct {
	Name       string               `json:"name"`
	Properties DelegationProperties `json:"properties"`
}

type DelegationProperties struct {
	ServiceName string `json:"serviceName"`
}

type ServiceEndpoints struct {
	Service string `json:"service"`
}

type SubResource struct {
	Id string `json:"id"`
}

type Ddos struct {
	ResourceId string `json:"id"`
	Enable     bool   `json:"enable"`
}

func CreateUpdateAzVnet(token string, resource_path string) {

	subnets := []SnetModel{
		SnetModel{
			SubnetName: "snet01",
			Properties: SnetProperties{
				AddressPrefixes: []string{"10.10.0.0/27"},
				Delegations: []Delegation{
					Delegation{
						Name: "Dell.Storage/fileSystems",
						Properties: DelegationProperties{
							ServiceName: "Dell.Storage/fileSystems",
						},
					},
				},
				ServiceEndpoints: []ServiceEndpoints{
					ServiceEndpoints{
						Service: "Microsoft.ServiceBus",
					},
					ServiceEndpoints{
						Service: "Microsoft.Sql",
					},
				},
			},
		},
		SnetModel{
			SubnetName: "snet02",
			Properties: SnetProperties{
				AddressPrefixes:                []string{"10.10.0.128/27"},
				PrivateEndpointNetworkPolicies: "Enabled",
				NatGateway: SubResource{
					Id: "/subscriptions/123/resourceGroups/rg-go/providers/Microsoft.Network/natGateways/ng",
				},
				Delegations: []Delegation{
					Delegation{
						Name: "Dell.Storage/fileSystems",
						Properties: DelegationProperties{
							ServiceName: "Dell.Storage/fileSystems",
						},
					},
				},
			},
		},
	}

	client := &http.Client{}
	vnetProperties := VnetModel{
		Location: "swedencentral",
		Tags: map[string]string{
			"ENV": "DEV",
		},
		Properties: VnetProperties{
			DhcpOptions: DhcpOptions{
				DnsServers: []string{"1.1.1.1"},
			},
			EableDdosProtection: false,
			Subnets:             []SnetModel{},
			// FlowTimeoutInMinutes: 10,
			AddressSpace: AddressSpaceModel{
				AddressPrefixes: []string{"10.10.0.0/24"},
			},
		},
	}

	// vnetProperties.Properties.Subnets = append(vnetProperties.Properties.Subnets, subnets...)

	for _, subnet := range subnets {

		if subnet.Properties.PrivateEndpointNetworkPolicies == "" {
			subnet.Properties.PrivateEndpointNetworkPolicies = "Enabled"
		}

		if subnet.Properties.PrivateLinkServiceNetworkPolicies == "" {
			subnet.Properties.PrivateLinkServiceNetworkPolicies = "Enabled"
		}

		vnetProperties.Properties.Subnets = append(vnetProperties.Properties.Subnets, subnet)

	}

	jsonData, _ := json.Marshal(vnetProperties)
	bodyReader := bytes.NewReader(jsonData)

	req, err := http.NewRequest("PUT", url+resource_path, bodyReader)
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-type", "application/json")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fmt.Println(url + resource_path)
	fmt.Println("client: status code: ", resp.StatusCode)
	fmt.Println("client: status code: ", resp.Status)
	fmt.Println(bytes.NewBuffer(jsonData))
}

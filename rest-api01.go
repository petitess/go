package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

var base_url string = "https://management.azure.com/"
var sub_id string = "123"
var url string = base_url + "subscriptions/" + sub_id

type ResourceGroup struct {
	Name     string            `json:"name"`
	Location string            `json:"location"`
	Tags     map[string]string `json:"tags"`
	// Id       string `json:"id"`
	// Type     string `json:"type"`
	// Properties RgProperties      `json:"properties"`
}

type RgProperties struct {
	ProvisioningState string `json:"provisioningState"`
}

type Tags map[string]string

func main() {
	token := GetAzToken()
	rg := GetAzResource(token, "/resourceGroups/rg-abc-default-01?api-version=2024-11-01")
	new_tags := make(Tags)
	new_tags["dada"] = "son"
	new_tags["baba"] = "yes"
	for key, value := range rg.Tags {
		new_tags[key] = value
	}
	rg.Tags = new_tags
	CreateUpdateAzResource(token, "/resourceGroups/rg-abc-default-01?api-version=2024-11-01", rg)

	new_rg := ResourceGroup{
		Name:     "rg-abc-default-03",
		Location: "swedencentral",
		Tags: map[string]string{
			"importan": "yes",
		},
	}
	CreateUpdateAzResource(token, "/resourceGroups/rg-abc-default-03?api-version=2024-11-01", new_rg)

	DeleteAzResource(token, "/resourceGroups/rg-abc-default-03?api-version=2024-11-01")
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
	fmt.Println(token.Token[1:15])
	return token.Token
}

func GetAzResource(token string, resource_path string) ResourceGroup {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url+resource_path, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	resp, _ := client.Do(req)
	// resp, _ := http.DefaultClient.Do(req)
	fmt.Printf("client: status code: %d\n", resp.StatusCode)

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", responseData)

	var rgObject ResourceGroup
	json.Unmarshal(responseData, &rgObject)
	fmt.Println(rgObject.Name)
	fmt.Println(len(rgObject.Name))

	return rgObject
}

func CreateUpdateAzResource(token string, resource_path string, properties ResourceGroup) {
	client := &http.Client{}

	// jsonData := []byte(`{"tags": {"hell": "server!"}}`)
	// bodyReader := bytes.NewReader(jsonData)

	jsonData, _ := json.Marshal(properties)
	bodyReader := bytes.NewReader(jsonData)

	req, err := http.NewRequest("PUT", url+resource_path, bodyReader)
	req.GetBody()
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-type", "application/json")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	resp, _ := client.Do(req)
	fmt.Printf("client: status : %d\n", resp.Status)
	fmt.Printf("client: status code: %d\n", resp.StatusCode)
	fmt.Println(bytes.NewBuffer(jsonData))
}

func DeleteAzResource(token string, resource_path string) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url+resource_path, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	resp, _ := client.Do(req)
	// resp, _ := http.DefaultClient.Do(req)
	fmt.Printf("client: status code: %d\n", resp.StatusCode)
}

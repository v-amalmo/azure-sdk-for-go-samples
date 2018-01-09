package containerinstance

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure-Samples/azure-sdk-for-go-samples/helpers"
	"github.com/Azure-Samples/azure-sdk-for-go-samples/resources"
)

var (
	containerGroupName string
)

func init() {
	err := parseArgs()
	if err != nil {
		log.Fatalf("cannot parse arguments: %v", err)
	}

}

func parseArgs() error {
	err := helpers.ParseArgs()
	if err != nil {
		return fmt.Errorf("cannot parse args: %v", err)
	}

	containerGroupName = os.Getenv("AZ_CONTAINERINSTANCE_CONTAINER_GROUP_NAME")
	if !(len(containerGroupName) > 0) {
		containerGroupName = "az-samples-go-container-group-" + helpers.GetRandomLetterSequence(10)
	}

	// Container instance is not yet available in many Azure locations
	helpers.OverrideLocation([]string{
		"westus",
		"eastus",
		"westeurope",
	})
	return nil
}

func ExampleCreateContainerGroup() {
	ctx := context.Background()
	defer resources.Cleanup(ctx)

	_, err := resources.CreateGroup(ctx, helpers.ResourceGroupName())
	if err != nil {
		log.Printf("cannot create resource group: %v", err)
	}
	helpers.PrintAndLog("created resource group")

	_, err = CreateContainerGroup(ctx, containerGroupName, helpers.Location(), helpers.ResourceGroupName())
	if err != nil {
		log.Fatalf("cannot create container group: %v", err)
	}

	helpers.PrintAndLog("created container group")

	c, err := GetContainerGroup(ctx, helpers.ResourceGroupName(), containerGroupName)
	if err != nil {
		log.Fatalf("cannot get container group %v from resource group %v", containerGroupName, helpers.ResourceGroupName())
	}

	if *c.Name != containerGroupName {
		log.Fatalf("incorrect name of container group: expected %v, got %v", containerGroupName, *c.Name)
	}

	helpers.PrintAndLog("retrieved container group")

	_, err = UpdateContainerGroup(ctx, helpers.ResourceGroupName(), containerGroupName)
	if err != nil {
		log.Fatalf("cannot upate container group: %v", err)
	}

	helpers.PrintAndLog("updated container group")

	_, err = DeleteContainerGroup(ctx, helpers.ResourceGroupName(), containerGroupName)
	if err != nil {
		log.Fatalf("cannot delete container group %v from resource group %v: %v", containerGroupName, helpers.ResourceGroupName(), err)
	}

	helpers.PrintAndLog("deleted container group")

	// Output:
	// created resource group
	// created container group
	// retrieved container group
	// updated container group
	// deleted container group
}
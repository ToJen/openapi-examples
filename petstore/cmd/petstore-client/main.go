package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tojen/openapi-examples/petstore/models"

	"github.com/go-openapi/strfmt"
	"github.com/tojen/openapi-examples/petstore/client/operations"

	httptransport "github.com/go-openapi/runtime/client"
	apiclient "github.com/tojen/openapi-examples/petstore/client"
)

func main() {

	// create the transport
	transport := httptransport.New("127.0.0.1:57229", "/api", nil)

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)
	timeout := time.Minute * 5

	// add pet
	petOneName := "Pet One"
	petOne := models.NewPet{
		Name: &petOneName,
	}
	addPetParams := operations.AddPetParams{
		Pet: &petOne,
	}
	addPetParams.SetTimeout(timeout)

	addPetsOK, err := client.Operations.AddPet(&addPetParams)
	if err != nil {
		log.Fatalf("add pets error: %v\n", err)
	}
	fmt.Printf("%#v\n", addPetsOK.Payload)

	// find all pets
	findPetsParams := operations.FindPetsParams{}
	findPetsParams.SetTimeout(timeout)

	findPetsOK, err := client.Operations.FindPets(&findPetsParams)
	if err != nil {
		log.Fatalf("find pets error: %v\n", err)
	}
	fmt.Printf("%#v\n", findPetsOK.Payload)
}

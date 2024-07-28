package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// dns records represents data about a machine.
type dnsRecords struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	PublicIp    string `json:"public_ip"`
}

// dns_records slice to seed record data.
var dnsRecordss = []dnsRecords{
	{ID: "1", Description: "Google DNS Server", PublicIp: "8.8.8.8"},
	{ID: "2", Description: "Static IP Server", PublicIp: "202.87.172.154"},
}

func main() {
	router := gin.Default()
	router.GET("/DNSRecords", getDNSRecords)
	router.GET("/DNSRecords/:id", getDNSRecordByID)
	router.POST("/DNSRecords", postDNSRecords)

	router.Run("localhost:8080")
}

// getDns_records responds with the list of all entries as JSON.
func getDNSRecords(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dnsRecordss)
}

// postDNSRecords adds a new dnsrecord from JSON received in the request body.
func postDNSRecords(c *gin.Context) {
	var newDNSRecords dnsRecords

	//Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newDNSRecords); err != nil {
		return
	}

	//Add the new dnsRecord to the slice.
	dnsRecordss = append(dnsRecordss, newDNSRecords)
	c.IndentedJSON(http.StatusCreated, newDNSRecords)
}

// getDNSRecords locates the DnsRecord whose ID value matches the ID parameter sent by the client, then returns that album as a response.
func getDNSRecordByID(c *gin.Context) {
	id := c.Param("id")

	//Loop over the list of dns Records, looking for a record whose ID value matches the parameter.
	for _, a := range dnsRecordss {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "DNSRecord not found."})
}

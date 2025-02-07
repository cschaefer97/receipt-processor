package main

import (
	"net/http"

	"github.com/cschaefer97/receipt-processor/model"
	"github.com/cschaefer97/receipt-processor/scoring"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// we store all processed receipts here in a UUID:points key:value pair.
var ProcessedReceipts = make(map[uuid.UUID]int)

func processReceipt(c *gin.Context) {
	//evaluates receipt for total points earned and returns generated UUID if valid.
	var receipt model.Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//point evaluation.
	points := scoring.CheckName(receipt.Retailer) + scoring.CheckPrice(receipt.Total) + scoring.CheckNumItems(receipt.Items) +
		scoring.CheckDescription(receipt.Items) + scoring.CheckTime(receipt.PurchaseTime) + scoring.CheckDate(receipt.PurchaseDate)

	//generate UUID and store UUID with point total in key:value pair.
	id := uuid.New()
	ProcessedReceipts[id] = points
	c.JSON(http.StatusOK, gin.H{"id": id.String()})
}

func getPoints(c *gin.Context) {
	//get a receipts total points earned based on generated UUID.
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	receiptPointsById := ProcessedReceipts[id]
	c.JSON(http.StatusOK, gin.H{"points": receiptPointsById})
}

func main() {
	//create GIN router
	router := gin.Default()

	//register routes
	router.POST("/receipts/process", processReceipt)
	router.GET("/receipts/:id/points", getPoints)

	//run server
	router.Run()
}

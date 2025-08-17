package endpoint

import (
	customer_dto "architecture/internal/domain/customer/dto"
	"architecture/internal/domain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Service interface {
		GetCustomers() *[]model.Customer
		CreateCustomer(customer_dto.CreateCustomerStruct) (*model.Customer, error)
	}

	CustomerEndpoint struct {
		service Service
	}
)

func New(service Service) *CustomerEndpoint {
	return &CustomerEndpoint{
		service: service,
	}
}

func (ce CustomerEndpoint) Route(r *gin.Engine) {
	customerRoute := r.Group("/customer")
	customerRoute.GET("/", ce.GetCustomers)
	customerRoute.POST("/create", ce.CreateCustomer)
}

func (ce *CustomerEndpoint) GetCustomers(c *gin.Context) {
	customers := *ce.service.GetCustomers()

	c.JSON(200, customers)
}

func (ce *CustomerEndpoint) CreateCustomer(c *gin.Context) {
	var customer customer_dto.CreateCustomerStruct

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := ce.service.CreateCustomer(customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer cadastrada com sucesso"})
}

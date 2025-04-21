package app

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func defaultHandler(c *gin.Context) {
    response := gin.H{
        "nodes": []string{
            "Node 1",
            "Node 2",
        },
        "edges": []gin.H{
            {
                "from": "Node 1", 
                "to": "Node 2",
            },
        },
    }
    c.IndentedJSON(http.StatusOK, response)
}


func getAllServices(c *gin.Context) {
	entrypoint := c.Query("entrypoint")
	if c.Query("entrypoint") == "" {
		c.IndentedJSON(http.StatusBadRequest, "Missing required parameter `entrypoint`")
		return
	}

	kustomization := ReadAndPrintKustomization(SearchForKustomizationInPath(entrypoint))
	print(kustomization)
	c.IndentedJSON(http.StatusOK, kustomization)
}

func getService(c *gin.Context) {
	entrypoint := c.Query("entrypoint")
	if c.Query("entrypoint") == "" {
		c.IndentedJSON(http.StatusBadRequest, "Missing required parameter `entrypoint`")
		return
	}
	basePath := GetBasePathFromEntrypoint(entrypoint)
	graph := NewGraph()
	err := TraverseKustomizations(basePath, SearchForKustomizationInPath(entrypoint), graph)
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{
            "error": "Error traversing kustomizations: " + err.Error(),
        })
        return
    }
	
	graph.PrintGraph()
	jsonGraph, err := GraphToJSON(graph)

	if err != nil {
		log.Fatal("Error converting graph to JSON: ", err)
	}
	graph.PrintGraph()
	// c.IndentedJSON(http.StatusOK, json.RawMessage(jsonGraph))
	c.Data(http.StatusOK, "application/json", []byte(jsonGraph))
}

func InitializeEndpoints() {
	// Configure CORS
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"*"} // Add your frontend URL
    config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
    config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
    
	router := gin.Default()
    router.Use(cors.New(config))
	router.GET("/getAllServices", getAllServices)
	router.GET("/getService", getService)
	router.GET("/", defaultHandler)
	router.Run(":8080")
}

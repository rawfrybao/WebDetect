package detect

import (
	"log"
	"net/http"
	"webdetect/internal/logger"

	"github.com/antchfx/htmlquery"
)

func GetContent(url, xpath string) string {
	//url := "https://bestvm.cloud/index.php?rp=/store/hkbgp-a"

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		log.Println(err)
		logger.Log(err)
	}
	defer resp.Body.Close()

	// Parse the HTML response
	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		log.Println(err)
		logger.Log(err)
	}

	// Use XPath to find the desired content
	//xpath := "/html/body/section[2]/div/div/div/div[2]/div/div[1]/div/div[4]/a"
	nodes, err := htmlquery.QueryAll(doc, xpath)
	if err != nil {
		log.Println(err)
		logger.Log(err)
	}
	if len(nodes) == 0 {
		log.Println("No nodes found")
		logger.Log("No nodes found")
	}

	return htmlquery.InnerText(nodes[0])
}

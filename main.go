package main

import ( 
	"fmt"
	"strings"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
)

type Article struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

type ArticleMap map[string]Article

type ArticleTitlesJSON struct {
	Titles []string `json:"titles"`
}

var articleDataMap = make(ArticleMap)

func HomePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Homepage Endpoint")
}

func ArticlePage(writer http.ResponseWriter, request *http.Request) {

	articleNameParam := ""
	if len(request.URL.Path) > len("/articles") {
		articleNameParam = request.URL.Path[len("/articles/"):]
	}
	
	if len(articleNameParam) == 0 {
	
		writer.Header().Set("Access-Control-Allow-Origin", "*")
	
		//We are looking at the main article list page
		switch request.Method {
		case "GET":     
			
			//Build the article titles json array payload
			var articleNameStrings []string
			
			for _, currentArticle := range articleDataMap {
				articleNameStrings = append(articleNameStrings, currentArticle.Title)
			}
			
			jsonArticleTitles := &ArticleTitlesJSON{
				Titles: []string(articleNameStrings) }
			
			encJsonArticleTitles, err := json.Marshal(jsonArticleTitles)
			
			if err != nil {
				//Encoding to JSON error
				log.Fatal(err)
			}
			
			//http status OK & send the article data map JSON
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte(encJsonArticleTitles)); //Payload is the JSON array of article titles
			
			return;
		default:
			fmt.Fprintf(writer, "Sorry, only GET is supported.")
		}
		
	} else {
	
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, PUT, OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	
		//We are dealing with an individual article
		switch request.Method {
		case "GET":     
		
			titleKeyUppercase := strings.ToUpper(articleNameParam);
			
			//We check our data map to see if the article exists
			if existingArticle, ok := articleDataMap[titleKeyUppercase]; ok {
				//We have found an existing article by this title
				
				writer.Header().Set("Content-Type", "application/json")
				writer.WriteHeader(http.StatusOK) //found article, http status OK
				json.NewEncoder(writer).Encode(existingArticle) //send individual article as payload, TODO potential rewrite as json.Marshal
			} else {
				//Article title did not exist in our map, return HTTP Not Found
				writer.WriteHeader(http.StatusNotFound) //updated existing status
				//no payload
			}
			
			return;
		case "PUT":
		
			articleContentBody, err := ioutil.ReadAll(request.Body)
			
			if err != nil {
				//Request body read in error
				log.Fatal(err)
			}
			
			//Convert to uppercase to make different string casings not identified as unique
			titleKeyUppercase := strings.ToUpper(articleNameParam);

			if existingArticle, ok := articleDataMap[titleKeyUppercase]; ok {
				//We have an existing article entry
				existingArticle.Title = articleNameParam;
				existingArticle.Content = string(articleContentBody);
				articleDataMap[titleKeyUppercase] = existingArticle;
				
				writer.WriteHeader(http.StatusOK) //updated existing status
			} else {
				//We have a new article entry
				var newArticle Article
				newArticle.Title = articleNameParam;
				newArticle.Content = string(articleContentBody);
				articleDataMap[titleKeyUppercase] = newArticle;
				
				writer.WriteHeader(http.StatusCreated) //created new status
			}
			
			return;
		case "OPTIONS":
		{
			writer.WriteHeader(http.StatusOK) //updated existing status
		}
		default:
			fmt.Printf("Only GET, PUT, and OPTIONS are supported.\n")
		}
		
		return;
	}
}

func HandleRequests() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/articles", ArticlePage) //avoid redirect on no slash
	http.HandleFunc("/articles/", ArticlePage)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func main() {
	//Listen for requests
	HandleRequests()
}
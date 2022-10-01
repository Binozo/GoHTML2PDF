package endpoint

import (
	"GoHTML2PDF/pkg/chrome"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const PORT = 7524

func RegisterRoutes() {
	http.HandleFunc("/convert", convert)
}

func Listen() {
	fmt.Println("Starting Server on Port " + strconv.Itoa(PORT) + "...")
	for true {
		err := http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
		if err != nil {
			fmt.Println("An error occured while listening to Port " + strconv.Itoa(PORT) + ". Restarting...")
			continue
		}
	}
}

func convert(w http.ResponseWriter, req *http.Request) {
	err := req.ParseMultipartForm(0)
	if err != nil {
		data, _ := json.Marshal(map[string]string{"error": "Couldn't parse Form"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(data))
		return
	}

	targetHTMLData := req.FormValue("html")
	if targetHTMLData == "" {
		data, _ := json.Marshal(map[string]string{"error": "Please post a form containing \"html\" with your html"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(data))
		return
	}

	pdfData, err := chrome.ConvertHTMLToPDF(targetHTMLData)
	if err != nil {
		data, _ := json.Marshal(map[string]string{"error": "Couldn't convert your html to pdf. Maybe your html is corrupt?", "data": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(data))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(pdfData)
}

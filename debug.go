// +build debug

package logger

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Debug(msg string, args ...interface{}) {

	info := getCallerInfo()

	log.Printf(info+msg, args...)
}

func DebugDocHtml(doc *goquery.Document) {
	html, _ := goquery.OuterHtml(doc.Selection)
	Debug("document: %v", html)
}
func DebugSelectionHtml(sel *goquery.Selection) {
	html, _ := goquery.OuterHtml(sel)
	Debug("\t------- Begin Selection -------\n%v\n\t------- End Selection -------", html)
}
func DebugSelectionHtmlWithTitle(title string, sel *goquery.Selection) {
	html, _ := goquery.OuterHtml(sel)
	Debug("\t------- Begin %v Selection -------\n%v\n\t------- End %v Selection -------", title, html, title)
}
func DebugSelectionText(sel *goquery.Selection) {
	Debug("selection: %v \n\t------- End Selection -------", sel.Text())
}
func DebugSelectionTextWithTitle(title string, sel *goquery.Selection) {
	Debug("\t------- Begin %v Selection -------\n%v\n\t------- End %v Selection -------", title, sel.Text(), title)
}

func DebugHttpResponseBody(resp *http.Response) {
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		Debug("Bad http response. No response body found.")
		return
	}

	bodyString := string(bodyBytes)
	Debug("http response: %q", bodyString)
}

// +build !debug

package logger

import "github.com/PuerkitoBio/goquery"

func Debug(fmt string, args ...interface{}) {}

func DebugDocHtml(doc *goquery.Document)                               {}
func DebugSelectionHtml(sel *goquery.Selection)                        {}
func DebugSelectionHtmlWithTitle(title string, sel *goquery.Selection) {}
func DebugSelectionText(sel *goquery.Selection)                        {}
func DebugSelectionTextWithTitle(title string, sel *goquery.Selection) {}

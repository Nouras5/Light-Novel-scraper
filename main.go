package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"time"
)

func main() {

    
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Novel Link: ")
    scanner.Scan()
    link := scanner.Text()


    resp, err := http.Get(link)
    if err != nil {
        fmt.Println(err)
    }
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)

    }

    
    file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
    }

    doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
    if err != nil {
        fmt.Println("Failed to parse HTML: ", err)
    }

    

    doc.Find("p").Each(func(i int, selection *goquery.Selection) {
        text := strings.TrimSpace(selection.Text())
        if text != "" {
            fmt.Println(reverseArabicText(text)) 

            _, err = file.WriteString(text + "\n")
            if err != nil {
                fmt.Println(err)
            }
        }
    })

    time.Sleep(1 * time.Hour)
    

}

func reverseArabicText(s string) string {
    runes := []rune(s)

    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }

    return string(runes)
}

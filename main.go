package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "strings"
    "sync"
)

// Simple tools to handle string and subdomain permutations
// cat list-of-subdomain.txt | str-replace -d '.' -j ','

// build the wordlist
// cat list-of-subdomain.txt | str-replace -d '.' -n

// append the wordlist to existing subdomain
// cat list-of-subdomain.txt | str-replace -W wordlists.txt -j '.' -s
// cat list-of-subdomain.txt | str-replace -W wordlists.txt -j '.' -e
// cat list-of-subdomain.txt | str-replace -W  wordlists.txt -tld example.com

var (
    delimiterString string
    joinString      string
    data            []string
    result          []string
    wordLists       string
    stripString     string
    word            string
    tld             string
    //newLine         bool
    startOfLine bool
    endOfLine   bool
    joinNewline bool
    concurrency int
)

func main() {
    // cli arguments
    flag.IntVar(&concurrency, "c", 20, "Set the concurrency level")
    flag.StringVar(&delimiterString, "d", ",", "Delimiter char to split")
    flag.StringVar(&joinString, "j", " ", "String to join after split")
    flag.StringVar(&stripString, "strip", "", "String to strip before split")

    flag.StringVar(&word, "w", "", "word to replace")
    flag.StringVar(&wordLists, "W", "", "Wordlist to replace")
    flag.StringVar(&tld, "tld", "", "Top level domain (e.g: example.com)")

    flag.BoolVar(&startOfLine, "s", false, "Adding word at start of line")
    flag.BoolVar(&endOfLine, "e", false, "Adding word at end of line")

    flag.BoolVar(&joinNewline, "n", false, "Join the result with new line after split")
    flag.Parse()

    if wordLists != "" {
        data = ReadingLines(wordLists)
    }
    if word != "" {
        data = append(data, word)
    }

    if joinNewline {
        joinString = "\n"
    }
    if joinString == "nN" {
        joinString = "\n"
    }

    var wg sync.WaitGroup
    jobs := make(chan string, concurrency)

    for i := 0; i < concurrency; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for line := range jobs {
                if stripString != "" {
                    line = strings.Trim(line, stripString)
                }
                handleString(line)
            }
        }()
    }

    sc := bufio.NewScanner(os.Stdin)
    go func() {
        for sc.Scan() {
            line := strings.TrimSpace(sc.Text())
            if err := sc.Err(); err == nil && line != "" {
                jobs <- line
            }
        }
        close(jobs)
    }()
    wg.Wait()

    if len(result) > 0 {
        fmt.Println(strings.Join(result, joinString))
    }
}

func handleStringWithWordlist(raw string) {
    for _, replaceWord := range data {
        replaceWord = strings.TrimSpace(replaceWord)
        if replaceWord == "" {
            continue
        }

        if startOfLine {
            line := fmt.Sprintf("%s%s%s", replaceWord, joinString, raw)
            fmt.Println(line)
            continue
        }

        if endOfLine {
            line := fmt.Sprintf("%s%s%s", raw, joinString, replaceWord)
            fmt.Println(line)
            continue
        }

        line := strings.Replace(raw, replaceWord, joinString, -1)
        fmt.Println(line)
    }
}

func handleStringWithTLD(raw string) {
    if strings.Count(raw, tld) < 1 {
        return
    }
    sub := strings.Trim(strings.Split(raw, tld)[0], ".")
    if strings.Count(sub, ".") < 1 {
        return
    }
    subWords := strings.Split(sub, ".")
    for _, subWord := range subWords {
        for _, replaceWord := range data {
            replaceWord = strings.TrimSpace(replaceWord)
            if replaceWord == "" {
                continue
            }
            subdomain := strings.Replace(raw, subWord, replaceWord, -1)
            if strings.Contains(subdomain, tld) {
                fmt.Println(subdomain)
            }
        }
    }
}

func handleString(raw string) {
    if wordLists != "" {
        if tld != "" {
            handleStringWithTLD(raw)
        } else {
            handleStringWithWordlist(raw)
        }
        return
    }

    if !strings.Contains(raw, delimiterString) {
        result = append(result, raw)
        fmt.Println(strings.Join(result, joinString))
        return
    }

    result = strings.Split(raw, delimiterString)
    fmt.Println(strings.Join(result, joinString))
}

// ReadingLines Reading file and return content as []string
func ReadingLines(filename string) []string {
    var result []string
    file, err := os.Open(filename)
    if err != nil {
        return result
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        val := strings.TrimSpace(scanner.Text())
        if val == "" {
            continue
        }
        result = append(result, val)
    }

    if err := scanner.Err(); err != nil {
        return result
    }
    return result
}

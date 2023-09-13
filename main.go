package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Duration struct {
	from time.Time
	to   time.Time
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	working := make([]*Duration, 0)

	for scanner.Scan() {
		// line: YYYY-MM-DD;<name>;<email>
		line := scanner.Text()
		chunks := strings.Split(line, ";")
		datestring := chunks[0]
		date, err := time.Parse("2006-01-02", datestring)

		if err != nil {
			log.Fatal(err)
		}

		if len(working) == 0 {
			working = append(working, &Duration{date, date})
			continue
		}

		last := working[len(working)-1]

		// if the last entry is less than 45 days ago
		if date.Sub(last.to).Hours() < 24*45 {
			last.to = date
			continue
		} else {
			working = append(working, &Duration{date, date})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, d := range working {
		fmt.Printf("%s - %s\n", d.from.Format("02.01.2006"), d.to.Format("02.01.2006"))
	}
}

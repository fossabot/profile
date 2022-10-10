package main

import (
	"bufio"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Domain struct {
	Payload []string `yaml:"payload"`
}

func main() {

	res, _ := http.Get("https://github.com/dler-io/Rules/raw/main/Surge/Surge%203/Provider/Reject.list")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	scanner := bufio.NewScanner(res.Body)

	domainSuffix := make([]string, 0)
	domain := make([]string, 0)

	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "DOMAIN-SUFFIX,") {
			split := strings.Split(text, ",")
			domainSuffix = append(domainSuffix, split[1])
		} else if strings.HasPrefix(text, "DOMAIN,") {
			split := strings.Split(text, ",")
			domain = append(domain, split[1])
		}
	}

	genSurgeFile(domain, domainSuffix)
	genClashFile(domain, domainSuffix)

}

func genSurgeFile(domain, domainSuffix []string) {
	_ = os.MkdirAll("./surge/list/", 0777)
	f, _ := os.Create("./surge/list/reject.list")

	for _, s := range domain {
		_, _ = f.WriteString(s + "\n")
	}

	for _, suffix := range domainSuffix {
		_, _ = f.WriteString("." + suffix + "\n")
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
}

func genClashFile(domain, domainSuffix []string) {
	_ = os.MkdirAll("./clash/provider/", 0777)
	f, _ := os.Create("./clash/provider/reject.yaml")

	all := make([]string, 0)
	all = append(all, domain...)

	for _, suffix := range domainSuffix {
		all = append(all, "."+suffix)
	}

	p := Domain{Payload: all}
	out, _ := yaml.Marshal(&p)
	_ = os.WriteFile(f.Name(), out, 0777)

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
}

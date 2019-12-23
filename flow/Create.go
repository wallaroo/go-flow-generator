package flow

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-shadow/moment"
	"github.com/lucasjones/reggen"
	yaml "gopkg.in/yaml.v2"
)

// Create creates a dat file that follows the schema passed as parameter
func Create(configSrc, dest string, rows int) error {
	descr := &FlowDescriptor{}
	yamlFile, err := ioutil.ReadFile(configSrc)
	if err != nil {
		log.Printf("File %s doesn't exists\n %s", configSrc, err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, descr)
	if err != nil {
		log.Printf("Error unmarshalling %s \n %s", configSrc, err)
		return err
	}

	outFile, err := os.Create(dest)
	if err != nil {
		log.Fatalf("Error creating out file %s\n%s", dest, err)
	}
	defer outFile.Close()

	//HEADERS
	for _, header := range descr.Headers {
		for _, headerField := range header.Fields {
			outFile.WriteString(generate(&headerField))
		}
		outFile.WriteString("\r\n")
	}

	//CONTENT
	for i := 0; i < rows; i++ {
		for _, contentField := range descr.Content.Fields {
			outFile.WriteString(generate(&contentField))
		}
		outFile.WriteString("\r\n")
	}

	//TRAILER
	for _, trailer := range descr.Trailers {
		for _, trailerField := range trailer.Fields {
			outFile.WriteString(generate(&trailerField))
		}
		outFile.WriteString("\r\n")
	}

	return nil
}

func generate(field *FieldDescriptor) string {
	var ret string
	log.Println("Generating", field.Name)
	if field.Const != "" {
		ret = field.Const
	} else if field.Fill != "" {
		ret = strings.Repeat(field.Fill, field.End-field.Start)
	} else if field.Enum != nil && len(field.Enum) > 0 {
		rand.Seed(time.Now().UnixNano())
		ret = field.Enum[rand.Intn(len(field.Enum))]
	} else {
		switch field.Type {
		case "string":
			ret = generateString(field)
		case "number":
			ret = generateNumber(field)
		case "date":
			ret = generateDate(field)
		}
	}
	log.Println("Generated", field.Name, ret)
	return ret
}
func generateString(field *FieldDescriptor) string {
	regex := "^\\w{" + strconv.Itoa(field.Size()) + "}$"
	log.Println("generating", regex)
	ret, err := reggen.Generate(regex, 10)
	if err != nil {
		log.Fatal(err)
	}
	return ret
}

func generateNumber(field *FieldDescriptor) string {
	regex := "^\\d{" + strconv.Itoa(field.Size()) + "}$"
	log.Println("generating", regex)
	ret, err := reggen.Generate(regex, 1)
	if err != nil {
		log.Fatal(err)
	}
	return ret
}

// TODO add an optional +/- in format field to drive future / past direction of randomicity
func generateDate(field *FieldDescriptor) string {
	now := moment.New()
	rand.Seed(time.Now().UnixNano())
	now.AddDays(rand.Intn(30) * -1)
	now.AddHours(rand.Intn(24) * -1)
	now.AddMinutes(rand.Intn(60) * -1)
	now.AddSeconds(rand.Intn(60) * -1)
	return now.Format(field.Format)
}

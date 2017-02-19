package main

import (
	"encoding/csv"
	"os"

	"gopkg.in/mgo.v2"

	"fmt"

	"github.com/urfave/cli"
)

const db = "ea"

func main() {
	app := cli.NewApp()
	app.Name = "economy-importer"
	app.Commands = []cli.Command{
		{
			Name:    "transform",
			Aliases: []string{"t"},
			Usage:   "transform. Transformed csv will be placed in the same folder with json as the extension",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file",
					Usage: "--file [filename]",
				},
				cli.StringFlag{
					Name:  "type",
					Value: "Classification",
					Usage: "--type [typename]",
				},
			},
			Action: transform,
		},
	}

	app.Run(os.Args)
}

func transform(c *cli.Context) error {
	records, err := readFile(c.String("file"))

	if err != nil {
		return nil
	}

	typeArg := c.String("type")
	writeData(typeArg, getData(typeArg, records))

	return nil
}

func readFile(filePath string) ([][]string, error) {
	fileReader, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(fileReader)
	reader.Comma = ';'

	reader.Read()

	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return records, nil
}

func getData(typeArg string, records [][]string) []interface{} {
	switch typeArg {
	case "classification":
		return new(classificationMap).TransformMany(records)
	case "transaction":
		return new(transactionMap).TransformMany(records)
	default:
		panic(fmt.Sprintf("Unrecognized type: %v", typeArg))
	}
}

func writeData(typeArg string, data []interface{}) {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	collection := session.DB(db).C(typeArg + "s")
	_, err = collection.RemoveAll(nil)

	if err != nil {
		panic(err)
	}

	err = collection.Insert(data...)

	if err != nil {
		panic(err)
	}
}

![coverage](./coverage_badge.png "Coverage")
![buildstatus](https://api.travis-ci.org/morrowc/contrib.svg?branch=master "BuildStatus")
![goreportcard](https://goreportcard.com/badge/github.com/manrs-tools/contrib "Go Report Card")

A library to parse IRR database data.

Reads an IRR/RPSL file, parsing each record into a returned slice of structs.

Example usage: (parse the contents of the file: /var/tmp/radb.db)

```golang
import irr

func main() {
  fd, err := os.Open("/var/tmp/radb.db")
	if err != nil {
		fmt.Printf("Failed to open irrFile(%v): %v\n", *irrFile, err)
		return
	}
  rdr := irr.NewReader(fd)
  r, _, err := rdr.Read()
	if err != nil {
		fmt.Printf("failed to readRune: %v\n", err)
		return
	}
	err = rdr.Unread()
	if err != nil {
		fmt.Printf("failed to unRead a rune(%v): %v\n", r, err)
		return
	}
	// The file must start with a letter, all IRR records start with a letter character.
	if !irr.IsLetter(r) {
		fmt.Printf("the initial character read(%v) is not a letter, file unparsable.", r)
		return
	}

	recs, err := irr.Parse(rdr)
	if err != nil {
		if err != io.EOF {
			fmt.Printf("Error in parsing(%v): %v\n", *irrFile, err)
			return
		}
		fmt.Printf("Reached end of file with: %v records\n", len(recs))
	}

	for _, rec := range recs {
		fmt.Printf("Record type(%v):\n", rec.Type)
		for k, v := range rec.Fields {
			fmt.Printf("Key(%v)\t-> Val(%v)\n", k, v)
		}
		fmt.Println()
	}
}
```

result upon run:

~~~~
Record type(63):
Key(63)	-> Val(202.30.78.0/24)
Key(15)	-> Val(Korea Advanced Institute of Science and Technology
               373-1
               Yusong-gu
               Kusong-dong
               Taejon
               305-701
               KOREA
               REPUBLIC OF)
Key(53)	-> Val(AS372)
Key(38)	-> Val(RS-COMM_NSFNET)
Key(41)	-> Val(MAINT-AS372)
Key(11)	-> Val(nsfnet-admin@merit.edu 19941103)
Key(69)	-> Val(RADB)

Record type(63):
Key(63)	-> Val(198.62.64.0/24)
Key(15)	-> Val(Naval Command Control and Ocean Surveillance Center
               Code 914
               San Diego
               CA 92152-5000, USA)
Key(53)	-> Val(AS22)
Key(38)	-> Val(RS-COMM_NSFNET)
Key(41)	-> Val(MAINT-AS22)
Key(11)	-> Val(nsfnet-admin@merit.edu 19950501)
Key(69)	-> Val(RADB)
~~~~

Please see the CONTRIBUTING file for information about contributing to this project.

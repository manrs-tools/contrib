![coverage](./coverage_badge.png "Coverage")
![buildstatus](https://api.travis-ci.org/morrowc/contrib.svg?branch=master "BuildStatus")
![goreportcard](https://goreportcard.com/badge/github.com/manrs-tools/contrib "Go Report Card")

A library to parse IRR database data.

Reads an IRR/RPSL file, parsing each record into a Google Protobuf,
the protobuf is defined in rpsl-parser/proto/.

Example usage of this library is in:
examples/...

the example parses the contents of the file: /tmp/radb.db
and outputs records when run:
TODO(morrowc): Add a more relevant example,
and fix the example main.go to be more relevant.

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

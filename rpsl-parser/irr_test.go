//
// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package rpsl

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRead(t *testing.T) {
	tests := []struct {
		desc    string
		input   string
		want    rune
		wantErr bool
	}{{
		desc:  "Successful Read",
		input: "aut-num: 7046\n",
		want:  rune('a'),
	}, {
		desc:    "Failed Read",
		input:   "",
		wantErr: true,
	}}

	for _, test := range tests {
		r := NewReader(strings.NewReader(test.input))
		got, _, err := r.Read()
		switch {
		case err != nil && !test.wantErr:
			t.Errorf("[%v]: got error when not expecting one: %v", test.desc, err)
		case err == nil && test.wantErr:
			t.Errorf("[%v]: did not get error, expected one.", test.desc)
		case err == nil:
			if got != test.want {
				t.Errorf("[%v]: got(%+v) does not equal want(%+v)", test.desc, got, test.want)
			}
		}
	}
}

func TestUnRead(t *testing.T) {
	tests := []struct {
		desc    string
		input   string
		wantErr bool
	}{{
		desc:  "Successful Unread",
		input: "aut-num: 7046\n",
	}, {
		desc:    "Failed UnRead",
		input:   "",
		wantErr: true,
	}}

	for _, test := range tests {
		r := NewReader(strings.NewReader(test.input))
		_, _, _ = r.Read()
		err := r.Unread()
		switch {
		case err != nil && !test.wantErr:
			t.Errorf("[%v]: got error when not expecting one: %v", test.desc, err)
		case err == nil && test.wantErr:
			t.Errorf("[%v]: did not get error, expected one.", test.desc)
		}
	}
}

func TestFindKey(t *testing.T) {
	tests := []struct {
		desc        string
		input       string
		wantKey     KeyWord
		wantLiteral string
	}{{
		desc:        "Success",
		input:       "aut-num: 7046\n",
		wantKey:     AUTNUM,
		wantLiteral: "aut-num",
	}, {
		desc:        "Success: address:",
		input:       "address: asdasdasd",
		wantKey:     ADDRESS,
		wantLiteral: "address",
	}, {
		desc:        "Success: admin-c:",
		input:       "admin-c: asdasdasd",
		wantKey:     ADMINC,
		wantLiteral: "admin-c",
	}, {
		desc:        "Success: aggr-bndry:",
		input:       "aggr-bndry: asdasdasd",
		wantKey:     AGGRBNDRY,
		wantLiteral: "aggr-bndry",
	}, {
		desc:        "Success: aggr-mtd:",
		input:       "aggr-mtd: asdasdasd",
		wantKey:     AGGRMTD,
		wantLiteral: "aggr-mtd",
	}, {
		desc:        "Success: alias:",
		input:       "alias: agr1-cpt-lo0.wolcomm.net",
		wantKey:     ALIAS,
		wantLiteral: "alias",
	}, {
		desc:        "Success: as-name:",
		input:       "as-name: CRYPt-PW",
		wantKey:     ASNAME,
		wantLiteral: "as-name",
	}, {
		desc:        "Success: as-set:",
		input:       "as-set: CRYPt-PW",
		wantKey:     ASSET,
		wantLiteral: "as-set",
	}, {
		desc:        "Success: auth:",
		input:       "auth: CRYPt-PW",
		wantKey:     AUTH,
		wantLiteral: "auth",
	}, {
		desc:        "Success: certif:",
		input:       "certif: asdasdasd",
		wantKey:     CERTIF,
		wantLiteral: "certif",
	}, {
		desc:        "Success: changed:",
		input:       "changed: 010018@onsetel.co.kr",
		wantKey:     CHANGED,
		wantLiteral: "changed",
	}, {
		desc:        "Success: components:",
		input:       "components: {104.37.0.0/21^21-24}",
		wantKey:     COMPONENTS,
		wantLiteral: "components",
	}, {
		desc:        "Success: country:",
		input:       "country: AE",
		wantKey:     COUNTRY,
		wantLiteral: "country",
	}, {
		desc:        "Success: default:",
		input:       "default: to",
		wantKey:     DEFAULT,
		wantLiteral: "default",
	}, {
		desc:        "Success: descr:",
		input:       "descr: asdasdasd",
		wantKey:     DESCR,
		wantLiteral: "descr",
	}, {
		desc:        "Success: email:",
		input:       "e-mail: {",
		wantKey:     EMAIL,
		wantLiteral: "e-mail",
	}, {
		desc:        "Success: export:",
		input:       "export: {",
		wantKey:     EXPORT,
		wantLiteral: "export",
	}, {
		desc:        "Success: export-comps:",
		input:       "export-comps: {",
		wantKey:     EXPORTCOMPS,
		wantLiteral: "export-comps",
	}, {
		desc:        "Success: export-via:",
		input:       "export-via: {",
		wantKey:     EXPORTVIA,
		wantLiteral: "export-via",
	}, {
		desc:        "Success: fax-no:",
		input:       "fax-no: {",
		wantKey:     FAXNO,
		wantLiteral: "fax-no",
	}, {
		desc:        "Success: filter:",
		input:       "filter: {",
		wantKey:     FILTER,
		wantLiteral: "filter",
	}, {
		desc:        "Success: filter-set:",
		input:       "filter-set: {",
		wantKey:     FILTERSET,
		wantLiteral: "filter-set",
	}, {
		desc:        "Success: fingerpr:",
		input:       "fingerpr: 0045",
		wantKey:     FINGERPR,
		wantLiteral: "fingerpr",
	}, {
		desc:        "Success: geoidx:",
		input:       "geoidx: -10.806195,-55.459549",
		wantKey:     GEOIDX,
		wantLiteral: "geoidx",
	}, {
		desc:        "Success: holes:",
		input:       "holes: 135.84.58.0/23",
		wantKey:     HOLES,
		wantLiteral: "holes",
	}, {
		desc:        "Success: ifaddr:",
		input:       "ifaddr: 146.129.242.213",
		wantKey:     IFADDR,
		wantLiteral: "ifaddr",
	}, {
		desc:        "Success: import:",
		input:       "import: {",
		wantKey:     IMPORT,
		wantLiteral: "import",
	}, {
		desc:        "Success: import-via:",
		input:       "import-via: {",
		wantKey:     IMPORTVIA,
		wantLiteral: "import-via",
	}, {
		desc:        "Success: inetnum:",
		input:       "inetnum: 104.132.0.0",
		wantKey:     INETNUM,
		wantLiteral: "inetnum",
	}, {
		desc:        "Success: inet6num:",
		input:       "inet6num: 2001:db8::/32",
		wantKey:     INET6NUM,
		wantLiteral: "inet6num",
	}, {
		desc:        "Success: inet-rtr:",
		input:       "inet-rtr: 2001:db8::/32",
		wantKey:     INETRTR,
		wantLiteral: "inet-rtr",
	}, {
		desc:        "Success: interface:",
		input:       "interface: 2001:1900:2100::3ce2",
		wantKey:     INTERFACE,
		wantLiteral: "interface",
	}, {
		desc:        "Success key-cert:",
		input:       "key-cert: asdsad",
		wantKey:     KEYCERT,
		wantLiteral: "key-cert",
	}, {
		desc:        "Success local-as:",
		input:       "local-as: asdsad",
		wantKey:     LOCALAS,
		wantLiteral: "local-as",
	}, {
		desc:        "Success: mbrs-by-ref:",
		input:       "mbrs-by-ref: asdasdasd",
		wantKey:     MBRSBYREF,
		wantLiteral: "mbrs-by-ref",
	}, {
		desc:        "Success: member-of:",
		input:       "member-of: asdasdasd",
		wantKey:     MEMBEROF,
		wantLiteral: "member-of",
	}, {
		desc:        "Success: members:",
		input:       "members: asdasdasd",
		wantKey:     MEMBERS,
		wantLiteral: "members",
	}, {
		desc:        "Success: method:",
		input:       "method: PGP",
		wantKey:     METHOD,
		wantLiteral: "method",
	}, {
		desc:        "Success: mnt-by:",
		input:       "mnt-by: AS71-MNT",
		wantKey:     MNTBY,
		wantLiteral: "mnt-by",
	}, {
		desc:        "Success: mnt-notify:",
		input:       "mnt-nfy: AS71-MNT",
		wantKey:     MNTNFY,
		wantLiteral: "mnt-nfy",
	}, {
		desc:        "Success: mntner:",
		input:       "mntner: AS71-MNT",
		wantKey:     MNTNER,
		wantLiteral: "mntner",
	}, {
		desc:        "Success: mp-export:",
		input:       "mp-export: AS71-MNT",
		wantKey:     MPEXPORT,
		wantLiteral: "mp-export",
	}, {
		desc:        "Success: mp-filter:",
		input:       "mp-filter: AS71-MNT",
		wantKey:     MPFILTER,
		wantLiteral: "mp-filter",
	}, {
		desc:        "Success: mp-import:",
		input:       "mp-import: AS71-MNT",
		wantKey:     MPIMPORT,
		wantLiteral: "mp-import",
	}, {
		desc:        "Success: mp-members:",
		input:       "mp-members: AS71-MNT",
		wantKey:     MPMEMBERS,
		wantLiteral: "mp-members",
	}, {
		desc:        "Success: mp-peer:",
		input:       "mp-peer: AS71-MNT",
		wantKey:     MPPEER,
		wantLiteral: "mp-peer",
	}, {
		desc:        "Success: mp-peering:",
		input:       "mp-peering: AS71-MNT",
		wantKey:     MPPEERING,
		wantLiteral: "mp-peering",
	}, {
		desc:        "Success: netname:",
		input:       "netname: GOOGLE-CORP-AARHUS",
		wantKey:     NETNAME,
		wantLiteral: "netname",
	}, {
		desc:        "Success: nic-hdl:",
		input:       "nic-hdl: GOOGLE-CORP-AARHUS",
		wantKey:     NICHDL,
		wantLiteral: "nic-hdl",
	}, {
		desc:        "Success: notify:",
		input:       "notify: Aaron.hanson@kent.k12.wa.us",
		wantKey:     NOTIFY,
		wantLiteral: "notify",
	}, {
		desc:        "Success: origin:",
		input:       "origin: AS1",
		wantKey:     ORIGIN,
		wantLiteral: "origin",
	}, {
		desc:        "Success: owner:",
		input:       "owner: Aaron",
		wantKey:     OWNER,
		wantLiteral: "owner",
	}, {
		desc:        "Success: peer:",
		input:       "peer: BGP4",
		wantKey:     PEER,
		wantLiteral: "peer",
	}, {
		desc:        "Success: peering:",
		input:       "peering: AS10310",
		wantKey:     PEERING,
		wantLiteral: "peering",
	}, {
		desc:        "Success: peering-set:",
		input:       "peering-set: AS10310",
		wantKey:     PEERINGSET,
		wantLiteral: "peering-set",
	}, {
		desc:        "Success: person:",
		input:       "person: Aaron",
		wantKey:     PERSON,
		wantLiteral: "person",
	}, {
		desc:        "Success: phone:",
		input:       "phone: 0040310800700",
		wantKey:     PHONE,
		wantLiteral: "phone",
	}, {
		desc:        "Success: remarks:",
		input:       "remarks: asdadasd",
		wantKey:     REMARKS,
		wantLiteral: "remarks",
	}, {
		desc:        "Success: role:",
		input:       "role: AAPT",
		wantKey:     ROLE,
		wantLiteral: "role",
	}, {
		desc:        "Success: roa-uri:",
		input:       "roa-uri: AAPT",
		wantKey:     ROAURI,
		wantLiteral: "roa-uri",
	}, {
		desc:        "Success: route:",
		input:       "route: 1.0.0.0/24",
		wantKey:     ROUTE,
		wantLiteral: "route",
	}, {
		desc:        "Success: route6:",
		input:       "route6: 2001:db8::/32",
		wantKey:     ROUTE6,
		wantLiteral: "route6",
	}, {
		desc:        "Success: route-set:",
		input:       "route-set: 1.0.0.0/24",
		wantKey:     ROUTESET,
		wantLiteral: "route-set",
	}, {
		desc:        "Success: rs-in:",
		input:       "rs-in: 1.0.0.0/24",
		wantKey:     RSIN,
		wantLiteral: "rs-in",
	}, {
		desc:        "Success: rs-out:",
		input:       "rs-out: 1.0.0.0/24",
		wantKey:     RSOUT,
		wantLiteral: "rs-out",
	}, {
		desc:        "Success: rtr-set:",
		input:       "rtr-set: 1.0.0.0/24",
		wantKey:     RTRSET,
		wantLiteral: "rtr-set",
	}, {
		desc:        "Success: route-set:",
		input:       "route-set: 1.0.0.0/24",
		wantKey:     ROUTESET,
		wantLiteral: "route-set",
	}, {
		desc:        "Success: source:",
		input:       "source: radb",
		wantKey:     SOURCE,
		wantLiteral: "source",
	}, {
		desc:        "Success: status:",
		input:       "status: ASSIGNED",
		wantKey:     STATUS,
		wantLiteral: "status",
	}, {
		desc:        "Success: tech-c:",
		input:       "tech-c: abuse@adnc.com",
		wantKey:     TECHC,
		wantLiteral: "tech-c",
	}, {
		desc:        "Success: trouble:",
		input:       "trouble: abuse@adnc.com",
		wantKey:     TROUBLE,
		wantLiteral: "trouble",
	}, {
		desc:        "Success: upd-to:",
		input:       "upd-to: abuse@adnc.com",
		wantKey:     UPDTO,
		wantLiteral: "upd-to",
	}, {
		desc:        "Success: *xxe:",
		input:       "*xxe: abuse@adnc.com",
		wantKey:     XXE,
		wantLiteral: "*xxe",
	}, {
		desc:        "Success: *xxner:",
		input:       "*xxner: abuse@adnc.com",
		wantKey:     XXNER,
		wantLiteral: "*xxner",
	}, {
		desc:        "Success: *xx-num:",
		input:       "*xx-num: abuse@adnc.com",
		wantKey:     XXNUM,
		wantLiteral: "*xx-num",
	}, {
		desc:        "Success: *xxring-set:",
		input:       "*xxring-set: abuse@adnc.com",
		wantKey:     XXRINGSET,
		wantLiteral: "*xxring-set",
	}, {
		desc:        "Success: *xxset:",
		input:       "*xxset: abuse@adnc.com",
		wantKey:     XXSET,
		wantLiteral: "*xxset",
	}, {
		desc:        "Success: *xxson:",
		input:       "*xxson: abuse@adnc.com",
		wantKey:     XXSON,
		wantLiteral: "*xxson",
	}, {
		desc:        "Success: *xxte:",
		input:       "*xxte: abuse@adnc.com",
		wantKey:     XXTE,
		wantLiteral: "*xxte",
	}, {
		desc:        "Success: *xxte6:",
		input:       "*xxte6: abuse@adnc.com",
		wantKey:     XXTE6,
		wantLiteral: "*xxte6",
	}, {
		desc:        "Success: *xxte-set:",
		input:       "*xxte-set: abuse@adnc.com",
		wantKey:     XXTESET,
		wantLiteral: "*xxte-set",
	}, {
		desc:        "Success ILLEGAL",
		input:       "aut: 7046\nas-name: UUNET Customer ASN\n\n",
		wantKey:     ILLEGAL,
		wantLiteral: "aut",
	}, {
		desc:        "Success EOF in key",
		input:       "aut",
		wantKey:     EOF,
		wantLiteral: "",
	}}

	for _, test := range tests {
		r := NewReader(strings.NewReader(test.input))
		gotKey, gotLiteral := r.findKey()

		if gotKey != test.wantKey {
			t.Errorf("[%v]: gotKey/wantKey do not match: %v/%v", test.desc, gotKey, test.wantKey)
		}
		if gotLiteral != test.wantLiteral {
			t.Errorf("[%v]: gotLiteral/wantLiteral do not match: %v/%v", test.desc, gotLiteral, test.wantLiteral)
		}
	}
}

func TestReadValue(t *testing.T) {
	tests := []struct {
		desc      string
		input     string
		wantValue string
		endRecord bool
		wantErr   bool
	}{{
		desc:      "Success",
		input:     "a new autnum\naut-num:",
		wantValue: "a new autnum",
	}, {
		desc:      "Success EndRecord",
		input:     "a new autnum\n\naut-num:",
		wantValue: "a new autnum",
		endRecord: true,
	}, {
		desc:    "Fail EOF",
		input:   "a new autnum\n",
		wantErr: true,
	}}

	for _, test := range tests {
		r := NewReader(strings.NewReader(test.input))
		gotValue, gotBool, err := r.readValue()

		switch {
		case err != nil && !test.wantErr:
			t.Errorf("[%v]: got error when not expecting one: %v", test.desc, err)
		case err == nil && test.wantErr:
			t.Errorf("[%v]: did not get error, expected one.", test.desc)
		case err == nil:
			if gotValue != test.wantValue {
				t.Errorf("[%v]: got/want mismatch: %v/%v", test.desc, gotValue, test.wantValue)
			}
			if gotBool != test.endRecord {
				t.Errorf("[%v]: gotBool/wantBool mismatch: %v/%v", test.desc, gotBool, test.endRecord)
			}
		}
	}

}

func TestPeek(t *testing.T) {
	tests := []struct {
		desc    string
		input   string
		want    rune
		wantErr bool
	}{{
		desc:  "Success",
		input: "s",
		want:  rune('s'),
	}, {
		desc:  "Sucess EOF read",
		input: "",
		want:  eof,
	}}

	for _, test := range tests {
		r := NewReader(strings.NewReader(test.input))
		got := r.Peek()

		if got != test.want {
			t.Errorf("[%v]: got/want do not match: %v/%v", test.desc, got, test.want)
		}
	}
}

func TestInitRecord(t *testing.T) {
	tests := []struct {
		desc    string
		input   string
		want    Record
		wantErr bool
	}{{
		desc:  "Successful initialization",
		input: "aut-num: 7046\nas-name: uunet customer\n\n",
		want:  Record{Type: AUTNUM, Fields: map[KeyWord]string{AUTNUM: "7046"}},
	}, {
		desc:    "Fail IllegalKey",
		input:   "foo: bar\nbaz: bling\n",
		wantErr: true,
	}, {
		desc:    "Fail: No Colon EOF",
		input:   "aut-num",
		wantErr: true,
	}, {
		desc:    "Fail No Value EOF",
		input:   "aut-num: \n",
		wantErr: true,
	}, {
		desc:    "Fail Value is end of record",
		input:   "aut-num: bar\n\n",
		wantErr: true,
	}}

	for _, test := range tests {
		r := NewReader(strings.NewReader(test.input))
		got, err := r.initRecord()
		switch {
		case err != nil && !test.wantErr:
			t.Errorf("[%v]: got error when not expecting one: %v", test.desc, err)
		case err == nil && test.wantErr:
			t.Errorf("[%v]: did not get error, expected one.", test.desc)
		case err == nil:
			if !reflect.DeepEqual(*got, test.want) {
				t.Errorf("[%v]: got(%+v) does not equal want(%+v)", test.desc, *got, test.want)
			}
		}
	}
}

func TestConsumeWS(t *testing.T) {
	tests := []struct {
		desc    string
		input   string
		wantErr bool
	}{{
		desc:  "Success",
		input: " ",
	}, {
		desc:  "Sucess doublespace",
		input: "  ",
	}, {
		desc: "Sucess colon and whitespace (tab)",
		input: ":	f",
	}, {
		desc:  "Success colon and no whitespace",
		input: ":f",
	}}

	for _, test := range tests {
		r := NewReader(strings.NewReader(test.input))
		got := r.ConsumeLeadingWS()

		switch {
		case got != nil && !test.wantErr:
			t.Errorf("[%v]: got error when not expecting one: %v", test.desc, got)
		case got == nil && test.wantErr:
			t.Errorf("[%v]: did not get error, expected one.", test.desc)
		}
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		desc    string
		file    string
		want    *Record
		rec     int
		recs    int
		wantErr bool
	}{{
		desc: "Success with single values only",
		file: "aut-num.txt",
		want: &Record{Type: AUTNUM, Fields: map[KeyWord]string{
			AUTNUM:  "AS372",
			ASNAME:  "UNSPECIFIED",
			DESCR:   "Nasa Science Network (FIX-West)",
			ADMINC:  "Not available",
			TECHC:   "See MAINT-AS372",
			MNTBY:   "MAINT-AS372",
			CHANGED: "DB-admin@merit.edu 19950201",
			SOURCE:  "RADB",
		}},
		rec:  0,
		recs: 1,
	}, {
		desc: "Success - with double value (members)",
		file: "as-set.txt",
		want: &Record{Type: ASSET, Fields: map[KeyWord]string{
			ASSET:   "AS-CENTRILOGIC-UK:AS-CUSTOMERS",
			DESCR:   "CentriLogic (UK Network) Customer ASNs",
			MEMBERS: "AS204018 AS33459",
			ADMINC:  "CentriLogic IP Tech",
			MNTBY:   "MAINT-AS19693",
			CHANGED: "IPtech@centrilogic.com 20180614  #14:42:22Z",
			TECHC:   "CentriLogic IP Tech",
			NOTIFY:  "iptech@centrilogic.com",
			SOURCE:  "RADB",
		}},
		rec:  0,
		recs: 1,
	}, {
		desc: "Success with two records read",
		file: "two-records.txt",
		want: &Record{Type: AUTNUM, Fields: map[KeyWord]string{
			AUTNUM:  "AS1263",
			ASNAME:  "NSN-TEST-AS",
			DESCR:   "NSN-TEST-AS",
			ADMINC:  "Not available",
			TECHC:   "See MAINT-AS1263",
			MNTBY:   "MAINT-AS1263",
			CHANGED: "DB-admin@merit.edu 19950201",
			SOURCE:  "RADB",
		}},
		rec:  1,
		recs: 2,
	}, {
		desc: "Success with a plus-sign continuation character",
		file: "route-set.txt",
		want: &Record{Type: ROUTESET, Fields: map[KeyWord]string{
			ROUTESET:  "RS-CTB-NOVA",
			MEMBERS:   "138.0.96.0/22,\n            143.0.128.0/22,\n            167.249.164.0/22,\n            168.195.104.0/22,\n            192.141.100.0/22,\n+\n            192.141.100.0/24",
			MPMEMBERS: "2804:2138::/32,\n            2804:2c04::/32,\n            2804:42bc::/32",
			DESCR:     "CTB's customer",
			REMARKS:   "ASN 264543",
			MNTBY:     "MAINT-AS36678",
			CHANGED:   "jiangz@ctamericas.com 20180614  #21:22:22Z",
			SOURCE:    "RADB",
		}},
		rec:  0,
		recs: 1,
	}}

	for _, test := range tests {
		fd, err := os.Open(filepath.Join("testdata", test.file))
		if err != nil {
			if !test.wantErr {
				t.Errorf("[%v]: wanted an error and did not get one.", test.desc)
			}
			continue
		}

		r := NewReader(fd)
		rc := make(chan *Record, 10)

		// Run the parse routine, pull the Records off the channel.
		Parse(r, rc)
		close(rc)

		if strings.HasSuffix(fmt.Sprintf("%s", err), "EOF") {
			err = nil
		}

		var got []*Record
		if err == nil {
			for i := 0; i < test.recs; i++ {
				got = append(got, <-rc)
			}
		}

		switch {
		case err != nil && !test.wantErr:
			t.Errorf("[%v]: got error when not expecting one: %v", test.desc, err)
		case err == nil && test.wantErr:
			t.Errorf("[%v]: did not get error, expected one.", test.desc)
		case err == nil:
			if !cmp.Equal(got[test.rec], test.want) {
				t.Errorf("[%v]: got/want mismatch:\n%s", test.desc, cmp.Diff(got[test.rec], test.want))
			}
		}
	}
}

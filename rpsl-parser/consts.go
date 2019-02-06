//
// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Hold all the constants used in the IRR package here, less clutter in the main IRR conde.
package rpsl

type KeyWord int

const (
	eof        = rune(0)
	octothorpe = rune('#')
	colon      = rune(':')

	// These keywords must be synced with the content in findKey.
	// Additionally, there will have to be a map of these to values kept/created.
	ADDRESS KeyWord = iota
	ADMINC
	AGGRBNDRY
	AGGRMTD
	ALIAS
	ASNAME
	ASSET
	AUTH
	AUTNUM
	CERTIF
	CHANGED
	COMPONENTS
	COUNTRY
	DEFAULT
	DESCR
	EMAIL
	EOF
	EXPORT
	EXPORTCOMPS
	EXPORTVIA
	FAXNO
	FILTER
	FILTERSET
	FINGERPR
	GEOIDX
	HOLES
	IFADDR
	ILLEGAL
	IMPORT
	IMPORTVIA
	INET6NUM
	INETNUM
	INETRTR
	INTERFACE
	KEYCERT
	LOCALAS
	MBRSBYREF
	MEMBEROF
	MEMBERS
	METHOD
	MNTBY
	MNTNER
	MNTNFY
	MPEXPORT
	MPFILTER
	MPIMPORT
	MPMEMBERS
	MPPEER
	MPPEERING
	NETNAME
	NICHDL
	NOTIFY
	ORIGIN
	OWNER
	PEER
	PEERING
	PEERINGSET
	PERSON
	PHONE
	REMARKS
	ROAURI
	ROLE
	ROUTE
	ROUTE6
	ROUTESET
	RSIN
	RSOUT
	RTRSET
	SOURCE
	STATUS
	TECHC
	TROUBLE
	UPDTO
	XXE
	XXNER
	XXNUM
	XXRINGSET
	XXSET
	XXSON
	XXTE
	XXTE6
	XXTESET
)

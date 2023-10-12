#!/usr/bin/env python
# MANRS API - developed by acidvegas in python (https://git.acid.vegas/manrs)

'''
Source             : https://www.manrs.org/
Request API Access : https://www.manrs.org/resources/api
API Doumentation   : https://manrs.stoplight.io/docs/manrs-public-api
'''

import http.client

class API(object):
	def __init__(self, api_key: str, dev: bool=False):
		'''
		MANRS API class.
		
		:param api_key: Your MANRS API key
		:param dev: Whether or not to use the development API (default: False)
		'''
		self.api_key = api_key
		self.dev = dev

	def call(self, endpoint: str) -> dict:
		'''
		Makes a call to the MANRS API

		:param endpoint: The endpoint you would like to query
		'''
		headers = {'Accept': 'application/json', 'Authorization': 'Bearer ' + self.api_key}
		if not self.dev:
			conn = http.client.HTTPSConnection('api.manrs.org')
		else:
			http.client.HTTPSConnection('api-dev.manrs.org')
		conn.request('GET', endpoint, headers=headers)
		res = conn.getresponse()
		data = res.read()
		return data.decode()

	def roa_by_asn(self, asn: str) -> dict:
		'''
		Retrieve data about ROAs by ASN
		
		:param asn: The ASN you would like to query either as a number or in AS12345 format
		'''
		return self.call('/roas/asn/'+asn)

	def roa_by_country(self, country: str) -> dict:
		'''
		Retrieve ROAs by country

		:param country: Two-letter ISO code for the country you wish to query
		'''
		return self.call('/roas/country/'+country)

	def asn_info(self) -> dict:
		'''Get a list of all known ASNs and info about them (e.g. holder name)'''
		return self.call('/asns/info')

	def ixp_info(self) -> dict:
		'''Query for info on IXPs'''
		return self.call('/ixps/info')

	def conformance_by_cdn(self) -> dict:
		'''List conformance for all CDNs that are participanting in MANRS'''
		return self.call('/conformance/cdns')

	def conformace_by_ixp(self) -> dict:
		'''List conformance for all IXPs that are participanting in MANRS'''
		return self.call('/conformance/ixps')
			
	def conformance_by_network_operator(self) -> dict:
		'''List conformance for all Network Operators that are participanting in MANRS'''
		return self.call('/conformance/net-ops')

	def conformance_by_vendor(self) -> dict:
		'''List conformance for all equipment vendors that are participanting in MANRS'''
		return self.call('/conformance/vendors')
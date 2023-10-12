# MANRS API
![](logo.png)

This is a Python wrapper for the [MANRS](https://www.manrs.org/) API, designed to simplify the process of making requests to the MANRS Public API. This class assists in querying data related to routing security, such as ROAs by ASN or country, ASN info, IXP info, and conformance details for CDNs, IXPs, network operators, and equipment vendors.

The official documentation for the MANRS API can be found [here](https://manrs.stoplight.io/docs/manrs-public-api)

## Features:
- Query ROAs by ASN or country.
- Retrieve information about all known ASNs and their holders.
- Fetch data about Internet Exchange Points *(IXPs)*.
- Get conformance details for different entities participating in MANRS.

## How to Use:
**Note:** You must [request access](https://www.manrs.org/resources/api) to get an API key!

1. **Initialization**: Instantiate the `MANRS` class with your API key.
```python
import manrs

api = manrs.API('YOUR_API_KEY')
```

2. **Making Calls:** Use the provided methods to make calls to the MANRS API. For instance, to get ROAs by ASN:
```python
response = api.roa_by_asn('AS16661')
```

3. **Development Mode:** If you're working in a development environment, set the `dev` flag to `True` during initialization.
```
api = manrs.API('YOUR_API_KEY', dev=True)
```

___

###### Mirrors
[acid.vegas](https://git.acid.vegas/manrs) • [GitHub](https://github.com/acidvegas/manrs) • [GitLab](https://gitlab.com/acidvegas/manrs) • [SuperNETs](https://git.supernets.org/acidvegas/manrs)

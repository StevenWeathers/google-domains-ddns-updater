# google-domains-ddns-updater

Little utility application to update GoogleDomains Dynamic DNS entries for a list of hostnames/domains.

Provide the application your Dynamic DNS username/password, and a comma seperated list of domains.


Running with Docker

```
docker build . -t gddu

docker run -v hostnames.json:/data/hostnames.json gddu
```


hostnames.json file format
```
{
    "hostnames": [
        {
            "domain": "google.com",
            "username": "google",
            "password": "elgoog"
        }
    ]
}
```
# google-domains-ddns-updater

API Server to update GoogleDomains Dynamic DNS entries for a list of domains.

Runs a cronlike job to make a request to google for each domain entered.

Exposes APIs to manage the domains entered, as well as manually trigger the job when needed.

## Running with Docker

```
docker build . -t gddu

docker run -v hostnames.json:/data/hostnames.json gddu
```

## hostnames.json file format
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

## API Endpoints (unsecured, unvalidated)
```
[GET]
/hostnames

[GET]
/hostnames/domain.com

[POST]
/hostnames
{
    "domain": "google.com",
    "username": "google",
    "password": "elgoog"
}

[PUT]
/hostnames/domain.com
{
    "username": "google",
    "password": "elgoog"
}

[DELETE]
/hostnames/domain.com
```

# google-domains-ddns-updater

API Server to update [GoogleDomains Dynamic DNS](https://support.google.com/domains/answer/6147083?hl=en) entries for a list of domains.

Runs a cron job to make a request to google for each domain entered.

Exposes APIs to manage the domains entered, as well as manually trigger the job when needed.

## Running with Docker

```
docker build . -t gddu

docker run -v hostnames.json:/data/hostnames.json -p 8000:8000 gddu

or on windows

docker run -v ${PWD}\hostnames.json:/data/hostnames.json -p 8000:8000 --name gddu gddu
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
/api/hostnames

[GET]
/api/hostnames/domain.com

[POST]
/api/hostnames
{
    "domain": "google.com",
    "username": "google",
    "password": "elgoog"
}

[PUT]
/api/hostnames/domain.com
{
    "username": "google",
    "password": "elgoog"
}

[DELETE]
/api/hostnames/domain.com

[GET]
/api/triggerUpdate
```

# Development

## Go App
```
currently using Docker, local directions to follow...
```

## Vue UI

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Run your tests
```
npm run test
```

### Lints and fixes files
```
npm run lint
```
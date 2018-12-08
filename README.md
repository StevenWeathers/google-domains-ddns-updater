# google-domains-ddns-updater

API Server to update [GoogleDomains Dynamic DNS](https://support.google.com/domains/answer/6147083?hl=en) entries for a list of domains.

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

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
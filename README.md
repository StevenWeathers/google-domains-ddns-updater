# google-domains-ddns-updater

Server to update [GoogleDomains Dynamic DNS](https://support.google.com/domains/answer/6147083?hl=en) entries for a list of domains.

Runs a cron job to make a request to google for each domain entered.

Exposes APIs to manage the domains entered, as well as manually trigger the job when needed.

Also exposes a WebUI (unprotected) that utilizes the APIs to provide easy management.

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

## WebUI
```
visit http://localhost:8000 in your browser
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

# Running in production

## Use latest docker image

```
docker pull stevenweathers/google-domains-ddns-updater
```

## Use latest released binary

[![](https://img.shields.io/github/v/release/stevenweathers/google-domains-ddns-updater?include_prereleases)](https://github.com/StevenWeathers/google-domains-ddns-updater/releases/latest)

# Running locally

## Building and running with Docker (preferred solution)

### Using Docker Compose

```
docker-compose up --build
```

### Using Docker without Compose

This solution will require you to pass environment variables or setup the config file, as well as setup and manage the DB yourself.

```
docker build ./ -f ./build/Dockerfile -t gddu:latest
docker run --publish 8080:8080 --name gddu gddu:latest
```

## Building

### Install dependencies
```
go get
npm install
```

## Build with Make
```
make build
```
### OR manual steps

### Build static assets
```
npm run build
```

### Build for current OS
```
go build
```

# Adding new Locale's
Using svelte-i18n **Google Domains DDNS Updater** now supports Locale selection on the UI (Default en-US)

Adding new locale's involves just a couple of steps.

1. First add the locale dictionary json files in ```web/public/lang/default/``` and ```web/public/lang/friendly/``` by copying the en.json and just changing the values of all keys
1. Second, the locale will need to be added to the locales list used by switcher component in ```web/config.js``` ```locales``` object

# Configuration
Thunderdome may be configured through environment variables or via a yaml file `config.yaml`
located in one of:

* `/etc/google-domains-ddns-updater/`
* `$HOME/.config/google-domains-ddns-updater/`
* Current working directory

The following configuration options exists:

| Option                     | Environment Variable | Description                                | Default Value           |
| -------------------------- | -------------------- | ------------------------------------------ | ------------------------|
| `http.port`                | PORT                 | Which port to listen for HTTP connections. | 8000 |
| `http.path_prefix`         | PATH_PREFIX          | Prefix added to all application urls for shared domain use, in format of `/{prefix}` e.g. `/gddu` | |
| `config.toast_timeout`     | CONFIG_TOAST_TIMEOUT | Number of milliseconds before notifications are hidden. | 1000 |
| `config.default_locale`   | CONFIG_DEFAULT_LOCALE | The default locale (language) for the UI | en |
| `config.json_path`   | JSONPATH | The path where the json file is stored | data/hostnames.json |
| `config.cadence`   | CADENCE | The cadence at which the cron will run | @hourly |

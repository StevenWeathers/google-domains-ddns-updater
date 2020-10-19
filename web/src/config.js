const locales = {
    en: 'English',
}

const { PathPrefix, DefaultLocale: fallbackLocale } = appConfig

const appRoutes = {
    landing: `${PathPrefix}/`,
}

export { locales, fallbackLocale, appRoutes, PathPrefix }

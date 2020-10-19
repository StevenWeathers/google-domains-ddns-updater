const appUrl = process.env.APP_URL || 'http://localhost:8080/'

describe('Google Domains DDNS Updater App', () => {
    beforeAll(async () => {
        await page.goto(appUrl) // @TODO - make this configurable for the endpoint
    })
})

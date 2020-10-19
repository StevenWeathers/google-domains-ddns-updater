<script>
    import Navaid from 'navaid'
    import { onDestroy } from 'svelte'

    import { _, locale, setupI18n, isLocaleLoaded } from './i18n'
    import { appRoutes } from './config'
    import LocaleSwitcher from './components/LocaleSwitcher.svelte'
    import Notifications from './components/Notifications.svelte'
    import HollowButton from './components/HollowButton.svelte'
    import Landing from './pages/Landing.svelte'
    import apiclient from './apiclient.js'

    setupI18n()

    const { AppVersion, PathPrefix } = appConfig

    const footerLinkClasses = 'no-underline text-teal-500 hover:text-teal-800'

    let notifications

    let currentPage = {
        route: Landing,
        params: {},
    }

    const router = Navaid('/')
        .on(appRoutes.landing, () => {
            currentPage = {
                route: Landing,
                params: {},
            }
        })
        .listen()

    const xfetch = apiclient(handle401)

    function handle401() {}

    function triggerJob() {
        xfetch('/api/triggerUpdate')
            .then(function() {
                notifications.success('Job triggered successfully')
            })
            .catch(function(error) {
                notifications.danger(
                    'Error encountered attempting to trigger job',
                )
            })
    }

    onDestroy(router.unlisten)
</script>

<style>
    :global(.nav-logo) {
        max-height: 3.75rem;
    }
</style>

<Notifications bind:this="{notifications}" />
{#if isLocaleLoaded}
    <nav
        class="flex items-center justify-between flex-wrap bg-white p-6"
        role="navigation"
        aria-label="main navigation">
        <div class="flex items-center flex-shrink-0 mr-6">
            <a href="{appRoutes.landing}">{$_('appName')}</a>
        </div>
        <div class="text-right mt-4 md:mt-0">
            <HollowButton color="red" onClick="{triggerJob}">
                Trigger Job
            </HollowButton>
            <LocaleSwitcher
                selectedLocale="{$locale}"
                on:locale-changed="{e => setupI18n({
                        withLocale: e.detail,
                    })}" />
        </div>
    </nav>

    <svelte:component
        this="{currentPage.route}"
        {...currentPage.params}
        {notifications}
        {router}
        {xfetch} />

    <footer class="p-6 text-center">
        <a
            href="https://github.com/StevenWeathers/google-domains-ddns-updater"
            class="{footerLinkClasses}">
            {$_('appName')}
        </a>
        {@html $_('footer.authoredBy', {
            values: {
                authorOpen: `<a href="http://stevenweathers.com" class="${footerLinkClasses}">`,
                authorClose: `</a>`,
            },
        })}
        {@html $_('footer.license', {
            values: {
                licenseOpen: `<a href="https://opensource.org/licenses/MIT" class="${footerLinkClasses}">`,
                licenseClose: `</a>`,
            },
        })}
        <br />
        {@html $_('footer.poweredBy', {
            values: {
                svelteOpen: `<a href="https://svelte.dev/" class="${footerLinkClasses}">`,
                svelteClose: `</a>`,
                goOpen: `<a href="https://golang.org/" class="${footerLinkClasses}">`,
                goClose: `</a>`,
            },
        })}
        <div class="text-sm text-gray-500">
            {$_('appVersion', { values: { version: AppVersion } })}
        </div>
    </footer>
{:else}
    <p>Loading...</p>
{/if}

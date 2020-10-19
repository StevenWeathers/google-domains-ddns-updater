<script>
    import SolidButton from '../components/SolidButton.svelte'
    import CloseIcon from './icons/CloseIcon.svelte'

    export let xfetch
    export let notifications
    export let hostnameToEdit = {}
    export let toggleModal = () => {}
    export let handleSave = () => {}

    let domain = hostnameToEdit.domain || ''
    let username = hostnameToEdit.username || ''
    let password = hostnameToEdit.password || ''

    const formSubmit =
        typeof hostnameToEdit.domain !== 'undefined'
            ? saveHostname
            : createHostname

    function createHostname(e) {
        e.preventDefault()
        const body = {
            domain,
            username,
            password,
        }

        if (!/^([a-z0-9]+(-[a-z0-9]+)*\.)+[a-z]{2,}$/.test(domain)) {
            notifications.danger('Domain is not valid')
            return
        }

        xfetch('/api/hostnames', { body })
            .then(function() {
                handleSave()
            })
            .catch(function(error) {
                notifications.danger(
                    `Error encountered creating new question: ${error.message}`,
                )
            })
    }

    function saveHostname(e) {
        e.preventDefault()
        const body = {
            username,
            password,
        }

        if (!/^([a-z0-9]+(-[a-z0-9]+)*\.)+[a-z]{2,}$/.test(domain)) {
            notifications.danger('Domain is not valid')
            return
        }

        xfetch(`/api/hostnames/${domain}`, { body, method: 'PUT' })
            .then(function() {
                handleSave()
            })
            .catch(function(error) {
                notifications.danger(
                    `Error encountered saving question changes: ${error.message}`,
                )
            })
    }

    $: submitDisabled = domain === '' || username === '' || password === ''
</script>

<div
    class="fixed inset-0 flex items-center z-40 max-h-screen overflow-y-scroll">
    <div class="fixed inset-0 bg-gray-900 opacity-75"></div>

    <div
        class="relative mx-4 md:mx-auto w-full md:w-2/3 lg:w-3/5 xl:w-1/2 z-50
        max-h-full">
        <div class="py-8">
            <div class="shadow-xl bg-white rounded-lg p-4 xl:p-6 max-h-full">
                <div class="flex justify-end mb-2">
                    <button
                        aria-label="close"
                        on:click="{toggleModal({})}"
                        class="text-gray-800">
                        <CloseIcon />
                    </button>
                </div>

                <form on:submit="{formSubmit}" name="saveHostname">
                    <div class="mb-4">
                        <label
                            class="block text-gray-700 text-sm font-bold mb-2"
                            for="domain">
                            Domain
                        </label>
                        <div class="control">
                            <input
                                name="domain"
                                bind:value="{domain}"
                                placeholder="Enter a domain"
                                class="bg-gray-200 border-gray-200 border-2
                                appearance-none rounded w-full py-2 px-3
                                text-gray-700 leading-tight focus:outline-none
                                focus:bg-white focus:border-orange-500"
                                id="domain"
                                required />
                        </div>
                    </div>

                    <div class="mb-4">
                        <label
                            class="block text-gray-700 text-sm font-bold mb-2"
                            for="username">
                            Username
                        </label>
                        <div class="control">
                            <textarea
                                name="username"
                                bind:value="{username}"
                                placeholder="Enter a username"
                                class="bg-gray-200 border-gray-200 border-2
                                appearance-none rounded w-full py-2 px-3
                                text-gray-700 leading-tight focus:outline-none
                                focus:bg-white focus:border-orange-500"
                                id="username"
                                required></textarea>
                        </div>
                    </div>

                    <div class="mb-4">
                        <label
                            class="block text-gray-700 text-sm font-bold mb-2"
                            for="password">
                            Password
                        </label>
                        <div class="control">
                            <textarea
                                name="password"
                                bind:value="{password}"
                                placeholder="Enter a password"
                                type="password"
                                class="bg-gray-200 border-gray-200 border-2
                                appearance-none rounded w-full py-2 px-3
                                text-gray-700 leading-tight focus:outline-none
                                focus:bg-white focus:border-orange-500"
                                id="password"
                                required></textarea>
                        </div>
                    </div>

                    <div class="text-right">
                        <SolidButton type="submit" disabled="{submitDisabled}">
                            Save
                        </SolidButton>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>

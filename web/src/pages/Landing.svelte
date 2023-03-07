<script>
    import PageLayout from '../components/PageLayout.svelte'
    import HollowButton from '../components/HollowButton.svelte'
    import HostnameForm from '../components/HostnameForm.svelte'

    export let xfetch
    export let notifications

    let hostnames = []
    let hostnameToEdit = {}
    let showEditHost = false

    function getHostnames() {
        xfetch('/api/hostnames')
            .then(res => res.json())
            .then(function (hs) {
                hostnames = hs.hostnames
            })
            .catch(function (error) {
                notifications.danger('Error retrieving hostnames')
            })
    }

    function deleteHostname(domain) {
        return function () {
            xfetch(`/api/hostnames/${domain}`, { method: 'DELETE' })
                .then(() => {
                    const hostnameIndex = hostnames.findIndex(
                        hostname => hostname.domain === domain,
                    )
                    hostnames.splice(hostnameIndex, 1)
                    hostnames = hostnames
                })
                .catch(function (error) {
                    notifications.danger(
                        `Error deleting hostname ${domain}: ${error.message}`,
                    )
                })
        }
    }

    function toggleHostnameEdit(hostname) {
        return function () {
            hostnameToEdit = hostname
            showEditHost = !showEditHost
        }
    }

    function handleHostnameSave() {
        getHostnames()
        toggleHostnameEdit({})()
    }

    getHostnames()
</script>

<PageLayout>
    <div class="w-full mb-4 lg:mb-0 p-4 md:p-6 bg-white shadow-lg rounded">
        <div class="flex w-full">
            <div class="w-4/5">
                <h2 class="text-2xl md:text-3xl font-bold text-center mb-4">
                    Hostnames
                </h2>
            </div>
            <div class="w-1/5">
                <div class="text-right">
                    <HollowButton
                        color="green"
                        onClick="{toggleHostnameEdit({})}"
                    >
                        Add Hostname
                    </HollowButton>
                </div>
            </div>
        </div>

        <table class="table-fixed w-full mb-4">
            <thead>
                <tr>
                    <th class="w-2/6 px-4 py-2">Domain</th>
                    <th class="w-2/6 px-4 py-2">Username</th>
                    <th class="w-1/6 px-4 py-2">Password</th>
                    <th class="w-1/6 px-4 py-2">Actions</th>
                </tr>
            </thead>
            <tbody>
                {#each hostnames as hostname}
                    <tr>
                        <td class="border px-4 py-2">{hostname.domain}</td>
                        <td class="border px-4 py-2">{hostname.username}</td>
                        <td class="border px-4 py-2">{hostname.password}</td>
                        <td class="border px-4 py-2 text-right">
                            <HollowButton
                                onClick="{toggleHostnameEdit(hostname)}"
                                color="blue"
                            >
                                Edit
                            </HollowButton>
                            <HollowButton
                                onClick="{deleteHostname(hostname.domain)}"
                                color="red"
                            >
                                Delete
                            </HollowButton>
                        </td>
                    </tr>
                {/each}
            </tbody>
        </table>
    </div>

    {#if showEditHost}
        <HostnameForm
            hostnameToEdit="{hostnameToEdit}"
            toggleModal="{toggleHostnameEdit}"
            handleSave="{handleHostnameSave}"
            xfetch="{xfetch}"
            notifications="{notifications}"
        />
    {/if}
</PageLayout>

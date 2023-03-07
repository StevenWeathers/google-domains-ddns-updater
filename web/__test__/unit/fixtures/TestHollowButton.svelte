<script>
    import HollowButton from '../../../src/components/HollowButton.svelte'

    export let onClick = () => {}
</script>

<HollowButton onClick="{onClick}">Test Button</HollowButton>

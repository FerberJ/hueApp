<script>
    import NavbarButton from "./NavbarButton.svelte";

    export let items;
    export let selectedIndex = -1
    export let size = "lg"
</script>


{#each items as item}
<NavbarButton title={item.title} icon={item.icon} index={item.index} bind:selectedIndex={selectedIndex} {size} /> 
{/each}


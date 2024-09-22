<script>
    import Button from "$lib/components/ui/button/button.svelte";

    export let size = "lg"; // lg, md
    export let title;
    export let icon;
    export let index;
    export let selectedIndex = -1;

    const setSelectedIndex = () => {
        selectedIndex = index;
    };
</script>

<div class="gap-4">
    <Button
        href="#"
        size="sm"
        variant={selectedIndex === index ? "default" : "ghost"}
        class="w-44 flex items-center justify-start my-1 font-mono"
        on:click={setSelectedIndex}
    >
        <svelte:component this={icon} class="mr-2 size-4" aria-hidden="true" />
        {#if size === "lg"}
            {title}
        {/if}
    </Button>
</div>

<style>
</style>

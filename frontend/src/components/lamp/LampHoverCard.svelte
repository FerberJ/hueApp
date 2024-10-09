<script lang="ts">
    import { Slider } from "$lib/components/ui/slider";
    import { createEventDispatcher } from 'svelte';


    export let brightness;

    
    const dispatch = createEventDispatcher();

    $: if (brightness !== undefined) {
        dispatch('brightnessChange', { brightness });
    }
     
</script>
    <Slider bind:value={brightness} max={100} step={5} class="pb-8"/>

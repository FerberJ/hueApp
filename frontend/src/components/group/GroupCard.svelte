<script>
    import * as Card from "$lib/components/ui/card";
    import { Lightbulb, Heart } from "lucide-svelte";
    import { Button } from "$lib/components/ui/button/index.js";
    import { ToggleLight } from "../../../wailsjs/go/hue/Hue.js";
    import { writable } from "svelte/store";
    import { ToggleLightLike } from "../../../wailsjs/go/dao/App.js";

    export let group;
    export const groupStore = writable(group);



    function toggleGroup(group) {
        ToggleLight(group);
        groupStore.update((l) => ({ ...l, on: !l.on }));
    }

    function toggleLike() {
        ToggleLightLike(group);
        groupStore.update((l) => ({ ...l, liked: !l.liked }));
    }
</script>

<Card.Root class="font-mono relative">
    <Card.Content class="grid grid-cols-3 grid-flow-row gap-4 py-4 pr-8">
        <Button
            class="row-span-2"
            variant="outline"
            on:click={() => {
                toggleGroup($groupStore);
            }}
        >
            {#if $groupStore.on}
                <Lightbulb fill="#eab308" />
            {:else}
                <Lightbulb />
            {/if}
        </Button>
        <p class="col-span-2 text-left">{$groupStore.name}</p>
    </Card.Content>
    <Button
        variant="ghost"
        class="absolute top-0 right-0"
        on:click={toggleLike}
    >
        {#if $groupStore.liked}
            <Heart fill="#f472b6" />
        {:else}
            <Heart />
        {/if}
    </Button>
</Card.Root>

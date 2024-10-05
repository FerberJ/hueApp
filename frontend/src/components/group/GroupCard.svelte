<script>
    import * as Card from "$lib/components/ui/card";
    import { Lightbulb, Heart } from "lucide-svelte";
    import { Button } from "$lib/components/ui/button/index.js";
    import { ToggleGroup } from "../../../wailsjs/go/hue/Hue.js";
    import { writable } from "svelte/store";
    import { ToggleGroupLike } from "../../../wailsjs/go/dao/App.js";

    export let group;
    export const groupStore = writable(group);

    function toggleGroup(group) {
        ToggleGroup(group);
        groupStore.update((l) => ({ ...l, on: !l.on }));
    }

    function toggleLike() {
        ToggleGroupLike(group);
        groupStore.update((l) => ({ ...l, liked: !l.liked }));
    }

    $: if (group) {
        groupStore.set(group);
    }
</script>

<Card.Root class="font-mono relative min-w-80">
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

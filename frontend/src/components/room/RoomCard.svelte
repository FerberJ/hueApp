<script>
    import * as Card from "$lib/components/ui/card";
    import { Lightbulb, Heart } from "lucide-svelte";
    import { Button } from "$lib/components/ui/button/index.js";
    import { ToggleRoom } from "../../../wailsjs/go/hue/Hue.js";
    import { writable } from "svelte/store";
    import { ToggleRoomLike } from "../../../wailsjs/go/dao/App.js";

    export let room;
    export const roomStore = writable(room);



    function toggleRoom(room) {
        ToggleRoom(room);
        roomStore.update((l) => ({ ...l, on: !l.on }));
    }

    function toggleLike() {
        ToggleRoomLike(room);
        roomStore.update((l) => ({ ...l, liked: !l.liked }));
    }

    $: if (room) {
        roomStore.set(room)
    }
</script>

<Card.Root class="font-mono relative min-w-80">
    <Card.Content class="grid grid-cols-3 grid-flow-row gap-4 py-4 pr-8">
        <Button
            class="row-span-2"
            variant="outline"
            on:click={() => {
                toggleRoom($roomStore);
            }}
        >
            {#if $roomStore.on}
                <Lightbulb fill="#eab308" />
            {:else}
                <Lightbulb />
            {/if}
        </Button>
        <p class="col-span-2 text-left">{$roomStore.name}</p>
    </Card.Content>
    <Button
        variant="ghost"
        class="absolute top-0 right-0"
        on:click={toggleLike}
    >
        {#if $roomStore.liked}
            <Heart fill="#f472b6" />
        {:else}
            <Heart />
        {/if}
    </Button>
</Card.Root>

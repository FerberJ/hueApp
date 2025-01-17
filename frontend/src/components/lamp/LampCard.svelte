<script>
    // @ts-ignore
    import * as ContextMenu from "$lib/components/ui/context-menu";
    import * as Card from "$lib/components/ui/card";
    import * as HoverCard from "$lib/components/ui/hover-card";
    import { Lightbulb, Heart } from "lucide-svelte";
    import { Button } from "$lib/components/ui/button/index.js";
    import { ToggleLight, DimLight } from "../../../wailsjs/go/hue/Hue.js";
    import { writable } from "svelte/store";
    import LampDialog from "./LampDialog.svelte";
    import { ToggleLightLike } from "../../../wailsjs/go/dao/App.js";
    import LampHoverCard from "./LampHoverCard.svelte";

    export let lamp;
    export const lampStore = writable(lamp);

    const isDialogOpen = writable(false);

    const isHovered = writable(false);

    const openDialog = () => {
        isDialogOpen.set(true);
    };

    function toggleLight(lamp) {
        ToggleLight(lamp, !lamp.on);
        lampStore.update((l) => ({ ...l, on: !l.on }));
    }

    function toggleLike() {
        ToggleLightLike(lamp);
        lampStore.update((l) => ({ ...l, liked: !l.liked }));
    }

    function dimLight(brightness) {
        console.log($lampStore, lamp)
        if (brightness === 0) {
            ToggleLight(lamp, false);
            lampStore.update((l) => ({ ...l, on: false }));
        } else {
            if (!lamp.on) {
                console.log('dimming')
                DimLight(lamp, brightness);
            }
        }
    }

    $: if (lamp) {
        $lampStore.brightness = lamp.brightness;
        $lampStore.on = lamp.on;
    }
</script>

<Card.Root class="font-mono relative min-w-80">
    <ContextMenu.Root>
        <ContextMenu.Trigger>
            <Card.Content
                class="grid grid-cols-3 grid-flow-row gap-4 py-4 pr-8"
            >
                <HoverCard.Root>
                    <HoverCard.Trigger>
                        <Button
                            class="row-span-2"
                            variant="outline"
                            on:click={() => {
                                toggleLight($lampStore);
                                isHovered.set(true);
                            }}
                        >
                            {#if $lampStore.on}
                                <Lightbulb fill="#eab308" />
                            {:else}
                                <Lightbulb />
                            {/if}
                        </Button>
                    </HoverCard.Trigger>
                    {#if $lampStore.brightness > -1}
                    <HoverCard.Content class="w-80">
                            <LampHoverCard
                                brightness={[$lampStore.brightness]}
                                on:brightnessChange={(e) =>
                                    dimLight(e.detail.brightness[0])}
                            />
                        </HoverCard.Content>
                        {/if}
                </HoverCard.Root>
                {#if $lampStore.customName === ""}
                    <p class="col-span-2 text-left">{$lampStore.name}</p>
                {:else}
                    <p class="col-span-2 text-left">{$lampStore.customName}</p>
                {/if}
                <p class="col-start-2 text-left text-xs">
                    {$lampStore.group.name}
                </p>
                <p class="text-left text-xs">
                    {#if $lampStore.room.name}
                        {$lampStore.room.name}
                    {:else}
                        <span style="visibility: hidden;">Placeholder</span>
                    {/if}
                </p>
            </Card.Content>
            <Button
                variant="ghost"
                class="absolute top-0 right-0"
                on:click={toggleLike}
            >
                {#if $lampStore.liked}
                    <Heart fill="#f472b6" />
                {:else}
                    <Heart />
                {/if}
            </Button>
        </ContextMenu.Trigger>
        <ContextMenu.Content class="w-32">
            <ContextMenu.Item inset on:click={openDialog} class="font-mono">
                Edit
            </ContextMenu.Item>
        </ContextMenu.Content>
    </ContextMenu.Root>
</Card.Root>

<LampDialog {isDialogOpen} lamp={lampStore} />

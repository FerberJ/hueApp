<script>
    import { onMount } from "svelte";
    import { GetLights } from "../../../wailsjs/go/hue/Hue";
    import {
        GetGroupOptions,
        GetGroups,
        GetOrCreateLight,
        GetRoomOptions,
    } from "../../../wailsjs/go/dao/App.js";
    import { groupList, roomList } from "../../store";
    import ScrollArea from "$lib/components/ui/scroll-area/scroll-area.svelte";
    import GroupCard from "./GroupCard.svelte";

    let lamps = [];

    onMount(async () => {
        // Get all lights
        const result = await GetGroups();
        lamps =result
    });
</script>

<ScrollArea
    class="h-dvh"
    style="height: calc(100vh-50px)"
    orientation="both"
>
    <div class="grid grid-flow-row grid-cols-3 gap-4 px-4 py-4">
        {#each lamps as group}
            <GroupCard {group} />
        {/each}
    </div>
</ScrollArea>

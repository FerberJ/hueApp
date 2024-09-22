<script>
    import { onMount } from "svelte";
    import { GetLights } from "../../../wailsjs/go/hue/Hue";
    import {
        GetGroupOptions,
        GetOrCreateLight,
        GetRoomOptions,
    } from "../../../wailsjs/go/dao/App.js";
    import LampCard from "./LampCard.svelte";
    import { groupList, roomList } from "../../store";
    import ScrollArea from "$lib/components/ui/scroll-area/scroll-area.svelte";

    let lamps = [];

    onMount(async () => {
        // get groups
        const groups = await GetGroupOptions();
        groupList.set(groups);

        // get rooms
        const rooms = await GetRoomOptions();
        roomList.set(rooms);

        // Get all lights
        const result = await GetLights();
        const lampPromises = result.map(async (lamp) => {
            const dbLamp = await GetOrCreateLight(lamp);
            console.log("result", dbLamp.name, dbLamp);
            return dbLamp;
        });

        let dbLamps = [];
        dbLamps.push(...(await Promise.all(lampPromises)));
        dbLamps.sort((a, b) => {
            let aName = "";
            let bName = "";

            aName = a.customName === "" ? a.name : a.customName;
            bName = b.customName === "" ? b.name : b.customName;

            return aName.localeCompare(bName);
        });
        lamps = dbLamps;
    });
</script>

<ScrollArea
    class="h-dvh"
    style="height: calc(100vh-50px)"
    orientation="both"
>
    <div class="grid grid-flow-row grid-cols-3 gap-4 px-4 py-4">
        {#each lamps as lamp}
            <LampCard {lamp} />
        {/each}
    </div>
</ScrollArea>

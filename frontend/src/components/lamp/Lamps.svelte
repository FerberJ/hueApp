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

    function updateLamps() {
        GetLights().then(hueLights => {
            lamps = lamps.map(lamp => {
                const hueLamp = hueLights.find(hue => hue.id === lamp.id);
                if (hueLamp) {
                    return {
                        ...lamp,
                        on: hueLamp.on,
                    };
                }
                return lamp;
            });
        });
    }

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

        setInterval(() => {
            updateLamps();
        }, 1500);
    });
</script>

<ScrollArea class="h-dvh" style="height: calc(100vh-50px)" orientation="both">
    <div
        class="grid grid-flow-row sm:grid-cols-1 md:grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 gap-4 px-4 py-4"
    >
        {#each lamps as lamp}
            <LampCard {lamp} />
        {/each}
    </div>
</ScrollArea>

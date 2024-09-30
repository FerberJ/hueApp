<script>
    import { onMount } from "svelte";
    import { GetLights } from "../../../wailsjs/go/hue/Hue";
    import { GetRooms } from "../../../wailsjs/go/dao/App.js";
    import ScrollArea from "$lib/components/ui/scroll-area/scroll-area.svelte";
    import RoomCard from "./RoomCard.svelte";

    let rooms = [];

    onMount(async () => {
        const [roomsRes, hueLights] = await Promise.all([
            GetRooms(),
            GetLights(),
        ]);

        console.log(roomsRes)

        roomsRes.forEach((group) => {
            if (group.lights) {
                const filteredLights = hueLights.filter((item1) =>
                    group.lights.some((item2) => item1.id === item2.id),
                );

                group.on = filteredLights.some((light) => light.on);
            }
        });

        rooms = roomsRes;
    });
</script>

<ScrollArea class="h-dvh" style="height: calc(100vh-50px)" orientation="both">
    <div
        class="grid grid-flow-row sm:grid-cols-1 md:grid-cols-1 lg:grid-cols-1 xl:grid-cols-2 2xl:grid-cols-3 gap-4 px-4 py-4"
    >
        {#each rooms as room}
            <RoomCard {room} />
        {/each}
    </div>
</ScrollArea>

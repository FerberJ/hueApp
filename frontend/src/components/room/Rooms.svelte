<script>
    import { onMount } from "svelte";
    import { GetLights } from "../../../wailsjs/go/hue/Hue";
    import { GetRooms } from "../../../wailsjs/go/dao/App.js";
    import ScrollArea from "$lib/components/ui/scroll-area/scroll-area.svelte";
    import RoomCard from "./RoomCard.svelte";

    let rooms = [];

    function updateRooms() {
        GetLights().then((hueLight) => {
            rooms = rooms.map((room) => {
                if (room.lights) {
                    const filteredLights = hueLight.filter((item1) =>
                        room.lights.some((item2) => item1.id === item2.id),
                    );
                    return {
                        ...room,
                        on: filteredLights.some((light) => light.on),
                    };
                }
                return room;
            });
        });
    }

    onMount(async () => {
        const [roomsRes, hueLights] = await Promise.all([
            GetRooms(),
            GetLights(),
        ]);

        roomsRes.forEach((room) => {
            if (room.lights) {
                const filteredLights = hueLights.filter((item1) =>
                    room.lights.some((item2) => item1.id === item2.id),
                );

                room.on = filteredLights.some((light) => light.on);
            }
        });

        rooms = roomsRes;

        setInterval(() => {
            updateRooms();
        }, 1500);
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

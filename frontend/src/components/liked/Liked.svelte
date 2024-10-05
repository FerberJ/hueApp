<script>
    import { onMount } from "svelte";
    import {
        GetLikedRooms,
        GetLikedGroups,
        GetLikedLights,
    } from "../../../wailsjs/go/dao/App.js";
    import { GetLights } from "../../../wailsjs/go/hue/Hue";
    import RoomCard from "../room/RoomCard.svelte";
    import GroupCard from "../group/GroupCard.svelte";
    import LampCard from "../lamp/LampCard.svelte";
    import Separator from "$lib/components/ui/separator/separator.svelte";
    import ScrollArea from "$lib/components/ui/scroll-area/scroll-area.svelte";

    let rooms = [];
    let groups = [];
    let lamps = [];

    function updateLamps() {
        GetLights().then((hueLights) => {
            lamps = lamps.map((lamp) => {
                const hueLamp = hueLights.find((hue) => hue.id === lamp.id);
                if (hueLamp) {
                    return {
                        ...lamp,
                        on: hueLamp.on,
                    };
                }
                return lamp;
            });

            rooms = rooms.map((room) => {
                if (room.lights) {
                    const filteredLights = hueLights.filter((item1) =>
                        room.lights.some((item2) => item1.id === item2.id),
                    );
                    return {
                        ...room,
                        on: filteredLights.some((light) => light.on),
                    };
                }
                return room;
            });

            groups = groups.map((group) => {
                if (group.lights) {
                    const filteredLights = hueLights.filter((item1) =>
                        group.lights.some((item2) => item1.id === item2.id),
                    );
                    return {
                        ...group,
                        on: filteredLights.some((light) => light.on),
                    };
                }
                return group;
            });
        });
    }

    onMount(async () => {
        const hueLights = await GetLights();

        const roomsRes = await GetLikedRooms();
        const groupsRes = await GetLikedGroups();
        const lampsRes = await GetLikedLights();

        roomsRes.forEach((room) => {
            if (room.lights) {
                const filteredLights = hueLights.filter((item1) =>
                    room.lights.some((item2) => item1.id === item2.id),
                );
                room.on = filteredLights.some((light) => light.on);
            }
        });

        groupsRes.forEach((group) => {
            if (group.lights) {
                const filteredLights = hueLights.filter((item1) =>
                    group.lights.some((item2) => item1.id === item2.id),
                );
                group.on = filteredLights.some((light) => light.on);
            }
        });

        lampsRes.forEach((lamp) => {
            lamp.on = hueLights.find((hue) => hue.id === lamp.id).on;
        });

        rooms = roomsRes;
        groups = groupsRes;
        lamps = lampsRes;
    });

    setInterval(() => {
        updateLamps();
    }, 1500);
</script>

<ScrollArea class="h-dvh" style="height: calc(100vh-50px)" orientation="both">
    <h1 class="font-mono pt-4">Rooms</h1>
    <div
        class="grid grid-flow-row sm:grid-cols-1 md:grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 gap-4 px-4 py-4"
    >
        {#each rooms as room}
            <RoomCard {room} />
        {/each}
    </div>
    <Separator />
    <h1 class="font-mono pt-4">Group</h1>
    <div
        class="grid grid-flow-row sm:grid-cols-1 md:grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 gap-4 px-4 py-4"
    >
        {#each groups as group}
            <GroupCard {group} />
        {/each}
    </div>
    <Separator />
    <h1 class="font-mono pt-4">Light</h1>
    <div
        class="grid grid-flow-row sm:grid-cols-1 md:grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 gap-4 px-4 py-4"
    >
        {#each lamps as lamp}
            <LampCard {lamp} />
        {/each}
    </div>
</ScrollArea>

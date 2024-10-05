<script>
    import { onMount } from "svelte";
    import { GetLights } from "../../../wailsjs/go/hue/Hue";
    import { GetGroups } from "../../../wailsjs/go/dao/App.js";
    import ScrollArea from "$lib/components/ui/scroll-area/scroll-area.svelte";
    import GroupCard from "./GroupCard.svelte";

    let groups = [];

    function updateGroups() {
        GetLights().then((hueLight) => {
            groups = groups.map((group) => {
                if (group.lights) {
                    const filteredLights = hueLight.filter((item1) =>
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
        const [groupsRes, hueLights] = await Promise.all([
            GetGroups(),
            GetLights(),
        ]);

        groupsRes.forEach((group) => {
            if (group.lights) {
                const filteredLights = hueLights.filter((item1) =>
                    group.lights.some((item2) => item1.id === item2.id),
                );

                group.on = filteredLights.some((light) => light.on);
            }
        });

        groups = groupsRes;

        setInterval(() => {
            updateGroups();
        }, 1500);
    });
</script>

<ScrollArea class="h-dvh" style="height: calc(100vh-50px)" orientation="both">
    <div
        class="grid grid-flow-row sm:grid-cols-1 md:grid-cols-1 lg:grid-cols-1 xl:grid-cols-2 2xl:grid-cols-3 gap-4 px-4 py-4"
    >
        {#each groups as group}
            <GroupCard {group} />
        {/each}
    </div>
</ScrollArea>

<script>
    import * as Dialog from "$lib/components/ui/dialog";
    import { Input } from "$lib/components/ui/input/index.js";
    import { Label } from "$lib/components/ui/label";
    import Select from "../select/Select.svelte";
    import { Button } from "$lib/components/ui/button/index.js";
    import { Save, Plus, Minus } from "lucide-svelte";
    import { writable } from "svelte/store";
    import { groupList, roomList } from "../../store";
    import {
        CreateGroupWithName,
        CreateRoomWithName,
        UpdateLight,
    } from "../../../wailsjs/go/dao/App";


    export let isDialogOpen;
    export let lamp

    let showAddGroup = writable(false);
    let showAddRoom = writable(false);

    let selectedGroup = $lamp.group.name
    let selectedRoom = $lamp.room.name

    let groupName = "";
    let roomName = "";
    let lampName = "";

    const toggleAddGroup = () => {
        showAddGroup.update((value) => !value);
    };

    const toggleAddRoom = () => {
        showAddRoom.update((value) => !value);
    };

    const addGroup = () => {
        if (groupName !== "") {
            CreateGroupWithName(groupName);
            groupList.update((list) => {
                return [...list, { value: groupName, label: groupName }];
            });
            groupName = "";
        }
    };

    const addRoom = () => {
        if (roomName !== "") {
            CreateRoomWithName(roomName);
            roomList.update((list) => {
                return [...list, { value: roomName, label: roomName }];
            });
            roomName = "";
        }
    };

    const saveLamp = () => {
        if (lampName !== "") {
            $lamp.customName = lampName;
        }
        $lamp.group.name = selectedGroup
        $lamp.room.name = selectedRoom

        UpdateLight($lamp, selectedRoom, selectedGroup);
        lampName = ""
        isDialogOpen.set(false);
    };
</script>

<Dialog.Root bind:open={$isDialogOpen}>
    <Dialog.Content class="custom-dialog-width">
        <Dialog.Header>
            <Dialog.Title class="font-mono">{$lamp.customName === "" ? $lamp.name : $lamp.customName}</Dialog.Title>
            <Dialog.Description class="font-mono">
                <div class="grid grid-cols-1 gap-4 pt-4">
                    <Label>Lamp</Label>
                    <Input
                        bind:value={lampName}
                        type="name"
                        id="name"
                        placeholder={$lamp.customName === "" ? $lamp.name : $lamp.customName}
                        class="w-64"
                    />
                    <Label>Group</Label>
                    <div class="flex gap-2">
                        <Select
                            items={$groupList}
                            title="Select Group"
                            selected={{
                                value: $lamp.group.name,
                                label: $lamp.group.name,
                            }}
                            bind:label={selectedGroup}
                        />
                        <Button variant="ghost" on:click={toggleAddGroup}>
                            {#if $showAddGroup}
                                <Minus />
                            {:else}
                                <Plus />
                            {/if}
                        </Button>
                    </div>
                    {#if $showAddGroup}
                        <div class="flex gap-2">
                            <Input bind:value={groupName} />
                            <Button variant="ghost" on:click={addGroup}>
                                <Save />
                            </Button>
                        </div>
                    {/if}
                    <Label>Room</Label>
                    <div class="flex gap-2">
                        <Select
                            items={$roomList}
                            title="Select Room"
                            selected={{
                                value: $lamp.room.name,
                                label: $lamp.room.name,
                            }}
                            bind:label={selectedRoom}
                        />
                        <Button variant="ghost" on:click={toggleAddRoom}>
                            {#if $showAddRoom}
                                <Minus />
                            {:else}
                                <Plus />
                            {/if}
                        </Button>
                    </div>
                    {#if $showAddRoom}
                        <div class="flex gap-2">
                            <Input bind:value={roomName} />
                            <Button variant="ghost" on:click={addRoom}>
                                <Save />
                            </Button>
                        </div>
                    {/if}

                    <Button on:click={saveLamp}>
                        <Save class="mr-2 h-4 w-4" />
                        Save Lamp
                    </Button>
                </div>
            </Dialog.Description>
        </Dialog.Header>
    </Dialog.Content>
</Dialog.Root>

<style>
    :global(.custom-dialog-width) {
        @apply w-auto;
    }
</style>

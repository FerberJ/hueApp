<script>
    import * as Select from "$lib/components/ui/select/index.js";

    export let width = "w-64";
    export let items = [];
    export let title = "";
    export let label = "";
    export let selected = { value: "", label: "" };

    function setLabel(name) {
        label = name;
    }
</script>

<Select.Root portal={null} {selected}>
    <Select.Trigger class="{width}">
        <Select.Value placeholder={title} />
    </Select.Trigger>
    <Select.Content>
        <Select.Group>
            {#if label !== ""}
                <Select.Label>{label}</Select.Label>
            {/if}
            {#each items as item}
                <Select.Item
                    value={item.value}
                    label={item.label}
                    on:click={() => setLabel(item.label)}
                >
                    {item.label}
                </Select.Item>
            {/each}
        </Select.Group>
    </Select.Content>
</Select.Root>

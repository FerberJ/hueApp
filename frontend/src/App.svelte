<script>
  import { Heart, LampFloor, Group, House } from "lucide-svelte";
  import "./app.css";
  import Lamps from "./components/lamp/Lamps.svelte";
  import Navbar from "./components/navbar/Navbar.svelte";
  import Groups from "./components/group/Groups.svelte";
  import Rooms from "./components/room/Rooms.svelte";
    import Liked from "./components/liked/Liked.svelte";

  const sidebarNavItems = [
    {
      title: "Favorite",
      icon: Heart,
      index: 0,
    },
    {
      title: "Rooms",
      icon: House,
      index: 1,
    },
    {
      title: "Groups",
      icon: Group,
      index: 2,
    },
    {
      title: "Lights",
      icon: LampFloor,
      index: 3,
    },
  ];

  let selectedIndex = 0;
</script>

<main>
  <div class="flex">
    <div class="gap-4 py-4 px-4">
      <div class="rounded-md border">
        <Navbar items={sidebarNavItems} bind:selectedIndex size="lg" />
      </div>
    </div>
    <div>
      {#if selectedIndex === 2}
        <Groups />
      {/if}
      {#if selectedIndex === 3}
        <Lamps />
      {/if}
      {#if selectedIndex === 1}
        <Rooms />
      {/if}
      {#if selectedIndex === 0}
      <Liked />
    {/if}
    </div>
  </div>
</main>

<style>
</style>

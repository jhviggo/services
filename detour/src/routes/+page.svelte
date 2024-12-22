<div class="page-content">
  <div class="map-area" bind:clientWidth={mapWidth}>
    {#if mapWidth > 0}
      <div class="map-container" style="width: {mapWidth}px;">
        <DraggableMap on:point-click={(e) => scrollToTimeline(e)} />
      </div>
    {/if}
  </div>
  <div class="divider"></div>
  <div>
    <div class="space">
      <Banner image="/images/banner.webp" imageAlt="Banner showing my motorcycle">
        <div class="over-banner"></div>
      </Banner>
    </div>
  
    {#each timeline as item}
      <div class="space">
        <SideBySide image="{item.image}" content="{item.content}" itemId={item.itemId} flipped={item.id % 2 == 1}/>
      </div>
    {/each}
  </div>
</div>

<button class="mobile-map" on:click={toggleMobileMap}>
  <div style:animation={pauseMobileMapAnimation ? '' : 'wiggle 2.5s infinite'}>
    See the route
    <icon class="material-symbols-outlined">
      map
    </icon>
  </div>
</button>

<Drawer open={mapDrawerOpen} side="left" on:close={() => (mapDrawerOpen = !mapDrawerOpen)}>
  <div slot="header" class="map-header">
    <h1 class="white">The route</h1>
  </div>
  <div class="mobile-map-container" bind:clientWidth={mobileMapWidth}>
    {#if mobileMapWidth > 0}
      <DraggableMap on:point-click={(e) => scrollToTimeline(e)} />
    {/if}
  </div>
</Drawer>
  
<script lang="ts">
  import SideBySide from '$components/SideBySide.svelte';
  import { onMount } from 'svelte';
  import { getTimelines, type Timeline } from '$lib/api'
  import DraggableMap from '$components/DraggableMap.svelte';
  import Banner from '$components/Banner.svelte';
  import Drawer from '$components/Drawer.svelte';

  let timeline: Timeline[] = [];
  let mapWidth: number = 0;
  let mobileMapWidth: number = 0;
  let mapDrawerOpen: boolean = false;
  let pauseMobileMapAnimation: boolean = false;

  onMount(async (): Promise<any> => {
    timeline = await getTimelines();
    timeline.sort((a, b) => b.id - a.id);
  });

  function scrollToTimeline(event: CustomEvent) {
    document.querySelector(`#${event.detail}`)?.scrollIntoView({ behavior: 'smooth', block: 'center', inline: 'center' });
  }

  function toggleMobileMap() {
    mapDrawerOpen = !mapDrawerOpen
    pauseMobileMapAnimation = true;
  }
</script>

<style>
  .page-content {
    display: grid;
    grid-template-columns: 1fr 4px 2fr;
  }

  .divider {
    background-color: var(--blue);
  }

  .over-banner {
    background: linear-gradient(0deg, rgba(0,0,0,1) 0%, rgba(0,0,0,0.3) 40%, rgba(0,0,0,0) 100%);
    height: 100%;
    width: 100%;
  }

  .space {
    margin-bottom: 2rem;
  }

  .map-container {
    position: fixed;
  }

  .mobile-map {
    display: none;
    position: fixed;
    font-weight: bold;
    width: 100%;
    height: 3rem;
    bottom: 0;
    background-color: var(--blue);
  }

  .mobile-map div {
    display: flex;
    justify-content: center;
    align-items: center;
    color: white;
  }

  .mobile-map div icon {
    color: white;
    margin-left: 0.5rem;
  }

  .mobile-map-container {
    width: 85dvw;
  }

  .map-header {
    height: 100%;
    display: flex;
    align-items: center;
  }

  .map-header h1 {
    margin: 0;
    margin-left: 1rem;
  }

  @media (max-width: 768px) {
    .page-content {
      display: block;
    }

    .map-area {
      display: none;
      width: 0;
    }

    .mobile-map {
      display: block;
    }
  }
</style>

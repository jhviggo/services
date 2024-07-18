<script lang="ts">
  import { type Refuel } from '@api/refuels';

  export let refuel: Refuel;

  let isOpen: boolean = false;

  function toggleFolder(): void {
    isOpen = !isOpen;
  }

  function formatDate(d: string | undefined): string {
    if (!d) return '';
    return d.split(' ')[0];
  }
</script>

<div class="container">
  <button on:click={toggleFolder}>
    <span class="text">
      <icon class="material-symbols-outlined">map</icon>
      {#if isOpen}
        ID: {refuel.id}
      {:else}
        {refuel.tripKM}Km trip ({refuel.cost}{refuel.currency})
      {/if}
    </span>

    {#if isOpen}
      <icon class="material-symbols-outlined">keyboard_arrow_up</icon>
    {:else}
      <icon class="material-symbols-outlined">keyboard_arrow_down</icon>
    {/if}
  </button>
  
 
  {#if isOpen}
    <div class="content">
      <span class="bold">Trip:</span><span>{refuel.tripKM} km</span>
      <span class="bold">Price:</span><span>{refuel.cost} {refuel.currency.toLocaleLowerCase()}</span>
      <span class="bold">Liters:</span><span>{refuel.liters} L</span>
      <span class="bold">Odometer:</span><span>{refuel.totalKM} Km</span>
      <span class="bold">Date:</span><span>{formatDate(refuel.createdAt)}</span>
    </div>
  {/if}

</div>

<style>
  .container {
    border: 1px solid var(--primary-blue);
    overflow-y: hidden;
  }

  button {
    background-color: unset;
    border: unset;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0;
    width: 100%;
    padding: 0.25rem;
  }

  .text {
    display: flex;
    align-items: center;
  }
  
  .content {
    padding: 0.25rem;
    display: grid;
    grid-template-columns: auto 1fr;
    gap: 0.5rem;
  }
</style>
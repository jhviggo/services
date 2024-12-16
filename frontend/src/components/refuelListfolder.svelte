<script lang="ts">
  import { type Refuel } from '@api/refuels';
  import { createEventDispatcher } from 'svelte';

  export let refuel: Refuel;

  const dispatch = createEventDispatcher();

  let isOpen: boolean = false;

  function toggleFolder(): void {
    isOpen = !isOpen;
  }

  function formatDate(d: string | undefined): string {
    if (!d) return '';
    return d.split('T')[0];
  }
</script>

<div class="container">
  <button class="refuel-button" on:click={toggleFolder}>
    <div class="map">
      <icon class="material-symbols-outlined">map</icon>
    </div>
    <span>
      {refuel.trip_km}Km trip ({refuel.cost}{refuel.currency})
    </span>

    {#if isOpen}
      <icon class="material-symbols-outlined arrow">keyboard_arrow_up</icon>
    {:else}
      <icon class="material-symbols-outlined arrow">keyboard_arrow_down</icon>
    {/if}
  </button>
  
 
  {#if isOpen}
    <div class="content">
      <span class="bold">Trip:</span><span>{refuel.trip_km} km</span>
      <span class="bold">Price:</span><span>{refuel.cost} {refuel.currency?.toLocaleLowerCase()}</span>
      <span class="bold">Liters:</span><span>{refuel.liters} L</span>
      <span class="bold">Odometer:</span><span>{refuel.total_km} Km</span>
      <span class="bold">Date:</span><span>{formatDate(refuel.created_at)}</span>
      <span class="bold">ID:</span><span>{refuel.id}</span>
      
      <button class="btn btn-danger" on:click={() => dispatch('delete', refuel)}><icon class="material-symbols-outlined">delete</icon> Delete</button>
    </div>
  {/if}
</div>

<style>
  .container {
    border: 1px solid gray;
    border-radius: 0.2rem;
    overflow-y: hidden;

    & .arrow {
      margin-right: 0.5rem;
    }

    & .delete {
      color: var(--red);
    }
  }

  .refuel-button {
    background-color: unset;
    border: unset;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0;
    width: 100%;
    height: 3rem;
    display: grid;
    grid-template-columns: auto 1fr auto auto;
    text-align: start;
    gap: 0.5rem;
  }

  .map {
    height: 3rem;
    width: 3rem;
    background-color: var(--primary-blue);
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .map icon {
    color: white;
  }
  
  .content {
    padding: 0.25rem;
    display: grid;
    grid-template-columns: auto 1fr;
    gap: 0.5rem;
    border-left: 4px solid var(--primary-blue);
  }
</style>
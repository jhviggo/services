<script lang="ts">
  import { type IVehicle } from '@api/vehicles';
 
  export let vehicles: IVehicle[];
  export let selectedVehicle: IVehicle;

  let vehicleListOpen = false;
  let vehiclesExceptSelected: IVehicle[];

  $: if (selectedVehicle) {
    updateVehiclesExceptSelected();
  }

  function toggleVehicleList() {
    vehicleListOpen = !vehicleListOpen;
  }

  function updateVehiclesExceptSelected() {
    vehiclesExceptSelected = vehicles.filter(value => value.id !== selectedVehicle.id);
    vehicleListOpen = false;
  }

  function selectVehicle(vehicleId: number | string) {
    selectedVehicle = vehicles.find(value => value.id === vehicleId) || selectedVehicle;
  }
</script>

{#if selectedVehicle}
  <div>
    <button class="model" on:click={toggleVehicleList}>
      <div class="selected">
        <icon class="material-symbols-outlined vehicle-icon blue">
          two_wheeler
        </icon>
      </div>
      <span>{selectedVehicle.model}</span>
      {#if vehicles.length > 1}
        <icon class="material-symbols-outlined">
          {#if vehicleListOpen}
            keyboard_arrow_up
          {:else}
            keyboard_arrow_down
          {/if}
        </icon>
      {/if}
    </button>
  </div>
{/if}

{#if vehicleListOpen}
  {#each vehiclesExceptSelected as vehicle}
    <button class="model" on:click={() => selectVehicle(vehicle.id)}>
      <icon class="material-symbols-outlined vehicle-icon">
        two_wheeler
      </icon>
      <span>{vehicle.model}</span>
    </button>
  {/each}
{/if}

<style>
  .model {
    display: grid;
    grid-template-columns: auto 1fr auto;
    align-items: center;
    background-color: unset;
    border: unset;
    padding: 0;
    width: 100%;
    height: 3rem;
    text-align: start;
    gap: 0.5rem;
    border: 1px solid gray;
    border-radius: 0.25rem;
    overflow: hidden;
  }

  .vehicle-icon {
    border-radius: 0.2rem;
    padding: 0.2rem;
  }

  .selected {
    background-color: var(--primary-blue);
    height: 3rem;
    width: 3rem;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .blue {
    background-color: var(--primary-blue);
    color: white;
  }
</style>

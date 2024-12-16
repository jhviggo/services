<script lang="ts">
  import Section from '@components/section.svelte';
  import ButtonIcon from '@components/buttonIcon.svelte';
  import { onMount } from 'svelte';
  import { addRefuel, fetchRefuels, deleteRefuel, type Refuel } from '@api/refuels';
  import { listVehicles, type IVehicle } from '@api/vehicles';
  import RefuelList from '@components/refuelList.svelte';
  import Spinner from '@components/spinner.svelte';
  import Vehicle from '@components/vehicle.svelte';
  import ConfirmationPrompt from '@components/confirmationPrompt.svelte';

  let loading: boolean = false;
  let refuels: Refuel[];
  let addRefuelIsOpen: boolean = false;
  let vehicleList: IVehicle[] = [];
  let selectedVehicle: IVehicle;
  let deletingRefuel: Refuel;
  let showDeletingRefuelPrompt: boolean = false;

  $: if(selectedVehicle) {
    updateRefuels();
  }

  onMount(async () => {
    loading = true;
    vehicleList = await listVehicles();
    console.log('vehicles', vehicleList);
    if (vehicleList.length > 0) {
      selectedVehicle = vehicleList[0];
      refuels = await fetchRefuels(vehicleList[0].id);
    }
    console.log('refuels', refuels);
    loading = false;
  });

  async function updateRefuels() {
    loading = true;
    refuels = await fetchRefuels(selectedVehicle.id);
    loading = false;
  }

  let values: { [key: string]: string | undefined } = {
    vehicle: undefined,
    total_km: undefined,
    trip_km: undefined,
    liters: undefined,
    cost: undefined,
  };

  async function handleAddRefuel() {
    const ok = await addRefuel({
      vehicle_id: selectedVehicle.id,
      total_km: Number(values.total_km),
      trip_km: Number(values.trip_km),
      liters: Number(values.liters),
      cost: Number(values.cost),
      currency: values.currency,
    } as Refuel);

    if (ok) {
      await updateRefuels();
      values = {};
      addRefuelIsOpen = false;
    }
    
  }

  function openDeleteRefuelPrompt(e: CustomEvent<Refuel>) {
    deletingRefuel = e.detail;
    showDeletingRefuelPrompt = true;
  }

  function closeDeleteRefuelPrompt() {
    deletingRefuel = {} as Refuel;
    showDeletingRefuelPrompt = false;
  }

  async function onDeleteRefuel() {
    await deleteRefuel(deletingRefuel);
    await updateRefuels();
    showDeletingRefuelPrompt = false;
  }
</script>

<Section>
  <h1>Selected bike</h1>
  <Vehicle bind:selectedVehicle={selectedVehicle} vehicles={vehicleList} />
</Section>

{#if addRefuelIsOpen}
  <Section>
    <small>Note: only numbers allowed. e.g. 155.25.</small>
    <form on:submit|preventDefault={handleAddRefuel} action="POST">
      <label>
        Vehicle
        <select bind:value={values.vehicle} class={loading ? 'skeleton' : ''}>
          {#if loading}
            <option value="Loading">Loading...</option>
          {:else}
            {#each vehicleList as vehicle}
              <option value={vehicle.id}>{ vehicle.model }</option>
            {/each}
          {/if}
        </select>
      </label>

      <label>
        Total vehicle KM
        <input bind:value={values.total_km} placeholder="Vehicle's kilometers" type="text" required pattern="([0-9]*[.])?[0-9]+">
        <span class="error">Must be a number</span>
      </label>

      <label>
        Trip KM
        <input bind:value={values.trip_km} placeholder="Trip distance" type="text" required pattern="([0-9]*[.])?[0-9]+">
        <span class="error">Must be a number</span>
      </label>

      <label>
        Liters
        <input bind:value={values.liters} placeholder="Liters refueled" type="text" required pattern="([0-9]*[.])?[0-9]+">
        <span class="error">Must be a number</span>
      </label>

      <label>
        Cost
        <input bind:value={values.cost} placeholder="Cost of refueling" type="text" required pattern="([0-9]*[.])?[0-9]+">
        <span class="error">Must be a number</span>
      </label>

      <label>
        Currency
        <select bind:value={values.currency}>
          <option value="dkk">DKK</option>
          <option value="eur">EUR</option>
          <option value="usd">USD</option>
        </select>
      </label>

      <div class="submit-container">
        <button class="btn btn-primary" type="submit">
          <span class="material-symbols-outlined">local_gas_station</span>
          Add refuel
        </button>
        <button class="btn btn-danger" on:click={() => addRefuelIsOpen = false}>Cancel</button>
      </div>
    </form>
  </Section>
{:else}
  <Section>
    <div class="refuels-header">
      <h1>Refuels</h1>
      <button class="btn btn-primary" on:click={() => addRefuelIsOpen = true}>
        <span class="material-symbols-outlined">add</span>
        Add refuel
      </button>
    </div>
    {#if loading}
      <Spinner />
    {:else if !loading && !refuels?.length}
      No refuels found
    {:else}
      <RefuelList {refuels} on:delete={(e) => openDeleteRefuelPrompt(e)} />
    {/if}
  </Section>
{/if}


{#if showDeletingRefuelPrompt}
  <ConfirmationPrompt message="Are you sure you want to delete this refuel?" on:yes={onDeleteRefuel} on:no={closeDeleteRefuelPrompt} />
{/if}

<style>
  form {
    display: flex;
    flex-direction: column;
    max-width: 25rem;
  }

  label {
    display: flex;
    flex-direction: column;
    width: 100%;
    margin-bottom: 1.25rem;

    & span {
      display: none;
    }
  }

  .refuels-header {
    display: flex;
    justify-content: space-between;
  }

  .submit-container {
    display: flex;
    gap: 0.5rem;
  }

  input, select {
    padding: 0.5rem 0.75rem;
  }

  input:user-invalid {
    border: 1px solid var(--red);
  }

  input:user-invalid + span {
    display: block;
  }

  .error {
    font-size: 0.85rem;
    height: 1.25rem;
  }

  .skeleton {
    background: linear-gradient(
        to right,
        rgba(255, 255, 255, 0),
        rgba(255, 255, 255, 0.5) 50%,
        rgba(255, 255, 255, 0) 80%
      ),
      #ddd;
    background-repeat: repeat-y;
    background-size: 50px 500px;
    animation: shine 1.5s infinite;
  }

  @keyframes shine {
    100% {
      background-position: 100%;
    }
  }
</style>
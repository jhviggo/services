<script lang="ts">
  import Section from '@components/section.svelte';
  import ButtonIcon from '@components/buttonIcon.svelte';
  import { onMount } from 'svelte';
  import { addRefuel, fetchRefuels, type Refuel } from '@api/refuels';
  import { listVehicles, type IVehicle } from '@api/vehicles';
  import RefuelList from '@components/refuelList.svelte';
  import Spinner from '@components/spinner.svelte';
  import Vehicle from '@components/vehicle.svelte';

  let loading: boolean = false;
  let refuels: Refuel[];
  let addRefuelIsOpen: boolean = false;
  let vehicleList: IVehicle[] = [];
  let selectedVehicle: IVehicle;

  $: if(selectedVehicle) {
    console.log(selectedVehicle);
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
    totalKM: undefined,
    tripKM: undefined,
    liters: undefined,
    cost: undefined,
  };

  let shownErrors: { [key: string]: boolean } = {
    vehicles: false,
    totalKM: false,
    tripKM: false,
    liters: false,
    cost: false,
  };


  onMount(async () => {
    shownErrors.vehicle = false;
  });

  function validateNumber(n: string | undefined, value: string) {
    if (n?.includes('.')) {
      shownErrors[value] = n !== Number(n).toFixed(2);
    } else {
      shownErrors[value] = n !== Number(n).toString();
    }
  }

  async function handleAddRefuel() {
    validateNumber(values.totalKM, 'totalKM');
    validateNumber(values.tripKM, 'tripKM');
    validateNumber(values.liters, 'liters');
    validateNumber(values.cost, 'cost');

    const allValid = Object.values(shownErrors).reduce((a, b) => !a && !b);

    if (allValid) {
      const ok = await addRefuel({
        vehicleId: selectedVehicle.id,
        totalKM: Number(values.totalKM),
        tripKM: Number(values.tripKM),
        liters: Number(values.liters),
        cost: Number(values.cost),
        currency: values.currency,
      } as Refuel);

      if (ok) {
        updateRefuels();
        values = {};
      }
    }
  }
</script>

<Section>
  <h1>Selected bike</h1>
  <Vehicle bind:selectedVehicle={selectedVehicle} vehicles={vehicleList} />
</Section>

<Section>
  <h1>Refuels</h1>
  {#if loading}
    <Spinner />
  {:else if !loading && !refuels?.length}
    No refuels found
  {:else}
    <RefuelList {refuels} />
  {/if}
</Section>
    
<Section>
  <button class="opener" on:click={() => addRefuelIsOpen = !addRefuelIsOpen}>
    <h1>Add a new refuel</h1>
    {#if addRefuelIsOpen}
      <icon class="material-symbols-outlined">keyboard_arrow_up</icon>
    {:else}
      <icon class="material-symbols-outlined">keyboard_arrow_down</icon>
    {/if}
  </button>
  {#if addRefuelIsOpen}
    <hr>
    <small>Note: only two digits after point is allowed. e.g. 155.25.</small>
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
        <input on:change={() => validateNumber(values.totalKM, 'totalKM')} bind:value={values.totalKM} placeholder="Vehicle's kilometers" class={shownErrors.totalKM ? 'invalid' : ''} type="text">
        <span hidden={!shownErrors.totalKM} class="error">Invalid total distance</span>
      </label>

      <label>
        Trip KM
        <input on:change={() => validateNumber(values.tripKM, 'tripKM')} bind:value={values.tripKM} placeholder="Trip distance" class={shownErrors.tripKM ? 'invalid' : ''} type="text">
        <span hidden={!shownErrors.tripKM} class="error">Invalid trip distance</span>
      </label>

      <label>
        Liters
        <input on:change={() => validateNumber(values.liters, 'liters')} bind:value={values.liters} placeholder="Liters refueled" class={shownErrors.liters ? 'invalid' : ''} type="text">
        <span hidden={!shownErrors.liters} class="error">Invalid liters fueled</span>
      </label>

      <label>
        Cost
        <input on:change={() => validateNumber(values.cost, 'cost')} bind:value={values.cost} placeholder="Cost of refueling" class={shownErrors.cost ? 'invalid' : ''} type="text">
        <span hidden={!shownErrors.cost} class="error">Invalid cost</span>
      </label>

      <label>
        Currency
        <select bind:value={values.currency}>
          <option value="dkk">DKK</option>
          <option value="eur">EUR</option>
          <option value="usd">USD</option>
        </select>
      </label>

      <ButtonIcon icon="local_gas_station">Add Refuel</ButtonIcon>
    </form>
  {/if}
</Section>

<style>
  .opener {
    background-color: unset;
    border: unset;
    text-align: left;
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }

  form {
    display: flex;
    flex-direction: column;
    max-width: 25rem;
  }

  label {
    display: flex;
    flex-direction: column;
    width: 100%;
  }

  input, select {
    padding: 0.5rem 0.75rem;
    margin-bottom: 1.25rem;
  }

  .error {
    font-size: 0.85rem;
    height: 1.25rem;
  }

  .invalid {
    border-color: var(--error);
    margin-bottom: 0;
  }

  .invalid:focus-visible {
    outline-color: var(--error);
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
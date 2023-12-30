<script lang="ts">
  import Section from '@components/section.svelte';
  import ButtonIcon from '@components/buttonIcon.svelte';
  import { onMount } from 'svelte';
  import { fetchVehicles } from '@api/vehicles';

  let loading: boolean = false;

  let values: { [key: string]: string | undefined } = {
    vehicle: undefined,
    totalkm: undefined,
    tripkm: undefined,
    liters: undefined,
    cost: undefined,
  };

  let shownErrors: { [key: string]: boolean } = {
    vehicles: true,
    totalkm: false,
    tripkm: false,
    liters: false,
    cost: false,
  };

  let vehicles: string[] = ['suzuki', 'kawasaki'];

  onMount(async () => {
    await fetchVehicles();
    shownErrors.vehicle = false;
  });

  function validateNumber(n: string | undefined, value: string) {
    shownErrors[value] = n !== Number(n).toFixed(2);
  }

  function addRefuel() {
    validateNumber(values.totalkm, 'totalkm');
    validateNumber(values.tripkm, 'tripkm');
    validateNumber(values.liters, 'liters');
    validateNumber(values.cost, 'cost');

    const allValid = Object.values(shownErrors).reduce((a, b) => !a && !b);

    if (allValid) {
      // POST
    }
    console.log('values', values);
  }
</script>

<Section>
  <h1>Add a new refuel</h1>
  <small>Note: only two digits after point is allowed. e.g. 155.25.</small>
  <hr>
  <form on:submit|preventDefault={addRefuel} action="POST">
    <label>
      Vehicle
      <select bind:value={values.vehicle} class={loading ? 'skeleton' : ''}>
        {#if loading}
          <option value="Loading">Loading...</option>
        {:else}
          {#each vehicles as vehicle}
            <option value={vehicle}>{ vehicle }</option>
          {/each}
        {/if}
      </select>
    </label>

    <label>
      Total vehicle KM
      <input on:change={() => validateNumber(values.totalkm, 'totalkm')} bind:value={values.totalkm} placeholder="Vehicle's kilometers" class={shownErrors.totalkm ? 'invalid' : ''} type="text">
      <span hidden={!shownErrors.totalkm} class="error">Invalid total distance</span>
    </label>

    <label>
      Trip KM
      <input on:change={() => validateNumber(values.tripkm, 'tripkm')} bind:value={values.tripkm} placeholder="Trip distance" class={shownErrors.tripkm ? 'invalid' : ''} type="text">
      <span hidden={!shownErrors.tripkm} class="error">Invalid trip distance</span>
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
</Section>

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
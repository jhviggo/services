<script lang="ts">
  import { type Refuel } from '@api/refuels';

  export let refuels: Refuel[];
</script>

<h1>Refuels</h1>
<div>
  {#each refuels as refuel}
    <div class="refuel">
      <p>{refuel.cost}{refuel.currency} - {refuel.tripKM}Km</p>
    </div>
  {/each}
</div>

<style>
  p {
    margin: 0;
  }

  .refuel {
    border: 1px solid black;
    border-radius: 0.25rem;
    margin: 0.5rem;
  }
</style>

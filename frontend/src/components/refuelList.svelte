<script lang="ts">
  import { type Refuel } from '@api/refuels';
  import RefuelListfolder from './refuelListfolder.svelte';

  export let refuels: Refuel[];
</script>

<div class="container">
  <div class="scroller">
    {#each refuels as refuel}
      <RefuelListfolder {refuel} />
    {/each}
  </div>
</div>

<style>
  .container {
    position: relative;
  }

  .scroller {
    max-height: 50dvh;
    overflow-y: auto;
  }

  .container:before {
    content: '';
    background: /* Shadow TOP */
      radial-gradient(
        farthest-side at 50% 0,
        rgba(0, 0, 0, 0.2),
        rgba(0, 0, 0, 0)
      ) center top;
    height: 20px;
    display: block;
    position: absolute;
    width: 100%;
  }

  .container:after {
    content: '';
    background: /* Shadow BOTTOM */
      radial-gradient(
        farthest-side at 50% 100%,
        rgba(0, 0, 0, 0.2),
        rgba(0, 0, 0, 0)
      ) center bottom;
    height: 20px;
    display: block;
    position: absolute;
    margin-top: -20px;
    width: 100%;
  }

  .scroller:before {
    content: '';
    background: white;
    display: block;
    position: sticky;
    height: 20px;
  }

  .scroller:after {
    content: '';
    background: white;
    display: block;
    position: sticky;
    height: 20px;
    z-index: 1;
  }
</style>

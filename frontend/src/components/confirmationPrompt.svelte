<div class="backdrop" aria-hidden="true" on:click={no}></div>

<div class="confirmation">
  <div class="box">
    <h2>{ message }</h2>
    <div class="box-buttons">
      <button class="btn btn-primary" on:click={yes}>Yes</button>
      <button class="btn btn-danger" on:click={no}>No</button>
    </div>
  </div>
</div>

<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  export let message: string;
  
  const dispatch = createEventDispatcher();

  function yes() {
    dispatch('yes');
  }

  function no() {
    dispatch('no');
  }
</script>

<style>
  .backdrop {
    position: fixed;
    width: 100dvw;
    height: 100dvh;
    background-color: rgba(0, 0, 0, 0.75);
    z-index: 110;
    top: 0;
  }

  .confirmation {
    z-index: 120;
    width: 100dvw;
    height: 100dvh;
    position: absolute;
    top: 0;
    display: flex;
    justify-content: center;
    align-items: center;

    & .box {
      width: 90dvw;
      height: 50dvh;
      max-width: 30rem;
      max-height: 15rem;
      background-color: white;
      display: grid;
      justify-content: center;
      border-radius: 0.2rem;
      border: 1px solid var(--primary-blue);

      & h2 {
        color: var(--primary-blue);
        text-align: center;
      }
    }
  }

  .box-buttons {
    display: grid;
    grid-template-columns: 1fr 1fr;
    padding: 2rem;
    gap: 0.5rem;
  }
</style>

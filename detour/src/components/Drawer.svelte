<div class="drawer" style={`${side}: 0;`} class:open={open}>
  <div class="header">
    <div>
      <slot name="header" />
    </div>
    <button class="close" on:click={close}>
      <icon class="material-symbols-outlined">
        close
      </icon>
    </button>
  </div>
  <slot />
</div>
<div class="shadow" class:open={open} on:click={close} aria-hidden="true"></div>

<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let open: boolean = false;
  export let side: 'left' | 'right' = 'right';

  function close() {
    dispatch('close');
  }
</script>

<style>
  .drawer {
    display: none;
    position: fixed;
    height: 100%;
    top: 0;
    background-color: var(--blue);
    z-index: 20;
  }

  .header {
    display: grid;
    grid-template-columns: 1fr auto;
    height: var(--header-height);
  }

  .header div {
    max-height: var(--header-height);
  }

  .close {
    border: unset;
    background-color: unset;
    padding: 0.5rem;
    padding-right: 1.5rem;
  }

  .close icon {
    color: white;
  }

  .shadow {
    display: none;
    position: fixed;
    background-color: black;
    opacity: 0.5;
    width: 100%;
    height: 100%;
    z-index: 19;
    top: 0;
    left: 0;
  }

  .open {
    display: block;
  }
</style>

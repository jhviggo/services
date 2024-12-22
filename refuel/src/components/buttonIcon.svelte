<script lang="ts">
  export let icon: string | undefined = undefined;
</script>

<button type="submit">
  {#if icon && icon !== ''}
    <span class="material-symbols-outlined icon">{ icon }</span>
  {/if}
  <span class="text">
    <slot />
  </span>
</button>

<style>
  button {
    padding: 0;
    width: 10rem;
    background-color: white;
    color: black;
    border: 2px solid var(--primary-blue);
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  button:hover {
    background-color: rgba(1, 95, 191, 0.1);
    cursor: pointer;
  }

  span {
    padding: 0.5rem;
  }

  .icon {
    border-right: 1px solid var(--primary-blue);
  }

  .text {
    display: flex;
    height: 100%;
    width: 100%;
    justify-content: center;
  }
</style>

<script lang="ts">
  import { errorStore, removeError } from '@lib/errorHandler';
</script>

<div class="toast-container">
  {#each $errorStore as error}
    <button class="toast {error.level}" on:click={() => removeError(error)}>
      <div>{ error.message }</div>
      <div>
        <span class="material-symbols-outlined icon">close</span>
      </div>
    </button>
  {/each}
</div>

<style>
  .toast-container {
    position: fixed;
    top: 0;
    right: 0;
  }

  .toast {
    margin: 1rem;
    padding: 1rem;
    border-radius: 0.25rem;
    background-color: var(--debug);
    display: grid;
    grid-template-columns: 1fr auto;
    border: unset;
    align-items: center;
    color: white;
  }

  .debug {
    background-color: var(--debug);
  }

  .info {
    background-color: var(--info);
  }

  .warn {
    background-color: var(--warn);
  }

  .error {
    background-color: var(--error);
  }

  .critical {
    background-color: var(--error);
  }
</style>

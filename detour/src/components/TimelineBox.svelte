<div class="timeline">
  <div class="time">
    <div class="vline"></div>
    {#if location}
      <div class="hline-text">
        <div class="location">{location}</div>
        <span>{city}</span>
      </div>
    {:else}
      <h1 class="header">{header}</h1>
    {/if}
  </div>
  <div class:spaced={spaced}><slot /></div>
</div>

<script lang="ts">
  export let city: string = '';
  export let location: string = '';
  export let header: string = '';
  export let spaced: boolean = false;
</script>

<style>
  .timeline {
    display: grid;
    grid-template-columns: var(--timeline-width) 1fr;
    justify-content: center;
  }

  .time {
    display: grid;
    grid-template-columns: 1fr 9fr;
    background-color: var(--orange-transparent);
  }

  @media (max-width: 768px) {
    .timeline {
      grid-template-columns: 1fr;
    }

    .time {
      display: none;
    }
  }

  .vline {
    border-right: 2px solid white;
  }

  .hline-text {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
  }

  .hline-text .location {
    width: 100%;
    text-align: center;
    font-weight: bold;
    border-bottom: 2px solid white;
  }

  .header {
    text-align: center;
    margin: 0;
  }

  .spaced {
    margin-top: 2rem;
    margin-bottom: 2rem;
  }
</style>

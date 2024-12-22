<div class="cookie-container">
  <div class="cookie">
    <div>
      This site uses cookies for Google Analytics so we can see how many users we have 😺
      <p>
        If you're okay with that please hit the OK button! Otherwise please click the No button.
      </p>
      <p>Google Uses the following cookies:</p>
      <ul>
        <li><strong>_ga</strong>: Used to distinguish users (expires after 2 years)</li>
        <li><strong>_ga_(container-id)</strong>: Used to persist session state (expires after 2 years)</li>
      </ul>
    </div>
    <div class="button-container">
      <button class="btn no" on:click={no}>No</button>
      <button class="btn ok" on:click={ok}>OK</button>
    </div>
  </div>
</div>

<script lang="ts">
  import { COOKIE_ACCEPTED, COOKIE_DENIED, cookiesStore } from '$lib/cookies';

  function no() {
    cookiesStore.set({
      status: COOKIE_DENIED,
    });
  }

  function ok() {
    cookiesStore.set({
      status: COOKIE_ACCEPTED,
    });
  }
</script>

<style>
  .cookie-container {
    position: fixed;
    bottom: 0;
    left: 0;
    background-color: var(--blue);
    width: 100%;
  }

  .cookie {
    display: grid;
    padding: 1rem;
    grid-template-columns: 1fr auto;
  }

  p, div, li, strong {
    color: white;
  }

  .button-container {
    display: grid;
    grid-template-columns: auto auto;
    align-items: end;
    gap: 1rem;
  }

  .btn {
    border: unset;
    padding: 0.5rem 2rem 0.5rem 2rem;
    font-weight: bold;
    font-size: 1.25rem;
  }

  .btn:hover {
    cursor: pointer;
  }

  .ok {
    background-color: var(--orange);
  }

  .no {
    background-color: var(--off-blue);
    color: white;
  }
</style>

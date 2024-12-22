<header>
  <a class="logo" href="/">
    <img src="/images/logo.jpg" alt="Viggo's detour logo">
  </a>
  <div class="links">
    {#each headers as header}
      <a href="{header.to}">
        <icon class="material-symbols-outlined">
          {header.icon || ''}
        </icon>
        {header.text}
      </a>
    {/each}
  </div>
  <button class="burger" on:click={toggleMenu}>
    <icon class="material-symbols-outlined">
      menu
    </icon>
  </button>
  <Drawer open={menuOpen} on:close={toggleMenu}>
    {#each headers as header}
    <div class="drawer-links">
      <a href="{header.to}" on:click={() => menuOpen = false}>
        <icon class="material-symbols-outlined">
          {header.icon || ''}
        </icon>
        {header.text}
      </a>
    </div>
    {/each}
  </Drawer>
</header>


<script lang="ts">
  import { headers } from '$lib/header';
  import Drawer from './Drawer.svelte';

  let menuOpen: boolean = false;

  function toggleMenu() {
    menuOpen = !menuOpen;
  }
</script>

<style>
  header {
    position: fixed;
    top: 0;
    left: 0;
    z-index: 10;
    display: grid;
    width: 100%;
    background-color: var(--blue);
    grid-template-columns: var(--timeline-width) 1fr;
    align-items: center;
    justify-items: end;
  }
  
  .logo {
    display: grid;
    justify-content: center;
    background-color: white;
    width: 100%;
    padding-top: 0.5rem;
    padding-bottom: 0.5rem;
  }

  img {
    height: calc(var(--header-height) - 1rem);
  }

  .links {
    display: flex;
    flex-direction: row;
    gap: 1rem;
    margin-right: 3rem;
  }

  .drawer-links a {
    padding-right: 2rem;
    padding-left: 2rem;
    padding-top: 0.5rem;
    padding-bottom: 0.5rem;
  }

  .burger {
    margin-right: 2rem;
    display: none;
    background-color: unset;
    border: unset;
  }

  @media (max-width: 768px) {
    .links {
      display: none;
    }

    .burger {
      display: block;
    }
  }

  a {
    display: flex;
    gap: 0.25rem;
    color: white;
    text-decoration: unset;
    font-weight: bold;
  }

  a:hover, a:hover > icon {
    color: var(--orange)
  }

  icon {
    color: white;
  }
</style>

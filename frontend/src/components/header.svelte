<script lang="ts">
  import { goto } from '$app/navigation';
  import { page } from '$app/stores';
  import { userStore, logout } from '@api/users';

  let isMenuOpen: boolean = true;

  page.subscribe(() => {
    isMenuOpen = false;
  })

  function toggleMenu() {
    isMenuOpen = !isMenuOpen;
  }

  function handleLogout() {
    logout();
    goto('/');
    isMenuOpen = false;
  }
</script>

<header>
  <a class="logo" href="/">
    <img src="/logo.png" alt="page logo"/>
    [services]
  </a>
  <button class="burger" on:click={toggleMenu}>
    <icon class="material-symbols-outlined">
      menu
    </icon>
  </button>
</header>

{#if isMenuOpen}
  <div class="backdrop" aria-hidden="true" on:click={toggleMenu}></div>
  <div class="menu">
    <div class="header">
      <span class="bold">{ $userStore.name || 'Menu' }</span>
      <button class="burger" on:click={toggleMenu}>
        <icon class="material-symbols-outlined">
          close
        </icon>
      </button>
    </div>
    <div class="items">
      <a href="/" class="btn btn-primary justify-left">
        <span class="material-symbols-outlined">
          house
        </span>
        Home
      </a>
      {#if $userStore?.token}
        <span class="bold">Refuel</span>
        <a href="/refuels" class="btn btn-primary justify-left">
          <span class="material-symbols-outlined">
            local_gas_station
          </span>
          Refuels
        </a>
        <a href="/dashboard" class="btn btn-primary justify-left">
          <span class="material-symbols-outlined">
            analytics
          </span>
          Dashboard
        </a>
      {/if}
      <span class="bold">User</span>
      {#if $userStore?.token}
        <a href="/profile" class="btn btn-primary justify-left">
          <span class="material-symbols-outlined">
            account_circle
          </span>
          Profile
        </a>
        <button class="btn btn-danger justify-left" on:click={handleLogout}>
          <span class="material-symbols-outlined">
            lock
          </span>
          Logout
        </button>
      {:else}
        <a href="/login" class="btn btn-primary justify-left">
          <span class="material-symbols-outlined">
            lock
          </span>
          Login
        </a>
      {/if}
    </div>
  </div>
{/if}

<style>
  header {
    display: flex;
    height: 4rem;
    justify-content: space-between;
    align-items: center;
    box-shadow: 0 4px 6px -6px #222;
    background-color: white;
  }

  .logo img {
    height: 3rem;
  }

  .logo {
    display: flex;
    padding: 0.5rem;
    align-items: center;
    gap: 0.5rem;
    font-size: 1.5rem;
    text-decoration: none;
    color: var(--primary-blue);
  }

  .burger {
    background-color: unset;
    border: unset;
    margin-right: 1rem;
  }

  .burger icon {
    font-size: 2rem;
  }

  .backdrop {
    position: fixed;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.75);
    z-index: 98;
  }

  @keyframes slide-in {
    from {right: -100%;}
    to {right: 0;}
  }

  .menu {
    position: fixed;
    z-index: 99;
    top: 0;
    right: 0;
    height: 100%;
    width: 70dvw;
    background-color: white;
    animation-name: slide-in;
    animation-duration: 0.25s;
  }

  .menu .header {
    display: flex;
    justify-content: space-between;
    padding-top: 1rem;
    padding-left: 1rem;
  }

  .menu .bold {
    font-weight: bold;
  }

  .menu .items {
    display: flex;
    flex-direction: column;
    padding-left: 0.5rem;
    padding-right: 0.5rem;
  }
  
  .btn {
    margin-bottom: 0.5rem;
  }

  .menu .items .link span {
    align-items: baseline;
    margin-right: 0.25rem;
  }
</style>

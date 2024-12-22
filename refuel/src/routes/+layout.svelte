<script lang="ts">
  import Header from '@components/header.svelte';
  import ErrorToast from '@components/errorToast.svelte';
  import { loadStoredUser } from '@api/users';
  import { beforeUpdate, } from 'svelte';

  let loaded = false;
  beforeUpdate(() => {
    if (!loaded) {
      loadStoredUser();
      loaded = true;
    }

    if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition((e) => {
      console.log('cool', e)
    });
  } else {
    console.log("Geolocation is not supported by this browser.");
  }
  });
</script>

<main>
  <Header />
  <div>
    <slot />
  </div>
  <footer>
    viggo.work
  </footer>
</main>

<ErrorToast />

<style>
  main {
    height: 100%;
    display: grid;
    grid-template-rows: auto 1fr auto;
    background-color: #eee;
  }

  footer {
    display: flex;
    justify-content: center;
    padding: 1rem;
  }
</style>
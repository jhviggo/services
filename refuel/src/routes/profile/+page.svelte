<Section>
  <div class="page">
    <h1>Profile page</h1>
    
    <p><span class="bold">Name: </span>{$userStore.name}</p>
    <p><span class="bold">ID: </span>{$userStore.id}</p>
    <p><span class="bold">Username: </span>{$userStore.username}</p>
  </div>
</Section>

<script lang="ts">
  import { userStore } from '@api/users'
  import Section from '@components/section.svelte';
</script>

<style>
  .page {
    padding: 1rem;
  }
</style>

<div class="page">
  <Banner image="/images/countries-banner.webp" imageAlt="collage showing differnet countries I've been to" />
  <div class="text">
    {@html text}
  </div>
</div>

<script lang="ts">
  import { onMount } from 'svelte';
  import { marked } from 'marked';
  import { getContentFile } from '$lib/api';
  import Banner from '$components/Banner.svelte';

  let text: string = '';

  onMount(async () => {
    const data = await getContentFile('page-about.md');
    text = await marked.parse(data);
  });
</script>

<style>
  .text {
    text-align: left;
  }

  :global(h1, h2, h3, h4, h5, h6) {
    color: var(--orange);
  }
</style>
<div class="side-by-side {!image || !content ? 'single-grid' : 'double-grid'}" id={itemId}>
  {#if image && !flipped}
    <img src={image} alt="">
  {/if}
  {#if content}
    <div class="content">
      {#if content}
        {@html parsedContent}
      {:else}
        <slot />
      {/if}
    </div>
  {/if}
  {#if image && flipped}
    <div class="picture">
      <img src={image} alt="">
      <slot name="hover" />
    </div>
  {/if}
</div>

<script lang="ts">
  import { getContentFile } from '$lib/api';
  import { marked } from 'marked';
  import { onMount } from 'svelte';

  export let image: string = '';
  export let content: string = '';
  export let flipped: boolean = false;
  export let itemId: string = '';

  let parsedContent: string;

  onMount(async () => {
    if (content !== '') {
      const data = await getContentFile(content);
      parsedContent = await marked.parse(data);
    }
  });
</script>

<style>
  .side-by-side {
    display: grid;
    background-color: white;
    scroll-margin-top: 5rem;
    margin: 0.5rem;
    border-radius: 0.5rem;
    overflow: hidden;
    border-bottom: 2px solid var(--blue);
    border-right: 2px solid var(--blue);
  }
  
  .single-grid { grid-template-columns: 1fr; }
  .double-grid { grid-template-columns: repeat(2, 1fr); }
  
  @media (max-width: 768px) {
    .double-grid {
      grid-template-rows: repeat(2, auto);
      grid-template-columns: 1fr;
    }
  }
  
  img {
    width: 100%;
  }
  
  .picture {
    background-color: var(--blue);
    position: relative;
    display: grid;
    align-items: center;
  }
  
  .content {
    background-color: white;
    padding: 0.5rem;
    overflow-y: auto;
  }
</style>

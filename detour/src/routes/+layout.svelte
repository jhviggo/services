<svelte:head>
  {#if import.meta.env.VITE_ENVIRONMENT !== 'dev'}
    <!-- Google tag (gtag.js) -->
    <script async src="https://www.googletagmanager.com/gtag/js?id=G-1JXNZFG1DP"></script>
    <script>
      window.dataLayer = window.dataLayer || [];
      function gtag(){dataLayer.push(arguments);}
      gtag('js', new Date());
      gtag('config', 'G-1JXNZFG1DP');
    </script>
  {/if}

  <link rel="stylesheet" type="text/css" href="theme.css?{Math.random().toString(36).slice(2, 18)}" />

  <meta name="description" content="My motorcycle trip from Denmark To Japan. My name is Viggo, I'm a software engineer who decided to take my Master's in Japan and will be driving there. I'll be documenting how much it costs, how to get the carnet de passage and all the details on the way">
  <meta name="keywords" content="Travel, Motorcycle, Camping, Software Engineer, Viggo">
  <meta name="author" content="Viggo Petersen">

  <meta property="og:title" content="Viggo's Detour" />
  <meta property="og:description" content="My motorcycle trip from Denmark To Japan. My name is Viggo, I'm a software engineer who decided to take my Master's in Japan and will be driving there. I'll be documenting how much it costs, how to get the carnet de passage and all the details on the way" />
  <meta property="og:image" content="https://viggosdetour.com/images/map.webp" />
  <meta property="og:image:alt" content="World map showing a route from Denmark to India, zig-zagging through many different countries on the way" />

  <meta name="twitter:title" content="Viggo's Detour" />
  <meta name="twitter:description" content="My motorcycle trip from Denmark To Japan. My name is Viggo, I'm a software engineer who decided to take my Master's in Japan and will be driving there. I'll be documenting how much it costs, how to get the carnet de passage and all the details on the way" />
  <meta name="twitter:url" content="https://viggosdetour.com/" />
  <meta name="twitter:image:src" content="https://viggosdetour.com/images/map.webp" />
  <meta name="twitter:image:alt" content="World map showing a route from Denmark to India, zig-zagging through many different countries on the way" />

  <title>Viggo's Detour</title>
</svelte:head>

<div class="main-layout">
  <div>
    <Header />
  </div>
  <div class="main-content">
    <slot />
  </div>
  <Footer />
</div>

<script lang="ts">
  import Header from '$components/Header.svelte';
  import Footer from '$components/Footer.svelte';
  import { onMount } from 'svelte';
  import { getVisitorCookie, setVisitorCookie } from '$lib/cookies';
  import { registerVisitor } from '$lib/api';
  import { v4 as uuidv4 } from 'uuid';

  onMount(() => {
    setTimeout(() => {
      const visitorId = getVisitorCookie();
      if (visitorId == '') {
        const newVisitorId = uuidv4();
        setVisitorCookie(newVisitorId);
        registerVisitor(newVisitorId);
      }
    }, 1000);
  });
</script>

<style>
  .main-layout {
    display: grid;
    grid-template-rows: var(--header-height) 1fr auto;
    min-height: 100%;
  }
</style>

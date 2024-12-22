<div class="banner">
  <img src={image} alt={imageAlt}>
  <div class="overlay">
    <div>
      <slot />
    </div>
  </div>
</div>

<script lang="ts">
  export let image: string;
  export let imageAlt: string = '';
</script>

<style>
  .banner {
    display: grid;
    position: relative;
    width: 100%;
  }

  .banner img {
    width: 100%;
  }

  .overlay {
    display: grid;
    position: absolute;
    height: 100%;
    width: 100%;
    top: 0;
    left: 0;
  }
</style>

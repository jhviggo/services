<main>
  <header>
    <a class="logo" href="/">
      <img src="/logo.png" alt="page logo"/>
      [Refuel service]
    </a>
    <div class="header-items">
      <a href="refuels">
        <span class="material-symbols-outlined">
          local_gas_station
        </span>
        Refuels
      </a>
      <div class="spacer"></div>
      <a href="about">
        <span class="material-symbols-outlined">
          analytics
        </span>
        Dashboard
      </a>
      <div class="spacer"></div>
      <a href="profile">
        <span class="material-symbols-outlined">
          account_circle
        </span>
        Profile
      </a>
    </div>
  </header>
  <div>
    <slot />
  </div>
  <footer>
    refuels.viggo.work
    <div class="spacer"></div>
    asd
  </footer>
</main>

<style>
  main {
    height: 100%;
    display: grid;
    grid-template-rows: auto 1fr auto;
    background-color: #eee;
  }

  header {
    display: flex;
    height: 4rem;
    align-items: end;
    justify-content: space-between;
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

  .header-items {
    display: flex;
    margin-right: 2rem;
    height: 100%;
    align-items: center;
  }

 .spacer::after {
    content: '•';
    color: #ccc;
    margin: 0.5rem 0.5rem;
    pointer-events: none;
  }

  .header-items a {
    font-size: large;
    text-decoration: none;
    color: black;
    display: flex;
  }

  .header-items a:hover {
    margin-top: 3px;
    border-bottom: 3px solid var(--primary-blue);
  }

  .header-items a:hover span {
    color: var(--primary-blue);
  }

  footer {
    display: flex;
    justify-content: center;
    padding: 1rem;
  }
</style>
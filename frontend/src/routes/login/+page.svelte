<script lang="ts">
  import ButtonIcon from '@components/buttonIcon.svelte';
  import Section from '@components/section.svelte';
  import { login } from '@api/users';
  import { goto } from '$app/navigation';

  const user = {
    username: '',
    password: '',
  };

  const shownErrors = {
    username: '',
    password: '',
  };

  async function doLogin() {
    const loggedIn = await login(user.username, user.password);
    if (loggedIn) {
      goto('/');
    }
  }
</script>

<Section>
  <h1>Login</h1>
  <hr>
  <form on:submit|preventDefault={doLogin} action="POST">
    <label>
      username
      <input
        bind:value={user.username}
        placeholder="Username"
        class={shownErrors.username ? 'invalid' : ''}
        required
        type="text">
      <span hidden={!shownErrors.username} class="error">Invalid total distance</span>
    </label>

    <label>
      Password
      <input
        bind:value={user.password}
        placeholder="password"
        class={shownErrors.password ? 'invalid' : ''}
        required
        type="password">
      <span hidden={!shownErrors.password} class="error">Invalid trip distance</span>
    </label>

    <ButtonIcon icon="lock">Login</ButtonIcon>
  </form>
</Section>

<style>
  form {
    display: flex;
    flex-direction: column;
    max-width: 25rem;
  }

  label {
    display: flex;
    flex-direction: column;
    width: 100%;
  }

  input {
    padding: 0.5rem 0.75rem;
    margin-bottom: 1.25rem;
  }

  .error {
    font-size: 0.85rem;
    height: 1.25rem;
  }

  .invalid {
    border-color: var(--error);
    margin-bottom: 0;
  }

  .invalid:focus-visible {
    outline-color: var(--error);
  }
</style>

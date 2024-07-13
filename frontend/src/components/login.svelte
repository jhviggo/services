<script lang="ts">
  import { login } from '@api/users';
  import { goto } from '$app/navigation';

  let username:string = '';
  let password: string = '';

  let usernameError: string = '';
  let passwordError: string = '';
  let authError: string = '';

  async function auth() {
    usernameError = username ? '' : 'Missing username';
    passwordError = password ? '' : 'Missing password';

    try {
      login(username, password);
      goto("/")
    } catch (e) {
      authError = 'Wrong username or password';
    }

  }
</script>

<form class="login-container">
  <label>
    Username
    <input required bind:value={username} placeholder="Username" type="text">
  </label>
  <span class="error">{usernameError}</span>

  <label>
    Password
    <input required bind:value={password} placeholder="Password" type="password">
  </label>
  <span class="error">{passwordError}</span>

  <button on:click={auth}>Login</button>
  <span class="error">{authError}</span>
</form>

<style>
  .login-container {
    display: flex;
    flex-direction: column;
  }
</style>

<script lang="ts">
  import { goto } from "$app/navigation";
  import Button, { Label } from "@smui/button";
  import type { User } from "../domain/user";
  export let path: string;
  export let user: User|null;
</script>

<nav>
  <div class="mdc-typography--headline2 title">Generate ORiginal Memo 1</div>
  {#if !user}
    <div style="text-align: right;">
      {#if path !== "login"}
        <Button on:click={() => goto("/login")}>
          <Label>Login</Label>
        </Button>
      {/if}
      {#if path !== "signup"}
        <Button on:click={() => goto("/signup")}>
          <Label>Signup</Label>
        </Button>
      {/if}
    </div>
  {/if}
</nav>

<style>
  nav {
    border-bottom: 1px solid rgba(255, 62, 0, 0.1);
    font-weight: 300;
    padding: 0 1em;
  }
  .title {
    margin: 0 auto;
    text-align: center;
  }
</style>

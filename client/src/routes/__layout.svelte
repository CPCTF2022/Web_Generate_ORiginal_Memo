<script lang="ts" context="module">
	export const load: Load = async ({ fetch }) => {
    const res = await fetch("/api/users/me", {
      method: "GET",
      headers: {
        Accept: "application/json",
      },
    });

    return {
      props: {
        status: res.status,
        userInfo: await res.json(),
      },
    };
  }
</script>

<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from '$app/stores';
  import type { Load } from "@sveltejs/kit";
  import Nav from "../components/Nav.svelte";
  import type { User } from "../domain/user";
  import { user } from "../store";

  export let status: number;
  export let userInfo: User;

  if (status === 200 && userInfo) {
    user.set(userInfo);
  }

  let path: string;
  page.subscribe((value) => {
    path = value.url.pathname;
  });
</script>

<Nav {path} user={userInfo} />

<main>
  <slot />
</main>

<style>
  main {
    position: relative;
    min-width: 50%;
    max-width: 100%;
    width: 750px;
    background-color: white;
    padding: 2em;
    margin: 0 auto;
    box-sizing: border-box;
  }
</style>

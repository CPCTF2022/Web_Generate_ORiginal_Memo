<script lang="ts" context="module">
	export const load: Load = async ({ fetch }) => {
    const res = await fetch("/api/users/me", {
      method: "GET",
      headers: {
        Accept: "application/json",
      },
    });

    let userInfo: User| null = null;
    if (res.status == 200) {
      userInfo = await res.json();
    }

    return {
      props: {
        userInfo,
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

  export let userInfo: User|null;

  let path: string = "";
  page.subscribe((value) => {
    path = value.url.pathname;
  });

  if (userInfo) {
    user.set(userInfo);
  } else if (path && path !== "/signup") {
    goto("/login");
  }

  let userValue: User|null = null;
  user.subscribe((value) => {
    userValue = value ?? userValue;
  });
</script>

<Nav {path} user={userValue} />

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

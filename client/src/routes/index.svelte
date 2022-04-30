<script lang="ts" context="module">
	export const load: Load = async ({ fetch }) => {
    const res = await fetch(`/api/memos`, {
      method: "GET",
      headers: {
        Accept: "application/json",
      },
    });

    let memos: Memo[];
    if (res.status == 200) {
      memos = (await res.json()).memos;
    } else {
      memos = [];
    }

    return {
      props: {
        memos,
      },
    };
  }
</script>

<script lang="ts">
  import type { Load } from "@sveltejs/kit";
  import type { Memo } from "../domain/memo";
  import Mess from "../components/Memo.svelte";
  import Input from "../components/Input.svelte";
  import { user } from "../store";
  import type { User } from "../domain/user";
  export let memos: Memo[];

  let memoValues: Memo[] = memos;
  const sendHandler = async (event: any) => {
    const response = await fetch("/api/memos", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify({ content: event.detail.text }),
    });
    if (response.status === 201) {
      const memo: Memo = await response.json();
      memoValues = [memo, ...memoValues];
    }
  };

  let userInfo: User| null = null;
  user.subscribe((value) => {
    userInfo = value ?? userInfo;
  });
</script>

{#if userInfo}
  <Input on:send={sendHandler} />
{/if}
<br />
{#each memoValues as memo}
  <Mess
    memo={memo}
  />
{/each}

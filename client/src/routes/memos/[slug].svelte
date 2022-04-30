<script lang="ts" context="module">
	export const load: Load = async ({ fetch, params }) => {
    const memoID: string = params.slug;

    const res = await fetch(`/api/memos/${memoID}`, {
      method: "GET",
      headers: {
        Accept: "application/json",
      },
    });

    const memo: Memo = (await res.json()).memo;
    return {
      props: {
        memo,
      },
    };
  }
</script>

<script lang="ts">
  import type { Load } from "@sveltejs/kit";
  import type { Memo } from "../../domain/memo";
  import Mess from "../../components/Memo.svelte";
  export let memo: Memo;
</script>

{#if memo}
  <Mess
    memo={memo}
  />
{:else}
  Invalid Memo
{/if}

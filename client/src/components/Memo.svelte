<script lang="ts">
  import { goto } from "$app/navigation";
  import Card, { Content } from "@smui/card";
  import type { Memo } from "../domain/memo";
  export let memo: Memo;
</script>

<Card
  style="width: 100%;margin:0 auto;"
  on:click={ () => goto(`/memos/${memo.id}`)}
>
  <Content>{memo?.content}</Content>
  <div class="mdc-typography--subtitle1 time">{memo.createdAt}</div>
</Card>

<style>
  .time {
    padding: 0 15px;
  }
</style>

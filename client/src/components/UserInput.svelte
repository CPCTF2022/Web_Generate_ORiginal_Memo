<script lang="ts">
  import Textfield from "@smui/textfield";
  import Button, { Label } from "@smui/button";
  import { createEventDispatcher } from "svelte";
  export let label: string;
  let name: string = "";
  let password: string = "";

  const dispatch = createEventDispatcher();

  const handleSubmit = () => {
    dispatch("submit", { name, password });
  };
</script>

<div style="width:100%;text-align:center;">
  <form on:submit|preventDefault={handleSubmit} method="post">
    <Textfield
      bind:value={name}
      type="text"
      label="Name"
      style="margin-top:5px;width:50%;"
    />
    <br />
    <Textfield
      bind:value={password}
      type="password"
      label="Password"
      style="margin-top:5px;width:50%;"
    />
    <br />
    <Button style="margin-top:10px;">
      <Label>{label}</Label>
    </Button>
  </form>
</div>

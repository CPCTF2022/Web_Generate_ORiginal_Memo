<script>
  import { createEventDispatcher } from "svelte";
  import Textfield from "@smui/textfield";
  import Button, { Label } from "@smui/button";
  let text = "";

  const dispatch = createEventDispatcher();
</script>

<div style="width: 100%;margin:0 auto;text-align:right;">
  <Textfield
    textarea
    bind:value={text}
    label="Label"
    style="width: 100%;"
    helperLine$style="width: 100%;"
  />
  <br />
  <Button
    on:click={() => {
      dispatch("send", { text });
      text = "";
    }}
  >
    <Label>Send</Label>
  </Button>
</div>

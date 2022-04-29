<script lang="ts">
  import { goto } from "$app/navigation";
  import type { User } from "../domain/user";
  import UserInput from "../components/UserInput.svelte";
  import { user } from "../store";

  const handleSignup = async (event: any) => {
    const response = await fetch("/api/signup", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify(event.detail),
    });
    if (response.status === 201) {
      const userInfo: User = await response.json();
      if (userInfo) {
        user.set(userInfo);
      }
      goto("/");
    }
  };
</script>

<UserInput label="Signup" on:submit={handleSignup} />

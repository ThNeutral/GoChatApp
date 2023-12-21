<script lang="ts">
	import { loginUser } from '$lib/queries';
	import type { PageData } from './$types';

	export let data: PageData;

	async function handleSubmit(
		event: SubmitEvent & {
			currentTarget: EventTarget & HTMLFormElement;
		}
	) {
		const target = event.target as HTMLFormElement;
		const inputData = {
			email: (target[0] as HTMLInputElement).value,
			password: (target[1] as HTMLInputElement).value
		};
		const json = await loginUser(inputData);
        if ("access_token" in json) {
            document.cookie = `Authorization=Bearer ${json.access_token}; HttpOnly; Path=/; Max-Age=2592000; SameSite=None`
        }
	}
</script>

<h1>{data.Authorization}</h1>
<button on:click={() => (document.cookie = 'test=cookie; HttpOnly')}>set cookies</button>
<form on:submit|preventDefault={handleSubmit}>
	<input name="email" required />
	<input name="password" required />
	<button type="submit">test</button>
</form>

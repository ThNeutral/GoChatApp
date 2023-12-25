<script lang="ts">
	import { getUserProfile, loginUser, registerUser } from '$lib/queries';

	async function handleRegisterSubmit(
		event: SubmitEvent & {
			currentTarget: EventTarget & HTMLFormElement;
		}
	) {
		const target = event.target as HTMLFormElement;
		const inputData = {
			username: (target[0] as HTMLInputElement).value,
			email: (target[1] as HTMLInputElement).value,
			password: (target[2] as HTMLInputElement).value
		};
		const json = await registerUser(inputData);
	}

	async function handleLoginSubmit(
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
	}

	async function handleGetUserSubmit(
		event: SubmitEvent & {
			currentTarget: EventTarget & HTMLFormElement;
		}
	) {
		const json = await getUserProfile();
		console.log(json);
	}
</script>

<form on:submit|preventDefault={handleRegisterSubmit}>
	<input name="username" required />
	<input name="email" required />
	<input name="password" required />
	<button type="submit">test</button>
</form>

<form on:submit|preventDefault={handleLoginSubmit}>
	<input name="email" required />
	<input name="password" required />
	<button type="submit">test</button>
</form>

<form on:submit|preventDefault={handleGetUserSubmit}>
	<button type="submit">test</button>
</form>

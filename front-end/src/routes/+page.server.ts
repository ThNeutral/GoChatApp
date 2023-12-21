export function load({ cookies }) {
	const cookie = cookies.get('Authorization');

	return {
		Authorization: cookie
	};
}
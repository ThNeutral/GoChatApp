const backendURL = 'http://localhost:8000/v1';

type LoginUserParams = {
	email: string;
	password: string;
};

type LoginUserResponseType = {
	error: string;
} & {
	message: string;
	access_token: string;
};

const loginUserEndpoint = '/login-user';

export async function loginUser(body: LoginUserParams): Promise<LoginUserResponseType> {
	const request = await fetch(backendURL + loginUserEndpoint, {
		method: 'POST',
		body: JSON.stringify(body),
        credentials: 'include'
	});
	return (await request.json()) as LoginUserResponseType;
}

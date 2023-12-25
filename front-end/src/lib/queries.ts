const backendURL = 'http://localhost:8000/v1';

type RegisterUserResponseType = {
	error: string;
} & {
	message: string;
};

type RegisterUserParams = {
	username: string;
	email: string;
	password: string;
};

const registrationUserEndpoint = '/register-user';

export async function registerUser(body: RegisterUserParams): Promise<RegisterUserResponseType> {
	const request = await fetch(backendURL + registrationUserEndpoint, {
		method: 'POST',
		body: JSON.stringify(body),
		credentials: 'include'
	});
	return (await request.json()) as RegisterUserResponseType;
}

const loginUserEndpoint = '/login-user';

type LoginUserParams = {
	email: string;
	password: string;
};

type LoginUserResponseType = {
	error: string;
} & {
	message: string;
};

export async function loginUser(body: LoginUserParams): Promise<LoginUserResponseType> {
	const request = await fetch(backendURL + loginUserEndpoint, {
		method: 'POST',
		body: JSON.stringify(body),
		credentials: 'include'
	});
	return (await request.json()) as LoginUserResponseType;
}

const userProfileEndpoint = '/user-profile';

export async function getUserProfile() {
	const request = await fetch(backendURL + userProfileEndpoint, {
		method: 'GET',
		credentials: 'include'
	});
	return await request.json();
}

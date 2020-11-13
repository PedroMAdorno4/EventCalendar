<template>
	<div>
		<h1 class="mb-4"> Cadastro </h1>
		<b-container fluid>
			<b-row>
					<b-form class="w-100" @submit.prevent="submitData">
						<b-form-group
							id="input-group-1"
							label="Nome de usuário:"
							label-for="input-1"
							label-class="text-left"
						>
							<b-form-input
								id="input-1"
								v-model="form.username"
								type="text"
								required
								placeholder="Digite seu nome de usuário"
							></b-form-input>
						</b-form-group>
							<b-form-group
								id="input-group-2"
								label="Nome:"
								label-for="input-2"
								label-class="text-left"
							>
							<div class="d-flex justify-content-around">
								<b-form-input
									id="input-2"
									class="mr-auto"
									style="width:45%"
									v-model="form.firstname"
									type="text"
									required
									placeholder="Primeiro nome"
								></b-form-input>
								<b-form-input
									id="input-3"
									style="width:45%"
									v-model="form.lastname"
									type="text"
									required
									placeholder="Último nome"
								></b-form-input>
							</div>
							</b-form-group>

						<b-form-group
							id="input-group-4"
							label="Senha:"
							label-for="input-4"
							label-class="text-left"
						>
							<b-form-input
								id="input-4"
								v-model="form.password"
								type="password"
								required
								placeholder="Digite sua senha"
							></b-form-input>
						</b-form-group>

						<b-form-group
							id="input-group-5"
							label="Confirme a senha:"
							label-for="input-5"
							label-class="text-left"
						>
							<b-form-input
								id="input-5"
								v-model="form.passwordConfirmation"
								type="password"
								required
								placeholder="Digite novamente sua senha"
							></b-form-input>
						</b-form-group>

						<b-overlay
							:show="registerStatus === 1"
							rounded
							opacity="0.6"
							spinner-small
							spinner-variant="primary"
						>
							<b-button
								type="submit"
								variant="primary"
								:disabled="registerStatus === 1"
								class="mt-4"
							>Cadastrar</b-button>
						</b-overlay>
					</b-form>
			</b-row>
		</b-container>
	</div>
</template>

<script>
import { mapState, mapActions } from "vuex";

export default {
	data() {
		return {
			form: {
				username: "",
				firstname: "",
				lastname: "",
				password: "",
				passwordConfirmation: ""
			},
			authenticated: false,
			registerStatus: 0
		};
	},
	methods: {
		...mapActions(["setToken"]),
		submitData() {
			this.registerStatus = 1;
			const userInfo = {
				username: this.form.username,
				firstname: this.form.firstname,
				lastname: this.form.lastname,
				password: this.form.password
			};

			fetch("/backend/API/user/create", {
				method: "POST",
				headers: {
					"Content-Type": "application/json"
				},
				body: JSON.stringify(userInfo)
			})
				.then(response => {
					return response;
				})
				.then(response => {
					if(response.status === 200) {
						this.login();
					} else {
						this.registerStatus = 0;
					}
				})
		},
		login() {
			const userInfo = {
				username: this.form.username,
				password: this.form.password
			};

			fetch("/backend/API/auth", {
				method: "POST",
				headers: {
					"Content-Type": "application/json"
				},
				body: JSON.stringify(userInfo)
			})
				.then(response => {
					return response.json();
				})
				.then(token => {
					this.setToken(token);
					this.registerStatus = 0;
					if (this.form.rememberMe == "remember") {
						this.$cookies.set("token", token, "30d", "", "localhost", true, "Strict")
					}
					if (token !== "") {
						this.$router.push({
							name: "Eventos",
							query: { redirect: "/events" }
						});
					}
				});
		},
	},
	computed: {
		...mapState(["token"])
	},
	mounted() {
	}
};
</script>

<style scoped>
.lab {
	text-align: left;
}
</style>
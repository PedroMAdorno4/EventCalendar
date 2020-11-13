<template>
	<div>
		<h1 class="mb-4"> Login </h1>
		<b-container fluid>
			<b-row align-h="center">
					<b-form class="w-100" @submit.prevent="submitLogin">
						<b-form-group
							id="input-group-1"
							label="Nome de usuario:"
							label-for="input-1"
							label-class="text-left"
						>
							<b-form-input
								id="input-1"
								v-model="form.username"
								type="text"
								required
								placeholder="Digite seu nome de usuario"
							></b-form-input>
						</b-form-group>

						<b-form-group
							id="input-group-2"
							label="Senha:"
							label-for="input-2"
							label-class="text-left"
						>
							<b-form-input
								id="input-2"
								v-model="form.password"
								type="password"
								required
								placeholder="Digite sua senha"
							></b-form-input>
						</b-form-group>

						<b-form-checkbox
							id="checkbox-1"
							switch
							v-model="form.rememberMe"
							name="checkbox-1"
							value="remember"
							unchecked-value="forget"
							class="mb-3"
						>
							Lembrar de mim por 30 dias
						</b-form-checkbox>

						<b-overlay
							:show="loginStatus === 1"
							rounded
							opacity="0.6"
							spinner-small
							spinner-variant="primary"
						>
							<b-button
								type="submit"
								variant="primary"
								:disabled="loginStatus === 1"
								class="mt-4"
							>Entrar</b-button>
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
				password: "",
				rememberMe: ""
			},
			authenticated: false,
			loginStatus: 0
		};
	},
	methods: {
		...mapActions(["setToken"]),
		submitLogin() {
			this.loginStatus = 1;
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
					this.loginStatus = 0;
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
		checkCookie() {
			if (this.$cookies.isKey("token")) {
				this.setToken(this.$cookies.get("token"));
				this.$router.push({
					name: "Eventos",
					query: { redirect: "/events" }
				});
			}
		}
	},
	computed: {
		...mapState(["token"])
	},
	mounted() {
		this.checkCookie();
	}
};
</script>
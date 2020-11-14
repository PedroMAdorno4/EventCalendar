<template>
	<div>
		<h1 class="mb-4"> Cadastro </h1>
		<b-container>
			<b-row>
				<b-form
					class="w-100"
					@submit.stop.prevent="submitData"
				>
					<b-form-group
						id="username-group"
						label="Nome de usuário:"
						label-for="username"
						label-class="text-left"
					>
						<b-form-input
							id="username"
							v-model="$v.form.username.$model"
							:state="validateState('username')"
							type="text"
							required
							placeholder="Digite seu nome de usuário"
						></b-form-input>
						<b-form-invalid-feedback id="username-live-feedback">Este campo é obrigatório e o nome de usuário deve ser único</b-form-invalid-feedback>
					</b-form-group>
					<b-form-group
						id="firstname-group"
						label="Nome:"
						label-for="input-2"
						label-class="text-left"
					>
						<div class="d-flex justify-content-between">
							<b-form-input
								id="firstname"
								style="width:48%"
								v-model="$v.form.firstname.$model"
								:state="validateState('firstname')"
								type="text"
								required
								placeholder="Primeiro nome"
							></b-form-input>
							<b-form-input
								id="lastname"
								style="width:48%"
								v-model="form.lastname"
								type="text"
								placeholder="Último nome"
							></b-form-input>
						</div>
						<b-form-invalid-feedback id="firstname-live-feedback">Campo obrigatório.</b-form-invalid-feedback>
					</b-form-group>

					<b-form-group
						id="password-group"
						label="Senha:"
						label-for="password"
						label-class="text-left"
					>
						<b-form-input
							id="password"
							v-model="$v.form.password.$model"
							:state="validateState('password')"
							type="password"
							required
							placeholder="Digite sua senha"
						></b-form-input>
						<b-form-invalid-feedback id="password-live-feedback">Campo obrigatório.</b-form-invalid-feedback>
					</b-form-group>

					<b-form-group
						id="passwordConfirmation-group"
						label="Confirme a senha:"
						label-for="passwordConfirmation"
						label-class="text-left"
					>
						<b-form-input
							id="passwordConfirmation"
							v-model="$v.form.passwordConfirmation.$model"
							:state="validateState('passwordConfirmation')"
							type="password"
							required
							placeholder="Digite novamente sua senha"
						></b-form-input>
						<b-form-invalid-feedback id="passwordConfirmation-live-feedback">As senhas devem ser iguais.</b-form-invalid-feedback>
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
import { required, sameAs } from "vuelidate/lib/validators";

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
	validations: {
		form: {
			username: {
				required,
			},
			firstname: {
				required,
			},
			password: {
				required,
			},
			passwordConfirmation: {
				required: sameAs('password'),
			},
		},
	},
	methods: {
		...mapActions(["setToken"]),
		validateState(name) {
			const { $dirty, $error } = this.$v.form[name];
			return $dirty ? !$error : null;
		},
		makeToast(title, message, type) {
			this.$bvToast.toast(message, {
				title: title,
				autoHideDelay: 5000,
				toaster: "b-toaster-bottom-right",
				variant: type,
				appendToast: false
			})
		},
		submitData() {
			this.$v.form.$touch();
      		if (this.$v.form.$anyError) {
        		return;
			  }
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
			.then((response) => response.text())
			.then((response) => {
				if (response.startsWith("Erro")) {
					this.makeToast("Erro criando conta!", response, "danger");
					this.form.username = '';
				} else {
					this.login();
				}
			})
			this.registerStatus = 0;
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
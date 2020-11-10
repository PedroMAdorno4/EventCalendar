<template>
	<div class="mb-5">
		<b-navbar
			v-if="authenticated"
			toggleable="sm"
			type="dark"
			variant="info"
			id="nav"
		>
			<b-navbar-brand to="/dashboard">
				<img
					src="../assets/logo.png"
					alt="Logo"
					width="30"
					height="30"
				>
				Calendar
			</b-navbar-brand>

			<b-collapse
				id="nav-collapse"
				is-nav
			>
				<b-navbar-nav>
					<b-nav-item
						to="/dashboard"
						:active="$route.path == '/dashboard'"
					>Home</b-nav-item>
					<b-nav-item
						to="/games"
						:active="$route.path.startsWith('/games')"
					>Jogos</b-nav-item>
				</b-navbar-nav>

				<!-- Right aligned nav items -->
				<b-navbar-nav class="ml-auto">

					<b-nav-text v-if="authenticated">Olá,</b-nav-text>
					<b-nav-item-dropdown
						v-if="authenticated"
						right
					>
						<!-- Using 'button-content' slot -->
						<template v-slot:button-content>
							<em>{{ userInfo.Name }}</em>
						</template>
						<!-- <b-dropdown-item-button href="#">Profile</b-dropdown-item-button> -->
						<b-dropdown-item-button @click="signOut">Sign Out</b-dropdown-item-button>
					</b-nav-item-dropdown>
					<b-nav-item
						to="/login"
						:active="$route.path == '/login'"
						v-else
					>Entrar</b-nav-item>
				</b-navbar-nav>
			</b-collapse>
		</b-navbar>

		<b-navbar
			v-else
			type="dark"
			variant="info"
			id="nav"
		>
			<b-navbar-brand to="/">
				<img
					src="../assets/logo.png"
					alt="Logo"
					width="30"
					height="30"
				>
				Calendário
			</b-navbar-brand>

			<b-collapse
				id="nav-collapse"
				is-nav
			>
				<b-navbar-nav>
					<b-nav-item
						to="/"
						:active="$route.path == '/'"
					>Início</b-nav-item>
					<b-nav-item
						to="/desktop"
						:active="$route.path == '/desktop'"
					>LepiDesktop</b-nav-item>
					<b-nav-item
						to="/board"
						:active="$route.path == '/board'"
					>LepiBoard</b-nav-item>
					<b-nav-item
						to="/team"
						:active="$route.path == '/team'"
					>Equipe</b-nav-item>
					<b-nav-item
						to="/publications"
						:active="$route.path == '/publications'"
					>Pulicacoes</b-nav-item>
				</b-navbar-nav>

				<!-- Right aligned nav items -->
				<b-navbar-nav class="ml-auto">
					<b-nav-item
						to="/login"
						:active="$route.path == '/login'"
					>Entrar</b-nav-item>
				</b-navbar-nav>
			</b-collapse>
		</b-navbar>
	</div>
</template>

<script>
import { mapState, mapGetters, mapActions } from "vuex";

export default {
	computed: {
		...mapState(["token", "authenticated"]),
		...mapGetters(["userInfo"]),
	},
	methods: {
		...mapActions(["logout"]),
		signOut() {
			// this.$cookies.remove("token", "", "localhost");
			this.$cookies.remove("token", "/lepi", "lifes.dc.ufscar.br")
			this.logout()
			this.$router.push({
				name: "PublicHome",
				query: { redirect: "/" }
			});
		}
	}
};
</script>

<template>
	<div class="header">
		<b-navbar
			v-if="authenticated"
			toggleable="sm"
			type="dark"
			variant="info"
			id="nav"
		>
			<b-navbar-brand to="/events">
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

				<CalendarHeader/>

				<!-- Right aligned nav items -->
				<b-navbar-nav>

					<b-nav-text v-if="authenticated" class="text-white">Olá,</b-nav-text>
					<b-nav-item-dropdown
						v-if="authenticated"
						right
					>
						<!-- Using 'button-content' slot -->
						<template v-slot:button-content>
							<em class="font-weight-bold text-white">{{ userInfo.Name }}</em>
						</template>
						<!-- <b-dropdown-item-button href="#">Profile</b-dropdown-item-button> -->
						<b-dropdown-item-button @click="signOut">Sair</b-dropdown-item-button>
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
					>Sobre</b-nav-item>
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
import CalendarHeader from '@/components/CalendarHeader.vue';

export default {
	components: {
		CalendarHeader,
	},
	computed: {
		...mapState(["token", "authenticated"]),
		...mapGetters(["userInfo"]),
	},
	methods: {
		...mapActions(["logout"]),
		signOut() {
			this.$cookies.remove("token", "", "localhost");
			this.logout();
			this.$router.push({
				name: "Inicio",
				query: { redirect: "/" },
			});
		},
	},
};
</script>

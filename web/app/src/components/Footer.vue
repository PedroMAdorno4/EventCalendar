<template>
	<div class="mt-5">

		<div
			class="bv-section bv-gray bv-p-reset"
			aria-hidden="true"
		>

		</div>
		<footer :class="['bd-footer', 'text-muted']">
			<b-container>
				<b-row class="justify-content-center">
					<b-col
						cols="12"
						sm
						order-sm="3"
						class="text-center text-sm-right"
					>
						<b-link
							to="/"
							class="d-inline-block mx-auto"
						>
							<img
								src='../assets/logo.png'
								draggable="false"
								ondragstart="return false;"
								class="mt-5"
								width="180"
								height="180"
								alt=' '
							>
						</b-link>
					</b-col>

					<b-col
						cols="auto"
						md="4"
						class="text-left mt-5"
					>
						<h5 class="bd-text-purple-bright mb-1">Documentation</h5>
						<ul class="list-unstyled ml-3">
							<li>
								<b-link
									to="/"
									exact
								>Home</b-link>
							</li>
							<li>
								<b-link
									to="/docs"
									exact
								>Getting started</b-link>
							</li>
							<li>
								<b-link
									to="/docs/components"
									exact
								>Components</b-link>
							</li>
							<li>
								<b-link
									to="/docs/directives"
									exact
								>Directives</b-link>
							</li>
							<li>
								<b-link
									to="/docs/icons"
									exact
								>Icons</b-link>
							</li>
							<li>
								<b-link
									to="/docs/reference"
									exact
								>Reference</b-link>
							</li>
							<li>
								<b-link
									to="/play"
									exact
								>Playground</b-link>
							</li>
						</ul>
						<!-- TODO: Uncomment when we have themes
          <h5 class="bd-text-purple-bright mb-1 mt-3">Themes</h5>
          <ul class="list-unstyled ml-3">
            <li><b-link to="/themes" exact>Themes and dashboards</b-link></li>
          </ul>
          -->
					</b-col>

					<b-col
						cols="auto"
						class="text-left"
					>
						<h5 class="bd-text-purple-bright mb-1 mt-5">Community</h5>
						<ul class="list-unstyled ml-3">
							<li><a
									href="https://github.com/bootstrap-vue/bootstrap-vue"
									target="_blank"
								>GitHub</a></li>
							<li><a
									href="https://discord.gg/j2Mtcny"
									target="_blank"
								>Discord chat</a></li>
							<li><a
									href="https://twitter.com/BootstrapVue"
									target="_blank"
								>Twitter</a></li>
							<li><a
									href="https://opencollective.com/bootstrap-vue/"
									target="_blank"
								>Open Collective</a></li>
						</ul>
						<h5 class="bd-text-purple-bright mb-1 mt-3">Release notes</h5>
						<ul class="list-unstyled ml-3">
							<li>
								<b-link
									to="/docs/reference/changelog"
									exact
								>Changelog</b-link>
							</li>
						</ul>
					</b-col>
				</b-row>

				<hr>

				<p class="mb-2">
					Designed and built with all the love in the world. Maintained by the
					<a
						href="https://github.com/orgs/bootstrap-vue/people"
						target="_blank"
					>core team</a>
					with the help of
					<a
						href="https://github.com/bootstrap-vue/bootstrap-vue/graphs/contributors"
						target="_blank"
					>our contributors</a>.
				</p>

				<p>
					Code licensed
					<a
						href="https://github.com/bootstrap-vue/bootstrap-vue/blob/master/LICENSE"
						target="_blank"
					>MIT</a>.
					Docs generated with
					<a
						href="https://nuxtjs.org/"
						target="_blank"
					>Nuxt.js</a>.
				</p>
			</b-container>
		</footer>
	</div>
</template>

<script>
export default {};
</script>

<style lang="scss" scoped>
.bv-logo {
	margin-top: -2rem;
}
.docs-border {
	border-top: 3px solid #dee2e6;
}
.bd-footer {
	background: #ad9bfd8c;
}
@media (min-width: 1200px) {
	.bd-footer > .container {
		max-width: 960px;
	}
}

$bv-angle-depth: 4rem;
$bv-angle-padding: 3rem;
$bv-angle-depth-md: 4rem;
$bv-angle-padding-md: 6rem;
.bv-section {
	background-size: 100% $bv-angle-depth;
	background-position: center top;
	background-repeat: no-repeat;
	padding-top: calc(#{$bv-angle-depth} + #{$bv-angle-padding});
	padding-bottom: 3rem;
	margin-top: -1px;
	&.bv-white {
		background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 27826 3336' preserveAspectRatio='none' version='1.1'%3E%3Cg%3E%3Cpolygon fill='%23eee' points='13913,3336 0,0 27826,0' /%3E%3C/g%3E%3C/svg%3E");
		background-color: #fff;
	}
	&.bv-gray {
		background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 27826 3336' preserveAspectRatio='none' version='1.1'%3E%3Cg%3E%3Cpolygon fill='%23fff' points='13913,3336 0,0 27826,0'/%3E%3C/g%3E%3C/svg%3E");
		background-color: #ad9bfd8c;
	}
	&.bv-gray-color {
		// background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 27826 3450' preserveAspectRatio='none'%3E%3Cdefs%3E%3Cfilter id='a' filterUnits='objectBoundingBox' x='-50%25' y='-50%25' width='200%25' height='200%25'%3E%3CfeOffset in='SourceGraphic' result='b' dy='20'/%3E%3CfeGaussianBlur in='b' stdDeviation='25'/%3E%3C/filter%3E%3C/defs%3E%3Cpath filter='url(%23a)' d='M13943 3336L5092 1668 0 708V0h27826v719l-5032 949z'/%3E%3Cpath fill='%2342B883' d='M13943 3336L5092 1668 0 708V0h27826v719l-5032 949z'/%3E%3Cpath filter='url(%23a)' d='M0 0h27826v371l-5067 955-8851 1668-8850-1668L0 373z'/%3E%3Cpath fill='%237952B3' d='M0 0h27826v371l-5067 955-8851 1668-8850-1668L0 373z'/%3E%3Cpath filter='url(%23a)' d='M0 0h27826l-5065 954-8851 1668L5059 954 0 1z'/%3E%3Cpath fill='%23fff' d='M0 0h27826l-5065 954-8851 1668L5059 954 0 1z'/%3E%3C/svg%3E%0A");
		background-image: url("data:image/svg+xml;charset=utf-8,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 27826 3450' preserveAspectRatio='none'%3E%3Cdefs%3E%3Cfilter id='a' filterUnits='objectBoundingBox' x='-50%25' y='-50%25' width='200%25' height='200%25'%3E%3CfeOffset in='SourceGraphic' result='b' dy='20'/%3E%3CfeGaussianBlur in='b' stdDeviation='25'/%3E%3C/filter%3E%3C/defs%3E%3Cpath filter='url(%23a)' d='M13943 3336L5092 1668 0 708V0h27826v719l-5032 949z'/%3E%3Cpath fill='%2342B883' d='M13943 3336L5092 1668 0 708V0h27826v719l-5032 949z'/%3E%3Cpath filter='url(%23a)' d='M0 0h27826v371l-5067 955-8851 1668-8850-1668L0 373z'/%3E%3Cpath fill='%237952B3' d='M0 0h27826v371l-5067 955-8851 1668-8850-1668L0 373z'/%3E%3Cpath filter='url(%23a)' d='M0 0h27826l-5065 954-8851 1668L5059 954 0 1z'/%3E%3Cpath fill='%23343a40' d='M0 0h27826l-5065 954-8851 1668L5059 954 0 1z'/%3E%3C/svg%3E");
		background-size: 175% 9rem;
		background-position-y: -2.75rem;
		background-color: #eee;
		padding-top: 9rem;
	}
	&.bv-p-reset {
		padding-top: $bv-angle-depth;
		padding-bottom: 0;
	}
	@media (min-width: 768px) {
		padding-bottom: 2rem;
		padding-top: calc(#{$bv-angle-depth-md} + #{$bv-angle-padding-md});
		&.bv-white,
		&.bv-gray {
			background-size: 100% $bv-angle-depth-md;
		}
		&.bv-gray-color {
			background-size: 100% 12rem;
			background-position-y: top;
			padding-top: 16rem;
		}
		&.bv-p-reset {
			padding-top: $bv-angle-depth-md;
			padding-bottom: 0;
		}
	}
}
</style>
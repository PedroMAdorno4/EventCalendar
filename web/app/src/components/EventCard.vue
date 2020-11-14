<template>
	<div>
		<!-- <b-card
			tag="article"
			style="max-width: 20rem;"
			bg-variant="success"
			text-variant="white"
			class="m-0"
			no-body
		>
			<!-- <template v-slot:header> -->
		<!-- <p class="mx-2 my-1 font-weight-bold" style="font-size: smaller">{{event.title}}
				</p> -->
		<!-- </template> 
		</b-card>  -->

		<b-button
			class="border-0 my-1 w-100"
			size="sm"
			variant="info"
            @click="setEvent"
		>
			<div class="w-100 d-flex">
				<div class="font-weight-bold text-align-center text-nowrap">
					{{hours}}
				</div>
				<div class="ml-2 text-truncate" style="display:block;">
					{{event.title}}
				</div>
			</div>
		</b-button>
	</div>
</template>

<script>
import { mapActions } from 'vuex';
export default {
	props: {
		event: {
			type: Object,
			default: {},
		},
	},
	computed: {
		hours() {
			var start = new Date(this.event.start * 1000);
            var end = new Date(this.event.end * 1000);
            var paddedStart = `${start.getHours().toString().padStart(2, '0')}:${start.getMinutes().toString().padStart(2, '0')}`;
            var paddedEnd = `${end.getHours().toString().padStart(2, '0')}:${end.getMinutes().toString().padStart(2, '0')}`;
			return `${paddedStart} - ${paddedEnd}`;
		},
    },
    methods: {
        ...mapActions(["setCurrentEvent", "setCurrentDate"]),
        setEvent() {
			this.setCurrentEvent(this.event);
			this.setCurrentDate(new Date(this.event.start * 1000));
            this.$bvModal.show('event-modal')
        }
    },
};
</script>

<style lang="scss" scoped>
</style>
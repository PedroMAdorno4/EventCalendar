<template>
	<transition-group name="eventListTransition" class="list">
		<!-- <div class="list"> -->
			<EventCard
				v-for="i in dayEvents"
				:key="i.start + i.end"
				:event="i"
			/>
		<!-- </div> -->
	</transition-group>
</template>

<script>
import { mapState } from "vuex";
import EventCard from "@/components/EventCard.vue";

export default {
	components: {
		EventCard,
	},
	props: {
		day: {
			type: Number,
			default: 1,
		},
	},
	data() {
		return {
			dayStart: Number,
			dayEnd: Number,
		};
	},
	computed: {
		...mapState(["events", "currentDate"]),
		dayEvents() {
			if (this.events === null) {
				return [];
			}
			var events = [];
			var time = new Date(
				this.currentDate.getFullYear(),
				this.currentDate.getMonth(),
				this.day,
				0,
				0
			);
			this.dayStart = time.getTime() / 1000;

			time = new Date(
				this.currentDate.getFullYear(),
				this.currentDate.getMonth(),
				this.day,
				23,
				59
			);
			this.dayEnd = time.getTime() / 1000;

			this.events.forEach((event) => {
				if (
					event["start"] >= this.dayStart &&
					event["start"] <= this.dayEnd &&
					event["end"] >= this.dayStart &&
					event["end"] <= this.dayEnd
				) {
					events.push(event);
				}
			});

			//sort is stable, sort from least important to most important
			events.sort((a, b) => (a.end - b.end)).sort((a, b) => (a.start - b.start));
			return events;
		},
	},
};
</script>

<style lang="scss" scoped>
.list {
	overflow-y: hidden;
	overflow-x: hidden;

	//Space for scrollbar
	padding-right: 0.75rem;

	&:hover {
		overflow-y: visible;
		overflow-x: hidden;
		scrollbar-width: thin;

		// padding-right: thin;
	}
}

.eventListTransition-enter-active,
.eventListTransition-leave-active {
	transition: all 1s;
}
.eventListTransition-enter, .eventListTransition-leave-to {
	opacity: 0;
	transform: translateY(30px);
}
</style>
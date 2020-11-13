<template>
	<div class="list pr-1">
		<EventCard
			v-for="i in dayEvents"
			:key="i.start"
			:event="i"
		/>
	</div>
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
			return events;
		},
	},
};
</script>

<style lang="scss" scoped>
.list {
	overflow-y: auto;
}
</style>
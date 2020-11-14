<template>
	<div class="h-100 w-100" style="overflow: hidden;">
		<EventModal />
		<transition name="calendarTransition" style="overflow: hidden;">
			<b-table
				class="h-100 mb-0 border-2"
				bordered
				outlined
				fixed
				block
				:key="`${currentDate.getMonth()}-${currentDate.getFullYear()}`"
				:fields="daysOfTheWeek"
				:items="days"
				tbody-class="events"
				thead-class="thead"
			>
				<template #cell()="data">
					<div class="d-flex flex-column h-100">
						<div class="w-100">
							<b-button
								:class="dayLabelClass(parseInt(data.value))"
								size="sm"
								:variant="dayLabelVariant(data.value)"
								@click="setCurrentDay(parseInt(data.value))"
							>
								{{data.value}}
							</b-button>
						</div>
						<EventList :day="parseInt(data.value)" />
					</div>
				</template>
			</b-table>
		</transition>
	</div>
</template>

<script>
import { mapActions, mapGetters, mapState } from "vuex";
import { addDays, addMonths, addYears } from "date-fns";

import EventList from "@/components/EventList.vue";
import EventModal from "@/components/EventModal.vue";

export default {
	components: {
		EventList,
		EventModal,
	},
	data() {
		return {
			daysOfTheWeek: [
				"Domingo",
				"Segunda",
				"Terça",
				"Quarta",
				"Quinta",
				"Sexta",
				"Sábado",
			],
		};
	},
	methods: {
		...mapActions(["setCurrentDate", "setEvents", "setCurrentEvent"]),
		monthDays(year, month) {
			return new Date(year, month + 1, 0).getDate();
		},
		dayOfTheWeek(year, month, day) {
			return new Date(year, month, day).getDay();
		},
		dayLabelVariant(day) {
			var today = new Date();
			if (
				today.getDate() === day &&
				this.currentDate.getMonth() === today.getMonth() &&
				this.currentDate.getFullYear() === today.getFullYear()
			) {
				return "primary";
			} else {
				return "outline-primary";
			}
		},
		dayLabelClass(day) {
			var labelClass = "border-0 mb-2 ";
			var today = new Date();
			if (
				today.getDate() === day &&
				this.currentDate.getMonth() === today.getMonth() &&
				this.currentDate.getFullYear() === today.getFullYear()
			) {
				return labelClass.concat("text-white font-weight-bold");
			} else {
				return labelClass;
			}
		},
		setCurrentDay(day) {
			var date = this.currentDate;
			date.setDate(day);
			this.setCurrentDate(date);
			this.setCurrentEvent({});
			this.$bvModal.show("event-modal");
		},
		getEvents() {
			fetch("/backend/API/event/getByOwner", {
				method: "GET",
				headers: {
					"Content-Type": "application/json",
					Authorization: "Bearer: " + this.token,
					_id: JSON.stringify(this.userInfo.ID),
				},
			})
				.then((response) => {
					return response.json();
				})
				.then((data) => {
					this.setEvents(data);
				});
		},
	},
	computed: {
		...mapGetters(["userInfo"]),
		...mapState(["currentDate", "token"]),
		days() {
			var weeks = [];

			var week = {};
			var monthLength =
				this.monthDays(
					this.currentDate.getFullYear(),
					this.currentDate.getMonth()
				) + 1;
			for (let d = 1; d < monthLength; d++) {
				var weekDay = this.dayOfTheWeek(
					this.currentDate.getFullYear(),
					this.currentDate.getMonth(),
					d
				);
				week[this.daysOfTheWeek[weekDay]] = d;

				if (
					weekDay === this.daysOfTheWeek.length - 1 ||
					d === monthLength - 1
				) {
					weeks.push(week);
					week = {};
				}
			}
			return weeks;
		},
	},
	mounted() {
		this.getEvents();
	},
};
</script>

<style lang="scss">
.calendarTransition-enter-active,
.calendarTransition-enter-active {
	transition: all 0.15s ease-in-out;
}
.calendarTransition-enter {
	opacity: 0;
	transform: translatey(1vw);
}
.calendarTransition-enter-to {
	opacity: 1;
	//   transform: translatex(100vw);
}

table {
	display: flex;
	flex-direction: column;
	width: 100%;
	border: 2px solid #dee2e6 !important;
}
table thead tr {
	display: flex;

	th {
		width: 100vw;
		padding-bottom: 0.75rem;
	}
}

table tbody {
	display: flex;
	flex-direction: column;
	width: 100%;

	tr {
		display: flex;
		height: 100%;
		max-height: 160px;

		td {
			width: 100vw;
			height: 100%;
			flex: 0 1 auto;
			overflow: hidden;
		}
	}
}

.table td {
	padding: 0.75rem 0rem 0rem 0.75rem;
}
</style>
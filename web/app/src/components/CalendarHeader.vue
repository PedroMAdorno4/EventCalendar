<template>
<div class="d-flex flex-fill align-items-center mr-auto">
    <b-button
			class="border-0 font-weight-bold"
            variant="outline-light"
		>
            Novo Evento
        </b-button>

	<div class="mx-auto d-flex align-items-center">
        <b-button
			class="border-0 mr-2"
            size="sm"
            variant="outline-light"
            @click="previousYear"
		>
            <b-icon icon="chevron-double-left"></b-icon>
        </b-button>
		<b-button
			class="border-0"
            size="sm"
            variant="outline-light"
            @click="previousMonth"
		>
            <b-icon icon="chevron-left"></b-icon>
        </b-button>

        <div class="mx-3 my-0 p-0 font-weight-bold text-white" style="width: 10em">{{this.strCurrentDate}}</div>

        <b-button
            class="border-0"
            size="sm"
            variant="outline-light"
            @click="nextMonth"
        >
            <b-icon icon="chevron-right"></b-icon>
        </b-button>
        <b-button
            class="border-0 ml-2"
            size="sm"
            variant="outline-light"
            @click="nextYear"
        >
            <b-icon icon="chevron-double-right"></b-icon>
        </b-button>
	</div>
</div>
</template>

<script>
import { mapActions, mapGetters, mapState } from 'vuex';
import { addMonths, subMonths, addYears, subYears } from 'date-fns';

export default {
    computed: {
        ...mapGetters(["currentDate", "strCurrentDate"]),
        ...mapState(["currentDate"])
    },
    methods: {
        ...mapActions(["setCurrentDate"]),
        nextMonth() {
            this.setCurrentDate(addMonths(this.currentDate, 1));
        },
        previousMonth() {
            this.setCurrentDate(subMonths(this.currentDate, 1));
        },
        nextYear() {
            this.setCurrentDate(addYears(this.currentDate, 1));
        },
        previousYear() {
            this.setCurrentDate(subYears(this.currentDate, 1));
        }
    },
};
</script>

<style lang="scss" scoped>
</style>
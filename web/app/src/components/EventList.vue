<template>
    <div>

    </div>
</template>

<script>
import { mapState } from 'vuex';
import EventCard from '@/components/EventCard.vue';

export default {
    components: {
        EventCard,
    },
    props: {
        day: {
            type: Number,
            default: 1
        },
    },
    data() {
        return {
            dayStart: Number,
            dayEnd: Number
        }
    },
    computed: {
        ...mapState(["events", "currentDate"]),
        // dayEvents() {
        //     events = [];
        //     var time = new Date(this.currentDate.getFullYear(), this.currentDate.getMonth(), this.day, 0, 0);
        //     this.dayStart = time.getTime() / 1000;

        //     time = new Date(this.currentDate.getFullYear(), this.currentDate.getMonth(), this.day, 23, 59);
        //     this.dayEnd = time.getTime() / 1000;

        //     this.events.forEach(event => {
        //         if(event["start"] >= this.dayStart && event["start"] <= this.dayEnd) {

        //         }
        //     });
        // }
    },
}
</script>

<style lang="scss" scoped>

</style>
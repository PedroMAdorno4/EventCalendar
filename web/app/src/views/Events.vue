<template>
    <div class="d-flex events">
        <Calendar/>
        
    </div>
</template>

<script>
import Calendar from '@/components/Calendar.vue';


export default {
    components: {
        Calendar,
    },
}
</script>

<style lang="scss">
.events {
    flex: 1 0 auto;
}
</style>
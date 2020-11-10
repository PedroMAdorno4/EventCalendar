<template>
    <b-container fluid>
        <b-row class="d-flex justify-content-around">
            <b-col cols="3">
			<Register></Register>
            </b-col>
			<b-col cols="3">
			<Login></Login>
            </b-col>
        </b-row>
    </b-container>
    </div>
</template>

<script>
import Login from '@/components/Login.vue';
import Register from '@/components/Register.vue';

export default {
    components: {
        Login,
        Register,
    }
}
</script>

<style lang="scss" scoped>

</style>
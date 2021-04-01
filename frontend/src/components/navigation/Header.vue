<template>
    <b-navbar type="is-primary" fixed-top="true">
        <template #brand>
            <b-navbar-item tag="router-link" :to="{ path: '/' }">
                SaN
            </b-navbar-item>
        </template>
        <template #start>
            <b-navbar-item href="#">
                Home
            </b-navbar-item>
<!--            <b-navbar-item href="#">-->
<!--                Documentation-->
<!--            </b-navbar-item>-->
<!--            <b-navbar-dropdown label="Info">-->
<!--                <b-navbar-item href="#">-->
<!--                    About-->
<!--                </b-navbar-item>-->
<!--                <b-navbar-item href="#">-->
<!--                    Contact-->
<!--                </b-navbar-item>-->
<!--            </b-navbar-dropdown>-->
        </template>
    </b-navbar>
</template>

<script>
    export default {
        name: "Header"
    }
</script>

<style scoped>

</style>

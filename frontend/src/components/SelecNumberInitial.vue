<template>
    <div>
        <button
                class="btn btn-primary"
                @click="confirmNumberSelected"
        >
            I have picked a number!
        </button>
    </div>
</template>

<script>
    export default {
        name: "SelecNumberInitial",
        methods: {
            confirmNumberSelected() {
                console.log("Button")
            }
        }
    }
</script>

<style scoped>

</style>

<template>
    <div class="number-option-box">
        <div class="number-option">
            {{ number }}
        </div>
    </div>
</template>

<script>
    export default {
        name: "NumberOptionBox",
        props: ["number"],
    }
</script>

<style scoped>
    .number-option-box {
        flex: 1 0 18%;
        margin: 5px;
        padding: 5px;
        border: 1px solid;
        border-radius: 5px;
    }

    .number-option {
        font-size: 1.5rem;
    }
</style>

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

<style scoped lang="scss">
    .number-option-box {
        flex: 1 0 15%;
        margin: 20px 30px;
        padding: 5px;
        border: 1px solid;
        border-radius: 5px;
        .number-option {
            font-size: 1.5rem
        }
    }
</style>
